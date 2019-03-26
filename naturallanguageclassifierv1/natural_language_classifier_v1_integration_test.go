// +build integration

package naturallanguageclassifierv1_test

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
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
	"os"
	"testing"
)

var service *naturallanguageclassifierv1.NaturalLanguageClassifierV1
var serviceErr error
var classifier *naturallanguageclassifierv1.Classifier

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = naturallanguageclassifierv1.
			NewNaturalLanguageClassifierV1(&naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
				URL:      os.Getenv("NATURAL_LANGUAGE_CLASSIFIER_URL"),
				Username: os.Getenv("NATURAL_LANGUAGE_CLASSIFIER_USERNAME"),
				Password: os.Getenv("NATURAL_LANGUAGE_CLASSIFIER_PASSWORD"),
			})
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}
func TestClassifier(t *testing.T) {
	shouldSkipTest(t)

	// create classifier
	pwd, _ := os.Getwd()

	metadata, metadataErr := os.Open(pwd + "/../resources/weather_training_metadata.json")
	assert.Nil(t, metadataErr)

	trainingData, trainingDataErr := os.Open(pwd + "/../resources/weather_training_data.csv")
	assert.Nil(t, trainingDataErr)

	response, responseErr := service.CreateClassifier(
		&naturallanguageclassifierv1.CreateClassifierOptions{
			TrainingData: trainingData,
			Metadata:     metadata,
		},
	)
	assert.Nil(t, responseErr)

	createClassifier := service.GetCreateClassifierResult(response)
	assert.NotNil(t, createClassifier)

	// List classifier
	response, responseErr = service.ListClassifiers(
		&naturallanguageclassifierv1.ListClassifiersOptions{},
	)
	assert.Nil(t, responseErr)

	listClassifer := service.GetListClassifiersResult(response)
	assert.NotNil(t, listClassifer)

	// Get classifier
	response, responseErr = service.GetClassifier(
		&naturallanguageclassifierv1.GetClassifierOptions{
			ClassifierID: createClassifier.ClassifierID,
		},
	)
	assert.Nil(t, responseErr)

	getClassifer := service.GetGetClassifierResult(response)
	assert.NotNil(t, getClassifer)

	classifier = createClassifier
}

func TestClassify(t *testing.T) {
	shouldSkipTest(t)

	// classify
	if *classifier.Status != "Available" {
		t.Skip("Skip test classify")
	}
	response, responseErr := service.Classify(
		&naturallanguageclassifierv1.ClassifyOptions{
			ClassifierID: classifier.ClassifierID,
			Text:         core.StringPtr("How hot will it be today?"),
		},
	)
	assert.Nil(t, responseErr)
	fmt.Println(response)

	classify := service.GetClassifyResult(response)
	assert.NotNil(t, classify)
}

func TestClassifyCollection(t *testing.T) {
	shouldSkipTest(t)

	if *classifier.Status != "Available" {
		t.Skip("Skip test classify collection.")
	}
	// classify collection
	response, responseErr := service.ClassifyCollection(
		&naturallanguageclassifierv1.ClassifyCollectionOptions{
			ClassifierID: classifier.ClassifierID,
			Collection: []naturallanguageclassifierv1.ClassifyInput{
				naturallanguageclassifierv1.ClassifyInput{
					Text: core.StringPtr("How hot will it be today?"),
				},
				naturallanguageclassifierv1.ClassifyInput{
					Text: core.StringPtr("Is it hot outside?"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	classifyCollection := service.GetClassifyCollectionResult(response)
	assert.NotNil(t, classifyCollection)
}

func TestDeleteClassifier(t *testing.T) {
	shouldSkipTest(t)

	// Delete classifier
	response, responseErr := service.DeleteClassifier(
		&naturallanguageclassifierv1.DeleteClassifierOptions{
			ClassifierID: classifier.ClassifierID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
}
