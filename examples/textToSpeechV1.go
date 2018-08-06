package main

import (
	"fmt"
	. "go-sdk/textToSpeechV1"
	"bytes"
	"os"
)

func main() {
	// Instantiate the Watson Text To Speech service
	textToSpeech, textToSpeechErr := NewTextToSpeechV1(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2017-09-21",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if textToSpeechErr != nil {
		fmt.Println(textToSpeechErr)
		return
	}


	/* SYNTHESIZE */

	synthesizeOptions := NewSynthesizeOptions("Hello World").
		SetAccept("audio/mp3").
		SetVoice("en-GB_KateVoice")

	// Call the textToSpeech Synthesize method
	synthesize, synthesizeErr := textToSpeech.Synthesize(synthesizeOptions)

	// Check successful call
	if synthesizeErr != nil {
		fmt.Println(synthesizeErr)
		return
	}

	// Cast synthesize.Result to the specific dataType returned by Synthesize
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	synthesizeResult := GetSynthesizeResult(synthesize)

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
}
