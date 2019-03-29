// +build integration

package toneanalyzerv3_test

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
	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/toneanalyzerv3"
	"os"
	"testing"
)

var service *toneanalyzerv3.ToneAnalyzerV3
var serviceErr error

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = toneanalyzerv3.
			NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
				URL:      os.Getenv("TONE_ANALYZER_URL"),
				Version:  "2017-09-21",
				Username: os.Getenv("TONE_ANALYZER_USERNAME"),
				Password: os.Getenv("TONE_ANALYZER_PASSWORD"),
			})
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestTone(t *testing.T) {
	shouldSkipTest(t)

	text := "Team, I know that times are tough! Product sales have been disappointing for the past three quarters. We have a competitive product, but we need to do a better job of selling it!"
	response, responseErr := service.Tone(
		&toneanalyzerv3.ToneOptions{
			ToneInput: &toneanalyzerv3.ToneInput{
				Text: &text,
			},
			ContentType: core.StringPtr(toneanalyzerv3.ToneOptions_ContentType_ApplicationJSON),
		},
	)
	assert.Nil(t, responseErr)

	tone := service.GetToneResult(response)
	assert.NotNil(t, tone)
}

func TestToneChat(t *testing.T) {
	shouldSkipTest(t)

	utterances := []toneanalyzerv3.Utterance{
		toneanalyzerv3.Utterance{
			Text: core.StringPtr("Hello, I'm having a problem with your product."),
			User: core.StringPtr("customer"),
		},
		toneanalyzerv3.Utterance{
			Text: core.StringPtr("OK, let me know what's going on, please."),
			User: core.StringPtr("agent"),
		},
		toneanalyzerv3.Utterance{
			Text: core.StringPtr("Well, nothing is working :("),
			User: core.StringPtr("customer"),
		},
		toneanalyzerv3.Utterance{
			Text: core.StringPtr("Sorry to hear that."),
			User: core.StringPtr("agent"),
		},
	}

	options := service.NewToneChatOptions(utterances)
	response, responseErr := service.ToneChat(options)
	assert.Nil(t, responseErr)

	toneChat := service.GetToneChatResult(response)
	assert.NotNil(t, toneChat)
}
