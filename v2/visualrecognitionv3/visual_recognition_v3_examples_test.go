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

package visualrecognitionv3_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/visualrecognitionv3"
)

//
// This file provides an example of how to use the Visual Recognition service.
//
// The following configuration properties are assumed to be defined:
// WATSON_VISION_COMBINED_URL=<service base url>
// WATSON_VISION_COMBINED_AUTH_TYPE=iam
// WATSON_VISION_COMBINED_APIKEY=<IAM apikey>
// WATSON_VISION_COMBINED_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../visual_recognition_v3.env"

var (
	visualRecognitionService *visualrecognitionv3.VisualRecognitionV3
	config                   map[string]string
	configLoaded             bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`VisualRecognitionV3 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(visualrecognitionv3.DefaultServiceName)
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

			visualRecognitionServiceOptions := &visualrecognitionv3.VisualRecognitionV3Options{
				Version: core.StringPtr("testString"),
			}

			visualRecognitionService, err = visualrecognitionv3.NewVisualRecognitionV3(visualRecognitionServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(visualRecognitionService).ToNot(BeNil())
		})
	})

	Describe(`VisualRecognitionV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Classify request example`, func() {
			fmt.Println("\nClassify() result:")
			// begin-classify

			classifyOptions := visualRecognitionService.NewClassifyOptions()

			classifiedImages, response, err := visualRecognitionService.Classify(classifyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifiedImages, "", "  ")
			fmt.Println(string(b))

			// end-classify

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifiedImages).ToNot(BeNil())

		})
		It(`CreateClassifier request example`, func() {
			fmt.Println("\nCreateClassifier() result:")
			// begin-createClassifier

			createClassifierOptions := visualRecognitionService.NewCreateClassifierOptions(
				"testString",
			)
			createClassifierOptions.AddPositiveExamples("key1", CreateMockReader("This is a mock file."))

			classifier, response, err := visualRecognitionService.CreateClassifier(createClassifierOptions)
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

			listClassifiersOptions := visualRecognitionService.NewListClassifiersOptions()

			classifiers, response, err := visualRecognitionService.ListClassifiers(listClassifiersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifiers, "", "  ")
			fmt.Println(string(b))

			// end-listClassifiers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifiers).ToNot(BeNil())

		})
		It(`GetClassifier request example`, func() {
			fmt.Println("\nGetClassifier() result:")
			// begin-getClassifier

			getClassifierOptions := visualRecognitionService.NewGetClassifierOptions(
				"testString",
			)

			classifier, response, err := visualRecognitionService.GetClassifier(getClassifierOptions)
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
		It(`UpdateClassifier request example`, func() {
			fmt.Println("\nUpdateClassifier() result:")
			// begin-updateClassifier

			updateClassifierOptions := visualRecognitionService.NewUpdateClassifierOptions(
				"testString",
			)

			classifier, response, err := visualRecognitionService.UpdateClassifier(updateClassifierOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(classifier, "", "  ")
			fmt.Println(string(b))

			// end-updateClassifier

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(classifier).ToNot(BeNil())

		})
		It(`GetCoreMlModel request example`, func() {
			fmt.Println("\nGetCoreMlModel() result:")
			// begin-getCoreMlModel

			getCoreMlModelOptions := visualRecognitionService.NewGetCoreMlModelOptions(
				"testString",
			)

			file, response, err := visualRecognitionService.GetCoreMlModel(getCoreMlModelOptions)
			if err != nil {
				panic(err)
			}
			if file != nil {
				defer file.Close()
				outFile, err := os.Create("file.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, file)
				if err != nil {
					panic(err)
				}
			}

			// end-getCoreMlModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(file).ToNot(BeNil())

		})
		It(`DeleteUserData request example`, func() {
			// begin-deleteUserData

			deleteUserDataOptions := visualRecognitionService.NewDeleteUserDataOptions(
				"testString",
			)

			response, err := visualRecognitionService.DeleteUserData(deleteUserDataOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteUserData
			fmt.Printf("\nDeleteUserData() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteClassifier request example`, func() {
			// begin-deleteClassifier

			deleteClassifierOptions := visualRecognitionService.NewDeleteClassifierOptions(
				"testString",
			)

			response, err := visualRecognitionService.DeleteClassifier(deleteClassifierOptions)
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
