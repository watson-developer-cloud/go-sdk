// +build integration

package comparecomplyv1_test

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

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/comparecomplyv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *comparecomplyv1.CompareComplyV1

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load(configFile)
	if err != nil {
		t.Skip(skipMessage)
	} else {
		configLoaded = true
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
		Version: "2018-10-15",
	})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestConvertToHTML(t *testing.T) {
	shouldSkipTest(t)

	testPDF, testPDFErr := os.Open("../resources/contract_A.pdf")
	if testPDFErr != nil {
		fmt.Println(testPDFErr)
	}

	html, _, responseErr := service.ConvertToHTML(
		&comparecomplyv1.ConvertToHTMLOptions{
			File:            testPDF,
			FileContentType: core.StringPtr("application/pdf"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, html)
}

func TestClassifyElements(t *testing.T) {
	shouldSkipTest(t)

	testPDF, testPDFErr := os.Open("../resources/contract_A.pdf")
	if testPDFErr != nil {
		fmt.Println(testPDFErr)
	}

	classifyElements, _, responseErr := service.ClassifyElements(
		&comparecomplyv1.ClassifyElementsOptions{
			File:            testPDF,
			FileContentType: core.StringPtr("application/pdf"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, classifyElements)
}

func TestExtractTables(t *testing.T) {
	shouldSkipTest(t)

	file, fileErr := os.Open("../resources/sample-tables.png")
	if fileErr != nil {
		fmt.Println(fileErr)
	}

	extractTables, _, responseErr := service.ExtractTables(
		&comparecomplyv1.ExtractTablesOptions{
			File:            file,
			FileContentType: core.StringPtr("image/png"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, extractTables)
}

func TestCompareDocuments(t *testing.T) {
	shouldSkipTest(t)

	file1, file1Err := os.Open("../resources/contract_A.pdf")
	if file1Err != nil {
		fmt.Println(file1Err)
	}

	file2, file2Err := os.Open("../resources/contract_B.pdf")
	if file2Err != nil {
		fmt.Println(file2Err)
	}

	compareDocuments, _, responseErr := service.CompareDocuments(
		&comparecomplyv1.CompareDocumentsOptions{
			File1:            file1,
			File2:            file2,
			File1ContentType: core.StringPtr("application/pdf"),
			File2ContentType: core.StringPtr("application/pdf"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, compareDocuments)
}

func TestFeedback(t *testing.T) {
	shouldSkipTest(t)

	// Add feedback
	// t.Skip()
	addFeedback, _, responseErr := service.AddFeedback(
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
	assert.NotNil(t, addFeedback)

	// List feedback
	listFeedback, _, responseErr := service.ListFeedback(
		&comparecomplyv1.ListFeedbackOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listFeedback)

	// Delete feedback
	deleteFeedback, _, responseErr := service.DeleteFeedback(
		&comparecomplyv1.DeleteFeedbackOptions{
			FeedbackID: addFeedback.FeedbackID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, deleteFeedback)
}

func TestBatch(t *testing.T) {
	shouldSkipTest(t)

	// Get batches
	listBatches, _, responseErr := service.ListBatches(
		&comparecomplyv1.ListBatchesOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listBatches)

	t.Skip()

	// Create batch
	inputCredentialsFile, inputCredentialsFileErr := os.Open("../resources/cloud-object-storage-credentials-input.json")
	if inputCredentialsFileErr != nil {
		fmt.Println(inputCredentialsFileErr)
	}

	outputCredentialsFile, outputCredentialsFileErr := os.Open("../resources/cloud-object-storage-credentials-output.json")
	if outputCredentialsFileErr != nil {
		fmt.Println(outputCredentialsFileErr)
	}

	createBatch, _, responseErr := service.CreateBatch(
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
	assert.NotNil(t, createBatch)

	// Get batch
	getBatch, _, responseErr := service.GetBatch(
		&comparecomplyv1.GetBatchOptions{
			BatchID: createBatch.BatchID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getBatch)

	// Update batch
	updateBatch, _, responseErr := service.UpdateBatch(
		&comparecomplyv1.UpdateBatchOptions{
			BatchID: createBatch.BatchID,
			Action:  core.StringPtr("rescan"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updateBatch)
}
