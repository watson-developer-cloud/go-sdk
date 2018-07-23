package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/toneAnalyzerV3"
)

func main() {
	// Instantiate the Watson Language Translator service
	toneAnalyzer, toneAnalyzerErr := toneAnalyzerV3.NewToneAnalyzerV3(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2017-09-21",
		APIkey: "YOUR API KEY",
	})

	// Check successful instantiation
	if toneAnalyzerErr != nil {
		fmt.Println(toneAnalyzer)
		return
	}


	//Tone Analyzer
	/* TONE CHAT */

	utterances := toneAnalyzerV3.ToneChatInput{
		[]toneAnalyzerV3.Utterance{
			{"Hello World", "Watson"},
			{"World Hello", "John Doe"},
			},
	}

	// Call the toneAnalyzer Tone Chat method
	toneChat, toneChatErr := toneAnalyzer.ToneChat(&utterances, "en", "en")

	// Check successful call
	if toneChatErr != nil {
		fmt.Println(toneChatErr)
		return
	}

	// Cast response from call to the specific struct returned by GetLToneChatResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	toneChatResult := toneAnalyzerV3.GetToneChatResult(toneChat)

	// Check successful casting
	if toneChatResult != nil {
		// Print result
		fmt.Println(toneChatResult)
	}
}
