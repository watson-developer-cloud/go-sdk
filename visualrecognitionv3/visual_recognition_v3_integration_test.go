// +build integration

package visualrecognitionv3_test

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
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
)

var service *visualrecognitionv3.VisualRecognitionV3
var serviceErr error

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = visualrecognitionv3.
			NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
				URL:     os.Getenv("VISUAL_RECOGNITION_URL"),
				Version: "2018-03-19",
				Authenticator: &core.IamAuthenticator{
					ApiKey: os.Getenv("VISUAL_RECOGNITION_APIKEY"),
				},
			})

		if serviceErr == nil {
			customHeaders := http.Header{}
			customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
			customHeaders.Add("X-Watson-Test", "1")
			service.Service.SetDefaultHeaders(customHeaders)
		}
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestClassify(t *testing.T) {
	shouldSkipTest(t)

	// Classify
	pwd, _ := os.Getwd()
	imageFile, imageFileErr := os.Open(pwd + "/../resources/kitty.jpg")
	assert.Nil(t, imageFileErr)
	defer imageFile.Close()

	classify, _, responseErr := service.Classify(
		&visualrecognitionv3.ClassifyOptions{
			ImagesFile:    imageFile,
			Threshold:     core.Float32Ptr(0.6),
			ClassifierIds: []string{"default"},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, classify)
}

func TestClassifiers(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()

	// Create classifier
	carsFile, carsFileErr := os.Open(pwd + "/../resources/cars.zip")
	assert.Nil(t, carsFileErr)
	defer carsFile.Close()

	trucksFile, trucksFileErr := os.Open(pwd + "/../resources/trucks.zip")
	assert.Nil(t, trucksFileErr)
	defer trucksFile.Close()

	positiveExamples := make(map[string]*os.File)
	positiveExamples["cars"] = carsFile

	createClassifier, _, responseErr := service.CreateClassifier(
		&visualrecognitionv3.CreateClassifierOptions{
			Name:             core.StringPtr("Cars vs trucks"),
			PositiveExamples: positiveExamples,
			NegativeExamples: trucksFile,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createClassifier)

	// List classifiers
	listClassifiers, _, responseErr := service.ListClassifiers(
		&visualrecognitionv3.ListClassifiersOptions{
			Verbose: core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listClassifiers)

	// Get classifier
	getClassifier, _, responseErr := service.GetClassifier(
		&visualrecognitionv3.GetClassifierOptions{
			ClassifierID: createClassifier.ClassifierID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getClassifier)

	// Delete classifier
	_, responseErr = service.DeleteClassifier(
		&visualrecognitionv3.DeleteClassifierOptions{
			ClassifierID: createClassifier.ClassifierID,
		},
	)
	assert.Nil(t, responseErr)
}
