//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2021, 2022.
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

package speechtotextv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/speechtotextv1"
)

//
// This file provides an example of how to use the Speech to Text service.
//
// The following configuration properties are assumed to be defined:
// SPEECH_TO_TEXT_URL=<service base url>
// SPEECH_TO_TEXT_AUTH_TYPE=iam
// SPEECH_TO_TEXT_APIKEY=<IAM apikey>
// SPEECH_TO_TEXT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../speech_to_text_v1.env"

var (
	speechToTextService *speechtotextv1.SpeechToTextV1
	config              map[string]string
	configLoaded        bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`SpeechToTextV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(speechtotextv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			speechToTextServiceOptions := &speechtotextv1.SpeechToTextV1Options{}

			speechToTextService, err = speechtotextv1.NewSpeechToTextV1(speechToTextServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(speechToTextService).ToNot(BeNil())
		})
	})

	Describe(`SpeechToTextV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListModels request example`, func() {
			fmt.Println("\nListModels() result:")
			// begin-listModels

			listModelsOptions := speechToTextService.NewListModelsOptions()

			speechModels, response, err := speechToTextService.ListModels(listModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(speechModels, "", "  ")
			fmt.Println(string(b))

			// end-listModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(speechModels).ToNot(BeNil())

		})
		It(`GetModel request example`, func() {
			fmt.Println("\nGetModel() result:")
			// begin-getModel

			getModelOptions := speechToTextService.NewGetModelOptions(
				"ar-AR_BroadbandModel",
			)

			speechModel, response, err := speechToTextService.GetModel(getModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(speechModel, "", "  ")
			fmt.Println(string(b))

			// end-getModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(speechModel).ToNot(BeNil())

		})
		It(`Recognize request example`, func() {
			fmt.Println("\nRecognize() result:")
			// begin-recognize

			recognizeOptions := speechToTextService.NewRecognizeOptions(
				CreateMockReader("This is a mock file."),
			)

			speechRecognitionResults, response, err := speechToTextService.Recognize(recognizeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(speechRecognitionResults, "", "  ")
			fmt.Println(string(b))

			// end-recognize

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(speechRecognitionResults).ToNot(BeNil())

		})
		It(`RegisterCallback request example`, func() {
			fmt.Println("\nRegisterCallback() result:")
			// begin-registerCallback

			registerCallbackOptions := speechToTextService.NewRegisterCallbackOptions(
				"testString",
			)

			registerStatus, response, err := speechToTextService.RegisterCallback(registerCallbackOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registerStatus, "", "  ")
			fmt.Println(string(b))

			// end-registerCallback

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registerStatus).ToNot(BeNil())

		})
		It(`UnregisterCallback request example`, func() {
			// begin-unregisterCallback

			unregisterCallbackOptions := speechToTextService.NewUnregisterCallbackOptions(
				"testString",
			)

			response, err := speechToTextService.UnregisterCallback(unregisterCallbackOptions)
			if err != nil {
				panic(err)
			}

			// end-unregisterCallback
			fmt.Printf("\nUnregisterCallback() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`CreateJob request example`, func() {
			fmt.Println("\nCreateJob() result:")
			// begin-createJob

			createJobOptions := speechToTextService.NewCreateJobOptions(
				CreateMockReader("This is a mock file."),
			)

			recognitionJob, response, err := speechToTextService.CreateJob(createJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(recognitionJob, "", "  ")
			fmt.Println(string(b))

			// end-createJob

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(recognitionJob).ToNot(BeNil())

		})
		It(`CheckJobs request example`, func() {
			fmt.Println("\nCheckJobs() result:")
			// begin-checkJobs

			checkJobsOptions := speechToTextService.NewCheckJobsOptions()

			recognitionJobs, response, err := speechToTextService.CheckJobs(checkJobsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(recognitionJobs, "", "  ")
			fmt.Println(string(b))

			// end-checkJobs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(recognitionJobs).ToNot(BeNil())

		})
		It(`CheckJob request example`, func() {
			fmt.Println("\nCheckJob() result:")
			// begin-checkJob

			checkJobOptions := speechToTextService.NewCheckJobOptions(
				"testString",
			)

			recognitionJob, response, err := speechToTextService.CheckJob(checkJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(recognitionJob, "", "  ")
			fmt.Println(string(b))

			// end-checkJob

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(recognitionJob).ToNot(BeNil())

		})
		It(`CreateLanguageModel request example`, func() {
			fmt.Println("\nCreateLanguageModel() result:")
			// begin-createLanguageModel

			createLanguageModelOptions := speechToTextService.NewCreateLanguageModelOptions(
				"testString",
				"ar-MS_Telephony",
			)

			languageModel, response, err := speechToTextService.CreateLanguageModel(createLanguageModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(languageModel, "", "  ")
			fmt.Println(string(b))

			// end-createLanguageModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(languageModel).ToNot(BeNil())

		})
		It(`ListLanguageModels request example`, func() {
			fmt.Println("\nListLanguageModels() result:")
			// begin-listLanguageModels

			listLanguageModelsOptions := speechToTextService.NewListLanguageModelsOptions()

			languageModels, response, err := speechToTextService.ListLanguageModels(listLanguageModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(languageModels, "", "  ")
			fmt.Println(string(b))

			// end-listLanguageModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(languageModels).ToNot(BeNil())

		})
		It(`GetLanguageModel request example`, func() {
			fmt.Println("\nGetLanguageModel() result:")
			// begin-getLanguageModel

			getLanguageModelOptions := speechToTextService.NewGetLanguageModelOptions(
				"testString",
			)

			languageModel, response, err := speechToTextService.GetLanguageModel(getLanguageModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(languageModel, "", "  ")
			fmt.Println(string(b))

			// end-getLanguageModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(languageModel).ToNot(BeNil())

		})
		It(`TrainLanguageModel request example`, func() {
			fmt.Println("\nTrainLanguageModel() result:")
			// begin-trainLanguageModel

			trainLanguageModelOptions := speechToTextService.NewTrainLanguageModelOptions(
				"testString",
			)

			trainingResponse, response, err := speechToTextService.TrainLanguageModel(trainLanguageModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingResponse, "", "  ")
			fmt.Println(string(b))

			// end-trainLanguageModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingResponse).ToNot(BeNil())

		})
		It(`ResetLanguageModel request example`, func() {
			// begin-resetLanguageModel

			resetLanguageModelOptions := speechToTextService.NewResetLanguageModelOptions(
				"testString",
			)

			response, err := speechToTextService.ResetLanguageModel(resetLanguageModelOptions)
			if err != nil {
				panic(err)
			}

			// end-resetLanguageModel
			fmt.Printf("\nResetLanguageModel() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`UpgradeLanguageModel request example`, func() {
			// begin-upgradeLanguageModel

			upgradeLanguageModelOptions := speechToTextService.NewUpgradeLanguageModelOptions(
				"testString",
			)

			response, err := speechToTextService.UpgradeLanguageModel(upgradeLanguageModelOptions)
			if err != nil {
				panic(err)
			}

			// end-upgradeLanguageModel
			fmt.Printf("\nUpgradeLanguageModel() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`ListCorpora request example`, func() {
			fmt.Println("\nListCorpora() result:")
			// begin-listCorpora

			listCorporaOptions := speechToTextService.NewListCorporaOptions(
				"testString",
			)

			corpora, response, err := speechToTextService.ListCorpora(listCorporaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(corpora, "", "  ")
			fmt.Println(string(b))

			// end-listCorpora

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(corpora).ToNot(BeNil())

		})
		It(`AddCorpus request example`, func() {
			// begin-addCorpus

			addCorpusOptions := speechToTextService.NewAddCorpusOptions(
				"testString",
				"testString",
				CreateMockReader("This is a mock file."),
			)

			response, err := speechToTextService.AddCorpus(addCorpusOptions)
			if err != nil {
				panic(err)
			}

			// end-addCorpus
			fmt.Printf("\nAddCorpus() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
		It(`GetCorpus request example`, func() {
			fmt.Println("\nGetCorpus() result:")
			// begin-getCorpus

			getCorpusOptions := speechToTextService.NewGetCorpusOptions(
				"testString",
				"testString",
			)

			corpus, response, err := speechToTextService.GetCorpus(getCorpusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(corpus, "", "  ")
			fmt.Println(string(b))

			// end-getCorpus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(corpus).ToNot(BeNil())

		})
		It(`ListWords request example`, func() {
			fmt.Println("\nListWords() result:")
			// begin-listWords

			listWordsOptions := speechToTextService.NewListWordsOptions(
				"testString",
			)

			words, response, err := speechToTextService.ListWords(listWordsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(words, "", "  ")
			fmt.Println(string(b))

			// end-listWords

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(words).ToNot(BeNil())

		})
		It(`AddWords request example`, func() {
			// begin-addWords

			customWordModel := &speechtotextv1.CustomWord{}

			addWordsOptions := speechToTextService.NewAddWordsOptions(
				"testString",
				[]speechtotextv1.CustomWord{*customWordModel},
			)

			response, err := speechToTextService.AddWords(addWordsOptions)
			if err != nil {
				panic(err)
			}

			// end-addWords
			fmt.Printf("\nAddWords() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
		It(`AddWord request example`, func() {
			// begin-addWord

			addWordOptions := speechToTextService.NewAddWordOptions(
				"testString",
				"testString",
			)

			response, err := speechToTextService.AddWord(addWordOptions)
			if err != nil {
				panic(err)
			}

			// end-addWord
			fmt.Printf("\nAddWord() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
		It(`GetWord request example`, func() {
			fmt.Println("\nGetWord() result:")
			// begin-getWord

			getWordOptions := speechToTextService.NewGetWordOptions(
				"testString",
				"testString",
			)

			word, response, err := speechToTextService.GetWord(getWordOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(word, "", "  ")
			fmt.Println(string(b))

			// end-getWord

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(word).ToNot(BeNil())

		})
		It(`ListGrammars request example`, func() {
			fmt.Println("\nListGrammars() result:")
			// begin-listGrammars

			listGrammarsOptions := speechToTextService.NewListGrammarsOptions(
				"testString",
			)

			grammars, response, err := speechToTextService.ListGrammars(listGrammarsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(grammars, "", "  ")
			fmt.Println(string(b))

			// end-listGrammars

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(grammars).ToNot(BeNil())

		})
		It(`AddGrammar request example`, func() {
			// begin-addGrammar

			addGrammarOptions := speechToTextService.NewAddGrammarOptions(
				"testString",
				"testString",
				CreateMockReader("This is a mock file."),
				"application/srgs",
			)

			response, err := speechToTextService.AddGrammar(addGrammarOptions)
			if err != nil {
				panic(err)
			}

			// end-addGrammar
			fmt.Printf("\nAddGrammar() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
		It(`GetGrammar request example`, func() {
			fmt.Println("\nGetGrammar() result:")
			// begin-getGrammar

			getGrammarOptions := speechToTextService.NewGetGrammarOptions(
				"testString",
				"testString",
			)

			grammar, response, err := speechToTextService.GetGrammar(getGrammarOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(grammar, "", "  ")
			fmt.Println(string(b))

			// end-getGrammar

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(grammar).ToNot(BeNil())

		})
		It(`CreateAcousticModel request example`, func() {
			fmt.Println("\nCreateAcousticModel() result:")
			// begin-createAcousticModel

			createAcousticModelOptions := speechToTextService.NewCreateAcousticModelOptions(
				"testString",
				"ar-AR_BroadbandModel",
			)

			acousticModel, response, err := speechToTextService.CreateAcousticModel(createAcousticModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(acousticModel, "", "  ")
			fmt.Println(string(b))

			// end-createAcousticModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(acousticModel).ToNot(BeNil())

		})
		It(`ListAcousticModels request example`, func() {
			fmt.Println("\nListAcousticModels() result:")
			// begin-listAcousticModels

			listAcousticModelsOptions := speechToTextService.NewListAcousticModelsOptions()

			acousticModels, response, err := speechToTextService.ListAcousticModels(listAcousticModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(acousticModels, "", "  ")
			fmt.Println(string(b))

			// end-listAcousticModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(acousticModels).ToNot(BeNil())

		})
		It(`GetAcousticModel request example`, func() {
			fmt.Println("\nGetAcousticModel() result:")
			// begin-getAcousticModel

			getAcousticModelOptions := speechToTextService.NewGetAcousticModelOptions(
				"testString",
			)

			acousticModel, response, err := speechToTextService.GetAcousticModel(getAcousticModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(acousticModel, "", "  ")
			fmt.Println(string(b))

			// end-getAcousticModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(acousticModel).ToNot(BeNil())

		})
		It(`TrainAcousticModel request example`, func() {
			fmt.Println("\nTrainAcousticModel() result:")
			// begin-trainAcousticModel

			trainAcousticModelOptions := speechToTextService.NewTrainAcousticModelOptions(
				"testString",
			)

			trainingResponse, response, err := speechToTextService.TrainAcousticModel(trainAcousticModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingResponse, "", "  ")
			fmt.Println(string(b))

			// end-trainAcousticModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingResponse).ToNot(BeNil())

		})
		It(`ResetAcousticModel request example`, func() {
			// begin-resetAcousticModel

			resetAcousticModelOptions := speechToTextService.NewResetAcousticModelOptions(
				"testString",
			)

			response, err := speechToTextService.ResetAcousticModel(resetAcousticModelOptions)
			if err != nil {
				panic(err)
			}

			// end-resetAcousticModel
			fmt.Printf("\nResetAcousticModel() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`UpgradeAcousticModel request example`, func() {
			// begin-upgradeAcousticModel

			upgradeAcousticModelOptions := speechToTextService.NewUpgradeAcousticModelOptions(
				"testString",
			)

			response, err := speechToTextService.UpgradeAcousticModel(upgradeAcousticModelOptions)
			if err != nil {
				panic(err)
			}

			// end-upgradeAcousticModel
			fmt.Printf("\nUpgradeAcousticModel() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`ListAudio request example`, func() {
			fmt.Println("\nListAudio() result:")
			// begin-listAudio

			listAudioOptions := speechToTextService.NewListAudioOptions(
				"testString",
			)

			audioResources, response, err := speechToTextService.ListAudio(listAudioOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(audioResources, "", "  ")
			fmt.Println(string(b))

			// end-listAudio

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(audioResources).ToNot(BeNil())

		})
		It(`AddAudio request example`, func() {
			// begin-addAudio

			addAudioOptions := speechToTextService.NewAddAudioOptions(
				"testString",
				"testString",
				CreateMockReader("This is a mock file."),
			)

			response, err := speechToTextService.AddAudio(addAudioOptions)
			if err != nil {
				panic(err)
			}

			// end-addAudio
			fmt.Printf("\nAddAudio() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
		It(`GetAudio request example`, func() {
			fmt.Println("\nGetAudio() result:")
			// begin-getAudio

			getAudioOptions := speechToTextService.NewGetAudioOptions(
				"testString",
				"testString",
			)

			audioListing, response, err := speechToTextService.GetAudio(getAudioOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(audioListing, "", "  ")
			fmt.Println(string(b))

			// end-getAudio

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(audioListing).ToNot(BeNil())

		})
		It(`DeleteWord request example`, func() {
			// begin-deleteWord

			deleteWordOptions := speechToTextService.NewDeleteWordOptions(
				"testString",
				"testString",
			)

			response, err := speechToTextService.DeleteWord(deleteWordOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteWord
			fmt.Printf("\nDeleteWord() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteUserData request example`, func() {
			// begin-deleteUserData

			deleteUserDataOptions := speechToTextService.NewDeleteUserDataOptions(
				"testString",
			)

			response, err := speechToTextService.DeleteUserData(deleteUserDataOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteUserData
			fmt.Printf("\nDeleteUserData() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteLanguageModel request example`, func() {
			// begin-deleteLanguageModel

			deleteLanguageModelOptions := speechToTextService.NewDeleteLanguageModelOptions(
				"testString",
			)

			response, err := speechToTextService.DeleteLanguageModel(deleteLanguageModelOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteLanguageModel
			fmt.Printf("\nDeleteLanguageModel() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteJob request example`, func() {
			// begin-deleteJob

			deleteJobOptions := speechToTextService.NewDeleteJobOptions(
				"testString",
			)

			response, err := speechToTextService.DeleteJob(deleteJobOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteJob
			fmt.Printf("\nDeleteJob() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteGrammar request example`, func() {
			// begin-deleteGrammar

			deleteGrammarOptions := speechToTextService.NewDeleteGrammarOptions(
				"testString",
				"testString",
			)

			response, err := speechToTextService.DeleteGrammar(deleteGrammarOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteGrammar
			fmt.Printf("\nDeleteGrammar() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteCorpus request example`, func() {
			// begin-deleteCorpus

			deleteCorpusOptions := speechToTextService.NewDeleteCorpusOptions(
				"testString",
				"testString",
			)

			response, err := speechToTextService.DeleteCorpus(deleteCorpusOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteCorpus
			fmt.Printf("\nDeleteCorpus() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteAudio request example`, func() {
			// begin-deleteAudio

			deleteAudioOptions := speechToTextService.NewDeleteAudioOptions(
				"testString",
				"testString",
			)

			response, err := speechToTextService.DeleteAudio(deleteAudioOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteAudio
			fmt.Printf("\nDeleteAudio() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteAcousticModel request example`, func() {
			// begin-deleteAcousticModel

			deleteAcousticModelOptions := speechToTextService.NewDeleteAcousticModelOptions(
				"testString",
			)

			response, err := speechToTextService.DeleteAcousticModel(deleteAcousticModelOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteAcousticModel
			fmt.Printf("\nDeleteAcousticModel() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
