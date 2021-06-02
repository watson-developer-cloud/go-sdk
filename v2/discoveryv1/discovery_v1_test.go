/**
 * (C) Copyright IBM Corp. 2021.
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
	"github.com/watson-developer-cloud/go-sdk/v2/discoveryv1"
)

var _ = Describe(`DiscoveryV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
			discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(discoveryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
				URL:     "https://discoveryv1/api",
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
			discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{})
			Expect(discoveryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISCOVERY_URL":       "https://discoveryv1/api",
				"DISCOVERY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				"DISCOVERY_URL":       "https://discoveryv1/api",
				"DISCOVERY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
			discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
			url, err = discoveryv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions) - Operation response error`, func() {
		version := "testString"
		createEnvironmentPath := "/v1/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEnvironment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(discoveryv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("testString")
				createEnvironmentOptionsModel.Description = core.StringPtr("testString")
				createEnvironmentOptionsModel.Size = core.StringPtr("LT")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions)`, func() {
		version := "testString"
		createEnvironmentPath := "/v1/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
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
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}`)
				}))
			})
			It(`Invoke CreateEnvironment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(discoveryv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("testString")
				createEnvironmentOptionsModel.Description = core.StringPtr("testString")
				createEnvironmentOptionsModel.Size = core.StringPtr("LT")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateEnvironmentWithContext(ctx, createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateEnvironmentWithContext(ctx, createEnvironmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
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
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}`)
				}))
			})
			It(`Invoke CreateEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(discoveryv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("testString")
				createEnvironmentOptionsModel.Description = core.StringPtr("testString")
				createEnvironmentOptionsModel.Size = core.StringPtr("LT")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEnvironment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(discoveryv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("testString")
				createEnvironmentOptionsModel.Description = core.StringPtr("testString")
				createEnvironmentOptionsModel.Size = core.StringPtr("LT")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEnvironmentOptions model with no property values
				createEnvironmentOptionsModelNew := new(discoveryv1.CreateEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateEnvironment(createEnvironmentOptionsModelNew)
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
			It(`Invoke CreateEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(discoveryv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("testString")
				createEnvironmentOptionsModel.Description = core.StringPtr("testString")
				createEnvironmentOptionsModel.Size = core.StringPtr("LT")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateEnvironment(createEnvironmentOptionsModel)
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
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) - Operation response error`, func() {
		version := "testString"
		listEnvironmentsPath := "/v1/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEnvironments with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(discoveryv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.StringPtr("testString")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions)`, func() {
		version := "testString"
		listEnvironmentsPath := "/v1/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}]}`)
				}))
			})
			It(`Invoke ListEnvironments successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(discoveryv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.StringPtr("testString")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}]}`)
				}))
			})
			It(`Invoke ListEnvironments successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListEnvironments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(discoveryv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.StringPtr("testString")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEnvironments with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(discoveryv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.StringPtr("testString")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListEnvironments(listEnvironmentsOptionsModel)
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
			It(`Invoke ListEnvironments successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(discoveryv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.StringPtr("testString")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListEnvironments(listEnvironmentsOptionsModel)
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
	Describe(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions) - Operation response error`, func() {
		version := "testString"
		getEnvironmentPath := "/v1/environments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnvironment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(discoveryv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions)`, func() {
		version := "testString"
		getEnvironmentPath := "/v1/environments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}`)
				}))
			})
			It(`Invoke GetEnvironment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(discoveryv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetEnvironmentWithContext(ctx, getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetEnvironmentWithContext(ctx, getEnvironmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}`)
				}))
			})
			It(`Invoke GetEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(discoveryv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEnvironment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(discoveryv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEnvironmentOptions model with no property values
				getEnvironmentOptionsModelNew := new(discoveryv1.GetEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetEnvironment(getEnvironmentOptionsModelNew)
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
			It(`Invoke GetEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(discoveryv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetEnvironment(getEnvironmentOptionsModel)
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
	Describe(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions) - Operation response error`, func() {
		version := "testString"
		updateEnvironmentPath := "/v1/environments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEnvironment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(discoveryv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Size = core.StringPtr("S")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions)`, func() {
		version := "testString"
		updateEnvironmentPath := "/v1/environments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}`)
				}))
			})
			It(`Invoke UpdateEnvironment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(discoveryv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Size = core.StringPtr("S")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateEnvironmentWithContext(ctx, updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateEnvironmentWithContext(ctx, updateEnvironmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "read_only": true, "size": "LT", "requested_size": "RequestedSize", "index_capacity": {"documents": {"available": 9, "maximum_allowed": 14}, "disk_usage": {"used_bytes": 9, "maximum_allowed_bytes": 19}, "collections": {"available": 9, "maximum_allowed": 14}}, "search_status": {"scope": "Scope", "status": "NO_DATA", "status_description": "StatusDescription", "last_trained": "2019-01-01"}}`)
				}))
			})
			It(`Invoke UpdateEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(discoveryv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Size = core.StringPtr("S")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEnvironment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(discoveryv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Size = core.StringPtr("S")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEnvironmentOptions model with no property values
				updateEnvironmentOptionsModelNew := new(discoveryv1.UpdateEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateEnvironment(updateEnvironmentOptionsModelNew)
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
			It(`Invoke UpdateEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(discoveryv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Size = core.StringPtr("S")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateEnvironment(updateEnvironmentOptionsModel)
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
	Describe(`DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions) - Operation response error`, func() {
		version := "testString"
		deleteEnvironmentPath := "/v1/environments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteEnvironment with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(discoveryv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions)`, func() {
		version := "testString"
		deleteEnvironmentPath := "/v1/environments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteEnvironment successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(discoveryv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.DeleteEnvironmentWithContext(ctx, deleteEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.DeleteEnvironmentWithContext(ctx, deleteEnvironmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.DeleteEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(discoveryv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteEnvironment with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(discoveryv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteEnvironmentOptions model with no property values
				deleteEnvironmentOptionsModelNew := new(discoveryv1.DeleteEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModelNew)
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
			It(`Invoke DeleteEnvironment successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(discoveryv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.DeleteEnvironment(deleteEnvironmentOptionsModel)
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
		listFieldsPath := "/v1/environments/testString/fields"
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv1.ListFieldsOptions)
				listFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
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
		listFieldsPath := "/v1/environments/testString/fields"
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
					fmt.Fprintf(res, "%s", `{"fields": [{"field": "Field", "type": "nested"}]}`)
				}))
			})
			It(`Invoke ListFields successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv1.ListFieldsOptions)
				listFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"fields": [{"field": "Field", "type": "nested"}]}`)
				}))
			})
			It(`Invoke ListFields successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				listFieldsOptionsModel := new(discoveryv1.ListFieldsOptions)
				listFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listFieldsOptionsModel.CollectionIds = []string{"testString"}
				listFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListFields(listFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListFields with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv1.ListFieldsOptions)
				listFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
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
				listFieldsOptionsModelNew := new(discoveryv1.ListFieldsOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListFieldsOptions model
				listFieldsOptionsModel := new(discoveryv1.ListFieldsOptions)
				listFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
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
	Describe(`CreateConfiguration(createConfigurationOptions *CreateConfigurationOptions) - Operation response error`, func() {
		version := "testString"
		createConfigurationPath := "/v1/environments/testString/configurations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigurationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateConfiguration with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the CreateConfigurationOptions model
				createConfigurationOptionsModel := new(discoveryv1.CreateConfigurationOptions)
				createConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				createConfigurationOptionsModel.Name = core.StringPtr("testString")
				createConfigurationOptionsModel.Description = core.StringPtr("testString")
				createConfigurationOptionsModel.Conversions = conversionsModel
				createConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				createConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				createConfigurationOptionsModel.Source = sourceModel
				createConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateConfiguration(createConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateConfiguration(createConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfiguration(createConfigurationOptions *CreateConfigurationOptions)`, func() {
		version := "testString"
		createConfigurationPath := "/v1/environments/testString/configurations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigurationPath))
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
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}`)
				}))
			})
			It(`Invoke CreateConfiguration successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the CreateConfigurationOptions model
				createConfigurationOptionsModel := new(discoveryv1.CreateConfigurationOptions)
				createConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				createConfigurationOptionsModel.Name = core.StringPtr("testString")
				createConfigurationOptionsModel.Description = core.StringPtr("testString")
				createConfigurationOptionsModel.Conversions = conversionsModel
				createConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				createConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				createConfigurationOptionsModel.Source = sourceModel
				createConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateConfigurationWithContext(ctx, createConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateConfiguration(createConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateConfigurationWithContext(ctx, createConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createConfigurationPath))
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
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}`)
				}))
			})
			It(`Invoke CreateConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the CreateConfigurationOptions model
				createConfigurationOptionsModel := new(discoveryv1.CreateConfigurationOptions)
				createConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				createConfigurationOptionsModel.Name = core.StringPtr("testString")
				createConfigurationOptionsModel.Description = core.StringPtr("testString")
				createConfigurationOptionsModel.Conversions = conversionsModel
				createConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				createConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				createConfigurationOptionsModel.Source = sourceModel
				createConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateConfiguration(createConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfiguration with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the CreateConfigurationOptions model
				createConfigurationOptionsModel := new(discoveryv1.CreateConfigurationOptions)
				createConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				createConfigurationOptionsModel.Name = core.StringPtr("testString")
				createConfigurationOptionsModel.Description = core.StringPtr("testString")
				createConfigurationOptionsModel.Conversions = conversionsModel
				createConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				createConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				createConfigurationOptionsModel.Source = sourceModel
				createConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateConfiguration(createConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigurationOptions model with no property values
				createConfigurationOptionsModelNew := new(discoveryv1.CreateConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateConfiguration(createConfigurationOptionsModelNew)
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
			It(`Invoke CreateConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the CreateConfigurationOptions model
				createConfigurationOptionsModel := new(discoveryv1.CreateConfigurationOptions)
				createConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				createConfigurationOptionsModel.Name = core.StringPtr("testString")
				createConfigurationOptionsModel.Description = core.StringPtr("testString")
				createConfigurationOptionsModel.Conversions = conversionsModel
				createConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				createConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				createConfigurationOptionsModel.Source = sourceModel
				createConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateConfiguration(createConfigurationOptionsModel)
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
	Describe(`ListConfigurations(listConfigurationsOptions *ListConfigurationsOptions) - Operation response error`, func() {
		version := "testString"
		listConfigurationsPath := "/v1/environments/testString/configurations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigurationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigurations with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListConfigurationsOptions model
				listConfigurationsOptionsModel := new(discoveryv1.ListConfigurationsOptions)
				listConfigurationsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listConfigurationsOptionsModel.Name = core.StringPtr("testString")
				listConfigurationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListConfigurations(listConfigurationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListConfigurations(listConfigurationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigurations(listConfigurationsOptions *ListConfigurationsOptions)`, func() {
		version := "testString"
		listConfigurationsPath := "/v1/environments/testString/configurations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigurationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configurations": [{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}]}`)
				}))
			})
			It(`Invoke ListConfigurations successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigurationsOptions model
				listConfigurationsOptionsModel := new(discoveryv1.ListConfigurationsOptions)
				listConfigurationsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listConfigurationsOptionsModel.Name = core.StringPtr("testString")
				listConfigurationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListConfigurationsWithContext(ctx, listConfigurationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListConfigurations(listConfigurationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListConfigurationsWithContext(ctx, listConfigurationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listConfigurationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configurations": [{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}]}`)
				}))
			})
			It(`Invoke ListConfigurations successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListConfigurations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigurationsOptions model
				listConfigurationsOptionsModel := new(discoveryv1.ListConfigurationsOptions)
				listConfigurationsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listConfigurationsOptionsModel.Name = core.StringPtr("testString")
				listConfigurationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListConfigurations(listConfigurationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigurations with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListConfigurationsOptions model
				listConfigurationsOptionsModel := new(discoveryv1.ListConfigurationsOptions)
				listConfigurationsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listConfigurationsOptionsModel.Name = core.StringPtr("testString")
				listConfigurationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListConfigurations(listConfigurationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigurationsOptions model with no property values
				listConfigurationsOptionsModelNew := new(discoveryv1.ListConfigurationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListConfigurations(listConfigurationsOptionsModelNew)
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
			It(`Invoke ListConfigurations successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListConfigurationsOptions model
				listConfigurationsOptionsModel := new(discoveryv1.ListConfigurationsOptions)
				listConfigurationsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listConfigurationsOptionsModel.Name = core.StringPtr("testString")
				listConfigurationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListConfigurations(listConfigurationsOptionsModel)
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
	Describe(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions) - Operation response error`, func() {
		version := "testString"
		getConfigurationPath := "/v1/environments/testString/configurations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfiguration with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(discoveryv1.GetConfigurationOptions)
				getConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				getConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions)`, func() {
		version := "testString"
		getConfigurationPath := "/v1/environments/testString/configurations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}`)
				}))
			})
			It(`Invoke GetConfiguration successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(discoveryv1.GetConfigurationOptions)
				getConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				getConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetConfigurationWithContext(ctx, getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetConfigurationWithContext(ctx, getConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}`)
				}))
			})
			It(`Invoke GetConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(discoveryv1.GetConfigurationOptions)
				getConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				getConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfiguration with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(discoveryv1.GetConfigurationOptions)
				getConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				getConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigurationOptions model with no property values
				getConfigurationOptionsModelNew := new(discoveryv1.GetConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetConfiguration(getConfigurationOptionsModelNew)
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
			It(`Invoke GetConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(discoveryv1.GetConfigurationOptions)
				getConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				getConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetConfiguration(getConfigurationOptionsModel)
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
	Describe(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions) - Operation response error`, func() {
		version := "testString"
		updateConfigurationPath := "/v1/environments/testString/configurations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfiguration with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(discoveryv1.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				updateConfigurationOptionsModel.Name = core.StringPtr("testString")
				updateConfigurationOptionsModel.Description = core.StringPtr("testString")
				updateConfigurationOptionsModel.Conversions = conversionsModel
				updateConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				updateConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				updateConfigurationOptionsModel.Source = sourceModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions)`, func() {
		version := "testString"
		updateConfigurationPath := "/v1/environments/testString/configurations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}`)
				}))
			})
			It(`Invoke UpdateConfiguration successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(discoveryv1.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				updateConfigurationOptionsModel.Name = core.StringPtr("testString")
				updateConfigurationOptionsModel.Description = core.StringPtr("testString")
				updateConfigurationOptionsModel.Conversions = conversionsModel
				updateConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				updateConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				updateConfigurationOptionsModel.Source = sourceModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateConfigurationWithContext(ctx, updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateConfigurationWithContext(ctx, updateConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "name": "Name", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "description": "Description", "conversions": {"pdf": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}]}}, "word": {"heading": {"fonts": [{"level": 5, "min_size": 7, "max_size": 7, "bold": true, "italic": true, "name": "Name"}], "styles": [{"level": 5, "names": ["Names"]}]}}, "html": {"exclude_tags_completely": ["ExcludeTagsCompletely"], "exclude_tags_keep_content": ["ExcludeTagsKeepContent"], "keep_content": {"xpaths": ["Xpaths"]}, "exclude_content": {"xpaths": ["Xpaths"]}, "keep_tag_attributes": ["KeepTagAttributes"], "exclude_tag_attributes": ["ExcludeTagAttributes"]}, "segment": {"enabled": false, "selector_tags": ["SelectorTags"], "annotated_fields": ["AnnotatedFields"]}, "json_normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "image_text_recognition": true}, "enrichments": [{"description": "Description", "destination_field": "DestinationField", "source_field": "SourceField", "overwrite": false, "enrichment": "Enrichment", "ignore_downstream_errors": true, "options": {"features": {"keywords": {"sentiment": false, "emotion": false, "limit": 5}, "entities": {"sentiment": false, "emotion": false, "limit": 5, "mentions": true, "mention_types": true, "sentence_locations": false, "model": "Model"}, "sentiment": {"document": true, "targets": ["Target"]}, "emotion": {"document": true, "targets": ["Target"]}, "categories": {"mapKey": "anyValue"}, "semantic_roles": {"entities": true, "keywords": true, "limit": 5}, "relations": {"model": "Model"}, "concepts": {"limit": 5}}, "language": "ar", "model": "Model"}}], "normalizations": [{"operation": "copy", "source_field": "SourceField", "destination_field": "DestinationField"}], "source": {"type": "box", "credential_id": "CredentialID", "schedule": {"enabled": false, "time_zone": "TimeZone", "frequency": "daily"}, "options": {"folders": [{"owner_user_id": "OwnerUserID", "folder_id": "FolderID", "limit": 5}], "objects": [{"name": "Name", "limit": 5}], "site_collections": [{"site_collection_path": "SiteCollectionPath", "limit": 5}], "urls": [{"url": "URL", "limit_to_starting_hosts": true, "crawl_speed": "gentle", "allow_untrusted_certificate": false, "maximum_hops": 11, "request_timeout": 14, "override_robots_txt": false, "blacklist": ["Blacklist"]}], "buckets": [{"name": "Name", "limit": 5}], "crawl_all_buckets": false}}}`)
				}))
			})
			It(`Invoke UpdateConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(discoveryv1.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				updateConfigurationOptionsModel.Name = core.StringPtr("testString")
				updateConfigurationOptionsModel.Description = core.StringPtr("testString")
				updateConfigurationOptionsModel.Conversions = conversionsModel
				updateConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				updateConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				updateConfigurationOptionsModel.Source = sourceModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfiguration with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(discoveryv1.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				updateConfigurationOptionsModel.Name = core.StringPtr("testString")
				updateConfigurationOptionsModel.Description = core.StringPtr("testString")
				updateConfigurationOptionsModel.Conversions = conversionsModel
				updateConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				updateConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				updateConfigurationOptionsModel.Source = sourceModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigurationOptions model with no property values
				updateConfigurationOptionsModelNew := new(discoveryv1.UpdateConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateConfiguration(updateConfigurationOptionsModelNew)
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
			It(`Invoke UpdateConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				pdfSettingsModel.Heading = pdfHeadingDetectionModel

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				wordSettingsModel.Heading = wordHeadingDetectionModel

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				xPathPatternsModel.Xpaths = []string{"testString"}

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(discoveryv1.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				updateConfigurationOptionsModel.Name = core.StringPtr("testString")
				updateConfigurationOptionsModel.Description = core.StringPtr("testString")
				updateConfigurationOptionsModel.Conversions = conversionsModel
				updateConfigurationOptionsModel.Enrichments = []discoveryv1.Enrichment{*enrichmentModel}
				updateConfigurationOptionsModel.Normalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				updateConfigurationOptionsModel.Source = sourceModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateConfiguration(updateConfigurationOptionsModel)
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
	Describe(`DeleteConfiguration(deleteConfigurationOptions *DeleteConfigurationOptions) - Operation response error`, func() {
		version := "testString"
		deleteConfigurationPath := "/v1/environments/testString/configurations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigurationPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteConfiguration with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigurationOptions model
				deleteConfigurationOptionsModel := new(discoveryv1.DeleteConfigurationOptions)
				deleteConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.DeleteConfiguration(deleteConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.DeleteConfiguration(deleteConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfiguration(deleteConfigurationOptions *DeleteConfigurationOptions)`, func() {
		version := "testString"
		deleteConfigurationPath := "/v1/environments/testString/configurations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigurationPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "status": "deleted", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke DeleteConfiguration successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteConfigurationOptions model
				deleteConfigurationOptionsModel := new(discoveryv1.DeleteConfigurationOptions)
				deleteConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.DeleteConfigurationWithContext(ctx, deleteConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.DeleteConfiguration(deleteConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.DeleteConfigurationWithContext(ctx, deleteConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigurationPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configuration_id": "ConfigurationID", "status": "deleted", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke DeleteConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.DeleteConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteConfigurationOptions model
				deleteConfigurationOptionsModel := new(discoveryv1.DeleteConfigurationOptions)
				deleteConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteConfiguration(deleteConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteConfiguration with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigurationOptions model
				deleteConfigurationOptionsModel := new(discoveryv1.DeleteConfigurationOptions)
				deleteConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.DeleteConfiguration(deleteConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteConfigurationOptions model with no property values
				deleteConfigurationOptionsModelNew := new(discoveryv1.DeleteConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.DeleteConfiguration(deleteConfigurationOptionsModelNew)
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
			It(`Invoke DeleteConfiguration successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigurationOptions model
				deleteConfigurationOptionsModel := new(discoveryv1.DeleteConfigurationOptions)
				deleteConfigurationOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.ConfigurationID = core.StringPtr("testString")
				deleteConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.DeleteConfiguration(deleteConfigurationOptionsModel)
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
		createCollectionPath := "/v1/environments/testString/collections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCollection with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv1.CreateCollectionOptions)
				createCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
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
		createCollectionPath := "/v1/environments/testString/collections"
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}`)
				}))
			})
			It(`Invoke CreateCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv1.CreateCollectionOptions)
				createCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}`)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv1.CreateCollectionOptions)
				createCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv1.CreateCollectionOptions)
				createCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
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
				createCollectionOptionsModelNew := new(discoveryv1.CreateCollectionOptions)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(discoveryv1.CreateCollectionOptions)
				createCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCollectionOptionsModel.Name = core.StringPtr("testString")
				createCollectionOptionsModel.Description = core.StringPtr("testString")
				createCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
				createCollectionOptionsModel.Language = core.StringPtr("en")
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
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions) - Operation response error`, func() {
		version := "testString"
		listCollectionsPath := "/v1/environments/testString/collections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCollections with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv1.ListCollectionsOptions)
				listCollectionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionsOptionsModel.Name = core.StringPtr("testString")
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
		listCollectionsPath := "/v1/environments/testString/collections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}]}`)
				}))
			})
			It(`Invoke ListCollections successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv1.ListCollectionsOptions)
				listCollectionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionsOptionsModel.Name = core.StringPtr("testString")
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
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}]}`)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				listCollectionsOptionsModel := new(discoveryv1.ListCollectionsOptions)
				listCollectionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionsOptionsModel.Name = core.StringPtr("testString")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCollections with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv1.ListCollectionsOptions)
				listCollectionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionsOptionsModel.Name = core.StringPtr("testString")
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
				listCollectionsOptionsModelNew := new(discoveryv1.ListCollectionsOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(discoveryv1.ListCollectionsOptions)
				listCollectionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionsOptionsModel.Name = core.StringPtr("testString")
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
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions) - Operation response error`, func() {
		version := "testString"
		getCollectionPath := "/v1/environments/testString/collections/testString"
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv1.GetCollectionOptions)
				getCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
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
		getCollectionPath := "/v1/environments/testString/collections/testString"
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
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}`)
				}))
			})
			It(`Invoke GetCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv1.GetCollectionOptions)
				getCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}`)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				getCollectionOptionsModel := new(discoveryv1.GetCollectionOptions)
				getCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv1.GetCollectionOptions)
				getCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
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
				getCollectionOptionsModelNew := new(discoveryv1.GetCollectionOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(discoveryv1.GetCollectionOptions)
				getCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
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
		updateCollectionPath := "/v1/environments/testString/collections/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCollection with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
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
		updateCollectionPath := "/v1/environments/testString/collections/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}`)
				}))
			})
			It(`Invoke UpdateCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
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
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "name": "Name", "description": "Description", "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z", "status": "active", "configuration_id": "ConfigurationID", "language": "Language", "document_counts": {"available": 9, "processing": 10, "failed": 6, "pending": 7}, "disk_usage": {"used_bytes": 9}, "training_status": {"total_examples": 13, "available": false, "processing": true, "minimum_queries_added": false, "minimum_examples_added": true, "sufficient_label_diversity": true, "notices": 7, "successfully_trained": "2019-01-01T12:00:00.000Z", "data_updated": "2019-01-01T12:00:00.000Z"}, "crawl_status": {"source_crawl": {"status": "running", "next_crawl": "2019-01-01T12:00:00.000Z"}}, "smart_document_understanding": {"enabled": false, "total_annotated_pages": 19, "total_pages": 10, "total_documents": 14, "custom_fields": {"defined": 7, "maximum_allowed": 14}}}`)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
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
				updateCollectionOptionsModelNew := new(discoveryv1.UpdateCollectionOptions)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(discoveryv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.ConfigurationID = core.StringPtr("testString")
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
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) - Operation response error`, func() {
		version := "testString"
		deleteCollectionPath := "/v1/environments/testString/collections/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectionPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCollection with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		version := "testString"
		deleteCollectionPath := "/v1/environments/testString/collections/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteCollection successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.DeleteCollectionWithContext(ctx, deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.DeleteCollectionWithContext(ctx, deleteCollectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collection_id": "CollectionID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCollection with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCollectionOptions model with no property values
				deleteCollectionOptionsModelNew := new(discoveryv1.DeleteCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.DeleteCollection(deleteCollectionOptionsModelNew)
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
			It(`Invoke DeleteCollection successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(discoveryv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.DeleteCollection(deleteCollectionOptionsModel)
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
	Describe(`ListCollectionFields(listCollectionFieldsOptions *ListCollectionFieldsOptions) - Operation response error`, func() {
		version := "testString"
		listCollectionFieldsPath := "/v1/environments/testString/collections/testString/fields"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionFieldsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCollectionFields with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionFieldsOptions model
				listCollectionFieldsOptionsModel := new(discoveryv1.ListCollectionFieldsOptions)
				listCollectionFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.CollectionID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListCollectionFields(listCollectionFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListCollectionFields(listCollectionFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCollectionFields(listCollectionFieldsOptions *ListCollectionFieldsOptions)`, func() {
		version := "testString"
		listCollectionFieldsPath := "/v1/environments/testString/collections/testString/fields"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionFieldsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"fields": [{"field": "Field", "type": "nested"}]}`)
				}))
			})
			It(`Invoke ListCollectionFields successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListCollectionFieldsOptions model
				listCollectionFieldsOptionsModel := new(discoveryv1.ListCollectionFieldsOptions)
				listCollectionFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.CollectionID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListCollectionFieldsWithContext(ctx, listCollectionFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListCollectionFields(listCollectionFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListCollectionFieldsWithContext(ctx, listCollectionFieldsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionFieldsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"fields": [{"field": "Field", "type": "nested"}]}`)
				}))
			})
			It(`Invoke ListCollectionFields successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListCollectionFields(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCollectionFieldsOptions model
				listCollectionFieldsOptionsModel := new(discoveryv1.ListCollectionFieldsOptions)
				listCollectionFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.CollectionID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListCollectionFields(listCollectionFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCollectionFields with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionFieldsOptions model
				listCollectionFieldsOptionsModel := new(discoveryv1.ListCollectionFieldsOptions)
				listCollectionFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.CollectionID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListCollectionFields(listCollectionFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCollectionFieldsOptions model with no property values
				listCollectionFieldsOptionsModelNew := new(discoveryv1.ListCollectionFieldsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListCollectionFields(listCollectionFieldsOptionsModelNew)
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
			It(`Invoke ListCollectionFields successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCollectionFieldsOptions model
				listCollectionFieldsOptionsModel := new(discoveryv1.ListCollectionFieldsOptions)
				listCollectionFieldsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.CollectionID = core.StringPtr("testString")
				listCollectionFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListCollectionFields(listCollectionFieldsOptionsModel)
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
	Describe(`ListExpansions(listExpansionsOptions *ListExpansionsOptions) - Operation response error`, func() {
		version := "testString"
		listExpansionsPath := "/v1/environments/testString/collections/testString/expansions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listExpansionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListExpansions with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListExpansionsOptions model
				listExpansionsOptionsModel := new(discoveryv1.ListExpansionsOptions)
				listExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				listExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListExpansions(listExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListExpansions(listExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListExpansions(listExpansionsOptions *ListExpansionsOptions)`, func() {
		version := "testString"
		listExpansionsPath := "/v1/environments/testString/collections/testString/expansions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listExpansionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"expansions": [{"input_terms": ["InputTerms"], "expanded_terms": ["ExpandedTerms"]}]}`)
				}))
			})
			It(`Invoke ListExpansions successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListExpansionsOptions model
				listExpansionsOptionsModel := new(discoveryv1.ListExpansionsOptions)
				listExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				listExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListExpansionsWithContext(ctx, listExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListExpansions(listExpansionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListExpansionsWithContext(ctx, listExpansionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listExpansionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"expansions": [{"input_terms": ["InputTerms"], "expanded_terms": ["ExpandedTerms"]}]}`)
				}))
			})
			It(`Invoke ListExpansions successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListExpansions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListExpansionsOptions model
				listExpansionsOptionsModel := new(discoveryv1.ListExpansionsOptions)
				listExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				listExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListExpansions(listExpansionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListExpansions with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListExpansionsOptions model
				listExpansionsOptionsModel := new(discoveryv1.ListExpansionsOptions)
				listExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				listExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListExpansions(listExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListExpansionsOptions model with no property values
				listExpansionsOptionsModelNew := new(discoveryv1.ListExpansionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListExpansions(listExpansionsOptionsModelNew)
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
			It(`Invoke ListExpansions successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListExpansionsOptions model
				listExpansionsOptionsModel := new(discoveryv1.ListExpansionsOptions)
				listExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				listExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListExpansions(listExpansionsOptionsModel)
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
	Describe(`CreateExpansions(createExpansionsOptions *CreateExpansionsOptions) - Operation response error`, func() {
		version := "testString"
		createExpansionsPath := "/v1/environments/testString/collections/testString/expansions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createExpansionsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateExpansions with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the Expansion model
				expansionModel := new(discoveryv1.Expansion)
				expansionModel.InputTerms = []string{"testString"}
				expansionModel.ExpandedTerms = []string{"testString"}

				// Construct an instance of the CreateExpansionsOptions model
				createExpansionsOptionsModel := new(discoveryv1.CreateExpansionsOptions)
				createExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				createExpansionsOptionsModel.Expansions = []discoveryv1.Expansion{*expansionModel}
				createExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateExpansions(createExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateExpansions(createExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateExpansions(createExpansionsOptions *CreateExpansionsOptions)`, func() {
		version := "testString"
		createExpansionsPath := "/v1/environments/testString/collections/testString/expansions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createExpansionsPath))
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
					fmt.Fprintf(res, "%s", `{"expansions": [{"input_terms": ["InputTerms"], "expanded_terms": ["ExpandedTerms"]}]}`)
				}))
			})
			It(`Invoke CreateExpansions successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the Expansion model
				expansionModel := new(discoveryv1.Expansion)
				expansionModel.InputTerms = []string{"testString"}
				expansionModel.ExpandedTerms = []string{"testString"}

				// Construct an instance of the CreateExpansionsOptions model
				createExpansionsOptionsModel := new(discoveryv1.CreateExpansionsOptions)
				createExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				createExpansionsOptionsModel.Expansions = []discoveryv1.Expansion{*expansionModel}
				createExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateExpansionsWithContext(ctx, createExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateExpansions(createExpansionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateExpansionsWithContext(ctx, createExpansionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createExpansionsPath))
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
					fmt.Fprintf(res, "%s", `{"expansions": [{"input_terms": ["InputTerms"], "expanded_terms": ["ExpandedTerms"]}]}`)
				}))
			})
			It(`Invoke CreateExpansions successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateExpansions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Expansion model
				expansionModel := new(discoveryv1.Expansion)
				expansionModel.InputTerms = []string{"testString"}
				expansionModel.ExpandedTerms = []string{"testString"}

				// Construct an instance of the CreateExpansionsOptions model
				createExpansionsOptionsModel := new(discoveryv1.CreateExpansionsOptions)
				createExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				createExpansionsOptionsModel.Expansions = []discoveryv1.Expansion{*expansionModel}
				createExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateExpansions(createExpansionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateExpansions with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the Expansion model
				expansionModel := new(discoveryv1.Expansion)
				expansionModel.InputTerms = []string{"testString"}
				expansionModel.ExpandedTerms = []string{"testString"}

				// Construct an instance of the CreateExpansionsOptions model
				createExpansionsOptionsModel := new(discoveryv1.CreateExpansionsOptions)
				createExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				createExpansionsOptionsModel.Expansions = []discoveryv1.Expansion{*expansionModel}
				createExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateExpansions(createExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateExpansionsOptions model with no property values
				createExpansionsOptionsModelNew := new(discoveryv1.CreateExpansionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateExpansions(createExpansionsOptionsModelNew)
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
			It(`Invoke CreateExpansions successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the Expansion model
				expansionModel := new(discoveryv1.Expansion)
				expansionModel.InputTerms = []string{"testString"}
				expansionModel.ExpandedTerms = []string{"testString"}

				// Construct an instance of the CreateExpansionsOptions model
				createExpansionsOptionsModel := new(discoveryv1.CreateExpansionsOptions)
				createExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				createExpansionsOptionsModel.Expansions = []discoveryv1.Expansion{*expansionModel}
				createExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateExpansions(createExpansionsOptionsModel)
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
	Describe(`DeleteExpansions(deleteExpansionsOptions *DeleteExpansionsOptions)`, func() {
		version := "testString"
		deleteExpansionsPath := "/v1/environments/testString/collections/testString/expansions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteExpansionsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteExpansions successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteExpansions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteExpansionsOptions model
				deleteExpansionsOptionsModel := new(discoveryv1.DeleteExpansionsOptions)
				deleteExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				deleteExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteExpansions(deleteExpansionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteExpansions with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteExpansionsOptions model
				deleteExpansionsOptionsModel := new(discoveryv1.DeleteExpansionsOptions)
				deleteExpansionsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteExpansionsOptionsModel.CollectionID = core.StringPtr("testString")
				deleteExpansionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteExpansions(deleteExpansionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteExpansionsOptions model with no property values
				deleteExpansionsOptionsModelNew := new(discoveryv1.DeleteExpansionsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteExpansions(deleteExpansionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions) - Operation response error`, func() {
		version := "testString"
		getTokenizationDictionaryStatusPath := "/v1/environments/testString/collections/testString/word_lists/tokenization_dictionary"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTokenizationDictionaryStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTokenizationDictionaryStatus with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTokenizationDictionaryStatusOptions model
				getTokenizationDictionaryStatusOptionsModel := new(discoveryv1.GetTokenizationDictionaryStatusOptions)
				getTokenizationDictionaryStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions *GetTokenizationDictionaryStatusOptions)`, func() {
		version := "testString"
		getTokenizationDictionaryStatusPath := "/v1/environments/testString/collections/testString/word_lists/tokenization_dictionary"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTokenizationDictionaryStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke GetTokenizationDictionaryStatus successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetTokenizationDictionaryStatusOptions model
				getTokenizationDictionaryStatusOptionsModel := new(discoveryv1.GetTokenizationDictionaryStatusOptions)
				getTokenizationDictionaryStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetTokenizationDictionaryStatusWithContext(ctx, getTokenizationDictionaryStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetTokenizationDictionaryStatusWithContext(ctx, getTokenizationDictionaryStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTokenizationDictionaryStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke GetTokenizationDictionaryStatus successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetTokenizationDictionaryStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTokenizationDictionaryStatusOptions model
				getTokenizationDictionaryStatusOptionsModel := new(discoveryv1.GetTokenizationDictionaryStatusOptions)
				getTokenizationDictionaryStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTokenizationDictionaryStatus with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTokenizationDictionaryStatusOptions model
				getTokenizationDictionaryStatusOptionsModel := new(discoveryv1.GetTokenizationDictionaryStatusOptions)
				getTokenizationDictionaryStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTokenizationDictionaryStatusOptions model with no property values
				getTokenizationDictionaryStatusOptionsModelNew := new(discoveryv1.GetTokenizationDictionaryStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModelNew)
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
			It(`Invoke GetTokenizationDictionaryStatus successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTokenizationDictionaryStatusOptions model
				getTokenizationDictionaryStatusOptionsModel := new(discoveryv1.GetTokenizationDictionaryStatusOptions)
				getTokenizationDictionaryStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getTokenizationDictionaryStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptionsModel)
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
	Describe(`CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions) - Operation response error`, func() {
		version := "testString"
		createTokenizationDictionaryPath := "/v1/environments/testString/collections/testString/word_lists/tokenization_dictionary"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTokenizationDictionaryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTokenizationDictionary with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TokenDictRule model
				tokenDictRuleModel := new(discoveryv1.TokenDictRule)
				tokenDictRuleModel.Text = core.StringPtr("testString")
				tokenDictRuleModel.Tokens = []string{"testString"}
				tokenDictRuleModel.Readings = []string{"testString"}
				tokenDictRuleModel.PartOfSpeech = core.StringPtr("testString")

				// Construct an instance of the CreateTokenizationDictionaryOptions model
				createTokenizationDictionaryOptionsModel := new(discoveryv1.CreateTokenizationDictionaryOptions)
				createTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.TokenizationRules = []discoveryv1.TokenDictRule{*tokenDictRuleModel}
				createTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTokenizationDictionary(createTokenizationDictionaryOptions *CreateTokenizationDictionaryOptions)`, func() {
		version := "testString"
		createTokenizationDictionaryPath := "/v1/environments/testString/collections/testString/word_lists/tokenization_dictionary"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTokenizationDictionaryPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateTokenizationDictionary successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the TokenDictRule model
				tokenDictRuleModel := new(discoveryv1.TokenDictRule)
				tokenDictRuleModel.Text = core.StringPtr("testString")
				tokenDictRuleModel.Tokens = []string{"testString"}
				tokenDictRuleModel.Readings = []string{"testString"}
				tokenDictRuleModel.PartOfSpeech = core.StringPtr("testString")

				// Construct an instance of the CreateTokenizationDictionaryOptions model
				createTokenizationDictionaryOptionsModel := new(discoveryv1.CreateTokenizationDictionaryOptions)
				createTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.TokenizationRules = []discoveryv1.TokenDictRule{*tokenDictRuleModel}
				createTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateTokenizationDictionaryWithContext(ctx, createTokenizationDictionaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateTokenizationDictionaryWithContext(ctx, createTokenizationDictionaryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTokenizationDictionaryPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateTokenizationDictionary successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateTokenizationDictionary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TokenDictRule model
				tokenDictRuleModel := new(discoveryv1.TokenDictRule)
				tokenDictRuleModel.Text = core.StringPtr("testString")
				tokenDictRuleModel.Tokens = []string{"testString"}
				tokenDictRuleModel.Readings = []string{"testString"}
				tokenDictRuleModel.PartOfSpeech = core.StringPtr("testString")

				// Construct an instance of the CreateTokenizationDictionaryOptions model
				createTokenizationDictionaryOptionsModel := new(discoveryv1.CreateTokenizationDictionaryOptions)
				createTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.TokenizationRules = []discoveryv1.TokenDictRule{*tokenDictRuleModel}
				createTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTokenizationDictionary with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TokenDictRule model
				tokenDictRuleModel := new(discoveryv1.TokenDictRule)
				tokenDictRuleModel.Text = core.StringPtr("testString")
				tokenDictRuleModel.Tokens = []string{"testString"}
				tokenDictRuleModel.Readings = []string{"testString"}
				tokenDictRuleModel.PartOfSpeech = core.StringPtr("testString")

				// Construct an instance of the CreateTokenizationDictionaryOptions model
				createTokenizationDictionaryOptionsModel := new(discoveryv1.CreateTokenizationDictionaryOptions)
				createTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.TokenizationRules = []discoveryv1.TokenDictRule{*tokenDictRuleModel}
				createTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTokenizationDictionaryOptions model with no property values
				createTokenizationDictionaryOptionsModelNew := new(discoveryv1.CreateTokenizationDictionaryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModelNew)
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
			It(`Invoke CreateTokenizationDictionary successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TokenDictRule model
				tokenDictRuleModel := new(discoveryv1.TokenDictRule)
				tokenDictRuleModel.Text = core.StringPtr("testString")
				tokenDictRuleModel.Tokens = []string{"testString"}
				tokenDictRuleModel.Readings = []string{"testString"}
				tokenDictRuleModel.PartOfSpeech = core.StringPtr("testString")

				// Construct an instance of the CreateTokenizationDictionaryOptions model
				createTokenizationDictionaryOptionsModel := new(discoveryv1.CreateTokenizationDictionaryOptions)
				createTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				createTokenizationDictionaryOptionsModel.TokenizationRules = []discoveryv1.TokenDictRule{*tokenDictRuleModel}
				createTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptionsModel)
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
	Describe(`DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions *DeleteTokenizationDictionaryOptions)`, func() {
		version := "testString"
		deleteTokenizationDictionaryPath := "/v1/environments/testString/collections/testString/word_lists/tokenization_dictionary"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTokenizationDictionaryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteTokenizationDictionary successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteTokenizationDictionary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTokenizationDictionaryOptions model
				deleteTokenizationDictionaryOptionsModel := new(discoveryv1.DeleteTokenizationDictionaryOptions)
				deleteTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				deleteTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteTokenizationDictionary(deleteTokenizationDictionaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTokenizationDictionary with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteTokenizationDictionaryOptions model
				deleteTokenizationDictionaryOptionsModel := new(discoveryv1.DeleteTokenizationDictionaryOptions)
				deleteTokenizationDictionaryOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteTokenizationDictionaryOptionsModel.CollectionID = core.StringPtr("testString")
				deleteTokenizationDictionaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteTokenizationDictionary(deleteTokenizationDictionaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTokenizationDictionaryOptions model with no property values
				deleteTokenizationDictionaryOptionsModelNew := new(discoveryv1.DeleteTokenizationDictionaryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteTokenizationDictionary(deleteTokenizationDictionaryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetStopwordListStatus(getStopwordListStatusOptions *GetStopwordListStatusOptions) - Operation response error`, func() {
		version := "testString"
		getStopwordListStatusPath := "/v1/environments/testString/collections/testString/word_lists/stopwords"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getStopwordListStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetStopwordListStatus with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetStopwordListStatusOptions model
				getStopwordListStatusOptionsModel := new(discoveryv1.GetStopwordListStatusOptions)
				getStopwordListStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetStopwordListStatus(getStopwordListStatusOptions *GetStopwordListStatusOptions)`, func() {
		version := "testString"
		getStopwordListStatusPath := "/v1/environments/testString/collections/testString/word_lists/stopwords"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getStopwordListStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke GetStopwordListStatus successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetStopwordListStatusOptions model
				getStopwordListStatusOptionsModel := new(discoveryv1.GetStopwordListStatusOptions)
				getStopwordListStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetStopwordListStatusWithContext(ctx, getStopwordListStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetStopwordListStatusWithContext(ctx, getStopwordListStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getStopwordListStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke GetStopwordListStatus successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetStopwordListStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetStopwordListStatusOptions model
				getStopwordListStatusOptionsModel := new(discoveryv1.GetStopwordListStatusOptions)
				getStopwordListStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetStopwordListStatus with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetStopwordListStatusOptions model
				getStopwordListStatusOptionsModel := new(discoveryv1.GetStopwordListStatusOptions)
				getStopwordListStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetStopwordListStatusOptions model with no property values
				getStopwordListStatusOptionsModelNew := new(discoveryv1.GetStopwordListStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModelNew)
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
			It(`Invoke GetStopwordListStatus successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetStopwordListStatusOptions model
				getStopwordListStatusOptionsModel := new(discoveryv1.GetStopwordListStatusOptions)
				getStopwordListStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getStopwordListStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetStopwordListStatus(getStopwordListStatusOptionsModel)
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
	Describe(`CreateStopwordList(createStopwordListOptions *CreateStopwordListOptions) - Operation response error`, func() {
		version := "testString"
		createStopwordListPath := "/v1/environments/testString/collections/testString/word_lists/stopwords"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createStopwordListPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateStopwordList with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateStopwordListOptions model
				createStopwordListOptionsModel := new(discoveryv1.CreateStopwordListOptions)
				createStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				createStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				createStopwordListOptionsModel.StopwordFile = CreateMockReader("This is a mock file.")
				createStopwordListOptionsModel.StopwordFilename = core.StringPtr("testString")
				createStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateStopwordList(createStopwordListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateStopwordList(createStopwordListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateStopwordList(createStopwordListOptions *CreateStopwordListOptions)`, func() {
		version := "testString"
		createStopwordListPath := "/v1/environments/testString/collections/testString/word_lists/stopwords"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createStopwordListPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateStopwordList successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CreateStopwordListOptions model
				createStopwordListOptionsModel := new(discoveryv1.CreateStopwordListOptions)
				createStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				createStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				createStopwordListOptionsModel.StopwordFile = CreateMockReader("This is a mock file.")
				createStopwordListOptionsModel.StopwordFilename = core.StringPtr("testString")
				createStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateStopwordListWithContext(ctx, createStopwordListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateStopwordList(createStopwordListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateStopwordListWithContext(ctx, createStopwordListOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createStopwordListPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "active", "type": "Type"}`)
				}))
			})
			It(`Invoke CreateStopwordList successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateStopwordList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateStopwordListOptions model
				createStopwordListOptionsModel := new(discoveryv1.CreateStopwordListOptions)
				createStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				createStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				createStopwordListOptionsModel.StopwordFile = CreateMockReader("This is a mock file.")
				createStopwordListOptionsModel.StopwordFilename = core.StringPtr("testString")
				createStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateStopwordList(createStopwordListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateStopwordList with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateStopwordListOptions model
				createStopwordListOptionsModel := new(discoveryv1.CreateStopwordListOptions)
				createStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				createStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				createStopwordListOptionsModel.StopwordFile = CreateMockReader("This is a mock file.")
				createStopwordListOptionsModel.StopwordFilename = core.StringPtr("testString")
				createStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateStopwordList(createStopwordListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateStopwordListOptions model with no property values
				createStopwordListOptionsModelNew := new(discoveryv1.CreateStopwordListOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateStopwordList(createStopwordListOptionsModelNew)
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
			It(`Invoke CreateStopwordList successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateStopwordListOptions model
				createStopwordListOptionsModel := new(discoveryv1.CreateStopwordListOptions)
				createStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				createStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				createStopwordListOptionsModel.StopwordFile = CreateMockReader("This is a mock file.")
				createStopwordListOptionsModel.StopwordFilename = core.StringPtr("testString")
				createStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateStopwordList(createStopwordListOptionsModel)
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
	Describe(`DeleteStopwordList(deleteStopwordListOptions *DeleteStopwordListOptions)`, func() {
		version := "testString"
		deleteStopwordListPath := "/v1/environments/testString/collections/testString/word_lists/stopwords"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteStopwordListPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteStopwordList successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteStopwordList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteStopwordListOptions model
				deleteStopwordListOptionsModel := new(discoveryv1.DeleteStopwordListOptions)
				deleteStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				deleteStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteStopwordList(deleteStopwordListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteStopwordList with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteStopwordListOptions model
				deleteStopwordListOptionsModel := new(discoveryv1.DeleteStopwordListOptions)
				deleteStopwordListOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteStopwordListOptionsModel.CollectionID = core.StringPtr("testString")
				deleteStopwordListOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteStopwordList(deleteStopwordListOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteStopwordListOptions model with no property values
				deleteStopwordListOptionsModelNew := new(discoveryv1.DeleteStopwordListOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteStopwordList(deleteStopwordListOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddDocument(addDocumentOptions *AddDocumentOptions) - Operation response error`, func() {
		version := "testString"
		addDocumentPath := "/v1/environments/testString/collections/testString/documents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv1.AddDocumentOptions)
				addDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
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
		addDocumentPath := "/v1/environments/testString/collections/testString/documents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke AddDocument successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv1.AddDocumentOptions)
				addDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke AddDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				addDocumentOptionsModel := new(discoveryv1.AddDocumentOptions)
				addDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
				addDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddDocument with error: Param validation error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv1.AddDocumentOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := discoveryService.AddDocument(addDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke AddDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv1.AddDocumentOptions)
				addDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
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
				addDocumentOptionsModelNew := new(discoveryv1.AddDocumentOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the AddDocumentOptions model
				addDocumentOptionsModel := new(discoveryv1.AddDocumentOptions)
				addDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				addDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				addDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				addDocumentOptionsModel.Filename = core.StringPtr("testString")
				addDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				addDocumentOptionsModel.Metadata = core.StringPtr("testString")
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
	Describe(`GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions) - Operation response error`, func() {
		version := "testString"
		getDocumentStatusPath := "/v1/environments/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocumentStatus with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(discoveryv1.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocumentStatus(getDocumentStatusOptions *GetDocumentStatusOptions)`, func() {
		version := "testString"
		getDocumentStatusPath := "/v1/environments/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "configuration_id": "ConfigurationID", "status": "available", "status_description": "StatusDescription", "filename": "Filename", "file_type": "pdf", "sha1": "Sha1", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetDocumentStatus successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(discoveryv1.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetDocumentStatusWithContext(ctx, getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetDocumentStatusWithContext(ctx, getDocumentStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "configuration_id": "ConfigurationID", "status": "available", "status_description": "StatusDescription", "filename": "Filename", "file_type": "pdf", "sha1": "Sha1", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetDocumentStatus successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetDocumentStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(discoveryv1.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDocumentStatus with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(discoveryv1.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetDocumentStatus(getDocumentStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentStatusOptions model with no property values
				getDocumentStatusOptionsModelNew := new(discoveryv1.GetDocumentStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetDocumentStatus(getDocumentStatusOptionsModelNew)
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
			It(`Invoke GetDocumentStatus successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetDocumentStatusOptions model
				getDocumentStatusOptionsModel := new(discoveryv1.GetDocumentStatusOptions)
				getDocumentStatusOptionsModel.EnvironmentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.CollectionID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetDocumentStatus(getDocumentStatusOptionsModel)
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
		updateDocumentPath := "/v1/environments/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv1.UpdateDocumentOptions)
				updateDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
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
		updateDocumentPath := "/v1/environments/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke UpdateDocument successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv1.UpdateDocumentOptions)
				updateDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "processing", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}`)
				}))
			})
			It(`Invoke UpdateDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				updateDocumentOptionsModel := new(discoveryv1.UpdateDocumentOptions)
				updateDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
				updateDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDocument with error: Param validation error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv1.UpdateDocumentOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := discoveryService.UpdateDocument(updateDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke UpdateDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv1.UpdateDocumentOptions)
				updateDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
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
				updateDocumentOptionsModelNew := new(discoveryv1.UpdateDocumentOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateDocumentOptions model
				updateDocumentOptionsModel := new(discoveryv1.UpdateDocumentOptions)
				updateDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				updateDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDocumentOptionsModel.File = CreateMockReader("This is a mock file.")
				updateDocumentOptionsModel.Filename = core.StringPtr("testString")
				updateDocumentOptionsModel.FileContentType = core.StringPtr("application/json")
				updateDocumentOptionsModel.Metadata = core.StringPtr("testString")
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
		deleteDocumentPath := "/v1/environments/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDocument with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
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
		deleteDocumentPath := "/v1/environments/testString/collections/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteDocument successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				deleteDocumentOptionsModel := new(discoveryv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDocument with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
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
				deleteDocumentOptionsModelNew := new(discoveryv1.DeleteDocumentOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(discoveryv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteDocumentOptionsModel.CollectionID = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocumentID = core.StringPtr("testString")
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
	Describe(`Query(queryOptions *QueryOptions) - Operation response error`, func() {
		version := "testString"
		queryPath := "/v1/environments/testString/collections/testString/query"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Watson-Logging-Opt-Out"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Logging-Opt-Out"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Query with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv1.QueryOptions)
				queryOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryOptionsModel.CollectionID = core.StringPtr("testString")
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Passages = core.BoolPtr(true)
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = core.StringPtr("testString")
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.PassagesFields = core.StringPtr("testString")
				queryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryOptionsModel.Deduplicate = core.BoolPtr(true)
				queryOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryOptionsModel.Similar = core.BoolPtr(true)
				queryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				queryOptionsModel.SimilarFields = core.StringPtr("testString")
				queryOptionsModel.Bias = core.StringPtr("testString")
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
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
		queryPath := "/v1/environments/testString/collections/testString/query"
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

					Expect(req.Header["X-Watson-Logging-Opt-Out"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Logging-Opt-Out"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17, "session_token": "SessionToken", "retrieval_details": {"document_retrieval_strategy": "untrained"}, "suggested_query": "SuggestedQuery"}`)
				}))
			})
			It(`Invoke Query successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv1.QueryOptions)
				queryOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryOptionsModel.CollectionID = core.StringPtr("testString")
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Passages = core.BoolPtr(true)
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = core.StringPtr("testString")
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.PassagesFields = core.StringPtr("testString")
				queryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryOptionsModel.Deduplicate = core.BoolPtr(true)
				queryOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryOptionsModel.Similar = core.BoolPtr(true)
				queryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				queryOptionsModel.SimilarFields = core.StringPtr("testString")
				queryOptionsModel.Bias = core.StringPtr("testString")
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
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

					Expect(req.Header["X-Watson-Logging-Opt-Out"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Logging-Opt-Out"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17, "session_token": "SessionToken", "retrieval_details": {"document_retrieval_strategy": "untrained"}, "suggested_query": "SuggestedQuery"}`)
				}))
			})
			It(`Invoke Query successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv1.QueryOptions)
				queryOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryOptionsModel.CollectionID = core.StringPtr("testString")
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Passages = core.BoolPtr(true)
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = core.StringPtr("testString")
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.PassagesFields = core.StringPtr("testString")
				queryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryOptionsModel.Deduplicate = core.BoolPtr(true)
				queryOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryOptionsModel.Similar = core.BoolPtr(true)
				queryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				queryOptionsModel.SimilarFields = core.StringPtr("testString")
				queryOptionsModel.Bias = core.StringPtr("testString")
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
				queryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.Query(queryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Query with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv1.QueryOptions)
				queryOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryOptionsModel.CollectionID = core.StringPtr("testString")
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Passages = core.BoolPtr(true)
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = core.StringPtr("testString")
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.PassagesFields = core.StringPtr("testString")
				queryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryOptionsModel.Deduplicate = core.BoolPtr(true)
				queryOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryOptionsModel.Similar = core.BoolPtr(true)
				queryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				queryOptionsModel.SimilarFields = core.StringPtr("testString")
				queryOptionsModel.Bias = core.StringPtr("testString")
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
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
				queryOptionsModelNew := new(discoveryv1.QueryOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryOptions model
				queryOptionsModel := new(discoveryv1.QueryOptions)
				queryOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryOptionsModel.CollectionID = core.StringPtr("testString")
				queryOptionsModel.Filter = core.StringPtr("testString")
				queryOptionsModel.Query = core.StringPtr("testString")
				queryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryOptionsModel.Passages = core.BoolPtr(true)
				queryOptionsModel.Aggregation = core.StringPtr("testString")
				queryOptionsModel.Count = core.Int64Ptr(int64(38))
				queryOptionsModel.Return = core.StringPtr("testString")
				queryOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryOptionsModel.Sort = core.StringPtr("testString")
				queryOptionsModel.Highlight = core.BoolPtr(true)
				queryOptionsModel.PassagesFields = core.StringPtr("testString")
				queryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryOptionsModel.Deduplicate = core.BoolPtr(true)
				queryOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryOptionsModel.Similar = core.BoolPtr(true)
				queryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				queryOptionsModel.SimilarFields = core.StringPtr("testString")
				queryOptionsModel.Bias = core.StringPtr("testString")
				queryOptionsModel.SpellingSuggestions = core.BoolPtr(true)
				queryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
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
	Describe(`QueryNotices(queryNoticesOptions *QueryNoticesOptions) - Operation response error`, func() {
		version := "testString"
		queryNoticesPath := "/v1/environments/testString/collections/testString/notices"
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
					// TODO: Add check for passages query parameter
					Expect(req.URL.Query()["aggregation"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for highlight query parameter
					Expect(req.URL.Query()["passages.count"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["passages.characters"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["deduplicate.field"]).To(Equal([]string{"testString"}))
					// TODO: Add check for similar query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke QueryNotices with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv1.QueryNoticesOptions)
				queryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Passages = core.BoolPtr(true)
				queryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Return = []string{"testString"}
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Sort = []string{"testString"}
				queryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				queryNoticesOptionsModel.PassagesFields = []string{"testString"}
				queryNoticesOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryNoticesOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryNoticesOptionsModel.Similar = core.BoolPtr(true)
				queryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				queryNoticesOptionsModel.SimilarFields = []string{"testString"}
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
		queryNoticesPath := "/v1/environments/testString/collections/testString/notices"
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
					// TODO: Add check for passages query parameter
					Expect(req.URL.Query()["aggregation"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for highlight query parameter
					Expect(req.URL.Query()["passages.count"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["passages.characters"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["deduplicate.field"]).To(Equal([]string{"testString"}))
					// TODO: Add check for similar query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}, "code": 4, "filename": "Filename", "file_type": "pdf", "sha1": "Sha1", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17}`)
				}))
			})
			It(`Invoke QueryNotices successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv1.QueryNoticesOptions)
				queryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Passages = core.BoolPtr(true)
				queryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Return = []string{"testString"}
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Sort = []string{"testString"}
				queryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				queryNoticesOptionsModel.PassagesFields = []string{"testString"}
				queryNoticesOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryNoticesOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryNoticesOptionsModel.Similar = core.BoolPtr(true)
				queryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				queryNoticesOptionsModel.SimilarFields = []string{"testString"}
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
					// TODO: Add check for passages query parameter
					Expect(req.URL.Query()["aggregation"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for highlight query parameter
					Expect(req.URL.Query()["passages.count"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["passages.characters"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["deduplicate.field"]).To(Equal([]string{"testString"}))
					// TODO: Add check for similar query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}, "code": 4, "filename": "Filename", "file_type": "pdf", "sha1": "Sha1", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17}`)
				}))
			})
			It(`Invoke QueryNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				queryNoticesOptionsModel := new(discoveryv1.QueryNoticesOptions)
				queryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Passages = core.BoolPtr(true)
				queryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Return = []string{"testString"}
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Sort = []string{"testString"}
				queryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				queryNoticesOptionsModel.PassagesFields = []string{"testString"}
				queryNoticesOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryNoticesOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryNoticesOptionsModel.Similar = core.BoolPtr(true)
				queryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				queryNoticesOptionsModel.SimilarFields = []string{"testString"}
				queryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.QueryNotices(queryNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke QueryNotices with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv1.QueryNoticesOptions)
				queryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Passages = core.BoolPtr(true)
				queryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Return = []string{"testString"}
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Sort = []string{"testString"}
				queryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				queryNoticesOptionsModel.PassagesFields = []string{"testString"}
				queryNoticesOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryNoticesOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryNoticesOptionsModel.Similar = core.BoolPtr(true)
				queryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				queryNoticesOptionsModel.SimilarFields = []string{"testString"}
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
				queryNoticesOptionsModelNew := new(discoveryv1.QueryNoticesOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryNoticesOptions model
				queryNoticesOptionsModel := new(discoveryv1.QueryNoticesOptions)
				queryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				queryNoticesOptionsModel.CollectionID = core.StringPtr("testString")
				queryNoticesOptionsModel.Filter = core.StringPtr("testString")
				queryNoticesOptionsModel.Query = core.StringPtr("testString")
				queryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				queryNoticesOptionsModel.Passages = core.BoolPtr(true)
				queryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				queryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Return = []string{"testString"}
				queryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryNoticesOptionsModel.Sort = []string{"testString"}
				queryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				queryNoticesOptionsModel.PassagesFields = []string{"testString"}
				queryNoticesOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				queryNoticesOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				queryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				queryNoticesOptionsModel.Similar = core.BoolPtr(true)
				queryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				queryNoticesOptionsModel.SimilarFields = []string{"testString"}
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
	Describe(`FederatedQuery(federatedQueryOptions *FederatedQueryOptions) - Operation response error`, func() {
		version := "testString"
		federatedQueryPath := "/v1/environments/testString/query"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(federatedQueryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Watson-Logging-Opt-Out"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Logging-Opt-Out"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke FederatedQuery with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FederatedQueryOptions model
				federatedQueryOptionsModel := new(discoveryv1.FederatedQueryOptions)
				federatedQueryOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryOptionsModel.CollectionIds = core.StringPtr("testString")
				federatedQueryOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryOptionsModel.Query = core.StringPtr("testString")
				federatedQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryOptionsModel.Passages = core.BoolPtr(true)
				federatedQueryOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Return = core.StringPtr("testString")
				federatedQueryOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Sort = core.StringPtr("testString")
				federatedQueryOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryOptionsModel.PassagesFields = core.StringPtr("testString")
				federatedQueryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				federatedQueryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				federatedQueryOptionsModel.Deduplicate = core.BoolPtr(true)
				federatedQueryOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				federatedQueryOptionsModel.SimilarFields = core.StringPtr("testString")
				federatedQueryOptionsModel.Bias = core.StringPtr("testString")
				federatedQueryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
				federatedQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.FederatedQuery(federatedQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.FederatedQuery(federatedQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`FederatedQuery(federatedQueryOptions *FederatedQueryOptions)`, func() {
		version := "testString"
		federatedQueryPath := "/v1/environments/testString/query"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(federatedQueryPath))
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

					Expect(req.Header["X-Watson-Logging-Opt-Out"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Logging-Opt-Out"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17, "session_token": "SessionToken", "retrieval_details": {"document_retrieval_strategy": "untrained"}, "suggested_query": "SuggestedQuery"}`)
				}))
			})
			It(`Invoke FederatedQuery successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the FederatedQueryOptions model
				federatedQueryOptionsModel := new(discoveryv1.FederatedQueryOptions)
				federatedQueryOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryOptionsModel.CollectionIds = core.StringPtr("testString")
				federatedQueryOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryOptionsModel.Query = core.StringPtr("testString")
				federatedQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryOptionsModel.Passages = core.BoolPtr(true)
				federatedQueryOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Return = core.StringPtr("testString")
				federatedQueryOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Sort = core.StringPtr("testString")
				federatedQueryOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryOptionsModel.PassagesFields = core.StringPtr("testString")
				federatedQueryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				federatedQueryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				federatedQueryOptionsModel.Deduplicate = core.BoolPtr(true)
				federatedQueryOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				federatedQueryOptionsModel.SimilarFields = core.StringPtr("testString")
				federatedQueryOptionsModel.Bias = core.StringPtr("testString")
				federatedQueryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
				federatedQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.FederatedQueryWithContext(ctx, federatedQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.FederatedQuery(federatedQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.FederatedQueryWithContext(ctx, federatedQueryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(federatedQueryPath))
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

					Expect(req.Header["X-Watson-Logging-Opt-Out"]).ToNot(BeNil())
					Expect(req.Header["X-Watson-Logging-Opt-Out"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17, "session_token": "SessionToken", "retrieval_details": {"document_retrieval_strategy": "untrained"}, "suggested_query": "SuggestedQuery"}`)
				}))
			})
			It(`Invoke FederatedQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.FederatedQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FederatedQueryOptions model
				federatedQueryOptionsModel := new(discoveryv1.FederatedQueryOptions)
				federatedQueryOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryOptionsModel.CollectionIds = core.StringPtr("testString")
				federatedQueryOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryOptionsModel.Query = core.StringPtr("testString")
				federatedQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryOptionsModel.Passages = core.BoolPtr(true)
				federatedQueryOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Return = core.StringPtr("testString")
				federatedQueryOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Sort = core.StringPtr("testString")
				federatedQueryOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryOptionsModel.PassagesFields = core.StringPtr("testString")
				federatedQueryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				federatedQueryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				federatedQueryOptionsModel.Deduplicate = core.BoolPtr(true)
				federatedQueryOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				federatedQueryOptionsModel.SimilarFields = core.StringPtr("testString")
				federatedQueryOptionsModel.Bias = core.StringPtr("testString")
				federatedQueryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
				federatedQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.FederatedQuery(federatedQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke FederatedQuery with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FederatedQueryOptions model
				federatedQueryOptionsModel := new(discoveryv1.FederatedQueryOptions)
				federatedQueryOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryOptionsModel.CollectionIds = core.StringPtr("testString")
				federatedQueryOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryOptionsModel.Query = core.StringPtr("testString")
				federatedQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryOptionsModel.Passages = core.BoolPtr(true)
				federatedQueryOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Return = core.StringPtr("testString")
				federatedQueryOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Sort = core.StringPtr("testString")
				federatedQueryOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryOptionsModel.PassagesFields = core.StringPtr("testString")
				federatedQueryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				federatedQueryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				federatedQueryOptionsModel.Deduplicate = core.BoolPtr(true)
				federatedQueryOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				federatedQueryOptionsModel.SimilarFields = core.StringPtr("testString")
				federatedQueryOptionsModel.Bias = core.StringPtr("testString")
				federatedQueryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
				federatedQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.FederatedQuery(federatedQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the FederatedQueryOptions model with no property values
				federatedQueryOptionsModelNew := new(discoveryv1.FederatedQueryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.FederatedQuery(federatedQueryOptionsModelNew)
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
			It(`Invoke FederatedQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FederatedQueryOptions model
				federatedQueryOptionsModel := new(discoveryv1.FederatedQueryOptions)
				federatedQueryOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryOptionsModel.CollectionIds = core.StringPtr("testString")
				federatedQueryOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryOptionsModel.Query = core.StringPtr("testString")
				federatedQueryOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryOptionsModel.Passages = core.BoolPtr(true)
				federatedQueryOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Return = core.StringPtr("testString")
				federatedQueryOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryOptionsModel.Sort = core.StringPtr("testString")
				federatedQueryOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryOptionsModel.PassagesFields = core.StringPtr("testString")
				federatedQueryOptionsModel.PassagesCount = core.Int64Ptr(int64(100))
				federatedQueryOptionsModel.PassagesCharacters = core.Int64Ptr(int64(50))
				federatedQueryOptionsModel.Deduplicate = core.BoolPtr(true)
				federatedQueryOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryOptionsModel.SimilarDocumentIds = core.StringPtr("testString")
				federatedQueryOptionsModel.SimilarFields = core.StringPtr("testString")
				federatedQueryOptionsModel.Bias = core.StringPtr("testString")
				federatedQueryOptionsModel.XWatsonLoggingOptOut = core.BoolPtr(true)
				federatedQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.FederatedQuery(federatedQueryOptionsModel)
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
	Describe(`FederatedQueryNotices(federatedQueryNoticesOptions *FederatedQueryNoticesOptions) - Operation response error`, func() {
		version := "testString"
		federatedQueryNoticesPath := "/v1/environments/testString/notices"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(federatedQueryNoticesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["aggregation"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for highlight query parameter
					Expect(req.URL.Query()["deduplicate.field"]).To(Equal([]string{"testString"}))
					// TODO: Add check for similar query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke FederatedQueryNotices with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FederatedQueryNoticesOptions model
				federatedQueryNoticesOptionsModel := new(discoveryv1.FederatedQueryNoticesOptions)
				federatedQueryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.CollectionIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Query = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Return = []string{"testString"}
				federatedQueryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Sort = []string{"testString"}
				federatedQueryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.SimilarFields = []string{"testString"}
				federatedQueryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`FederatedQueryNotices(federatedQueryNoticesOptions *FederatedQueryNoticesOptions)`, func() {
		version := "testString"
		federatedQueryNoticesPath := "/v1/environments/testString/notices"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(federatedQueryNoticesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["aggregation"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for highlight query parameter
					Expect(req.URL.Query()["deduplicate.field"]).To(Equal([]string{"testString"}))
					// TODO: Add check for similar query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}, "code": 4, "filename": "Filename", "file_type": "pdf", "sha1": "Sha1", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17}`)
				}))
			})
			It(`Invoke FederatedQueryNotices successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the FederatedQueryNoticesOptions model
				federatedQueryNoticesOptionsModel := new(discoveryv1.FederatedQueryNoticesOptions)
				federatedQueryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.CollectionIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Query = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Return = []string{"testString"}
				federatedQueryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Sort = []string{"testString"}
				federatedQueryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.SimilarFields = []string{"testString"}
				federatedQueryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.FederatedQueryNoticesWithContext(ctx, federatedQueryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.FederatedQueryNoticesWithContext(ctx, federatedQueryNoticesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(federatedQueryNoticesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["natural_language_query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["aggregation"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// TODO: Add check for highlight query parameter
					Expect(req.URL.Query()["deduplicate.field"]).To(Equal([]string{"testString"}))
					// TODO: Add check for similar query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"id": "ID", "metadata": {"mapKey": "anyValue"}, "collection_id": "CollectionID", "result_metadata": {"score": 5, "confidence": 10}, "code": 4, "filename": "Filename", "file_type": "pdf", "sha1": "Sha1", "notices": [{"notice_id": "NoticeID", "created": "2019-01-01T12:00:00.000Z", "document_id": "DocumentID", "query_id": "QueryID", "severity": "warning", "step": "Step", "description": "Description"}]}], "aggregations": [{"type": "histogram", "matching_results": 15, "field": "Field", "interval": 8}], "passages": [{"document_id": "DocumentID", "passage_score": 12, "passage_text": "PassageText", "start_offset": 11, "end_offset": 9, "field": "Field"}], "duplicates_removed": 17}`)
				}))
			})
			It(`Invoke FederatedQueryNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.FederatedQueryNotices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FederatedQueryNoticesOptions model
				federatedQueryNoticesOptionsModel := new(discoveryv1.FederatedQueryNoticesOptions)
				federatedQueryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.CollectionIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Query = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Return = []string{"testString"}
				federatedQueryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Sort = []string{"testString"}
				federatedQueryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.SimilarFields = []string{"testString"}
				federatedQueryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke FederatedQueryNotices with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FederatedQueryNoticesOptions model
				federatedQueryNoticesOptionsModel := new(discoveryv1.FederatedQueryNoticesOptions)
				federatedQueryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.CollectionIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Query = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Return = []string{"testString"}
				federatedQueryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Sort = []string{"testString"}
				federatedQueryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.SimilarFields = []string{"testString"}
				federatedQueryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the FederatedQueryNoticesOptions model with no property values
				federatedQueryNoticesOptionsModelNew := new(discoveryv1.FederatedQueryNoticesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModelNew)
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
			It(`Invoke FederatedQueryNotices successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the FederatedQueryNoticesOptions model
				federatedQueryNoticesOptionsModel := new(discoveryv1.FederatedQueryNoticesOptions)
				federatedQueryNoticesOptionsModel.EnvironmentID = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.CollectionIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.Filter = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Query = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Aggregation = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Count = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Return = []string{"testString"}
				federatedQueryNoticesOptionsModel.Offset = core.Int64Ptr(int64(38))
				federatedQueryNoticesOptionsModel.Sort = []string{"testString"}
				federatedQueryNoticesOptionsModel.Highlight = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.DeduplicateField = core.StringPtr("testString")
				federatedQueryNoticesOptionsModel.Similar = core.BoolPtr(true)
				federatedQueryNoticesOptionsModel.SimilarDocumentIds = []string{"testString"}
				federatedQueryNoticesOptionsModel.SimilarFields = []string{"testString"}
				federatedQueryNoticesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.FederatedQueryNotices(federatedQueryNoticesOptionsModel)
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
		getAutocompletionPath := "/v1/environments/testString/collections/testString/autocompletion"
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv1.GetAutocompletionOptions)
				getAutocompletionOptionsModel.EnvironmentID = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
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
		getAutocompletionPath := "/v1/environments/testString/collections/testString/autocompletion"
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv1.GetAutocompletionOptions)
				getAutocompletionOptionsModel.EnvironmentID = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				getAutocompletionOptionsModel := new(discoveryv1.GetAutocompletionOptions)
				getAutocompletionOptionsModel.EnvironmentID = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv1.GetAutocompletionOptions)
				getAutocompletionOptionsModel.EnvironmentID = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
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
				getAutocompletionOptionsModelNew := new(discoveryv1.GetAutocompletionOptions)
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetAutocompletionOptions model
				getAutocompletionOptionsModel := new(discoveryv1.GetAutocompletionOptions)
				getAutocompletionOptionsModel.EnvironmentID = core.StringPtr("testString")
				getAutocompletionOptionsModel.CollectionID = core.StringPtr("testString")
				getAutocompletionOptionsModel.Prefix = core.StringPtr("testString")
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
	Describe(`ListTrainingData(listTrainingDataOptions *ListTrainingDataOptions) - Operation response error`, func() {
		version := "testString"
		listTrainingDataPath := "/v1/environments/testString/collections/testString/training_data"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingDataPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTrainingData with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingDataOptions model
				listTrainingDataOptionsModel := new(discoveryv1.ListTrainingDataOptions)
				listTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListTrainingData(listTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListTrainingData(listTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTrainingData(listTrainingDataOptions *ListTrainingDataOptions)`, func() {
		version := "testString"
		listTrainingDataPath := "/v1/environments/testString/collections/testString/training_data"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingDataPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "collection_id": "CollectionID", "queries": [{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}]}`)
				}))
			})
			It(`Invoke ListTrainingData successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListTrainingDataOptions model
				listTrainingDataOptionsModel := new(discoveryv1.ListTrainingDataOptions)
				listTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListTrainingDataWithContext(ctx, listTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListTrainingData(listTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListTrainingDataWithContext(ctx, listTrainingDataOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingDataPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_id": "EnvironmentID", "collection_id": "CollectionID", "queries": [{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}]}`)
				}))
			})
			It(`Invoke ListTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTrainingDataOptions model
				listTrainingDataOptionsModel := new(discoveryv1.ListTrainingDataOptions)
				listTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListTrainingData(listTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTrainingData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingDataOptions model
				listTrainingDataOptionsModel := new(discoveryv1.ListTrainingDataOptions)
				listTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListTrainingData(listTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTrainingDataOptions model with no property values
				listTrainingDataOptionsModelNew := new(discoveryv1.ListTrainingDataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListTrainingData(listTrainingDataOptionsModelNew)
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
			It(`Invoke ListTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingDataOptions model
				listTrainingDataOptionsModel := new(discoveryv1.ListTrainingDataOptions)
				listTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListTrainingData(listTrainingDataOptionsModel)
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
	Describe(`AddTrainingData(addTrainingDataOptions *AddTrainingDataOptions) - Operation response error`, func() {
		version := "testString"
		addTrainingDataPath := "/v1/environments/testString/collections/testString/training_data"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addTrainingDataPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddTrainingData with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv1.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CrossReference = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the AddTrainingDataOptions model
				addTrainingDataOptionsModel := new(discoveryv1.AddTrainingDataOptions)
				addTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				addTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addTrainingDataOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				addTrainingDataOptionsModel.Filter = core.StringPtr("testString")
				addTrainingDataOptionsModel.Examples = []discoveryv1.TrainingExample{*trainingExampleModel}
				addTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.AddTrainingData(addTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.AddTrainingData(addTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddTrainingData(addTrainingDataOptions *AddTrainingDataOptions)`, func() {
		version := "testString"
		addTrainingDataPath := "/v1/environments/testString/collections/testString/training_data"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addTrainingDataPath))
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
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}`)
				}))
			})
			It(`Invoke AddTrainingData successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv1.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CrossReference = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the AddTrainingDataOptions model
				addTrainingDataOptionsModel := new(discoveryv1.AddTrainingDataOptions)
				addTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				addTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addTrainingDataOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				addTrainingDataOptionsModel.Filter = core.StringPtr("testString")
				addTrainingDataOptionsModel.Examples = []discoveryv1.TrainingExample{*trainingExampleModel}
				addTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.AddTrainingDataWithContext(ctx, addTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.AddTrainingData(addTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.AddTrainingDataWithContext(ctx, addTrainingDataOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addTrainingDataPath))
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
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}`)
				}))
			})
			It(`Invoke AddTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.AddTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv1.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CrossReference = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the AddTrainingDataOptions model
				addTrainingDataOptionsModel := new(discoveryv1.AddTrainingDataOptions)
				addTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				addTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addTrainingDataOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				addTrainingDataOptionsModel.Filter = core.StringPtr("testString")
				addTrainingDataOptionsModel.Examples = []discoveryv1.TrainingExample{*trainingExampleModel}
				addTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.AddTrainingData(addTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddTrainingData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv1.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CrossReference = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the AddTrainingDataOptions model
				addTrainingDataOptionsModel := new(discoveryv1.AddTrainingDataOptions)
				addTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				addTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addTrainingDataOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				addTrainingDataOptionsModel.Filter = core.StringPtr("testString")
				addTrainingDataOptionsModel.Examples = []discoveryv1.TrainingExample{*trainingExampleModel}
				addTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.AddTrainingData(addTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddTrainingDataOptions model with no property values
				addTrainingDataOptionsModelNew := new(discoveryv1.AddTrainingDataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.AddTrainingData(addTrainingDataOptionsModelNew)
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
			It(`Invoke AddTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv1.TrainingExample)
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CrossReference = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))

				// Construct an instance of the AddTrainingDataOptions model
				addTrainingDataOptionsModel := new(discoveryv1.AddTrainingDataOptions)
				addTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				addTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				addTrainingDataOptionsModel.NaturalLanguageQuery = core.StringPtr("testString")
				addTrainingDataOptionsModel.Filter = core.StringPtr("testString")
				addTrainingDataOptionsModel.Examples = []discoveryv1.TrainingExample{*trainingExampleModel}
				addTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.AddTrainingData(addTrainingDataOptionsModel)
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
	Describe(`DeleteAllTrainingData(deleteAllTrainingDataOptions *DeleteAllTrainingDataOptions)`, func() {
		version := "testString"
		deleteAllTrainingDataPath := "/v1/environments/testString/collections/testString/training_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllTrainingDataPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAllTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteAllTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAllTrainingDataOptions model
				deleteAllTrainingDataOptionsModel := new(discoveryv1.DeleteAllTrainingDataOptions)
				deleteAllTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteAllTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				deleteAllTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteAllTrainingData(deleteAllTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAllTrainingData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteAllTrainingDataOptions model
				deleteAllTrainingDataOptionsModel := new(discoveryv1.DeleteAllTrainingDataOptions)
				deleteAllTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteAllTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				deleteAllTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteAllTrainingData(deleteAllTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAllTrainingDataOptions model with no property values
				deleteAllTrainingDataOptionsModelNew := new(discoveryv1.DeleteAllTrainingDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteAllTrainingData(deleteAllTrainingDataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrainingData(getTrainingDataOptions *GetTrainingDataOptions) - Operation response error`, func() {
		version := "testString"
		getTrainingDataPath := "/v1/environments/testString/collections/testString/training_data/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingDataPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTrainingData with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingDataOptions model
				getTrainingDataOptionsModel := new(discoveryv1.GetTrainingDataOptions)
				getTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetTrainingData(getTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetTrainingData(getTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrainingData(getTrainingDataOptions *GetTrainingDataOptions)`, func() {
		version := "testString"
		getTrainingDataPath := "/v1/environments/testString/collections/testString/training_data/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingDataPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}`)
				}))
			})
			It(`Invoke GetTrainingData successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetTrainingDataOptions model
				getTrainingDataOptionsModel := new(discoveryv1.GetTrainingDataOptions)
				getTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetTrainingDataWithContext(ctx, getTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetTrainingData(getTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetTrainingDataWithContext(ctx, getTrainingDataOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingDataPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"query_id": "QueryID", "natural_language_query": "NaturalLanguageQuery", "filter": "Filter", "examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}`)
				}))
			})
			It(`Invoke GetTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTrainingDataOptions model
				getTrainingDataOptionsModel := new(discoveryv1.GetTrainingDataOptions)
				getTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetTrainingData(getTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTrainingData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingDataOptions model
				getTrainingDataOptionsModel := new(discoveryv1.GetTrainingDataOptions)
				getTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetTrainingData(getTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTrainingDataOptions model with no property values
				getTrainingDataOptionsModelNew := new(discoveryv1.GetTrainingDataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetTrainingData(getTrainingDataOptionsModelNew)
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
			It(`Invoke GetTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingDataOptions model
				getTrainingDataOptionsModel := new(discoveryv1.GetTrainingDataOptions)
				getTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetTrainingData(getTrainingDataOptionsModel)
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
	Describe(`DeleteTrainingData(deleteTrainingDataOptions *DeleteTrainingDataOptions)`, func() {
		version := "testString"
		deleteTrainingDataPath := "/v1/environments/testString/collections/testString/training_data/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrainingDataPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTrainingData successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTrainingDataOptions model
				deleteTrainingDataOptionsModel := new(discoveryv1.DeleteTrainingDataOptions)
				deleteTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				deleteTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				deleteTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteTrainingData(deleteTrainingDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTrainingData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteTrainingDataOptions model
				deleteTrainingDataOptionsModel := new(discoveryv1.DeleteTrainingDataOptions)
				deleteTrainingDataOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteTrainingDataOptionsModel.CollectionID = core.StringPtr("testString")
				deleteTrainingDataOptionsModel.QueryID = core.StringPtr("testString")
				deleteTrainingDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteTrainingData(deleteTrainingDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTrainingDataOptions model with no property values
				deleteTrainingDataOptionsModelNew := new(discoveryv1.DeleteTrainingDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteTrainingData(deleteTrainingDataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTrainingExamples(listTrainingExamplesOptions *ListTrainingExamplesOptions) - Operation response error`, func() {
		version := "testString"
		listTrainingExamplesPath := "/v1/environments/testString/collections/testString/training_data/testString/examples"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingExamplesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTrainingExamples with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingExamplesOptions model
				listTrainingExamplesOptionsModel := new(discoveryv1.ListTrainingExamplesOptions)
				listTrainingExamplesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.QueryID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTrainingExamples(listTrainingExamplesOptions *ListTrainingExamplesOptions)`, func() {
		version := "testString"
		listTrainingExamplesPath := "/v1/environments/testString/collections/testString/training_data/testString/examples"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingExamplesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}`)
				}))
			})
			It(`Invoke ListTrainingExamples successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListTrainingExamplesOptions model
				listTrainingExamplesOptionsModel := new(discoveryv1.ListTrainingExamplesOptions)
				listTrainingExamplesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.QueryID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListTrainingExamplesWithContext(ctx, listTrainingExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListTrainingExamplesWithContext(ctx, listTrainingExamplesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTrainingExamplesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"examples": [{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}]}`)
				}))
			})
			It(`Invoke ListTrainingExamples successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListTrainingExamples(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTrainingExamplesOptions model
				listTrainingExamplesOptionsModel := new(discoveryv1.ListTrainingExamplesOptions)
				listTrainingExamplesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.QueryID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTrainingExamples with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingExamplesOptions model
				listTrainingExamplesOptionsModel := new(discoveryv1.ListTrainingExamplesOptions)
				listTrainingExamplesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.QueryID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTrainingExamplesOptions model with no property values
				listTrainingExamplesOptionsModelNew := new(discoveryv1.ListTrainingExamplesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModelNew)
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
			It(`Invoke ListTrainingExamples successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListTrainingExamplesOptions model
				listTrainingExamplesOptionsModel := new(discoveryv1.ListTrainingExamplesOptions)
				listTrainingExamplesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.CollectionID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.QueryID = core.StringPtr("testString")
				listTrainingExamplesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListTrainingExamples(listTrainingExamplesOptionsModel)
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
	Describe(`CreateTrainingExample(createTrainingExampleOptions *CreateTrainingExampleOptions) - Operation response error`, func() {
		version := "testString"
		createTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrainingExamplePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTrainingExample with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateTrainingExampleOptions model
				createTrainingExampleOptionsModel := new(discoveryv1.CreateTrainingExampleOptions)
				createTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.DocumentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				createTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				createTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateTrainingExample(createTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateTrainingExample(createTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTrainingExample(createTrainingExampleOptions *CreateTrainingExampleOptions)`, func() {
		version := "testString"
		createTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTrainingExamplePath))
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
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}`)
				}))
			})
			It(`Invoke CreateTrainingExample successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CreateTrainingExampleOptions model
				createTrainingExampleOptionsModel := new(discoveryv1.CreateTrainingExampleOptions)
				createTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.DocumentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				createTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				createTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateTrainingExampleWithContext(ctx, createTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateTrainingExample(createTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateTrainingExampleWithContext(ctx, createTrainingExampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTrainingExamplePath))
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
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}`)
				}))
			})
			It(`Invoke CreateTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTrainingExampleOptions model
				createTrainingExampleOptionsModel := new(discoveryv1.CreateTrainingExampleOptions)
				createTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.DocumentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				createTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				createTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateTrainingExample(createTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTrainingExample with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateTrainingExampleOptions model
				createTrainingExampleOptionsModel := new(discoveryv1.CreateTrainingExampleOptions)
				createTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.DocumentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				createTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				createTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateTrainingExample(createTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTrainingExampleOptions model with no property values
				createTrainingExampleOptionsModelNew := new(discoveryv1.CreateTrainingExampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateTrainingExample(createTrainingExampleOptionsModelNew)
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
			It(`Invoke CreateTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateTrainingExampleOptions model
				createTrainingExampleOptionsModel := new(discoveryv1.CreateTrainingExampleOptions)
				createTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.DocumentID = core.StringPtr("testString")
				createTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				createTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				createTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateTrainingExample(createTrainingExampleOptionsModel)
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
	Describe(`DeleteTrainingExample(deleteTrainingExampleOptions *DeleteTrainingExampleOptions)`, func() {
		version := "testString"
		deleteTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTrainingExamplePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := discoveryService.DeleteTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTrainingExampleOptions model
				deleteTrainingExampleOptionsModel := new(discoveryv1.DeleteTrainingExampleOptions)
				deleteTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteTrainingExample(deleteTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTrainingExample with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteTrainingExampleOptions model
				deleteTrainingExampleOptionsModel := new(discoveryv1.DeleteTrainingExampleOptions)
				deleteTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				deleteTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := discoveryService.DeleteTrainingExample(deleteTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTrainingExampleOptions model with no property values
				deleteTrainingExampleOptionsModelNew := new(discoveryv1.DeleteTrainingExampleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = discoveryService.DeleteTrainingExample(deleteTrainingExampleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTrainingExample(updateTrainingExampleOptions *UpdateTrainingExampleOptions) - Operation response error`, func() {
		version := "testString"
		updateTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrainingExamplePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTrainingExample with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateTrainingExampleOptions model
				updateTrainingExampleOptionsModel := new(discoveryv1.UpdateTrainingExampleOptions)
				updateTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				updateTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTrainingExample(updateTrainingExampleOptions *UpdateTrainingExampleOptions)`, func() {
		version := "testString"
		updateTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrainingExamplePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}`)
				}))
			})
			It(`Invoke UpdateTrainingExample successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTrainingExampleOptions model
				updateTrainingExampleOptionsModel := new(discoveryv1.UpdateTrainingExampleOptions)
				updateTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				updateTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateTrainingExampleWithContext(ctx, updateTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateTrainingExampleWithContext(ctx, updateTrainingExampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTrainingExamplePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}`)
				}))
			})
			It(`Invoke UpdateTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTrainingExampleOptions model
				updateTrainingExampleOptionsModel := new(discoveryv1.UpdateTrainingExampleOptions)
				updateTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				updateTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTrainingExample with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateTrainingExampleOptions model
				updateTrainingExampleOptionsModel := new(discoveryv1.UpdateTrainingExampleOptions)
				updateTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				updateTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTrainingExampleOptions model with no property values
				updateTrainingExampleOptionsModelNew := new(discoveryv1.UpdateTrainingExampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModelNew)
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
			It(`Invoke UpdateTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the UpdateTrainingExampleOptions model
				updateTrainingExampleOptionsModel := new(discoveryv1.UpdateTrainingExampleOptions)
				updateTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.CrossReference = core.StringPtr("testString")
				updateTrainingExampleOptionsModel.Relevance = core.Int64Ptr(int64(38))
				updateTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateTrainingExample(updateTrainingExampleOptionsModel)
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
	Describe(`GetTrainingExample(getTrainingExampleOptions *GetTrainingExampleOptions) - Operation response error`, func() {
		version := "testString"
		getTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingExamplePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTrainingExample with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingExampleOptions model
				getTrainingExampleOptionsModel := new(discoveryv1.GetTrainingExampleOptions)
				getTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetTrainingExample(getTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetTrainingExample(getTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrainingExample(getTrainingExampleOptions *GetTrainingExampleOptions)`, func() {
		version := "testString"
		getTrainingExamplePath := "/v1/environments/testString/collections/testString/training_data/testString/examples/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingExamplePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}`)
				}))
			})
			It(`Invoke GetTrainingExample successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetTrainingExampleOptions model
				getTrainingExampleOptionsModel := new(discoveryv1.GetTrainingExampleOptions)
				getTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetTrainingExampleWithContext(ctx, getTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetTrainingExample(getTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetTrainingExampleWithContext(ctx, getTrainingExampleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTrainingExamplePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"document_id": "DocumentID", "cross_reference": "CrossReference", "relevance": 9}`)
				}))
			})
			It(`Invoke GetTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetTrainingExample(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTrainingExampleOptions model
				getTrainingExampleOptionsModel := new(discoveryv1.GetTrainingExampleOptions)
				getTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetTrainingExample(getTrainingExampleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTrainingExample with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingExampleOptions model
				getTrainingExampleOptionsModel := new(discoveryv1.GetTrainingExampleOptions)
				getTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetTrainingExample(getTrainingExampleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTrainingExampleOptions model with no property values
				getTrainingExampleOptionsModelNew := new(discoveryv1.GetTrainingExampleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetTrainingExample(getTrainingExampleOptionsModelNew)
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
			It(`Invoke GetTrainingExample successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetTrainingExampleOptions model
				getTrainingExampleOptionsModel := new(discoveryv1.GetTrainingExampleOptions)
				getTrainingExampleOptionsModel.EnvironmentID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.CollectionID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.QueryID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.ExampleID = core.StringPtr("testString")
				getTrainingExampleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetTrainingExample(getTrainingExampleOptionsModel)
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
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		version := "testString"
		deleteUserDataPath := "/v1/user_data"
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
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
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
				deleteUserDataOptionsModel := new(discoveryv1.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = discoveryService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(discoveryv1.DeleteUserDataOptions)
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
				deleteUserDataOptionsModelNew := new(discoveryv1.DeleteUserDataOptions)
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
	Describe(`CreateEvent(createEventOptions *CreateEventOptions) - Operation response error`, func() {
		version := "testString"
		createEventPath := "/v1/events"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEventPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEvent with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the EventData model
				eventDataModel := new(discoveryv1.EventData)
				eventDataModel.EnvironmentID = core.StringPtr("testString")
				eventDataModel.SessionToken = core.StringPtr("testString")
				eventDataModel.ClientTimestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventDataModel.DisplayRank = core.Int64Ptr(int64(38))
				eventDataModel.CollectionID = core.StringPtr("testString")
				eventDataModel.DocumentID = core.StringPtr("testString")

				// Construct an instance of the CreateEventOptions model
				createEventOptionsModel := new(discoveryv1.CreateEventOptions)
				createEventOptionsModel.Type = core.StringPtr("click")
				createEventOptionsModel.Data = eventDataModel
				createEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateEvent(createEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateEvent(createEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEvent(createEventOptions *CreateEventOptions)`, func() {
		version := "testString"
		createEventPath := "/v1/events"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEventPath))
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
					fmt.Fprintf(res, "%s", `{"type": "click", "data": {"environment_id": "EnvironmentID", "session_token": "SessionToken", "client_timestamp": "2019-01-01T12:00:00.000Z", "display_rank": 11, "collection_id": "CollectionID", "document_id": "DocumentID", "query_id": "QueryID"}}`)
				}))
			})
			It(`Invoke CreateEvent successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the EventData model
				eventDataModel := new(discoveryv1.EventData)
				eventDataModel.EnvironmentID = core.StringPtr("testString")
				eventDataModel.SessionToken = core.StringPtr("testString")
				eventDataModel.ClientTimestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventDataModel.DisplayRank = core.Int64Ptr(int64(38))
				eventDataModel.CollectionID = core.StringPtr("testString")
				eventDataModel.DocumentID = core.StringPtr("testString")

				// Construct an instance of the CreateEventOptions model
				createEventOptionsModel := new(discoveryv1.CreateEventOptions)
				createEventOptionsModel.Type = core.StringPtr("click")
				createEventOptionsModel.Data = eventDataModel
				createEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateEventWithContext(ctx, createEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateEvent(createEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateEventWithContext(ctx, createEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEventPath))
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
					fmt.Fprintf(res, "%s", `{"type": "click", "data": {"environment_id": "EnvironmentID", "session_token": "SessionToken", "client_timestamp": "2019-01-01T12:00:00.000Z", "display_rank": 11, "collection_id": "CollectionID", "document_id": "DocumentID", "query_id": "QueryID"}}`)
				}))
			})
			It(`Invoke CreateEvent successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EventData model
				eventDataModel := new(discoveryv1.EventData)
				eventDataModel.EnvironmentID = core.StringPtr("testString")
				eventDataModel.SessionToken = core.StringPtr("testString")
				eventDataModel.ClientTimestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventDataModel.DisplayRank = core.Int64Ptr(int64(38))
				eventDataModel.CollectionID = core.StringPtr("testString")
				eventDataModel.DocumentID = core.StringPtr("testString")

				// Construct an instance of the CreateEventOptions model
				createEventOptionsModel := new(discoveryv1.CreateEventOptions)
				createEventOptionsModel.Type = core.StringPtr("click")
				createEventOptionsModel.Data = eventDataModel
				createEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateEvent(createEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEvent with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the EventData model
				eventDataModel := new(discoveryv1.EventData)
				eventDataModel.EnvironmentID = core.StringPtr("testString")
				eventDataModel.SessionToken = core.StringPtr("testString")
				eventDataModel.ClientTimestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventDataModel.DisplayRank = core.Int64Ptr(int64(38))
				eventDataModel.CollectionID = core.StringPtr("testString")
				eventDataModel.DocumentID = core.StringPtr("testString")

				// Construct an instance of the CreateEventOptions model
				createEventOptionsModel := new(discoveryv1.CreateEventOptions)
				createEventOptionsModel.Type = core.StringPtr("click")
				createEventOptionsModel.Data = eventDataModel
				createEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateEvent(createEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEventOptions model with no property values
				createEventOptionsModelNew := new(discoveryv1.CreateEventOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateEvent(createEventOptionsModelNew)
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
			It(`Invoke CreateEvent successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the EventData model
				eventDataModel := new(discoveryv1.EventData)
				eventDataModel.EnvironmentID = core.StringPtr("testString")
				eventDataModel.SessionToken = core.StringPtr("testString")
				eventDataModel.ClientTimestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventDataModel.DisplayRank = core.Int64Ptr(int64(38))
				eventDataModel.CollectionID = core.StringPtr("testString")
				eventDataModel.DocumentID = core.StringPtr("testString")

				// Construct an instance of the CreateEventOptions model
				createEventOptionsModel := new(discoveryv1.CreateEventOptions)
				createEventOptionsModel.Type = core.StringPtr("click")
				createEventOptionsModel.Data = eventDataModel
				createEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateEvent(createEventOptionsModel)
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
	Describe(`QueryLog(queryLogOptions *QueryLogOptions) - Operation response error`, func() {
		version := "testString"
		queryLogPath := "/v1/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryLogPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke QueryLog with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryLogOptions model
				queryLogOptionsModel := new(discoveryv1.QueryLogOptions)
				queryLogOptionsModel.Filter = core.StringPtr("testString")
				queryLogOptionsModel.Query = core.StringPtr("testString")
				queryLogOptionsModel.Count = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Sort = []string{"testString"}
				queryLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.QueryLog(queryLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.QueryLog(queryLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`QueryLog(queryLogOptions *QueryLogOptions)`, func() {
		version := "testString"
		queryLogPath := "/v1/logs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryLogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"environment_id": "EnvironmentID", "customer_id": "CustomerID", "document_type": "query", "natural_language_query": "NaturalLanguageQuery", "document_results": {"results": [{"position": 8, "document_id": "DocumentID", "score": 5, "confidence": 10, "collection_id": "CollectionID"}], "count": 5}, "created_timestamp": "2019-01-01T12:00:00.000Z", "client_timestamp": "2019-01-01T12:00:00.000Z", "query_id": "QueryID", "session_token": "SessionToken", "collection_id": "CollectionID", "display_rank": 11, "document_id": "DocumentID", "event_type": "click", "result_type": "document"}]}`)
				}))
			})
			It(`Invoke QueryLog successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the QueryLogOptions model
				queryLogOptionsModel := new(discoveryv1.QueryLogOptions)
				queryLogOptionsModel.Filter = core.StringPtr("testString")
				queryLogOptionsModel.Query = core.StringPtr("testString")
				queryLogOptionsModel.Count = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Sort = []string{"testString"}
				queryLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.QueryLogWithContext(ctx, queryLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.QueryLog(queryLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.QueryLogWithContext(ctx, queryLogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(queryLogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"matching_results": 15, "results": [{"environment_id": "EnvironmentID", "customer_id": "CustomerID", "document_type": "query", "natural_language_query": "NaturalLanguageQuery", "document_results": {"results": [{"position": 8, "document_id": "DocumentID", "score": 5, "confidence": 10, "collection_id": "CollectionID"}], "count": 5}, "created_timestamp": "2019-01-01T12:00:00.000Z", "client_timestamp": "2019-01-01T12:00:00.000Z", "query_id": "QueryID", "session_token": "SessionToken", "collection_id": "CollectionID", "display_rank": 11, "document_id": "DocumentID", "event_type": "click", "result_type": "document"}]}`)
				}))
			})
			It(`Invoke QueryLog successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.QueryLog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the QueryLogOptions model
				queryLogOptionsModel := new(discoveryv1.QueryLogOptions)
				queryLogOptionsModel.Filter = core.StringPtr("testString")
				queryLogOptionsModel.Query = core.StringPtr("testString")
				queryLogOptionsModel.Count = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Sort = []string{"testString"}
				queryLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.QueryLog(queryLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke QueryLog with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryLogOptions model
				queryLogOptionsModel := new(discoveryv1.QueryLogOptions)
				queryLogOptionsModel.Filter = core.StringPtr("testString")
				queryLogOptionsModel.Query = core.StringPtr("testString")
				queryLogOptionsModel.Count = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Sort = []string{"testString"}
				queryLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.QueryLog(queryLogOptionsModel)
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
			It(`Invoke QueryLog successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the QueryLogOptions model
				queryLogOptionsModel := new(discoveryv1.QueryLogOptions)
				queryLogOptionsModel.Filter = core.StringPtr("testString")
				queryLogOptionsModel.Query = core.StringPtr("testString")
				queryLogOptionsModel.Count = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Offset = core.Int64Ptr(int64(38))
				queryLogOptionsModel.Sort = []string{"testString"}
				queryLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.QueryLog(queryLogOptionsModel)
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
	Describe(`GetMetricsQuery(getMetricsQueryOptions *GetMetricsQueryOptions) - Operation response error`, func() {
		version := "testString"
		getMetricsQueryPath := "/v1/metrics/number_of_queries"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetricsQuery with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryOptions model
				getMetricsQueryOptionsModel := new(discoveryv1.GetMetricsQueryOptions)
				getMetricsQueryOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetMetricsQuery(getMetricsQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetMetricsQuery(getMetricsQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetricsQuery(getMetricsQueryOptions *GetMetricsQueryOptions)`, func() {
		version := "testString"
		getMetricsQueryPath := "/v1/metrics/number_of_queries"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQuery successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetMetricsQueryOptions model
				getMetricsQueryOptionsModel := new(discoveryv1.GetMetricsQueryOptions)
				getMetricsQueryOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetMetricsQueryWithContext(ctx, getMetricsQueryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetMetricsQuery(getMetricsQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetMetricsQueryWithContext(ctx, getMetricsQueryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetMetricsQuery(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetricsQueryOptions model
				getMetricsQueryOptionsModel := new(discoveryv1.GetMetricsQueryOptions)
				getMetricsQueryOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetMetricsQuery(getMetricsQueryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetricsQuery with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryOptions model
				getMetricsQueryOptionsModel := new(discoveryv1.GetMetricsQueryOptions)
				getMetricsQueryOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetMetricsQuery(getMetricsQueryOptionsModel)
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
			It(`Invoke GetMetricsQuery successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryOptions model
				getMetricsQueryOptionsModel := new(discoveryv1.GetMetricsQueryOptions)
				getMetricsQueryOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetMetricsQuery(getMetricsQueryOptionsModel)
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
	Describe(`GetMetricsQueryEvent(getMetricsQueryEventOptions *GetMetricsQueryEventOptions) - Operation response error`, func() {
		version := "testString"
		getMetricsQueryEventPath := "/v1/metrics/number_of_queries_with_event"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryEventPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetricsQueryEvent with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryEventOptions model
				getMetricsQueryEventOptionsModel := new(discoveryv1.GetMetricsQueryEventOptions)
				getMetricsQueryEventOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetricsQueryEvent(getMetricsQueryEventOptions *GetMetricsQueryEventOptions)`, func() {
		version := "testString"
		getMetricsQueryEventPath := "/v1/metrics/number_of_queries_with_event"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryEventPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQueryEvent successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetMetricsQueryEventOptions model
				getMetricsQueryEventOptionsModel := new(discoveryv1.GetMetricsQueryEventOptions)
				getMetricsQueryEventOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetMetricsQueryEventWithContext(ctx, getMetricsQueryEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetMetricsQueryEventWithContext(ctx, getMetricsQueryEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryEventPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQueryEvent successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetMetricsQueryEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetricsQueryEventOptions model
				getMetricsQueryEventOptionsModel := new(discoveryv1.GetMetricsQueryEventOptions)
				getMetricsQueryEventOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetricsQueryEvent with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryEventOptions model
				getMetricsQueryEventOptionsModel := new(discoveryv1.GetMetricsQueryEventOptions)
				getMetricsQueryEventOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptionsModel)
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
			It(`Invoke GetMetricsQueryEvent successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryEventOptions model
				getMetricsQueryEventOptionsModel := new(discoveryv1.GetMetricsQueryEventOptions)
				getMetricsQueryEventOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryEventOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptionsModel)
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
	Describe(`GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions) - Operation response error`, func() {
		version := "testString"
		getMetricsQueryNoResultsPath := "/v1/metrics/number_of_queries_with_no_search_results"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryNoResultsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetricsQueryNoResults with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryNoResultsOptions model
				getMetricsQueryNoResultsOptionsModel := new(discoveryv1.GetMetricsQueryNoResultsOptions)
				getMetricsQueryNoResultsOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryNoResultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions *GetMetricsQueryNoResultsOptions)`, func() {
		version := "testString"
		getMetricsQueryNoResultsPath := "/v1/metrics/number_of_queries_with_no_search_results"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryNoResultsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQueryNoResults successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetMetricsQueryNoResultsOptions model
				getMetricsQueryNoResultsOptionsModel := new(discoveryv1.GetMetricsQueryNoResultsOptions)
				getMetricsQueryNoResultsOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryNoResultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetMetricsQueryNoResultsWithContext(ctx, getMetricsQueryNoResultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetMetricsQueryNoResultsWithContext(ctx, getMetricsQueryNoResultsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryNoResultsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQueryNoResults successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetMetricsQueryNoResults(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetricsQueryNoResultsOptions model
				getMetricsQueryNoResultsOptionsModel := new(discoveryv1.GetMetricsQueryNoResultsOptions)
				getMetricsQueryNoResultsOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryNoResultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetricsQueryNoResults with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryNoResultsOptions model
				getMetricsQueryNoResultsOptionsModel := new(discoveryv1.GetMetricsQueryNoResultsOptions)
				getMetricsQueryNoResultsOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryNoResultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptionsModel)
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
			It(`Invoke GetMetricsQueryNoResults successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryNoResultsOptions model
				getMetricsQueryNoResultsOptionsModel := new(discoveryv1.GetMetricsQueryNoResultsOptions)
				getMetricsQueryNoResultsOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsQueryNoResultsOptionsModel.ResultType = core.StringPtr("document")
				getMetricsQueryNoResultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptionsModel)
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
	Describe(`GetMetricsEventRate(getMetricsEventRateOptions *GetMetricsEventRateOptions) - Operation response error`, func() {
		version := "testString"
		getMetricsEventRatePath := "/v1/metrics/event_rate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsEventRatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetricsEventRate with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsEventRateOptions model
				getMetricsEventRateOptionsModel := new(discoveryv1.GetMetricsEventRateOptions)
				getMetricsEventRateOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.ResultType = core.StringPtr("document")
				getMetricsEventRateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetMetricsEventRate(getMetricsEventRateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetMetricsEventRate(getMetricsEventRateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetricsEventRate(getMetricsEventRateOptions *GetMetricsEventRateOptions)`, func() {
		version := "testString"
		getMetricsEventRatePath := "/v1/metrics/event_rate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsEventRatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsEventRate successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetMetricsEventRateOptions model
				getMetricsEventRateOptionsModel := new(discoveryv1.GetMetricsEventRateOptions)
				getMetricsEventRateOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.ResultType = core.StringPtr("document")
				getMetricsEventRateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetMetricsEventRateWithContext(ctx, getMetricsEventRateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetMetricsEventRate(getMetricsEventRateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetMetricsEventRateWithContext(ctx, getMetricsEventRateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsEventRatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// TODO: Add check for start_time query parameter
					// TODO: Add check for end_time query parameter
					Expect(req.URL.Query()["result_type"]).To(Equal([]string{"document"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"interval": "Interval", "event_type": "EventType", "results": [{"key_as_string": "2019-01-01T12:00:00.000Z", "key": 3, "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsEventRate successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetMetricsEventRate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetricsEventRateOptions model
				getMetricsEventRateOptionsModel := new(discoveryv1.GetMetricsEventRateOptions)
				getMetricsEventRateOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.ResultType = core.StringPtr("document")
				getMetricsEventRateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetMetricsEventRate(getMetricsEventRateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetricsEventRate with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsEventRateOptions model
				getMetricsEventRateOptionsModel := new(discoveryv1.GetMetricsEventRateOptions)
				getMetricsEventRateOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.ResultType = core.StringPtr("document")
				getMetricsEventRateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetMetricsEventRate(getMetricsEventRateOptionsModel)
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
			It(`Invoke GetMetricsEventRate successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsEventRateOptions model
				getMetricsEventRateOptionsModel := new(discoveryv1.GetMetricsEventRateOptions)
				getMetricsEventRateOptionsModel.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				getMetricsEventRateOptionsModel.ResultType = core.StringPtr("document")
				getMetricsEventRateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetMetricsEventRate(getMetricsEventRateOptionsModel)
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
	Describe(`GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions) - Operation response error`, func() {
		version := "testString"
		getMetricsQueryTokenEventPath := "/v1/metrics/top_query_tokens_with_event_rate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryTokenEventPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMetricsQueryTokenEvent with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryTokenEventOptions model
				getMetricsQueryTokenEventOptionsModel := new(discoveryv1.GetMetricsQueryTokenEventOptions)
				getMetricsQueryTokenEventOptionsModel.Count = core.Int64Ptr(int64(38))
				getMetricsQueryTokenEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions *GetMetricsQueryTokenEventOptions)`, func() {
		version := "testString"
		getMetricsQueryTokenEventPath := "/v1/metrics/top_query_tokens_with_event_rate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryTokenEventPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"event_type": "EventType", "results": [{"key": "Key", "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQueryTokenEvent successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetMetricsQueryTokenEventOptions model
				getMetricsQueryTokenEventOptionsModel := new(discoveryv1.GetMetricsQueryTokenEventOptions)
				getMetricsQueryTokenEventOptionsModel.Count = core.Int64Ptr(int64(38))
				getMetricsQueryTokenEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetMetricsQueryTokenEventWithContext(ctx, getMetricsQueryTokenEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetMetricsQueryTokenEventWithContext(ctx, getMetricsQueryTokenEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMetricsQueryTokenEventPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["count"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"aggregations": [{"event_type": "EventType", "results": [{"key": "Key", "matching_results": 15, "event_rate": 9}]}]}`)
				}))
			})
			It(`Invoke GetMetricsQueryTokenEvent successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetMetricsQueryTokenEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMetricsQueryTokenEventOptions model
				getMetricsQueryTokenEventOptionsModel := new(discoveryv1.GetMetricsQueryTokenEventOptions)
				getMetricsQueryTokenEventOptionsModel.Count = core.Int64Ptr(int64(38))
				getMetricsQueryTokenEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMetricsQueryTokenEvent with error: Operation request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryTokenEventOptions model
				getMetricsQueryTokenEventOptionsModel := new(discoveryv1.GetMetricsQueryTokenEventOptions)
				getMetricsQueryTokenEventOptionsModel.Count = core.Int64Ptr(int64(38))
				getMetricsQueryTokenEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptionsModel)
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
			It(`Invoke GetMetricsQueryTokenEvent successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetMetricsQueryTokenEventOptions model
				getMetricsQueryTokenEventOptionsModel := new(discoveryv1.GetMetricsQueryTokenEventOptions)
				getMetricsQueryTokenEventOptionsModel.Count = core.Int64Ptr(int64(38))
				getMetricsQueryTokenEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptionsModel)
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
	Describe(`ListCredentials(listCredentialsOptions *ListCredentialsOptions) - Operation response error`, func() {
		version := "testString"
		listCredentialsPath := "/v1/environments/testString/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCredentialsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCredentials with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(discoveryv1.ListCredentialsOptions)
				listCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCredentials(listCredentialsOptions *ListCredentialsOptions)`, func() {
		version := "testString"
		listCredentialsPath := "/v1/environments/testString/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credentials": [{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}]}`)
				}))
			})
			It(`Invoke ListCredentials successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(discoveryv1.ListCredentialsOptions)
				listCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListCredentialsWithContext(ctx, listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListCredentialsWithContext(ctx, listCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credentials": [{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}]}`)
				}))
			})
			It(`Invoke ListCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(discoveryv1.ListCredentialsOptions)
				listCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCredentials with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(discoveryv1.ListCredentialsOptions)
				listCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListCredentials(listCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCredentialsOptions model with no property values
				listCredentialsOptionsModelNew := new(discoveryv1.ListCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListCredentials(listCredentialsOptionsModelNew)
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
			It(`Invoke ListCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListCredentialsOptions model
				listCredentialsOptionsModel := new(discoveryv1.ListCredentialsOptions)
				listCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				listCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListCredentials(listCredentialsOptionsModel)
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
	Describe(`CreateCredentials(createCredentialsOptions *CreateCredentialsOptions) - Operation response error`, func() {
		version := "testString"
		createCredentialsPath := "/v1/environments/testString/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCredentials with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the CreateCredentialsOptions model
				createCredentialsOptionsModel := new(discoveryv1.CreateCredentialsOptions)
				createCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCredentialsOptionsModel.SourceType = core.StringPtr("box")
				createCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				createCredentialsOptionsModel.Status = core.StringPtr("connected")
				createCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateCredentials(createCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateCredentials(createCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCredentials(createCredentialsOptions *CreateCredentialsOptions)`, func() {
		version := "testString"
		createCredentialsPath := "/v1/environments/testString/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialsPath))
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
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}`)
				}))
			})
			It(`Invoke CreateCredentials successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the CreateCredentialsOptions model
				createCredentialsOptionsModel := new(discoveryv1.CreateCredentialsOptions)
				createCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCredentialsOptionsModel.SourceType = core.StringPtr("box")
				createCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				createCredentialsOptionsModel.Status = core.StringPtr("connected")
				createCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateCredentialsWithContext(ctx, createCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateCredentials(createCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateCredentialsWithContext(ctx, createCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCredentialsPath))
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
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}`)
				}))
			})
			It(`Invoke CreateCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the CreateCredentialsOptions model
				createCredentialsOptionsModel := new(discoveryv1.CreateCredentialsOptions)
				createCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCredentialsOptionsModel.SourceType = core.StringPtr("box")
				createCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				createCredentialsOptionsModel.Status = core.StringPtr("connected")
				createCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateCredentials(createCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCredentials with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the CreateCredentialsOptions model
				createCredentialsOptionsModel := new(discoveryv1.CreateCredentialsOptions)
				createCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCredentialsOptionsModel.SourceType = core.StringPtr("box")
				createCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				createCredentialsOptionsModel.Status = core.StringPtr("connected")
				createCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateCredentials(createCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCredentialsOptions model with no property values
				createCredentialsOptionsModelNew := new(discoveryv1.CreateCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateCredentials(createCredentialsOptionsModelNew)
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
			It(`Invoke CreateCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the CreateCredentialsOptions model
				createCredentialsOptionsModel := new(discoveryv1.CreateCredentialsOptions)
				createCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				createCredentialsOptionsModel.SourceType = core.StringPtr("box")
				createCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				createCredentialsOptionsModel.Status = core.StringPtr("connected")
				createCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateCredentials(createCredentialsOptionsModel)
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
	Describe(`GetCredentials(getCredentialsOptions *GetCredentialsOptions) - Operation response error`, func() {
		version := "testString"
		getCredentialsPath := "/v1/environments/testString/credentials/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCredentialsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCredentials with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCredentialsOptions model
				getCredentialsOptionsModel := new(discoveryv1.GetCredentialsOptions)
				getCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				getCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				getCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetCredentials(getCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetCredentials(getCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCredentials(getCredentialsOptions *GetCredentialsOptions)`, func() {
		version := "testString"
		getCredentialsPath := "/v1/environments/testString/credentials/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}`)
				}))
			})
			It(`Invoke GetCredentials successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetCredentialsOptions model
				getCredentialsOptionsModel := new(discoveryv1.GetCredentialsOptions)
				getCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				getCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				getCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetCredentialsWithContext(ctx, getCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetCredentials(getCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetCredentialsWithContext(ctx, getCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}`)
				}))
			})
			It(`Invoke GetCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCredentialsOptions model
				getCredentialsOptionsModel := new(discoveryv1.GetCredentialsOptions)
				getCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				getCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				getCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetCredentials(getCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCredentials with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCredentialsOptions model
				getCredentialsOptionsModel := new(discoveryv1.GetCredentialsOptions)
				getCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				getCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				getCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetCredentials(getCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCredentialsOptions model with no property values
				getCredentialsOptionsModelNew := new(discoveryv1.GetCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetCredentials(getCredentialsOptionsModelNew)
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
			It(`Invoke GetCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetCredentialsOptions model
				getCredentialsOptionsModel := new(discoveryv1.GetCredentialsOptions)
				getCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				getCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				getCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetCredentials(getCredentialsOptionsModel)
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
	Describe(`UpdateCredentials(updateCredentialsOptions *UpdateCredentialsOptions) - Operation response error`, func() {
		version := "testString"
		updateCredentialsPath := "/v1/environments/testString/credentials/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCredentialsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCredentials with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the UpdateCredentialsOptions model
				updateCredentialsOptionsModel := new(discoveryv1.UpdateCredentialsOptions)
				updateCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				updateCredentialsOptionsModel.SourceType = core.StringPtr("box")
				updateCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				updateCredentialsOptionsModel.Status = core.StringPtr("connected")
				updateCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.UpdateCredentials(updateCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.UpdateCredentials(updateCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCredentials(updateCredentialsOptions *UpdateCredentialsOptions)`, func() {
		version := "testString"
		updateCredentialsPath := "/v1/environments/testString/credentials/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCredentialsPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}`)
				}))
			})
			It(`Invoke UpdateCredentials successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the UpdateCredentialsOptions model
				updateCredentialsOptionsModel := new(discoveryv1.UpdateCredentialsOptions)
				updateCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				updateCredentialsOptionsModel.SourceType = core.StringPtr("box")
				updateCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				updateCredentialsOptionsModel.Status = core.StringPtr("connected")
				updateCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.UpdateCredentialsWithContext(ctx, updateCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.UpdateCredentials(updateCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.UpdateCredentialsWithContext(ctx, updateCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCredentialsPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "source_type": "box", "credential_details": {"credential_type": "oauth2", "client_id": "ClientID", "enterprise_id": "EnterpriseID", "url": "URL", "username": "Username", "organization_url": "OrganizationURL", "site_collection.path": "SiteCollectionPath", "client_secret": "ClientSecret", "public_key_id": "PublicKeyID", "private_key": "PrivateKey", "passphrase": "Passphrase", "password": "Password", "gateway_id": "GatewayID", "source_version": "online", "web_application_url": "WebApplicationURL", "domain": "Domain", "endpoint": "Endpoint", "access_key_id": "AccessKeyID", "secret_access_key": "SecretAccessKey"}, "status": "connected"}`)
				}))
			})
			It(`Invoke UpdateCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.UpdateCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the UpdateCredentialsOptions model
				updateCredentialsOptionsModel := new(discoveryv1.UpdateCredentialsOptions)
				updateCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				updateCredentialsOptionsModel.SourceType = core.StringPtr("box")
				updateCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				updateCredentialsOptionsModel.Status = core.StringPtr("connected")
				updateCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.UpdateCredentials(updateCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCredentials with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the UpdateCredentialsOptions model
				updateCredentialsOptionsModel := new(discoveryv1.UpdateCredentialsOptions)
				updateCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				updateCredentialsOptionsModel.SourceType = core.StringPtr("box")
				updateCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				updateCredentialsOptionsModel.Status = core.StringPtr("connected")
				updateCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.UpdateCredentials(updateCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCredentialsOptions model with no property values
				updateCredentialsOptionsModelNew := new(discoveryv1.UpdateCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.UpdateCredentials(updateCredentialsOptionsModelNew)
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
			It(`Invoke UpdateCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")

				// Construct an instance of the UpdateCredentialsOptions model
				updateCredentialsOptionsModel := new(discoveryv1.UpdateCredentialsOptions)
				updateCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				updateCredentialsOptionsModel.SourceType = core.StringPtr("box")
				updateCredentialsOptionsModel.CredentialDetails = credentialDetailsModel
				updateCredentialsOptionsModel.Status = core.StringPtr("connected")
				updateCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.UpdateCredentials(updateCredentialsOptionsModel)
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
	Describe(`DeleteCredentials(deleteCredentialsOptions *DeleteCredentialsOptions) - Operation response error`, func() {
		version := "testString"
		deleteCredentialsPath := "/v1/environments/testString/credentials/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCredentialsPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCredentials with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCredentialsOptions model
				deleteCredentialsOptionsModel := new(discoveryv1.DeleteCredentialsOptions)
				deleteCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.DeleteCredentials(deleteCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.DeleteCredentials(deleteCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCredentials(deleteCredentialsOptions *DeleteCredentialsOptions)`, func() {
		version := "testString"
		deleteCredentialsPath := "/v1/environments/testString/credentials/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCredentialsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteCredentials successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCredentialsOptions model
				deleteCredentialsOptionsModel := new(discoveryv1.DeleteCredentialsOptions)
				deleteCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.DeleteCredentialsWithContext(ctx, deleteCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.DeleteCredentials(deleteCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.DeleteCredentialsWithContext(ctx, deleteCredentialsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteCredentialsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credential_id": "CredentialID", "status": "deleted"}`)
				}))
			})
			It(`Invoke DeleteCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.DeleteCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCredentialsOptions model
				deleteCredentialsOptionsModel := new(discoveryv1.DeleteCredentialsOptions)
				deleteCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteCredentials(deleteCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCredentials with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCredentialsOptions model
				deleteCredentialsOptionsModel := new(discoveryv1.DeleteCredentialsOptions)
				deleteCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.DeleteCredentials(deleteCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCredentialsOptions model with no property values
				deleteCredentialsOptionsModelNew := new(discoveryv1.DeleteCredentialsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.DeleteCredentials(deleteCredentialsOptionsModelNew)
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
			It(`Invoke DeleteCredentials successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteCredentialsOptions model
				deleteCredentialsOptionsModel := new(discoveryv1.DeleteCredentialsOptions)
				deleteCredentialsOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.CredentialID = core.StringPtr("testString")
				deleteCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.DeleteCredentials(deleteCredentialsOptionsModel)
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
	Describe(`ListGateways(listGatewaysOptions *ListGatewaysOptions) - Operation response error`, func() {
		version := "testString"
		listGatewaysPath := "/v1/environments/testString/gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGateways with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(discoveryv1.ListGatewaysOptions)
				listGatewaysOptionsModel.EnvironmentID = core.StringPtr("testString")
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListGateways(listGatewaysOptions *ListGatewaysOptions)`, func() {
		version := "testString"
		listGatewaysPath := "/v1/environments/testString/gateways"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateways": [{"gateway_id": "GatewayID", "name": "Name", "status": "connected", "token": "Token", "token_id": "TokenID"}]}`)
				}))
			})
			It(`Invoke ListGateways successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(discoveryv1.ListGatewaysOptions)
				listGatewaysOptionsModel.EnvironmentID = core.StringPtr("testString")
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.ListGatewaysWithContext(ctx, listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.ListGatewaysWithContext(ctx, listGatewaysOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateways": [{"gateway_id": "GatewayID", "name": "Name", "status": "connected", "token": "Token", "token_id": "TokenID"}]}`)
				}))
			})
			It(`Invoke ListGateways successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.ListGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(discoveryv1.ListGatewaysOptions)
				listGatewaysOptionsModel.EnvironmentID = core.StringPtr("testString")
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListGateways with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(discoveryv1.ListGatewaysOptions)
				listGatewaysOptionsModel.EnvironmentID = core.StringPtr("testString")
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.ListGateways(listGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListGatewaysOptions model with no property values
				listGatewaysOptionsModelNew := new(discoveryv1.ListGatewaysOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.ListGateways(listGatewaysOptionsModelNew)
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
			It(`Invoke ListGateways successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the ListGatewaysOptions model
				listGatewaysOptionsModel := new(discoveryv1.ListGatewaysOptions)
				listGatewaysOptionsModel.EnvironmentID = core.StringPtr("testString")
				listGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.ListGateways(listGatewaysOptionsModel)
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
	Describe(`CreateGateway(createGatewayOptions *CreateGatewayOptions) - Operation response error`, func() {
		version := "testString"
		createGatewayPath := "/v1/environments/testString/gateways"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGateway with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(discoveryv1.CreateGatewayOptions)
				createGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				createGatewayOptionsModel.Name = core.StringPtr("testString")
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGateway(createGatewayOptions *CreateGatewayOptions)`, func() {
		version := "testString"
		createGatewayPath := "/v1/environments/testString/gateways"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"gateway_id": "GatewayID", "name": "Name", "status": "connected", "token": "Token", "token_id": "TokenID"}`)
				}))
			})
			It(`Invoke CreateGateway successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(discoveryv1.CreateGatewayOptions)
				createGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				createGatewayOptionsModel.Name = core.StringPtr("testString")
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.CreateGatewayWithContext(ctx, createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.CreateGatewayWithContext(ctx, createGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createGatewayPath))
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
					fmt.Fprintf(res, "%s", `{"gateway_id": "GatewayID", "name": "Name", "status": "connected", "token": "Token", "token_id": "TokenID"}`)
				}))
			})
			It(`Invoke CreateGateway successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.CreateGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(discoveryv1.CreateGatewayOptions)
				createGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				createGatewayOptionsModel.Name = core.StringPtr("testString")
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGateway with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(discoveryv1.CreateGatewayOptions)
				createGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				createGatewayOptionsModel.Name = core.StringPtr("testString")
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.CreateGateway(createGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGatewayOptions model with no property values
				createGatewayOptionsModelNew := new(discoveryv1.CreateGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.CreateGateway(createGatewayOptionsModelNew)
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
			It(`Invoke CreateGateway successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the CreateGatewayOptions model
				createGatewayOptionsModel := new(discoveryv1.CreateGatewayOptions)
				createGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				createGatewayOptionsModel.Name = core.StringPtr("testString")
				createGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.CreateGateway(createGatewayOptionsModel)
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
	Describe(`GetGateway(getGatewayOptions *GetGatewayOptions) - Operation response error`, func() {
		version := "testString"
		getGatewayPath := "/v1/environments/testString/gateways/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGateway with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(discoveryv1.GetGatewayOptions)
				getGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				getGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGateway(getGatewayOptions *GetGatewayOptions)`, func() {
		version := "testString"
		getGatewayPath := "/v1/environments/testString/gateways/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateway_id": "GatewayID", "name": "Name", "status": "connected", "token": "Token", "token_id": "TokenID"}`)
				}))
			})
			It(`Invoke GetGateway successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(discoveryv1.GetGatewayOptions)
				getGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				getGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.GetGatewayWithContext(ctx, getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.GetGatewayWithContext(ctx, getGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateway_id": "GatewayID", "name": "Name", "status": "connected", "token": "Token", "token_id": "TokenID"}`)
				}))
			})
			It(`Invoke GetGateway successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.GetGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(discoveryv1.GetGatewayOptions)
				getGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				getGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGateway with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(discoveryv1.GetGatewayOptions)
				getGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				getGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.GetGateway(getGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayOptions model with no property values
				getGatewayOptionsModelNew := new(discoveryv1.GetGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.GetGateway(getGatewayOptionsModelNew)
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
			It(`Invoke GetGateway successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the GetGatewayOptions model
				getGatewayOptionsModel := new(discoveryv1.GetGatewayOptions)
				getGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				getGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				getGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.GetGateway(getGatewayOptionsModel)
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
	Describe(`DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions) - Operation response error`, func() {
		version := "testString"
		deleteGatewayPath := "/v1/environments/testString/gateways/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteGateway with error: Operation response processing error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(discoveryv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := discoveryService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				discoveryService.EnableRetries(0, 0)
				result, response, operationErr = discoveryService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGateway(deleteGatewayOptions *DeleteGatewayOptions)`, func() {
		version := "testString"
		deleteGatewayPath := "/v1/environments/testString/gateways/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateway_id": "GatewayID", "status": "Status"}`)
				}))
			})
			It(`Invoke DeleteGateway successfully with retries`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())
				discoveryService.EnableRetries(0, 0)

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(discoveryv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := discoveryService.DeleteGatewayWithContext(ctx, deleteGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				discoveryService.DisableRetries()
				result, response, operationErr := discoveryService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = discoveryService.DeleteGatewayWithContext(ctx, deleteGatewayOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"gateway_id": "GatewayID", "status": "Status"}`)
				}))
			})
			It(`Invoke DeleteGateway successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := discoveryService.DeleteGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(discoveryv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = discoveryService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteGateway with error: Operation validation and request error`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(discoveryv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := discoveryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := discoveryService.DeleteGateway(deleteGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteGatewayOptions model with no property values
				deleteGatewayOptionsModelNew := new(discoveryv1.DeleteGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = discoveryService.DeleteGateway(deleteGatewayOptionsModelNew)
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
			It(`Invoke DeleteGateway successfully`, func() {
				discoveryService, serviceErr := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(discoveryService).ToNot(BeNil())

				// Construct an instance of the DeleteGatewayOptions model
				deleteGatewayOptionsModel := new(discoveryv1.DeleteGatewayOptions)
				deleteGatewayOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteGatewayOptionsModel.GatewayID = core.StringPtr("testString")
				deleteGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := discoveryService.DeleteGateway(deleteGatewayOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			discoveryService, _ := discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
				URL:           "http://discoveryv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAddDocumentOptions successfully`, func() {
				// Construct an instance of the AddDocumentOptions model
				environmentID := "testString"
				collectionID := "testString"
				addDocumentOptionsModel := discoveryService.NewAddDocumentOptions(environmentID, collectionID)
				addDocumentOptionsModel.SetEnvironmentID("testString")
				addDocumentOptionsModel.SetCollectionID("testString")
				addDocumentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				addDocumentOptionsModel.SetFilename("testString")
				addDocumentOptionsModel.SetFileContentType("application/json")
				addDocumentOptionsModel.SetMetadata("testString")
				addDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addDocumentOptionsModel).ToNot(BeNil())
				Expect(addDocumentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(addDocumentOptionsModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.FileContentType).To(Equal(core.StringPtr("application/json")))
				Expect(addDocumentOptionsModel.Metadata).To(Equal(core.StringPtr("testString")))
				Expect(addDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddTrainingDataOptions successfully`, func() {
				// Construct an instance of the TrainingExample model
				trainingExampleModel := new(discoveryv1.TrainingExample)
				Expect(trainingExampleModel).ToNot(BeNil())
				trainingExampleModel.DocumentID = core.StringPtr("testString")
				trainingExampleModel.CrossReference = core.StringPtr("testString")
				trainingExampleModel.Relevance = core.Int64Ptr(int64(38))
				Expect(trainingExampleModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(trainingExampleModel.CrossReference).To(Equal(core.StringPtr("testString")))
				Expect(trainingExampleModel.Relevance).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the AddTrainingDataOptions model
				environmentID := "testString"
				collectionID := "testString"
				addTrainingDataOptionsModel := discoveryService.NewAddTrainingDataOptions(environmentID, collectionID)
				addTrainingDataOptionsModel.SetEnvironmentID("testString")
				addTrainingDataOptionsModel.SetCollectionID("testString")
				addTrainingDataOptionsModel.SetNaturalLanguageQuery("testString")
				addTrainingDataOptionsModel.SetFilter("testString")
				addTrainingDataOptionsModel.SetExamples([]discoveryv1.TrainingExample{*trainingExampleModel})
				addTrainingDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addTrainingDataOptionsModel).ToNot(BeNil())
				Expect(addTrainingDataOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(addTrainingDataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(addTrainingDataOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(addTrainingDataOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(addTrainingDataOptionsModel.Examples).To(Equal([]discoveryv1.TrainingExample{*trainingExampleModel}))
				Expect(addTrainingDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConfiguration successfully`, func() {
				name := "testString"
				model, err := discoveryService.NewConfiguration(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCollectionOptions successfully`, func() {
				// Construct an instance of the CreateCollectionOptions model
				environmentID := "testString"
				createCollectionOptionsName := "testString"
				createCollectionOptionsModel := discoveryService.NewCreateCollectionOptions(environmentID, createCollectionOptionsName)
				createCollectionOptionsModel.SetEnvironmentID("testString")
				createCollectionOptionsModel.SetName("testString")
				createCollectionOptionsModel.SetDescription("testString")
				createCollectionOptionsModel.SetConfigurationID("testString")
				createCollectionOptionsModel.SetLanguage("en")
				createCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectionOptionsModel).ToNot(BeNil())
				Expect(createCollectionOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.ConfigurationID).To(Equal(core.StringPtr("testString")))
				Expect(createCollectionOptionsModel.Language).To(Equal(core.StringPtr("en")))
				Expect(createCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateConfigurationOptions successfully`, func() {
				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				Expect(fontSettingModel).ToNot(BeNil())
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")
				Expect(fontSettingModel.Level).To(Equal(core.Int64Ptr(int64(38))))
				Expect(fontSettingModel.MinSize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(fontSettingModel.MaxSize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(fontSettingModel.Bold).To(Equal(core.BoolPtr(true)))
				Expect(fontSettingModel.Italic).To(Equal(core.BoolPtr(true)))
				Expect(fontSettingModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				Expect(pdfHeadingDetectionModel).ToNot(BeNil())
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				Expect(pdfHeadingDetectionModel.Fonts).To(Equal([]discoveryv1.FontSetting{*fontSettingModel}))

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				Expect(pdfSettingsModel).ToNot(BeNil())
				pdfSettingsModel.Heading = pdfHeadingDetectionModel
				Expect(pdfSettingsModel.Heading).To(Equal(pdfHeadingDetectionModel))

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				Expect(wordStyleModel).ToNot(BeNil())
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}
				Expect(wordStyleModel.Level).To(Equal(core.Int64Ptr(int64(38))))
				Expect(wordStyleModel.Names).To(Equal([]string{"testString"}))

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				Expect(wordHeadingDetectionModel).ToNot(BeNil())
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}
				Expect(wordHeadingDetectionModel.Fonts).To(Equal([]discoveryv1.FontSetting{*fontSettingModel}))
				Expect(wordHeadingDetectionModel.Styles).To(Equal([]discoveryv1.WordStyle{*wordStyleModel}))

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				Expect(wordSettingsModel).ToNot(BeNil())
				wordSettingsModel.Heading = wordHeadingDetectionModel
				Expect(wordSettingsModel.Heading).To(Equal(wordHeadingDetectionModel))

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				Expect(xPathPatternsModel).ToNot(BeNil())
				xPathPatternsModel.Xpaths = []string{"testString"}
				Expect(xPathPatternsModel.Xpaths).To(Equal([]string{"testString"}))

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				Expect(htmlSettingsModel).ToNot(BeNil())
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}
				Expect(htmlSettingsModel.ExcludeTagsCompletely).To(Equal([]string{"testString"}))
				Expect(htmlSettingsModel.ExcludeTagsKeepContent).To(Equal([]string{"testString"}))
				Expect(htmlSettingsModel.KeepContent).To(Equal(xPathPatternsModel))
				Expect(htmlSettingsModel.ExcludeContent).To(Equal(xPathPatternsModel))
				Expect(htmlSettingsModel.KeepTagAttributes).To(Equal([]string{"testString"}))
				Expect(htmlSettingsModel.ExcludeTagAttributes).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				Expect(segmentSettingsModel).ToNot(BeNil())
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}
				Expect(segmentSettingsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(segmentSettingsModel.SelectorTags).To(Equal([]string{"testString"}))
				Expect(segmentSettingsModel.AnnotatedFields).To(Equal([]string{"testString"}))

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				Expect(normalizationOperationModel).ToNot(BeNil())
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")
				Expect(normalizationOperationModel.Operation).To(Equal(core.StringPtr("copy")))
				Expect(normalizationOperationModel.SourceField).To(Equal(core.StringPtr("testString")))
				Expect(normalizationOperationModel.DestinationField).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				Expect(conversionsModel).ToNot(BeNil())
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)
				Expect(conversionsModel.PDF).To(Equal(pdfSettingsModel))
				Expect(conversionsModel.Word).To(Equal(wordSettingsModel))
				Expect(conversionsModel.HTML).To(Equal(htmlSettingsModel))
				Expect(conversionsModel.Segment).To(Equal(segmentSettingsModel))
				Expect(conversionsModel.JSONNormalizations).To(Equal([]discoveryv1.NormalizationOperation{*normalizationOperationModel}))
				Expect(conversionsModel.ImageTextRecognition).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				Expect(nluEnrichmentKeywordsModel).ToNot(BeNil())
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))
				Expect(nluEnrichmentKeywordsModel.Sentiment).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentKeywordsModel.Emotion).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentKeywordsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				Expect(nluEnrichmentEntitiesModel).ToNot(BeNil())
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")
				Expect(nluEnrichmentEntitiesModel.Sentiment).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.Emotion).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(nluEnrichmentEntitiesModel.Mentions).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.MentionTypes).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.SentenceLocations).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				Expect(nluEnrichmentSentimentModel).ToNot(BeNil())
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}
				Expect(nluEnrichmentSentimentModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentSentimentModel.Targets).To(Equal([]string{"testString"}))

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				Expect(nluEnrichmentEmotionModel).ToNot(BeNil())
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}
				Expect(nluEnrichmentEmotionModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEmotionModel.Targets).To(Equal([]string{"testString"}))

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				Expect(nluEnrichmentSemanticRolesModel).ToNot(BeNil())
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))
				Expect(nluEnrichmentSemanticRolesModel.Entities).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentSemanticRolesModel.Keywords).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentSemanticRolesModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				Expect(nluEnrichmentRelationsModel).ToNot(BeNil())
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")
				Expect(nluEnrichmentRelationsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				Expect(nluEnrichmentConceptsModel).ToNot(BeNil())
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))
				Expect(nluEnrichmentConceptsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				Expect(nluEnrichmentFeaturesModel).ToNot(BeNil())
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel
				Expect(nluEnrichmentFeaturesModel.Keywords).To(Equal(nluEnrichmentKeywordsModel))
				Expect(nluEnrichmentFeaturesModel.Entities).To(Equal(nluEnrichmentEntitiesModel))
				Expect(nluEnrichmentFeaturesModel.Sentiment).To(Equal(nluEnrichmentSentimentModel))
				Expect(nluEnrichmentFeaturesModel.Emotion).To(Equal(nluEnrichmentEmotionModel))
				Expect(nluEnrichmentFeaturesModel.Categories).To(Equal(make(map[string]interface{})))
				Expect(nluEnrichmentFeaturesModel.SemanticRoles).To(Equal(nluEnrichmentSemanticRolesModel))
				Expect(nluEnrichmentFeaturesModel.Relations).To(Equal(nluEnrichmentRelationsModel))
				Expect(nluEnrichmentFeaturesModel.Concepts).To(Equal(nluEnrichmentConceptsModel))

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				Expect(enrichmentOptionsModel).ToNot(BeNil())
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")
				Expect(enrichmentOptionsModel.Features).To(Equal(nluEnrichmentFeaturesModel))
				Expect(enrichmentOptionsModel.Language).To(Equal(core.StringPtr("ar")))
				Expect(enrichmentOptionsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				Expect(enrichmentModel).ToNot(BeNil())
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel
				Expect(enrichmentModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.DestinationField).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.SourceField).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.Overwrite).To(Equal(core.BoolPtr(true)))
				Expect(enrichmentModel.Enrichment).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.IgnoreDownstreamErrors).To(Equal(core.BoolPtr(true)))
				Expect(enrichmentModel.Options).To(Equal(enrichmentOptionsModel))

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				Expect(sourceScheduleModel).ToNot(BeNil())
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")
				Expect(sourceScheduleModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(sourceScheduleModel.TimeZone).To(Equal(core.StringPtr("testString")))
				Expect(sourceScheduleModel.Frequency).To(Equal(core.StringPtr("daily")))

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				Expect(sourceOptionsFolderModel).ToNot(BeNil())
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsFolderModel.OwnerUserID).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsFolderModel.FolderID).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsFolderModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				Expect(sourceOptionsObjectModel).ToNot(BeNil())
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsObjectModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsObjectModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				Expect(sourceOptionsSiteCollModel).ToNot(BeNil())
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsSiteCollModel.SiteCollectionPath).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsSiteCollModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				Expect(sourceOptionsWebCrawlModel).ToNot(BeNil())
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}
				Expect(sourceOptionsWebCrawlModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsWebCrawlModel.LimitToStartingHosts).To(Equal(core.BoolPtr(true)))
				Expect(sourceOptionsWebCrawlModel.CrawlSpeed).To(Equal(core.StringPtr("gentle")))
				Expect(sourceOptionsWebCrawlModel.AllowUntrustedCertificate).To(Equal(core.BoolPtr(true)))
				Expect(sourceOptionsWebCrawlModel.MaximumHops).To(Equal(core.Int64Ptr(int64(38))))
				Expect(sourceOptionsWebCrawlModel.RequestTimeout).To(Equal(core.Int64Ptr(int64(38))))
				Expect(sourceOptionsWebCrawlModel.OverrideRobotsTxt).To(Equal(core.BoolPtr(true)))
				Expect(sourceOptionsWebCrawlModel.Blacklist).To(Equal([]string{"testString"}))

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				Expect(sourceOptionsBucketsModel).ToNot(BeNil())
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsBucketsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsBucketsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				Expect(sourceOptionsModel).ToNot(BeNil())
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)
				Expect(sourceOptionsModel.Folders).To(Equal([]discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}))
				Expect(sourceOptionsModel.Objects).To(Equal([]discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}))
				Expect(sourceOptionsModel.SiteCollections).To(Equal([]discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}))
				Expect(sourceOptionsModel.Urls).To(Equal([]discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}))
				Expect(sourceOptionsModel.Buckets).To(Equal([]discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}))
				Expect(sourceOptionsModel.CrawlAllBuckets).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				Expect(sourceModel).ToNot(BeNil())
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel
				Expect(sourceModel.Type).To(Equal(core.StringPtr("box")))
				Expect(sourceModel.CredentialID).To(Equal(core.StringPtr("testString")))
				Expect(sourceModel.Schedule).To(Equal(sourceScheduleModel))
				Expect(sourceModel.Options).To(Equal(sourceOptionsModel))

				// Construct an instance of the CreateConfigurationOptions model
				environmentID := "testString"
				createConfigurationOptionsName := "testString"
				createConfigurationOptionsModel := discoveryService.NewCreateConfigurationOptions(environmentID, createConfigurationOptionsName)
				createConfigurationOptionsModel.SetEnvironmentID("testString")
				createConfigurationOptionsModel.SetName("testString")
				createConfigurationOptionsModel.SetDescription("testString")
				createConfigurationOptionsModel.SetConversions(conversionsModel)
				createConfigurationOptionsModel.SetEnrichments([]discoveryv1.Enrichment{*enrichmentModel})
				createConfigurationOptionsModel.SetNormalizations([]discoveryv1.NormalizationOperation{*normalizationOperationModel})
				createConfigurationOptionsModel.SetSource(sourceModel)
				createConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigurationOptionsModel).ToNot(BeNil())
				Expect(createConfigurationOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigurationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createConfigurationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createConfigurationOptionsModel.Conversions).To(Equal(conversionsModel))
				Expect(createConfigurationOptionsModel.Enrichments).To(Equal([]discoveryv1.Enrichment{*enrichmentModel}))
				Expect(createConfigurationOptionsModel.Normalizations).To(Equal([]discoveryv1.NormalizationOperation{*normalizationOperationModel}))
				Expect(createConfigurationOptionsModel.Source).To(Equal(sourceModel))
				Expect(createConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCredentialsOptions successfully`, func() {
				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				Expect(credentialDetailsModel).ToNot(BeNil())
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")
				Expect(credentialDetailsModel.CredentialType).To(Equal(core.StringPtr("oauth2")))
				Expect(credentialDetailsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.OrganizationURL).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.SiteCollectionPath).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.ClientSecret).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.PublicKeyID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.PrivateKey).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Passphrase).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.SourceVersion).To(Equal(core.StringPtr("online")))
				Expect(credentialDetailsModel.WebApplicationURL).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Endpoint).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.AccessKeyID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.SecretAccessKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateCredentialsOptions model
				environmentID := "testString"
				createCredentialsOptionsModel := discoveryService.NewCreateCredentialsOptions(environmentID)
				createCredentialsOptionsModel.SetEnvironmentID("testString")
				createCredentialsOptionsModel.SetSourceType("box")
				createCredentialsOptionsModel.SetCredentialDetails(credentialDetailsModel)
				createCredentialsOptionsModel.SetStatus("connected")
				createCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCredentialsOptionsModel).ToNot(BeNil())
				Expect(createCredentialsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createCredentialsOptionsModel.SourceType).To(Equal(core.StringPtr("box")))
				Expect(createCredentialsOptionsModel.CredentialDetails).To(Equal(credentialDetailsModel))
				Expect(createCredentialsOptionsModel.Status).To(Equal(core.StringPtr("connected")))
				Expect(createCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEnvironmentOptions successfully`, func() {
				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsName := "testString"
				createEnvironmentOptionsModel := discoveryService.NewCreateEnvironmentOptions(createEnvironmentOptionsName)
				createEnvironmentOptionsModel.SetName("testString")
				createEnvironmentOptionsModel.SetDescription("testString")
				createEnvironmentOptionsModel.SetSize("LT")
				createEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEnvironmentOptionsModel).ToNot(BeNil())
				Expect(createEnvironmentOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createEnvironmentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createEnvironmentOptionsModel.Size).To(Equal(core.StringPtr("LT")))
				Expect(createEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEventOptions successfully`, func() {
				// Construct an instance of the EventData model
				eventDataModel := new(discoveryv1.EventData)
				Expect(eventDataModel).ToNot(BeNil())
				eventDataModel.EnvironmentID = core.StringPtr("testString")
				eventDataModel.SessionToken = core.StringPtr("testString")
				eventDataModel.ClientTimestamp = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				eventDataModel.DisplayRank = core.Int64Ptr(int64(38))
				eventDataModel.CollectionID = core.StringPtr("testString")
				eventDataModel.DocumentID = core.StringPtr("testString")
				Expect(eventDataModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(eventDataModel.SessionToken).To(Equal(core.StringPtr("testString")))
				Expect(eventDataModel.ClientTimestamp).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(eventDataModel.DisplayRank).To(Equal(core.Int64Ptr(int64(38))))
				Expect(eventDataModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(eventDataModel.DocumentID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateEventOptions model
				createEventOptionsType := "click"
				var createEventOptionsData *discoveryv1.EventData = nil
				createEventOptionsModel := discoveryService.NewCreateEventOptions(createEventOptionsType, createEventOptionsData)
				createEventOptionsModel.SetType("click")
				createEventOptionsModel.SetData(eventDataModel)
				createEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEventOptionsModel).ToNot(BeNil())
				Expect(createEventOptionsModel.Type).To(Equal(core.StringPtr("click")))
				Expect(createEventOptionsModel.Data).To(Equal(eventDataModel))
				Expect(createEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateExpansionsOptions successfully`, func() {
				// Construct an instance of the Expansion model
				expansionModel := new(discoveryv1.Expansion)
				Expect(expansionModel).ToNot(BeNil())
				expansionModel.InputTerms = []string{"testString"}
				expansionModel.ExpandedTerms = []string{"testString"}
				Expect(expansionModel.InputTerms).To(Equal([]string{"testString"}))
				Expect(expansionModel.ExpandedTerms).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateExpansionsOptions model
				environmentID := "testString"
				collectionID := "testString"
				createExpansionsOptionsExpansions := []discoveryv1.Expansion{}
				createExpansionsOptionsModel := discoveryService.NewCreateExpansionsOptions(environmentID, collectionID, createExpansionsOptionsExpansions)
				createExpansionsOptionsModel.SetEnvironmentID("testString")
				createExpansionsOptionsModel.SetCollectionID("testString")
				createExpansionsOptionsModel.SetExpansions([]discoveryv1.Expansion{*expansionModel})
				createExpansionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createExpansionsOptionsModel).ToNot(BeNil())
				Expect(createExpansionsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createExpansionsOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(createExpansionsOptionsModel.Expansions).To(Equal([]discoveryv1.Expansion{*expansionModel}))
				Expect(createExpansionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGatewayOptions successfully`, func() {
				// Construct an instance of the CreateGatewayOptions model
				environmentID := "testString"
				createGatewayOptionsModel := discoveryService.NewCreateGatewayOptions(environmentID)
				createGatewayOptionsModel.SetEnvironmentID("testString")
				createGatewayOptionsModel.SetName("testString")
				createGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGatewayOptionsModel).ToNot(BeNil())
				Expect(createGatewayOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateStopwordListOptions successfully`, func() {
				// Construct an instance of the CreateStopwordListOptions model
				environmentID := "testString"
				collectionID := "testString"
				stopwordFile := CreateMockReader("This is a mock file.")
				stopwordFilename := "testString"
				createStopwordListOptionsModel := discoveryService.NewCreateStopwordListOptions(environmentID, collectionID, stopwordFile, stopwordFilename)
				createStopwordListOptionsModel.SetEnvironmentID("testString")
				createStopwordListOptionsModel.SetCollectionID("testString")
				createStopwordListOptionsModel.SetStopwordFile(CreateMockReader("This is a mock file."))
				createStopwordListOptionsModel.SetStopwordFilename("testString")
				createStopwordListOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createStopwordListOptionsModel).ToNot(BeNil())
				Expect(createStopwordListOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createStopwordListOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(createStopwordListOptionsModel.StopwordFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createStopwordListOptionsModel.StopwordFilename).To(Equal(core.StringPtr("testString")))
				Expect(createStopwordListOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTokenizationDictionaryOptions successfully`, func() {
				// Construct an instance of the TokenDictRule model
				tokenDictRuleModel := new(discoveryv1.TokenDictRule)
				Expect(tokenDictRuleModel).ToNot(BeNil())
				tokenDictRuleModel.Text = core.StringPtr("testString")
				tokenDictRuleModel.Tokens = []string{"testString"}
				tokenDictRuleModel.Readings = []string{"testString"}
				tokenDictRuleModel.PartOfSpeech = core.StringPtr("testString")
				Expect(tokenDictRuleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(tokenDictRuleModel.Tokens).To(Equal([]string{"testString"}))
				Expect(tokenDictRuleModel.Readings).To(Equal([]string{"testString"}))
				Expect(tokenDictRuleModel.PartOfSpeech).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateTokenizationDictionaryOptions model
				environmentID := "testString"
				collectionID := "testString"
				createTokenizationDictionaryOptionsModel := discoveryService.NewCreateTokenizationDictionaryOptions(environmentID, collectionID)
				createTokenizationDictionaryOptionsModel.SetEnvironmentID("testString")
				createTokenizationDictionaryOptionsModel.SetCollectionID("testString")
				createTokenizationDictionaryOptionsModel.SetTokenizationRules([]discoveryv1.TokenDictRule{*tokenDictRuleModel})
				createTokenizationDictionaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTokenizationDictionaryOptionsModel).ToNot(BeNil())
				Expect(createTokenizationDictionaryOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createTokenizationDictionaryOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(createTokenizationDictionaryOptionsModel.TokenizationRules).To(Equal([]discoveryv1.TokenDictRule{*tokenDictRuleModel}))
				Expect(createTokenizationDictionaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTrainingExampleOptions successfully`, func() {
				// Construct an instance of the CreateTrainingExampleOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				createTrainingExampleOptionsModel := discoveryService.NewCreateTrainingExampleOptions(environmentID, collectionID, queryID)
				createTrainingExampleOptionsModel.SetEnvironmentID("testString")
				createTrainingExampleOptionsModel.SetCollectionID("testString")
				createTrainingExampleOptionsModel.SetQueryID("testString")
				createTrainingExampleOptionsModel.SetDocumentID("testString")
				createTrainingExampleOptionsModel.SetCrossReference("testString")
				createTrainingExampleOptionsModel.SetRelevance(int64(38))
				createTrainingExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTrainingExampleOptionsModel).ToNot(BeNil())
				Expect(createTrainingExampleOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingExampleOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingExampleOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingExampleOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingExampleOptionsModel.CrossReference).To(Equal(core.StringPtr("testString")))
				Expect(createTrainingExampleOptionsModel.Relevance).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createTrainingExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllTrainingDataOptions successfully`, func() {
				// Construct an instance of the DeleteAllTrainingDataOptions model
				environmentID := "testString"
				collectionID := "testString"
				deleteAllTrainingDataOptionsModel := discoveryService.NewDeleteAllTrainingDataOptions(environmentID, collectionID)
				deleteAllTrainingDataOptionsModel.SetEnvironmentID("testString")
				deleteAllTrainingDataOptionsModel.SetCollectionID("testString")
				deleteAllTrainingDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllTrainingDataOptionsModel).ToNot(BeNil())
				Expect(deleteAllTrainingDataOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAllTrainingDataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAllTrainingDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCollectionOptions successfully`, func() {
				// Construct an instance of the DeleteCollectionOptions model
				environmentID := "testString"
				collectionID := "testString"
				deleteCollectionOptionsModel := discoveryService.NewDeleteCollectionOptions(environmentID, collectionID)
				deleteCollectionOptionsModel.SetEnvironmentID("testString")
				deleteCollectionOptionsModel.SetCollectionID("testString")
				deleteCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCollectionOptionsModel).ToNot(BeNil())
				Expect(deleteCollectionOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigurationOptions successfully`, func() {
				// Construct an instance of the DeleteConfigurationOptions model
				environmentID := "testString"
				configurationID := "testString"
				deleteConfigurationOptionsModel := discoveryService.NewDeleteConfigurationOptions(environmentID, configurationID)
				deleteConfigurationOptionsModel.SetEnvironmentID("testString")
				deleteConfigurationOptionsModel.SetConfigurationID("testString")
				deleteConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigurationOptionsModel).ToNot(BeNil())
				Expect(deleteConfigurationOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigurationOptionsModel.ConfigurationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCredentialsOptions successfully`, func() {
				// Construct an instance of the DeleteCredentialsOptions model
				environmentID := "testString"
				credentialID := "testString"
				deleteCredentialsOptionsModel := discoveryService.NewDeleteCredentialsOptions(environmentID, credentialID)
				deleteCredentialsOptionsModel.SetEnvironmentID("testString")
				deleteCredentialsOptionsModel.SetCredentialID("testString")
				deleteCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCredentialsOptionsModel).ToNot(BeNil())
				Expect(deleteCredentialsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCredentialsOptionsModel.CredentialID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDocumentOptions model
				environmentID := "testString"
				collectionID := "testString"
				documentID := "testString"
				deleteDocumentOptionsModel := discoveryService.NewDeleteDocumentOptions(environmentID, collectionID, documentID)
				deleteDocumentOptionsModel.SetEnvironmentID("testString")
				deleteDocumentOptionsModel.SetCollectionID("testString")
				deleteDocumentOptionsModel.SetDocumentID("testString")
				deleteDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDocumentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEnvironmentOptions successfully`, func() {
				// Construct an instance of the DeleteEnvironmentOptions model
				environmentID := "testString"
				deleteEnvironmentOptionsModel := discoveryService.NewDeleteEnvironmentOptions(environmentID)
				deleteEnvironmentOptionsModel.SetEnvironmentID("testString")
				deleteEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEnvironmentOptionsModel).ToNot(BeNil())
				Expect(deleteEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteExpansionsOptions successfully`, func() {
				// Construct an instance of the DeleteExpansionsOptions model
				environmentID := "testString"
				collectionID := "testString"
				deleteExpansionsOptionsModel := discoveryService.NewDeleteExpansionsOptions(environmentID, collectionID)
				deleteExpansionsOptionsModel.SetEnvironmentID("testString")
				deleteExpansionsOptionsModel.SetCollectionID("testString")
				deleteExpansionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteExpansionsOptionsModel).ToNot(BeNil())
				Expect(deleteExpansionsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteExpansionsOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteExpansionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteGatewayOptions model
				environmentID := "testString"
				gatewayID := "testString"
				deleteGatewayOptionsModel := discoveryService.NewDeleteGatewayOptions(environmentID, gatewayID)
				deleteGatewayOptionsModel.SetEnvironmentID("testString")
				deleteGatewayOptionsModel.SetGatewayID("testString")
				deleteGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteGatewayOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGatewayOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteStopwordListOptions successfully`, func() {
				// Construct an instance of the DeleteStopwordListOptions model
				environmentID := "testString"
				collectionID := "testString"
				deleteStopwordListOptionsModel := discoveryService.NewDeleteStopwordListOptions(environmentID, collectionID)
				deleteStopwordListOptionsModel.SetEnvironmentID("testString")
				deleteStopwordListOptionsModel.SetCollectionID("testString")
				deleteStopwordListOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteStopwordListOptionsModel).ToNot(BeNil())
				Expect(deleteStopwordListOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteStopwordListOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteStopwordListOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTokenizationDictionaryOptions successfully`, func() {
				// Construct an instance of the DeleteTokenizationDictionaryOptions model
				environmentID := "testString"
				collectionID := "testString"
				deleteTokenizationDictionaryOptionsModel := discoveryService.NewDeleteTokenizationDictionaryOptions(environmentID, collectionID)
				deleteTokenizationDictionaryOptionsModel.SetEnvironmentID("testString")
				deleteTokenizationDictionaryOptionsModel.SetCollectionID("testString")
				deleteTokenizationDictionaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTokenizationDictionaryOptionsModel).ToNot(BeNil())
				Expect(deleteTokenizationDictionaryOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTokenizationDictionaryOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTokenizationDictionaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTrainingDataOptions successfully`, func() {
				// Construct an instance of the DeleteTrainingDataOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				deleteTrainingDataOptionsModel := discoveryService.NewDeleteTrainingDataOptions(environmentID, collectionID, queryID)
				deleteTrainingDataOptionsModel.SetEnvironmentID("testString")
				deleteTrainingDataOptionsModel.SetCollectionID("testString")
				deleteTrainingDataOptionsModel.SetQueryID("testString")
				deleteTrainingDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTrainingDataOptionsModel).ToNot(BeNil())
				Expect(deleteTrainingDataOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingDataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingDataOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTrainingExampleOptions successfully`, func() {
				// Construct an instance of the DeleteTrainingExampleOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				exampleID := "testString"
				deleteTrainingExampleOptionsModel := discoveryService.NewDeleteTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				deleteTrainingExampleOptionsModel.SetEnvironmentID("testString")
				deleteTrainingExampleOptionsModel.SetCollectionID("testString")
				deleteTrainingExampleOptionsModel.SetQueryID("testString")
				deleteTrainingExampleOptionsModel.SetExampleID("testString")
				deleteTrainingExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTrainingExampleOptionsModel).ToNot(BeNil())
				Expect(deleteTrainingExampleOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingExampleOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingExampleOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingExampleOptionsModel.ExampleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTrainingExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewEnrichment successfully`, func() {
				destinationField := "testString"
				sourceField := "testString"
				enrichment := "testString"
				model, err := discoveryService.NewEnrichment(destinationField, sourceField, enrichment)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewEventData successfully`, func() {
				environmentID := "testString"
				sessionToken := "testString"
				collectionID := "testString"
				documentID := "testString"
				model, err := discoveryService.NewEventData(environmentID, sessionToken, collectionID, documentID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExpansion successfully`, func() {
				expandedTerms := []string{"testString"}
				model, err := discoveryService.NewExpansion(expandedTerms)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExpansions successfully`, func() {
				expansions := []discoveryv1.Expansion{}
				model, err := discoveryService.NewExpansions(expansions)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFederatedQueryNoticesOptions successfully`, func() {
				// Construct an instance of the FederatedQueryNoticesOptions model
				environmentID := "testString"
				collectionIds := []string{"testString"}
				federatedQueryNoticesOptionsModel := discoveryService.NewFederatedQueryNoticesOptions(environmentID, collectionIds)
				federatedQueryNoticesOptionsModel.SetEnvironmentID("testString")
				federatedQueryNoticesOptionsModel.SetCollectionIds([]string{"testString"})
				federatedQueryNoticesOptionsModel.SetFilter("testString")
				federatedQueryNoticesOptionsModel.SetQuery("testString")
				federatedQueryNoticesOptionsModel.SetNaturalLanguageQuery("testString")
				federatedQueryNoticesOptionsModel.SetAggregation("testString")
				federatedQueryNoticesOptionsModel.SetCount(int64(38))
				federatedQueryNoticesOptionsModel.SetReturn([]string{"testString"})
				federatedQueryNoticesOptionsModel.SetOffset(int64(38))
				federatedQueryNoticesOptionsModel.SetSort([]string{"testString"})
				federatedQueryNoticesOptionsModel.SetHighlight(true)
				federatedQueryNoticesOptionsModel.SetDeduplicateField("testString")
				federatedQueryNoticesOptionsModel.SetSimilar(true)
				federatedQueryNoticesOptionsModel.SetSimilarDocumentIds([]string{"testString"})
				federatedQueryNoticesOptionsModel.SetSimilarFields([]string{"testString"})
				federatedQueryNoticesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(federatedQueryNoticesOptionsModel).ToNot(BeNil())
				Expect(federatedQueryNoticesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryNoticesOptionsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(federatedQueryNoticesOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryNoticesOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryNoticesOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryNoticesOptionsModel.Aggregation).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryNoticesOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(federatedQueryNoticesOptionsModel.Return).To(Equal([]string{"testString"}))
				Expect(federatedQueryNoticesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(federatedQueryNoticesOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(federatedQueryNoticesOptionsModel.Highlight).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryNoticesOptionsModel.DeduplicateField).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryNoticesOptionsModel.Similar).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryNoticesOptionsModel.SimilarDocumentIds).To(Equal([]string{"testString"}))
				Expect(federatedQueryNoticesOptionsModel.SimilarFields).To(Equal([]string{"testString"}))
				Expect(federatedQueryNoticesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFederatedQueryOptions successfully`, func() {
				// Construct an instance of the FederatedQueryOptions model
				environmentID := "testString"
				federatedQueryOptionsCollectionIds := "testString"
				federatedQueryOptionsModel := discoveryService.NewFederatedQueryOptions(environmentID, federatedQueryOptionsCollectionIds)
				federatedQueryOptionsModel.SetEnvironmentID("testString")
				federatedQueryOptionsModel.SetCollectionIds("testString")
				federatedQueryOptionsModel.SetFilter("testString")
				federatedQueryOptionsModel.SetQuery("testString")
				federatedQueryOptionsModel.SetNaturalLanguageQuery("testString")
				federatedQueryOptionsModel.SetPassages(true)
				federatedQueryOptionsModel.SetAggregation("testString")
				federatedQueryOptionsModel.SetCount(int64(38))
				federatedQueryOptionsModel.SetReturn("testString")
				federatedQueryOptionsModel.SetOffset(int64(38))
				federatedQueryOptionsModel.SetSort("testString")
				federatedQueryOptionsModel.SetHighlight(true)
				federatedQueryOptionsModel.SetPassagesFields("testString")
				federatedQueryOptionsModel.SetPassagesCount(int64(100))
				federatedQueryOptionsModel.SetPassagesCharacters(int64(50))
				federatedQueryOptionsModel.SetDeduplicate(true)
				federatedQueryOptionsModel.SetDeduplicateField("testString")
				federatedQueryOptionsModel.SetSimilar(true)
				federatedQueryOptionsModel.SetSimilarDocumentIds("testString")
				federatedQueryOptionsModel.SetSimilarFields("testString")
				federatedQueryOptionsModel.SetBias("testString")
				federatedQueryOptionsModel.SetXWatsonLoggingOptOut(true)
				federatedQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(federatedQueryOptionsModel).ToNot(BeNil())
				Expect(federatedQueryOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.CollectionIds).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Passages).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryOptionsModel.Aggregation).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(federatedQueryOptionsModel.Return).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(federatedQueryOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Highlight).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryOptionsModel.PassagesFields).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.PassagesCount).To(Equal(core.Int64Ptr(int64(100))))
				Expect(federatedQueryOptionsModel.PassagesCharacters).To(Equal(core.Int64Ptr(int64(50))))
				Expect(federatedQueryOptionsModel.Deduplicate).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryOptionsModel.DeduplicateField).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Similar).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryOptionsModel.SimilarDocumentIds).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.SimilarFields).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.Bias).To(Equal(core.StringPtr("testString")))
				Expect(federatedQueryOptionsModel.XWatsonLoggingOptOut).To(Equal(core.BoolPtr(true)))
				Expect(federatedQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAutocompletionOptions successfully`, func() {
				// Construct an instance of the GetAutocompletionOptions model
				environmentID := "testString"
				collectionID := "testString"
				prefix := "testString"
				getAutocompletionOptionsModel := discoveryService.NewGetAutocompletionOptions(environmentID, collectionID, prefix)
				getAutocompletionOptionsModel.SetEnvironmentID("testString")
				getAutocompletionOptionsModel.SetCollectionID("testString")
				getAutocompletionOptionsModel.SetPrefix("testString")
				getAutocompletionOptionsModel.SetField("testString")
				getAutocompletionOptionsModel.SetCount(int64(38))
				getAutocompletionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAutocompletionOptionsModel).ToNot(BeNil())
				Expect(getAutocompletionOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.Prefix).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.Field).To(Equal(core.StringPtr("testString")))
				Expect(getAutocompletionOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getAutocompletionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCollectionOptions successfully`, func() {
				// Construct an instance of the GetCollectionOptions model
				environmentID := "testString"
				collectionID := "testString"
				getCollectionOptionsModel := discoveryService.NewGetCollectionOptions(environmentID, collectionID)
				getCollectionOptionsModel.SetEnvironmentID("testString")
				getCollectionOptionsModel.SetCollectionID("testString")
				getCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionOptionsModel).ToNot(BeNil())
				Expect(getCollectionOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigurationOptions successfully`, func() {
				// Construct an instance of the GetConfigurationOptions model
				environmentID := "testString"
				configurationID := "testString"
				getConfigurationOptionsModel := discoveryService.NewGetConfigurationOptions(environmentID, configurationID)
				getConfigurationOptionsModel.SetEnvironmentID("testString")
				getConfigurationOptionsModel.SetConfigurationID("testString")
				getConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigurationOptionsModel).ToNot(BeNil())
				Expect(getConfigurationOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigurationOptionsModel.ConfigurationID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCredentialsOptions successfully`, func() {
				// Construct an instance of the GetCredentialsOptions model
				environmentID := "testString"
				credentialID := "testString"
				getCredentialsOptionsModel := discoveryService.NewGetCredentialsOptions(environmentID, credentialID)
				getCredentialsOptionsModel.SetEnvironmentID("testString")
				getCredentialsOptionsModel.SetCredentialID("testString")
				getCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCredentialsOptionsModel).ToNot(BeNil())
				Expect(getCredentialsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getCredentialsOptionsModel.CredentialID).To(Equal(core.StringPtr("testString")))
				Expect(getCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentStatusOptions successfully`, func() {
				// Construct an instance of the GetDocumentStatusOptions model
				environmentID := "testString"
				collectionID := "testString"
				documentID := "testString"
				getDocumentStatusOptionsModel := discoveryService.NewGetDocumentStatusOptions(environmentID, collectionID, documentID)
				getDocumentStatusOptionsModel.SetEnvironmentID("testString")
				getDocumentStatusOptionsModel.SetCollectionID("testString")
				getDocumentStatusOptionsModel.SetDocumentID("testString")
				getDocumentStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentStatusOptionsModel).ToNot(BeNil())
				Expect(getDocumentStatusOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentStatusOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentStatusOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnvironmentOptions successfully`, func() {
				// Construct an instance of the GetEnvironmentOptions model
				environmentID := "testString"
				getEnvironmentOptionsModel := discoveryService.NewGetEnvironmentOptions(environmentID)
				getEnvironmentOptionsModel.SetEnvironmentID("testString")
				getEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnvironmentOptionsModel).ToNot(BeNil())
				Expect(getEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayOptions successfully`, func() {
				// Construct an instance of the GetGatewayOptions model
				environmentID := "testString"
				gatewayID := "testString"
				getGatewayOptionsModel := discoveryService.NewGetGatewayOptions(environmentID, gatewayID)
				getGatewayOptionsModel.SetEnvironmentID("testString")
				getGatewayOptionsModel.SetGatewayID("testString")
				getGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayOptionsModel).ToNot(BeNil())
				Expect(getGatewayOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayOptionsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetricsEventRateOptions successfully`, func() {
				// Construct an instance of the GetMetricsEventRateOptions model
				getMetricsEventRateOptionsModel := discoveryService.NewGetMetricsEventRateOptions()
				getMetricsEventRateOptionsModel.SetStartTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsEventRateOptionsModel.SetEndTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsEventRateOptionsModel.SetResultType("document")
				getMetricsEventRateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetricsEventRateOptionsModel).ToNot(BeNil())
				Expect(getMetricsEventRateOptionsModel.StartTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsEventRateOptionsModel.EndTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsEventRateOptionsModel.ResultType).To(Equal(core.StringPtr("document")))
				Expect(getMetricsEventRateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetricsQueryEventOptions successfully`, func() {
				// Construct an instance of the GetMetricsQueryEventOptions model
				getMetricsQueryEventOptionsModel := discoveryService.NewGetMetricsQueryEventOptions()
				getMetricsQueryEventOptionsModel.SetStartTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsQueryEventOptionsModel.SetEndTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsQueryEventOptionsModel.SetResultType("document")
				getMetricsQueryEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetricsQueryEventOptionsModel).ToNot(BeNil())
				Expect(getMetricsQueryEventOptionsModel.StartTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsQueryEventOptionsModel.EndTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsQueryEventOptionsModel.ResultType).To(Equal(core.StringPtr("document")))
				Expect(getMetricsQueryEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetricsQueryNoResultsOptions successfully`, func() {
				// Construct an instance of the GetMetricsQueryNoResultsOptions model
				getMetricsQueryNoResultsOptionsModel := discoveryService.NewGetMetricsQueryNoResultsOptions()
				getMetricsQueryNoResultsOptionsModel.SetStartTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsQueryNoResultsOptionsModel.SetEndTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsQueryNoResultsOptionsModel.SetResultType("document")
				getMetricsQueryNoResultsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetricsQueryNoResultsOptionsModel).ToNot(BeNil())
				Expect(getMetricsQueryNoResultsOptionsModel.StartTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsQueryNoResultsOptionsModel.EndTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsQueryNoResultsOptionsModel.ResultType).To(Equal(core.StringPtr("document")))
				Expect(getMetricsQueryNoResultsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetricsQueryOptions successfully`, func() {
				// Construct an instance of the GetMetricsQueryOptions model
				getMetricsQueryOptionsModel := discoveryService.NewGetMetricsQueryOptions()
				getMetricsQueryOptionsModel.SetStartTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsQueryOptionsModel.SetEndTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				getMetricsQueryOptionsModel.SetResultType("document")
				getMetricsQueryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetricsQueryOptionsModel).ToNot(BeNil())
				Expect(getMetricsQueryOptionsModel.StartTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsQueryOptionsModel.EndTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(getMetricsQueryOptionsModel.ResultType).To(Equal(core.StringPtr("document")))
				Expect(getMetricsQueryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMetricsQueryTokenEventOptions successfully`, func() {
				// Construct an instance of the GetMetricsQueryTokenEventOptions model
				getMetricsQueryTokenEventOptionsModel := discoveryService.NewGetMetricsQueryTokenEventOptions()
				getMetricsQueryTokenEventOptionsModel.SetCount(int64(38))
				getMetricsQueryTokenEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMetricsQueryTokenEventOptionsModel).ToNot(BeNil())
				Expect(getMetricsQueryTokenEventOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getMetricsQueryTokenEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetStopwordListStatusOptions successfully`, func() {
				// Construct an instance of the GetStopwordListStatusOptions model
				environmentID := "testString"
				collectionID := "testString"
				getStopwordListStatusOptionsModel := discoveryService.NewGetStopwordListStatusOptions(environmentID, collectionID)
				getStopwordListStatusOptionsModel.SetEnvironmentID("testString")
				getStopwordListStatusOptionsModel.SetCollectionID("testString")
				getStopwordListStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getStopwordListStatusOptionsModel).ToNot(BeNil())
				Expect(getStopwordListStatusOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getStopwordListStatusOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getStopwordListStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTokenizationDictionaryStatusOptions successfully`, func() {
				// Construct an instance of the GetTokenizationDictionaryStatusOptions model
				environmentID := "testString"
				collectionID := "testString"
				getTokenizationDictionaryStatusOptionsModel := discoveryService.NewGetTokenizationDictionaryStatusOptions(environmentID, collectionID)
				getTokenizationDictionaryStatusOptionsModel.SetEnvironmentID("testString")
				getTokenizationDictionaryStatusOptionsModel.SetCollectionID("testString")
				getTokenizationDictionaryStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTokenizationDictionaryStatusOptionsModel).ToNot(BeNil())
				Expect(getTokenizationDictionaryStatusOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getTokenizationDictionaryStatusOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getTokenizationDictionaryStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTrainingDataOptions successfully`, func() {
				// Construct an instance of the GetTrainingDataOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				getTrainingDataOptionsModel := discoveryService.NewGetTrainingDataOptions(environmentID, collectionID, queryID)
				getTrainingDataOptionsModel.SetEnvironmentID("testString")
				getTrainingDataOptionsModel.SetCollectionID("testString")
				getTrainingDataOptionsModel.SetQueryID("testString")
				getTrainingDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTrainingDataOptionsModel).ToNot(BeNil())
				Expect(getTrainingDataOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingDataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingDataOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTrainingExampleOptions successfully`, func() {
				// Construct an instance of the GetTrainingExampleOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				exampleID := "testString"
				getTrainingExampleOptionsModel := discoveryService.NewGetTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				getTrainingExampleOptionsModel.SetEnvironmentID("testString")
				getTrainingExampleOptionsModel.SetCollectionID("testString")
				getTrainingExampleOptionsModel.SetQueryID("testString")
				getTrainingExampleOptionsModel.SetExampleID("testString")
				getTrainingExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTrainingExampleOptionsModel).ToNot(BeNil())
				Expect(getTrainingExampleOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingExampleOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingExampleOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingExampleOptionsModel.ExampleID).To(Equal(core.StringPtr("testString")))
				Expect(getTrainingExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCollectionFieldsOptions successfully`, func() {
				// Construct an instance of the ListCollectionFieldsOptions model
				environmentID := "testString"
				collectionID := "testString"
				listCollectionFieldsOptionsModel := discoveryService.NewListCollectionFieldsOptions(environmentID, collectionID)
				listCollectionFieldsOptionsModel.SetEnvironmentID("testString")
				listCollectionFieldsOptionsModel.SetCollectionID("testString")
				listCollectionFieldsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectionFieldsOptionsModel).ToNot(BeNil())
				Expect(listCollectionFieldsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listCollectionFieldsOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(listCollectionFieldsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCollectionsOptions successfully`, func() {
				// Construct an instance of the ListCollectionsOptions model
				environmentID := "testString"
				listCollectionsOptionsModel := discoveryService.NewListCollectionsOptions(environmentID)
				listCollectionsOptionsModel.SetEnvironmentID("testString")
				listCollectionsOptionsModel.SetName("testString")
				listCollectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectionsOptionsModel).ToNot(BeNil())
				Expect(listCollectionsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listCollectionsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listCollectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigurationsOptions successfully`, func() {
				// Construct an instance of the ListConfigurationsOptions model
				environmentID := "testString"
				listConfigurationsOptionsModel := discoveryService.NewListConfigurationsOptions(environmentID)
				listConfigurationsOptionsModel.SetEnvironmentID("testString")
				listConfigurationsOptionsModel.SetName("testString")
				listConfigurationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigurationsOptionsModel).ToNot(BeNil())
				Expect(listConfigurationsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigurationsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listConfigurationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCredentialsOptions successfully`, func() {
				// Construct an instance of the ListCredentialsOptions model
				environmentID := "testString"
				listCredentialsOptionsModel := discoveryService.NewListCredentialsOptions(environmentID)
				listCredentialsOptionsModel.SetEnvironmentID("testString")
				listCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCredentialsOptionsModel).ToNot(BeNil())
				Expect(listCredentialsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnvironmentsOptions successfully`, func() {
				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := discoveryService.NewListEnvironmentsOptions()
				listEnvironmentsOptionsModel.SetName("testString")
				listEnvironmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnvironmentsOptionsModel).ToNot(BeNil())
				Expect(listEnvironmentsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listEnvironmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListExpansionsOptions successfully`, func() {
				// Construct an instance of the ListExpansionsOptions model
				environmentID := "testString"
				collectionID := "testString"
				listExpansionsOptionsModel := discoveryService.NewListExpansionsOptions(environmentID, collectionID)
				listExpansionsOptionsModel.SetEnvironmentID("testString")
				listExpansionsOptionsModel.SetCollectionID("testString")
				listExpansionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listExpansionsOptionsModel).ToNot(BeNil())
				Expect(listExpansionsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listExpansionsOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(listExpansionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFieldsOptions successfully`, func() {
				// Construct an instance of the ListFieldsOptions model
				environmentID := "testString"
				collectionIds := []string{"testString"}
				listFieldsOptionsModel := discoveryService.NewListFieldsOptions(environmentID, collectionIds)
				listFieldsOptionsModel.SetEnvironmentID("testString")
				listFieldsOptionsModel.SetCollectionIds([]string{"testString"})
				listFieldsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFieldsOptionsModel).ToNot(BeNil())
				Expect(listFieldsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listFieldsOptionsModel.CollectionIds).To(Equal([]string{"testString"}))
				Expect(listFieldsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewaysOptions successfully`, func() {
				// Construct an instance of the ListGatewaysOptions model
				environmentID := "testString"
				listGatewaysOptionsModel := discoveryService.NewListGatewaysOptions(environmentID)
				listGatewaysOptionsModel.SetEnvironmentID("testString")
				listGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewaysOptionsModel).ToNot(BeNil())
				Expect(listGatewaysOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTrainingDataOptions successfully`, func() {
				// Construct an instance of the ListTrainingDataOptions model
				environmentID := "testString"
				collectionID := "testString"
				listTrainingDataOptionsModel := discoveryService.NewListTrainingDataOptions(environmentID, collectionID)
				listTrainingDataOptionsModel.SetEnvironmentID("testString")
				listTrainingDataOptionsModel.SetCollectionID("testString")
				listTrainingDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTrainingDataOptionsModel).ToNot(BeNil())
				Expect(listTrainingDataOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listTrainingDataOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(listTrainingDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTrainingExamplesOptions successfully`, func() {
				// Construct an instance of the ListTrainingExamplesOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				listTrainingExamplesOptionsModel := discoveryService.NewListTrainingExamplesOptions(environmentID, collectionID, queryID)
				listTrainingExamplesOptionsModel.SetEnvironmentID("testString")
				listTrainingExamplesOptionsModel.SetCollectionID("testString")
				listTrainingExamplesOptionsModel.SetQueryID("testString")
				listTrainingExamplesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTrainingExamplesOptionsModel).ToNot(BeNil())
				Expect(listTrainingExamplesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listTrainingExamplesOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(listTrainingExamplesOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(listTrainingExamplesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryLogOptions successfully`, func() {
				// Construct an instance of the QueryLogOptions model
				queryLogOptionsModel := discoveryService.NewQueryLogOptions()
				queryLogOptionsModel.SetFilter("testString")
				queryLogOptionsModel.SetQuery("testString")
				queryLogOptionsModel.SetCount(int64(38))
				queryLogOptionsModel.SetOffset(int64(38))
				queryLogOptionsModel.SetSort([]string{"testString"})
				queryLogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryLogOptionsModel).ToNot(BeNil())
				Expect(queryLogOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(queryLogOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(queryLogOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryLogOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryLogOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(queryLogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryNoticesOptions successfully`, func() {
				// Construct an instance of the QueryNoticesOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryNoticesOptionsModel := discoveryService.NewQueryNoticesOptions(environmentID, collectionID)
				queryNoticesOptionsModel.SetEnvironmentID("testString")
				queryNoticesOptionsModel.SetCollectionID("testString")
				queryNoticesOptionsModel.SetFilter("testString")
				queryNoticesOptionsModel.SetQuery("testString")
				queryNoticesOptionsModel.SetNaturalLanguageQuery("testString")
				queryNoticesOptionsModel.SetPassages(true)
				queryNoticesOptionsModel.SetAggregation("testString")
				queryNoticesOptionsModel.SetCount(int64(38))
				queryNoticesOptionsModel.SetReturn([]string{"testString"})
				queryNoticesOptionsModel.SetOffset(int64(38))
				queryNoticesOptionsModel.SetSort([]string{"testString"})
				queryNoticesOptionsModel.SetHighlight(true)
				queryNoticesOptionsModel.SetPassagesFields([]string{"testString"})
				queryNoticesOptionsModel.SetPassagesCount(int64(100))
				queryNoticesOptionsModel.SetPassagesCharacters(int64(50))
				queryNoticesOptionsModel.SetDeduplicateField("testString")
				queryNoticesOptionsModel.SetSimilar(true)
				queryNoticesOptionsModel.SetSimilarDocumentIds([]string{"testString"})
				queryNoticesOptionsModel.SetSimilarFields([]string{"testString"})
				queryNoticesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryNoticesOptionsModel).ToNot(BeNil())
				Expect(queryNoticesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Passages).To(Equal(core.BoolPtr(true)))
				Expect(queryNoticesOptionsModel.Aggregation).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryNoticesOptionsModel.Return).To(Equal([]string{"testString"}))
				Expect(queryNoticesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryNoticesOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(queryNoticesOptionsModel.Highlight).To(Equal(core.BoolPtr(true)))
				Expect(queryNoticesOptionsModel.PassagesFields).To(Equal([]string{"testString"}))
				Expect(queryNoticesOptionsModel.PassagesCount).To(Equal(core.Int64Ptr(int64(100))))
				Expect(queryNoticesOptionsModel.PassagesCharacters).To(Equal(core.Int64Ptr(int64(50))))
				Expect(queryNoticesOptionsModel.DeduplicateField).To(Equal(core.StringPtr("testString")))
				Expect(queryNoticesOptionsModel.Similar).To(Equal(core.BoolPtr(true)))
				Expect(queryNoticesOptionsModel.SimilarDocumentIds).To(Equal([]string{"testString"}))
				Expect(queryNoticesOptionsModel.SimilarFields).To(Equal([]string{"testString"}))
				Expect(queryNoticesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryOptions successfully`, func() {
				// Construct an instance of the QueryOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryOptionsModel := discoveryService.NewQueryOptions(environmentID, collectionID)
				queryOptionsModel.SetEnvironmentID("testString")
				queryOptionsModel.SetCollectionID("testString")
				queryOptionsModel.SetFilter("testString")
				queryOptionsModel.SetQuery("testString")
				queryOptionsModel.SetNaturalLanguageQuery("testString")
				queryOptionsModel.SetPassages(true)
				queryOptionsModel.SetAggregation("testString")
				queryOptionsModel.SetCount(int64(38))
				queryOptionsModel.SetReturn("testString")
				queryOptionsModel.SetOffset(int64(38))
				queryOptionsModel.SetSort("testString")
				queryOptionsModel.SetHighlight(true)
				queryOptionsModel.SetPassagesFields("testString")
				queryOptionsModel.SetPassagesCount(int64(100))
				queryOptionsModel.SetPassagesCharacters(int64(50))
				queryOptionsModel.SetDeduplicate(true)
				queryOptionsModel.SetDeduplicateField("testString")
				queryOptionsModel.SetSimilar(true)
				queryOptionsModel.SetSimilarDocumentIds("testString")
				queryOptionsModel.SetSimilarFields("testString")
				queryOptionsModel.SetBias("testString")
				queryOptionsModel.SetSpellingSuggestions(true)
				queryOptionsModel.SetXWatsonLoggingOptOut(true)
				queryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryOptionsModel).ToNot(BeNil())
				Expect(queryOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.NaturalLanguageQuery).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Passages).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.Aggregation).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Count).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryOptionsModel.Return).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(queryOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Highlight).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.PassagesFields).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.PassagesCount).To(Equal(core.Int64Ptr(int64(100))))
				Expect(queryOptionsModel.PassagesCharacters).To(Equal(core.Int64Ptr(int64(50))))
				Expect(queryOptionsModel.Deduplicate).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.DeduplicateField).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Similar).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.SimilarDocumentIds).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.SimilarFields).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.Bias).To(Equal(core.StringPtr("testString")))
				Expect(queryOptionsModel.SpellingSuggestions).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.XWatsonLoggingOptOut).To(Equal(core.BoolPtr(true)))
				Expect(queryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSourceOptionsBuckets successfully`, func() {
				name := "testString"
				model, err := discoveryService.NewSourceOptionsBuckets(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceOptionsFolder successfully`, func() {
				ownerUserID := "testString"
				folderID := "testString"
				model, err := discoveryService.NewSourceOptionsFolder(ownerUserID, folderID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceOptionsObject successfully`, func() {
				name := "testString"
				model, err := discoveryService.NewSourceOptionsObject(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceOptionsSiteColl successfully`, func() {
				siteCollectionPath := "testString"
				model, err := discoveryService.NewSourceOptionsSiteColl(siteCollectionPath)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSourceOptionsWebCrawl successfully`, func() {
				url := "testString"
				model, err := discoveryService.NewSourceOptionsWebCrawl(url)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTokenDictRule successfully`, func() {
				text := "testString"
				tokens := []string{"testString"}
				partOfSpeech := "testString"
				model, err := discoveryService.NewTokenDictRule(text, tokens, partOfSpeech)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCollectionOptions successfully`, func() {
				// Construct an instance of the UpdateCollectionOptions model
				environmentID := "testString"
				collectionID := "testString"
				updateCollectionOptionsName := "testString"
				updateCollectionOptionsModel := discoveryService.NewUpdateCollectionOptions(environmentID, collectionID, updateCollectionOptionsName)
				updateCollectionOptionsModel.SetEnvironmentID("testString")
				updateCollectionOptionsModel.SetCollectionID("testString")
				updateCollectionOptionsModel.SetName("testString")
				updateCollectionOptionsModel.SetDescription("testString")
				updateCollectionOptionsModel.SetConfigurationID("testString")
				updateCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectionOptionsModel).ToNot(BeNil())
				Expect(updateCollectionOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.ConfigurationID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigurationOptions successfully`, func() {
				// Construct an instance of the FontSetting model
				fontSettingModel := new(discoveryv1.FontSetting)
				Expect(fontSettingModel).ToNot(BeNil())
				fontSettingModel.Level = core.Int64Ptr(int64(38))
				fontSettingModel.MinSize = core.Int64Ptr(int64(38))
				fontSettingModel.MaxSize = core.Int64Ptr(int64(38))
				fontSettingModel.Bold = core.BoolPtr(true)
				fontSettingModel.Italic = core.BoolPtr(true)
				fontSettingModel.Name = core.StringPtr("testString")
				Expect(fontSettingModel.Level).To(Equal(core.Int64Ptr(int64(38))))
				Expect(fontSettingModel.MinSize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(fontSettingModel.MaxSize).To(Equal(core.Int64Ptr(int64(38))))
				Expect(fontSettingModel.Bold).To(Equal(core.BoolPtr(true)))
				Expect(fontSettingModel.Italic).To(Equal(core.BoolPtr(true)))
				Expect(fontSettingModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PDFHeadingDetection model
				pdfHeadingDetectionModel := new(discoveryv1.PDFHeadingDetection)
				Expect(pdfHeadingDetectionModel).ToNot(BeNil())
				pdfHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				Expect(pdfHeadingDetectionModel.Fonts).To(Equal([]discoveryv1.FontSetting{*fontSettingModel}))

				// Construct an instance of the PDFSettings model
				pdfSettingsModel := new(discoveryv1.PDFSettings)
				Expect(pdfSettingsModel).ToNot(BeNil())
				pdfSettingsModel.Heading = pdfHeadingDetectionModel
				Expect(pdfSettingsModel.Heading).To(Equal(pdfHeadingDetectionModel))

				// Construct an instance of the WordStyle model
				wordStyleModel := new(discoveryv1.WordStyle)
				Expect(wordStyleModel).ToNot(BeNil())
				wordStyleModel.Level = core.Int64Ptr(int64(38))
				wordStyleModel.Names = []string{"testString"}
				Expect(wordStyleModel.Level).To(Equal(core.Int64Ptr(int64(38))))
				Expect(wordStyleModel.Names).To(Equal([]string{"testString"}))

				// Construct an instance of the WordHeadingDetection model
				wordHeadingDetectionModel := new(discoveryv1.WordHeadingDetection)
				Expect(wordHeadingDetectionModel).ToNot(BeNil())
				wordHeadingDetectionModel.Fonts = []discoveryv1.FontSetting{*fontSettingModel}
				wordHeadingDetectionModel.Styles = []discoveryv1.WordStyle{*wordStyleModel}
				Expect(wordHeadingDetectionModel.Fonts).To(Equal([]discoveryv1.FontSetting{*fontSettingModel}))
				Expect(wordHeadingDetectionModel.Styles).To(Equal([]discoveryv1.WordStyle{*wordStyleModel}))

				// Construct an instance of the WordSettings model
				wordSettingsModel := new(discoveryv1.WordSettings)
				Expect(wordSettingsModel).ToNot(BeNil())
				wordSettingsModel.Heading = wordHeadingDetectionModel
				Expect(wordSettingsModel.Heading).To(Equal(wordHeadingDetectionModel))

				// Construct an instance of the XPathPatterns model
				xPathPatternsModel := new(discoveryv1.XPathPatterns)
				Expect(xPathPatternsModel).ToNot(BeNil())
				xPathPatternsModel.Xpaths = []string{"testString"}
				Expect(xPathPatternsModel.Xpaths).To(Equal([]string{"testString"}))

				// Construct an instance of the HTMLSettings model
				htmlSettingsModel := new(discoveryv1.HTMLSettings)
				Expect(htmlSettingsModel).ToNot(BeNil())
				htmlSettingsModel.ExcludeTagsCompletely = []string{"testString"}
				htmlSettingsModel.ExcludeTagsKeepContent = []string{"testString"}
				htmlSettingsModel.KeepContent = xPathPatternsModel
				htmlSettingsModel.ExcludeContent = xPathPatternsModel
				htmlSettingsModel.KeepTagAttributes = []string{"testString"}
				htmlSettingsModel.ExcludeTagAttributes = []string{"testString"}
				Expect(htmlSettingsModel.ExcludeTagsCompletely).To(Equal([]string{"testString"}))
				Expect(htmlSettingsModel.ExcludeTagsKeepContent).To(Equal([]string{"testString"}))
				Expect(htmlSettingsModel.KeepContent).To(Equal(xPathPatternsModel))
				Expect(htmlSettingsModel.ExcludeContent).To(Equal(xPathPatternsModel))
				Expect(htmlSettingsModel.KeepTagAttributes).To(Equal([]string{"testString"}))
				Expect(htmlSettingsModel.ExcludeTagAttributes).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentSettings model
				segmentSettingsModel := new(discoveryv1.SegmentSettings)
				Expect(segmentSettingsModel).ToNot(BeNil())
				segmentSettingsModel.Enabled = core.BoolPtr(true)
				segmentSettingsModel.SelectorTags = []string{"testString"}
				segmentSettingsModel.AnnotatedFields = []string{"testString"}
				Expect(segmentSettingsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(segmentSettingsModel.SelectorTags).To(Equal([]string{"testString"}))
				Expect(segmentSettingsModel.AnnotatedFields).To(Equal([]string{"testString"}))

				// Construct an instance of the NormalizationOperation model
				normalizationOperationModel := new(discoveryv1.NormalizationOperation)
				Expect(normalizationOperationModel).ToNot(BeNil())
				normalizationOperationModel.Operation = core.StringPtr("copy")
				normalizationOperationModel.SourceField = core.StringPtr("testString")
				normalizationOperationModel.DestinationField = core.StringPtr("testString")
				Expect(normalizationOperationModel.Operation).To(Equal(core.StringPtr("copy")))
				Expect(normalizationOperationModel.SourceField).To(Equal(core.StringPtr("testString")))
				Expect(normalizationOperationModel.DestinationField).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Conversions model
				conversionsModel := new(discoveryv1.Conversions)
				Expect(conversionsModel).ToNot(BeNil())
				conversionsModel.PDF = pdfSettingsModel
				conversionsModel.Word = wordSettingsModel
				conversionsModel.HTML = htmlSettingsModel
				conversionsModel.Segment = segmentSettingsModel
				conversionsModel.JSONNormalizations = []discoveryv1.NormalizationOperation{*normalizationOperationModel}
				conversionsModel.ImageTextRecognition = core.BoolPtr(true)
				Expect(conversionsModel.PDF).To(Equal(pdfSettingsModel))
				Expect(conversionsModel.Word).To(Equal(wordSettingsModel))
				Expect(conversionsModel.HTML).To(Equal(htmlSettingsModel))
				Expect(conversionsModel.Segment).To(Equal(segmentSettingsModel))
				Expect(conversionsModel.JSONNormalizations).To(Equal([]discoveryv1.NormalizationOperation{*normalizationOperationModel}))
				Expect(conversionsModel.ImageTextRecognition).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the NluEnrichmentKeywords model
				nluEnrichmentKeywordsModel := new(discoveryv1.NluEnrichmentKeywords)
				Expect(nluEnrichmentKeywordsModel).ToNot(BeNil())
				nluEnrichmentKeywordsModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Emotion = core.BoolPtr(true)
				nluEnrichmentKeywordsModel.Limit = core.Int64Ptr(int64(38))
				Expect(nluEnrichmentKeywordsModel.Sentiment).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentKeywordsModel.Emotion).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentKeywordsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NluEnrichmentEntities model
				nluEnrichmentEntitiesModel := new(discoveryv1.NluEnrichmentEntities)
				Expect(nluEnrichmentEntitiesModel).ToNot(BeNil())
				nluEnrichmentEntitiesModel.Sentiment = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Emotion = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Limit = core.Int64Ptr(int64(38))
				nluEnrichmentEntitiesModel.Mentions = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.MentionTypes = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.SentenceLocations = core.BoolPtr(true)
				nluEnrichmentEntitiesModel.Model = core.StringPtr("testString")
				Expect(nluEnrichmentEntitiesModel.Sentiment).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.Emotion).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(nluEnrichmentEntitiesModel.Mentions).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.MentionTypes).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.SentenceLocations).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEntitiesModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NluEnrichmentSentiment model
				nluEnrichmentSentimentModel := new(discoveryv1.NluEnrichmentSentiment)
				Expect(nluEnrichmentSentimentModel).ToNot(BeNil())
				nluEnrichmentSentimentModel.Document = core.BoolPtr(true)
				nluEnrichmentSentimentModel.Targets = []string{"testString"}
				Expect(nluEnrichmentSentimentModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentSentimentModel.Targets).To(Equal([]string{"testString"}))

				// Construct an instance of the NluEnrichmentEmotion model
				nluEnrichmentEmotionModel := new(discoveryv1.NluEnrichmentEmotion)
				Expect(nluEnrichmentEmotionModel).ToNot(BeNil())
				nluEnrichmentEmotionModel.Document = core.BoolPtr(true)
				nluEnrichmentEmotionModel.Targets = []string{"testString"}
				Expect(nluEnrichmentEmotionModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentEmotionModel.Targets).To(Equal([]string{"testString"}))

				// Construct an instance of the NluEnrichmentSemanticRoles model
				nluEnrichmentSemanticRolesModel := new(discoveryv1.NluEnrichmentSemanticRoles)
				Expect(nluEnrichmentSemanticRolesModel).ToNot(BeNil())
				nluEnrichmentSemanticRolesModel.Entities = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Keywords = core.BoolPtr(true)
				nluEnrichmentSemanticRolesModel.Limit = core.Int64Ptr(int64(38))
				Expect(nluEnrichmentSemanticRolesModel.Entities).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentSemanticRolesModel.Keywords).To(Equal(core.BoolPtr(true)))
				Expect(nluEnrichmentSemanticRolesModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NluEnrichmentRelations model
				nluEnrichmentRelationsModel := new(discoveryv1.NluEnrichmentRelations)
				Expect(nluEnrichmentRelationsModel).ToNot(BeNil())
				nluEnrichmentRelationsModel.Model = core.StringPtr("testString")
				Expect(nluEnrichmentRelationsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NluEnrichmentConcepts model
				nluEnrichmentConceptsModel := new(discoveryv1.NluEnrichmentConcepts)
				Expect(nluEnrichmentConceptsModel).ToNot(BeNil())
				nluEnrichmentConceptsModel.Limit = core.Int64Ptr(int64(38))
				Expect(nluEnrichmentConceptsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NluEnrichmentFeatures model
				nluEnrichmentFeaturesModel := new(discoveryv1.NluEnrichmentFeatures)
				Expect(nluEnrichmentFeaturesModel).ToNot(BeNil())
				nluEnrichmentFeaturesModel.Keywords = nluEnrichmentKeywordsModel
				nluEnrichmentFeaturesModel.Entities = nluEnrichmentEntitiesModel
				nluEnrichmentFeaturesModel.Sentiment = nluEnrichmentSentimentModel
				nluEnrichmentFeaturesModel.Emotion = nluEnrichmentEmotionModel
				nluEnrichmentFeaturesModel.Categories = make(map[string]interface{})
				nluEnrichmentFeaturesModel.SemanticRoles = nluEnrichmentSemanticRolesModel
				nluEnrichmentFeaturesModel.Relations = nluEnrichmentRelationsModel
				nluEnrichmentFeaturesModel.Concepts = nluEnrichmentConceptsModel
				Expect(nluEnrichmentFeaturesModel.Keywords).To(Equal(nluEnrichmentKeywordsModel))
				Expect(nluEnrichmentFeaturesModel.Entities).To(Equal(nluEnrichmentEntitiesModel))
				Expect(nluEnrichmentFeaturesModel.Sentiment).To(Equal(nluEnrichmentSentimentModel))
				Expect(nluEnrichmentFeaturesModel.Emotion).To(Equal(nluEnrichmentEmotionModel))
				Expect(nluEnrichmentFeaturesModel.Categories).To(Equal(make(map[string]interface{})))
				Expect(nluEnrichmentFeaturesModel.SemanticRoles).To(Equal(nluEnrichmentSemanticRolesModel))
				Expect(nluEnrichmentFeaturesModel.Relations).To(Equal(nluEnrichmentRelationsModel))
				Expect(nluEnrichmentFeaturesModel.Concepts).To(Equal(nluEnrichmentConceptsModel))

				// Construct an instance of the EnrichmentOptions model
				enrichmentOptionsModel := new(discoveryv1.EnrichmentOptions)
				Expect(enrichmentOptionsModel).ToNot(BeNil())
				enrichmentOptionsModel.Features = nluEnrichmentFeaturesModel
				enrichmentOptionsModel.Language = core.StringPtr("ar")
				enrichmentOptionsModel.Model = core.StringPtr("testString")
				Expect(enrichmentOptionsModel.Features).To(Equal(nluEnrichmentFeaturesModel))
				Expect(enrichmentOptionsModel.Language).To(Equal(core.StringPtr("ar")))
				Expect(enrichmentOptionsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Enrichment model
				enrichmentModel := new(discoveryv1.Enrichment)
				Expect(enrichmentModel).ToNot(BeNil())
				enrichmentModel.Description = core.StringPtr("testString")
				enrichmentModel.DestinationField = core.StringPtr("testString")
				enrichmentModel.SourceField = core.StringPtr("testString")
				enrichmentModel.Overwrite = core.BoolPtr(true)
				enrichmentModel.Enrichment = core.StringPtr("testString")
				enrichmentModel.IgnoreDownstreamErrors = core.BoolPtr(true)
				enrichmentModel.Options = enrichmentOptionsModel
				Expect(enrichmentModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.DestinationField).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.SourceField).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.Overwrite).To(Equal(core.BoolPtr(true)))
				Expect(enrichmentModel.Enrichment).To(Equal(core.StringPtr("testString")))
				Expect(enrichmentModel.IgnoreDownstreamErrors).To(Equal(core.BoolPtr(true)))
				Expect(enrichmentModel.Options).To(Equal(enrichmentOptionsModel))

				// Construct an instance of the SourceSchedule model
				sourceScheduleModel := new(discoveryv1.SourceSchedule)
				Expect(sourceScheduleModel).ToNot(BeNil())
				sourceScheduleModel.Enabled = core.BoolPtr(true)
				sourceScheduleModel.TimeZone = core.StringPtr("testString")
				sourceScheduleModel.Frequency = core.StringPtr("daily")
				Expect(sourceScheduleModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(sourceScheduleModel.TimeZone).To(Equal(core.StringPtr("testString")))
				Expect(sourceScheduleModel.Frequency).To(Equal(core.StringPtr("daily")))

				// Construct an instance of the SourceOptionsFolder model
				sourceOptionsFolderModel := new(discoveryv1.SourceOptionsFolder)
				Expect(sourceOptionsFolderModel).ToNot(BeNil())
				sourceOptionsFolderModel.OwnerUserID = core.StringPtr("testString")
				sourceOptionsFolderModel.FolderID = core.StringPtr("testString")
				sourceOptionsFolderModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsFolderModel.OwnerUserID).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsFolderModel.FolderID).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsFolderModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptionsObject model
				sourceOptionsObjectModel := new(discoveryv1.SourceOptionsObject)
				Expect(sourceOptionsObjectModel).ToNot(BeNil())
				sourceOptionsObjectModel.Name = core.StringPtr("testString")
				sourceOptionsObjectModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsObjectModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsObjectModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptionsSiteColl model
				sourceOptionsSiteCollModel := new(discoveryv1.SourceOptionsSiteColl)
				Expect(sourceOptionsSiteCollModel).ToNot(BeNil())
				sourceOptionsSiteCollModel.SiteCollectionPath = core.StringPtr("testString")
				sourceOptionsSiteCollModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsSiteCollModel.SiteCollectionPath).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsSiteCollModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptionsWebCrawl model
				sourceOptionsWebCrawlModel := new(discoveryv1.SourceOptionsWebCrawl)
				Expect(sourceOptionsWebCrawlModel).ToNot(BeNil())
				sourceOptionsWebCrawlModel.URL = core.StringPtr("testString")
				sourceOptionsWebCrawlModel.LimitToStartingHosts = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.CrawlSpeed = core.StringPtr("gentle")
				sourceOptionsWebCrawlModel.AllowUntrustedCertificate = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.MaximumHops = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.RequestTimeout = core.Int64Ptr(int64(38))
				sourceOptionsWebCrawlModel.OverrideRobotsTxt = core.BoolPtr(true)
				sourceOptionsWebCrawlModel.Blacklist = []string{"testString"}
				Expect(sourceOptionsWebCrawlModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsWebCrawlModel.LimitToStartingHosts).To(Equal(core.BoolPtr(true)))
				Expect(sourceOptionsWebCrawlModel.CrawlSpeed).To(Equal(core.StringPtr("gentle")))
				Expect(sourceOptionsWebCrawlModel.AllowUntrustedCertificate).To(Equal(core.BoolPtr(true)))
				Expect(sourceOptionsWebCrawlModel.MaximumHops).To(Equal(core.Int64Ptr(int64(38))))
				Expect(sourceOptionsWebCrawlModel.RequestTimeout).To(Equal(core.Int64Ptr(int64(38))))
				Expect(sourceOptionsWebCrawlModel.OverrideRobotsTxt).To(Equal(core.BoolPtr(true)))
				Expect(sourceOptionsWebCrawlModel.Blacklist).To(Equal([]string{"testString"}))

				// Construct an instance of the SourceOptionsBuckets model
				sourceOptionsBucketsModel := new(discoveryv1.SourceOptionsBuckets)
				Expect(sourceOptionsBucketsModel).ToNot(BeNil())
				sourceOptionsBucketsModel.Name = core.StringPtr("testString")
				sourceOptionsBucketsModel.Limit = core.Int64Ptr(int64(38))
				Expect(sourceOptionsBucketsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(sourceOptionsBucketsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the SourceOptions model
				sourceOptionsModel := new(discoveryv1.SourceOptions)
				Expect(sourceOptionsModel).ToNot(BeNil())
				sourceOptionsModel.Folders = []discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}
				sourceOptionsModel.Objects = []discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}
				sourceOptionsModel.SiteCollections = []discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}
				sourceOptionsModel.Urls = []discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}
				sourceOptionsModel.Buckets = []discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}
				sourceOptionsModel.CrawlAllBuckets = core.BoolPtr(true)
				Expect(sourceOptionsModel.Folders).To(Equal([]discoveryv1.SourceOptionsFolder{*sourceOptionsFolderModel}))
				Expect(sourceOptionsModel.Objects).To(Equal([]discoveryv1.SourceOptionsObject{*sourceOptionsObjectModel}))
				Expect(sourceOptionsModel.SiteCollections).To(Equal([]discoveryv1.SourceOptionsSiteColl{*sourceOptionsSiteCollModel}))
				Expect(sourceOptionsModel.Urls).To(Equal([]discoveryv1.SourceOptionsWebCrawl{*sourceOptionsWebCrawlModel}))
				Expect(sourceOptionsModel.Buckets).To(Equal([]discoveryv1.SourceOptionsBuckets{*sourceOptionsBucketsModel}))
				Expect(sourceOptionsModel.CrawlAllBuckets).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Source model
				sourceModel := new(discoveryv1.Source)
				Expect(sourceModel).ToNot(BeNil())
				sourceModel.Type = core.StringPtr("box")
				sourceModel.CredentialID = core.StringPtr("testString")
				sourceModel.Schedule = sourceScheduleModel
				sourceModel.Options = sourceOptionsModel
				Expect(sourceModel.Type).To(Equal(core.StringPtr("box")))
				Expect(sourceModel.CredentialID).To(Equal(core.StringPtr("testString")))
				Expect(sourceModel.Schedule).To(Equal(sourceScheduleModel))
				Expect(sourceModel.Options).To(Equal(sourceOptionsModel))

				// Construct an instance of the UpdateConfigurationOptions model
				environmentID := "testString"
				configurationID := "testString"
				updateConfigurationOptionsName := "testString"
				updateConfigurationOptionsModel := discoveryService.NewUpdateConfigurationOptions(environmentID, configurationID, updateConfigurationOptionsName)
				updateConfigurationOptionsModel.SetEnvironmentID("testString")
				updateConfigurationOptionsModel.SetConfigurationID("testString")
				updateConfigurationOptionsModel.SetName("testString")
				updateConfigurationOptionsModel.SetDescription("testString")
				updateConfigurationOptionsModel.SetConversions(conversionsModel)
				updateConfigurationOptionsModel.SetEnrichments([]discoveryv1.Enrichment{*enrichmentModel})
				updateConfigurationOptionsModel.SetNormalizations([]discoveryv1.NormalizationOperation{*normalizationOperationModel})
				updateConfigurationOptionsModel.SetSource(sourceModel)
				updateConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigurationOptionsModel).ToNot(BeNil())
				Expect(updateConfigurationOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigurationOptionsModel.ConfigurationID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigurationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigurationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigurationOptionsModel.Conversions).To(Equal(conversionsModel))
				Expect(updateConfigurationOptionsModel.Enrichments).To(Equal([]discoveryv1.Enrichment{*enrichmentModel}))
				Expect(updateConfigurationOptionsModel.Normalizations).To(Equal([]discoveryv1.NormalizationOperation{*normalizationOperationModel}))
				Expect(updateConfigurationOptionsModel.Source).To(Equal(sourceModel))
				Expect(updateConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCredentialsOptions successfully`, func() {
				// Construct an instance of the CredentialDetails model
				credentialDetailsModel := new(discoveryv1.CredentialDetails)
				Expect(credentialDetailsModel).ToNot(BeNil())
				credentialDetailsModel.CredentialType = core.StringPtr("oauth2")
				credentialDetailsModel.ClientID = core.StringPtr("testString")
				credentialDetailsModel.EnterpriseID = core.StringPtr("testString")
				credentialDetailsModel.URL = core.StringPtr("testString")
				credentialDetailsModel.Username = core.StringPtr("testString")
				credentialDetailsModel.OrganizationURL = core.StringPtr("testString")
				credentialDetailsModel.SiteCollectionPath = core.StringPtr("testString")
				credentialDetailsModel.ClientSecret = core.StringPtr("testString")
				credentialDetailsModel.PublicKeyID = core.StringPtr("testString")
				credentialDetailsModel.PrivateKey = core.StringPtr("testString")
				credentialDetailsModel.Passphrase = core.StringPtr("testString")
				credentialDetailsModel.Password = core.StringPtr("testString")
				credentialDetailsModel.GatewayID = core.StringPtr("testString")
				credentialDetailsModel.SourceVersion = core.StringPtr("online")
				credentialDetailsModel.WebApplicationURL = core.StringPtr("testString")
				credentialDetailsModel.Domain = core.StringPtr("testString")
				credentialDetailsModel.Endpoint = core.StringPtr("testString")
				credentialDetailsModel.AccessKeyID = core.StringPtr("testString")
				credentialDetailsModel.SecretAccessKey = core.StringPtr("testString")
				Expect(credentialDetailsModel.CredentialType).To(Equal(core.StringPtr("oauth2")))
				Expect(credentialDetailsModel.ClientID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.OrganizationURL).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.SiteCollectionPath).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.ClientSecret).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.PublicKeyID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.PrivateKey).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Passphrase).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.GatewayID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.SourceVersion).To(Equal(core.StringPtr("online")))
				Expect(credentialDetailsModel.WebApplicationURL).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.Endpoint).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.AccessKeyID).To(Equal(core.StringPtr("testString")))
				Expect(credentialDetailsModel.SecretAccessKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateCredentialsOptions model
				environmentID := "testString"
				credentialID := "testString"
				updateCredentialsOptionsModel := discoveryService.NewUpdateCredentialsOptions(environmentID, credentialID)
				updateCredentialsOptionsModel.SetEnvironmentID("testString")
				updateCredentialsOptionsModel.SetCredentialID("testString")
				updateCredentialsOptionsModel.SetSourceType("box")
				updateCredentialsOptionsModel.SetCredentialDetails(credentialDetailsModel)
				updateCredentialsOptionsModel.SetStatus("connected")
				updateCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCredentialsOptionsModel).ToNot(BeNil())
				Expect(updateCredentialsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateCredentialsOptionsModel.CredentialID).To(Equal(core.StringPtr("testString")))
				Expect(updateCredentialsOptionsModel.SourceType).To(Equal(core.StringPtr("box")))
				Expect(updateCredentialsOptionsModel.CredentialDetails).To(Equal(credentialDetailsModel))
				Expect(updateCredentialsOptionsModel.Status).To(Equal(core.StringPtr("connected")))
				Expect(updateCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDocumentOptions successfully`, func() {
				// Construct an instance of the UpdateDocumentOptions model
				environmentID := "testString"
				collectionID := "testString"
				documentID := "testString"
				updateDocumentOptionsModel := discoveryService.NewUpdateDocumentOptions(environmentID, collectionID, documentID)
				updateDocumentOptionsModel.SetEnvironmentID("testString")
				updateDocumentOptionsModel.SetCollectionID("testString")
				updateDocumentOptionsModel.SetDocumentID("testString")
				updateDocumentOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				updateDocumentOptionsModel.SetFilename("testString")
				updateDocumentOptionsModel.SetFileContentType("application/json")
				updateDocumentOptionsModel.SetMetadata("testString")
				updateDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDocumentOptionsModel).ToNot(BeNil())
				Expect(updateDocumentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDocumentOptionsModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.FileContentType).To(Equal(core.StringPtr("application/json")))
				Expect(updateDocumentOptionsModel.Metadata).To(Equal(core.StringPtr("testString")))
				Expect(updateDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnvironmentOptions successfully`, func() {
				// Construct an instance of the UpdateEnvironmentOptions model
				environmentID := "testString"
				updateEnvironmentOptionsModel := discoveryService.NewUpdateEnvironmentOptions(environmentID)
				updateEnvironmentOptionsModel.SetEnvironmentID("testString")
				updateEnvironmentOptionsModel.SetName("testString")
				updateEnvironmentOptionsModel.SetDescription("testString")
				updateEnvironmentOptionsModel.SetSize("S")
				updateEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnvironmentOptionsModel).ToNot(BeNil())
				Expect(updateEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnvironmentOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateEnvironmentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateEnvironmentOptionsModel.Size).To(Equal(core.StringPtr("S")))
				Expect(updateEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTrainingExampleOptions successfully`, func() {
				// Construct an instance of the UpdateTrainingExampleOptions model
				environmentID := "testString"
				collectionID := "testString"
				queryID := "testString"
				exampleID := "testString"
				updateTrainingExampleOptionsModel := discoveryService.NewUpdateTrainingExampleOptions(environmentID, collectionID, queryID, exampleID)
				updateTrainingExampleOptionsModel.SetEnvironmentID("testString")
				updateTrainingExampleOptionsModel.SetCollectionID("testString")
				updateTrainingExampleOptionsModel.SetQueryID("testString")
				updateTrainingExampleOptionsModel.SetExampleID("testString")
				updateTrainingExampleOptionsModel.SetCrossReference("testString")
				updateTrainingExampleOptionsModel.SetRelevance(int64(38))
				updateTrainingExampleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTrainingExampleOptionsModel).ToNot(BeNil())
				Expect(updateTrainingExampleOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingExampleOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingExampleOptionsModel.QueryID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingExampleOptionsModel.ExampleID).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingExampleOptionsModel.CrossReference).To(Equal(core.StringPtr("testString")))
				Expect(updateTrainingExampleOptionsModel.Relevance).To(Equal(core.Int64Ptr(int64(38))))
				Expect(updateTrainingExampleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
