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

package naturallanguageunderstandingv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/naturallanguageunderstandingv1"
)

//
// This file provides an example of how to use the Natural Language Understanding service.
//
// The following configuration properties are assumed to be defined:
// NATURAL-LANGUAGE-UNDERSTANDING_URL=<service base url>
// NATURAL-LANGUAGE-UNDERSTANDING_AUTH_TYPE=iam
// NATURAL-LANGUAGE-UNDERSTANDING_APIKEY=<IAM apikey>
// NATURAL-LANGUAGE-UNDERSTANDING_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../natural_language_understanding_v1.env"

var (
	naturalLanguageUnderstandingService *naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1
	config                              map[string]string
	configLoaded                        bool = false
)

// Globlal variables to hold link values
var (
	deleteCategoriesModelLink      string
	deleteClassificationsModelLink string
	deleteSentimentModelLink       string
	getCategoriesModelLink         string
	getClassificationsModelLink    string
	getSentimentModelLink          string
	updateCategoriesModelLink      string
	updateClassificationsModelLink string
	updateSentimentModelLink       string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`NaturalLanguageUnderstandingV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(naturallanguageunderstandingv1.DefaultServiceName)
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

			naturalLanguageUnderstandingServiceOptions := &naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				Version: core.StringPtr("testString"),
			}

			naturalLanguageUnderstandingService, err = naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(naturalLanguageUnderstandingServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
		})
	})

	Describe(`NaturalLanguageUnderstandingV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCategoriesModel request example`, func() {
			fmt.Println("\nCreateCategoriesModel() result:")
			// begin-createCategoriesModel

			createCategoriesModelOptions := naturalLanguageUnderstandingService.NewCreateCategoriesModelOptions(
				"testString",
				CreateMockReader("This is a mock file."),
			)

			categoriesModel, response, err := naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(categoriesModel, "", "  ")
			fmt.Println(string(b))

			// end-createCategoriesModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(categoriesModel).ToNot(BeNil())

			getCategoriesModelLink = *categoriesModel.ModelID
			updateCategoriesModelLink = *categoriesModel.ModelID
			deleteCategoriesModelLink = *categoriesModel.ModelID

		})
		It(`CreateClassificationsModel request example`, func() {
			fmt.Println("\nCreateClassificationsModel() result:")
			// begin-createClassificationsModel

			createClassificationsModelOptions := naturalLanguageUnderstandingService.NewCreateClassificationsModelOptions(
				"testString",
				CreateMockReader("This is a mock file."),
			)

			classificationsModel, response, err := naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classificationsModel, "", "  ")
			fmt.Println(string(b))

			// end-createClassificationsModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(classificationsModel).ToNot(BeNil())

			getClassificationsModelLink = *classificationsModel.ModelID
			updateClassificationsModelLink = *classificationsModel.ModelID
			deleteClassificationsModelLink = *classificationsModel.ModelID

		})
		It(`CreateSentimentModel request example`, func() {
			fmt.Println("\nCreateSentimentModel() result:")
			// begin-createSentimentModel

			createSentimentModelOptions := naturalLanguageUnderstandingService.NewCreateSentimentModelOptions(
				"testString",
				CreateMockReader("This is a mock file."),
			)

			sentimentModel, response, err := naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sentimentModel, "", "  ")
			fmt.Println(string(b))

			// end-createSentimentModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sentimentModel).ToNot(BeNil())

			getSentimentModelLink = *sentimentModel.ModelID
			updateSentimentModelLink = *sentimentModel.ModelID
			deleteSentimentModelLink = *sentimentModel.ModelID

		})
		It(`Analyze request example`, func() {
			fmt.Println("\nAnalyze() result:")
			// begin-analyze

			featuresModel := &naturallanguageunderstandingv1.Features{}

			analyzeOptions := naturalLanguageUnderstandingService.NewAnalyzeOptions(
				featuresModel,
			)

			analysisResults, response, err := naturalLanguageUnderstandingService.Analyze(analyzeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analysisResults, "", "  ")
			fmt.Println(string(b))

			// end-analyze

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analysisResults).ToNot(BeNil())

		})
		It(`ListModels request example`, func() {
			fmt.Println("\nListModels() result:")
			// begin-listModels

			listModelsOptions := naturalLanguageUnderstandingService.NewListModelsOptions()

			listModelsResults, response, err := naturalLanguageUnderstandingService.ListModels(listModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listModelsResults, "", "  ")
			fmt.Println(string(b))

			// end-listModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listModelsResults).ToNot(BeNil())

		})
		It(`ListSentimentModels request example`, func() {
			fmt.Println("\nListSentimentModels() result:")
			// begin-listSentimentModels

			listSentimentModelsOptions := naturalLanguageUnderstandingService.NewListSentimentModelsOptions()

			listSentimentModelsResponse, response, err := naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listSentimentModelsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listSentimentModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listSentimentModelsResponse).ToNot(BeNil())

		})
		It(`GetSentimentModel request example`, func() {
			fmt.Println("\nGetSentimentModel() result:")
			// begin-getSentimentModel

			getSentimentModelOptions := naturalLanguageUnderstandingService.NewGetSentimentModelOptions(
				getSentimentModelLink,
			)

			sentimentModel, response, err := naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sentimentModel, "", "  ")
			fmt.Println(string(b))

			// end-getSentimentModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sentimentModel).ToNot(BeNil())

		})
		It(`UpdateSentimentModel request example`, func() {
			fmt.Println("\nUpdateSentimentModel() result:")
			// begin-updateSentimentModel

			updateSentimentModelOptions := naturalLanguageUnderstandingService.NewUpdateSentimentModelOptions(
				updateSentimentModelLink,
				"testString",
				CreateMockReader("This is a mock file."),
			)

			sentimentModel, response, err := naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sentimentModel, "", "  ")
			fmt.Println(string(b))

			// end-updateSentimentModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sentimentModel).ToNot(BeNil())

		})
		It(`ListCategoriesModels request example`, func() {
			fmt.Println("\nListCategoriesModels() result:")
			// begin-listCategoriesModels

			listCategoriesModelsOptions := naturalLanguageUnderstandingService.NewListCategoriesModelsOptions()

			categoriesModelList, response, err := naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(categoriesModelList, "", "  ")
			fmt.Println(string(b))

			// end-listCategoriesModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(categoriesModelList).ToNot(BeNil())

		})
		It(`GetCategoriesModel request example`, func() {
			fmt.Println("\nGetCategoriesModel() result:")
			// begin-getCategoriesModel

			getCategoriesModelOptions := naturalLanguageUnderstandingService.NewGetCategoriesModelOptions(
				getCategoriesModelLink,
			)

			categoriesModel, response, err := naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(categoriesModel, "", "  ")
			fmt.Println(string(b))

			// end-getCategoriesModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(categoriesModel).ToNot(BeNil())

		})
		It(`UpdateCategoriesModel request example`, func() {
			fmt.Println("\nUpdateCategoriesModel() result:")
			// begin-updateCategoriesModel

			updateCategoriesModelOptions := naturalLanguageUnderstandingService.NewUpdateCategoriesModelOptions(
				updateCategoriesModelLink,
				"testString",
				CreateMockReader("This is a mock file."),
			)

			categoriesModel, response, err := naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(categoriesModel, "", "  ")
			fmt.Println(string(b))

			// end-updateCategoriesModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(categoriesModel).ToNot(BeNil())

		})
		It(`ListClassificationsModels request example`, func() {
			fmt.Println("\nListClassificationsModels() result:")
			// begin-listClassificationsModels

			listClassificationsModelsOptions := naturalLanguageUnderstandingService.NewListClassificationsModelsOptions()

			classificationsModelList, response, err := naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classificationsModelList, "", "  ")
			fmt.Println(string(b))

			// end-listClassificationsModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classificationsModelList).ToNot(BeNil())

		})
		It(`GetClassificationsModel request example`, func() {
			fmt.Println("\nGetClassificationsModel() result:")
			// begin-getClassificationsModel

			getClassificationsModelOptions := naturalLanguageUnderstandingService.NewGetClassificationsModelOptions(
				getClassificationsModelLink,
			)

			classificationsModel, response, err := naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classificationsModel, "", "  ")
			fmt.Println(string(b))

			// end-getClassificationsModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classificationsModel).ToNot(BeNil())

		})
		It(`UpdateClassificationsModel request example`, func() {
			fmt.Println("\nUpdateClassificationsModel() result:")
			// begin-updateClassificationsModel

			updateClassificationsModelOptions := naturalLanguageUnderstandingService.NewUpdateClassificationsModelOptions(
				updateClassificationsModelLink,
				"testString",
				CreateMockReader("This is a mock file."),
			)

			classificationsModel, response, err := naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classificationsModel, "", "  ")
			fmt.Println(string(b))

			// end-updateClassificationsModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classificationsModel).ToNot(BeNil())

		})
		It(`DeleteSentimentModel request example`, func() {
			fmt.Println("\nDeleteSentimentModel() result:")
			// begin-deleteSentimentModel

			deleteSentimentModelOptions := naturalLanguageUnderstandingService.NewDeleteSentimentModelOptions(
				deleteSentimentModelLink,
			)

			deleteModelResults, response, err := naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteModelResults, "", "  ")
			fmt.Println(string(b))

			// end-deleteSentimentModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteModelResults).ToNot(BeNil())

		})
		It(`DeleteClassificationsModel request example`, func() {
			fmt.Println("\nDeleteClassificationsModel() result:")
			// begin-deleteClassificationsModel

			deleteClassificationsModelOptions := naturalLanguageUnderstandingService.NewDeleteClassificationsModelOptions(
				deleteClassificationsModelLink,
			)

			deleteModelResults, response, err := naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteModelResults, "", "  ")
			fmt.Println(string(b))

			// end-deleteClassificationsModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteModelResults).ToNot(BeNil())

		})
		It(`DeleteCategoriesModel request example`, func() {
			fmt.Println("\nDeleteCategoriesModel() result:")
			// begin-deleteCategoriesModel

			deleteCategoriesModelOptions := naturalLanguageUnderstandingService.NewDeleteCategoriesModelOptions(
				deleteCategoriesModelLink,
			)

			deleteModelResults, response, err := naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteModelResults, "", "  ")
			fmt.Println(string(b))

			// end-deleteCategoriesModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteModelResults).ToNot(BeNil())

		})
		It(`DeleteModel request example`, func() {
			fmt.Println("\nDeleteModel() result:")
			// begin-deleteModel

			deleteModelOptions := naturalLanguageUnderstandingService.NewDeleteModelOptions(
				"testString",
			)

			deleteModelResults, response, err := naturalLanguageUnderstandingService.DeleteModel(deleteModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteModelResults, "", "  ")
			fmt.Println(string(b))

			// end-deleteModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteModelResults).ToNot(BeNil())

		})
	})
})
