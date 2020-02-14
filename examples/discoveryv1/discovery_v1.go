package main

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v3/core"

	discovery "github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func main() {
	// Instantiate the Watson Discovery service
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("YOUR IAM API KEY"),
	}
	service, serviceErr := discovery.NewDiscoveryV1(&discovery.DiscoveryV1Options{
		URL:           "YOUR SERVICE URL",
		Version:       "2018-03-05",
		Authenticator: authenticator,
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
	listEnvironmentResult, response, responseErr := service.ListEnvironments(listEnvironmentsOptions)

	// Check successful call
	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	fmt.Println(response)

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

	_, response, responseErr = service.AddDocument(addDocumentOptions)

	if responseErr != nil {
		panic(responseErr)
	}

	defer file.Close()

	core.PrettyPrint(response.GetResult(), "Add document: ")

	/* QUERY */

	queryOptions := service.NewQueryOptions(environmentID, collectionID).
		SetFilter("extracted_metadata.sha1::9181d244*").
		SetReturn("extracted_metadata.sha1")

	_, response, responseErr = service.Query(queryOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)
}
