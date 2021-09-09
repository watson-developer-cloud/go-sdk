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

package assistantv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv1"
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
const externalConfigFile = "../assistant_v1.env"

var (
	assistantService *assistantv1.AssistantV1
	config           map[string]string
	configLoaded     bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`AssistantV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(assistantv1.DefaultServiceName)
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

			assistantServiceOptions := &assistantv1.AssistantV1Options{
				Version: core.StringPtr("testString"),
			}

			assistantService, err = assistantv1.NewAssistantV1(assistantServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(assistantService).ToNot(BeNil())
		})
	})

	Describe(`AssistantV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Message request example`, func() {
			fmt.Println("\nMessage() result:")
			// begin-message

			messageOptions := assistantService.NewMessageOptions(
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
		It(`ListWorkspaces request example`, func() {
			fmt.Println("\nListWorkspaces() result:")
			// begin-listWorkspaces

			listWorkspacesOptions := assistantService.NewListWorkspacesOptions()

			workspaceCollection, response, err := assistantService.ListWorkspaces(listWorkspacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceCollection, "", "  ")
			fmt.Println(string(b))

			// end-listWorkspaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceCollection).ToNot(BeNil())

		})
		It(`CreateWorkspace request example`, func() {
			fmt.Println("\nCreateWorkspace() result:")
			// begin-createWorkspace

			createWorkspaceOptions := assistantService.NewCreateWorkspaceOptions()

			workspace, response, err := assistantService.CreateWorkspace(createWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-createWorkspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(workspace).ToNot(BeNil())

		})
		It(`GetWorkspace request example`, func() {
			fmt.Println("\nGetWorkspace() result:")
			// begin-getWorkspace

			getWorkspaceOptions := assistantService.NewGetWorkspaceOptions(
				"testString",
			)

			workspace, response, err := assistantService.GetWorkspace(getWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-getWorkspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`UpdateWorkspace request example`, func() {
			fmt.Println("\nUpdateWorkspace() result:")
			// begin-updateWorkspace

			updateWorkspaceOptions := assistantService.NewUpdateWorkspaceOptions(
				"testString",
			)

			workspace, response, err := assistantService.UpdateWorkspace(updateWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-updateWorkspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`ListIntents request example`, func() {
			fmt.Println("\nListIntents() result:")
			// begin-listIntents

			listIntentsOptions := assistantService.NewListIntentsOptions(
				"testString",
			)

			intentCollection, response, err := assistantService.ListIntents(listIntentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(intentCollection, "", "  ")
			fmt.Println(string(b))

			// end-listIntents

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(intentCollection).ToNot(BeNil())

		})
		It(`CreateIntent request example`, func() {
			fmt.Println("\nCreateIntent() result:")
			// begin-createIntent

			createIntentOptions := assistantService.NewCreateIntentOptions(
				"testString",
				"testString",
			)

			intent, response, err := assistantService.CreateIntent(createIntentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(intent, "", "  ")
			fmt.Println(string(b))

			// end-createIntent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(intent).ToNot(BeNil())

		})
		It(`GetIntent request example`, func() {
			fmt.Println("\nGetIntent() result:")
			// begin-getIntent

			getIntentOptions := assistantService.NewGetIntentOptions(
				"testString",
				"testString",
			)

			intent, response, err := assistantService.GetIntent(getIntentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(intent, "", "  ")
			fmt.Println(string(b))

			// end-getIntent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(intent).ToNot(BeNil())

		})
		It(`UpdateIntent request example`, func() {
			fmt.Println("\nUpdateIntent() result:")
			// begin-updateIntent

			updateIntentOptions := assistantService.NewUpdateIntentOptions(
				"testString",
				"testString",
			)

			intent, response, err := assistantService.UpdateIntent(updateIntentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(intent, "", "  ")
			fmt.Println(string(b))

			// end-updateIntent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(intent).ToNot(BeNil())

		})
		It(`ListExamples request example`, func() {
			fmt.Println("\nListExamples() result:")
			// begin-listExamples

			listExamplesOptions := assistantService.NewListExamplesOptions(
				"testString",
				"testString",
			)

			exampleCollection, response, err := assistantService.ListExamples(listExamplesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(exampleCollection, "", "  ")
			fmt.Println(string(b))

			// end-listExamples

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(exampleCollection).ToNot(BeNil())

		})
		It(`CreateExample request example`, func() {
			fmt.Println("\nCreateExample() result:")
			// begin-createExample

			createExampleOptions := assistantService.NewCreateExampleOptions(
				"testString",
				"testString",
				"testString",
			)

			example, response, err := assistantService.CreateExample(createExampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(example, "", "  ")
			fmt.Println(string(b))

			// end-createExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(example).ToNot(BeNil())

		})
		It(`GetExample request example`, func() {
			fmt.Println("\nGetExample() result:")
			// begin-getExample

			getExampleOptions := assistantService.NewGetExampleOptions(
				"testString",
				"testString",
				"testString",
			)

			example, response, err := assistantService.GetExample(getExampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(example, "", "  ")
			fmt.Println(string(b))

			// end-getExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(example).ToNot(BeNil())

		})
		It(`UpdateExample request example`, func() {
			fmt.Println("\nUpdateExample() result:")
			// begin-updateExample

			updateExampleOptions := assistantService.NewUpdateExampleOptions(
				"testString",
				"testString",
				"testString",
			)

			example, response, err := assistantService.UpdateExample(updateExampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(example, "", "  ")
			fmt.Println(string(b))

			// end-updateExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(example).ToNot(BeNil())

		})
		It(`ListCounterexamples request example`, func() {
			fmt.Println("\nListCounterexamples() result:")
			// begin-listCounterexamples

			listCounterexamplesOptions := assistantService.NewListCounterexamplesOptions(
				"testString",
			)

			counterexampleCollection, response, err := assistantService.ListCounterexamples(listCounterexamplesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(counterexampleCollection, "", "  ")
			fmt.Println(string(b))

			// end-listCounterexamples

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(counterexampleCollection).ToNot(BeNil())

		})
		It(`CreateCounterexample request example`, func() {
			fmt.Println("\nCreateCounterexample() result:")
			// begin-createCounterexample

			createCounterexampleOptions := assistantService.NewCreateCounterexampleOptions(
				"testString",
				"testString",
			)

			counterexample, response, err := assistantService.CreateCounterexample(createCounterexampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(counterexample, "", "  ")
			fmt.Println(string(b))

			// end-createCounterexample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(counterexample).ToNot(BeNil())

		})
		It(`GetCounterexample request example`, func() {
			fmt.Println("\nGetCounterexample() result:")
			// begin-getCounterexample

			getCounterexampleOptions := assistantService.NewGetCounterexampleOptions(
				"testString",
				"testString",
			)

			counterexample, response, err := assistantService.GetCounterexample(getCounterexampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(counterexample, "", "  ")
			fmt.Println(string(b))

			// end-getCounterexample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(counterexample).ToNot(BeNil())

		})
		It(`UpdateCounterexample request example`, func() {
			fmt.Println("\nUpdateCounterexample() result:")
			// begin-updateCounterexample

			updateCounterexampleOptions := assistantService.NewUpdateCounterexampleOptions(
				"testString",
				"testString",
			)

			counterexample, response, err := assistantService.UpdateCounterexample(updateCounterexampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(counterexample, "", "  ")
			fmt.Println(string(b))

			// end-updateCounterexample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(counterexample).ToNot(BeNil())

		})
		It(`ListEntities request example`, func() {
			fmt.Println("\nListEntities() result:")
			// begin-listEntities

			listEntitiesOptions := assistantService.NewListEntitiesOptions(
				"testString",
			)

			entityCollection, response, err := assistantService.ListEntities(listEntitiesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entityCollection, "", "  ")
			fmt.Println(string(b))

			// end-listEntities

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entityCollection).ToNot(BeNil())

		})
		It(`CreateEntity request example`, func() {
			fmt.Println("\nCreateEntity() result:")
			// begin-createEntity

			createEntityOptions := assistantService.NewCreateEntityOptions(
				"testString",
				"testString",
			)

			entity, response, err := assistantService.CreateEntity(createEntityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(b))

			// end-createEntity

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(entity).ToNot(BeNil())

		})
		It(`GetEntity request example`, func() {
			fmt.Println("\nGetEntity() result:")
			// begin-getEntity

			getEntityOptions := assistantService.NewGetEntityOptions(
				"testString",
				"testString",
			)

			entity, response, err := assistantService.GetEntity(getEntityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(b))

			// end-getEntity

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entity).ToNot(BeNil())

		})
		It(`UpdateEntity request example`, func() {
			fmt.Println("\nUpdateEntity() result:")
			// begin-updateEntity

			updateEntityOptions := assistantService.NewUpdateEntityOptions(
				"testString",
				"testString",
			)

			entity, response, err := assistantService.UpdateEntity(updateEntityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entity, "", "  ")
			fmt.Println(string(b))

			// end-updateEntity

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entity).ToNot(BeNil())

		})
		It(`ListMentions request example`, func() {
			fmt.Println("\nListMentions() result:")
			// begin-listMentions

			listMentionsOptions := assistantService.NewListMentionsOptions(
				"testString",
				"testString",
			)

			entityMentionCollection, response, err := assistantService.ListMentions(listMentionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(entityMentionCollection, "", "  ")
			fmt.Println(string(b))

			// end-listMentions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entityMentionCollection).ToNot(BeNil())

		})
		It(`ListValues request example`, func() {
			fmt.Println("\nListValues() result:")
			// begin-listValues

			listValuesOptions := assistantService.NewListValuesOptions(
				"testString",
				"testString",
			)

			valueCollection, response, err := assistantService.ListValues(listValuesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueCollection, "", "  ")
			fmt.Println(string(b))

			// end-listValues

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueCollection).ToNot(BeNil())

		})
		It(`CreateValue request example`, func() {
			fmt.Println("\nCreateValue() result:")
			// begin-createValue

			createValueOptions := assistantService.NewCreateValueOptions(
				"testString",
				"testString",
				"testString",
			)

			value, response, err := assistantService.CreateValue(createValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(value, "", "  ")
			fmt.Println(string(b))

			// end-createValue

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(value).ToNot(BeNil())

		})
		It(`GetValue request example`, func() {
			fmt.Println("\nGetValue() result:")
			// begin-getValue

			getValueOptions := assistantService.NewGetValueOptions(
				"testString",
				"testString",
				"testString",
			)

			value, response, err := assistantService.GetValue(getValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(value, "", "  ")
			fmt.Println(string(b))

			// end-getValue

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(value).ToNot(BeNil())

		})
		It(`UpdateValue request example`, func() {
			fmt.Println("\nUpdateValue() result:")
			// begin-updateValue

			updateValueOptions := assistantService.NewUpdateValueOptions(
				"testString",
				"testString",
				"testString",
			)

			value, response, err := assistantService.UpdateValue(updateValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(value, "", "  ")
			fmt.Println(string(b))

			// end-updateValue

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(value).ToNot(BeNil())

		})
		It(`ListSynonyms request example`, func() {
			fmt.Println("\nListSynonyms() result:")
			// begin-listSynonyms

			listSynonymsOptions := assistantService.NewListSynonymsOptions(
				"testString",
				"testString",
				"testString",
			)

			synonymCollection, response, err := assistantService.ListSynonyms(listSynonymsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(synonymCollection, "", "  ")
			fmt.Println(string(b))

			// end-listSynonyms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(synonymCollection).ToNot(BeNil())

		})
		It(`CreateSynonym request example`, func() {
			fmt.Println("\nCreateSynonym() result:")
			// begin-createSynonym

			createSynonymOptions := assistantService.NewCreateSynonymOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			synonym, response, err := assistantService.CreateSynonym(createSynonymOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(synonym, "", "  ")
			fmt.Println(string(b))

			// end-createSynonym

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(synonym).ToNot(BeNil())

		})
		It(`GetSynonym request example`, func() {
			fmt.Println("\nGetSynonym() result:")
			// begin-getSynonym

			getSynonymOptions := assistantService.NewGetSynonymOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			synonym, response, err := assistantService.GetSynonym(getSynonymOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(synonym, "", "  ")
			fmt.Println(string(b))

			// end-getSynonym

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(synonym).ToNot(BeNil())

		})
		It(`UpdateSynonym request example`, func() {
			fmt.Println("\nUpdateSynonym() result:")
			// begin-updateSynonym

			updateSynonymOptions := assistantService.NewUpdateSynonymOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			synonym, response, err := assistantService.UpdateSynonym(updateSynonymOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(synonym, "", "  ")
			fmt.Println(string(b))

			// end-updateSynonym

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(synonym).ToNot(BeNil())

		})
		It(`ListDialogNodes request example`, func() {
			fmt.Println("\nListDialogNodes() result:")
			// begin-listDialogNodes

			listDialogNodesOptions := assistantService.NewListDialogNodesOptions(
				"testString",
			)

			dialogNodeCollection, response, err := assistantService.ListDialogNodes(listDialogNodesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dialogNodeCollection, "", "  ")
			fmt.Println(string(b))

			// end-listDialogNodes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dialogNodeCollection).ToNot(BeNil())

		})
		It(`CreateDialogNode request example`, func() {
			fmt.Println("\nCreateDialogNode() result:")
			// begin-createDialogNode

			createDialogNodeOptions := assistantService.NewCreateDialogNodeOptions(
				"testString",
				"testString",
			)

			dialogNode, response, err := assistantService.CreateDialogNode(createDialogNodeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dialogNode, "", "  ")
			fmt.Println(string(b))

			// end-createDialogNode

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dialogNode).ToNot(BeNil())

		})
		It(`GetDialogNode request example`, func() {
			fmt.Println("\nGetDialogNode() result:")
			// begin-getDialogNode

			getDialogNodeOptions := assistantService.NewGetDialogNodeOptions(
				"testString",
				"testString",
			)

			dialogNode, response, err := assistantService.GetDialogNode(getDialogNodeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dialogNode, "", "  ")
			fmt.Println(string(b))

			// end-getDialogNode

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dialogNode).ToNot(BeNil())

		})
		It(`UpdateDialogNode request example`, func() {
			fmt.Println("\nUpdateDialogNode() result:")
			// begin-updateDialogNode

			updateDialogNodeOptions := assistantService.NewUpdateDialogNodeOptions(
				"testString",
				"testString",
			)

			dialogNode, response, err := assistantService.UpdateDialogNode(updateDialogNodeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dialogNode, "", "  ")
			fmt.Println(string(b))

			// end-updateDialogNode

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dialogNode).ToNot(BeNil())

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
		It(`ListAllLogs request example`, func() {
			fmt.Println("\nListAllLogs() result:")
			// begin-listAllLogs

			listAllLogsOptions := assistantService.NewListAllLogsOptions(
				"testString",
			)

			logCollection, response, err := assistantService.ListAllLogs(listAllLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logCollection, "", "  ")
			fmt.Println(string(b))

			// end-listAllLogs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logCollection).ToNot(BeNil())

		})
		It(`DeleteWorkspace request example`, func() {
			// begin-deleteWorkspace

			deleteWorkspaceOptions := assistantService.NewDeleteWorkspaceOptions(
				"testString",
			)

			response, err := assistantService.DeleteWorkspace(deleteWorkspaceOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteWorkspace
			fmt.Printf("\nDeleteWorkspace() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteValue request example`, func() {
			// begin-deleteValue

			deleteValueOptions := assistantService.NewDeleteValueOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteValue(deleteValueOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteValue
			fmt.Printf("\nDeleteValue() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

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
		It(`DeleteSynonym request example`, func() {
			// begin-deleteSynonym

			deleteSynonymOptions := assistantService.NewDeleteSynonymOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteSynonym(deleteSynonymOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteSynonym
			fmt.Printf("\nDeleteSynonym() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteIntent request example`, func() {
			// begin-deleteIntent

			deleteIntentOptions := assistantService.NewDeleteIntentOptions(
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteIntent(deleteIntentOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteIntent
			fmt.Printf("\nDeleteIntent() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteExample request example`, func() {
			// begin-deleteExample

			deleteExampleOptions := assistantService.NewDeleteExampleOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteExample(deleteExampleOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteExample
			fmt.Printf("\nDeleteExample() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteEntity request example`, func() {
			// begin-deleteEntity

			deleteEntityOptions := assistantService.NewDeleteEntityOptions(
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteEntity(deleteEntityOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteEntity
			fmt.Printf("\nDeleteEntity() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteDialogNode request example`, func() {
			// begin-deleteDialogNode

			deleteDialogNodeOptions := assistantService.NewDeleteDialogNodeOptions(
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteDialogNode(deleteDialogNodeOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteDialogNode
			fmt.Printf("\nDeleteDialogNode() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteCounterexample request example`, func() {
			// begin-deleteCounterexample

			deleteCounterexampleOptions := assistantService.NewDeleteCounterexampleOptions(
				"testString",
				"testString",
			)

			response, err := assistantService.DeleteCounterexample(deleteCounterexampleOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteCounterexample
			fmt.Printf("\nDeleteCounterexample() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
