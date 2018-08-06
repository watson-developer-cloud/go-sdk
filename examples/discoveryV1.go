package main

import (
	"fmt"
	. "go-sdk/discoveryV1"
	"encoding/json"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Discovery service
	discovery, discoveryV1Err := NewDiscoveryV1(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-03-05",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if discoveryV1Err != nil {
		fmt.Println(discoveryV1Err)
		return
	}


	/* LIST ENVIRONMENTS */

	// Create a new ListEnvironmentsOptions and set optional parameter Name
	listEnvironmentsOptions := NewListEnvironmentsOptions().
		SetName("Watson Discovery Environment")

	// Call the discovery ListEnvironments method
	listEnvironment, listEnvironmentErr := discovery.ListEnvironments(listEnvironmentsOptions)

	// Check successful call
	if listEnvironmentErr != nil {
		fmt.Println(listEnvironmentErr)
		return
	}

	// Cast listEnvironment.Result to the specific dataType returned by ListEnvironments
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	listEnvironmentResult := GetListEnvironmentsResult(listEnvironment)

	// Check successful casting
	if listEnvironmentResult != nil {
		prettyPrint(listEnvironmentResult, "List Environments")
	}
}
