package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/assistantv1"
)

func main() {
	// Instantiate the Watson Assistant service
	assistant, assistantErr := assistantv1.NewAssistantV1(watson.Credentials{
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
	list, listErr := assistant.ListWorkspaces(assistantv1.NewListWorkspacesOptions())

	// Check successful call
	if listErr != nil {
		fmt.Println(listErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListWorkspacesResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listResult := assistantv1.GetListWorkspacesResult(list)

	// Check successful casting
	if listResult != nil {
		// Print result
		fmt.Printf("FOUND %v WORKSPACES\n", len(listResult.Workspaces))
		fmt.Println(listResult)
	}


	/* GET WORKSPACE */

	// Call the assistant GetWorkspace method
	getWorkspaceOptions := assistantv1.NewGetWorkspaceOptions("0a0c06c1-8e31-4655-9067-58fcac5134fc")
	get, getErr := assistant.GetWorkspace(getWorkspaceOptions)

	// Check successful call
	if getErr != nil {
		fmt.Println(getErr)
		return
	}

	// Cast response from call to the specific struct returned by GetGetWorkspaceResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	getResult := assistantv1.GetGetWorkspaceResult(get)

	// Check successful casting
	if getResult != nil {
		// Print result
		fmt.Println(getResult)
	}
}
