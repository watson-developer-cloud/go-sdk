//go:build integration
// +build integration

package naturallanguageunderstandingv1_test

/**
 * (C) Copyright IBM Corp. 2018, 2022.
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
	"strconv"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v3/naturallanguageunderstandingv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../../.env"

var service *naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1

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

	service, err = naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(
		&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			Version: core.StringPtr("2019-07-12"),
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

func TestAnalyze(t *testing.T) {
	shouldSkipTest(t)

	text := `IBM is an American multinational technology company
					 headquartered in Armonk, New York, United States
					with operations in over 170 countries.`
	analyze, response, responseErr := service.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			Text: &text,
			Features: &naturallanguageunderstandingv1.Features{
				Emotion: &naturallanguageunderstandingv1.EmotionOptions{
					Document: core.BoolPtr(true),
				},
				Sentiment: &naturallanguageunderstandingv1.SentimentOptions{
					Document: core.BoolPtr(true),
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, analyze)
}

func TestListModels(t *testing.T) {
	shouldSkipTest(t)

	// list models
	listModels, response, responseErr := service.ListModels(
		&naturallanguageunderstandingv1.ListModelsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, listModels)
}

func TestSentimentModelsCRUD(t *testing.T) {
	shouldSkipTest(t)

	trainingData, err := os.Open("../resources/nlu_training_data.csv")
	assert.Nil(t, err)

	createdModel, _, createModelErr := service.CreateSentimentModel(
		&naturallanguageunderstandingv1.CreateSentimentModelOptions{
			Name:         core.StringPtr("Go test sentiment model"),
			Description:  core.StringPtr("go test model"),
			Language:     core.StringPtr("en"),
			TrainingData: trainingData,
		},
	)

	assert.Nil(t, createModelErr)
	assert.Equal(t, "Go test sentiment model", *createdModel.Name)

	trainingData.Close()

	readModel, _, readModelErr := service.GetSentimentModel(
		&naturallanguageunderstandingv1.GetSentimentModelOptions{
			ModelID: createdModel.ModelID,
		},
	)

	assert.Nil(t, readModelErr)
	assert.Equal(t, "Go test sentiment model", *readModel.Name)

	listedModels, _, listModelsErr := service.ListSentimentModels(
		&naturallanguageunderstandingv1.ListSentimentModelsOptions{},
	)

	assert.Nil(t, listModelsErr)
	assert.NotNil(t, listedModels)

	updatedTrainingData, err := os.Open("../resources/nlu_training_data.csv")
	assert.Nil(t, err)
	defer updatedTrainingData.Close()

	updatedModel, _, updateModelErr := service.UpdateSentimentModel(
		&naturallanguageunderstandingv1.UpdateSentimentModelOptions{
			ModelID:      createdModel.ModelID,
			Name:         core.StringPtr("Go updated model"),
			Description:  core.StringPtr("I'm updated"),
			Language:     core.StringPtr("en"),
			TrainingData: updatedTrainingData,
		},
	)

	assert.Nil(t, updateModelErr)
	assert.Equal(t, "Go updated model", *updatedModel.Name)

	deletedModel, _, deleteModelErr := service.DeleteSentimentModel(
		&naturallanguageunderstandingv1.DeleteSentimentModelOptions{
			ModelID: createdModel.ModelID,
		},
	)

	assert.Nil(t, deleteModelErr)
	assert.NotNil(t, deletedModel)
}

func TestClassificationsModelCRUD(t *testing.T) {
	shouldSkipTest(t)

	trainingData, err := os.Open("../resources/nlu_classifications_training.json")
	assert.Nil(t, err)

	hashedModelName := strconv.FormatInt(time.Now().UnixNano(), 10)

	createdModel, _, createModelErr := service.CreateClassificationsModel(
		&naturallanguageunderstandingv1.CreateClassificationsModelOptions{
			Name:                    core.StringPtr("Go" + hashedModelName),
			Description:             core.StringPtr("Description"),
			TrainingData:            trainingData,
			TrainingDataContentType: core.StringPtr("application/json"),
			Language:                core.StringPtr("en"),
		},
	)

	assert.Nil(t, createModelErr)
	assert.Equal(t, "Go"+hashedModelName, *createdModel.Name)

	trainingData.Close()

	readModel, _, readModelErr := service.GetClassificationsModel(
		&naturallanguageunderstandingv1.GetClassificationsModelOptions{
			ModelID: createdModel.ModelID,
		},
	)

	assert.Nil(t, readModelErr)
	assert.Equal(t, *createdModel.Name, *readModel.Name)

	listedModels, _, listModelsErr := service.ListClassificationsModels(
		&naturallanguageunderstandingv1.ListClassificationsModelsOptions{},
	)

	assert.Nil(t, listModelsErr)
	assert.NotNil(t, listedModels)

	updateTrainingData, err := os.Open("../resources/nlu_classifications_training.json")
	assert.Nil(t, err)
	defer updateTrainingData.Close()

	updatedModel, _, updateModelErr := service.UpdateClassificationsModel(
		&naturallanguageunderstandingv1.UpdateClassificationsModelOptions{
			ModelID:                 createdModel.ModelID,
			Name:                    core.StringPtr("Go update" + hashedModelName),
			Description:             core.StringPtr("Description"),
			TrainingData:            updateTrainingData,
			TrainingDataContentType: core.StringPtr("application/json"),
			Language:                core.StringPtr("en"),
		},
	)

	assert.Nil(t, updateModelErr)
	assert.Equal(t, "Go update"+hashedModelName, *updatedModel.Name)

	deletedModel, _, deleteModelErr := service.DeleteClassificationsModel(
		&naturallanguageunderstandingv1.DeleteClassificationsModelOptions{
			ModelID: createdModel.ModelID,
		},
	)

	assert.Nil(t, deleteModelErr)
	assert.NotNil(t, deletedModel)
}

func TestCategoriesModels(t *testing.T) {
	shouldSkipTest(t)

	trainingData, err := os.Open("../resources/nlu_categories_training.json")
	assert.Nil(t, err)

	hashedModelName := strconv.FormatInt(time.Now().UnixNano(), 10)

	createdModel, _, createModelErr := service.CreateCategoriesModel(
		&naturallanguageunderstandingv1.CreateCategoriesModelOptions{
			Name:                    core.StringPtr("Go" + hashedModelName),
			Description:             core.StringPtr("Description"),
			TrainingData:            trainingData,
			TrainingDataContentType: core.StringPtr("application/json"),
			Language:                core.StringPtr("en"),
		},
	)

	assert.Nil(t, createModelErr)
	assert.Equal(t, "Go"+hashedModelName, *createdModel.Name)

	trainingData.Close()

	readModel, _, readModelErr := service.GetCategoriesModel(
		&naturallanguageunderstandingv1.GetCategoriesModelOptions{
			ModelID: createdModel.ModelID,
		},
	)

	assert.Nil(t, readModelErr)
	assert.Equal(t, *createdModel.Name, *readModel.Name)

	listedModels, _, listModelsErr := service.ListCategoriesModels(
		&naturallanguageunderstandingv1.ListCategoriesModelsOptions{},
	)

	assert.Nil(t, listModelsErr)
	assert.NotNil(t, listedModels)

	updateTrainingData, err := os.Open("../resources/nlu_categories_training.json")
	assert.Nil(t, err)
	defer updateTrainingData.Close()

	updatedModel, _, updateModelErr := service.UpdateCategoriesModel(
		&naturallanguageunderstandingv1.UpdateCategoriesModelOptions{
			ModelID:                 createdModel.ModelID,
			Name:                    core.StringPtr("Go update" + hashedModelName),
			Description:             core.StringPtr("Description"),
			TrainingData:            updateTrainingData,
			TrainingDataContentType: core.StringPtr("application/json"),
			Language:                core.StringPtr("en"),
		},
	)

	assert.Nil(t, updateModelErr)
	assert.Equal(t, "Go update"+hashedModelName, *updatedModel.Name)

	deletedModel, _, deleteModelErr := service.DeleteCategoriesModel(
		&naturallanguageunderstandingv1.DeleteCategoriesModelOptions{
			ModelID: createdModel.ModelID,
		},
	)

	assert.Nil(t, deleteModelErr)
	assert.NotNil(t, deletedModel)
}
