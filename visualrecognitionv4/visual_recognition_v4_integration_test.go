// +build integration

package visualrecognitionv4_test

/**
 * (C) Copyright IBM Corp. 2019, 2020.
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
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv4"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *visualrecognitionv4.VisualRecognitionV4
var collectionID string

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load(configFile)
	if err != nil {
		t.Skip(skipMessage)
	}

	collectionID = os.Getenv("VISUAL_RECOGNITION_COLLECTION_ID")
	assert.NotEmpty(t, collectionID)
	if collectionID != "" {
		configLoaded = true
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Version:     "2019-02-11",
			ServiceName: "visual_recognition",
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

func TestCollection(t *testing.T) {
	shouldSkipTest(t)

	collection, _, responseErr := service.CreateCollection(
		&visualrecognitionv4.CreateCollectionOptions{
			Name:        core.StringPtr("my_go_collection"),
			Description: core.StringPtr("simple collection for go"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, collection)
	collectionId := collection.CollectionID

	myGOCollection, _, responseErr := service.GetCollection(
		&visualrecognitionv4.GetCollectionOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, myGOCollection)

	updatedCollection, _, responseErr := service.UpdateCollection(
		&visualrecognitionv4.UpdateCollectionOptions{
			CollectionID: collectionId,
			Description:  core.StringPtr("some random description"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updatedCollection)

	listCollections, _, responseErr := service.ListCollections(
		&visualrecognitionv4.ListCollectionsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listCollections)

	_, responseErr = service.DeleteCollection(
		&visualrecognitionv4.DeleteCollectionOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
}

func TestImages(t *testing.T) {
	shouldSkipTest(t)

	collection, _, responseErr := service.CreateCollection(
		&visualrecognitionv4.CreateCollectionOptions{
			Name:        core.StringPtr("my_go_collection"),
			Description: core.StringPtr("simple collection for go"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, collection)
	collectionId := collection.CollectionID

	kittyFile, err := os.Open("../resources/kitty.jpg")
	assert.Nil(t, err)
	defer kittyFile.Close()

	giraffeFile, err := os.Open("../resources/my-giraffe.jpeg")
	assert.Nil(t, err)
	defer giraffeFile.Close()
	addImages, _, responseErr := service.AddImages(
		&visualrecognitionv4.AddImagesOptions{
			CollectionID: collectionId,
			ImagesFile: []visualrecognitionv4.FileWithMetadata{
				visualrecognitionv4.FileWithMetadata{
					Data:     kittyFile,
					Filename: core.StringPtr("hello kitty"),
				},
				visualrecognitionv4.FileWithMetadata{
					Data:     giraffeFile,
					Filename: core.StringPtr("hello giraffe"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, addImages)

	listImages, _, responseErr := service.ListImages(
		&visualrecognitionv4.ListImagesOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listImages)

	imageDetails, _, responseErr := service.GetImageDetails(
		&visualrecognitionv4.GetImageDetailsOptions{
			CollectionID: collectionId,
			ImageID:      addImages.Images[0].ImageID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, imageDetails)

	readCloser, _, responseErr := service.GetJpegImage(
		&visualrecognitionv4.GetJpegImageOptions{
			CollectionID: collectionId,
			ImageID:      imageDetails.ImageID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, readCloser)

	_, responseErr = service.DeleteImage(
		&visualrecognitionv4.DeleteImageOptions{
			CollectionID: collectionId,
			ImageID:      imageDetails.ImageID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, readCloser)

	_, responseErr = service.DeleteCollection(
		&visualrecognitionv4.DeleteCollectionOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
}

func TestAnalyze(t *testing.T) {
	shouldSkipTest(t)

	giraffeFile, err := os.Open("../resources/my-giraffe.jpeg")
	assert.Nil(t, err)
	defer giraffeFile.Close()

	analyzeResult, _, responseErr := service.Analyze(
		&visualrecognitionv4.AnalyzeOptions{
			CollectionIds: []string{collectionID},
			Features:      []string{visualrecognitionv4.AnalyzeOptions_Features_Objects},
			ImagesFile: []visualrecognitionv4.FileWithMetadata{
				visualrecognitionv4.FileWithMetadata{
					Data:     giraffeFile,
					Filename: core.StringPtr("my giraffe"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, analyzeResult)
}

func TestTraining(t *testing.T) {
	shouldSkipTest(t)

	collection, _, responseErr := service.CreateCollection(
		&visualrecognitionv4.CreateCollectionOptions{
			Name:        core.StringPtr("my_go_collection_for_training"),
			Description: core.StringPtr("simple collection for go"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, collection)
	collectionId := collection.CollectionID

	giraffeFile, err := os.Open("../resources/South_Africa_Luca_Galuzzi_2004.jpeg")
	assert.Nil(t, err)
	defer giraffeFile.Close()
	addImages, _, responseErr := service.AddImages(
		&visualrecognitionv4.AddImagesOptions{
			CollectionID: collectionId,
			ImagesFile: []visualrecognitionv4.FileWithMetadata{
				visualrecognitionv4.FileWithMetadata{
					Data:     giraffeFile,
					Filename: core.StringPtr("hello giraffe"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, addImages)

	trainingData, _, responseErr := service.AddImageTrainingData(
		&visualrecognitionv4.AddImageTrainingDataOptions{
			CollectionID: collectionId,
			ImageID:      addImages.Images[0].ImageID,
			Objects: []visualrecognitionv4.TrainingDataObject{
				visualrecognitionv4.TrainingDataObject{
					Object: core.StringPtr("giraffe training data"),
					Location: &visualrecognitionv4.Location{
						Top:    core.Int64Ptr(64),
						Left:   core.Int64Ptr(270),
						Width:  core.Int64Ptr(755),
						Height: core.Int64Ptr(784),
					},
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, trainingData)

	objectMetadataList, _, responseErr := service.ListObjectMetadata(
		&visualrecognitionv4.ListObjectMetadataOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, objectMetadataList)

	objectMetadata := objectMetadataList.Objects[0]
	updatedObjectMetadata, _, responseErr := service.UpdateObjectMetadata(
		&visualrecognitionv4.UpdateObjectMetadataOptions{
			CollectionID: collectionId,
			Object:       objectMetadata.Object,
			NewObject:    core.StringPtr("updated giraffe training data"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updatedObjectMetadata)

	objMetadata, _, responseErr := service.GetObjectMetadata(
		&visualrecognitionv4.GetObjectMetadataOptions{
			CollectionID: collectionId,
			Object:       core.StringPtr("updated giraffe training data"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, objMetadata)

	train, _, responseErr := service.Train(
		&visualrecognitionv4.TrainOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, train)

	trainingUsage, response, trainingUsageErr := service.GetTrainingUsage(
		&visualrecognitionv4.GetTrainingUsageOptions{},
	)
	assert.Nil(t, trainingUsageErr)
	assert.NotNil(t, response)
	assert.NotNil(t, trainingUsage)

	_, responseErr = service.DeleteObject(
		&visualrecognitionv4.DeleteObjectOptions{
			CollectionID: collectionId,
			Object:       core.StringPtr("updated giraffe training data"),
		},
	)
	assert.Nil(t, responseErr)

	_, responseErr = service.DeleteCollection(
		&visualrecognitionv4.DeleteCollectionOptions{
			CollectionID: collectionId,
		},
	)
	assert.Nil(t, responseErr)
}

func TestGetModel(t *testing.T) {
	shouldSkipTest(t)

	result, response, err := service.GetModelFile(
		&visualrecognitionv4.GetModelFileOptions{
			CollectionID: core.StringPtr(collectionID),
			Feature:      core.StringPtr("objects"),
			ModelFormat:  core.StringPtr("rscnn"),
		},
	)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, response)
}
