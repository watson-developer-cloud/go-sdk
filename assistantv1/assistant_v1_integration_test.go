// +build integration

package assistantv1

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
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/watson-developer-cloud/go-sdk/core"
	"os"
	"testing"
)

var service *AssistantV1
var serviceErr error

func TestInitialization(t *testing.T) {
	err := godotenv.Load("../.env")
	require.Nil(t, err)

	service, serviceErr = NewAssistantV1(&AssistantV1Options{
		URL:      os.Getenv("ASSISTANT_URL"),
		Version:  "2018-09-20",
		Username: os.Getenv("ASSISTANT_USERNAME"),
		Password: os.Getenv("ASSISTANT_PASSWORD"),
	})
	require.Nil(t, serviceErr)
}

func TestListCounterexamples(t *testing.T) {
	response, responseErr := service.ListCounterexamples(service.
		NewListCounterexamplesOptions(os.Getenv("ASSISTANT_WORKSPACE_ID")))
	require.Nil(t, responseErr)

	result := service.GetListCounterexamplesResult(response)
	assert.NotNil(t, result)
}

func TestCounterexamples(t *testing.T) {
	// List Counter Examples
	response, responseErr := service.ListCounterexamples(service.
		NewListCounterexamplesOptions(os.Getenv("ASSISTANT_WORKSPACE_ID")))
	assert.Nil(t, responseErr)

	result := service.GetListCounterexamplesResult(response)
	assert.NotNil(t, result)

	// Create counter example
	response, responseErr = service.CreateCounterexample(service.
		NewCreateCounterexampleOptions(os.Getenv("ASSISTANT_WORKSPACE_ID"), "Make me a lemonade"))
	assert.Nil(t, responseErr)

	createCounterExample := service.GetCreateCounterexampleResult(response)
	assert.NotNil(t, createCounterExample)

	// Get counter example
	response, responseErr = service.GetCounterexample(service.
		NewGetCounterexampleOptions(os.Getenv("ASSISTANT_WORKSPACE_ID"), "Make me a lemonade"))
	assert.Nil(t, responseErr)

	getCounterExample := service.GetGetCounterexampleResult(response)
	assert.NotNil(t, getCounterExample)

	// Update counter example
	options := service.NewUpdateCounterexampleOptions(os.Getenv("ASSISTANT_WORKSPACE_ID"),
		"Make me a lemonade").
		SetNewText("Make me a smoothie")
	response, responseErr = service.UpdateCounterexample(options)
	assert.Nil(t, responseErr)

	updateCounterExample := service.GetUpdateCounterexampleResult(response)
	assert.NotNil(t, updateCounterExample)

	// Delete counter example
	response, responseErr = service.DeleteCounterexample(service.
		NewDeleteCounterexampleOptions(os.Getenv("ASSISTANT_WORKSPACE_ID"), "Make me a smoothie"))
	assert.NotNil(t, response)
}

func TestEntity(t *testing.T) {
	// List entities
	response, responseErr := service.ListEntities(
		&ListEntitiesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listEntities := service.GetListEntitiesResult(response)
	assert.NotNil(t, listEntities)

	// Create entity
	response, responseErr = service.CreateEntity(
		&CreateEntityOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Values: []CreateValue{
				CreateValue{
					Value: core.StringPtr("expresso"),
				},
				CreateValue{
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
		&GetEntityOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Export:      core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)

	getEntity := service.GetGetEntityResult(response)
	assert.NotNil(t, getEntity)

	// Update entity
	response, responseErr = service.UpdateEntity(
		&UpdateEntityOptions{
			WorkspaceID:    core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:         core.StringPtr("coffee"),
			NewDescription: core.StringPtr("cafe"),
		},
	)
	assert.Nil(t, responseErr)

	updateEntity := service.GetUpdateEntityResult(response)
	assert.NotNil(t, updateEntity)
}

func TestValues(t *testing.T) {
	// List values
	response, responseErr := service.ListValues(
		&ListValuesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)

	listValues := service.GetListValuesResult(response)
	assert.NotNil(t, listValues)

	// Create value
	response, responseErr = service.CreateValue(
		&CreateValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
		},
	)
	assert.Nil(t, responseErr)

	createValue := service.GetCreateValueResult(response)
	assert.NotNil(t, createValue)

	//Get value
	response, responseErr = service.GetValue(
		&GetValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
		},
	)
	assert.Nil(t, responseErr)

	getValue := service.GetGetValueResult(response)
	assert.NotNil(t, getValue)

	// Update value
	response, responseErr = service.UpdateValue(
		&UpdateValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
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
	// List mentions
	response, responseErr := service.ListMentions(
		&ListMentionsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)

	listMentions := service.GetListMentionsResult(response)
	assert.NotNil(t, listMentions)
}

func TestSynonyms(t *testing.T) {
	// List synonyms
	response, responseErr := service.ListSynonyms(
		&ListSynonymsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	listSynonyms := service.GetListSynonymsResult(response)
	assert.NotNil(t, listSynonyms)

	// Create synonym
	response, responseErr = service.CreateSynonym(
		&CreateSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
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
		&GetSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
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
		&UpdateSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
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
		&DeleteSynonymOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("N.M."),
		},
	)
	assert.Nil(t, responseErr)

	// Delete value
	response, responseErr = service.DeleteValue(
		&DeleteValueOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete entity
	response, responseErr = service.DeleteEntity(
		&DeleteEntityOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestIntents(t *testing.T) {
	// List intents
	response, responseErr := service.ListIntents(
		&ListIntentsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listIntents := service.GetListIntentsResult(response)
	assert.NotNil(t, listIntents)

	// Create intent
	response, responseErr = service.CreateIntent(
		&CreateIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hello"),
			Description: core.StringPtr("greetings"),
		},
	)
	assert.Nil(t, responseErr)

	createIntent := service.GetCreateIntentResult(response)
	assert.NotNil(t, createIntent)

	//Get intent
	response, responseErr = service.GetIntent(
		&GetIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hello"),
		},
	)
	assert.Nil(t, responseErr)

	getIntent := service.GetGetIntentResult(response)
	assert.NotNil(t, getIntent)

	// Update intent
	response, responseErr = service.UpdateIntent(
		&UpdateIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hello"),
			NewIntent:   core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)

	updateIntent := service.GetUpdateIntentResult(response)
	assert.NotNil(t, updateIntent)
}

func TestExamples(t *testing.T) {
	// List examples
	response, responseErr := service.ListExamples(
		&ListExamplesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)

	listExamples := service.GetListExamplesResult(response)
	assert.NotNil(t, listExamples)

	// Create example
	response, responseErr = service.CreateExample(
		&CreateExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
		},
	)
	assert.Nil(t, responseErr)

	createExample := service.GetCreateExampleResult(response)
	assert.NotNil(t, createExample)

	//Get example
	response, responseErr = service.GetExample(
		&GetExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
		},
	)
	assert.Nil(t, responseErr)

	getExample := service.GetGetExampleResult(response)
	assert.NotNil(t, getExample)

	// Update example
	response, responseErr = service.UpdateExample(
		&UpdateExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
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
		&DeleteExampleOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Hello there!"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete intent
	response, responseErr = service.DeleteIntent(
		&DeleteIntentOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Intent:      core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestDialogNodes(t *testing.T) {
	// List dialog nodes
	response, responseErr := service.ListDialogNodes(
		&ListDialogNodesOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listDialogNodes := service.GetListDialogNodesResult(response)
	assert.NotNil(t, listDialogNodes)

	// Create dialog node
	response, responseErr = service.CreateDialogNode(
		&CreateDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
			Conditions:  core.StringPtr("#hello"),
			Output: &DialogNodeOutput{
				Generic: []DialogNodeOutputGeneric{
					DialogNodeOutputGeneric{
						ResponseType: core.StringPtr(DialogNodeOutputGeneric_ResponseType_Text),
						Values: []DialogNodeOutputTextValuesElement{
							DialogNodeOutputTextValuesElement{
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
		&GetDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
		},
	)
	assert.Nil(t, responseErr)

	getDialogNode := service.GetGetDialogNodeResult(response)
	assert.NotNil(t, getDialogNode)

	// Update dialog node
	response, responseErr = service.UpdateDialogNode(
		&UpdateDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
			NewTitle:    core.StringPtr("Greeting."),
		},
	)
	assert.Nil(t, responseErr)

	updateDialogNode := service.GetUpdateDialogNodeResult(response)
	assert.NotNil(t, updateDialogNode)

	// Delete dialog node
	response, responseErr = service.DeleteDialogNode(
		&DeleteDialogNodeOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			DialogNode:  core.StringPtr("greeting"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestWorkspaces(t *testing.T) {
	// List workspaces
	response, responseErr := service.ListWorkspaces(
		&ListWorkspacesOptions{},
	)
	assert.Nil(t, responseErr)

	listWorkspaces := service.GetListWorkspacesResult(response)
	assert.NotNil(t, listWorkspaces)

	// Create workspace
	response, responseErr = service.CreateWorkspace(
		&CreateWorkspaceOptions{
			Name:        core.StringPtr("API test"),
			Description: core.StringPtr("Example workspace created via API"),
		},
	)
	assert.Nil(t, responseErr)

	createWorkspace := service.GetCreateWorkspaceResult(response)
	assert.NotNil(t, createWorkspace)

	//Get workspace
	response, responseErr = service.GetWorkspace(
		&GetWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
		},
	)
	assert.Nil(t, responseErr)

	getWorkspace := service.GetGetWorkspaceResult(response)
	assert.NotNil(t, getWorkspace)

	// Update workspace
	response, responseErr = service.UpdateWorkspace(
		&UpdateWorkspaceOptions{
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
		&DeleteWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
		},
	)
	assert.Nil(t, responseErr)
}

func TestMessage(t *testing.T) {
	response, responseErr := service.Message(
		&MessageOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
			Input: &InputData{
				Text: core.StringPtr("Hello World"),
			},
		},
	)
	assert.Nil(t, responseErr)

	message := service.GetMessageResult(response)
	assert.NotNil(t, message)
}

func TestLogs(t *testing.T) {
	// List logs
	response, responseErr := service.ListLogs(
		&ListLogsOptions{
			WorkspaceID: core.StringPtr(os.Getenv("ASSISTANT_WORKSPACE_ID")),
		},
	)
	assert.Nil(t, responseErr)

	listLogs := service.GetListLogsResult(response)
	assert.NotNil(t, listLogs)

	// list all logs
	response, responseErr = service.ListAllLogs(
		&ListAllLogsOptions{
			Filter: core.StringPtr("language::en,request.context.metadata.deployment::testDeployment"),
		},
	)
	assert.Nil(t, responseErr)

	listAllLogs := service.GetListAllLogsResult(response)
	assert.NotNil(t, listAllLogs)
}
