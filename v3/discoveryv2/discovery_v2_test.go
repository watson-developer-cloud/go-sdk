/**
 * (C) Copyright IBM Corp. 2019, 2022.
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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v3/discoveryv2"
)

var _ = Describe(`DiscoveryV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(discoveryService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(discoveryService.Service.IsSSLDisabled()).To(BeFalse())
			discoveryService.DisableSSLVerification()
			Expect(discoveryService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(discoveryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				URL:     "https://discoveryv2/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(discoveryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{})
			Expect(discoveryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISCOVERY_URL":       "https://discoveryv2/api",
				"DISCOVERY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					Version: core.StringPtr(version),
				})
				Expect(discoveryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := discoveryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != discoveryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(discoveryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(discoveryService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(discoveryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := discoveryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != discoveryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(discoveryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(discoveryService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					Version: core.StringPtr(version),
				})
				err := discoveryService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := discoveryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != discoveryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(discoveryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(discoveryService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISCOVERY_URL":       "https://discoveryv2/api",
				"DISCOVERY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(discoveryService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISCOVERY_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(discoveryService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = discoveryv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions) - Operation response error`, func() {
		version := "testString"
		listCollectionsPath := "/v2/projects/testString/collections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCollections with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv2.ListCollectionsOptions)
				listCollectionsOptionsModel.ProjectID = core.StringPtr("testString")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions)`, func() {
		version := "testString"
		listCollectionsPath := "/v2/projects/testString/collections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"collection_id": "CollectionID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke ListCollections successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv2.ListCollectionsOptions)
				listCollectionsOptionsModel.ProjectID = core.StringPtr("testString")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"collection_id": "CollectionID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv2.ListCollectionsOptions)
				listCollectionsOptionsModel.ProjectID = core.StringPtr("testString")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCollections with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv2.ListCollectionsOptions)
				listCollectionsOptionsModel.ProjectID = core.StringPtr("testString")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCollectionsOptions model with no property values
				listCollectionsOptionsModelNew := new(discoveryv2.ListCollectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListCollections(listCollectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv2.ListCollectionsOptions)
				listCollectionsOptionsModel.ProjectID = core.StringPtr("testString")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions) - Operation response error`, func() {
		version := "testString"
		createCollectionPath := "/v2/projects/testString/collections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCollection with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv2.CreateCollectionOptions)
				createCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
				createCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
		version := "testString"
		createCollectionPath := "/v2/projects/testString/collections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "language": "en", "enrichments": [{"enrichment_id": "EnrichmentID", "fields": ["Fields"]}]}`)
				}))
			})
			It(`Invoke CreateCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv2.CreateCollectionOptions)
				createCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
				createCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "language": "en", "enrichments": [{"enrichment_id": "EnrichmentID", "fields": ["Fields"]}]}`)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv2.CreateCollectionOptions)
				createCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
				createCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv2.CreateCollectionOptions)
				createCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
				createCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCollectionOptions model with no property values
				createCollectionOptionsModelNew := new(discoveryv2.CreateCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateCollection(createCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv2.CreateCollectionOptions)
				createCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
				createCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions) - Operation response error`, func() {
		version := "testString"
		getCollectionPath := "/v2/projects/testString/collections/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCollection with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv2.GetCollectionOptions)
				getCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
		version := "testString"
		getCollectionPath := "/v2/projects/testString/collections/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "language": "en", "enrichments": [{"enrichment_id": "EnrichmentID", "fields": ["Fields"]}]}`)
				}))
			})
			It(`Invoke GetCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv2.GetCollectionOptions)
				getCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "language": "en", "enrichments": [{"enrichment_id": "EnrichmentID", "fields": ["Fields"]}]}`)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv2.GetCollectionOptions)
				getCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv2.GetCollectionOptions)
				getCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCollectionOptions model with no property values
				getCollectionOptionsModelNew := new(discoveryv2.GetCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetCollection(getCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv2.GetCollectionOptions)
				getCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) - Operation response error`, func() {
		version := "testString"
		updateCollectionPath := "/v2/projects/testString/collections/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCollection with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv2.UpdateCollectionOptions)
				updateCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
		version := "testString"
		updateCollectionPath := "/v2/projects/testString/collections/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "language": "en", "enrichments": [{"enrichment_id": "EnrichmentID", "fields": ["Fields"]}]}`)
				}))
			})
			It(`Invoke UpdateCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv2.UpdateCollectionOptions)
				updateCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "language": "en", "enrichments": [{"enrichment_id": "EnrichmentID", "fields": ["Fields"]}]}`)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv2.UpdateCollectionOptions)
				updateCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv2.UpdateCollectionOptions)
				updateCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCollectionOptions model with no property values
				updateCollectionOptionsModelNew := new(discoveryv2.UpdateCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateCollection(updateCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv2.UpdateCollectionOptions)
				updateCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Enrichments = []discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		version := "testString"
		deleteCollectionPath := "/v2/projects/testString/collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv2.DeleteCollectionOptions)
				deleteCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv2.DeleteCollectionOptions)
				deleteCollectionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCollectionOptions model with no property values
				deleteCollectionOptionsModelNew := new(discoveryv2.DeleteCollectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteCollection(deleteCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Query(queryOptions *QueryOptions) - Operation response error`, func() {
		version := "testString"
		queryPath := "/v2/projects/testString/query"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Query with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryLargeTableResults model
				queryLargeTableResultsModel := new(discoveryv2.QueryLargeTableResults)
				queryLargeTableResultsModel.Enabled = core.BoolPtr(true)
				queryLargeTableResultsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryLargeSuggestedRefinements model
				queryLargeSuggestedRefinementsModel := new(discoveryv2.QueryLargeSuggestedRefinements)
				queryLargeSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				queryLargeSuggestedRefinementsModel.Count = core.Int64Ptr(int64(1))

				// Construct an instance of the QueryLargePassages model
				queryLargePassagesModel := new(discoveryv2.QueryLargePassages)
				queryLargePassagesModel.Enabled = core.BoolPtr(true)
				queryLargePassagesModel.PerDocument = core.BoolPtr(true)
				queryLargePassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				queryLargePassagesModel.Fields = []string{"testString"}
				queryLargePassagesModel.Count = core.Int64Ptr(int64(400))
				queryLargePassagesModel.Characters = core.Int64Ptr(int64(50))
				queryLargePassagesModel.FindAnswers = core.BoolPtr(false)
				queryLargePassagesModel.MaxAnswersPerPassage = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv2.QueryOptions)
				queryOptionsModel.ProjectID = core.StringPtr("testString")
				queryOptionsModel.CollectionIds = []string{"testString"}
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = []string{"testString"}
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.TableResults = queryLargeTableResultsModel
				queryOptionsModel.SuggestedRefinements = queryLargeSuggestedRefinementsModel
				queryOptionsModel.Passages = queryLargePassagesModel
				queryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.Query(queryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.Query(queryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Query(queryOptions *QueryOptions)`, func() {
		version := "testString"
		queryPath := "/v2/projects/testString/query"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"document_id": "DocumentID", "metadata": {"mapKey": "anyValue"}, "result_metadata": {"document_retrieval_source": "search", "collection_id": "CollectionID", "confidence": 10}, "document_passages": [{"passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field", "confidence": 0, "answers": [{"answer_text": "AnswerText", "start_offset": 11, "end_offset": 9, "confidence": 0}]}]}], "aggregations": [{"type": "filter", "match": "Match", "matching_results": 15}], "retrieval_details": {"document_retrieval_strategy": "untrained"}, "suggested_query": "SuggestedQuery", "suggested_refinements": [{"text": "Text"}], "table_results": [{"table_id": "TableID", "source_document_id": "SourceDocumentID", "collection_id": "CollectionID", "table_html": "TableHTML", "table_html_offset": 15, "table": {"location": {"begin": 5, "end": 3}, "text": "Text", "section_title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "table_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "row_headers": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "column_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "key_value_pairs": [{"key": {"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}, "value": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}]}], "body_cells": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14, "row_header_ids": [{"id": "ID"}], "row_header_texts": [{"text": "Text"}], "row_header_texts_normalized": [{"text_normalized": "TextNormalized"}], "column_header_ids": [{"id": "ID"}], "column_header_texts": [{"text": "Text"}], "column_header_texts_normalized": [{"text_normalized": "TextNormalized"}], "attributes": [{"type": "Type", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "contexts": [{"text": "Text", "location": {"begin": 5, "end": 3}}]}}], "passages": [{"passage_text": "PassageText", "passage_score": 12, "document_id": "DocumentID", "collection_id": "CollectionID", "start_offset": 11, "end_offset": 9, "field": "Field", "confidence": 0, "answers": [{"answer_text": "AnswerText", "start_offset": 11, "end_offset": 9, "confidence": 0}]}]}`)
				}))
			})
			It(`Invoke Query successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the QueryLargeTableResults model
				queryLargeTableResultsModel := new(discoveryv2.QueryLargeTableResults)
				queryLargeTableResultsModel.Enabled = core.BoolPtr(true)
				queryLargeTableResultsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryLargeSuggestedRefinements model
				queryLargeSuggestedRefinementsModel := new(discoveryv2.QueryLargeSuggestedRefinements)
				queryLargeSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				queryLargeSuggestedRefinementsModel.Count = core.Int64Ptr(int64(1))

				// Construct an instance of the QueryLargePassages model
				queryLargePassagesModel := new(discoveryv2.QueryLargePassages)
				queryLargePassagesModel.Enabled = core.BoolPtr(true)
				queryLargePassagesModel.PerDocument = core.BoolPtr(true)
				queryLargePassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				queryLargePassagesModel.Fields = []string{"testString"}
				queryLargePassagesModel.Count = core.Int64Ptr(int64(400))
				queryLargePassagesModel.Characters = core.Int64Ptr(int64(50))
				queryLargePassagesModel.FindAnswers = core.BoolPtr(false)
				queryLargePassagesModel.MaxAnswersPerPassage = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv2.QueryOptions)
				queryOptionsModel.ProjectID = core.StringPtr("testString")
				queryOptionsModel.CollectionIds = []string{"testString"}
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = []string{"testString"}
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.TableResults = queryLargeTableResultsModel
				queryOptionsModel.SuggestedRefinements = queryLargeSuggestedRefinementsModel
				queryOptionsModel.Passages = queryLargePassagesModel
				queryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.QueryWithContext(ctx, queryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.Query(queryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.QueryWithContext(ctx, queryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"document_id": "DocumentID", "metadata": {"mapKey": "anyValue"}, "result_metadata": {"document_retrieval_source": "search", "collection_id": "CollectionID", "confidence": 10}, "document_passages": [{"passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field", "confidence": 0, "answers": [{"answer_text": "AnswerText", "start_offset": 11, "end_offset": 9, "confidence": 0}]}]}], "aggregations": [{"type": "filter", "match": "Match", "matching_results": 15}], "retrieval_details": {"document_retrieval_strategy": "untrained"}, "suggested_query": "SuggestedQuery", "suggested_refinements": [{"text": "Text"}], "table_results": [{"table_id": "TableID", "source_document_id": "SourceDocumentID", "collection_id": "CollectionID", "table_html": "TableHTML", "table_html_offset": 15, "table": {"location": {"begin": 5, "end": 3}, "text": "Text", "section_title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "title": {"text": "Text", "location": {"begin": 5, "end": 3}}, "table_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "row_headers": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "column_headers": [{"cell_id": "CellID", "location": {"anyKey": "anyValue"}, "text": "Text", "text_normalized": "TextNormalized", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14}], "key_value_pairs": [{"key": {"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}, "value": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text"}]}], "body_cells": [{"cell_id": "CellID", "location": {"begin": 5, "end": 3}, "text": "Text", "row_index_begin": 13, "row_index_end": 11, "column_index_begin": 16, "column_index_end": 14, "row_header_ids": [{"id": "ID"}], "row_header_texts": [{"text": "Text"}], "row_header_texts_normalized": [{"text_normalized": "TextNormalized"}], "column_header_ids": [{"id": "ID"}], "column_header_texts": [{"text": "Text"}], "column_header_texts_normalized": [{"text_normalized": "TextNormalized"}], "attributes": [{"type": "Type", "text": "Text", "location": {"begin": 5, "end": 3}}]}], "contexts": [{"text": "Text", "location": {"begin": 5, "end": 3}}]}}], "passages": [{"passage_text": "PassageText", "passage_score": 12, "document_id": "DocumentID", "collection_id": "CollectionID", "start_offset": 11, "end_offset": 9, "field": "Field", "confidence": 0, "answers": [{"answer_text": "AnswerText", "start_offset": 11, "end_offset": 9, "confidence": 0}]}]}`)
				}))
			})
			It(`Invoke Query successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.Query(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the QueryLargeTableResults model
				queryLargeTableResultsModel := new(discoveryv2.QueryLargeTableResults)
				queryLargeTableResultsModel.Enabled = core.BoolPtr(true)
				queryLargeTableResultsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryLargeSuggestedRefinements model
				queryLargeSuggestedRefinementsModel := new(discoveryv2.QueryLargeSuggestedRefinements)
				queryLargeSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				queryLargeSuggestedRefinementsModel.Count = core.Int64Ptr(int64(1))

				// Construct an instance of the QueryLargePassages model
				queryLargePassagesModel := new(discoveryv2.QueryLargePassages)
				queryLargePassagesModel.Enabled = core.BoolPtr(true)
				queryLargePassagesModel.PerDocument = core.BoolPtr(true)
				queryLargePassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				queryLargePassagesModel.Fields = []string{"testString"}
				queryLargePassagesModel.Count = core.Int64Ptr(int64(400))
				queryLargePassagesModel.Characters = core.Int64Ptr(int64(50))
				queryLargePassagesModel.FindAnswers = core.BoolPtr(false)
				queryLargePassagesModel.MaxAnswersPerPassage = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv2.QueryOptions)
				queryOptionsModel.ProjectID = core.StringPtr("testString")
				queryOptionsModel.CollectionIds = []string{"testString"}
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = []string{"testString"}
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.TableResults = queryLargeTableResultsModel
				queryOptionsModel.SuggestedRefinements = queryLargeSuggestedRefinementsModel
				queryOptionsModel.Passages = queryLargePassagesModel
				queryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.Query(queryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Query with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryLargeTableResults model
				queryLargeTableResultsModel := new(discoveryv2.QueryLargeTableResults)
				queryLargeTableResultsModel.Enabled = core.BoolPtr(true)
				queryLargeTableResultsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryLargeSuggestedRefinements model
				queryLargeSuggestedRefinementsModel := new(discoveryv2.QueryLargeSuggestedRefinements)
				queryLargeSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				queryLargeSuggestedRefinementsModel.Count = core.Int64Ptr(int64(1))

				// Construct an instance of the QueryLargePassages model
				queryLargePassagesModel := new(discoveryv2.QueryLargePassages)
				queryLargePassagesModel.Enabled = core.BoolPtr(true)
				queryLargePassagesModel.PerDocument = core.BoolPtr(true)
				queryLargePassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				queryLargePassagesModel.Fields = []string{"testString"}
				queryLargePassagesModel.Count = core.Int64Ptr(int64(400))
				queryLargePassagesModel.Characters = core.Int64Ptr(int64(50))
				queryLargePassagesModel.FindAnswers = core.BoolPtr(false)
				queryLargePassagesModel.MaxAnswersPerPassage = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv2.QueryOptions)
				queryOptionsModel.ProjectID = core.StringPtr("testString")
				queryOptionsModel.CollectionIds = []string{"testString"}
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = []string{"testString"}
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.TableResults = queryLargeTableResultsModel
				queryOptionsModel.SuggestedRefinements = queryLargeSuggestedRefinementsModel
				queryOptionsModel.Passages = queryLargePassagesModel
				queryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.Query(queryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the QueryOptions model with no property values
				queryOptionsModelNew := new(discoveryv2.QueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.Query(queryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke Query successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryLargeTableResults model
				queryLargeTableResultsModel := new(discoveryv2.QueryLargeTableResults)
				queryLargeTableResultsModel.Enabled = core.BoolPtr(true)
				queryLargeTableResultsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryLargeSuggestedRefinements model
				queryLargeSuggestedRefinementsModel := new(discoveryv2.QueryLargeSuggestedRefinements)
				queryLargeSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				queryLargeSuggestedRefinementsModel.Count = core.Int64Ptr(int64(1))

				// Construct an instance of the QueryLargePassages model
				queryLargePassagesModel := new(discoveryv2.QueryLargePassages)
				queryLargePassagesModel.Enabled = core.BoolPtr(true)
				queryLargePassagesModel.PerDocument = core.BoolPtr(true)
				queryLargePassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				queryLargePassagesModel.Fields = []string{"testString"}
				queryLargePassagesModel.Count = core.Int64Ptr(int64(400))
				queryLargePassagesModel.Characters = core.Int64Ptr(int64(50))
				queryLargePassagesModel.FindAnswers = core.BoolPtr(false)
				queryLargePassagesModel.MaxAnswersPerPassage = core.Int64Ptr(int64(38))

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv2.QueryOptions)
				queryOptionsModel.ProjectID = core.StringPtr("testString")
				queryOptionsModel.CollectionIds = []string{"testString"}
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = []string{"testString"}
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.TableResults = queryLargeTableResultsModel
				queryOptionsModel.SuggestedRefinements = queryLargeSuggestedRefinementsModel
				queryOptionsModel.Passages = queryLargePassagesModel
				queryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.Query(queryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions) - Operation response error`, func() {
		version := "testString"
		getAutocompletionPath := "/v2/projects/testString/autocompletion"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutocompletionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["field"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAutocompletion with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv2.GetAutocompletionOptions)
				getAutocompletionOptionsModel.ProjectID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionIds = []string{"testString"}
				getAutocompletionOptionsModel.Field = core.StringPtr("testString")
				getAutocompletionOptionsModel.Count = core.Int64Ptr(int64(38))
				getAutocompletionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetAutocompletion(getAutocompletionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetAutocompletion(getAutocompletionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAutocompletion(getAutocompletionOptions *GetAutocompletionOptions)`, func() {
		version := "testString"
		getAutocompletionPath := "/v2/projects/testString/autocompletion"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutocompletionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["field"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"completions": ["Completions"]}`)
				}))
			})
			It(`Invoke GetAutocompletion successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv2.GetAutocompletionOptions)
				getAutocompletionOptionsModel.ProjectID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionIds = []string{"testString"}
				getAutocompletionOptionsModel.Field = core.StringPtr("testString")
				getAutocompletionOptionsModel.Count = core.Int64Ptr(int64(38))
				getAutocompletionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetAutocompletionWithContext(ctx, getAutocompletionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetAutocompletion(getAutocompletionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetAutocompletionWithContext(ctx, getAutocompletionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutocompletionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["field"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"completions": ["Completions"]}`)
				}))
			})
			It(`Invoke GetAutocompletion successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetAutocompletion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv2.GetAutocompletionOptions)
				getAutocompletionOptionsModel.ProjectID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionIds = []string{"testString"}
				getAutocompletionOptionsModel.Field = core.StringPtr("testString")
				getAutocompletionOptionsModel.Count = core.Int64Ptr(int64(38))
				getAutocompletionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetAutocompletion(getAutocompletionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAutocompletion with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv2.GetAutocompletionOptions)
				getAutocompletionOptionsModel.ProjectID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionIds = []string{"testString"}
				getAutocompletionOptionsModel.Field = core.StringPtr("testString")
				getAutocompletionOptionsModel.Count = core.Int64Ptr(int64(38))
				getAutocompletionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetAutocompletion(getAutocompletionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAutocompletionOptions model with no property values
				getAutocompletionOptionsModelNew := new(discoveryv2.GetAutocompletionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetAutocompletion(getAutocompletionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAutocompletion successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv2.GetAutocompletionOptions)
				getAutocompletionOptionsModel.ProjectID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionIds = []string{"testString"}
				getAutocompletionOptionsModel.Field = core.StringPtr("testString")
				getAutocompletionOptionsModel.Count = core.Int64Ptr(int64(38))
				getAutocompletionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetAutocompletion(getAutocompletionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`QueryCollectionNotices(queryCollectionNoticesOptions *QueryCollectionNoticesOptions) - Operation response error`, func() {
		version := "testString"
		queryCollectionNoticesPath := "/v2/projects/testString/collections/testString/notices"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryCollectionNoticesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke QueryCollectionNotices with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryCollectionNoticesOptions model
				queryCollectionNoticesOptionsModel := new(discoveryv2.QueryCollectionNoticesOptions)
				queryCollectionNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Query = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`QueryCollectionNotices(queryCollectionNoticesOptions *QueryCollectionNoticesOptions)`, func() {
		version := "testString"
		queryCollectionNoticesPath := "/v2/projects/testString/collections/testString/notices"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryCollectionNoticesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "collection_id": "CollectionID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke QueryCollectionNotices successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the QueryCollectionNoticesOptions model
				queryCollectionNoticesOptionsModel := new(discoveryv2.QueryCollectionNoticesOptions)
				queryCollectionNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Query = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.QueryCollectionNoticesWithContext(ctx, queryCollectionNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.QueryCollectionNoticesWithContext(ctx, queryCollectionNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryCollectionNoticesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "collection_id": "CollectionID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke QueryCollectionNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.QueryCollectionNotices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the QueryCollectionNoticesOptions model
				queryCollectionNoticesOptionsModel := new(discoveryv2.QueryCollectionNoticesOptions)
				queryCollectionNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Query = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke QueryCollectionNotices with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryCollectionNoticesOptions model
				queryCollectionNoticesOptionsModel := new(discoveryv2.QueryCollectionNoticesOptions)
				queryCollectionNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Query = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the QueryCollectionNoticesOptions model with no property values
				queryCollectionNoticesOptionsModelNew := new(discoveryv2.QueryCollectionNoticesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke QueryCollectionNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryCollectionNoticesOptions model
				queryCollectionNoticesOptionsModel := new(discoveryv2.QueryCollectionNoticesOptions)
				queryCollectionNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Query = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryCollectionNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryCollectionNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.QueryCollectionNotices(queryCollectionNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`QueryNotices(queryNoticesOptions *QueryNoticesOptions) - Operation response error`, func() {
		version := "testString"
		queryNoticesPath := "/v2/projects/testString/notices"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryNoticesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke QueryNotices with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv2.QueryNoticesOptions)
				queryNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`QueryNotices(queryNoticesOptions *QueryNoticesOptions)`, func() {
		version := "testString"
		queryNoticesPath := "/v2/projects/testString/notices"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryNoticesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "collection_id": "CollectionID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke QueryNotices successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv2.QueryNoticesOptions)
				queryNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.QueryNoticesWithContext(ctx, queryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.QueryNoticesWithContext(ctx, queryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryNoticesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "collection_id": "CollectionID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke QueryNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.QueryNotices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv2.QueryNoticesOptions)
				queryNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke QueryNotices with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv2.QueryNoticesOptions)
				queryNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the QueryNoticesOptions model with no property values
				queryNoticesOptionsModelNew := new(discoveryv2.QueryNoticesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.QueryNotices(queryNoticesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke QueryNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv2.QueryNoticesOptions)
				queryNoticesOptionsModel.ProjectID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFields(listFieldsOptions *ListFieldsOptions) - Operation response error`, func() {
		version := "testString"
		listFieldsPath := "/v2/projects/testString/fields"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListFields with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv2.ListFieldsOptions)
				listFieldsOptionsModel.ProjectID = core.StringPtr("testString")
				listFieldsOptionsModel.CollectionIds = []string{"testString"}
				listFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFields(listFieldsOptions *ListFieldsOptions)`, func() {
		version := "testString"
		listFieldsPath := "/v2/projects/testString/fields"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"fields": [{"field": "Field", "type": "nested", "collection_id": "CollectionID"}]}`)
				}))
			})
			It(`Invoke ListFields successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv2.ListFieldsOptions)
				listFieldsOptionsModel.ProjectID = core.StringPtr("testString")
				listFieldsOptionsModel.CollectionIds = []string{"testString"}
				listFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListFieldsWithContext(ctx, listFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListFieldsWithContext(ctx, listFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFieldsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"fields": [{"field": "Field", "type": "nested", "collection_id": "CollectionID"}]}`)
				}))
			})
			It(`Invoke ListFields successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListFields(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv2.ListFieldsOptions)
				listFieldsOptionsModel.ProjectID = core.StringPtr("testString")
				listFieldsOptionsModel.CollectionIds = []string{"testString"}
				listFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListFields with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv2.ListFieldsOptions)
				listFieldsOptionsModel.ProjectID = core.StringPtr("testString")
				listFieldsOptionsModel.CollectionIds = []string{"testString"}
				listFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListFieldsOptions model with no property values
				listFieldsOptionsModelNew := new(discoveryv2.ListFieldsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListFields(listFieldsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListFields successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv2.ListFieldsOptions)
				listFieldsOptionsModel.ProjectID = core.StringPtr("testString")
				listFieldsOptionsModel.CollectionIds = []string{"testString"}
				listFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetComponentSettings(getComponentSettingsOptions *GetComponentSettingsOptions) - Operation response error`, func() {
		version := "testString"
		getComponentSettingsPath := "/v2/projects/testString/component_settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponentSettings with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetComponentSettingsOptions model
				getComponentSettingsOptionsModel := new(discoveryv2.GetComponentSettingsOptions)
				getComponentSettingsOptionsModel.ProjectID = core.StringPtr("testString")
				getComponentSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetComponentSettings(getComponentSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetComponentSettings(getComponentSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetComponentSettings(getComponentSettingsOptions *GetComponentSettingsOptions)`, func() {
		version := "testString"
		getComponentSettingsPath := "/v2/projects/testString/component_settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"fields_shown": {"body": {"use_passage": true, "field": "Field"}, "title": {"field": "Field"}}, "autocomplete": true, "structured_search": true, "results_per_page": 14, "aggregations": [{"name": "Name", "label": "Label", "multiple_selections_allowed": false, "visualization_type": "auto"}]}`)
				}))
			})
			It(`Invoke GetComponentSettings successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetComponentSettingsOptions model
				getComponentSettingsOptionsModel := new(discoveryv2.GetComponentSettingsOptions)
				getComponentSettingsOptionsModel.ProjectID = core.StringPtr("testString")
				getComponentSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetComponentSettingsWithContext(ctx, getComponentSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetComponentSettings(getComponentSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetComponentSettingsWithContext(ctx, getComponentSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"fields_shown": {"body": {"use_passage": true, "field": "Field"}, "title": {"field": "Field"}}, "autocomplete": true, "structured_search": true, "results_per_page": 14, "aggregations": [{"name": "Name", "label": "Label", "multiple_selections_allowed": false, "visualization_type": "auto"}]}`)
				}))
			})
			It(`Invoke GetComponentSettings successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetComponentSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentSettingsOptions model
				getComponentSettingsOptionsModel := new(discoveryv2.GetComponentSettingsOptions)
				getComponentSettingsOptionsModel.ProjectID = core.StringPtr("testString")
				getComponentSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetComponentSettings(getComponentSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetComponentSettings with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetComponentSettingsOptions model
				getComponentSettingsOptionsModel := new(discoveryv2.GetComponentSettingsOptions)
				getComponentSettingsOptionsModel.ProjectID = core.StringPtr("testString")
				getComponentSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetComponentSettings(getComponentSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentSettingsOptions model with no property values
				getComponentSettingsOptionsModelNew := new(discoveryv2.GetComponentSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetComponentSettings(getComponentSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetComponentSettings successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetComponentSettingsOptions model
				getComponentSettingsOptionsModel := new(discoveryv2.GetComponentSettingsOptions)
				getComponentSettingsOptionsModel.ProjectID = core.StringPtr("testString")
				getComponentSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetComponentSettings(getComponentSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddDocument(addDocumentOptions *AddDocumentOptions) - Operation response error`, func() {
		version := "testString"
		addDocumentPath := "/v2/projects/testString/collections/testString/documents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv2.AddDocumentOptions)
				addDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
				addDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				addDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddDocument(addDocumentOptions *AddDocumentOptions)`, func() {
		version := "testString"
		addDocumentPath := "/v2/projects/testString/collections/testString/documents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing"}`)
				}))
			})
			It(`Invoke AddDocument successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv2.AddDocumentOptions)
				addDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
				addDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				addDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.AddDocumentWithContext(ctx, addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.AddDocumentWithContext(ctx, addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing"}`)
				}))
			})
			It(`Invoke AddDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.AddDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv2.AddDocumentOptions)
				addDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
				addDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				addDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddDocument with error: Param validation error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv2.AddDocumentOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke AddDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv2.AddDocumentOptions)
				addDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
				addDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				addDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddDocumentOptions model with no property values
				addDocumentOptionsModelNew := new(discoveryv2.AddDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.AddDocument(addDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke AddDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv2.AddDocumentOptions)
				addDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
				addDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				addDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDocument(updateDocumentOptions *UpdateDocumentOptions) - Operation response error`, func() {
		version := "testString"
		updateDocumentPath := "/v2/projects/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv2.UpdateDocumentOptions)
				updateDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
				updateDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				updateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDocument(updateDocumentOptions *UpdateDocumentOptions)`, func() {
		version := "testString"
		updateDocumentPath := "/v2/projects/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing"}`)
				}))
			})
			It(`Invoke UpdateDocument successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv2.UpdateDocumentOptions)
				updateDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
				updateDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				updateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateDocumentWithContext(ctx, updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateDocumentWithContext(ctx, updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing"}`)
				}))
			})
			It(`Invoke UpdateDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv2.UpdateDocumentOptions)
				updateDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
				updateDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				updateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDocument with error: Param validation error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv2.UpdateDocumentOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke UpdateDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv2.UpdateDocumentOptions)
				updateDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
				updateDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				updateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDocumentOptions model with no property values
				updateDocumentOptionsModelNew := new(discoveryv2.UpdateDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateDocument(updateDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke UpdateDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv2.UpdateDocumentOptions)
				updateDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
				updateDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				updateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) - Operation response error`, func() {
		version := "testString"
		deleteDocumentPath := "/v2/projects/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv2.DeleteDocumentOptions)
				deleteDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
		version := "testString"
		deleteDocumentPath := "/v2/projects/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteDocument successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv2.DeleteDocumentOptions)
				deleteDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.DeleteDocumentWithContext(ctx, deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.DeleteDocumentWithContext(ctx, deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Watson-Discovery-Force"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Discovery-Force"][0]).To(Equal(fmt.Sprintf("%v", false)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.DeleteDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv2.DeleteDocumentOptions)
				deleteDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv2.DeleteDocumentOptions)
				deleteDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDocumentOptions model with no property values
				deleteDocumentOptionsModelNew := new(discoveryv2.DeleteDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.DeleteDocument(deleteDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv2.DeleteDocumentOptions)
				deleteDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.XWatsonDiscoveryForce = core.BoolPtr(false)
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTrainingQueries(listTrainingQueriesOptions *ListTrainingQueriesOptions) - Operation response error`, func() {
		version := "testString"
		listTrainingQueriesPath := "/v2/projects/testString/training_data/queries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingQueriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTrainingQueries with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingQueriesOptions model
				listTrainingQueriesOptionsModel := new(discoveryv2.ListTrainingQueriesOptions)
				listTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				listTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTrainingQueries(listTrainingQueriesOptions *ListTrainingQueriesOptions)`, func() {
		version := "testString"
		listTrainingQueriesPath := "/v2/projects/testString/training_data/queries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingQueriesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"queries": [{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke ListTrainingQueries successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListTrainingQueriesOptions model
				listTrainingQueriesOptionsModel := new(discoveryv2.ListTrainingQueriesOptions)
				listTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				listTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListTrainingQueriesWithContext(ctx, listTrainingQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListTrainingQueriesWithContext(ctx, listTrainingQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingQueriesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"queries": [{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke ListTrainingQueries successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListTrainingQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTrainingQueriesOptions model
				listTrainingQueriesOptionsModel := new(discoveryv2.ListTrainingQueriesOptions)
				listTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				listTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTrainingQueries with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingQueriesOptions model
				listTrainingQueriesOptionsModel := new(discoveryv2.ListTrainingQueriesOptions)
				listTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				listTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTrainingQueriesOptions model with no property values
				listTrainingQueriesOptionsModelNew := new(discoveryv2.ListTrainingQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListTrainingQueries successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingQueriesOptions model
				listTrainingQueriesOptionsModel := new(discoveryv2.ListTrainingQueriesOptions)
				listTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				listTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListTrainingQueries(listTrainingQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTrainingQueries(deleteTrainingQueriesOptions *DeleteTrainingQueriesOptions)`, func() {
		version := "testString"
		deleteTrainingQueriesPath := "/v2/projects/testString/training_data/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrainingQueriesPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTrainingQueries successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteTrainingQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTrainingQueriesOptions model
				deleteTrainingQueriesOptionsModel := new(discoveryv2.DeleteTrainingQueriesOptions)
				deleteTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				deleteTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteTrainingQueries(deleteTrainingQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTrainingQueries with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteTrainingQueriesOptions model
				deleteTrainingQueriesOptionsModel := new(discoveryv2.DeleteTrainingQueriesOptions)
				deleteTrainingQueriesOptionsModel.ProjectID = core.StringPtr("testString")
				deleteTrainingQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteTrainingQueries(deleteTrainingQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTrainingQueriesOptions model with no property values
				deleteTrainingQueriesOptionsModelNew := new(discoveryv2.DeleteTrainingQueriesOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteTrainingQueries(deleteTrainingQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTrainingQuery(createTrainingQueryOptions *CreateTrainingQueryOptions) - Operation response error`, func() {
		version := "testString"
		createTrainingQueryPath := "/v2/projects/testString/training_data/queries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrainingQueryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTrainingQuery with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateTrainingQueryOptions model
				createTrainingQueryOptionsModel := new(discoveryv2.CreateTrainingQueryOptions)
				createTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				createTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				createTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTrainingQuery(createTrainingQueryOptions *CreateTrainingQueryOptions)`, func() {
		version := "testString"
		createTrainingQueryPath := "/v2/projects/testString/training_data/queries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrainingQueryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke CreateTrainingQuery successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateTrainingQueryOptions model
				createTrainingQueryOptionsModel := new(discoveryv2.CreateTrainingQueryOptions)
				createTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				createTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				createTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateTrainingQueryWithContext(ctx, createTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateTrainingQueryWithContext(ctx, createTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrainingQueryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke CreateTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateTrainingQueryOptions model
				createTrainingQueryOptionsModel := new(discoveryv2.CreateTrainingQueryOptions)
				createTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				createTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				createTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTrainingQuery with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateTrainingQueryOptions model
				createTrainingQueryOptionsModel := new(discoveryv2.CreateTrainingQueryOptions)
				createTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				createTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				createTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTrainingQueryOptions model with no property values
				createTrainingQueryOptionsModelNew := new(discoveryv2.CreateTrainingQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the CreateTrainingQueryOptions model
				createTrainingQueryOptionsModel := new(discoveryv2.CreateTrainingQueryOptions)
				createTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				createTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				createTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				createTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateTrainingQuery(createTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrainingQuery(getTrainingQueryOptions *GetTrainingQueryOptions) - Operation response error`, func() {
		version := "testString"
		getTrainingQueryPath := "/v2/projects/testString/training_data/queries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingQueryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTrainingQuery with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingQueryOptions model
				getTrainingQueryOptionsModel := new(discoveryv2.GetTrainingQueryOptions)
				getTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetTrainingQuery(getTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetTrainingQuery(getTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrainingQuery(getTrainingQueryOptions *GetTrainingQueryOptions)`, func() {
		version := "testString"
		getTrainingQueryPath := "/v2/projects/testString/training_data/queries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingQueryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetTrainingQuery successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetTrainingQueryOptions model
				getTrainingQueryOptionsModel := new(discoveryv2.GetTrainingQueryOptions)
				getTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetTrainingQueryWithContext(ctx, getTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetTrainingQuery(getTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetTrainingQueryWithContext(ctx, getTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingQueryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTrainingQueryOptions model
				getTrainingQueryOptionsModel := new(discoveryv2.GetTrainingQueryOptions)
				getTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetTrainingQuery(getTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTrainingQuery with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingQueryOptions model
				getTrainingQueryOptionsModel := new(discoveryv2.GetTrainingQueryOptions)
				getTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetTrainingQuery(getTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTrainingQueryOptions model with no property values
				getTrainingQueryOptionsModelNew := new(discoveryv2.GetTrainingQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetTrainingQuery(getTrainingQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingQueryOptions model
				getTrainingQueryOptionsModel := new(discoveryv2.GetTrainingQueryOptions)
				getTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetTrainingQuery(getTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTrainingQuery(updateTrainingQueryOptions *UpdateTrainingQueryOptions) - Operation response error`, func() {
		version := "testString"
		updateTrainingQueryPath := "/v2/projects/testString/training_data/queries/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrainingQueryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTrainingQuery with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateTrainingQueryOptions model
				updateTrainingQueryOptionsModel := new(discoveryv2.UpdateTrainingQueryOptions)
				updateTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				updateTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTrainingQuery(updateTrainingQueryOptions *UpdateTrainingQueryOptions)`, func() {
		version := "testString"
		updateTrainingQueryPath := "/v2/projects/testString/training_data/queries/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrainingQueryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke UpdateTrainingQuery successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateTrainingQueryOptions model
				updateTrainingQueryOptionsModel := new(discoveryv2.UpdateTrainingQueryOptions)
				updateTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				updateTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateTrainingQueryWithContext(ctx, updateTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateTrainingQueryWithContext(ctx, updateTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrainingQueryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "examples": [{"document_id": "DocumentID", "collection_id": "CollectionID", "relevance": 9, "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke UpdateTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateTrainingQueryOptions model
				updateTrainingQueryOptionsModel := new(discoveryv2.UpdateTrainingQueryOptions)
				updateTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				updateTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTrainingQuery with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateTrainingQueryOptions model
				updateTrainingQueryOptionsModel := new(discoveryv2.UpdateTrainingQueryOptions)
				updateTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				updateTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTrainingQueryOptions model with no property values
				updateTrainingQueryOptionsModelNew := new(discoveryv2.UpdateTrainingQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateTrainingQueryOptions model
				updateTrainingQueryOptionsModel := new(discoveryv2.UpdateTrainingQueryOptions)
				updateTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Examples = []discoveryv2.TrainingExample{*trainingExampleModel}
				updateTrainingQueryOptionsModel.Filter = core.StringPtr("testString")
				updateTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateTrainingQuery(updateTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTrainingQuery(deleteTrainingQueryOptions *DeleteTrainingQueryOptions)`, func() {
		version := "testString"
		deleteTrainingQueryPath := "/v2/projects/testString/training_data/queries/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrainingQueryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTrainingQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteTrainingQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTrainingQueryOptions model
				deleteTrainingQueryOptionsModel := new(discoveryv2.DeleteTrainingQueryOptions)
				deleteTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				deleteTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				deleteTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteTrainingQuery(deleteTrainingQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTrainingQuery with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteTrainingQueryOptions model
				deleteTrainingQueryOptionsModel := new(discoveryv2.DeleteTrainingQueryOptions)
				deleteTrainingQueryOptionsModel.ProjectID = core.StringPtr("testString")
				deleteTrainingQueryOptionsModel.QueryID = core.StringPtr("testString")
				deleteTrainingQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteTrainingQuery(deleteTrainingQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTrainingQueryOptions model with no property values
				deleteTrainingQueryOptionsModelNew := new(discoveryv2.DeleteTrainingQueryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteTrainingQuery(deleteTrainingQueryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AnalyzeDocument(analyzeDocumentOptions *AnalyzeDocumentOptions) - Operation response error`, func() {
		version := "testString"
		analyzeDocumentPath := "/v2/projects/testString/collections/testString/analyze"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzeDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AnalyzeDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeDocumentOptions model
				analyzeDocumentOptionsModel := new(discoveryv2.AnalyzeDocumentOptions)
				analyzeDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				analyzeDocumentOptionsModel.Filename = core.StringPtr("testString")
				analyzeDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				analyzeDocumentOptionsModel.Metadata = core.StringPtr("testString")
				analyzeDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AnalyzeDocument(analyzeDocumentOptions *AnalyzeDocumentOptions)`, func() {
		version := "testString"
		analyzeDocumentPath := "/v2/projects/testString/collections/testString/analyze"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzeDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "collection_id": "CollectionID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}], "result": {"metadata": {"mapKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke AnalyzeDocument successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the AnalyzeDocumentOptions model
				analyzeDocumentOptionsModel := new(discoveryv2.AnalyzeDocumentOptions)
				analyzeDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				analyzeDocumentOptionsModel.Filename = core.StringPtr("testString")
				analyzeDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				analyzeDocumentOptionsModel.Metadata = core.StringPtr("testString")
				analyzeDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.AnalyzeDocumentWithContext(ctx, analyzeDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.AnalyzeDocumentWithContext(ctx, analyzeDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzeDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "collection_id": "CollectionID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}], "result": {"metadata": {"mapKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke AnalyzeDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.AnalyzeDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AnalyzeDocumentOptions model
				analyzeDocumentOptionsModel := new(discoveryv2.AnalyzeDocumentOptions)
				analyzeDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				analyzeDocumentOptionsModel.Filename = core.StringPtr("testString")
				analyzeDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				analyzeDocumentOptionsModel.Metadata = core.StringPtr("testString")
				analyzeDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AnalyzeDocument with error: Param validation error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeDocumentOptions model
				analyzeDocumentOptionsModel := new(discoveryv2.AnalyzeDocumentOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke AnalyzeDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeDocumentOptions model
				analyzeDocumentOptionsModel := new(discoveryv2.AnalyzeDocumentOptions)
				analyzeDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				analyzeDocumentOptionsModel.Filename = core.StringPtr("testString")
				analyzeDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				analyzeDocumentOptionsModel.Metadata = core.StringPtr("testString")
				analyzeDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AnalyzeDocumentOptions model with no property values
				analyzeDocumentOptionsModelNew := new(discoveryv2.AnalyzeDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.AnalyzeDocument(analyzeDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke AnalyzeDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AnalyzeDocumentOptions model
				analyzeDocumentOptionsModel := new(discoveryv2.AnalyzeDocumentOptions)
				analyzeDocumentOptionsModel.ProjectID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				analyzeDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				analyzeDocumentOptionsModel.Filename = core.StringPtr("testString")
				analyzeDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				analyzeDocumentOptionsModel.Metadata = core.StringPtr("testString")
				analyzeDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.AnalyzeDocument(analyzeDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnrichments(listEnrichmentsOptions *ListEnrichmentsOptions) - Operation response error`, func() {
		version := "testString"
		listEnrichmentsPath := "/v2/projects/testString/enrichments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnrichmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEnrichments with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListEnrichmentsOptions model
				listEnrichmentsOptionsModel := new(discoveryv2.ListEnrichmentsOptions)
				listEnrichmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listEnrichmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListEnrichments(listEnrichmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListEnrichments(listEnrichmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnrichments(listEnrichmentsOptions *ListEnrichmentsOptions)`, func() {
		version := "testString"
		listEnrichmentsPath := "/v2/projects/testString/enrichments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnrichmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enrichments": [{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}]}`)
				}))
			})
			It(`Invoke ListEnrichments successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListEnrichmentsOptions model
				listEnrichmentsOptionsModel := new(discoveryv2.ListEnrichmentsOptions)
				listEnrichmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listEnrichmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListEnrichmentsWithContext(ctx, listEnrichmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListEnrichments(listEnrichmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListEnrichmentsWithContext(ctx, listEnrichmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnrichmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enrichments": [{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}]}`)
				}))
			})
			It(`Invoke ListEnrichments successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListEnrichments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnrichmentsOptions model
				listEnrichmentsOptionsModel := new(discoveryv2.ListEnrichmentsOptions)
				listEnrichmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listEnrichmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListEnrichments(listEnrichmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEnrichments with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListEnrichmentsOptions model
				listEnrichmentsOptionsModel := new(discoveryv2.ListEnrichmentsOptions)
				listEnrichmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listEnrichmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListEnrichments(listEnrichmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListEnrichmentsOptions model with no property values
				listEnrichmentsOptionsModelNew := new(discoveryv2.ListEnrichmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListEnrichments(listEnrichmentsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListEnrichments successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListEnrichmentsOptions model
				listEnrichmentsOptionsModel := new(discoveryv2.ListEnrichmentsOptions)
				listEnrichmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listEnrichmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListEnrichments(listEnrichmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEnrichment(createEnrichmentOptions *CreateEnrichmentOptions) - Operation response error`, func() {
		version := "testString"
		createEnrichmentPath := "/v2/projects/testString/enrichments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnrichmentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEnrichment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv2.EnrichmentOptions)
				enrichmentOptionsModel.Languages = []string{"testString"}
				enrichmentOptionsModel.EntityType = core.StringPtr("testString")
				enrichmentOptionsModel.RegularExpression = core.StringPtr("testString")
				enrichmentOptionsModel.ResultField = core.StringPtr("testString")

				// Construct an instance of the CreateEnrichment model
				createEnrichmentModel := new(discoveryv2.CreateEnrichment)
				createEnrichmentModel.Name = core.StringPtr("testString")
				createEnrichmentModel.Description = core.StringPtr("testString")
				createEnrichmentModel.Type = core.StringPtr("dictionary")
				createEnrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the CreateEnrichmentOptions model
				createEnrichmentOptionsModel := new(discoveryv2.CreateEnrichmentOptions)
				createEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				createEnrichmentOptionsModel.Enrichment = createEnrichmentModel
				createEnrichmentOptionsModel.File = CreateMockReader("This is a mock file.")
				createEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateEnrichment(createEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateEnrichment(createEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEnrichment(createEnrichmentOptions *CreateEnrichmentOptions)`, func() {
		version := "testString"
		createEnrichmentPath := "/v2/projects/testString/enrichments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnrichmentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}`)
				}))
			})
			It(`Invoke CreateEnrichment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv2.EnrichmentOptions)
				enrichmentOptionsModel.Languages = []string{"testString"}
				enrichmentOptionsModel.EntityType = core.StringPtr("testString")
				enrichmentOptionsModel.RegularExpression = core.StringPtr("testString")
				enrichmentOptionsModel.ResultField = core.StringPtr("testString")

				// Construct an instance of the CreateEnrichment model
				createEnrichmentModel := new(discoveryv2.CreateEnrichment)
				createEnrichmentModel.Name = core.StringPtr("testString")
				createEnrichmentModel.Description = core.StringPtr("testString")
				createEnrichmentModel.Type = core.StringPtr("dictionary")
				createEnrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the CreateEnrichmentOptions model
				createEnrichmentOptionsModel := new(discoveryv2.CreateEnrichmentOptions)
				createEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				createEnrichmentOptionsModel.Enrichment = createEnrichmentModel
				createEnrichmentOptionsModel.File = CreateMockReader("This is a mock file.")
				createEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateEnrichmentWithContext(ctx, createEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateEnrichment(createEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateEnrichmentWithContext(ctx, createEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnrichmentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}`)
				}))
			})
			It(`Invoke CreateEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv2.EnrichmentOptions)
				enrichmentOptionsModel.Languages = []string{"testString"}
				enrichmentOptionsModel.EntityType = core.StringPtr("testString")
				enrichmentOptionsModel.RegularExpression = core.StringPtr("testString")
				enrichmentOptionsModel.ResultField = core.StringPtr("testString")

				// Construct an instance of the CreateEnrichment model
				createEnrichmentModel := new(discoveryv2.CreateEnrichment)
				createEnrichmentModel.Name = core.StringPtr("testString")
				createEnrichmentModel.Description = core.StringPtr("testString")
				createEnrichmentModel.Type = core.StringPtr("dictionary")
				createEnrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the CreateEnrichmentOptions model
				createEnrichmentOptionsModel := new(discoveryv2.CreateEnrichmentOptions)
				createEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				createEnrichmentOptionsModel.Enrichment = createEnrichmentModel
				createEnrichmentOptionsModel.File = CreateMockReader("This is a mock file.")
				createEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateEnrichment(createEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEnrichment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv2.EnrichmentOptions)
				enrichmentOptionsModel.Languages = []string{"testString"}
				enrichmentOptionsModel.EntityType = core.StringPtr("testString")
				enrichmentOptionsModel.RegularExpression = core.StringPtr("testString")
				enrichmentOptionsModel.ResultField = core.StringPtr("testString")

				// Construct an instance of the CreateEnrichment model
				createEnrichmentModel := new(discoveryv2.CreateEnrichment)
				createEnrichmentModel.Name = core.StringPtr("testString")
				createEnrichmentModel.Description = core.StringPtr("testString")
				createEnrichmentModel.Type = core.StringPtr("dictionary")
				createEnrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the CreateEnrichmentOptions model
				createEnrichmentOptionsModel := new(discoveryv2.CreateEnrichmentOptions)
				createEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				createEnrichmentOptionsModel.Enrichment = createEnrichmentModel
				createEnrichmentOptionsModel.File = CreateMockReader("This is a mock file.")
				createEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateEnrichment(createEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEnrichmentOptions model with no property values
				createEnrichmentOptionsModelNew := new(discoveryv2.CreateEnrichmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateEnrichment(createEnrichmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv2.EnrichmentOptions)
				enrichmentOptionsModel.Languages = []string{"testString"}
				enrichmentOptionsModel.EntityType = core.StringPtr("testString")
				enrichmentOptionsModel.RegularExpression = core.StringPtr("testString")
				enrichmentOptionsModel.ResultField = core.StringPtr("testString")

				// Construct an instance of the CreateEnrichment model
				createEnrichmentModel := new(discoveryv2.CreateEnrichment)
				createEnrichmentModel.Name = core.StringPtr("testString")
				createEnrichmentModel.Description = core.StringPtr("testString")
				createEnrichmentModel.Type = core.StringPtr("dictionary")
				createEnrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the CreateEnrichmentOptions model
				createEnrichmentOptionsModel := new(discoveryv2.CreateEnrichmentOptions)
				createEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				createEnrichmentOptionsModel.Enrichment = createEnrichmentModel
				createEnrichmentOptionsModel.File = CreateMockReader("This is a mock file.")
				createEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateEnrichment(createEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnrichment(getEnrichmentOptions *GetEnrichmentOptions) - Operation response error`, func() {
		version := "testString"
		getEnrichmentPath := "/v2/projects/testString/enrichments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnrichmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnrichment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetEnrichmentOptions model
				getEnrichmentOptionsModel := new(discoveryv2.GetEnrichmentOptions)
				getEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				getEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				getEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetEnrichment(getEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetEnrichment(getEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnrichment(getEnrichmentOptions *GetEnrichmentOptions)`, func() {
		version := "testString"
		getEnrichmentPath := "/v2/projects/testString/enrichments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnrichmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}`)
				}))
			})
			It(`Invoke GetEnrichment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetEnrichmentOptions model
				getEnrichmentOptionsModel := new(discoveryv2.GetEnrichmentOptions)
				getEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				getEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				getEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetEnrichmentWithContext(ctx, getEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetEnrichment(getEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetEnrichmentWithContext(ctx, getEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnrichmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}`)
				}))
			})
			It(`Invoke GetEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnrichmentOptions model
				getEnrichmentOptionsModel := new(discoveryv2.GetEnrichmentOptions)
				getEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				getEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				getEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetEnrichment(getEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEnrichment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetEnrichmentOptions model
				getEnrichmentOptionsModel := new(discoveryv2.GetEnrichmentOptions)
				getEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				getEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				getEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetEnrichment(getEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEnrichmentOptions model with no property values
				getEnrichmentOptionsModelNew := new(discoveryv2.GetEnrichmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetEnrichment(getEnrichmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetEnrichmentOptions model
				getEnrichmentOptionsModel := new(discoveryv2.GetEnrichmentOptions)
				getEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				getEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				getEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetEnrichment(getEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEnrichment(updateEnrichmentOptions *UpdateEnrichmentOptions) - Operation response error`, func() {
		version := "testString"
		updateEnrichmentPath := "/v2/projects/testString/enrichments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnrichmentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEnrichment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateEnrichmentOptions model
				updateEnrichmentOptionsModel := new(discoveryv2.UpdateEnrichmentOptions)
				updateEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Name = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Description = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateEnrichment(updateEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateEnrichment(updateEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEnrichment(updateEnrichmentOptions *UpdateEnrichmentOptions)`, func() {
		version := "testString"
		updateEnrichmentPath := "/v2/projects/testString/enrichments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnrichmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}`)
				}))
			})
			It(`Invoke UpdateEnrichment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateEnrichmentOptions model
				updateEnrichmentOptionsModel := new(discoveryv2.UpdateEnrichmentOptions)
				updateEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Name = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Description = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateEnrichmentWithContext(ctx, updateEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateEnrichment(updateEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateEnrichmentWithContext(ctx, updateEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnrichmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"enrichment_id": "EnrichmentID", "name": "Name", "description": "Description", "type": "part_of_speech", "options": {"languages": ["Languages"], "entity_type": "EntityType", "regular_expression": "RegularExpression", "result_field": "ResultField"}}`)
				}))
			})
			It(`Invoke UpdateEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEnrichmentOptions model
				updateEnrichmentOptionsModel := new(discoveryv2.UpdateEnrichmentOptions)
				updateEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Name = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Description = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateEnrichment(updateEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEnrichment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateEnrichmentOptions model
				updateEnrichmentOptionsModel := new(discoveryv2.UpdateEnrichmentOptions)
				updateEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Name = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Description = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateEnrichment(updateEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEnrichmentOptions model with no property values
				updateEnrichmentOptionsModelNew := new(discoveryv2.UpdateEnrichmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateEnrichment(updateEnrichmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateEnrichmentOptions model
				updateEnrichmentOptionsModel := new(discoveryv2.UpdateEnrichmentOptions)
				updateEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Name = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Description = core.StringPtr("testString")
				updateEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateEnrichment(updateEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEnrichment(deleteEnrichmentOptions *DeleteEnrichmentOptions)`, func() {
		version := "testString"
		deleteEnrichmentPath := "/v2/projects/testString/enrichments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnrichmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteEnrichment successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteEnrichment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEnrichmentOptions model
				deleteEnrichmentOptionsModel := new(discoveryv2.DeleteEnrichmentOptions)
				deleteEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				deleteEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteEnrichment(deleteEnrichmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEnrichment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteEnrichmentOptions model
				deleteEnrichmentOptionsModel := new(discoveryv2.DeleteEnrichmentOptions)
				deleteEnrichmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteEnrichmentOptionsModel.EnrichmentID = core.StringPtr("testString")
				deleteEnrichmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteEnrichment(deleteEnrichmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEnrichmentOptions model with no property values
				deleteEnrichmentOptionsModelNew := new(discoveryv2.DeleteEnrichmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteEnrichment(deleteEnrichmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions) - Operation response error`, func() {
		version := "testString"
		listProjectsPath := "/v2/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjects with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(discoveryv2.ListProjectsOptions)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions)`, func() {
		version := "testString"
		listProjectsPath := "/v2/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"projects": [{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15}]}`)
				}))
			})
			It(`Invoke ListProjects successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(discoveryv2.ListProjectsOptions)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"projects": [{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15}]}`)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListProjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(discoveryv2.ListProjectsOptions)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjects with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(discoveryv2.ListProjectsOptions)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(discoveryv2.ListProjectsOptions)
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions) - Operation response error`, func() {
		version := "testString"
		createProjectPath := "/v2/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProject with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DefaultQueryParamsPassages model
				defaultQueryParamsPassagesModel := new(discoveryv2.DefaultQueryParamsPassages)
				defaultQueryParamsPassagesModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.Fields = []string{"testString"}
				defaultQueryParamsPassagesModel.Characters = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.PerDocument = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsTableResults model
				defaultQueryParamsTableResultsModel := new(discoveryv2.DefaultQueryParamsTableResults)
				defaultQueryParamsTableResultsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsTableResultsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsTableResultsModel.PerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsSuggestedRefinements model
				defaultQueryParamsSuggestedRefinementsModel := new(discoveryv2.DefaultQueryParamsSuggestedRefinements)
				defaultQueryParamsSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsSuggestedRefinementsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParams model
				defaultQueryParamsModel := new(discoveryv2.DefaultQueryParams)
				defaultQueryParamsModel.CollectionIds = []string{"testString"}
				defaultQueryParamsModel.Passages = defaultQueryParamsPassagesModel
				defaultQueryParamsModel.TableResults = defaultQueryParamsTableResultsModel
				defaultQueryParamsModel.Aggregation = core.StringPtr("testString")
				defaultQueryParamsModel.SuggestedRefinements = defaultQueryParamsSuggestedRefinementsModel
				defaultQueryParamsModel.SpellingSuggestions = core.BoolPtr(true)
				defaultQueryParamsModel.Highlight = core.BoolPtr(true)
				defaultQueryParamsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsModel.Sort = core.StringPtr("testString")
				defaultQueryParamsModel.Return = []string{"testString"}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(discoveryv2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("testString")
				createProjectOptionsModel.Type = core.StringPtr("document_retrieval")
				createProjectOptionsModel.DefaultQueryParameters = defaultQueryParamsModel
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
		version := "testString"
		createProjectPath := "/v2/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15, "default_query_parameters": {"collection_ids": ["CollectionIds"], "passages": {"enabled": false, "count": 5, "fields": ["Fields"], "characters": 10, "per_document": false, "max_per_document": 14}, "table_results": {"enabled": false, "count": 5, "per_document": 11}, "aggregation": "Aggregation", "suggested_refinements": {"enabled": false, "count": 5}, "spelling_suggestions": false, "highlight": false, "count": 5, "sort": "Sort", "return": ["Return"]}}`)
				}))
			})
			It(`Invoke CreateProject successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DefaultQueryParamsPassages model
				defaultQueryParamsPassagesModel := new(discoveryv2.DefaultQueryParamsPassages)
				defaultQueryParamsPassagesModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.Fields = []string{"testString"}
				defaultQueryParamsPassagesModel.Characters = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.PerDocument = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsTableResults model
				defaultQueryParamsTableResultsModel := new(discoveryv2.DefaultQueryParamsTableResults)
				defaultQueryParamsTableResultsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsTableResultsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsTableResultsModel.PerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsSuggestedRefinements model
				defaultQueryParamsSuggestedRefinementsModel := new(discoveryv2.DefaultQueryParamsSuggestedRefinements)
				defaultQueryParamsSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsSuggestedRefinementsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParams model
				defaultQueryParamsModel := new(discoveryv2.DefaultQueryParams)
				defaultQueryParamsModel.CollectionIds = []string{"testString"}
				defaultQueryParamsModel.Passages = defaultQueryParamsPassagesModel
				defaultQueryParamsModel.TableResults = defaultQueryParamsTableResultsModel
				defaultQueryParamsModel.Aggregation = core.StringPtr("testString")
				defaultQueryParamsModel.SuggestedRefinements = defaultQueryParamsSuggestedRefinementsModel
				defaultQueryParamsModel.SpellingSuggestions = core.BoolPtr(true)
				defaultQueryParamsModel.Highlight = core.BoolPtr(true)
				defaultQueryParamsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsModel.Sort = core.StringPtr("testString")
				defaultQueryParamsModel.Return = []string{"testString"}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(discoveryv2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("testString")
				createProjectOptionsModel.Type = core.StringPtr("document_retrieval")
				createProjectOptionsModel.DefaultQueryParameters = defaultQueryParamsModel
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15, "default_query_parameters": {"collection_ids": ["CollectionIds"], "passages": {"enabled": false, "count": 5, "fields": ["Fields"], "characters": 10, "per_document": false, "max_per_document": 14}, "table_results": {"enabled": false, "count": 5, "per_document": 11}, "aggregation": "Aggregation", "suggested_refinements": {"enabled": false, "count": 5}, "spelling_suggestions": false, "highlight": false, "count": 5, "sort": "Sort", "return": ["Return"]}}`)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DefaultQueryParamsPassages model
				defaultQueryParamsPassagesModel := new(discoveryv2.DefaultQueryParamsPassages)
				defaultQueryParamsPassagesModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.Fields = []string{"testString"}
				defaultQueryParamsPassagesModel.Characters = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.PerDocument = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsTableResults model
				defaultQueryParamsTableResultsModel := new(discoveryv2.DefaultQueryParamsTableResults)
				defaultQueryParamsTableResultsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsTableResultsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsTableResultsModel.PerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsSuggestedRefinements model
				defaultQueryParamsSuggestedRefinementsModel := new(discoveryv2.DefaultQueryParamsSuggestedRefinements)
				defaultQueryParamsSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsSuggestedRefinementsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParams model
				defaultQueryParamsModel := new(discoveryv2.DefaultQueryParams)
				defaultQueryParamsModel.CollectionIds = []string{"testString"}
				defaultQueryParamsModel.Passages = defaultQueryParamsPassagesModel
				defaultQueryParamsModel.TableResults = defaultQueryParamsTableResultsModel
				defaultQueryParamsModel.Aggregation = core.StringPtr("testString")
				defaultQueryParamsModel.SuggestedRefinements = defaultQueryParamsSuggestedRefinementsModel
				defaultQueryParamsModel.SpellingSuggestions = core.BoolPtr(true)
				defaultQueryParamsModel.Highlight = core.BoolPtr(true)
				defaultQueryParamsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsModel.Sort = core.StringPtr("testString")
				defaultQueryParamsModel.Return = []string{"testString"}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(discoveryv2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("testString")
				createProjectOptionsModel.Type = core.StringPtr("document_retrieval")
				createProjectOptionsModel.DefaultQueryParameters = defaultQueryParamsModel
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProject with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DefaultQueryParamsPassages model
				defaultQueryParamsPassagesModel := new(discoveryv2.DefaultQueryParamsPassages)
				defaultQueryParamsPassagesModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.Fields = []string{"testString"}
				defaultQueryParamsPassagesModel.Characters = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.PerDocument = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsTableResults model
				defaultQueryParamsTableResultsModel := new(discoveryv2.DefaultQueryParamsTableResults)
				defaultQueryParamsTableResultsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsTableResultsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsTableResultsModel.PerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsSuggestedRefinements model
				defaultQueryParamsSuggestedRefinementsModel := new(discoveryv2.DefaultQueryParamsSuggestedRefinements)
				defaultQueryParamsSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsSuggestedRefinementsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParams model
				defaultQueryParamsModel := new(discoveryv2.DefaultQueryParams)
				defaultQueryParamsModel.CollectionIds = []string{"testString"}
				defaultQueryParamsModel.Passages = defaultQueryParamsPassagesModel
				defaultQueryParamsModel.TableResults = defaultQueryParamsTableResultsModel
				defaultQueryParamsModel.Aggregation = core.StringPtr("testString")
				defaultQueryParamsModel.SuggestedRefinements = defaultQueryParamsSuggestedRefinementsModel
				defaultQueryParamsModel.SpellingSuggestions = core.BoolPtr(true)
				defaultQueryParamsModel.Highlight = core.BoolPtr(true)
				defaultQueryParamsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsModel.Sort = core.StringPtr("testString")
				defaultQueryParamsModel.Return = []string{"testString"}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(discoveryv2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("testString")
				createProjectOptionsModel.Type = core.StringPtr("document_retrieval")
				createProjectOptionsModel.DefaultQueryParameters = defaultQueryParamsModel
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectOptions model with no property values
				createProjectOptionsModelNew := new(discoveryv2.CreateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateProject(createProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DefaultQueryParamsPassages model
				defaultQueryParamsPassagesModel := new(discoveryv2.DefaultQueryParamsPassages)
				defaultQueryParamsPassagesModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.Fields = []string{"testString"}
				defaultQueryParamsPassagesModel.Characters = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.PerDocument = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsTableResults model
				defaultQueryParamsTableResultsModel := new(discoveryv2.DefaultQueryParamsTableResults)
				defaultQueryParamsTableResultsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsTableResultsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsTableResultsModel.PerDocument = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParamsSuggestedRefinements model
				defaultQueryParamsSuggestedRefinementsModel := new(discoveryv2.DefaultQueryParamsSuggestedRefinements)
				defaultQueryParamsSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsSuggestedRefinementsModel.Count = core.Int64Ptr(int64(38))

				// Construct an instance of the DefaultQueryParams model
				defaultQueryParamsModel := new(discoveryv2.DefaultQueryParams)
				defaultQueryParamsModel.CollectionIds = []string{"testString"}
				defaultQueryParamsModel.Passages = defaultQueryParamsPassagesModel
				defaultQueryParamsModel.TableResults = defaultQueryParamsTableResultsModel
				defaultQueryParamsModel.Aggregation = core.StringPtr("testString")
				defaultQueryParamsModel.SuggestedRefinements = defaultQueryParamsSuggestedRefinementsModel
				defaultQueryParamsModel.SpellingSuggestions = core.BoolPtr(true)
				defaultQueryParamsModel.Highlight = core.BoolPtr(true)
				defaultQueryParamsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsModel.Sort = core.StringPtr("testString")
				defaultQueryParamsModel.Return = []string{"testString"}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(discoveryv2.CreateProjectOptions)
				createProjectOptionsModel.Name = core.StringPtr("testString")
				createProjectOptionsModel.Type = core.StringPtr("document_retrieval")
				createProjectOptionsModel.DefaultQueryParameters = defaultQueryParamsModel
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions) - Operation response error`, func() {
		version := "testString"
		getProjectPath := "/v2/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProject with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(discoveryv2.GetProjectOptions)
				getProjectOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
		version := "testString"
		getProjectPath := "/v2/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15, "default_query_parameters": {"collection_ids": ["CollectionIds"], "passages": {"enabled": false, "count": 5, "fields": ["Fields"], "characters": 10, "per_document": false, "max_per_document": 14}, "table_results": {"enabled": false, "count": 5, "per_document": 11}, "aggregation": "Aggregation", "suggested_refinements": {"enabled": false, "count": 5}, "spelling_suggestions": false, "highlight": false, "count": 5, "sort": "Sort", "return": ["Return"]}}`)
				}))
			})
			It(`Invoke GetProject successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(discoveryv2.GetProjectOptions)
				getProjectOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15, "default_query_parameters": {"collection_ids": ["CollectionIds"], "passages": {"enabled": false, "count": 5, "fields": ["Fields"], "characters": 10, "per_document": false, "max_per_document": 14}, "table_results": {"enabled": false, "count": 5, "per_document": 11}, "aggregation": "Aggregation", "suggested_refinements": {"enabled": false, "count": 5}, "spelling_suggestions": false, "highlight": false, "count": 5, "sort": "Sort", "return": ["Return"]}}`)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(discoveryv2.GetProjectOptions)
				getProjectOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProject with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(discoveryv2.GetProjectOptions)
				getProjectOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectOptions model with no property values
				getProjectOptionsModelNew := new(discoveryv2.GetProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetProject(getProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(discoveryv2.GetProjectOptions)
				getProjectOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions) - Operation response error`, func() {
		version := "testString"
		updateProjectPath := "/v2/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProject with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(discoveryv2.UpdateProjectOptions)
				updateProjectOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("testString")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions)`, func() {
		version := "testString"
		updateProjectPath := "/v2/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15, "default_query_parameters": {"collection_ids": ["CollectionIds"], "passages": {"enabled": false, "count": 5, "fields": ["Fields"], "characters": 10, "per_document": false, "max_per_document": 14}, "table_results": {"enabled": false, "count": 5, "per_document": 11}, "aggregation": "Aggregation", "suggested_refinements": {"enabled": false, "count": 5}, "spelling_suggestions": false, "highlight": false, "count": 5, "sort": "Sort", "return": ["Return"]}}`)
				}))
			})
			It(`Invoke UpdateProject successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(discoveryv2.UpdateProjectOptions)
				updateProjectOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("testString")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"project_id": "ProjectID", "name": "Name", "type": "document_retrieval", "relevancy_training_status": {"data_updated": "DataUpdated", "total_examples": 13, "sufficient_label_diversity": true, "processing": true, "minimum_examples_added": true, "successfully_trained": "SuccessfullyTrained", "available": false, "notices": 7, "minimum_queries_added": false}, "collection_count": 15, "default_query_parameters": {"collection_ids": ["CollectionIds"], "passages": {"enabled": false, "count": 5, "fields": ["Fields"], "characters": 10, "per_document": false, "max_per_document": 14}, "table_results": {"enabled": false, "count": 5, "per_document": 11}, "aggregation": "Aggregation", "suggested_refinements": {"enabled": false, "count": 5}, "spelling_suggestions": false, "highlight": false, "count": 5, "sort": "Sort", "return": ["Return"]}}`)
				}))
			})
			It(`Invoke UpdateProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(discoveryv2.UpdateProjectOptions)
				updateProjectOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("testString")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProject with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(discoveryv2.UpdateProjectOptions)
				updateProjectOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("testString")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectOptions model with no property values
				updateProjectOptionsModelNew := new(discoveryv2.UpdateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateProject(updateProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(discoveryv2.UpdateProjectOptions)
				updateProjectOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectOptionsModel.Name = core.StringPtr("testString")
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
		version := "testString"
		deleteProjectPath := "/v2/projects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProject successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(discoveryv2.DeleteProjectOptions)
				deleteProjectOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProject with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(discoveryv2.DeleteProjectOptions)
				deleteProjectOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProjectOptions model with no property values
				deleteProjectOptionsModelNew := new(discoveryv2.DeleteProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteProject(deleteProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		version := "testString"
		deleteUserDataPath := "/v2/user_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteUserDataPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteUserData successfully`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(discoveryv2.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(discoveryv2.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteUserDataOptions model with no property values
				deleteUserDataOptionsModelNew := new(discoveryv2.DeleteUserDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteUserData(deleteUserDataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			discoveryService, _ := discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
				URL:           "http://discoveryv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAddDocumentOptions successfully`, func() {
				// Construct an instance of the AddDocumentOptions model
				projectID := "testString"
				collectionID := "testString"
				addDocumentOptionsModel := discoveryService.NewAddDocumentOptions(projectID, collectionID)
				addDocumentOptionsModel.SetProjectID("testString")
				addDocumentOptionsModel.SetCollectionID("testString")
				addDocumentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				addDocumentOptionsModel.SetFilename("testString")
				addDocumentOptionsModel.SetFileContentType("application/json")
				addDocumentOptionsModel.SetMetadata("testString")
				addDocumentOptionsModel.SetXWatsonDiscoveryForce(false)
				addDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addDocumentOptionsModel).ToNot(BeNil())
				Expect(addDocumentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(addDocumentOptionsModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.FileContentType).To(Equal(core.StringPtr("application/json")))
				Expect(addDocumentOptionsModel.Metadata).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.XWatsonDiscoveryForce).To(Equal(core.BoolPtr(false)))
				Expect(addDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAnalyzeDocumentOptions successfully`, func() {
				// Construct an instance of the AnalyzeDocumentOptions model
				projectID := "testString"
				collectionID := "testString"
				analyzeDocumentOptionsModel := discoveryService.NewAnalyzeDocumentOptions(projectID, collectionID)
				analyzeDocumentOptionsModel.SetProjectID("testString")
				analyzeDocumentOptionsModel.SetCollectionID("testString")
				analyzeDocumentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				analyzeDocumentOptionsModel.SetFilename("testString")
				analyzeDocumentOptionsModel.SetFileContentType("application/json")
				analyzeDocumentOptionsModel.SetMetadata("testString")
				analyzeDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(analyzeDocumentOptionsModel).ToNot(BeNil())
				Expect(analyzeDocumentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(analyzeDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(analyzeDocumentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(analyzeDocumentOptionsModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(analyzeDocumentOptionsModel.FileContentType).To(Equal(core.StringPtr("application/json")))
				Expect(analyzeDocumentOptionsModel.Metadata).To(Equal(core.StringPtr("testString")))
				Expect(analyzeDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCollectionDetails successfully`, func() {
				name := "testString"
				_model, err := discoveryService.NewCollectionDetails(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCollectionOptions successfully`, func() {
				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				Expect(collectionEnrichmentModel).ToNot(BeNil())
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}
				Expect(collectionEnrichmentModel.EnrichmentID).To(Equal(core.StringPtr("testString")))
				Expect(collectionEnrichmentModel.Fields).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateCollectionOptions model
				projectID := "testString"
				createCollectionOptionsName := "testString"
				createCollectionOptionsModel := discoveryService.NewCreateCollectionOptions(projectID, createCollectionOptionsName)
				createCollectionOptionsModel.SetProjectID("testString")
				createCollectionOptionsModel.SetName("testString")
				createCollectionOptionsModel.SetDescription("testString")
				createCollectionOptionsModel.SetLanguage("en")
				createCollectionOptionsModel.SetEnrichments([]discoveryv2.CollectionEnrichment{*collectionEnrichmentModel})
				createCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectionOptionsModel).ToNot(BeNil())
				Expect(createCollectionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Language).To(Equal(core.StringPtr("en")))
				Expect(createCollectionOptionsModel.Enrichments).To(Equal([]discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}))
				Expect(createCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEnrichmentOptions successfully`, func() {
				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv2.EnrichmentOptions)
				Expect(enrichmentOptionsModel).ToNot(BeNil())
				enrichmentOptionsModel.Languages = []string{"testString"}
				enrichmentOptionsModel.EntityType = core.StringPtr("testString")
				enrichmentOptionsModel.RegularExpression = core.StringPtr("testString")
				enrichmentOptionsModel.ResultField = core.StringPtr("testString")
				Expect(enrichmentOptionsModel.Languages).To(Equal([]string{"testString"}))
				Expect(enrichmentOptionsModel.EntityType).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentOptionsModel.RegularExpression).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentOptionsModel.ResultField).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateEnrichment model
				createEnrichmentModel := new(discoveryv2.CreateEnrichment)
				Expect(createEnrichmentModel).ToNot(BeNil())
				createEnrichmentModel.Name = core.StringPtr("testString")
				createEnrichmentModel.Description = core.StringPtr("testString")
				createEnrichmentModel.Type = core.StringPtr("dictionary")
				createEnrichmentModel.Options = enrichmentOptionsModel
				Expect(createEnrichmentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createEnrichmentModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createEnrichmentModel.Type).To(Equal(core.StringPtr("dictionary")))
				Expect(createEnrichmentModel.Options).To(Equal(enrichmentOptionsModel))

				// Construct an instance of the CreateEnrichmentOptions model
				projectID := "testString"
				var enrichment *discoveryv2.CreateEnrichment = nil
				createEnrichmentOptionsModel := discoveryService.NewCreateEnrichmentOptions(projectID, enrichment)
				createEnrichmentOptionsModel.SetProjectID("testString")
				createEnrichmentOptionsModel.SetEnrichment(createEnrichmentModel)
				createEnrichmentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				createEnrichmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEnrichmentOptionsModel).ToNot(BeNil())
				Expect(createEnrichmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(createEnrichmentOptionsModel.Enrichment).To(Equal(createEnrichmentModel))
				Expect(createEnrichmentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createEnrichmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the DefaultQueryParamsPassages model
				defaultQueryParamsPassagesModel := new(discoveryv2.DefaultQueryParamsPassages)
				Expect(defaultQueryParamsPassagesModel).ToNot(BeNil())
				defaultQueryParamsPassagesModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.Fields = []string{"testString"}
				defaultQueryParamsPassagesModel.Characters = core.Int64Ptr(int64(38))
				defaultQueryParamsPassagesModel.PerDocument = core.BoolPtr(true)
				defaultQueryParamsPassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				Expect(defaultQueryParamsPassagesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(defaultQueryParamsPassagesModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(defaultQueryParamsPassagesModel.Fields).To(Equal([]string{"testString"}))
				Expect(defaultQueryParamsPassagesModel.Characters).To(Equal(core.Int64Ptr(int64(38))))
				Expect(defaultQueryParamsPassagesModel.PerDocument).To(Equal(core.BoolPtr(true)))
				Expect(defaultQueryParamsPassagesModel.MaxPerDocument).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the DefaultQueryParamsTableResults model
				defaultQueryParamsTableResultsModel := new(discoveryv2.DefaultQueryParamsTableResults)
				Expect(defaultQueryParamsTableResultsModel).ToNot(BeNil())
				defaultQueryParamsTableResultsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsTableResultsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsTableResultsModel.PerDocument = core.Int64Ptr(int64(38))
				Expect(defaultQueryParamsTableResultsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(defaultQueryParamsTableResultsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(defaultQueryParamsTableResultsModel.PerDocument).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the DefaultQueryParamsSuggestedRefinements model
				defaultQueryParamsSuggestedRefinementsModel := new(discoveryv2.DefaultQueryParamsSuggestedRefinements)
				Expect(defaultQueryParamsSuggestedRefinementsModel).ToNot(BeNil())
				defaultQueryParamsSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				defaultQueryParamsSuggestedRefinementsModel.Count = core.Int64Ptr(int64(38))
				Expect(defaultQueryParamsSuggestedRefinementsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(defaultQueryParamsSuggestedRefinementsModel.Count).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the DefaultQueryParams model
				defaultQueryParamsModel := new(discoveryv2.DefaultQueryParams)
				Expect(defaultQueryParamsModel).ToNot(BeNil())
				defaultQueryParamsModel.CollectionIds = []string{"testString"}
				defaultQueryParamsModel.Passages = defaultQueryParamsPassagesModel
				defaultQueryParamsModel.TableResults = defaultQueryParamsTableResultsModel
				defaultQueryParamsModel.Aggregation = core.StringPtr("testString")
				defaultQueryParamsModel.SuggestedRefinements = defaultQueryParamsSuggestedRefinementsModel
				defaultQueryParamsModel.SpellingSuggestions = core.BoolPtr(true)
				defaultQueryParamsModel.Highlight = core.BoolPtr(true)
				defaultQueryParamsModel.Count = core.Int64Ptr(int64(38))
				defaultQueryParamsModel.Sort = core.StringPtr("testString")
				defaultQueryParamsModel.Return = []string{"testString"}
				Expect(defaultQueryParamsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(defaultQueryParamsModel.Passages).To(Equal(defaultQueryParamsPassagesModel))
				Expect(defaultQueryParamsModel.TableResults).To(Equal(defaultQueryParamsTableResultsModel))
				Expect(defaultQueryParamsModel.Aggregation).To(Equal(core.StringPtr("testString")))
				Expect(defaultQueryParamsModel.SuggestedRefinements).To(Equal(defaultQueryParamsSuggestedRefinementsModel))
				Expect(defaultQueryParamsModel.SpellingSuggestions).To(Equal(core.BoolPtr(true)))
				Expect(defaultQueryParamsModel.Highlight).To(Equal(core.BoolPtr(true)))
				Expect(defaultQueryParamsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(defaultQueryParamsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(defaultQueryParamsModel.Return).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsName := "testString"
				createProjectOptionsType := "document_retrieval"
				createProjectOptionsModel := discoveryService.NewCreateProjectOptions(createProjectOptionsName, createProjectOptionsType)
				createProjectOptionsModel.SetName("testString")
				createProjectOptionsModel.SetType("document_retrieval")
				createProjectOptionsModel.SetDefaultQueryParameters(defaultQueryParamsModel)
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProjectOptionsModel.Type).To(Equal(core.StringPtr("document_retrieval")))
				Expect(createProjectOptionsModel.DefaultQueryParameters).To(Equal(defaultQueryParamsModel))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTrainingQueryOptions successfully`, func() {
				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				Expect(trainingExampleModel).ToNot(BeNil())
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))
				Expect(trainingExampleModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(trainingExampleModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(trainingExampleModel.Relevance).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the CreateTrainingQueryOptions model
				projectID := "testString"
				createTrainingQueryOptionsNaturalLanguageQuery := "testString"
				createTrainingQueryOptionsExamples := []discoveryv2.TrainingExample{}
				createTrainingQueryOptionsModel := discoveryService.NewCreateTrainingQueryOptions(projectID, createTrainingQueryOptionsNaturalLanguageQuery, createTrainingQueryOptionsExamples)
				createTrainingQueryOptionsModel.SetProjectID("testString")
				createTrainingQueryOptionsModel.SetNaturalLanguageQuery("testString")
				createTrainingQueryOptionsModel.SetExamples([]discoveryv2.TrainingExample{*trainingExampleModel})
				createTrainingQueryOptionsModel.SetFilter("testString")
				createTrainingQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTrainingQueryOptionsModel).ToNot(BeNil())
				Expect(createTrainingQueryOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingQueryOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingQueryOptionsModel.Examples).To(Equal([]discoveryv2.TrainingExample{*trainingExampleModel}))
				Expect(createTrainingQueryOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCollectionOptions successfully`, func() {
				// Construct an instance of the DeleteCollectionOptions model
				projectID := "testString"
				collectionID := "testString"
				deleteCollectionOptionsModel := discoveryService.NewDeleteCollectionOptions(projectID, collectionID)
				deleteCollectionOptionsModel.SetProjectID("testString")
				deleteCollectionOptionsModel.SetCollectionID("testString")
				deleteCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCollectionOptionsModel).ToNot(BeNil())
				Expect(deleteCollectionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDocumentOptions model
				projectID := "testString"
				collectionID := "testString"
				documentID := "testString"
				deleteDocumentOptionsModel := discoveryService.NewDeleteDocumentOptions(projectID, collectionID, documentID)
				deleteDocumentOptionsModel.SetProjectID("testString")
				deleteDocumentOptionsModel.SetCollectionID("testString")
				deleteDocumentOptionsModel.SetDocumentID("testString")
				deleteDocumentOptionsModel.SetXWatsonDiscoveryForce(false)
				deleteDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDocumentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.XWatsonDiscoveryForce).To(Equal(core.BoolPtr(false)))
				Expect(deleteDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEnrichmentOptions successfully`, func() {
				// Construct an instance of the DeleteEnrichmentOptions model
				projectID := "testString"
				enrichmentID := "testString"
				deleteEnrichmentOptionsModel := discoveryService.NewDeleteEnrichmentOptions(projectID, enrichmentID)
				deleteEnrichmentOptionsModel.SetProjectID("testString")
				deleteEnrichmentOptionsModel.SetEnrichmentID("testString")
				deleteEnrichmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEnrichmentOptionsModel).ToNot(BeNil())
				Expect(deleteEnrichmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEnrichmentOptionsModel.EnrichmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEnrichmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				projectID := "testString"
				deleteProjectOptionsModel := discoveryService.NewDeleteProjectOptions(projectID)
				deleteProjectOptionsModel.SetProjectID("testString")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTrainingQueriesOptions successfully`, func() {
				// Construct an instance of the DeleteTrainingQueriesOptions model
				projectID := "testString"
				deleteTrainingQueriesOptionsModel := discoveryService.NewDeleteTrainingQueriesOptions(projectID)
				deleteTrainingQueriesOptionsModel.SetProjectID("testString")
				deleteTrainingQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTrainingQueriesOptionsModel).ToNot(BeNil())
				Expect(deleteTrainingQueriesOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTrainingQueryOptions successfully`, func() {
				// Construct an instance of the DeleteTrainingQueryOptions model
				projectID := "testString"
				queryID := "testString"
				deleteTrainingQueryOptionsModel := discoveryService.NewDeleteTrainingQueryOptions(projectID, queryID)
				deleteTrainingQueryOptionsModel.SetProjectID("testString")
				deleteTrainingQueryOptionsModel.SetQueryID("testString")
				deleteTrainingQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTrainingQueryOptionsModel).ToNot(BeNil())
				Expect(deleteTrainingQueryOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingQueryOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteUserDataOptions successfully`, func() {
				// Construct an instance of the DeleteUserDataOptions model
				customerID := "testString"
				deleteUserDataOptionsModel := discoveryService.NewDeleteUserDataOptions(customerID)
				deleteUserDataOptionsModel.SetCustomerID("testString")
				deleteUserDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteUserDataOptionsModel).ToNot(BeNil())
				Expect(deleteUserDataOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteUserDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAutocompletionOptions successfully`, func() {
				// Construct an instance of the GetAutocompletionOptions model
				projectID := "testString"
				prefix := "testString"
				getAutocompletionOptionsModel := discoveryService.NewGetAutocompletionOptions(projectID, prefix)
				getAutocompletionOptionsModel.SetProjectID("testString")
				getAutocompletionOptionsModel.SetPrefix("testString")
				getAutocompletionOptionsModel.SetCollectionIds([]string{"testString"})
				getAutocompletionOptionsModel.SetField("testString")
				getAutocompletionOptionsModel.SetCount(int64(38))
				getAutocompletionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAutocompletionOptionsModel).ToNot(BeNil())
				Expect(getAutocompletionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.Prefix).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(getAutocompletionOptionsModel.Field).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getAutocompletionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCollectionOptions successfully`, func() {
				// Construct an instance of the GetCollectionOptions model
				projectID := "testString"
				collectionID := "testString"
				getCollectionOptionsModel := discoveryService.NewGetCollectionOptions(projectID, collectionID)
				getCollectionOptionsModel.SetProjectID("testString")
				getCollectionOptionsModel.SetCollectionID("testString")
				getCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionOptionsModel).ToNot(BeNil())
				Expect(getCollectionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentSettingsOptions successfully`, func() {
				// Construct an instance of the GetComponentSettingsOptions model
				projectID := "testString"
				getComponentSettingsOptionsModel := discoveryService.NewGetComponentSettingsOptions(projectID)
				getComponentSettingsOptionsModel.SetProjectID("testString")
				getComponentSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentSettingsOptionsModel).ToNot(BeNil())
				Expect(getComponentSettingsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getComponentSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnrichmentOptions successfully`, func() {
				// Construct an instance of the GetEnrichmentOptions model
				projectID := "testString"
				enrichmentID := "testString"
				getEnrichmentOptionsModel := discoveryService.NewGetEnrichmentOptions(projectID, enrichmentID)
				getEnrichmentOptionsModel.SetProjectID("testString")
				getEnrichmentOptionsModel.SetEnrichmentID("testString")
				getEnrichmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnrichmentOptionsModel).ToNot(BeNil())
				Expect(getEnrichmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getEnrichmentOptionsModel.EnrichmentID).To(Equal(core.StringPtr("testString")))
				Expect(getEnrichmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				projectID := "testString"
				getProjectOptionsModel := discoveryService.NewGetProjectOptions(projectID)
				getProjectOptionsModel.SetProjectID("testString")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTrainingQueryOptions successfully`, func() {
				// Construct an instance of the GetTrainingQueryOptions model
				projectID := "testString"
				queryID := "testString"
				getTrainingQueryOptionsModel := discoveryService.NewGetTrainingQueryOptions(projectID, queryID)
				getTrainingQueryOptionsModel.SetProjectID("testString")
				getTrainingQueryOptionsModel.SetQueryID("testString")
				getTrainingQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTrainingQueryOptionsModel).ToNot(BeNil())
				Expect(getTrainingQueryOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingQueryOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCollectionsOptions successfully`, func() {
				// Construct an instance of the ListCollectionsOptions model
				projectID := "testString"
				listCollectionsOptionsModel := discoveryService.NewListCollectionsOptions(projectID)
				listCollectionsOptionsModel.SetProjectID("testString")
				listCollectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectionsOptionsModel).ToNot(BeNil())
				Expect(listCollectionsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listCollectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnrichmentsOptions successfully`, func() {
				// Construct an instance of the ListEnrichmentsOptions model
				projectID := "testString"
				listEnrichmentsOptionsModel := discoveryService.NewListEnrichmentsOptions(projectID)
				listEnrichmentsOptionsModel.SetProjectID("testString")
				listEnrichmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnrichmentsOptionsModel).ToNot(BeNil())
				Expect(listEnrichmentsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listEnrichmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFieldsOptions successfully`, func() {
				// Construct an instance of the ListFieldsOptions model
				projectID := "testString"
				listFieldsOptionsModel := discoveryService.NewListFieldsOptions(projectID)
				listFieldsOptionsModel.SetProjectID("testString")
				listFieldsOptionsModel.SetCollectionIds([]string{"testString"})
				listFieldsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFieldsOptionsModel).ToNot(BeNil())
				Expect(listFieldsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listFieldsOptionsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(listFieldsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := discoveryService.NewListProjectsOptions()
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTrainingQueriesOptions successfully`, func() {
				// Construct an instance of the ListTrainingQueriesOptions model
				projectID := "testString"
				listTrainingQueriesOptionsModel := discoveryService.NewListTrainingQueriesOptions(projectID)
				listTrainingQueriesOptionsModel.SetProjectID("testString")
				listTrainingQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTrainingQueriesOptionsModel).ToNot(BeNil())
				Expect(listTrainingQueriesOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listTrainingQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryCollectionNoticesOptions successfully`, func() {
				// Construct an instance of the QueryCollectionNoticesOptions model
				projectID := "testString"
				collectionID := "testString"
				queryCollectionNoticesOptionsModel := discoveryService.NewQueryCollectionNoticesOptions(projectID, collectionID)
				queryCollectionNoticesOptionsModel.SetProjectID("testString")
				queryCollectionNoticesOptionsModel.SetCollectionID("testString")
				queryCollectionNoticesOptionsModel.SetFilter("testString")
				queryCollectionNoticesOptionsModel.SetQuery("testString")
				queryCollectionNoticesOptionsModel.SetNaturalLanguageQuery("testString")
				queryCollectionNoticesOptionsModel.SetCount(int64(38))
				queryCollectionNoticesOptionsModel.SetOffset(int64(38))
				queryCollectionNoticesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryCollectionNoticesOptionsModel).ToNot(BeNil())
				Expect(queryCollectionNoticesOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(queryCollectionNoticesOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(queryCollectionNoticesOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(queryCollectionNoticesOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(queryCollectionNoticesOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(queryCollectionNoticesOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryCollectionNoticesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryCollectionNoticesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryNoticesOptions successfully`, func() {
				// Construct an instance of the QueryNoticesOptions model
				projectID := "testString"
				queryNoticesOptionsModel := discoveryService.NewQueryNoticesOptions(projectID)
				queryNoticesOptionsModel.SetProjectID("testString")
				queryNoticesOptionsModel.SetFilter("testString")
				queryNoticesOptionsModel.SetQuery("testString")
				queryNoticesOptionsModel.SetNaturalLanguageQuery("testString")
				queryNoticesOptionsModel.SetCount(int64(38))
				queryNoticesOptionsModel.SetOffset(int64(38))
				queryNoticesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryNoticesOptionsModel).ToNot(BeNil())
				Expect(queryNoticesOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryNoticesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryNoticesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryOptions successfully`, func() {
				// Construct an instance of the QueryLargeTableResults model
				queryLargeTableResultsModel := new(discoveryv2.QueryLargeTableResults)
				Expect(queryLargeTableResultsModel).ToNot(BeNil())
				queryLargeTableResultsModel.Enabled = core.BoolPtr(true)
				queryLargeTableResultsModel.Count = core.Int64Ptr(int64(38))
				Expect(queryLargeTableResultsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(queryLargeTableResultsModel.Count).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the QueryLargeSuggestedRefinements model
				queryLargeSuggestedRefinementsModel := new(discoveryv2.QueryLargeSuggestedRefinements)
				Expect(queryLargeSuggestedRefinementsModel).ToNot(BeNil())
				queryLargeSuggestedRefinementsModel.Enabled = core.BoolPtr(true)
				queryLargeSuggestedRefinementsModel.Count = core.Int64Ptr(int64(1))
				Expect(queryLargeSuggestedRefinementsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(queryLargeSuggestedRefinementsModel.Count).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the QueryLargePassages model
				queryLargePassagesModel := new(discoveryv2.QueryLargePassages)
				Expect(queryLargePassagesModel).ToNot(BeNil())
				queryLargePassagesModel.Enabled = core.BoolPtr(true)
				queryLargePassagesModel.PerDocument = core.BoolPtr(true)
				queryLargePassagesModel.MaxPerDocument = core.Int64Ptr(int64(38))
				queryLargePassagesModel.Fields = []string{"testString"}
				queryLargePassagesModel.Count = core.Int64Ptr(int64(400))
				queryLargePassagesModel.Characters = core.Int64Ptr(int64(50))
				queryLargePassagesModel.FindAnswers = core.BoolPtr(false)
				queryLargePassagesModel.MaxAnswersPerPassage = core.Int64Ptr(int64(38))
				Expect(queryLargePassagesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(queryLargePassagesModel.PerDocument).To(Equal(core.BoolPtr(true)))
				Expect(queryLargePassagesModel.MaxPerDocument).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryLargePassagesModel.Fields).To(Equal([]string{"testString"}))
				Expect(queryLargePassagesModel.Count).To(Equal(core.Int64Ptr(int64(400))))
				Expect(queryLargePassagesModel.Characters).To(Equal(core.Int64Ptr(int64(50))))
				Expect(queryLargePassagesModel.FindAnswers).To(Equal(core.BoolPtr(false)))
				Expect(queryLargePassagesModel.MaxAnswersPerPassage).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the QueryOptions model
				projectID := "testString"
				queryOptionsModel := discoveryService.NewQueryOptions(projectID)
				queryOptionsModel.SetProjectID("testString")
				queryOptionsModel.SetCollectionIds([]string{"testString"})
				queryOptionsModel.SetFilter("testString")
				queryOptionsModel.SetQuery("testString")
				queryOptionsModel.SetNaturalLanguageQuery("testString")
				queryOptionsModel.SetAggregation("testString")
				queryOptionsModel.SetCount(int64(38))
				queryOptionsModel.SetReturn([]string{"testString"})
				queryOptionsModel.SetOffset(int64(38))
				queryOptionsModel.SetSort("testString")
				queryOptionsModel.SetHighlight(true)
				queryOptionsModel.SetSpellingSuggestions(true)
				queryOptionsModel.SetTableResults(queryLargeTableResultsModel)
				queryOptionsModel.SetSuggestedRefinements(queryLargeSuggestedRefinementsModel)
				queryOptionsModel.SetPassages(queryLargePassagesModel)
				queryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryOptionsModel).ToNot(BeNil())
				Expect(queryOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(queryOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Aggregation).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryOptionsModel.Return).To(Equal([]string{"testString"}))
				Expect(queryOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Highlight).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.SpellingSuggestions).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.TableResults).To(Equal(queryLargeTableResultsModel))
				Expect(queryOptionsModel.SuggestedRefinements).To(Equal(queryLargeSuggestedRefinementsModel))
				Expect(queryOptionsModel.Passages).To(Equal(queryLargePassagesModel))
				Expect(queryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTrainingExample successfully`, func() {
				documentID := "testString"
				collectionID := "testString"
				relevance := int64(38)
				_model, err := discoveryService.NewTrainingExample(documentID, collectionID, relevance)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTrainingQuery successfully`, func() {
				naturalLanguageQuery := "testString"
				examples := []discoveryv2.TrainingExample{}
				_model, err := discoveryService.NewTrainingQuery(naturalLanguageQuery, examples)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCollectionOptions successfully`, func() {
				// Construct an instance of the CollectionEnrichment model
				collectionEnrichmentModel := new(discoveryv2.CollectionEnrichment)
				Expect(collectionEnrichmentModel).ToNot(BeNil())
				collectionEnrichmentModel.EnrichmentID = core.StringPtr("testString")
				collectionEnrichmentModel.Fields = []string{"testString"}
				Expect(collectionEnrichmentModel.EnrichmentID).To(Equal(core.StringPtr("testString")))
				Expect(collectionEnrichmentModel.Fields).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateCollectionOptions model
				projectID := "testString"
				collectionID := "testString"
				updateCollectionOptionsModel := discoveryService.NewUpdateCollectionOptions(projectID, collectionID)
				updateCollectionOptionsModel.SetProjectID("testString")
				updateCollectionOptionsModel.SetCollectionID("testString")
				updateCollectionOptionsModel.SetName("testString")
				updateCollectionOptionsModel.SetDescription("testString")
				updateCollectionOptionsModel.SetEnrichments([]discoveryv2.CollectionEnrichment{*collectionEnrichmentModel})
				updateCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectionOptionsModel).ToNot(BeNil())
				Expect(updateCollectionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Enrichments).To(Equal([]discoveryv2.CollectionEnrichment{*collectionEnrichmentModel}))
				Expect(updateCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDocumentOptions successfully`, func() {
				// Construct an instance of the UpdateDocumentOptions model
				projectID := "testString"
				collectionID := "testString"
				documentID := "testString"
				updateDocumentOptionsModel := discoveryService.NewUpdateDocumentOptions(projectID, collectionID, documentID)
				updateDocumentOptionsModel.SetProjectID("testString")
				updateDocumentOptionsModel.SetCollectionID("testString")
				updateDocumentOptionsModel.SetDocumentID("testString")
				updateDocumentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				updateDocumentOptionsModel.SetFilename("testString")
				updateDocumentOptionsModel.SetFileContentType("application/json")
				updateDocumentOptionsModel.SetMetadata("testString")
				updateDocumentOptionsModel.SetXWatsonDiscoveryForce(false)
				updateDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDocumentOptionsModel).ToNot(BeNil())
				Expect(updateDocumentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDocumentOptionsModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.FileContentType).To(Equal(core.StringPtr("application/json")))
				Expect(updateDocumentOptionsModel.Metadata).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.XWatsonDiscoveryForce).To(Equal(core.BoolPtr(false)))
				Expect(updateDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnrichmentOptions successfully`, func() {
				// Construct an instance of the UpdateEnrichmentOptions model
				projectID := "testString"
				enrichmentID := "testString"
				updateEnrichmentOptionsName := "testString"
				updateEnrichmentOptionsModel := discoveryService.NewUpdateEnrichmentOptions(projectID, enrichmentID, updateEnrichmentOptionsName)
				updateEnrichmentOptionsModel.SetProjectID("testString")
				updateEnrichmentOptionsModel.SetEnrichmentID("testString")
				updateEnrichmentOptionsModel.SetName("testString")
				updateEnrichmentOptionsModel.SetDescription("testString")
				updateEnrichmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnrichmentOptionsModel).ToNot(BeNil())
				Expect(updateEnrichmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnrichmentOptionsModel.EnrichmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnrichmentOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateEnrichmentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateEnrichmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectOptions successfully`, func() {
				// Construct an instance of the UpdateProjectOptions model
				projectID := "testString"
				updateProjectOptionsModel := discoveryService.NewUpdateProjectOptions(projectID)
				updateProjectOptionsModel.SetProjectID("testString")
				updateProjectOptionsModel.SetName("testString")
				updateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectOptionsModel).ToNot(BeNil())
				Expect(updateProjectOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTrainingQueryOptions successfully`, func() {
				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv2.TrainingExample)
				Expect(trainingExampleModel).ToNot(BeNil())
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CollectionID = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))
				Expect(trainingExampleModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(trainingExampleModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(trainingExampleModel.Relevance).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the UpdateTrainingQueryOptions model
				projectID := "testString"
				queryID := "testString"
				updateTrainingQueryOptionsNaturalLanguageQuery := "testString"
				updateTrainingQueryOptionsExamples := []discoveryv2.TrainingExample{}
				updateTrainingQueryOptionsModel := discoveryService.NewUpdateTrainingQueryOptions(projectID, queryID, updateTrainingQueryOptionsNaturalLanguageQuery, updateTrainingQueryOptionsExamples)
				updateTrainingQueryOptionsModel.SetProjectID("testString")
				updateTrainingQueryOptionsModel.SetQueryID("testString")
				updateTrainingQueryOptionsModel.SetNaturalLanguageQuery("testString")
				updateTrainingQueryOptionsModel.SetExamples([]discoveryv2.TrainingExample{*trainingExampleModel})
				updateTrainingQueryOptionsModel.SetFilter("testString")
				updateTrainingQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTrainingQueryOptionsModel).ToNot(BeNil())
				Expect(updateTrainingQueryOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingQueryOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingQueryOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingQueryOptionsModel.Examples).To(Equal([]discoveryv2.TrainingExample{*trainingExampleModel}))
				Expect(updateTrainingQueryOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
