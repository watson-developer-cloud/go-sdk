package main

import (
	"io/ioutil"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/toneanalyzerv3"
)

func main() {
	// Instantiate the Watson Tone Analyzer service
		authenticator := &core.IamAuthenticator{
        	ApiKey:     os.Getenv("YOUR API KEY"),
    	}
	service, serviceErr := toneanalyzerv3.
		NewToneAnalyzerV3(&toneanalyzerv3.ToneAnalyzerV3Options{
			URL:       "YOUR SERVICE URL",
			Version:   "2017-09-21",
			Authenticator: authenticator,
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* TONE CHAT */

	utterances := []toneanalyzerv3.Utterance{
		toneanalyzerv3.Utterance{
			Text: core.StringPtr("Hello World"),
			User: core.StringPtr("Watson"),
		},
		toneanalyzerv3.Utterance{
			Text: core.StringPtr("World Hello"),
			User: core.StringPtr("John Doe"),
		},
	}

	toneChatOptions := service.NewToneChatOptions(utterances).
		SetAcceptLanguage("en").
		SetContentLanguage("en")

	// Call the toneAnalyzer ToneChat method
	toneChatResult, _, responseErr := service.ToneChat(toneChatOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if toneChatResult != nil {
		core.PrettyPrint(toneChatResult, "Tone Chat")
	}

	/* TONE */

	// Call the toneAnalyzer Tone method
	toneOptions := service.NewToneOptions().
		SetBody("I am very happy. It is a good day").
		SetContentType("text/plain")
	toneResult, response, responseErr := service.Tone(toneOptions)

	if responseErr != nil {
		panic(responseErr)
	}

	if toneResult != nil {
		core.PrettyPrint(toneResult, "Tone using plain text")
	}

	// Call the toneAnalyzer Tone method
	toneInput := &toneanalyzerv3.ToneInput{
		Text: core.StringPtr("Team, I know that times are tough! Product sales have been disappointing for the past three quarters. We have a competitive product, but we need to do a better job of selling it!"),
	}
	toneOptions = service.
		NewToneOptions().
		SetToneInput(toneInput).
		SetContentType("application/json")
	toneResult, response, responseErr = service.Tone(toneOptions)

	if responseErr != nil {
		panic(responseErr)
	}

	if toneResult != nil {
		core.PrettyPrint(toneResult, "Tone using toneInput")
	}

	// Call the toneAnalyzer Tone method
	pwd, _ := os.Getwd()
	htmlByte, htmlByteErr := ioutil.ReadFile(pwd + "/../../resources/tone-example.html")
	if htmlByteErr != nil {
		panic(htmlByteErr)
	}
	toneOptions = service.NewToneOptions().
		SetBody(string(htmlByte)).
		SetContentType("text/html")
	toneResult, response, responseErr = service.Tone(toneOptions)

	if responseErr != nil {
		panic(responseErr)
	}

	if toneResult != nil {
		core.PrettyPrint(toneResult, "Tone using toneInput")
	}

}
