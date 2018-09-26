package main

import (
	"fmt"
	"os"

	"github.com/ibm-watson/go-sdk/discoveryv1"

	discovery "github.com/ibm-watson/go-sdk/discoveryv1"
)

func main() {
	// Instantiate the Watson Discovery service
	service, serviceErr := discovery.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
		URL:       "YOUR SERVICE URL",
		Version:   "2018-03-05",
		IAMApiKey: "YOUR IAM API KEY",
	})
	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println(serviceErr)
		return
	}

	/* LIST ENVIRONMENTS */

	// Create a new ListEnvironmentsOptions and set optional parameter Name
	listEnvironmentsOptions := service.NewListEnvironmentsOptions()

	// Call the discovery ListEnvironments method
	response, responseErr := service.ListEnvironments(listEnvironmentsOptions)

	// Check successful call
	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	fmt.Println(response)

	// Cast listEnvironment.Result to the specific dataType returned by ListEnvironments
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	listEnvironmentResult := service.GetListEnvironmentsResult(response)

	// Check successful casting
	if listEnvironmentResult != nil {
		fmt.Println(listEnvironmentResult.Environments[0])
	}

	/* ADD DOCUMENT */

	file, err := os.Open("<PATH TO YOUR FILE>")
	if err != nil {
		panic(err)
	}

	addDocumentOptions := service.NewAddDocumentOptions("<ENVIRONMENT ID>",
		"<COLLECTION ID>").
		SetFile(file)

	response, responseErr = service.AddDocument(addDocumentOptions)

	if responseErr != nil {
		panic(responseErr)
	}

	defer file.Close()

	fmt.Println(response)

	/* QUERY */

	queryOptions := service.NewQueryOptions("<ENVIRONMENT ID>", "<COLLECTION ID>").
		SetFilter("extracted_metadata.sha1::9181d244*").
		SetReturnFields([]string{"extracted_metadata.sha1"})

	response, responseErr = service.Query(queryOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)
}
