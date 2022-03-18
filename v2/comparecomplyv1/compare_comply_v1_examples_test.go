//go:build examples
// +build examples

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

package comparecomplyv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/comparecomplyv1"
)

//
// This file provides an example of how to use the Compare Comply service.
//
// The following configuration properties are assumed to be defined:
// COMPARE-COMPLY_URL=<service base url>
// COMPARE-COMPLY_AUTH_TYPE=iam
// COMPARE-COMPLY_APIKEY=<IAM apikey>
// COMPARE-COMPLY_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../compare_comply_v1.env"

var (
	compareComplyService *comparecomplyv1.CompareComplyV1
	config               map[string]string
	configLoaded         bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`CompareComplyV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(comparecomplyv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			compareComplyServiceOptions := &comparecomplyv1.CompareComplyV1Options{
				Version: core.StringPtr("testString"),
			}

			compareComplyService, err = comparecomplyv1.NewCompareComplyV1(compareComplyServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(compareComplyService).ToNot(BeNil())
		})
	})

	Describe(`CompareComplyV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ConvertToHTML request example`, func() {
			fmt.Println("\nConvertToHTML() result:")
			// begin-convertToHtml

			convertToHTMLOptions := compareComplyService.NewConvertToHTMLOptions(
				CreateMockReader("This is a mock file."),
			)

			htmlReturn, response, err := compareComplyService.ConvertToHTML(convertToHTMLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(htmlReturn, "", "  ")
			fmt.Println(string(b))

			// end-convertToHtml

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(htmlReturn).ToNot(BeNil())

		})
		It(`ClassifyElements request example`, func() {
			fmt.Println("\nClassifyElements() result:")
			// begin-classifyElements

			classifyElementsOptions := compareComplyService.NewClassifyElementsOptions(
				CreateMockReader("This is a mock file."),
			)

			classifyReturn, response, err := compareComplyService.ClassifyElements(classifyElementsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifyReturn, "", "  ")
			fmt.Println(string(b))

			// end-classifyElements

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifyReturn).ToNot(BeNil())

		})
		It(`ExtractTables request example`, func() {
			fmt.Println("\nExtractTables() result:")
			// begin-extractTables

			extractTablesOptions := compareComplyService.NewExtractTablesOptions(
				CreateMockReader("This is a mock file."),
			)

			tableReturn, response, err := compareComplyService.ExtractTables(extractTablesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tableReturn, "", "  ")
			fmt.Println(string(b))

			// end-extractTables

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tableReturn).ToNot(BeNil())

		})
		It(`CompareDocuments request example`, func() {
			fmt.Println("\nCompareDocuments() result:")
			// begin-compareDocuments

			compareDocumentsOptions := compareComplyService.NewCompareDocumentsOptions(
				CreateMockReader("This is a mock file."),
				CreateMockReader("This is a mock file."),
			)

			compareReturn, response, err := compareComplyService.CompareDocuments(compareDocumentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(compareReturn, "", "  ")
			fmt.Println(string(b))

			// end-compareDocuments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(compareReturn).ToNot(BeNil())

		})
		It(`AddFeedback request example`, func() {
			fmt.Println("\nAddFeedback() result:")
			// begin-addFeedback

			locationModel := &comparecomplyv1.Location{
				Begin: core.Int64Ptr(int64(26)),
				End:   core.Int64Ptr(int64(26)),
			}

			typeLabelModel := &comparecomplyv1.TypeLabel{}

			categoryModel := &comparecomplyv1.Category{}

			originalLabelsInModel := &comparecomplyv1.OriginalLabelsIn{
				Types:      []comparecomplyv1.TypeLabel{*typeLabelModel},
				Categories: []comparecomplyv1.Category{*categoryModel},
			}

			updatedLabelsInModel := &comparecomplyv1.UpdatedLabelsIn{
				Types:      []comparecomplyv1.TypeLabel{*typeLabelModel},
				Categories: []comparecomplyv1.Category{*categoryModel},
			}

			feedbackDataInputModel := &comparecomplyv1.FeedbackDataInput{
				FeedbackType:   core.StringPtr("testString"),
				Location:       locationModel,
				Text:           core.StringPtr("testString"),
				OriginalLabels: originalLabelsInModel,
				UpdatedLabels:  updatedLabelsInModel,
			}

			addFeedbackOptions := compareComplyService.NewAddFeedbackOptions(
				feedbackDataInputModel,
			)

			feedbackReturn, response, err := compareComplyService.AddFeedback(addFeedbackOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feedbackReturn, "", "  ")
			fmt.Println(string(b))

			// end-addFeedback

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feedbackReturn).ToNot(BeNil())

		})
		It(`ListFeedback request example`, func() {
			fmt.Println("\nListFeedback() result:")
			// begin-listFeedback

			listFeedbackOptions := compareComplyService.NewListFeedbackOptions()

			feedbackList, response, err := compareComplyService.ListFeedback(listFeedbackOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feedbackList, "", "  ")
			fmt.Println(string(b))

			// end-listFeedback

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feedbackList).ToNot(BeNil())

		})
		It(`GetFeedback request example`, func() {
			fmt.Println("\nGetFeedback() result:")
			// begin-getFeedback

			getFeedbackOptions := compareComplyService.NewGetFeedbackOptions(
				"testString",
			)

			getFeedback, response, err := compareComplyService.GetFeedback(getFeedbackOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getFeedback, "", "  ")
			fmt.Println(string(b))

			// end-getFeedback

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getFeedback).ToNot(BeNil())

		})
		It(`CreateBatch request example`, func() {
			fmt.Println("\nCreateBatch() result:")
			// begin-createBatch

			createBatchOptions := compareComplyService.NewCreateBatchOptions(
				"html_conversion",
				CreateMockReader("This is a mock file."),
				"testString",
				"testString",
				CreateMockReader("This is a mock file."),
				"testString",
				"testString",
			)

			batchStatus, response, err := compareComplyService.CreateBatch(createBatchOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(batchStatus, "", "  ")
			fmt.Println(string(b))

			// end-createBatch

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(batchStatus).ToNot(BeNil())

		})
		It(`ListBatches request example`, func() {
			fmt.Println("\nListBatches() result:")
			// begin-listBatches

			listBatchesOptions := compareComplyService.NewListBatchesOptions()

			batches, response, err := compareComplyService.ListBatches(listBatchesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(batches, "", "  ")
			fmt.Println(string(b))

			// end-listBatches

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(batches).ToNot(BeNil())

		})
		It(`GetBatch request example`, func() {
			fmt.Println("\nGetBatch() result:")
			// begin-getBatch

			getBatchOptions := compareComplyService.NewGetBatchOptions(
				"testString",
			)

			batchStatus, response, err := compareComplyService.GetBatch(getBatchOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(batchStatus, "", "  ")
			fmt.Println(string(b))

			// end-getBatch

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(batchStatus).ToNot(BeNil())

		})
		It(`UpdateBatch request example`, func() {
			fmt.Println("\nUpdateBatch() result:")
			// begin-updateBatch

			updateBatchOptions := compareComplyService.NewUpdateBatchOptions(
				"testString",
				"rescan",
			)

			batchStatus, response, err := compareComplyService.UpdateBatch(updateBatchOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(batchStatus, "", "  ")
			fmt.Println(string(b))

			// end-updateBatch

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(batchStatus).ToNot(BeNil())

		})
		It(`DeleteFeedback request example`, func() {
			fmt.Println("\nDeleteFeedback() result:")
			// begin-deleteFeedback

			deleteFeedbackOptions := compareComplyService.NewDeleteFeedbackOptions(
				"testString",
			)

			feedbackDeleted, response, err := compareComplyService.DeleteFeedback(deleteFeedbackOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feedbackDeleted, "", "  ")
			fmt.Println(string(b))

			// end-deleteFeedback

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feedbackDeleted).ToNot(BeNil())

		})
	})
})
