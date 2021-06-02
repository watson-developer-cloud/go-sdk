// +build integration

package personalityinsightsv3_test

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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/personalityinsightsv3"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *personalityinsightsv3.PersonalityInsightsV3

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

	service, err = personalityinsightsv3.NewPersonalityInsightsV3(
		&personalityinsightsv3.PersonalityInsightsV3Options{
			Version: core.StringPtr("2017-10-13"),
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

func TestProfile(t *testing.T) {
	shouldSkipTest(t)

	file, err := ioutil.ReadFile("../resources/personality-v3.json")
	assert.Nil(t, err)

	// Unmarshal JSON into Content struct
	content := new(personalityinsightsv3.Content)
	err = json.Unmarshal(file, content)
	assert.Nil(t, err)

	profile, response, responseErr := service.Profile(
		&personalityinsightsv3.ProfileOptions{
			Content:                content,
			ContentType:            core.StringPtr("application/json"),
			RawScores:              core.BoolPtr(true),
			ConsumptionPreferences: core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, profile)
}
