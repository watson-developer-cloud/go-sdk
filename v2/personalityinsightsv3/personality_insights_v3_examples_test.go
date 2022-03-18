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

package personalityinsightsv3_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/personalityinsightsv3"
)

//
// This file provides an example of how to use the Personality Insights service.
//
// The following configuration properties are assumed to be defined:
// PERSONALITY_INSIGHTS_URL=<service base url>
// PERSONALITY_INSIGHTS_AUTH_TYPE=iam
// PERSONALITY_INSIGHTS_APIKEY=<IAM apikey>
// PERSONALITY_INSIGHTS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../personality_insights_v3.env"

var (
	personalityInsightsService *personalityinsightsv3.PersonalityInsightsV3
	config                     map[string]string
	configLoaded               bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`PersonalityInsightsV3 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(personalityinsightsv3.DefaultServiceName)
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

			personalityInsightsServiceOptions := &personalityinsightsv3.PersonalityInsightsV3Options{
				Version: core.StringPtr("testString"),
			}

			personalityInsightsService, err = personalityinsightsv3.NewPersonalityInsightsV3(personalityInsightsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(personalityInsightsService).ToNot(BeNil())
		})
	})

	Describe(`PersonalityInsightsV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Profile request example`, func() {
			fmt.Println("\nProfile() result:")
			// begin-profile

			contentItemModel := &personalityinsightsv3.ContentItem{
				Content: core.StringPtr("testString"),
			}

			contentModel := &personalityinsightsv3.Content{
				ContentItems: []personalityinsightsv3.ContentItem{*contentItemModel},
			}

			profileOptions := personalityInsightsService.NewProfileOptions()
			profileOptions.SetContent(contentModel)

			profile, response, err := personalityInsightsService.Profile(profileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ProfileAsCSV request example`, func() {
			fmt.Println("\nProfileAsCSV() result:")
			// begin-profileAsCsv

			contentItemModel := &personalityinsightsv3.ContentItem{
				Content: core.StringPtr("testString"),
			}

			contentModel := &personalityinsightsv3.Content{
				ContentItems: []personalityinsightsv3.ContentItem{*contentItemModel},
			}

			profileOptions := personalityInsightsService.NewProfileOptions()
			profileOptions.SetContent(contentModel)

			csvFile, response, err := personalityInsightsService.ProfileAsCSV(profileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(csvFile, "", "  ")
			fmt.Println(string(b))

			// end-profileAsCsv

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(csvFile).ToNot(BeNil())

		})
	})
})
