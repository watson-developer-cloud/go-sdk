// +build integration

package personalityinsightsv3_test

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
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/watson-developer-cloud/go-sdk/core"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
	"io/ioutil"
	"os"
	"testing"
)

var service *personalityinsightsv3.PersonalityInsightsV3
var serviceErr error

func TestInitialization(t *testing.T) {
	err := godotenv.Load("../.env")
	require.Nil(t, err)

	service, serviceErr = personalityinsightsv3.
		NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
			URL:      os.Getenv("PERSONALITY_INSIGHTS_URL"),
			Version:  "2017-10-13",
			Username: os.Getenv("PERSONALITY_INSIGHTS_USERNAME"),
			Password: os.Getenv("PERSONALITY_INSIGHTS_PASSWORD"),
		})
	require.Nil(t, serviceErr)
}

func TestProfile(t *testing.T) {
	pwd, _ := os.Getwd()
	file, fileErr := ioutil.ReadFile(pwd + "/../resources/personality-v3.json")
	assert.Nil(t, fileErr)

	// Unmarshal JSON into Content struct
	content := new(personalityinsightsv3.Content)
	json.Unmarshal(file, content)

	response, responseErr := service.Profile(
		&personalityinsightsv3.ProfileOptions{
			Content:                content,
			ContentType:            core.StringPtr(personalityinsightsv3.ProfileOptions_ContentType_ApplicationJSON),
			RawScores:              core.BoolPtr(true),
			ConsumptionPreferences: core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)

	profile := service.GetProfileResult(response)
	assert.NotNil(t, profile)
}
