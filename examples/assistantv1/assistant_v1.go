package main

import (
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	assistant "github.com/watson-developer-cloud/go-sdk/assistantv1"
)

func main() {
	// Instantiate the Watson Assistant service
	authenticator := &core.IamAuthenticator{
		ApiKey: "my-iam-apikey",
	}
	service, serviceErr := assistant.NewAssistantV1(&assistant.AssistantV1Options{
		URL:           "YOUR SERVICE URL",
		Version:       "2018-07-10",
		Authenticator: authenticator,
	})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* LIST WORKSPACES */

	// Call the assistant ListWorkspaces method
	_, response, responseErr := service.ListWorkspaces(&assistant.ListWorkspacesOptions{})

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

	createWorkspaceResult, response, responseErr := service.CreateWorkspace(createWorkspaceOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)
	workspaceID := createWorkspaceResult.WorkspaceID

	// 	/* GET WORKSPACE */

	// Call the assistant GetWorkspace method
	_, response, responseErr = service.GetWorkspace(service.
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

	_, response, responseErr = service.UpdateWorkspace(updateWorkspaceOptions)

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
	messageResult, response, responseErr := service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// To continue with the same assistant, pass in the context from the previous call
	context := messageResult.Context

	input.SetText(core.StringPtr("What's the weather right now?"))
	messageOptions.SetContext(context).
		SetInput(input)

	_, response, responseErr = service.Message(messageOptions)

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
