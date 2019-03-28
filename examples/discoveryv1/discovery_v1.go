package main

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"os"

	discovery "github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func main() {
	// Instantiate the Watson Discovery service
	service, serviceErr := discovery.NewDiscoveryV1(&discovery.DiscoveryV1Options{
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

	environmentID := "<YOUR ENVIRONEMNT ID>"
	collectionID := "<YOUR COLLECTION ID>"

	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../../resources/example.html")

	if fileErr != nil {
		panic(fileErr)
	}

	addDocumentOptions := service.NewAddDocumentOptions(environmentID,
		collectionID).
		SetFile(file).
		SetMetadata("{\"Creator\": \"Johnny Appleseed\", \"Subject\": \"Apples\" }")

	response, responseErr = service.AddDocument(addDocumentOptions)

	if responseErr != nil {
		panic(responseErr)
	}

	defer file.Close()

	core.PrettyPrint(response.GetResult(), "Add document: ")

	/* QUERY */

	queryOptions := service.NewQueryOptions(environmentID, collectionID).
		SetFilter("extracted_metadata.sha1::9181d244*").
		SetReturnFields("extracted_metadata.sha1")

	response, responseErr = service.Query(queryOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)
}
