/**
 * (C) Copyright IBM Corp. 2020.
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

package comparecomplyv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/comparecomplyv1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe(`CompareComplyV1`, func() {
	Describe(`ConvertToHTML(convertToHTMLOptions *ConvertToHTMLOptions)`, func() {
		convertToHTMLPath := "/v1/html_conversion"
		version := "exampleString"
		bearerToken := "0ui9876453"
		file := new(os.File)
		Context(`Successfully - Convert document to HTML`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(convertToHTMLPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ConvertToHTML`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ConvertToHTML(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				convertToHTMLOptions := testService.NewConvertToHTMLOptions(file)
				result, response, operationErr = testService.ConvertToHTML(convertToHTMLOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ClassifyElements(classifyElementsOptions *ClassifyElementsOptions)`, func() {
		classifyElementsPath := "/v1/element_classification"
		version := "exampleString"
		bearerToken := "0ui9876453"
		file := new(os.File)
		Context(`Successfully - Classify the elements of a document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(classifyElementsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ClassifyElements`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ClassifyElements(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				classifyElementsOptions := testService.NewClassifyElementsOptions(file)
				result, response, operationErr = testService.ClassifyElements(classifyElementsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ExtractTables(extractTablesOptions *ExtractTablesOptions)`, func() {
		extractTablesPath := "/v1/tables"
		version := "exampleString"
		bearerToken := "0ui9876453"
		file := new(os.File)
		Context(`Successfully - Extract a document's tables`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(extractTablesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ExtractTables`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ExtractTables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				extractTablesOptions := testService.NewExtractTablesOptions(file)
				result, response, operationErr = testService.ExtractTables(extractTablesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CompareDocuments(compareDocumentsOptions *CompareDocumentsOptions)`, func() {
		compareDocumentsPath := "/v1/comparison"
		version := "exampleString"
		bearerToken := "0ui9876453"
		file1 := new(os.File)
		file2 := new(os.File)
		Context(`Successfully - Compare two documents`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(compareDocumentsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CompareDocuments`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CompareDocuments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				compareDocumentsOptions := testService.NewCompareDocumentsOptions(file1, file2)
				result, response, operationErr = testService.CompareDocuments(compareDocumentsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddFeedback(addFeedbackOptions *AddFeedbackOptions)`, func() {
		addFeedbackPath := "/v1/feedback"
		version := "exampleString"
		bearerToken := "0ui9876453"
		feedbackData := &comparecomplyv1.FeedbackDataInput{FeedbackType: core.StringPtr("exampleString"), Location: &comparecomplyv1.Location{Begin: core.Int64Ptr(int64(1234)), End: core.Int64Ptr(int64(1234))}, Text: core.StringPtr("exampleString"), OriginalLabels: &comparecomplyv1.OriginalLabelsIn{Types: []comparecomplyv1.TypeLabel{}, Categories: []comparecomplyv1.Category{}}, UpdatedLabels: &comparecomplyv1.UpdatedLabelsIn{Types: []comparecomplyv1.TypeLabel{}, Categories: []comparecomplyv1.Category{}}}
		Context(`Successfully - Add feedback`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addFeedbackPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call AddFeedback`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.AddFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				addFeedbackOptions := testService.NewAddFeedbackOptions(feedbackData)
				result, response, operationErr = testService.AddFeedback(addFeedbackOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListFeedback(listFeedbackOptions *ListFeedbackOptions)`, func() {
		listFeedbackPath := "/v1/feedback"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List the feedback in a document`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listFeedbackPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListFeedback`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listFeedbackOptions := testService.NewListFeedbackOptions()
				result, response, operationErr = testService.ListFeedback(listFeedbackOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetFeedback(getFeedbackOptions *GetFeedbackOptions)`, func() {
		getFeedbackPath := "/v1/feedback/{feedback_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		feedbackID := "exampleString"
		getFeedbackPath = strings.Replace(getFeedbackPath, "{feedback_id}", feedbackID, 1)
		Context(`Successfully - Get a specified feedback entry`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getFeedbackPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetFeedback`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getFeedbackOptions := testService.NewGetFeedbackOptions(feedbackID)
				result, response, operationErr = testService.GetFeedback(getFeedbackOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteFeedback(deleteFeedbackOptions *DeleteFeedbackOptions)`, func() {
		deleteFeedbackPath := "/v1/feedback/{feedback_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		feedbackID := "exampleString"
		deleteFeedbackPath = strings.Replace(deleteFeedbackPath, "{feedback_id}", feedbackID, 1)
		Context(`Successfully - Delete a specified feedback entry`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteFeedbackPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call DeleteFeedback`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteFeedback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteFeedbackOptions := testService.NewDeleteFeedbackOptions(feedbackID)
				result, response, operationErr = testService.DeleteFeedback(deleteFeedbackOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateBatch(createBatchOptions *CreateBatchOptions)`, func() {
		createBatchPath := "/v1/batches"
		version := "exampleString"
		bearerToken := "0ui9876453"
		function := "exampleString"
		inputCredentialsFile := new(os.File)
		inputBucketLocation := "exampleString"
		inputBucketName := "exampleString"
		outputCredentialsFile := new(os.File)
		outputBucketLocation := "exampleString"
		outputBucketName := "exampleString"
		Context(`Successfully - Submit a batch-processing request`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createBatchPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["function"]).To(Equal([]string{function}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call CreateBatch`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateBatch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createBatchOptions := testService.NewCreateBatchOptions(function, inputCredentialsFile, inputBucketLocation, inputBucketName, outputCredentialsFile, outputBucketLocation, outputBucketName)
				result, response, operationErr = testService.CreateBatch(createBatchOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListBatches(listBatchesOptions *ListBatchesOptions)`, func() {
		listBatchesPath := "/v1/batches"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List submitted batch-processing jobs`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listBatchesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call ListBatches`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListBatches(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listBatchesOptions := testService.NewListBatchesOptions()
				result, response, operationErr = testService.ListBatches(listBatchesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetBatch(getBatchOptions *GetBatchOptions)`, func() {
		getBatchPath := "/v1/batches/{batch_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		batchID := "exampleString"
		getBatchPath = strings.Replace(getBatchPath, "{batch_id}", batchID, 1)
		Context(`Successfully - Get information about a specific batch-processing job`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getBatchPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetBatch`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetBatch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getBatchOptions := testService.NewGetBatchOptions(batchID)
				result, response, operationErr = testService.GetBatch(getBatchOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateBatch(updateBatchOptions *UpdateBatchOptions)`, func() {
		updateBatchPath := "/v1/batches/{batch_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		batchID := "exampleString"
		action := "exampleString"
		updateBatchPath = strings.Replace(updateBatchPath, "{batch_id}", batchID, 1)
		Context(`Successfully - Update a pending or active batch-processing job`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateBatchPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["action"]).To(Equal([]string{action}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call UpdateBatch`, func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateBatch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateBatchOptions := testService.NewUpdateBatchOptions(batchID, action)
				result, response, operationErr = testService.UpdateBatch(updateBatchOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				URL:           "http://comparecomplyv1modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewFeedbackDataInput successfully", func() {
				feedbackType := "exampleString"
				location := &comparecomplyv1.Location{Begin: core.Int64Ptr(int64(1234)), End: core.Int64Ptr(int64(1234))}
				text := "exampleString"
				originalLabels := &comparecomplyv1.OriginalLabelsIn{Types: []comparecomplyv1.TypeLabel{}, Categories: []comparecomplyv1.Category{}}
				updatedLabels := &comparecomplyv1.UpdatedLabelsIn{Types: []comparecomplyv1.TypeLabel{}, Categories: []comparecomplyv1.Category{}}
				model, err := testService.NewFeedbackDataInput(feedbackType, location, text, originalLabels, updatedLabels)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewLabel successfully", func() {
				nature := "exampleString"
				party := "exampleString"
				model, err := testService.NewLabel(nature, party)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewLocation successfully", func() {
				begin := int64(1234)
				end := int64(1234)
				model, err := testService.NewLocation(begin, end)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewOriginalLabelsIn successfully", func() {
				types := []comparecomplyv1.TypeLabel{}
				categories := []comparecomplyv1.Category{}
				model, err := testService.NewOriginalLabelsIn(types, categories)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewUpdatedLabelsIn successfully", func() {
				types := []comparecomplyv1.TypeLabel{}
				categories := []comparecomplyv1.Category{}
				model, err := testService.NewUpdatedLabelsIn(types, categories)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
