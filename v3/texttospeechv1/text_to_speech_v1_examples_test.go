//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
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

package texttospeechv1_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v3/texttospeechv1"
)

//
// This file provides an example of how to use the Text to Speech service.
//
// The following configuration properties are assumed to be defined:
// TEXT_TO_SPEECH_URL=<service base url>
// TEXT_TO_SPEECH_AUTH_TYPE=iam
// TEXT_TO_SPEECH_APIKEY=<IAM apikey>
// TEXT_TO_SPEECH_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`TextToSpeechV1 Examples Tests`, func() {

	const externalConfigFile = "../text_to_speech_v1.env"

	var (
		textToSpeechService *texttospeechv1.TextToSpeechV1
		config              map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(texttospeechv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			textToSpeechServiceOptions := &texttospeechv1.TextToSpeechV1Options{}

			textToSpeechService, err = texttospeechv1.NewTextToSpeechV1(textToSpeechServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(textToSpeechService).ToNot(BeNil())
		})
	})

	Describe(`TextToSpeechV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVoices request example`, func() {
			fmt.Println("\nListVoices() result:")
			// begin-listVoices

			listVoicesOptions := textToSpeechService.NewListVoicesOptions()

			voices, response, err := textToSpeechService.ListVoices(listVoicesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(voices, "", "  ")
			fmt.Println(string(b))

			// end-listVoices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(voices).ToNot(BeNil())

		})
		It(`GetVoice request example`, func() {
			fmt.Println("\nGetVoice() result:")
			// begin-getVoice

			getVoiceOptions := textToSpeechService.NewGetVoiceOptions(
				"ar-AR_OmarVoice",
			)

			voice, response, err := textToSpeechService.GetVoice(getVoiceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(voice, "", "  ")
			fmt.Println(string(b))

			// end-getVoice

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(voice).ToNot(BeNil())

		})
		It(`Synthesize request example`, func() {
			fmt.Println("\nSynthesize() result:")
			// begin-synthesize

			synthesizeOptions := textToSpeechService.NewSynthesizeOptions(
				"testString",
			)

			file, response, err := textToSpeechService.Synthesize(synthesizeOptions)
			if err != nil {
				panic(err)
			}
			if file != nil {
				defer file.Close()
				outFile, err := os.Create("file.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, file)
				if err != nil {
					panic(err)
				}
			}

			// end-synthesize

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(file).ToNot(BeNil())

		})
		It(`GetPronunciation request example`, func() {
			fmt.Println("\nGetPronunciation() result:")
			// begin-getPronunciation

			getPronunciationOptions := textToSpeechService.NewGetPronunciationOptions(
				"testString",
			)

			pronunciation, response, err := textToSpeechService.GetPronunciation(getPronunciationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pronunciation, "", "  ")
			fmt.Println(string(b))

			// end-getPronunciation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pronunciation).ToNot(BeNil())

		})
		It(`CreateCustomModel request example`, func() {
			fmt.Println("\nCreateCustomModel() result:")
			// begin-createCustomModel

			createCustomModelOptions := textToSpeechService.NewCreateCustomModelOptions(
				"testString",
			)

			customModel, response, err := textToSpeechService.CreateCustomModel(createCustomModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customModel, "", "  ")
			fmt.Println(string(b))

			// end-createCustomModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(customModel).ToNot(BeNil())

		})
		It(`ListCustomModels request example`, func() {
			fmt.Println("\nListCustomModels() result:")
			// begin-listCustomModels

			listCustomModelsOptions := textToSpeechService.NewListCustomModelsOptions()

			customModels, response, err := textToSpeechService.ListCustomModels(listCustomModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customModels, "", "  ")
			fmt.Println(string(b))

			// end-listCustomModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customModels).ToNot(BeNil())

		})
		It(`UpdateCustomModel request example`, func() {
			// begin-updateCustomModel

			updateCustomModelOptions := textToSpeechService.NewUpdateCustomModelOptions(
				"testString",
			)

			response, err := textToSpeechService.UpdateCustomModel(updateCustomModelOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from UpdateCustomModel(): %d\n", response.StatusCode)
			}

			// end-updateCustomModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetCustomModel request example`, func() {
			fmt.Println("\nGetCustomModel() result:")
			// begin-getCustomModel

			getCustomModelOptions := textToSpeechService.NewGetCustomModelOptions(
				"testString",
			)

			customModel, response, err := textToSpeechService.GetCustomModel(getCustomModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customModel, "", "  ")
			fmt.Println(string(b))

			// end-getCustomModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customModel).ToNot(BeNil())

		})
		It(`AddWords request example`, func() {
			// begin-addWords

			wordModel := &texttospeechv1.Word{
				Word:        core.StringPtr("testString"),
				Translation: core.StringPtr("testString"),
			}

			addWordsOptions := textToSpeechService.NewAddWordsOptions(
				"testString",
				[]texttospeechv1.Word{*wordModel},
			)

			response, err := textToSpeechService.AddWords(addWordsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from AddWords(): %d\n", response.StatusCode)
			}

			// end-addWords

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`ListWords request example`, func() {
			fmt.Println("\nListWords() result:")
			// begin-listWords

			listWordsOptions := textToSpeechService.NewListWordsOptions(
				"testString",
			)

			words, response, err := textToSpeechService.ListWords(listWordsOptions)
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
		It(`AddWord request example`, func() {
			// begin-addWord

			addWordOptions := textToSpeechService.NewAddWordOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := textToSpeechService.AddWord(addWordOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from AddWord(): %d\n", response.StatusCode)
			}

			// end-addWord

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`GetWord request example`, func() {
			fmt.Println("\nGetWord() result:")
			// begin-getWord

			getWordOptions := textToSpeechService.NewGetWordOptions(
				"testString",
				"testString",
			)

			translation, response, err := textToSpeechService.GetWord(getWordOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(translation, "", "  ")
			fmt.Println(string(b))

			// end-getWord

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(translation).ToNot(BeNil())

		})
		It(`ListCustomPrompts request example`, func() {
			fmt.Println("\nListCustomPrompts() result:")
			// begin-listCustomPrompts

			listCustomPromptsOptions := textToSpeechService.NewListCustomPromptsOptions(
				"testString",
			)

			prompts, response, err := textToSpeechService.ListCustomPrompts(listCustomPromptsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(prompts, "", "  ")
			fmt.Println(string(b))

			// end-listCustomPrompts

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prompts).ToNot(BeNil())

		})
		It(`AddCustomPrompt request example`, func() {
			fmt.Println("\nAddCustomPrompt() result:")
			// begin-addCustomPrompt

			promptMetadataModel := &texttospeechv1.PromptMetadata{
				PromptText: core.StringPtr("testString"),
			}

			addCustomPromptOptions := textToSpeechService.NewAddCustomPromptOptions(
				"testString",
				"testString",
				promptMetadataModel,
				CreateMockReader("This is a mock file."),
			)

			prompt, response, err := textToSpeechService.AddCustomPrompt(addCustomPromptOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(prompt, "", "  ")
			fmt.Println(string(b))

			// end-addCustomPrompt

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(prompt).ToNot(BeNil())

		})
		It(`GetCustomPrompt request example`, func() {
			fmt.Println("\nGetCustomPrompt() result:")
			// begin-getCustomPrompt

			getCustomPromptOptions := textToSpeechService.NewGetCustomPromptOptions(
				"testString",
				"testString",
			)

			prompt, response, err := textToSpeechService.GetCustomPrompt(getCustomPromptOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(prompt, "", "  ")
			fmt.Println(string(b))

			// end-getCustomPrompt

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prompt).ToNot(BeNil())

		})
		It(`ListSpeakerModels request example`, func() {
			fmt.Println("\nListSpeakerModels() result:")
			// begin-listSpeakerModels

			listSpeakerModelsOptions := textToSpeechService.NewListSpeakerModelsOptions()

			speakers, response, err := textToSpeechService.ListSpeakerModels(listSpeakerModelsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(speakers, "", "  ")
			fmt.Println(string(b))

			// end-listSpeakerModels

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(speakers).ToNot(BeNil())

		})
		It(`CreateSpeakerModel request example`, func() {
			fmt.Println("\nCreateSpeakerModel() result:")
			// begin-createSpeakerModel

			createSpeakerModelOptions := textToSpeechService.NewCreateSpeakerModelOptions(
				"testString",
				CreateMockReader("This is a mock file."),
			)

			speakerModel, response, err := textToSpeechService.CreateSpeakerModel(createSpeakerModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(speakerModel, "", "  ")
			fmt.Println(string(b))

			// end-createSpeakerModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(speakerModel).ToNot(BeNil())

		})
		It(`GetSpeakerModel request example`, func() {
			fmt.Println("\nGetSpeakerModel() result:")
			// begin-getSpeakerModel

			getSpeakerModelOptions := textToSpeechService.NewGetSpeakerModelOptions(
				"testString",
			)

			speakerCustomModels, response, err := textToSpeechService.GetSpeakerModel(getSpeakerModelOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(speakerCustomModels, "", "  ")
			fmt.Println(string(b))

			// end-getSpeakerModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(speakerCustomModels).ToNot(BeNil())

		})
		It(`DeleteWord request example`, func() {
			// begin-deleteWord

			deleteWordOptions := textToSpeechService.NewDeleteWordOptions(
				"testString",
				"testString",
			)

			response, err := textToSpeechService.DeleteWord(deleteWordOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteWord(): %d\n", response.StatusCode)
			}

			// end-deleteWord

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteUserData request example`, func() {
			// begin-deleteUserData

			deleteUserDataOptions := textToSpeechService.NewDeleteUserDataOptions(
				"testString",
			)

			response, err := textToSpeechService.DeleteUserData(deleteUserDataOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteUserData(): %d\n", response.StatusCode)
			}

			// end-deleteUserData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteSpeakerModel request example`, func() {
			// begin-deleteSpeakerModel

			deleteSpeakerModelOptions := textToSpeechService.NewDeleteSpeakerModelOptions(
				"testString",
			)

			response, err := textToSpeechService.DeleteSpeakerModel(deleteSpeakerModelOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSpeakerModel(): %d\n", response.StatusCode)
			}

			// end-deleteSpeakerModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteCustomPrompt request example`, func() {
			// begin-deleteCustomPrompt

			deleteCustomPromptOptions := textToSpeechService.NewDeleteCustomPromptOptions(
				"testString",
				"testString",
			)

			response, err := textToSpeechService.DeleteCustomPrompt(deleteCustomPromptOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCustomPrompt(): %d\n", response.StatusCode)
			}

			// end-deleteCustomPrompt

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteCustomModel request example`, func() {
			// begin-deleteCustomModel

			deleteCustomModelOptions := textToSpeechService.NewDeleteCustomModelOptions(
				"testString",
			)

			response, err := textToSpeechService.DeleteCustomModel(deleteCustomModelOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCustomModel(): %d\n", response.StatusCode)
			}

			// end-deleteCustomModel

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
