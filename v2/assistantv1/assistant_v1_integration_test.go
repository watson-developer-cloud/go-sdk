// +build integration

package assistantv1_test

/**
 * (C) Copyright IBM Corp. 2018, 2020.
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
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *assistantv1.AssistantV1
var workspaceId string

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load(configFile)
	if err != nil {
		t.Skip(skipMessage)
	}

	workspaceId = os.Getenv("ASSISTANT_WORKSPACE_ID")
	assert.NotEmpty(t, workspaceId)
	if workspaceId != "" {
		configLoaded = true
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
		Version:     core.StringPtr("2020-04-01"),
		ServiceName: "assistant",
	})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestCounterexamples(t *testing.T) {
	shouldSkipTest(t)

	// List Counter Examples
	result, _, responseErr := service.ListCounterexamples(service.
		NewListCounterexamplesOptions(workspaceId))
	assert.Nil(t, responseErr)

	assert.NotNil(t, result)

	// Create counter example
	createCounterExample, _, responseErr := service.CreateCounterexample(service.
		NewCreateCounterexampleOptions(workspaceId, "Make me a lemonade?"))
	assert.Nil(t, responseErr)

	assert.NotNil(t, createCounterExample)

	// Get counter example
	getCounterExample, _, responseErr := service.GetCounterexample(service.
		NewGetCounterexampleOptions(workspaceId, "Make me a lemonade?"))
	assert.Nil(t, responseErr)

	assert.NotNil(t, getCounterExample)

	// Update counter example
	options := service.NewUpdateCounterexampleOptions(workspaceId,
		"Make me a lemonade?").
		SetNewText("Make me a smoothie?")
	updateCounterExample, _, responseErr := service.UpdateCounterexample(options)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateCounterExample)

	// Delete counter example
	response, responseErr := service.DeleteCounterexample(service.
		NewDeleteCounterexampleOptions(workspaceId, "Make me a smoothie?"))
	assert.NotNil(t, response)
	assert.Nil(t, responseErr)
}

func TestEntity(t *testing.T) {
	shouldSkipTest(t)

	// List entities
	listEntities, _, responseErr := service.ListEntities(
		&assistantv1.ListEntitiesOptions{
			WorkspaceID: core.StringPtr(workspaceId),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listEntities)

	// Create entity
	createEntity, _, responseErr := service.CreateEntity(
		&assistantv1.CreateEntityOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Values: []assistantv1.CreateValue{
				{
					Value: core.StringPtr("expresso"),
				},
				{
					Value: core.StringPtr("latte"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createEntity)

	//Get entity
	getEntity, _, responseErr := service.GetEntity(
		&assistantv1.GetEntityOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Export:      core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getEntity)

	// Update entity
	updateEntity, _, responseErr := service.UpdateEntity(
		&assistantv1.UpdateEntityOptions{
			WorkspaceID:    core.StringPtr(workspaceId),
			Entity:         core.StringPtr("coffee"),
			NewDescription: core.StringPtr("cafe"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateEntity)
}

func TestValues(t *testing.T) {
	shouldSkipTest(t)

	// List values
	listValues, _, responseErr := service.ListValues(
		&assistantv1.ListValuesOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listValues)

	// Create value
	createValue, _, responseErr := service.CreateValue(
		&assistantv1.CreateValueOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createValue)

	//Get value
	getValue, _, responseErr := service.GetValue(
		&assistantv1.GetValueOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getValue)

	// Update value
	updateValue, _, responseErr := service.UpdateValue(
		&assistantv1.UpdateValueOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("mocha"),
			NewValue:    core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateValue)
}

func TestListMentions(t *testing.T) {
	shouldSkipTest(t)

	// List mentions
	listMentions, _, responseErr := service.ListMentions(
		&assistantv1.ListMentionsOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listMentions)
}

func TestSynonyms(t *testing.T) {
	shouldSkipTest(t)

	// List synonyms
	listSynonyms, _, responseErr := service.ListSynonyms(
		&assistantv1.ListSynonymsOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listSynonyms)

	// Create synonym
	createSynonym, _, responseErr := service.CreateSynonym(
		&assistantv1.CreateSynonymOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("NM"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createSynonym)

	//Get synonym
	getSynonym, _, responseErr := service.GetSynonym(
		&assistantv1.GetSynonymOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("NM"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getSynonym)

	// Update synonym
	updateSynonym, _, responseErr := service.UpdateSynonym(
		&assistantv1.UpdateSynonymOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("NM"),
			NewSynonym:  core.StringPtr("N.M."),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateSynonym)

	// Delete synonym
	_, responseErr = service.DeleteSynonym(
		&assistantv1.DeleteSynonymOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
			Synonym:     core.StringPtr("N.M."),
		},
	)
	assert.Nil(t, responseErr)

	// Delete value
	_, responseErr = service.DeleteValue(
		&assistantv1.DeleteValueOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
			Value:       core.StringPtr("new mocha"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete entity
	_, responseErr = service.DeleteEntity(
		&assistantv1.DeleteEntityOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Entity:      core.StringPtr("coffee"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestIntents(t *testing.T) {
	shouldSkipTest(t)

	// List intents
	listIntents, _, responseErr := service.ListIntents(
		&assistantv1.ListIntentsOptions{
			WorkspaceID: core.StringPtr(workspaceId),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listIntents)

	// Create intent
	createIntent, _, responseErr := service.CreateIntent(
		&assistantv1.CreateIntentOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hello"),
			Description: core.StringPtr("greetings"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createIntent)

	//Get intent
	getIntent, _, responseErr := service.GetIntent(
		&assistantv1.GetIntentOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hello"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getIntent)

	// Update intent
	updateIntent, _, responseErr := service.UpdateIntent(
		&assistantv1.UpdateIntentOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hello"),
			NewIntent:   core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateIntent)
}

func TestExamples(t *testing.T) {
	shouldSkipTest(t)

	// List examples
	listExamples, _, responseErr := service.ListExamples(
		&assistantv1.ListExamplesOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listExamples)

	// Create example
	createExample, _, responseErr := service.CreateExample(
		&assistantv1.CreateExampleOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createExample)

	//Get example
	getExample, _, responseErr := service.GetExample(
		&assistantv1.GetExampleOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getExample)

	// Update example
	updateExample, _, responseErr := service.UpdateExample(
		&assistantv1.UpdateExampleOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Howdy"),
			NewText:     core.StringPtr("Hello there!"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateExample)

	// Delete example
	_, responseErr = service.DeleteExample(
		&assistantv1.DeleteExampleOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hi"),
			Text:        core.StringPtr("Hello there!"),
		},
	)
	assert.Nil(t, responseErr)

	// Delete intent
	_, responseErr = service.DeleteIntent(
		&assistantv1.DeleteIntentOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Intent:      core.StringPtr("hi"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestDialogNodes(t *testing.T) {
	shouldSkipTest(t)

	// List dialog nodes
	listDialogNodes, _, responseErr := service.ListDialogNodes(
		&assistantv1.ListDialogNodesOptions{
			WorkspaceID: core.StringPtr(workspaceId),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listDialogNodes)

	// Create output type to be passed in as interface{} type
	outputNode := new(assistantv1.DialogNodeOutputGenericDialogNodeOutputResponseTypeText)
	outputNode.ResponseType = core.StringPtr("text")
	outputNode.Values = []assistantv1.DialogNodeOutputTextValuesElement{
		{
			Text: core.StringPtr("Hi! How can I help you?"),
		},
	}

	// Create dialog node
	createDialog, _, responseErr := service.CreateDialogNode(
		&assistantv1.CreateDialogNodeOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			DialogNode:  core.StringPtr("greeting"),
			Conditions:  core.StringPtr("#hello"),
			Output: &assistantv1.DialogNodeOutput{
				Generic: []assistantv1.DialogNodeOutputGenericIntf{outputNode},
			},
			Title:                core.StringPtr("Greeting"),
			DisambiguationOptOut: core.BoolPtr(true),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createDialog)

	//Get dialog node
	getDialogNode, _, responseErr := service.GetDialogNode(
		&assistantv1.GetDialogNodeOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			DialogNode:  core.StringPtr("greeting"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getDialogNode)

	// Update dialog node
	updateDialogNode, _, responseErr := service.UpdateDialogNode(
		&assistantv1.UpdateDialogNodeOptions{
			WorkspaceID:             core.StringPtr(workspaceId),
			DialogNode:              core.StringPtr("greeting"),
			NewTitle:                core.StringPtr("Greeting."),
			NewDisambiguationOptOut: core.BoolPtr(false),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateDialogNode)

	// Delete dialog node
	_, responseErr = service.DeleteDialogNode(
		&assistantv1.DeleteDialogNodeOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			DialogNode:  core.StringPtr("greeting"),
		},
	)
	assert.Nil(t, responseErr)
}

func TestWorkspaces(t *testing.T) {
	shouldSkipTest(t)

	// List workspaces
	listWorkspaces, _, responseErr := service.ListWorkspaces(
		&assistantv1.ListWorkspacesOptions{},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, listWorkspaces)

	// Create workspace
	createWorkspace, _, responseErr := service.CreateWorkspace(
		&assistantv1.CreateWorkspaceOptions{
			Name:        core.StringPtr("API test"),
			Description: core.StringPtr("Example workspace created via SDK"),
			Webhooks: []assistantv1.Webhook{
				{
					URL:  core.StringPtr("https://test-webhook"),
					Name: core.StringPtr("Dwight Schrute"),
				},
			},
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, createWorkspace)

	//Get workspace
	getWorkspace, _, responseErr := service.GetWorkspace(
		&assistantv1.GetWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, getWorkspace)

	// Update workspace
	updateWorkspace, _, responseErr := service.UpdateWorkspace(
		&assistantv1.UpdateWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
			Name:        core.StringPtr("Updated workspace for GO"),
			Description: core.StringPtr("Example workspace updated via API"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateWorkspace)

	// Delete workspace
	_, responseErr = service.DeleteWorkspace(
		&assistantv1.DeleteWorkspaceOptions{
			WorkspaceID: core.StringPtr(*createWorkspace.WorkspaceID),
		},
	)
	assert.Nil(t, responseErr)
}

func TestMessage(t *testing.T) {
	shouldSkipTest(t)

	message, _, responseErr := service.Message(
		&assistantv1.MessageOptions{
			WorkspaceID: core.StringPtr(workspaceId),
			Input: &assistantv1.MessageInput{
				Text: core.StringPtr("Hello World"),
			},
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, message)
}

func TestLogs(t *testing.T) {
	shouldSkipTest(t)

	// List logs
	listLogs, response, responseErr := service.ListLogs(
		&assistantv1.ListLogsOptions{
			WorkspaceID: core.StringPtr(workspaceId),
		},
	)

	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, listLogs)

	// list all logs
	listAllLogs, response, responseErr := service.ListAllLogs(
		&assistantv1.ListAllLogsOptions{
			Filter: core.StringPtr("language::en,request.context.metadata.deployment::testDeployment"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
	assert.NotNil(t, listAllLogs)
}
