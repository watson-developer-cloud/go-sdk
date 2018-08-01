package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/toneanalyzerv3"
	"encoding/json"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Tone Analyzer service
	toneAnalyzer, toneAnalyzerErr := toneanalyzerv3.NewToneAnalyzerV3(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2017-09-21",
		Username: "YOUR SERVICE USERNAME",
		Password: "YOUR SERVICE PASSWORD",
	})

	// Check successful instantiation
	if toneAnalyzerErr != nil {
		fmt.Println(toneAnalyzer)
		return
	}

	/* TONE CHAT */

	utterances := []toneanalyzerv3.Utterance{
		{
			Text: "Hello World",
			User: "Watson",
		},
		{
			Text: "World Hello",
			User: "John Doe",
		},
	}

	toneChatOptions := toneanalyzerv3.NewToneChatOptions(utterances).
		SetAcceptLanguage("en").
		SetContentLanguage("en")

	// Call the toneAnalyzer ToneChat method
	toneChat, toneChatErr := toneAnalyzer.ToneChat(toneChatOptions)

	// Check successful call
	if toneChatErr != nil {
		fmt.Println(toneChatErr)
		return
	}

	// Cast response from call to the specific struct returned by GetToneChatResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	toneChatResult := toneanalyzerv3.GetToneChatResult(toneChat)

	// Check successful casting
	if toneChatResult != nil {
		prettyPrint(toneChatResult, "Tone Chat")
	}
}
