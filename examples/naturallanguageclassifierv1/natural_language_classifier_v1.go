package main

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	nlc "github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
)

func main() {
	// Instantiate the Watson Natural Language Classifier service
	service, serviceErr := nlc.NewNaturalLanguageClassifierV1(&nlc.NaturalLanguageClassifierV1Options{
		URL:       "YOUR SERVICE URL",
		IAMApiKey: "YOUR IAM API KEY",
	})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* CREATE CLASSIFIER */

	pwd, _ := os.Getwd()

	metadata, metadataErr := os.Open(pwd + "/../../resources/weather_training_metadata.json")
	if metadataErr != nil {
		fmt.Println(metadataErr)
	}

	data, dataErr := os.Open(pwd + "/../../resources/weather_training_data.csv")
	if dataErr != nil {
		fmt.Println(dataErr)
	}

	createClassifierOptions := service.NewCreateClassifierOptions(metadata, data)

	// Call the natural language classifier CreateClassifier method
	response, responseErr := service.CreateClassifier(createClassifierOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast create.Result to the specific dataType returned by CreateClassifier
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	createResult := service.GetCreateClassifierResult(response)

	// Check successful casting
	if createResult != nil {
		core.PrettyPrint(createResult, "Create Classifier")
	}

	/* CLASSIFY */

	if *createResult.Status == "Available" {
		classifyOptions := service.NewClassifyOptions(*createResult.ClassifierID, "How hot will it be tomorrow?")

		// Call the natural language classifier Classify method
		response, responseErr := service.Classify(classifyOptions)

		// Check successful call
		if responseErr != nil {
			panic(responseErr)
		}

		// Cast result
		classifyResult := service.GetClassifyResult(response)

		// Check successful casting
		if classifyResult != nil {
			core.PrettyPrint(classifyResult, "Classify")
		}
	}
}
