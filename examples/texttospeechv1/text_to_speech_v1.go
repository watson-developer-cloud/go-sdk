package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
)

func main() {
	// Instantiate the Watson Text To Speech service
	service, serviceErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			URL:       "YOUR SERVICE URL",
			IAMApiKey: "YOUR SERVICE API KEY",
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* SYNTHESIZE */

	synthesizeOptions := service.NewSynthesizeOptions("Hello World").
		SetAccept("audio/mp3").
		SetVoice("en-GB_KateVoice")

	// Call the textToSpeech Synthesize method
	response, responseErr := service.Synthesize(synthesizeOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast synthesize.Result to the specific dataType returned by Synthesize
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	synthesizeResult := service.GetSynthesizeResult(response)

	// Check successful casting
	if synthesizeResult != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(synthesizeResult)

		fileName := "synthesize_example_output.mp3"
		file, _ := os.Create(fileName)
		file.Write(buff.Bytes())
		file.Close()

		fmt.Println("Wrote synthesized text to " + fileName)
	}
	synthesizeResult.Close()
}
