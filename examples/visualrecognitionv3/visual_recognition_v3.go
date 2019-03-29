package main

import (
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
)

func main() {
	// Instantiate the Watson Visual Recognition service
	service, serviceErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			URL:       "YOUR SERVICE URL",
			Version:   "2018-03-19",
			IAMApiKey: "YOUR API KEY",
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* CLASSIFY */

	// Open file with image to classify
	pwd, _ := os.Getwd()
	imageFile, imageFileErr := os.Open(pwd + "/../../resources/kitty.jpg")

	// Check successful file read
	if imageFileErr != nil {
		panic(imageFileErr)
	}

	classifyOptions := service.NewClassifyOptions()
	classifyOptions.ImagesFile = imageFile
	classifyOptions.URL = core.StringPtr("https://www.readersdigest.ca/wp-content/uploads/sites/14/2011/01/4-ways-cheer-up-depressed-cat.jpg")
	classifyOptions.Threshold = core.Float32Ptr(0.6)
	classifyOptions.ClassifierIds = []string{"default", "food", "explicit"}

	// Call the visual recognition Classify method
	response, responseErr := service.Classify(classifyOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast classify.Result to the specific dataType returned by Classify
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	classifyResult := service.GetClassifyResult(response)

	// Check successful casting
	if classifyResult != nil {
		core.PrettyPrint(classifyResult, "Classify")
	}

	/* CREATE CLASSIFIER */

	carsFile, carsFileErr := os.Open(pwd + "/../../resources/cars.zip")
	if carsFileErr != nil {
		panic(carsFileErr)
	}

	trucksFile, trucksFileErr := os.Open(pwd + "/../../resources/trucks.zip")
	if trucksFileErr != nil {
		panic(trucksFileErr)
	}

	createClassifierOptions := service.
		NewCreateClassifierOptions("Cars vs Trucks").
		AddPositiveExamples("cars", carsFile)
	createClassifierOptions.NegativeExamples = trucksFile

	response, responseErr = service.CreateClassifier(createClassifierOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	createResult := service.GetCreateClassifierResult(response)
	if createResult != nil {
		core.PrettyPrint(createResult, "Create Classifier")
	}

	// Test classifier
	imageFile, imageFileErr = os.Open(pwd + "/../../resources/car.jpg")
	if imageFileErr != nil {
		panic(imageFileErr)
	}

	classifyOptions = service.NewClassifyOptions()
	classifyOptions.ImagesFile = imageFile

	response, responseErr = service.Classify(classifyOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	classifyResult = service.GetClassifyResult(response)
	if classifyResult != nil {
		core.PrettyPrint(classifyResult, "Classify")
	}

	// /* DETECT FACES */

	imageFile, imageFileErr = os.Open(pwd + "/../../resources/face.jpg")
	if imageFileErr != nil {
		panic(imageFileErr)
	}

	detectFacesOptions := service.NewDetectFacesOptions()
	detectFacesOptions.ImagesFile = imageFile
	detectFacesOptions.URL = core.StringPtr("https://www.ibm.com/ibm/ginni/images/ginni_bio_780x981_v4_03162016.jpg")

	response, responseErr = service.DetectFaces(detectFacesOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	detectResult := service.GetDetectFacesResult(response)
	if detectResult != nil {
		core.PrettyPrint(detectResult, "Detect Faces")
	}
}
