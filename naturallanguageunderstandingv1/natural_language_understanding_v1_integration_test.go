// +build integration

package naturallanguageunderstandingv1_test

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
	"net/http"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

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
			Version: "2019-07-12",
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
