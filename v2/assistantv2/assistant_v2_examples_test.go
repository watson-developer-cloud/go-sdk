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

package assistantv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

//
// This file provides an example of how to use the Assistant service.
//
// The following configuration properties are assumed to be defined:
// CONVERSATION_URL=<service base url>
// CONVERSATION_AUTH_TYPE=iam
// CONVERSATION_APIKEY=<IAM apikey>
// CONVERSATION_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../assistant_v2.env"

var (
	assistantService *assistantv2.AssistantV2
	config           map[string]string
	configLoaded     bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`AssistantV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(assistantv2.DefaultServiceName)
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

			assistantServiceOptions := &assistantv2.AssistantV2Options{
				Version: core.StringPtr("testString"),
			}

			assistantService, err = assistantv2.NewAssistantV2(assistantServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(assistantService).ToNot(BeNil())
		})
	})

	Describe(`AssistantV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSession request example`, func() {
			fmt.Println("\nCreateSession() result:")
			// begin-createSession

			createSessionOptions := assistantService.NewCreateSessionOptions(
				"testString",
			)

			sessionResponse, response, err := assistantService.CreateSession(createSessionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sessionResponse, "", "  ")
			fmt.Println(string(b))

			// end-createSession

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sessionResponse).ToNot(BeNil())

		})
		It(`Message request example`, func() {
			fmt.Println("\nMessage() result:")
			// begin-message

			messageOptions := assistantService.NewMessageOptions(
				"testString",
				"testString",
			)

			messageResponse, response, err := assistantService.Message(messageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(messageResponse, "", "  ")
			fmt.Println(string(b))

			// end-message

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(messageResponse).ToNot(BeNil())

		})
		It(`MessageStateless request example`, func() {
			fmt.Println("\nMessageStateless() result:")
			// begin-messageStateless

			messageStatelessOptions := assistantService.NewMessageStatelessOptions(
				"testString",
			)

			messageResponseStateless, response, err := assistantService.MessageStateless(messageStatelessOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(messageResponseStateless, "", "  ")
			fmt.Println(string(b))

			// end-messageStateless

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(messageResponseStateless).ToNot(BeNil())

		})
		It(`BulkClassify request example`, func() {
			fmt.Println("\nBulkClassify() result:")
			// begin-bulkClassify

			bulkClassifyOptions := assistantService.NewBulkClassifyOptions(
				"testString",
			)

			bulkClassifyResponse, response, err := assistantService.BulkClassify(bulkClassifyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(bulkClassifyResponse, "", "  ")
			fmt.Println(string(b))

			// end-bulkClassify

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bulkClassifyResponse).ToNot(BeNil())

		})
		It(`ListLogs request example`, func() {
			fmt.Println("\nListLogs() result:")
			// begin-listLogs

			listLogsOptions := assistantService.NewListLogsOptions(
				"testString",
			)

			logCollection, response, err := assistantService.ListLogs(listLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logCollection, "", "  ")
			fmt.Println(string(b))

			// end-listLogs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logCollection).ToNot(BeNil())

		})
		It(`DeleteUserData request example`, func() {
			// begin-deleteUserData

			deleteUserDataOptions := assistantService.NewDeleteUserDataOptions(
				"testString",
			)

			response, err := assistantService.DeleteUserData(deleteUserDataOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteUserData
			fmt.Printf("\nDeleteUserData() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteSession request example`, func() {
			// begin-deleteSession

			deleteSessionOptions := assistantService.NewDeleteSessionOptions(
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteSession(deleteSessionOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteSession
			fmt.Printf("\nDeleteSession() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
