package main

import (
	"fmt"
	. "go-sdk/speechToTextV1"
	"os"
	"encoding/json"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Speech To Text service
	speechToText, speechToTextErr := NewSpeechToTextV1(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2017-09-21",
		APIkey: "YOUR SERVICE API KEY",
	})

	// Check successful instantiation
	if speechToTextErr != nil {
		fmt.Println(speechToTextErr)
		return
	}


	/* RECOGNIZE */

	pwd, _ := os.Getwd()

	// Open file with mp3 to recognize
	audio, audioErr := os.Open(pwd + "/resources/audio_example.mp3")
	if audioErr != nil {
		fmt.Println(audioErr)
		return
	}

	// Create a new RecognizeOptions for ContentType "audio/mp3"
	recognizeOptions := NewRecognizeOptionsForMp3(audio)

	// Call the speechToText Recognize method
	recognize, recognizeErr := speechToText.Recognize(recognizeOptions)

	// Check successful call
	if recognizeErr != nil {
		fmt.Println(recognizeErr)
		return
	}

	// Cast recognize.Result to the specific dataType returned by Recognize
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	recognizeResult := GetRecognizeResult(recognize)

	// Check successful casting
	if recognizeResult != nil {
		prettyPrint(recognizeResult, "Recognize")
	}
}
