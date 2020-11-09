// +build integration

package toneanalyzerv3_test

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

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/toneanalyzerv3"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *toneanalyzerv3.ToneAnalyzerV3

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

	service, err = toneanalyzerv3.NewToneAnalyzerV3(
		&toneanalyzerv3.ToneAnalyzerV3Options{
			Version: core.StringPtr("2017-09-21"),
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

func TestTone(t *testing.T) {
	shouldSkipTest(t)

	text := "Team, I know that times are tough! Product sales have been disappointing for the past three quarters. We have a competitive product, but we need to do a better job of selling it!"
	tone, _, responseErr := service.Tone(
		&toneanalyzerv3.ToneOptions{
			ToneInput: &toneanalyzerv3.ToneInput{
				Text: &text,
			},
			ContentType: core.StringPtr("application/json"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, tone)
}

func TestToneChat(t *testing.T) {
	shouldSkipTest(t)

	utterances := []toneanalyzerv3.Utterance{
		{
			Text: core.StringPtr("Hello, I'm having a problem with your product."),
			User: core.StringPtr("customer"),
		},
		{
			Text: core.StringPtr("OK, let me know what's going on, please."),
			User: core.StringPtr("agent"),
		},
		{
			Text: core.StringPtr("Well, nothing is working :("),
			User: core.StringPtr("customer"),
		},
		{
			Text: core.StringPtr("Sorry to hear that."),
			User: core.StringPtr("agent"),
		},
	}

	options := service.NewToneChatOptions(utterances)
	toneChat, _, responseErr := service.ToneChat(options)
	assert.Nil(t, responseErr)
	assert.NotNil(t, toneChat)
}
