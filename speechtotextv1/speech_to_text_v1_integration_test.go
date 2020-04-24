// +build integration

package speechtotextv1_test

/**
 * (C) Copyright IBM Corp. 2018, 2019.
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
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *speechtotextv1.SpeechToTextV1
var languageModel *speechtotextv1.LanguageModel

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

	service, err = speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestModel(t *testing.T) {
	shouldSkipTest(t)

	// List models
	listModels, _, responseErr := service.ListModels(
		&speechtotextv1.ListModelsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listModels)

	//Get model
	getModel, _, responseErr := service.GetModel(
		&speechtotextv1.GetModelOptions{
			ModelID: core.StringPtr("en-US_BroadbandModel"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getModel)
}

func TestRecognize(t *testing.T) {
	shouldSkipTest(t)

	files := [1]string{"audio_example.mp3"}
	for _, fileName := range files {
		var audio io.ReadCloser
		var audioErr error
		audio, audioErr = os.Open("../resources/" + fileName)
		assert.Nil(t, audioErr)

		recognize, _, responseErr := service.Recognize(
			&speechtotextv1.RecognizeOptions{
				Audio:                     audio,
				ContentType:               core.StringPtr("audio/mp3"),
				Timestamps:                core.BoolPtr(true),
				WordAlternativesThreshold: core.Float32Ptr(0.9),
				Keywords:                  []string{"colorado", "tornado", "tornadoes"},
				KeywordsThreshold:         core.Float32Ptr(0.5),
			},
		)
		assert.Nil(t, responseErr)
		assert.NotNil(t, recognize)
	}
}

func TestJobs(t *testing.T) {
	shouldSkipTest(t)

	t.Skip("Skipping time consuming test")
	checkJobs, _, responseErr := service.CheckJobs(
		&speechtotextv1.CheckJobsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, checkJobs)
}

func TestLanguageModel(t *testing.T) {
	shouldSkipTest(t)

	// create language model
	createLanguageModel, _, responseErr := service.CreateLanguageModel(
		&speechtotextv1.CreateLanguageModelOptions{
			Name:          core.StringPtr("First example language model for GO"),
			BaseModelName: core.StringPtr(speechtotextv1.CreateLanguageModelOptions_BaseModelName_EnUsBroadbandmodel),
			Description:   core.StringPtr("First custom language model example"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createLanguageModel)

	// List language model
	listLanguageModel, _, responseErr := service.ListLanguageModels(
		&speechtotextv1.ListLanguageModelsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listLanguageModel)

	// Get language model
	getLanguageModel, _, responseErr := service.GetLanguageModel(
		&speechtotextv1.GetLanguageModelOptions{
			CustomizationID: createLanguageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getLanguageModel)

	// store in global variable
	languageModel = getLanguageModel
}

func TestCorpora(t *testing.T) {
	shouldSkipTest(t)

	// List corpora
	listCorpora, _, responseErr := service.ListCorpora(
		&speechtotextv1.ListCorporaOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listCorpora)

	// Add corpora
	if *languageModel.Status != "Available" {
		t.Skip("Skipping the rest of the corpora tests")
	}

	corpusFile, corpusFileErr := os.Open("../resources/corpus-short-1.txt")
	if corpusFileErr != nil {
		panic(corpusFileErr)
	}
	_, responseErr = service.AddCorpus(
		&speechtotextv1.AddCorpusOptions{
			CustomizationID: languageModel.CustomizationID,
			CorpusName:      core.StringPtr("corpus for GO"),
			CorpusFile:      corpusFile,
		},
	)
	assert.Nil(t, responseErr)

	// Get corpus
	getLanguageModel, _, responseErr := service.GetCorpus(
		&speechtotextv1.GetCorpusOptions{
			CustomizationID: languageModel.CustomizationID,
			CorpusName:      core.StringPtr("corpus for GO"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getLanguageModel)

	// Delete corpus
	_, responseErr = service.DeleteCorpus(
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
	listWords, _, responseErr := service.ListWords(
		&speechtotextv1.ListWordsOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listWords)

	if *languageModel.Status != "Available" {
		t.Skip("Skipping the rest of the words tests")
	}

	// Add words
	_, responseErr = service.AddWords(
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
	_, responseErr = service.AddWord(
		&speechtotextv1.AddWordOptions{
			CustomizationID: languageModel.CustomizationID,
			WordName:        core.StringPtr("NCAA"),
			SoundsLike:      []string{"N. C. A. A.", "N. C. double A."},
			DisplayAs:       core.StringPtr("NCAA"),
		},
	)
	assert.Nil(t, responseErr)

	// Get word
	getWord, _, responseErr := service.GetWord(
		&speechtotextv1.GetWordOptions{
			CustomizationID: languageModel.CustomizationID,
			WordName:        core.StringPtr("NCAA"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getWord)

	// Delete word
	_, responseErr = service.DeleteWord(
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
	createAcousticModel, _, responseErr := service.CreateAcousticModel(
		&speechtotextv1.CreateAcousticModelOptions{
			Name:          core.StringPtr("First example acoustic model for GO"),
			BaseModelName: core.StringPtr(speechtotextv1.CreateAcousticModelOptions_BaseModelName_EnUsBroadbandmodel),
			Description:   core.StringPtr("First custom acoustic model example"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createAcousticModel)

	// List acoustic model
	listAcousticModel, _, responseErr := service.ListAcousticModels(
		&speechtotextv1.ListAcousticModelsOptions{
			Language: core.StringPtr("en-US"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listAcousticModel)

	// Get acoustic model
	getAcousticModel, _, responseErr := service.GetAcousticModel(
		&speechtotextv1.GetAcousticModelOptions{
			CustomizationID: createAcousticModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getAcousticModel)

	// Delete acoustic model
	_, responseErr = service.DeleteAcousticModel(
		&speechtotextv1.DeleteAcousticModelOptions{
			CustomizationID: createAcousticModel.CustomizationID,
		},
	)
	assert.Nil(t, responseErr)

	t.Skip("Skip upgrade acoustic model")
	_, responseErr = service.UpgradeAcousticModel(
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
	listAudio, _, responseErr := service.ListAudio(
		&speechtotextv1.ListAudioOptions{
			CustomizationID: languageModel.CustomizationID,
		},
	)
	assert.NotNil(t, listAudio)

	// Add audio
	pwd, _ := os.Getwd()

	var audioFile io.ReadCloser
	var audioFileErr error
	audioFile, audioFileErr = os.Open(pwd + "/../resources/output.wav")
	if audioFileErr != nil {
		panic(audioFileErr)
	}
	_, responseErr = service.AddAudio(
		&speechtotextv1.AddAudioOptions{
			CustomizationID: languageModel.CustomizationID,
			AudioName:       core.StringPtr("audio1"),
			AudioResource:   audioFile,
			ContentType:     core.StringPtr("audio/wav"),
		},
	)
	assert.Nil(t, responseErr)

	// Get audio
	getAudio, _, responseErr := service.GetAudio(
		&speechtotextv1.GetAudioOptions{
			CustomizationID: languageModel.CustomizationID,
			AudioName:       core.StringPtr("audio1"),
		},
	)
	assert.NotNil(t, getAudio)

	// Delete audio
	_, responseErr = service.DeleteAudio(
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

func TestRecognizeUsingWebsocket(t *testing.T) {
	shouldSkipTest(t)
	f, _ := os.Open("../resources/audio_example.mp3")
	callback := myCallBack{T: t}

	recognizeOptions := service.NewRecognizeUsingWebsocketOptions(f, "audio/mp3")

	recognizeOptions.SetModel("en-US_BroadbandModel").SetWordConfidence(true).SetSpeakerLabels(true).SetTimestamps(true)

	service.RecognizeUsingWebsocket(recognizeOptions, callback)

}
