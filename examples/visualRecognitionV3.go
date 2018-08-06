package main

import (
	"os"
	"fmt"
	. "go-sdk/visualRecognitionV3"
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
	visualRecognition, visualRecognitionErr := NewVisualRecognitionV3(&ServiceCredentials{
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

	// Open file with image to classify
	pwd, _ := os.Getwd()
	imageFile, imageFileErr := os.Open(pwd + "/resources/kitty.jpg")

	// Check successful file read
	if imageFileErr != nil {
		fmt.Println(imageFileErr)
		return
	}

	classifyOptions := NewClassifyOptions().
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

	// Cast classify.Result to the specific dataType returned by Classify
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	classifyResult := GetClassifyResult(classify)

	// Check successful casting
	if classifyResult != nil {
		prettyPrint(classifyResult, "Classify")
	}


	/* CREATE CLASSIFIER */

	carsFile, carsFileErr := os.Open(pwd + "/resources/cars.zip")
	if carsFileErr != nil {
		fmt.Println(carsFileErr)
		return
	}

	trucksFile, trucksFileErr := os.Open(pwd + "/resources/trucks.zip")
	if trucksFileErr != nil {
		fmt.Println(trucksFileErr)
		return
	}

	createClassifierOptions := NewCreateClassifierOptions("Cars vs Trucks", "cars", *carsFile).
		SetNegativeExamples(*trucksFile)

	create, createErr := visualRecognition.CreateClassifier(createClassifierOptions)
	if createErr != nil {
		fmt.Println(createErr)
		return
	}

	createResult := GetCreateClassifierResult(create)
	if createResult != nil {
		prettyPrint(createResult, "Create Classifier")
	}

	// Test classifier
	imageFile, imageFileErr = os.Open(pwd + "/resources/car.jpg")
	if imageFileErr != nil {
		fmt.Println(imageFileErr)
		return
	}

	classifyOptions = NewClassifyOptions().
		SetImagesFile(*imageFile)

	classify, classifyErr = visualRecognition.Classify(classifyOptions)
	if classifyErr != nil {
		fmt.Println(classifyErr)
		return
	}

	classifyResult = GetClassifyResult(classify)
	if classifyResult != nil {
		prettyPrint(classifyResult, "Classify")
	}


	/* DETECT FACES */

	imageFile, imageFileErr = os.Open(pwd + "/resources/face.jpg")
	if imageFileErr != nil {
		fmt.Println(imageFileErr)
		return
	}

	detectFacesOptions := NewDetectFacesOptions().
		SetImagesFile(*imageFile).
		SetURL("https://www.ibm.com/ibm/ginni/images/ginni_bio_780x981_v4_03162016.jpg")

	detect, detectErr := visualRecognition.DetectFaces(detectFacesOptions)
	if detectErr != nil {
		fmt.Println(detectErr)
		return
	}

	detectResult := GetDetectFacesResult(detect)
	if detectResult != nil {
		prettyPrint(detectResult, "Detect Faces")
	}
}
