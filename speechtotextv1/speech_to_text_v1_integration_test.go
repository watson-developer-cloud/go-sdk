// +build integration

package speechtotextv1_test

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
	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"io"
	"os"
	"testing"
)

var service *speechtotextv1.SpeechToTextV1
var serviceErr error
var languageModel *speechtotextv1.LanguageModel

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = speechtotextv1.
			NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL:      os.Getenv("SPEECH_TO_TEXT_URL"),
				Username: os.Getenv("SPEECH_TO_TEXT_USERNAME"),
				Password: os.Getenv("SPEECH_TO_TEXT_PASSWORD"),
			})
	}
}
func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}
func TestModel(t *testing.T) {
	shouldSkipTest(t)

	// List models
	response, responseErr := service.ListModels(
		&speechtotextv1.ListModelsOptions{},
	)
	assert.Nil(t, responseErr)

	listModels := service.GetListModelsResult(response)
	assert.NotNil(t, listModels)

	//Get model
	response, responseErr = service.GetModel(
		&speechtotextv1.GetModelOptions{
			ModelID: core.StringPtr("en-US_BroadbandModel"),
		},
	)
	assert.Nil(t, responseErr)

	getModel := service.GetGetModelResult(response)
	assert.NotNil(t, getModel)
}

func TestRecognize(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()
	files := [1]string{"audio_example.mp3"}
	for _, fileName := range files {
		var audio io.ReadCloser
		var audioErr error
		audio, audioErr = os.Open(pwd + "/../resources/" + fileName)
		assert.Nil(t, audioErr)

		response, responseErr := service.Recognize(
			&speechtotextv1.RecognizeOptions{
				Audio:                     &audio,
				ContentType:               core.StringPtr(speechtotextv1.RecognizeOptions_ContentType_AudioMp3),
				Timestamps:                core.BoolPtr(true),
				WordAlternativesThreshold: core.Float32Ptr(0.9),
				Keywords:                  []string{"colorado", "tornado", "tornadoes"},
				KeywordsThreshold:         core.Float32Ptr(0.5),
			},
		)
		assert.Nil(t, responseErr)
		recognize := service.GetRecognizeResult(response)
		assert.NotNil(t, recognize)
	}
}

func TestJobs(t *testing.T) {
	shouldSkipTest(t)

	t.Skip("Skipping time consuming test")
	response, responseErr := service.CheckJobs(
		&speechtotextv1.CheckJobsOptions{},
	)
	assert.Nil(t, responseErr)

	checkJobs := service.GetCheckJobsResult(response)
	assert.NotNil(t, checkJobs)
}

func TestLanguageModel(t *testing.T) {
	shouldSkipTest(t)

	// create language model
	response, responseErr := service.CreateLanguageModel(
		&speechtotextv1.CreateLanguageModelOptions{
			Name:          core.StringPtr("First example language model for GO"),
			BaseModelName: core.StringPtr(speechtotextv1.CreateLanguageModelOptions_BaseModelName_EnUsBroadbandmodel),
			Description:   core.StringPtr("First custom language model example"),
		},
	)
	assert.Nil(t, responseErr)
	createLanguageModel := service.GetCreateLanguageModelResult(response)
	assert.NotNil(t, createLanguageModel)

	// List language model
	response, responseErr = service.ListLanguageModels(
		&speechtotextv1.ListLanguageModelsOptions{},
	)
	assert.Nil(t, responseErr)

	listLanguageModel := service.GetListLanguageModelsResult(response)
	assert.NotNil(t, listLanguageModel)

	// Get language model
	response, responseErr = service.GetLanguageModel(
		&speechtotextv1.GetLanguageModelOptions{
			CustomizationID: createLanguageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
	getLanguageModel := service.GetGetLanguageModelResult(response)
	assert.NotNil(t, getLanguageModel)

	// store in global variable
	languageModel = getLanguageModel
}

func TestCorpora(t *testing.T) {
	shouldSkipTest(t)

	// List corpora
	response, responseErr := service.ListCorpora(
		&speechtotextv1.ListCorporaOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)

	listCorpora := service.GetListCorporaResult(response)
	assert.NotNil(t, listCorpora)

	// Add corpora
	if *languageModel.Status != "Available" {
		t.Skip("Skipping the rest of the corpora tests")
	}

	pwd, _ := os.Getwd()
	corpusFile, corpusFileErr := os.Open(pwd + "/../resources/corpus-short-1.txt")
	if corpusFileErr != nil {
		panic(corpusFileErr)
	}
	response, responseErr = service.AddCorpus(
		&speechtotextv1.AddCorpusOptions{
			CustomizationID: languageModel.CustomizationID,
			CorpusName:      core.StringPtr("corpus for GO"),
			CorpusFile:      corpusFile,
		},
	)
	assert.Nil(t, responseErr)

	// Get corpus
	response, responseErr = service.GetCorpus(
		&speechtotextv1.GetCorpusOptions{
			CustomizationID: languageModel.CustomizationID,
			CorpusName:      core.StringPtr("corpus for GO"),
		},
	)
	assert.Nil(t, responseErr)

	getLanguageModel := service.GetGetCorpusResult(response)
	assert.NotNil(t, getLanguageModel)

	// Delete corpus
	response, responseErr = service.DeleteCorpus(
		&speechtotextv1.DeleteCorpusOptions{
			CustomizationID: languageModel.CustomizationID,
			CorpusName:      core.StringPtr("corpus for GO"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestWords(t *testing.T) {
	shouldSkipTest(t)

	// List words
	response, responseErr := service.ListWords(
		&speechtotextv1.ListWordsOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)

	listWords := service.GetListWordsResult(response)
	assert.NotNil(t, listWords)

	if *languageModel.Status != "Available" {
		t.Skip("Skipping the rest of the words tests")
	}

	// Add words
	response, responseErr = service.AddWords(
		&speechtotextv1.AddWordsOptions{
			CustomizationID: languageModel.CustomizationID,
			Words: []speechtotextv1.CustomWord{
				speechtotextv1.CustomWord{
					Word:       core.StringPtr("HHonors"),
					SoundsLike: []string{"hilton honors", "H. honors"},
					DisplayAs:  core.StringPtr("HHonors"),
				},
				speechtotextv1.CustomWord{
					Word:       core.StringPtr("IEEE"),
					SoundsLike: []string{"I. triple E."},
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	// Add word
	response, responseErr = service.AddWord(
		&speechtotextv1.AddWordOptions{
			CustomizationID: languageModel.CustomizationID,
			WordName:        core.StringPtr("NCAA"),
			SoundsLike:      []string{"N. C. A. A.", "N. C. double A."},
			DisplayAs:       core.StringPtr("NCAA"),
		},
	)
	assert.Nil(t, responseErr)

	// Get word
	response, responseErr = service.GetWord(
		&speechtotextv1.GetWordOptions{
			CustomizationID: languageModel.CustomizationID,
			WordName:        core.StringPtr("NCAA"),
		},
	)
	assert.Nil(t, responseErr)

	getWord := service.GetGetWordResult(response)
	assert.NotNil(t, getWord)

	// Delete word
	response, responseErr = service.DeleteWord(
		&speechtotextv1.DeleteWordOptions{
			CustomizationID: languageModel.CustomizationID,
			WordName:        core.StringPtr("NCAA"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestAcousticModel(t *testing.T) {
	shouldSkipTest(t)

	// create acoustic model
	response, responseErr := service.CreateAcousticModel(
		&speechtotextv1.CreateAcousticModelOptions{
			Name:          core.StringPtr("First example acoustic model for GO"),
			BaseModelName: core.StringPtr(speechtotextv1.CreateAcousticModelOptions_BaseModelName_EnUsBroadbandmodel),
			Description:   core.StringPtr("First custom acoustic model example"),
		},
	)
	assert.Nil(t, responseErr)

	createAcousticModel := service.GetCreateAcousticModelResult(response)
	assert.NotNil(t, createAcousticModel)

	// List acoustic model
	response, responseErr = service.ListAcousticModels(
		&speechtotextv1.ListAcousticModelsOptions{
			Language: core.StringPtr("en-US"),
		},
	)
	assert.Nil(t, responseErr)

	listAcousticModel := service.GetListAcousticModelsResult(response)
	assert.NotNil(t, listAcousticModel)

	// Get acoustic model
	response, responseErr = service.GetAcousticModel(
		&speechtotextv1.GetAcousticModelOptions{
			CustomizationID: createAcousticModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)

	getAcousticModel := service.GetGetAcousticModelResult(response)
	assert.NotNil(t, getAcousticModel)

	// Delete acoustic model
	response, responseErr = service.DeleteAcousticModel(
		&speechtotextv1.DeleteAcousticModelOptions{
			CustomizationID: createAcousticModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)

	t.Skip("Skip upgrade acoustic model")
	response, responseErr = service.UpgradeAcousticModel(
		&speechtotextv1.UpgradeAcousticModelOptions{
			CustomizationID: createAcousticModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestAudio(t *testing.T) {
	shouldSkipTest(t)

	if *languageModel.Status != "Available" {
		t.Skip("Skipping the rest of the audio tests")
	}

	// List audio
	response, responseErr := service.ListAudio(
		&speechtotextv1.ListAudioOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)

	listAudio := service.GetListAudioResult(response)
	assert.NotNil(t, listAudio)

	// Add audio
	pwd, _ := os.Getwd()

	var audioFile io.ReadCloser
	var audioFileErr error
	audioFile, audioFileErr = os.Open(pwd + "/../resources/output.wav")
	if audioFileErr != nil {
		panic(audioFileErr)
	}
	response, responseErr = service.AddAudio(
		&speechtotextv1.AddAudioOptions{
			CustomizationID: languageModel.CustomizationID,
			AudioName:       core.StringPtr("audio1"),
			AudioResource:   &audioFile,
			ContentType:     core.StringPtr(speechtotextv1.AddAudioOptions_ContentType_AudioWav),
		},
	)
	assert.Nil(t, responseErr)

	// Get audio
	response, responseErr = service.GetAudio(
		&speechtotextv1.GetAudioOptions{
			CustomizationID: languageModel.CustomizationID,
			AudioName:       core.StringPtr("audio1"),
		},
	)

	getAudio := service.GetGetAudioResult(response)
	assert.NotNil(t, getAudio)

	// Delete audio
	response, responseErr = service.DeleteAudio(
		&speechtotextv1.DeleteAudioOptions{
			CustomizationID: languageModel.CustomizationID,
			AudioName:       core.StringPtr("audio1"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestDeleteLanguageModel(t *testing.T) {
	shouldSkipTest(t)

	// Delete language model
	_, responseErr := service.DeleteLanguageModel(
		&speechtotextv1.DeleteLanguageModelOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
}

type myCallBack struct {
	T *testing.T
}

func (cb myCallBack) OnOpen() {
}
func (cb myCallBack) OnClose() {
}
func (cb myCallBack) OnData(resp *core.DetailedResponse) {
	var speechResults speechtotextv1.SpeechRecognitionResults
	result := resp.GetResult().([]byte)
	json.Unmarshal(result, &speechResults)
	assert.NotNil(cb.T, speechResults)
	assert.NotNil(cb.T, speechResults.Results[0].Alternatives[0].WordConfidence)
	assert.NotNil(cb.T, speechResults.SpeakerLabels)
	assert.NotNil(cb.T, speechResults.Results[0].Alternatives[0].Timestamps)
}
func (cb myCallBack) OnError(err error) {
	cb.T.Fail()
}

func TestRecognizeUsingWebsockets(t *testing.T) {
	shouldSkipTest(t)
	f, _ := os.Open("../resources/audio_example.mp3")
	callback := myCallBack{T: t}

	recognizeOptions := service.NewRecognizeUsingWebsocketOptions(f, "audio/mp3")

	recognizeOptions.SetModel("en-US_BroadbandModel").SetWordConfidence(true).SetSpeakerLabels(true).SetTimestamps(true)

	service.RecognizeUsingWebsockets(recognizeOptions, callback)

}
