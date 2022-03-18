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

package naturallanguageclassifierv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/naturallanguageclassifierv1"
)

//
// This file provides an example of how to use the Natural Language Classifier service.
//
// The following configuration properties are assumed to be defined:
// NATURAL_LANGUAGE_CLASSIFIER_URL=<service base url>
// NATURAL_LANGUAGE_CLASSIFIER_AUTH_TYPE=iam
// NATURAL_LANGUAGE_CLASSIFIER_APIKEY=<IAM apikey>
// NATURAL_LANGUAGE_CLASSIFIER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../natural_language_classifier_v1.env"

var (
	naturalLanguageClassifierService *naturallanguageclassifierv1.NaturalLanguageClassifierV1
	config                           map[string]string
	configLoaded                     bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`NaturalLanguageClassifierV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(naturallanguageclassifierv1.DefaultServiceName)
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

			naturalLanguageClassifierServiceOptions := &naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{}

			naturalLanguageClassifierService, err = naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(naturalLanguageClassifierServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(naturalLanguageClassifierService).ToNot(BeNil())
		})
	})

	Describe(`NaturalLanguageClassifierV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Classify request example`, func() {
			fmt.Println("\nClassify() result:")
			// begin-classify

			classifyOptions := naturalLanguageClassifierService.NewClassifyOptions(
				"testString",
				"testString",
			)

			classification, response, err := naturalLanguageClassifierService.Classify(classifyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classification, "", "  ")
			fmt.Println(string(b))

			// end-classify

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classification).ToNot(BeNil())

		})
		It(`ClassifyCollection request example`, func() {
			fmt.Println("\nClassifyCollection() result:")
			// begin-classifyCollection

			classifyInputModel := &naturallanguageclassifierv1.ClassifyInput{
				Text: core.StringPtr("How hot will it be today?"),
			}

			classifyCollectionOptions := naturalLanguageClassifierService.NewClassifyCollectionOptions(
				"testString",
				[]naturallanguageclassifierv1.ClassifyInput{*classifyInputModel},
			)

			classificationCollection, response, err := naturalLanguageClassifierService.ClassifyCollection(classifyCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classificationCollection, "", "  ")
			fmt.Println(string(b))

			// end-classifyCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classificationCollection).ToNot(BeNil())

		})
		It(`CreateClassifier request example`, func() {
			fmt.Println("\nCreateClassifier() result:")
			// begin-createClassifier

			createClassifierOptions := naturalLanguageClassifierService.NewCreateClassifierOptions(
				CreateMockReader("This is a mock file."),
				CreateMockReader("This is a mock file."),
			)

			classifier, response, err := naturalLanguageClassifierService.CreateClassifier(createClassifierOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifier, "", "  ")
			fmt.Println(string(b))

			// end-createClassifier

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifier).ToNot(BeNil())

		})
		It(`ListClassifiers request example`, func() {
			fmt.Println("\nListClassifiers() result:")
			// begin-listClassifiers

			listClassifiersOptions := naturalLanguageClassifierService.NewListClassifiersOptions()

			classifierList, response, err := naturalLanguageClassifierService.ListClassifiers(listClassifiersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifierList, "", "  ")
			fmt.Println(string(b))

			// end-listClassifiers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifierList).ToNot(BeNil())

		})
		It(`GetClassifier request example`, func() {
			fmt.Println("\nGetClassifier() result:")
			// begin-getClassifier

			getClassifierOptions := naturalLanguageClassifierService.NewGetClassifierOptions(
				"testString",
			)

			classifier, response, err := naturalLanguageClassifierService.GetClassifier(getClassifierOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifier, "", "  ")
			fmt.Println(string(b))

			// end-getClassifier

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifier).ToNot(BeNil())

		})
		It(`DeleteClassifier request example`, func() {
			// begin-deleteClassifier

			deleteClassifierOptions := naturalLanguageClassifierService.NewDeleteClassifierOptions(
				"testString",
			)

			response, err := naturalLanguageClassifierService.DeleteClassifier(deleteClassifierOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteClassifier
			fmt.Printf("\nDeleteClassifier() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
