// +build examples

/**
 * (C) Copyright IBM Corp. 2021.
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

package toneanalyzerv3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/toneanalyzerv3"
)

//
// This file provides an example of how to use the Tone Analyzer service.
//
// The following configuration properties are assumed to be defined:
// TONE_ANALYZER_URL=<service base url>
// TONE_ANALYZER_AUTH_TYPE=iam
// TONE_ANALYZER_APIKEY=<IAM apikey>
// TONE_ANALYZER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../tone_analyzer_v3.env"

var (
	toneAnalyzerService *toneanalyzerv3.ToneAnalyzerV3
	config              map[string]string
	configLoaded        bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ToneAnalyzerV3 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(toneanalyzerv3.DefaultServiceName)
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

			toneAnalyzerServiceOptions := &toneanalyzerv3.ToneAnalyzerV3Options{
				Version: core.StringPtr("testString"),
			}

			toneAnalyzerService, err = toneanalyzerv3.NewToneAnalyzerV3(toneAnalyzerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(toneAnalyzerService).ToNot(BeNil())
		})
	})

	Describe(`ToneAnalyzerV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Tone request example`, func() {
			fmt.Println("\nTone() result:")
			// begin-tone

			toneInputModel := &toneanalyzerv3.ToneInput{
				Text: core.StringPtr("testString"),
			}

			toneOptions := toneAnalyzerService.NewToneOptions()
			toneOptions.SetToneInput(toneInputModel)

			toneAnalysis, response, err := toneAnalyzerService.Tone(toneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(toneAnalysis, "", "  ")
			fmt.Println(string(b))

			// end-tone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(toneAnalysis).ToNot(BeNil())

		})
		It(`ToneChat request example`, func() {
			fmt.Println("\nToneChat() result:")
			// begin-toneChat

			utteranceModel := &toneanalyzerv3.Utterance{
				Text: core.StringPtr("testString"),
			}

			toneChatOptions := toneAnalyzerService.NewToneChatOptions(
				[]toneanalyzerv3.Utterance{*utteranceModel},
			)

			utteranceAnalyses, response, err := toneAnalyzerService.ToneChat(toneChatOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(utteranceAnalyses, "", "  ")
			fmt.Println(string(b))

			// end-toneChat

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(utteranceAnalyses).ToNot(BeNil())

		})
	})
})
