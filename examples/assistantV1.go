package main

import (
	"fmt"
	. "go-sdk/assistantV1"
	"encoding/json"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Assistant service
	assistant, assistantErr := NewAssistantV1(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-07-10",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if assistantErr != nil {
		fmt.Println(assistantErr)
		return
	}


	/* LIST WORKSPACES */

	// Call the assistant ListWorkspaces method
	list, listErr := assistant.ListWorkspaces(NewListWorkspacesOptions())

	// Check successful call
	if listErr != nil {
		fmt.Println(listErr)
		return
	}

	// Cast list.Result to the specific dataType returned by ListWorkspaces
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	listResult := GetListWorkspacesResult(list)

	// Check successful casting
	if listResult != nil {
		prettyPrint(listResult, "List Workspaces")
	}


	/* GET WORKSPACE */

	// Call the assistant GetWorkspace method
	getWorkspaceOptions := NewGetWorkspaceOptions(listResult.Workspaces[0].WorkspaceID)
	get, getErr := assistant.GetWorkspace(getWorkspaceOptions)

	// Check successful call
	if getErr != nil {
		fmt.Println(getErr)
		return
	}

	// Cast result
	getResult := GetGetWorkspaceResult(get)

	// Check successful casting
	if getResult != nil {
		prettyPrint(getResult, "Get Workspace")
	}


	/* DELETE WORKSPACE */

	// Call the assistant DeleteWorkspace method
	deleteWorkspaceOptions := NewDeleteWorkspaceOptions(getResult.WorkspaceID)
	del, delErr := assistant.DeleteWorkspace(deleteWorkspaceOptions)

	// Check successful call
	if delErr != nil {
		fmt.Println(delErr)
		return
	}

	// NOTE: this method has no corresponding GetDeleteWorkspaceResult() function because DeleteWorkspace returns nothing

	prettyPrint(del, "Delete Workspace")


	/* CREATE WORKSPACE */

	createWorkspaceOptions := NewCreateWorkspaceOptions().
		SetName("Test Workspace")

	create, createErr := assistant.CreateWorkspace(createWorkspaceOptions)

	// Check successful call
	if createErr != nil {
		fmt.Println(createErr)
		return
	}

	// Cast result
	createResult := GetCreateWorkspaceResult(create)

	// Check successful casting
	if createResult != nil {
		prettyPrint(createResult, "Create Workspace")
	}


	/* UPDATE WORKSPACE */

	updateWorkspaceOptions := NewUpdateWorkspaceOptions(createResult.WorkspaceID).
		SetName("Updated workspace name").
		SetDescription("Updated description")

	update, updateErr := assistant.UpdateWorkspace(updateWorkspaceOptions)

	// Check successful call
	if updateErr != nil {
		fmt.Println(updateErr)
		return
	}

	// Cast result
	updateResult := GetUpdateWorkspaceResult(update)

	// Check successful casting
	if updateResult != nil {
		prettyPrint(updateResult, "Update Workspace")
	}
}
