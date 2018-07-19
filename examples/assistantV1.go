package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/assistantV1"
)

func main() {
	// Instantiate the Watson Assistant service
	assistant, assistantErr := assistantV1.NewAssistantV1(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-02-16",
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
	list, listErr := assistant.ListWorkspaces(0, true, "", "", false)

	// Check successful call
	if listErr != nil {
		fmt.Println(listErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListWorkspacesResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	result := assistantV1.GetListWorkspacesResult(list)

	// Check successful casting
	if result != nil {
		// Print result
		fmt.Printf("FOUND %v WORKSPACES\n", len(result.Workspaces))
		fmt.Println(result)
	}
}
