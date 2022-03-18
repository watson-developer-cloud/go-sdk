//go:build integration
// +build integration

package naturallanguageclassifierv1_test

/**
 * (C) Copyright IBM Corp. 2018, 2021.
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
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/naturallanguageclassifierv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../../.env"

var service *naturallanguageclassifierv1.NaturalLanguageClassifierV1
var testClassifierID string
var availableClassifierID string

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

	serviceOptions := &naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{}
	service, err = naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(serviceOptions)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestCreateClassifier(t *testing.T) {
	shouldSkipTest(t)

	// create classifier
	metadata, metadataErr := os.Open("../resources/weather_training_metadata.json")
	assert.Nil(t, metadataErr)

	trainingData, trainingDataErr := os.Open("../resources/weather_training_data.csv")
	assert.Nil(t, trainingDataErr)

	classifier, response, responseErr := service.CreateClassifier(
		&naturallanguageclassifierv1.CreateClassifierOptions{
			TrainingData:     trainingData,
			TrainingMetadata: metadata,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, classifier)
	assert.Equal(t, "Training", *classifier.Status)

	testClassifierID = *classifier.ClassifierID
}

func TestGetClassifier(t *testing.T) {
	shouldSkipTest(t)

	if testClassifierID == "" {
		t.Skip("Classifier ID not available, skipping test")
	}

	// Get classifier
	classifier, response, responseErr := service.GetClassifier(
		&naturallanguageclassifierv1.GetClassifierOptions{
			ClassifierID: &testClassifierID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, classifier)
	if classifier != nil {
		b, _ := json.MarshalIndent(classifier, "", "  ")
		t.Log("GetClassifier response: ", string(b))
	}
}

func TestListClassifiers(t *testing.T) {
	shouldSkipTest(t)

	// List classifier
	classifierList, response, responseErr := service.ListClassifiers(
		&naturallanguageclassifierv1.ListClassifiersOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, classifierList)
	if classifierList != nil {
		// Now fetch each classifier (to get the status) and find the first one
		// with a status of "Available".
		for _, c := range classifierList.Classifiers {
			t.Logf("Fetching classifier id=%s\n", *c.ClassifierID)
			classifier, _, err := service.GetClassifier(
				&naturallanguageclassifierv1.GetClassifierOptions{
					ClassifierID: c.ClassifierID,
				},
			)
			if err == nil {
				t.Logf("   status=%s\n", *classifier.Status)
				if *classifier.Status == naturallanguageclassifierv1.ClassifierStatusAvailableConst {
					availableClassifierID = *classifier.ClassifierID
					t.Logf("Found available classifier id: %s\n", availableClassifierID)
					break
				}
			}
		}
	}
}

func TestClassify(t *testing.T) {
	shouldSkipTest(t)

	if availableClassifierID == "" {
		t.Skip("Available classifier not found, skipping test...")
	}

	// classify
	classification, response, responseErr := service.Classify(
		&naturallanguageclassifierv1.ClassifyOptions{
			ClassifierID: &availableClassifierID,
			Text:         core.StringPtr("How hot will it be today?"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, classification)
}

func TestClassifyCollection(t *testing.T) {
	shouldSkipTest(t)

	if availableClassifierID == "" {
		t.Skip("Available classifier not found, skipping test...")
	}

	// classify collection
	classificationCollection, response, responseErr := service.ClassifyCollection(
		&naturallanguageclassifierv1.ClassifyCollectionOptions{
			ClassifierID: &availableClassifierID,
			Collection: []naturallanguageclassifierv1.ClassifyInput{
				{
					Text: core.StringPtr("How hot will it be today?"),
				},
				{
					Text: core.StringPtr("Is it hot outside?"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, classificationCollection)
}

func TestDeleteClassifier(t *testing.T) {
	shouldSkipTest(t)

	if testClassifierID == "" {
		t.Skip("No classifier to delete, skipping test...")
	}

	// Delete classifier
	response, responseErr := service.DeleteClassifier(
		&naturallanguageclassifierv1.DeleteClassifierOptions{
			ClassifierID: &testClassifierID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
}
