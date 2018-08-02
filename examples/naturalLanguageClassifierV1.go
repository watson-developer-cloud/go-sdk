package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/naturallanguageclassifierv1"
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
	nlc, nlcErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(watson.Credentials{
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

	createClassifierOptions := naturallanguageclassifierv1.NewCreateClassifierOptions(*metadata, *data)

	// Call the natural language classifier CreateClassifier method
	create, createErr := nlc.CreateClassifier(createClassifierOptions)

	// Check successful call
	if createErr != nil {
		fmt.Println(createErr)
		return
	}

	// Cast response from call to the specific struct returned by GetCreateClassifierResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	createResult := naturallanguageclassifierv1.GetCreateClassifierResult(create)

	// Check successful casting
	if createResult != nil {
		prettyPrint(createResult, "Create Classifier")
	}
}
