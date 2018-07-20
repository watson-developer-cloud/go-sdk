package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/languageTranslatorV3"
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


	/* TONE */

	// Call the toneAnalyzer Tone method
	tone, toneErr := toneAnalyzer.Tone("YOUR BODY", "YOUR CONTENT TYPE", true, true)

	// Check successful call
	if toneErr != nil {
		fmt.Println(toneErr)
		return
	}

	// Cast response from call to the specific struct returned by GetToneResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	toneResult := toneAnalyzerV3.GetToneResult(tone)

	// Check successful casting
	if toneResult != nil {
		// Print result
		fmt.Println(toneResult)
	}


	/* TONE CHAT */

	// Call the toneAnalyzer Tone Chat method
	toneChat, toneChatErr := toneAnalyzer.ToneChat("YOUR BODY", "YOUR CONTENT LANGUAGE", true,)

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
