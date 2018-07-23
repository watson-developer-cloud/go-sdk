package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/discoveryV1"
)

func main() {
	// Instantiate the Watson Discovery service
	discovery, discoveryV1Err := discoveryV1.NewDiscoveryV1(watson.Credentials{
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

	// Call the discovery List Environments method
	listEnvironment, listEnvironmentErr := discovery.ListEnvironments("Watson Discovery Environment")

	// Check successful call
	if listEnvironmentErr != nil {
		fmt.Println(listEnvironmentErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListEnvironmentsResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listEnvironmentResult := discoveryV1.GetListEnvironmentsResult(listEnvironment)

	// Check successful casting
	if listEnvironmentResult != nil {
		// Print result
		fmt.Println(listEnvironmentResult)
	}

}
