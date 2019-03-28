package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"os"
)

func main() {
	// Instantiate the Watson Speech To Text service
	service, serviceErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			URL:       "YOUR SERVICE URL",
			IAMApiKey: "YOUR API KEY",
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
	recognizeOptions := service.
		NewRecognizeOptions(audio).
		SetContentType(speechtotextv1.RecognizeOptions_ContentType_AudioMp3)

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

	// Example using websockets!
	audio, _ = os.Open(pwd + "/../../resources/audio_example.mp3")

	// callbook can have `OnOpen`, `onData`, `OnClose` and `onError` functions
	callback := myCallBack{}

	recognizeUsingWebsocketOptions := service.
		NewRecognizeUsingWebsocketOptions(audio, "audio/mp3")

	recognizeUsingWebsocketOptions.
		SetModel("en-US_BroadbandModel").
		SetWordConfidence(true).
		SetSpeakerLabels(true).
		SetTimestamps(true)

	service.RecognizeUsingWebsockets(recognizeUsingWebsocketOptions, callback)
}

type myCallBack struct{}

func (cb myCallBack) OnOpen() {
	fmt.Println("Handshake successful")
}

func (cb myCallBack) OnClose() {
	fmt.Println("Closing connection")
}

func (cb myCallBack) OnData(resp *core.DetailedResponse) {
	var speechResults speechtotextv1.SpeechRecognitionResults
	result := resp.GetResult().([]byte)
	json.Unmarshal(result, &speechResults)
	core.PrettyPrint(speechResults, "Recognized audio: ")
}

func (cb myCallBack) OnError(err error) {
	panic(err)
}
