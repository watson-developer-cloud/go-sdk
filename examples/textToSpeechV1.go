package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/texttospeechv1"
	"bytes"
	"os"
)

func main() {
	// Instantiate the Watson Text To Speech service
	textToSpeech, textToSpeechErr := texttospeechv1.NewTextToSpeechV1(watson.Credentials{
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

	synthesizeOptions := texttospeechv1.NewSynthesizeOptions("Hello World").
		SetAccept("audio/mp3").
		SetVoice("en-GB_KateVoice")

	// Call the textToSpeech Synthesize method
	synthesize, synthesizeErr := textToSpeech.Synthesize(synthesizeOptions)

	// Check successful call
	if synthesizeErr != nil {
		fmt.Println(synthesizeErr)
		return
	}

	// Cast response from call to the specific struct returned by GetSynthesizeResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	synthesizeResult := texttospeechv1.GetSynthesizeResult(synthesize)

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
