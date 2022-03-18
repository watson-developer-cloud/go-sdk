//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2021, 2022.
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

package visualrecognitionv4_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/visualrecognitionv4"
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
const externalConfigFile = "../visual_recognition_v4.env"

var (
	visualRecognitionService *visualrecognitionv4.VisualRecognitionV4
	config                   map[string]string
	configLoaded             bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`VisualRecognitionV4 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(visualrecognitionv4.DefaultServiceName)
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

			visualRecognitionServiceOptions := &visualrecognitionv4.VisualRecognitionV4Options{
				Version: core.StringPtr("testString"),
			}

			visualRecognitionService, err = visualrecognitionv4.NewVisualRecognitionV4(visualRecognitionServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(visualRecognitionService).ToNot(BeNil())
		})
	})

	Describe(`VisualRecognitionV4 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Analyze request example`, func() {
			fmt.Println("\nAnalyze() result:")
			// begin-analyze

			analyzeOptions := visualRecognitionService.NewAnalyzeOptions(
				[]string{"testString"},
				[]string{"objects"},
			)

			analyzeResponse, response, err := visualRecognitionService.Analyze(analyzeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyzeResponse, "", "  ")
			fmt.Println(string(b))

			// end-analyze

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyzeResponse).ToNot(BeNil())

		})
		It(`CreateCollection request example`, func() {
			fmt.Println("\nCreateCollection() result:")
			// begin-createCollection

			createCollectionOptions := visualRecognitionService.NewCreateCollectionOptions()

			collection, response, err := visualRecognitionService.CreateCollection(createCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-createCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collection).ToNot(BeNil())

		})
		It(`ListCollections request example`, func() {
			fmt.Println("\nListCollections() result:")
			// begin-listCollections

			listCollectionsOptions := visualRecognitionService.NewListCollectionsOptions()

			collectionsList, response, err := visualRecognitionService.ListCollections(listCollectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectionsList, "", "  ")
			fmt.Println(string(b))

			// end-listCollections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectionsList).ToNot(BeNil())

		})
		It(`GetCollection request example`, func() {
			fmt.Println("\nGetCollection() result:")
			// begin-getCollection

			getCollectionOptions := visualRecognitionService.NewGetCollectionOptions(
				"testString",
			)

			collection, response, err := visualRecognitionService.GetCollection(getCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-getCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collection).ToNot(BeNil())

		})
		It(`UpdateCollection request example`, func() {
			fmt.Println("\nUpdateCollection() result:")
			// begin-updateCollection

			updateCollectionOptions := visualRecognitionService.NewUpdateCollectionOptions(
				"testString",
			)

			collection, response, err := visualRecognitionService.UpdateCollection(updateCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-updateCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collection).ToNot(BeNil())

		})
		It(`GetModelFile request example`, func() {
			fmt.Println("\nGetModelFile() result:")
			// begin-getModelFile

			getModelFileOptions := visualRecognitionService.NewGetModelFileOptions(
				"testString",
				"objects",
				"rscnn",
			)

			file, response, err := visualRecognitionService.GetModelFile(getModelFileOptions)
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

			// end-getModelFile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(file).ToNot(BeNil())

		})
		It(`AddImages request example`, func() {
			fmt.Println("\nAddImages() result:")
			// begin-addImages

			addImagesOptions := visualRecognitionService.NewAddImagesOptions(
				"testString",
			)
			addImagesOptions.SetTrainingData(`{"objects":[{"object":"2018-Fit","location":{"left":33,"top":8,"width":760,"height":419}}]}`)

			imageDetailsList, response, err := visualRecognitionService.AddImages(addImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageDetailsList, "", "  ")
			fmt.Println(string(b))

			// end-addImages

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDetailsList).ToNot(BeNil())

		})
		It(`ListImages request example`, func() {
			fmt.Println("\nListImages() result:")
			// begin-listImages

			listImagesOptions := visualRecognitionService.NewListImagesOptions(
				"testString",
			)

			imageSummaryList, response, err := visualRecognitionService.ListImages(listImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageSummaryList, "", "  ")
			fmt.Println(string(b))

			// end-listImages

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageSummaryList).ToNot(BeNil())

		})
		It(`GetImageDetails request example`, func() {
			fmt.Println("\nGetImageDetails() result:")
			// begin-getImageDetails

			getImageDetailsOptions := visualRecognitionService.NewGetImageDetailsOptions(
				"testString",
				"testString",
			)

			imageDetails, response, err := visualRecognitionService.GetImageDetails(getImageDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageDetails, "", "  ")
			fmt.Println(string(b))

			// end-getImageDetails

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDetails).ToNot(BeNil())

		})
		It(`GetJpegImage request example`, func() {
			fmt.Println("\nGetJpegImage() result:")
			// begin-getJpegImage

			getJpegImageOptions := visualRecognitionService.NewGetJpegImageOptions(
				"testString",
				"testString",
			)

			jpegImage, response, err := visualRecognitionService.GetJpegImage(getJpegImageOptions)
			if err != nil {
				panic(err)
			}
			if jpegImage != nil {
				defer jpegImage.Close()
				outFile, err := os.Create("jpegImage.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, jpegImage)
				if err != nil {
					panic(err)
				}
			}

			// end-getJpegImage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jpegImage).ToNot(BeNil())

		})
		It(`ListObjectMetadata request example`, func() {
			fmt.Println("\nListObjectMetadata() result:")
			// begin-listObjectMetadata

			listObjectMetadataOptions := visualRecognitionService.NewListObjectMetadataOptions(
				"testString",
			)

			objectMetadataList, response, err := visualRecognitionService.ListObjectMetadata(listObjectMetadataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectMetadataList, "", "  ")
			fmt.Println(string(b))

			// end-listObjectMetadata

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectMetadataList).ToNot(BeNil())

		})
		It(`UpdateObjectMetadata request example`, func() {
			fmt.Println("\nUpdateObjectMetadata() result:")
			// begin-updateObjectMetadata

			updateObjectMetadataOptions := visualRecognitionService.NewUpdateObjectMetadataOptions(
				"testString",
				"testString",
				"testString",
			)

			updateObjectMetadata, response, err := visualRecognitionService.UpdateObjectMetadata(updateObjectMetadataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updateObjectMetadata, "", "  ")
			fmt.Println(string(b))

			// end-updateObjectMetadata

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateObjectMetadata).ToNot(BeNil())

		})
		It(`GetObjectMetadata request example`, func() {
			fmt.Println("\nGetObjectMetadata() result:")
			// begin-getObjectMetadata

			getObjectMetadataOptions := visualRecognitionService.NewGetObjectMetadataOptions(
				"testString",
				"testString",
			)

			objectMetadata, response, err := visualRecognitionService.GetObjectMetadata(getObjectMetadataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(objectMetadata, "", "  ")
			fmt.Println(string(b))

			// end-getObjectMetadata

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(objectMetadata).ToNot(BeNil())

		})
		It(`Train request example`, func() {
			fmt.Println("\nTrain() result:")
			// begin-train

			trainOptions := visualRecognitionService.NewTrainOptions(
				"testString",
			)

			collection, response, err := visualRecognitionService.Train(trainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-train

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(collection).ToNot(BeNil())

		})
		It(`AddImageTrainingData request example`, func() {
			fmt.Println("\nAddImageTrainingData() result:")
			// begin-addImageTrainingData

			addImageTrainingDataOptions := visualRecognitionService.NewAddImageTrainingDataOptions(
				"testString",
				"testString",
			)

			trainingDataObjects, response, err := visualRecognitionService.AddImageTrainingData(addImageTrainingDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingDataObjects, "", "  ")
			fmt.Println(string(b))

			// end-addImageTrainingData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingDataObjects).ToNot(BeNil())

		})
		It(`GetTrainingUsage request example`, func() {
			fmt.Println("\nGetTrainingUsage() result:")
			// begin-getTrainingUsage

			getTrainingUsageOptions := visualRecognitionService.NewGetTrainingUsageOptions()

			trainingEvents, response, err := visualRecognitionService.GetTrainingUsage(getTrainingUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingEvents, "", "  ")
			fmt.Println(string(b))

			// end-getTrainingUsage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingEvents).ToNot(BeNil())

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
		It(`DeleteObject request example`, func() {
			// begin-deleteObject

			deleteObjectOptions := visualRecognitionService.NewDeleteObjectOptions(
				"testString",
				"testString",
			)

			response, err := visualRecognitionService.DeleteObject(deleteObjectOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteObject
			fmt.Printf("\nDeleteObject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteImage request example`, func() {
			// begin-deleteImage

			deleteImageOptions := visualRecognitionService.NewDeleteImageOptions(
				"testString",
				"testString",
			)

			response, err := visualRecognitionService.DeleteImage(deleteImageOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteImage
			fmt.Printf("\nDeleteImage() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteCollection request example`, func() {
			// begin-deleteCollection

			deleteCollectionOptions := visualRecognitionService.NewDeleteCollectionOptions(
				"testString",
			)

			response, err := visualRecognitionService.DeleteCollection(deleteCollectionOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteCollection
			fmt.Printf("\nDeleteCollection() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
