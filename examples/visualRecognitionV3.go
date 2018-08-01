package main

import (
	"os"
	"fmt"
	watson "golang-sdk"
	"golang-sdk/visualrecognitionv3"
	"encoding/json"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Visual Recognition service
	visualRecognition, visualRecognitionErr := visualrecognitionv3.NewVisualRecognitionV3(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-03-19",
		APIkey: "YOUR SERVICE API KEY",
	})

	// Check successful instantiation
	if visualRecognitionErr != nil {
		fmt.Println(visualRecognitionErr)
		return
	}

	/* CLASSIFY */

	// Read file with image to classify
	pwd, _ := os.Getwd()
	imageFile, imageFileErr := os.Open(pwd + "/resources/visualRecognition.jpg")

	// Check successful file read
	if imageFileErr != nil {
		fmt.Println(imageFileErr)
		return
	}

	classifyOptions := visualrecognitionv3.NewClassifyOptions().
		SetImagesFile(*imageFile).
		SetURL("https://www.readersdigest.ca/wp-content/uploads/2011/01/4-ways-cheer-up-depressed-cat.jpg").
		SetThreshold(0.6).
		SetClassifierIds([]string{ "default", "food", "explicit" })

	// Call the visual recognition Classify method
	classify, classifyErr := visualRecognition.Classify(classifyOptions)

	// Check successful call
	if classifyErr != nil {
		fmt.Println(classifyErr)
		return
	}

	// Cast response from call to the specific struct returned by GetClassifyResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	classifyResult := visualrecognitionv3.GetClassifyResult(classify)

	// Check successful casting
	if classifyResult != nil {
		prettyPrint(classifyResult, "Classify")
	}
}
