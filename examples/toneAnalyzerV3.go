package main

import (
	"fmt"
	. "go-sdk/toneAnalyzerV3"
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
	toneAnalyzer, toneAnalyzerErr := NewToneAnalyzerV3(&ServiceCredentials{
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

	utterances := []Utterance{
		{
			Text: "Hello World",
			User: "Watson",
		},
		{
			Text: "World Hello",
			User: "John Doe",
		},
	}

	toneChatOptions := NewToneChatOptions(utterances).
		SetAcceptLanguage("en").
		SetContentLanguage("en")

	// Call the toneAnalyzer ToneChat method
	toneChat, toneChatErr := toneAnalyzer.ToneChat(toneChatOptions)

	// Check successful call
	if toneChatErr != nil {
		fmt.Println(toneChatErr)
		return
	}

	// Cast toneChat.Result to the specific dataType returned by ToneChat
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	toneChatResult := GetToneChatResult(toneChat)

	// Check successful casting
	if toneChatResult != nil {
		prettyPrint(toneChatResult, "Tone Chat")
	}
}
