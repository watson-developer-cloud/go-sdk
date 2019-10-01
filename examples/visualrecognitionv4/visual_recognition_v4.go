package main

import (
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv4"
)

func main() {
	// Instantiate the Watson Visual Recognition service
	service, serviceErr := visualrecognitionv4.
		NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
			Version: "2019-02-11",
			Authenticator: &core.IamAuthenticator{
				ApiKey: "YOUR APIKEY",
			},
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}
	service.SetServiceURL("https://gateway.watsonplatform.net/visual-recognition/api")

	/* CREATE COLLLECTION */
	collection, _, responseErr := service.CreateCollection(
		&visualrecognitionv4.CreateCollectionOptions{
			Name:        core.StringPtr("my_go_collection_for_training"),
			Description: core.StringPtr("simple collection for go"),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	collectionId := collection.CollectionID
	core.PrettyPrint(collection, "Collection")

	/* ADD IMAGES */
	pwd, _ := os.Getwd()
	giraffeFile, _ := os.Open(pwd + "/../../resources/South_Africa_Luca_Galuzzi_2004.jpeg")
	defer giraffeFile.Close()

	addImages, _, responseErr := service.AddImages(
		&visualrecognitionv4.AddImagesOptions{
			CollectionID: collectionId,
			ImagesFile: []visualrecognitionv4.FileWithMetadata{
				visualrecognitionv4.FileWithMetadata{
					Data:     giraffeFile,
					Filename: core.StringPtr("hello giraffe"),
				},
			},
		},
	)
	core.PrettyPrint(addImages, "Add images result: ")

	/* ADD IMAGE TRAINING DATA */
	trainingData, _, responseErr := service.AddImageTrainingData(
		&visualrecognitionv4.AddImageTrainingDataOptions{
			CollectionID: collectionId,
			ImageID:      addImages.Images[0].ImageID,
			Objects: []visualrecognitionv4.TrainingDataObject{
				visualrecognitionv4.TrainingDataObject{
					Object: core.StringPtr("giraffe training data"),
					Location: &visualrecognitionv4.Location{
						Top:    core.Int64Ptr(64),
						Left:   core.Int64Ptr(270),
						Width:  core.Int64Ptr(755),
						Height: core.Int64Ptr(784),
					},
				},
			},
		},
	)
	core.PrettyPrint(trainingData, "Training data: ")

	/* TRAIN */
	train, _, responseErr := service.Train(
		&visualrecognitionv4.TrainOptions{
			CollectionID: collectionId,
		},
	)
	core.PrettyPrint(train, "Training result: ")

	/* ANALYZE */
	imageFile, _ := os.Open(pwd + "/../../resources/my-giraffe.jpeg")
	defer imageFile.Close()

	result, _, _ := service.Analyze(
		&visualrecognitionv4.AnalyzeOptions{
			CollectionIds: []string{*collectionId},
			Features:      []string{visualrecognitionv4.AnalyzeOptions_Features_Objects},
			ImagesFile: []visualrecognitionv4.FileWithMetadata{
				visualrecognitionv4.FileWithMetadata{
					Data:     imageFile,
					Filename: core.StringPtr("random name"),
				},
			},
		},
	)
	core.PrettyPrint(result, "Analyze result: ")
}
