//go:build integration
// +build integration

package texttospeechv1_test

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
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/texttospeechv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../../.env"

var service *texttospeechv1.TextToSpeechV1
var customizationID *string

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

	service, err = texttospeechv1.NewTextToSpeechV1(
		&texttospeechv1.TextToSpeechV1Options{})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestVoice(t *testing.T) {
	shouldSkipTest(t)

	// list voices
	listVoices, _, responseErr := service.ListVoices(
		&texttospeechv1.ListVoicesOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listVoices)

	// Get voice
	voice, _, _ := service.GetVoice(
		&texttospeechv1.GetVoiceOptions{
			Voice: core.StringPtr(texttospeechv1.GetVoiceOptionsVoiceEnUsAllisonvoiceConst),
		},
	)
	assert.NotNil(t, voice)
}

func TestSynthesize(t *testing.T) {
	shouldSkipTest(t)

	// synthesize
	synthesize, _, responseErr := service.Synthesize(
		&texttospeechv1.SynthesizeOptions{
			Text:   core.StringPtr("Hello world"),
			Accept: core.StringPtr("audio/wav"),
			Voice:  core.StringPtr(texttospeechv1.SynthesizeOptionsVoiceEnUsAllisonvoiceConst),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, synthesize)
	synthesize.Close()
}

type myCallBack struct {
	T *testing.T
}

func (cb myCallBack) OnOpen() {}

func (cb myCallBack) OnClose() {}

func (cb myCallBack) OnAudioStream(b []byte) {}

func (cb myCallBack) OnError(err error) {
	cb.T.Fail()
}

func (cb myCallBack) OnTimingInformation(timings texttospeechv1.Timings) {
	assert.NotNil(cb.T, timings)
}

func (cb myCallBack) OnMarks(marks texttospeechv1.Marks) {
	assert.NotNil(cb.T, marks)
}

func (cb myCallBack) OnContentType(contentType string) {
	assert.NotNil(cb.T, contentType)
	assert.Equal(cb.T, contentType, "audio/mpeg")
}

func (cb myCallBack) OnData(resp *core.DetailedResponse) {
	result := resp.GetResult().([]byte)
	assert.NotNil(cb.T, result)
}

func TestSynthesizeUsingWebsocket(t *testing.T) {
	shouldSkipTest(t)

	callback := myCallBack{T: t}
	synthesizeUsingWebsocketOptions := service.
		NewSynthesizeUsingWebsocketOptions("This is a <mark name=\"SIMPLE\"/>simple <mark name=\"EXAMPLE\"/> example.", callback)

	synthesizeUsingWebsocketOptions.
		SetAccept("audio/mp3").
		SetVoice("en-US_AllisonVoice")
	synthesizeUsingWebsocketOptions.SetTimings([]string{"words"})
	err := service.SynthesizeUsingWebsocket(synthesizeUsingWebsocketOptions)
	assert.Nil(t, err)
}

func TestPronunciation(t *testing.T) {
	shouldSkipTest(t)

	// get pronunciation
	pronunciation, _, responseErr := service.GetPronunciation(
		&texttospeechv1.GetPronunciationOptions{
			Text:   core.StringPtr("IEEE"),
			Voice:  core.StringPtr(texttospeechv1.GetPronunciationOptionsVoiceEnUsAllisonvoiceConst),
			Format: core.StringPtr("ibm"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, pronunciation)
}

func TestVoiceModel(t *testing.T) {
	shouldSkipTest(t)

	// create voice model
	createVoiceModel, _, responseErr := service.CreateCustomModel(
		&texttospeechv1.CreateCustomModelOptions{
			Name:        core.StringPtr("First model for GO"),
			Language:    core.StringPtr(texttospeechv1.CreateCustomModelOptionsLanguageEnUsConst),
			Description: core.StringPtr("First custom voice model"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createVoiceModel)

	// List voice models
	listVoiceModels, _, responseErr := service.ListCustomModels(
		&texttospeechv1.ListCustomModelsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listVoiceModels)

	// Update voice model
	_, responseErr = service.UpdateCustomModel(
		&texttospeechv1.UpdateCustomModelOptions{
			CustomizationID: createVoiceModel.CustomizationID,
			Name:            core.StringPtr("First Model Update for GO"),
			Description:     core.StringPtr("First custom voice model update"),
			Words: []texttospeechv1.Word{
				{
					Word:        core.StringPtr("NCAA"),
					Translation: core.StringPtr("N C double A"),
				},
				{
					Word:        core.StringPtr("iPhone"),
					Translation: core.StringPtr("I phone"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	// Get voice model
	getVoiceModel, _, responseErr := service.GetCustomModel(
		&texttospeechv1.GetCustomModelOptions{
			CustomizationID: createVoiceModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getVoiceModel)

	customizationID = createVoiceModel.CustomizationID
}

func TestWords(t *testing.T) {
	shouldSkipTest(t)

	// List Words
	listWords, _, responseErr := service.ListWords(
		&texttospeechv1.ListWordsOptions{
			CustomizationID: customizationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listWords)

	// Add words
	_, responseErr = service.AddWords(
		&texttospeechv1.AddWordsOptions{
			CustomizationID: customizationID,
			Words: []texttospeechv1.Word{
				{
					Word:        core.StringPtr("EEE"),
					Translation: core.StringPtr("<phoneme alphabet=\"ibm\" ph=\"tr1Ipxl.1i\"></phoneme>"),
				},
				{
					Word:        core.StringPtr("IEEE"),
					Translation: core.StringPtr("<phoneme alphabet=\"ibm\" ph=\"1Y.tr1Ipxl.1i\"></phoneme>"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	// Add word
	_, responseErr = service.AddWord(
		&texttospeechv1.AddWordOptions{
			CustomizationID: customizationID,
			Word:            core.StringPtr("ACLs"),
			Translation:     core.StringPtr("ackles"),
		},
	)
	assert.Nil(t, responseErr)

	// Get word
	getWord, _, responseErr := service.GetWord(
		&texttospeechv1.GetWordOptions{
			CustomizationID: customizationID,
			Word:            core.StringPtr("ACLs"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getWord)

	// Delete word
	_, responseErr = service.DeleteWord(
		&texttospeechv1.DeleteWordOptions{
			CustomizationID: customizationID,
			Word:            core.StringPtr("ACLs"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete voice model
	_, responseErr = service.DeleteCustomModel(
		&texttospeechv1.DeleteCustomModelOptions{
			CustomizationID: customizationID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestCustomPromptsCRUD(t *testing.T) {
	shouldSkipTest(t)

	hashedConfigurationName := strconv.FormatInt(time.Now().UnixNano(), 10)

	createVoiceModel, _, responseErr := service.CreateCustomModel(
		&texttospeechv1.CreateCustomModelOptions{
			Name:        core.StringPtr("Go custom prompt model" + hashedConfigurationName),
			Language:    core.StringPtr(texttospeechv1.CreateCustomModelOptionsLanguageEnUsConst),
			Description: core.StringPtr("First custom voice model"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createVoiceModel)

	audio, audioErr := os.Open("../resources/tts_audio.wav")
	assert.Nil(t, audioErr)
	defer audio.Close()

	hashedPromptName := strconv.FormatInt(rand.Int63n(1000), 10)

	addedPrompt, _, addPromptErr := service.AddCustomPrompt(
		&texttospeechv1.AddCustomPromptOptions{
			CustomizationID: createVoiceModel.CustomizationID,
			PromptID:        core.StringPtr("gosdkprompt" + hashedPromptName),
			File:            audio,
			Metadata: &texttospeechv1.PromptMetadata{
				PromptText: core.StringPtr("Hello world"),
			},
		},
	)

	assert.Nil(t, addPromptErr)
	assert.NotNil(t, addedPrompt)

	readPrompt, _, readPromptErr := service.GetCustomPrompt(
		&texttospeechv1.GetCustomPromptOptions{
			CustomizationID: createVoiceModel.CustomizationID,
			PromptID:        addedPrompt.PromptID,
		},
	)

	assert.Nil(t, readPromptErr)
	assert.Equal(t, *addedPrompt.PromptID, *readPrompt.PromptID)

	listedPrompts, _, listPromptsErr := service.ListCustomPrompts(
		&texttospeechv1.ListCustomPromptsOptions{
			CustomizationID: createVoiceModel.CustomizationID,
		},
	)

	assert.Nil(t, listPromptsErr)
	assert.NotNil(t, listedPrompts)

	response, deletePromptErr := service.DeleteCustomPrompt(
		&texttospeechv1.DeleteCustomPromptOptions{
			CustomizationID: createVoiceModel.CustomizationID,
			PromptID:        addedPrompt.PromptID,
		},
	)

	assert.Nil(t, deletePromptErr)
	assert.NotNil(t, response)

	// Delete voice model
	_, responseErr = service.DeleteCustomModel(
		&texttospeechv1.DeleteCustomModelOptions{
			CustomizationID: createVoiceModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestSpeakerModelsCRUD(t *testing.T) {
	shouldSkipTest(t)

	audio, audioErr := os.Open("../resources/tts_audio.wav")
	assert.Nil(t, audioErr)
	defer audio.Close()

	hashedSpeakerName := strconv.FormatInt(rand.Int63n(1000), 10)

	createdModel, _, createModelErr := service.CreateSpeakerModel(
		&texttospeechv1.CreateSpeakerModelOptions{
			SpeakerName: core.StringPtr("gospeaker" + hashedSpeakerName),
			Audio:       audio,
		},
	)

	assert.Nil(t, createModelErr)
	assert.NotNil(t, createdModel)

	readModel, _, readModelErr := service.GetSpeakerModel(
		&texttospeechv1.GetSpeakerModelOptions{
			SpeakerID: createdModel.SpeakerID,
		},
	)

	assert.Nil(t, readModelErr)
	assert.NotNil(t, readModel)

	listedModels, _, listModelErr := service.ListSpeakerModels(
		&texttospeechv1.ListSpeakerModelsOptions{},
	)

	assert.Nil(t, listModelErr)
	assert.NotNil(t, listedModels)

	response, deleteModelErr := service.DeleteSpeakerModel(
		&texttospeechv1.DeleteSpeakerModelOptions{
			SpeakerID: createdModel.SpeakerID,
		},
	)

	assert.Nil(t, deleteModelErr)
	assert.NotNil(t, response)
}
