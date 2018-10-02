package main

import (
	"os"

	"github.com/ibm-watson/go-sdk/core"
	"github.com/ibm-watson/go-sdk/speechtotextv1"
)

func main() {
	// Instantiate the Watson Speech To Text service
	service, serviceErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			URL:      "YOUR SERVICE URL",
			Username: "YOUR SERVICE URL",
			Password: "YOUR SERVICE URL",
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* RECOGNIZE */

	pwd, _ := os.Getwd()

	// Open file with mp3 to recognize
	audio, audioErr := os.Open(pwd + "/../../resources/audio_example.mp3")
	if audioErr != nil {
		panic(audioErr)
	}

	// Create a new RecognizeOptions for ContentType "audio/mp3"
	recognizeOptions := service.NewRecognizeOptionsForMp3(audio)

	// Call the speechToText Recognize method
	response, responseErr := service.Recognize(recognizeOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast recognize.Result to the specific dataType returned by Recognize
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	recognizeResult := service.GetRecognizeResult(response)

	// Check successful casting
	if recognizeResult != nil {
		core.PrettyPrint(recognizeResult, "Recognize")
	}
}
