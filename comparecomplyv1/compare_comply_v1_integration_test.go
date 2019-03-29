// +build integration

package comparecomplyv1_test

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
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/comparecomplyv1"
	"os"
	"testing"
)

var service *comparecomplyv1.CompareComplyV1
var serviceErr error

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = comparecomplyv1.
			NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
				URL:       os.Getenv("COMPARE_COMPLY_URL"),
				Version:   "2018-10-15",
				IAMApiKey: os.Getenv("COMPARE_COMPLY_IAMAPIKEY"),
			})
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestConvertToHTML(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()
	testPDF, testPDFErr := os.Open(pwd + "/../resources/contract_A.pdf")
	if testPDFErr != nil {
		fmt.Println(testPDFErr)
	}

	response, responseErr := service.ConvertToHTML(
		&comparecomplyv1.ConvertToHTMLOptions{
			File:     testPDF,
			Filename: core.StringPtr("contract_A.pdf"),
		},
	)
	assert.Nil(t, responseErr)

	html := service.GetConvertToHTMLResult(response)
	assert.NotNil(t, html)
}

func TestClassifyElements(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()
	testPDF, testPDFErr := os.Open(pwd + "/../resources/contract_A.pdf")
	if testPDFErr != nil {
		fmt.Println(testPDFErr)
	}

	response, responseErr := service.ClassifyElements(
		&comparecomplyv1.ClassifyElementsOptions{
			File:            testPDF,
			FileContentType: core.StringPtr("application/pdf"),
		},
	)
	assert.Nil(t, responseErr)

	classifyElements := service.GetClassifyElementsResult(response)
	assert.NotNil(t, classifyElements)
}

func TestExtractTables(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/sample-tables.pdf")
	if fileErr != nil {
		fmt.Println(fileErr)
	}

	response, responseErr := service.ExtractTables(
		&comparecomplyv1.ExtractTablesOptions{
			File:            file,
			FileContentType: core.StringPtr("application/pdf"),
		},
	)
	assert.Nil(t, responseErr)

	extractTables := service.GetExtractTablesResult(response)
	assert.NotNil(t, extractTables)
}

func TestCompareDocuments(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()
	file1, file1Err := os.Open(pwd + "/../resources/contract_A.pdf")
	if file1Err != nil {
		fmt.Println(file1Err)
	}

	file2, file2Err := os.Open(pwd + "/../resources/contract_B.pdf")
	if file2Err != nil {
		fmt.Println(file2Err)
	}

	response, responseErr := service.CompareDocuments(
		&comparecomplyv1.CompareDocumentsOptions{
			File1:            file1,
			File2:            file2,
			File1ContentType: core.StringPtr("application/pdf"),
			File2ContentType: core.StringPtr("application/pdf"),
		},
	)
	assert.Nil(t, responseErr)

	compareDocuments := service.GetCompareDocumentsResult(response)
	assert.NotNil(t, compareDocuments)
}

func TestFeedback(t *testing.T) {
	shouldSkipTest(t)

	// Add feedback
	t.Skip()
	response, responseErr := service.AddFeedback(
		&comparecomplyv1.AddFeedbackOptions{
			UserID:  core.StringPtr("wonder woman"),
			Comment: core.StringPtr("test comment"),
			FeedbackData: &comparecomplyv1.FeedbackDataInput{
				FeedbackType: core.StringPtr("element_classification"),
				Document: &comparecomplyv1.ShortDoc{
					Title: core.StringPtr("title"),
				},
				ModelID:      core.StringPtr("contracts"),
				ModelVersion: core.StringPtr("11.00"),
				Location: &comparecomplyv1.Location{
					Begin: core.Int64Ptr(214),
					End:   core.Int64Ptr(237),
				},
				Text: core.StringPtr("1. IBM will provide a Senior Managing Consultant / expert resource, for up to 80 hours, to assist Florida Power & Light (FPL) with the creation of an IT infrastructure unit cost model for existing infrastructure."),
				OriginalLabels: &comparecomplyv1.OriginalLabelsIn{
					Types: []comparecomplyv1.TypeLabel{
						comparecomplyv1.TypeLabel{
							Label: &comparecomplyv1.Label{
								Nature: core.StringPtr("Obligation"),
								Party:  core.StringPtr("IBM"),
							},
							ProvenanceIds: []string{"85f5981a-ba91-44f5-9efa-0bd22e64b7bc", "ce0480a1-5ef1-4c3e-9861-3743b5610795"},
						},
					},
					Categories: []comparecomplyv1.Category{
						comparecomplyv1.Category{
							Label: core.StringPtr(comparecomplyv1.Category_Label_Amendments),
						},
					},
				},
				UpdatedLabels: &comparecomplyv1.UpdatedLabelsIn{
					Types: []comparecomplyv1.TypeLabel{
						comparecomplyv1.TypeLabel{
							Label: &comparecomplyv1.Label{
								Nature: core.StringPtr("Obligation"),
								Party:  core.StringPtr("IBM"),
							},
						},
						comparecomplyv1.TypeLabel{
							Label: &comparecomplyv1.Label{
								Nature: core.StringPtr("Disclaimer"),
								Party:  core.StringPtr("Buyer"),
							},
						},
					},
					Categories: []comparecomplyv1.Category{
						comparecomplyv1.Category{
							Label: core.StringPtr("Responsibilities"),
						},
					},
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	addFeedback := service.GetAddFeedbackResult(response)
	assert.NotNil(t, addFeedback)

	// Get feedback
	// response, responseErr = service.GetFeedback(
	// 	&comparecomplyv1.GetFeedbackOptions{
	// 		FeedbackID: addFeedback.FeedbackID,
	// 	},
	// )
	// assert.Nil(t, responseErr)

	// getFeedback := service.GetGetFeedbackResult(response)
	// assert.NotNil(t, getFeedback)

	// List feedback
	response, responseErr = service.ListFeedback(
		&comparecomplyv1.ListFeedbackOptions{},
	)
	assert.Nil(t, responseErr)

	listFeedback := service.GetListFeedbackResult(response)
	assert.NotNil(t, listFeedback)

	// Delete feedback
	response, responseErr = service.DeleteFeedback(
		&comparecomplyv1.DeleteFeedbackOptions{
			FeedbackID: addFeedback.FeedbackID,
		},
	)
	assert.Nil(t, responseErr)

	deleteFeedback := service.GetDeleteFeedbackResult(response)
	assert.NotNil(t, deleteFeedback)
}

func TestBatch(t *testing.T) {
	shouldSkipTest(t)

	// Get batches
	response, responseErr := service.ListBatches(
		&comparecomplyv1.ListBatchesOptions{},
	)
	assert.Nil(t, responseErr)

	listBatches := service.GetListBatchesResult(response)
	assert.NotNil(t, listBatches)

	t.Skip()
	// Create batch
	pwd, _ := os.Getwd()
	inputCredentialsFile, inputCredentialsFileErr := os.Open(pwd + "/../resources/cloud-object-storage-credentials-input.json")
	if inputCredentialsFileErr != nil {
		fmt.Println(inputCredentialsFileErr)
	}

	outputCredentialsFile, outputCredentialsFileErr := os.Open(pwd + "/../resources/cloud-object-storage-credentials-output.json")
	if outputCredentialsFileErr != nil {
		fmt.Println(outputCredentialsFileErr)
	}

	response, responseErr = service.CreateBatch(
		&comparecomplyv1.CreateBatchOptions{
			Function:              core.StringPtr("html_conversion"),
			InputCredentialsFile:  inputCredentialsFile,
			InputBucketLocation:   core.StringPtr("us-south"),
			InputBucketName:       core.StringPtr("compare-comply-integration-test-bucket-input"),
			OutputCredentialsFile: outputCredentialsFile,
			OutputBucketLocation:  core.StringPtr("us-south"),
			OutputBucketName:      core.StringPtr("compare-comply-integration-test-bucket-output"),
		},
	)
	assert.Nil(t, responseErr)

	createBatch := service.GetCreateBatchResult(response)
	assert.NotNil(t, createBatch)

	// Get batch
	response, responseErr = service.GetBatch(
		&comparecomplyv1.GetBatchOptions{
			BatchID: createBatch.BatchID,
		},
	)
	assert.Nil(t, responseErr)

	getBatch := service.GetGetBatchResult(response)
	assert.NotNil(t, getBatch)

	// Update batch
	response, responseErr = service.UpdateBatch(
		&comparecomplyv1.UpdateBatchOptions{
			BatchID: createBatch.BatchID,
			Action:  core.StringPtr("rescan"),
		},
	)
	assert.Nil(t, responseErr)

	updateBatch := service.GetUpdateBatchResult(response)
	assert.NotNil(t, updateBatch)
}
