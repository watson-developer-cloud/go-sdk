package main

import (
	watson "go-sdk/assistantv1" //TODO: Update with the full path
	"go-sdk/core"
)

func main() {
	// Instantiate the Watson Assistant service
	assistant, assistantErr := watson.NewAssistantV1(&watson.AssistantV1Options{
		URL:      "YOUR SERVICE URL",
		Version:  "2018-07-10",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if assistantErr != nil {
		panic(assistantErr)
	}

	/* LIST WORKSPACES */

	// Call the assistant ListWorkspaces method
	response, responseErr := assistant.ListWorkspaces(&watson.ListWorkspacesOptions{})

	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful call
	response.PrettyPrint(response.GetResult())
	response.PrettyPrint(response.GetHeaders())
	response.PrettyPrint(response.GetStatusCode())

	/* CREATE WORKSPACE */

	createEntity := watson.CreateEntity{
		Entity:      core.StringPtr("pizzatoppingstest"),
		Description: core.StringPtr("Tasty pizza topping"),
		Metadata:    map[string]string{"property": "value"},
	}
	createWorkspaceOptions := watson.NewCreateWorkspaceOptions().
		SetName("Test Workspace").
		SetDescription("GO example workspace").
		SetEntities([]watson.CreateEntity{createEntity})

	response, responseErr = assistant.CreateWorkspace(createWorkspaceOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	response.PrettyPrint(response.GetResult())

	// 	/* GET WORKSPACE */

	// Call the assistant GetWorkspace method
	workspaceID := watson.GetCreateWorkspaceResult(response).WorkspaceID // TODO: see if we can get it directly
	response, responseErr = assistant.GetWorkspace(watson.
		NewGetWorkspaceOptions(*workspaceID).
		SetExport(true))

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	response.PrettyPrint(response.GetResult())

	// 	/* UPDATE WORKSPACE */

	updateWorkspaceOptions := watson.NewUpdateWorkspaceOptions(*workspaceID).
		SetName("Updated workspace name").
		SetDescription("Updated description")

	response, responseErr = assistant.UpdateWorkspace(updateWorkspaceOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	response.PrettyPrint(response.GetResult())

	// 	/* MESSAGE */

	inputData := watson.InputData{
		Text: core.StringPtr("Hello, how are you?"),
	}

	messageOptions := watson.NewMessageOptions(*workspaceID).
		SetInput(inputData)

	// Call the Message method with no specified context
	response, responseErr = assistant.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	response.PrettyPrint(response.GetResult())

	// To continue with the same assistant, pass in the context from the previous call
	conversationID := watson.GetMessageResult(response).Context.ConversationID // TODO: see if we can get it directly
	context := watson.Context{
		ConversationID: conversationID,
	}

	inputData.Text = core.StringPtr("What's the weather right now?")
	messageOptions.SetContext(context).
		SetInput(inputData)

	response, responseErr = assistant.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	response.PrettyPrint(response.GetResult())

	// 	/* DELETE WORKSPACE */

	// Call the assistant DeleteWorkspace method
	response, responseErr = assistant.DeleteWorkspace(watson.NewDeleteWorkspaceOptions(*workspaceID))

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}
	response.PrettyPrint(response.GetResult())
}
