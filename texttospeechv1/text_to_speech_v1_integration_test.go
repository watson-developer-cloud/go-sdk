// +build integration

package texttospeechv1_test

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
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
	"os"
	"testing"
)

var service *texttospeechv1.TextToSpeechV1
var serviceErr error
var customizationID *string

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = texttospeechv1.
			NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				URL:      os.Getenv("TEXT_TO_SPEECH_URL"),
				Username: os.Getenv("TEXT_TO_SPEECH_USERNAME"),
				Password: os.Getenv("TEXT_TO_SPEECH_PASSWORD"),
			})
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestVoice(t *testing.T) {
	shouldSkipTest(t)

	// list voices
	response, responseErr := service.ListVoices(
		&texttospeechv1.ListVoicesOptions{},
	)
	assert.Nil(t, responseErr)

	listVoices := service.GetListVoicesResult(response)
	assert.NotNil(t, listVoices)

	// Get voice
	response, responseErr = service.GetVoice(
		&texttospeechv1.GetVoiceOptions{
			Voice: core.StringPtr(texttospeechv1.GetVoiceOptions_Voice_EnUsAllisonvoice),
		},
	)

	voice := service.GetGetVoiceResult(response)
	assert.NotNil(t, voice)
}

func TestSynthesize(t *testing.T) {
	shouldSkipTest(t)

	// synthesize
	response, responseErr := service.Synthesize(
		&texttospeechv1.SynthesizeOptions{
			Text:   core.StringPtr("Hello world"),
			Accept: core.StringPtr(texttospeechv1.SynthesizeOptions_Accept_AudioWav),
			Voice:  core.StringPtr(texttospeechv1.SynthesizeOptions_Voice_EnUsAllisonvoice),
		},
	)
	assert.Nil(t, responseErr)

	synthesize := service.GetSynthesizeResult(response)
	assert.NotNil(t, synthesize)
	synthesize.Close()
}

func TestPronunciation(t *testing.T) {
	shouldSkipTest(t)

	// get pronunciation
	response, responseErr := service.GetPronunciation(
		&texttospeechv1.GetPronunciationOptions{
			Text:   core.StringPtr("IEEE"),
			Voice:  core.StringPtr(texttospeechv1.GetPronunciationOptions_Voice_EnUsAllisonvoice),
			Format: core.StringPtr("ibm"),
		},
	)
	assert.Nil(t, responseErr)

	pronunciation := service.GetGetPronunciationResult(response)
	assert.NotNil(t, pronunciation)
}

func TestVoiceModel(t *testing.T) {
	shouldSkipTest(t)

	// create voice model
	response, responseErr := service.CreateVoiceModel(
		&texttospeechv1.CreateVoiceModelOptions{
			Name:        core.StringPtr("First model for GO"),
			Language:    core.StringPtr(texttospeechv1.CreateVoiceModelOptions_Language_EnUs),
			Description: core.StringPtr("First custom voice model"),
		},
	)
	assert.Nil(t, responseErr)

	createVoiceModel := service.GetCreateVoiceModelResult(response)
	assert.NotNil(t, createVoiceModel)

	// List voice models
	response, responseErr = service.ListVoiceModels(
		&texttospeechv1.ListVoiceModelsOptions{},
	)
	assert.Nil(t, responseErr)

	listVoiceModels := service.GetListVoiceModelsResult(response)
	assert.NotNil(t, listVoiceModels)

	// Update voice model
	response, responseErr = service.UpdateVoiceModel(
		&texttospeechv1.UpdateVoiceModelOptions{
			CustomizationID: createVoiceModel.CustomizationID,
			Name:            core.StringPtr("First Model Update for GO"),
			Description:     core.StringPtr("First custom voice model update"),
			Words: []texttospeechv1.Word{
				texttospeechv1.Word{
					Word:        core.StringPtr("NCAA"),
					Translation: core.StringPtr("N C double A"),
				},
				texttospeechv1.Word{
					Word:        core.StringPtr("iPhone"),
					Translation: core.StringPtr("I phone"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	// Get voice model
	response, responseErr = service.GetVoiceModel(
		&texttospeechv1.GetVoiceModelOptions{
			CustomizationID: createVoiceModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)

	getVoiceModel := service.GetGetVoiceModelResult(response)
	assert.NotNil(t, getVoiceModel)

	customizationID = createVoiceModel.CustomizationID
}

func TestWords(t *testing.T) {
	shouldSkipTest(t)

	// List Words
	response, responseErr := service.ListWords(
		&texttospeechv1.ListWordsOptions{
			CustomizationID: customizationID,
		},
	)
	assert.Nil(t, responseErr)

	listWords := service.GetListWordsResult(response)
	assert.NotNil(t, listWords)

	// Add words
	response, responseErr = service.AddWords(
		&texttospeechv1.AddWordsOptions{
			CustomizationID: customizationID,
			Words: []texttospeechv1.Word{
				texttospeechv1.Word{
					Word:        core.StringPtr("EEE"),
					Translation: core.StringPtr("<phoneme alphabet=\"ibm\" ph=\"tr1Ipxl.1i\"></phoneme>"),
				},
				texttospeechv1.Word{
					Word:        core.StringPtr("IEEE"),
					Translation: core.StringPtr("<phoneme alphabet=\"ibm\" ph=\"1Y.tr1Ipxl.1i\"></phoneme>"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	// Add word
	response, responseErr = service.AddWord(
		&texttospeechv1.AddWordOptions{
			CustomizationID: customizationID,
			Word:            core.StringPtr("ACLs"),
			Translation:     core.StringPtr("ackles"),
		},
	)
	assert.Nil(t, responseErr)

	// Get word
	response, responseErr = service.GetWord(
		&texttospeechv1.GetWordOptions{
			CustomizationID: customizationID,
			Word:            core.StringPtr("ACLs"),
		},
	)
	assert.Nil(t, responseErr)

	getWord := service.GetGetWordResult(response)
	assert.NotNil(t, getWord)

	// Delete word
	response, responseErr = service.DeleteWord(
		&texttospeechv1.DeleteWordOptions{
			CustomizationID: customizationID,
			Word:            core.StringPtr("ACLs"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete voice model
	response, responseErr = service.DeleteVoiceModel(
		&texttospeechv1.DeleteVoiceModelOptions{
			CustomizationID: customizationID,
		},
	)
	assert.Nil(t, responseErr)
}
