// +build integration

package naturallanguageunderstandingv1_test

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
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/watson-developer-cloud/go-sdk/core"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
	"os"
	"testing"
)

var service *naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1
var serviceErr error

func TestInitialization(t *testing.T) {
	err := godotenv.Load("../.env")
	require.Nil(t, err)

	service, serviceErr = naturallanguageunderstandingv1.
		NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
			URL:      os.Getenv("NATURAL_LANGUAGE_UNDERSTANDING_URL"),
			Version:  "2018-03-16",
			Username: os.Getenv("NATURAL_LANGUAGE_UNDERSTANDING_USERNAME"),
			Password: os.Getenv("NATURAL_LANGUAGE_UNDERSTANDING_PASSWORD"),
		})
	require.Nil(t, serviceErr)
}

func TestAnalyze(t *testing.T) {
	text := `IBM is an American multinational technology company
					 headquartered in Armonk, New York, United States
					with operations in over 170 countries.`
	response, responseErr := service.Analyze(
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

	analyze := service.GetAnalyzeResult(response)
	assert.NotNil(t, analyze)
}

func TestListModels(t *testing.T) {
	// list models
	t.Skip("Skip list models temporarily")
	response, responseErr := service.ListModels(
		&naturallanguageunderstandingv1.ListModelsOptions{},
	)
	assert.Nil(t, responseErr)

	listModels := service.GetListModelsResult(response)
	assert.NotNil(t, listModels)
}
