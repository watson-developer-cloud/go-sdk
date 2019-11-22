package main

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"

	discovery "github.com/watson-developer-cloud/go-sdk/discoveryv2"
)

func main() {
	// Discovery v2 is only available on Cloud Pak for Data.
	// Instantiate the Watson Discovery service
	authenticator := &core.CloudPakForDataAuthenticator{
		URL:                    "<authenticating url>",
		Username:               "<your username>",
		Password:               "<your password>",
		DisableSSLVerification: true,
	}
	service, serviceErr := discovery.NewDiscoveryV2(&discovery.DiscoveryV2Options{
		URL:           "<your service url>",
		Version:       "2019-11-22",
		Authenticator: authenticator,
	})
	service.DisableSSLVerification()

	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println(serviceErr)
		return
	}

	PROJECT_ID := "9558dc01-8554-4d18-b0a5-70196f9f2fe6"

	// LIST COLLECTIONS
	listCollectionsResult, _, responseErr := service.ListCollections(&discovery.ListCollectionsOptions{
		ProjectID: core.StringPtr(PROJECT_ID),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if listCollectionsResult != nil {
		core.PrettyPrint(listCollectionsResult, "Collections: ")
	}

	// COMPONENT SETTINGS
	settingsResult, _, responseErr := service.GetComponentSettings(&discovery.GetComponentSettingsOptions{
		ProjectID: core.StringPtr(PROJECT_ID),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if settingsResult != nil {
		core.PrettyPrint(settingsResult, "Component Settings: ")
	}

	// ADD DOCUMENT
	COLLECTION_ID := "161d1e47-9651-e657-0000-016e8e6fb5b6"
	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../../resources/example.html")

	if fileErr != nil {
		panic(fileErr)
	}

	documentResult, response, responseErr := service.AddDocument(&discovery.AddDocumentOptions{
		ProjectID:       core.StringPtr(PROJECT_ID),
		CollectionID:    core.StringPtr(COLLECTION_ID),
		File:            file,
		Filename:        core.StringPtr("example.html"),
		FileContentType: core.StringPtr("text/html"),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		fmt.Println(response)
		return
	}

	if documentResult != nil {
		core.PrettyPrint(documentResult, "Document: ")
	}
	documentID := documentResult.DocumentID

	// CREATE TRAINING DATA
	trainingResult, _, responseErr := service.CreateTrainingQuery(&discovery.CreateTrainingQueryOptions{
		ProjectID:            core.StringPtr(PROJECT_ID),
		NaturalLanguageQuery: core.StringPtr("How is the weather today?"),
		Examples: []discovery.TrainingExample{
			discovery.TrainingExample{
				DocumentID:   documentID,
				CollectionID: core.StringPtr(COLLECTION_ID),
				Relevance:    core.Int64Ptr(1),
			},
		},
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if trainingResult != nil {
		core.PrettyPrint(trainingResult, "Training : ")
	}

	// TRAINING QUERIES
	trainingQueriesResult, _, responseErr := service.ListTrainingQueries(&discovery.ListTrainingQueriesOptions{
		ProjectID: core.StringPtr(PROJECT_ID),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if trainingQueriesResult != nil {
		core.PrettyPrint(trainingQueriesResult, "Training queries: ")
	}

	// QUERIES //
	// query
	queryResult, _, responseErr := service.Query(&discovery.QueryOptions{
		ProjectID:            core.StringPtr(PROJECT_ID),
		CollectionIds:        []string{COLLECTION_ID},
		NaturalLanguageQuery: core.StringPtr("How is the weather today?"),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if queryResult != nil {
		core.PrettyPrint(queryResult, "Query: ")
	}

	// get autocomplete
	getAutocompletionResult, _, responseErr := service.GetAutocompletion(&discovery.GetAutocompletionOptions{
		ProjectID: core.StringPtr(PROJECT_ID),
		Prefix:    core.StringPtr("IBM"),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if getAutocompletionResult != nil {
		core.PrettyPrint(getAutocompletionResult, "Autocompletion: ")
	}

	// get notices
	noticesResult, _, responseErr := service.QueryNotices(&discovery.QueryNoticesOptions{
		ProjectID:            core.StringPtr(PROJECT_ID),
		NaturalLanguageQuery: core.StringPtr("warning"),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if noticesResult != nil {
		core.PrettyPrint(noticesResult, "Notices: ")
	}

	// list fields
	fieldsResult, _, responseErr := service.ListFields(&discovery.ListFieldsOptions{
		ProjectID: core.StringPtr(PROJECT_ID),
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if fieldsResult != nil {
		core.PrettyPrint(fieldsResult, "List fields: ")
	}

	// CLEANUP
	//delete training queries
	_, responseErr = service.DeleteTrainingQueries(&discovery.DeleteTrainingQueriesOptions{
		ProjectID: core.StringPtr(PROJECT_ID),
	})

	// delete document
	deleteDocumentResult, _, responseErr := service.DeleteDocument(&discovery.DeleteDocumentOptions{
		ProjectID:    core.StringPtr(PROJECT_ID),
		CollectionID: core.StringPtr(COLLECTION_ID),
		DocumentID:   documentID,
	})

	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	if deleteDocumentResult != nil {
		core.PrettyPrint(deleteDocumentResult, "Document result: ")
	}
}
