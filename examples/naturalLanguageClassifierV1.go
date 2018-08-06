package main

import (
	"fmt"
	. "go-sdk/naturalLanguageClassifierV1"
	"encoding/json"
	"os"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Natural Language Classifier service
	nlc, nlcErr := NewNaturalLanguageClassifierV1(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-07-10",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if nlcErr != nil {
		fmt.Println(nlcErr)
		return
	}


	/* CREATE CLASSIFIER */

	pwd, _ := os.Getwd()

	metadata, metadataErr := os.Open(pwd + "/resources/weather_training_metadata.json")
	if metadataErr != nil {
		fmt.Println(metadataErr)
	}

	data, dataErr := os.Open(pwd + "/resources/weather_training_data.csv")
	if dataErr != nil {
		fmt.Println(dataErr)
	}

	createClassifierOptions := NewCreateClassifierOptions(*metadata, *data)

	// Call the natural language classifier CreateClassifier method
	create, createErr := nlc.CreateClassifier(createClassifierOptions)

	// Check successful call
	if createErr != nil {
		fmt.Println(createErr)
		return
	}

	// Cast create.Result to the specific dataType returned by CreateClassifier
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	createResult := GetCreateClassifierResult(create)

	// Check successful casting
	if createResult != nil {
		prettyPrint(createResult, "Create Classifier")
	}


	/* CLASSIFY */

	if createResult.Status == "Available" {
		classifyOptions := NewClassifyOptions(createResult.ClassifierID, "How hot will it be tomorrow?")

		// Call the natural language classifier Classify method
		classify, classifyErr := nlc.Classify(classifyOptions)

		// Check successful call
		if classifyErr != nil {
			fmt.Println(classifyErr)
			return
		}

		// Cast result
		classifyResult := GetClassifyResult(classify)

		// Check successful casting
		if classifyResult != nil {
			prettyPrint(classifyResult, "Classify")
		}
	}
}
