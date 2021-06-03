// +build integration

package visualrecognitionv3_test

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
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/visualrecognitionv3"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../../.env"

var service *visualrecognitionv3.VisualRecognitionV3

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

	service, err = visualrecognitionv3.NewVisualRecognitionV3(
		&visualrecognitionv3.VisualRecognitionV3Options{
			Version:     core.StringPtr("2018-03-19"),
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

func TestClassify(t *testing.T) {
	shouldSkipTest(t)

	// Classify
	imageFile, imageFileErr := os.Open("../resources/kitty.jpg")
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

	// Create classifier
	carsFile, carsFileErr := os.Open("../resources/cars.zip")
	assert.Nil(t, carsFileErr)
	defer carsFile.Close()

	trucksFile, trucksFileErr := os.Open("../resources/trucks.zip")
	assert.Nil(t, trucksFileErr)
	defer trucksFile.Close()

	positiveExamples := make(map[string]io.ReadCloser)
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
