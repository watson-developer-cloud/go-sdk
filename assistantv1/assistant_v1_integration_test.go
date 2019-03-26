// +build integration

package assistantv1_test

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
	"github.com/watson-developer-cloud/go-sdk/assistantv1"
	"os"
	"testing"
)

var service *assistantv1.AssistantV1
var serviceErr error

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = assistantv1.
			NewAssistantV1(&assistantv1.AssistantV1Options{
				URL:      os.Getenv("ASSISTANT_GO_SDK_URL"),
				Version:  "2018-09-20",
				Username: os.Getenv("ASSISTANT_GO_SDK_USERNAME"),
				Password: os.Getenv("ASSISTANT_GO_SDK_PASSWORD"),
			})
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestCounterexamples(t *testing.T) {
	shouldSkipTest(t)

	// List Counter Examples
	response, responseErr := service.ListCounterexamples(service.
		NewListCounterexamplesOptions(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")))
	assert.Nil(t, responseErr)

	result := service.GetListCounterexamplesResult(response)
	assert.NotNil(t, result)

	// Create counter example
	response, responseErr = service.CreateCounterexample(service.
		NewCreateCounterexampleOptions(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID"), "Make me a lemonade?"))
	assert.Nil(t, responseErr)

	createCounterExample := service.GetCreateCounterexampleResult(response)
	assert.NotNil(t, createCounterExample)

	// Get counter example
	response, responseErr = service.GetCounterexample(service.
		NewGetCounterexampleOptions(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID"), "Make me a lemonade?"))
	assert.Nil(t, responseErr)

	getCounterExample := service.GetGetCounterexampleResult(response)
	assert.NotNil(t, getCounterExample)

	// Update counter example
	options := service.NewUpdateCounterexampleOptions(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID"),
		"Make me a lemonade?").
		SetNewText("Make me a smoothie?")
	response, responseErr = service.UpdateCounterexample(options)
	assert.Nil(t, responseErr)

	updateCounterExample := service.GetUpdateCounterexampleResult(response)
	assert.NotNil(t, updateCounterExample)

	// Delete counter example
	response, responseErr = service.DeleteCounterexample(service.
		NewDeleteCounterexampleOptions(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID"), "Make me a smoothie?"))
	assert.NotNil(t, response)
}

func TestEntity(t *testing.T) {
	shouldSkipTest(t)

	// List entities
	response, responseErr := service.ListEntities(
		&assistantv1.ListEntitiesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listEntities := service.GetListEntitiesResult(response)
	assert.NotNil(t, listEntities)

	// Create entity
	response, responseErr = service.CreateEntity(
		&assistantv1.CreateEntityOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Values: []assistantv1.CreateValue{
				assistantv1.CreateValue{
					Value: core.StringPtr("expresso"),
				},
				assistantv1.CreateValue{
					Value: core.StringPtr("latte"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	createEntity := service.GetCreateEntityResult(response)
	assert.NotNil(t, createEntity)

	//Get entity
	response, responseErr = service.GetEntity(
		&assistantv1.GetEntityOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Export:      core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)

	getEntity := service.GetGetEntityResult(response)
	assert.NotNil(t, getEntity)

	// Update entity
	response, responseErr = service.UpdateEntity(
		&assistantv1.UpdateEntityOptions{
			WorkspaceID:    core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:         core.StringPtr("coffee"),
			NewDescription: core.StringPtr("cafe"),
		},
	)
	assert.Nil(t, responseErr)

	updateEntity := service.GetUpdateEntityResult(response)
	assert.NotNil(t, updateEntity)
}

func TestValues(t *testing.T) {
	shouldSkipTest(t)

	// List values
	response, responseErr := service.ListValues(
		&assistantv1.ListValuesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)

	listValues := service.GetListValuesResult(response)
	assert.NotNil(t, listValues)

	// Create value
	response, responseErr = service.CreateValue(
		&assistantv1.CreateValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
		},
	)
	assert.Nil(t, responseErr)

	createValue := service.GetCreateValueResult(response)
	assert.NotNil(t, createValue)

	//Get value
	response, responseErr = service.GetValue(
		&assistantv1.GetValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
		},
	)
	assert.Nil(t, responseErr)

	getValue := service.GetGetValueResult(response)
	assert.NotNil(t, getValue)

	// Update value
	response, responseErr = service.UpdateValue(
		&assistantv1.UpdateValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
			NewValue:    core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	updateValue := service.GetUpdateValueResult(response)
	assert.NotNil(t, updateValue)
}

func TestListMentions(t *testing.T) {
	shouldSkipTest(t)

	// List mentions
	response, responseErr := service.ListMentions(
		&assistantv1.ListMentionsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)

	listMentions := service.GetListMentionsResult(response)
	assert.NotNil(t, listMentions)
}

func TestSynonyms(t *testing.T) {
	shouldSkipTest(t)

	// List synonyms
	response, responseErr := service.ListSynonyms(
		&assistantv1.ListSynonymsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	listSynonyms := service.GetListSynonymsResult(response)
	assert.NotNil(t, listSynonyms)

	// Create synonym
	response, responseErr = service.CreateSynonym(
		&assistantv1.CreateSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("NM"),
		},
	)
	assert.Nil(t, responseErr)

	createSynonym := service.GetCreateSynonymResult(response)
	assert.NotNil(t, createSynonym)

	//Get synonym
	response, responseErr = service.GetSynonym(
		&assistantv1.GetSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("NM"),
		},
	)
	assert.Nil(t, responseErr)

	getSynonym := service.GetGetSynonymResult(response)
	assert.NotNil(t, getSynonym)

	// Update synonym
	response, responseErr = service.UpdateSynonym(
		&assistantv1.UpdateSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("NM"),
			NewSynonym:  core.StringPtr("N.M."),
		},
	)
	assert.Nil(t, responseErr)

	updateSynonym := service.GetUpdateSynonymResult(response)
	assert.NotNil(t, updateSynonym)

	// Delete synonym
	response, responseErr = service.DeleteSynonym(
		&assistantv1.DeleteSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("N.M."),
		},
	)
	assert.Nil(t, responseErr)

	// Delete value
	response, responseErr = service.DeleteValue(
		&assistantv1.DeleteValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete entity
	response, responseErr = service.DeleteEntity(
		&assistantv1.DeleteEntityOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestIntents(t *testing.T) {
	shouldSkipTest(t)

	// List intents
	response, responseErr := service.ListIntents(
		&assistantv1.ListIntentsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listIntents := service.GetListIntentsResult(response)
	assert.NotNil(t, listIntents)

	// Create intent
	response, responseErr = service.CreateIntent(
		&assistantv1.CreateIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hello"),
			Description: core.StringPtr("greetings"),
		},
	)
	assert.Nil(t, responseErr)

	createIntent := service.GetCreateIntentResult(response)
	assert.NotNil(t, createIntent)

	//Get intent
	response, responseErr = service.GetIntent(
		&assistantv1.GetIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hello"),
		},
	)
	assert.Nil(t, responseErr)

	getIntent := service.GetGetIntentResult(response)
	assert.NotNil(t, getIntent)

	// Update intent
	response, responseErr = service.UpdateIntent(
		&assistantv1.UpdateIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hello"),
			NewIntent:   core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)

	updateIntent := service.GetUpdateIntentResult(response)
	assert.NotNil(t, updateIntent)
}

func TestExamples(t *testing.T) {
	shouldSkipTest(t)

	// List examples
	response, responseErr := service.ListExamples(
		&assistantv1.ListExamplesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)

	listExamples := service.GetListExamplesResult(response)
	assert.NotNil(t, listExamples)

	// Create example
	response, responseErr = service.CreateExample(
		&assistantv1.CreateExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
		},
	)
	assert.Nil(t, responseErr)

	createExample := service.GetCreateExampleResult(response)
	assert.NotNil(t, createExample)

	//Get example
	response, responseErr = service.GetExample(
		&assistantv1.GetExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
		},
	)
	assert.Nil(t, responseErr)

	getExample := service.GetGetExampleResult(response)
	assert.NotNil(t, getExample)

	// Update example
	response, responseErr = service.UpdateExample(
		&assistantv1.UpdateExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
			NewText:     core.StringPtr("Hello there!"),
		},
	)
	assert.Nil(t, responseErr)

	updateExample := service.GetUpdateExampleResult(response)
	assert.NotNil(t, updateExample)

	// Delete example
	response, responseErr = service.DeleteExample(
		&assistantv1.DeleteExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Hello there!"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete intent
	response, responseErr = service.DeleteIntent(
		&assistantv1.DeleteIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestDialogNodes(t *testing.T) {
	shouldSkipTest(t)

	// List dialog nodes
	response, responseErr := service.ListDialogNodes(
		&assistantv1.ListDialogNodesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listDialogNodes := service.GetListDialogNodesResult(response)
	assert.NotNil(t, listDialogNodes)

	// Create dialog node
	response, responseErr = service.CreateDialogNode(
		&assistantv1.CreateDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
			Conditions:  core.StringPtr("#hello"),
			Output: &assistantv1.DialogNodeOutput{
				"generic": []assistantv1.DialogNodeOutputGeneric{
					assistantv1.DialogNodeOutputGeneric{
						ResponseType: core.StringPtr(assistantv1.DialogNodeOutputGeneric_ResponseType_Text),
						Values: []assistantv1.DialogNodeOutputTextValuesElement{
							assistantv1.DialogNodeOutputTextValuesElement{
								Text: core.StringPtr("Hi! How can I help you?"),
							},
						},
					},
				},
			},
			Title: core.StringPtr("Greeting"),
		},
	)
	assert.Nil(t, responseErr)

	createDialog := service.GetCreateDialogNodeResult(response)
	assert.NotNil(t, createDialog)

	//Get dialog node
	response, responseErr = service.GetDialogNode(
		&assistantv1.GetDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
		},
	)
	assert.Nil(t, responseErr)

	getDialogNode := service.GetGetDialogNodeResult(response)
	assert.NotNil(t, getDialogNode)

	// Update dialog node
	response, responseErr = service.UpdateDialogNode(
		&assistantv1.UpdateDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
			NewTitle:    core.StringPtr("Greeting."),
		},
	)
	assert.Nil(t, responseErr)

	updateDialogNode := service.GetUpdateDialogNodeResult(response)
	assert.NotNil(t, updateDialogNode)

	// Delete dialog node
	response, responseErr = service.DeleteDialogNode(
		&assistantv1.DeleteDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestWorkspaces(t *testing.T) {
	shouldSkipTest(t)

	// List workspaces
	response, responseErr := service.ListWorkspaces(
		&assistantv1.ListWorkspacesOptions{},
	)
	assert.Nil(t, responseErr)

	listWorkspaces := service.GetListWorkspacesResult(response)
	assert.NotNil(t, listWorkspaces)

	// Create workspace
	response, responseErr = service.CreateWorkspace(
		&assistantv1.CreateWorkspaceOptions{
			Name:        core.StringPtr("API test"),
			Description: core.StringPtr("Example workspace created via API"),
		},
	)
	assert.Nil(t, responseErr)

	createWorkspace := service.GetCreateWorkspaceResult(response)
	assert.NotNil(t, createWorkspace)

	//Get workspace
	response, responseErr = service.GetWorkspace(
		&assistantv1.GetWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
		},
	)
	assert.Nil(t, responseErr)

	getWorkspace := service.GetGetWorkspaceResult(response)
	assert.NotNil(t, getWorkspace)

	// Update workspace
	response, responseErr = service.UpdateWorkspace(
		&assistantv1.UpdateWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
			Name:        core.StringPtr("Updated workspace for GO"),
			Description: core.StringPtr("Example workspace updated via API"),
		},
	)
	assert.Nil(t, responseErr)

	updateWorkspace := service.GetUpdateWorkspaceResult(response)
	assert.NotNil(t, updateWorkspace)

	// Delete workspace
	response, responseErr = service.DeleteWorkspace(
		&assistantv1.DeleteWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
		},
	)
	assert.Nil(t, responseErr)
}

func TestMessage(t *testing.T) {
	shouldSkipTest(t)

	response, responseErr := service.Message(
		&assistantv1.MessageOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
			Input: &assistantv1.MessageInput{
				"text": "Hello World",
			},
		},
	)
	assert.Nil(t, responseErr)

	message := service.GetMessageResult(response)
	assert.NotNil(t, message)
}

func TestLogs(t *testing.T) {
	shouldSkipTest(t)

	// List logs
	response, responseErr := service.ListLogs(
		&assistantv1.ListLogsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_GO_SDK_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listLogs := service.GetListLogsResult(response)
	assert.NotNil(t, listLogs)

	// list all logs
	response, responseErr = service.ListAllLogs(
		&assistantv1.ListAllLogsOptions{
			Filter: core.StringPtr("language::en,request.context.metadata.deployment::testDeployment"),
		},
	)
	assert.Nil(t, responseErr)

	listAllLogs := service.GetListAllLogsResult(response)
	assert.NotNil(t, listAllLogs)
}
