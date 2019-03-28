package main

import (
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	assistant "github.com/watson-developer-cloud/go-sdk/assistantv1"
)

func main() {
	// Instantiate the Watson Assistant service
	service, serviceErr := assistant.NewAssistantV1(&assistant.AssistantV1Options{
		URL:       "YOUR SERVICE URL",
		Version:   "2018-07-10",
		IAMApiKey: "YOUR API KEY",
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

	metadata := make(map[string]interface{})
	metadata["property"] = "value"

	createEntity := assistant.CreateEntity{
		Entity:      core.StringPtr("pizzatoppingstest"),
		Description: core.StringPtr("Tasty pizza topping"),
		Metadata:    metadata,
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
	input := &assistant.MessageInput{}
	input.SetText(core.StringPtr("Hello, how are you?"))

	messageOptions := service.NewMessageOptions(*workspaceID).
		SetInput(input)

	// Call the Message method with no specified context
	response, responseErr = service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// To continue with the same assistant, pass in the context from the previous call
	context := service.GetMessageResult(response).Context

	input.SetText(core.StringPtr("What's the weather right now?"))
	messageOptions.SetContext(context).
		SetInput(input)

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
