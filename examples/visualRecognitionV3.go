package main

import (
	"os"
	"fmt"
	watson "golang-sdk"
	"golang-sdk/visualrecognitionv3"
)

func main() {
	// Instantiate the Watson Visual Recognition service
	visualRecognition, visualRecognitionErr := visualrecognitionv3.NewVisualRecognitionV3(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-03-19",
		APIkey: "YOUR SERVICE API KEY",
	})

	// Check successful instantiation
	if visualRecognitionErr != nil {
		fmt.Println(visualRecognition)
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

	// Call the visual recognition Classify  method
	classify, classifyErr := visualRecognition.Classify(*imageFile, "en", "https://www.readersdigest.ca/wp-content/uploads/2011/01/4-ways-cheer-up-depressed-cat.jpg", 0.5, []string{"IBM"}, []string{"default"}, "JPEG")

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
		// Print result
		fmt.Println(classifyResult)
	}

}
