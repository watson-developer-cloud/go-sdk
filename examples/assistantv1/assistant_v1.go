package main

import (
	"fmt"

	assistant "github.com/ibm-watson/go-sdk/assistantv1"
	core "github.com/ibm-watson/go-sdk/core"
)

func main() {
	// Instantiate the Watson Assistant service
	service, serviceErr := assistant.NewAssistantV1(&assistant.AssistantV1Options{
		URL:      "YOUR SERVICE URL",
		Version:  "2018-07-10",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* LIST WORKSPACES */

	// Call the assistant ListWorkspaces method
	response, responseErr := service.ListWorkspaces(&assistant.ListWorkspacesOptions{})

	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	/* CREATE WORKSPACE */

	createEntity := assistant.CreateEntity{
		Entity:      core.StringPtr("pizzatoppingstest"),
		Description: core.StringPtr("Tasty pizza topping"),
		Metadata:    map[string]string{"property": "value"},
	}
	createWorkspaceOptions := service.NewCreateWorkspaceOptions().
		SetName("Test Workspace").
		SetDescription("GO example workspace").
		SetEntities([]assistant.CreateEntity{createEntity})

	response, responseErr = service.CreateWorkspace(createWorkspaceOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// Cast response.Result to the specific dataType
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	createWorkspaceResult := service.GetCreateWorkspaceResult(response)
	workspaceID := createWorkspaceResult.WorkspaceID

	// 	/* GET WORKSPACE */

	// Call the assistant GetWorkspace method
	response, responseErr = service.GetWorkspace(service.
		NewGetWorkspaceOptions(*workspaceID).
		SetExport(true))

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// 	/* UPDATE WORKSPACE */

	updateWorkspaceOptions := service.NewUpdateWorkspaceOptions(*workspaceID).
		SetName("Updated workspace name").
		SetDescription("Updated description")

	response, responseErr = service.UpdateWorkspace(updateWorkspaceOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// 	/* MESSAGE */

	inputData := &assistant.InputData{
		Text: core.StringPtr("Hello, how are you?"),
	}

	messageOptions := service.NewMessageOptions(*workspaceID).
		SetInput(inputData)

	// Call the Message method with no specified context
	response, responseErr = service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// To continue with the same assistant, pass in the context from the previous call
	conversationID := service.GetMessageResult(response).Context.ConversationID
	context := &assistant.Context{
		ConversationID: conversationID,
	}

	inputData.Text = core.StringPtr("What's the weather right now?")
	messageOptions.SetContext(context).
		SetInput(inputData)

	response, responseErr = service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// 	/* DELETE WORKSPACE */

	// Call the assistant DeleteWorkspace method
	response, responseErr = service.DeleteWorkspace(service.NewDeleteWorkspaceOptions(*workspaceID))

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}
	fmt.Println(response)
}
