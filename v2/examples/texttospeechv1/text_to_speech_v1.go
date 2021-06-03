package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/watson-developer-cloud/go-sdk/v2/texttospeechv1"
)

func main() {
	// Instantiate the Watson Text To Speech service
	authenticator := &core.IamAuthenticator{
		ApiKey: "YOUR SERVICE API KEY",
	}

	service, serviceErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			Authenticator: authenticator,
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
	synthesizeResult, _, responseErr := service.Synthesize(synthesizeOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

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

	/* SYNTHESIZE USING WEBSOCKET*/
	// create a file for websocket output
	fileName := "synthesize_ws_example_output.mp3"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	callback := myCallBack{f: file}

	synthesizeUsingWebsocketOptions := service.
		NewSynthesizeUsingWebsocketOptions("This is a <mark name=\"SIMPLE\"/>simple <mark name=\"EXAMPLE\"/> example.", callback)

	synthesizeUsingWebsocketOptions.
		SetAccept("audio/mp3").
		SetVoice("en-US_AllisonVoice")
	synthesizeUsingWebsocketOptions.SetTimings([]string{"words"})
	err = service.SynthesizeUsingWebsocket(synthesizeUsingWebsocketOptions)
	if err != nil {
		fmt.Println(err)
	}
}

type myCallBack struct {
	f *os.File
}

func (cb myCallBack) OnOpen() {
	fmt.Println("Handshake successful")
}

func (cb myCallBack) OnClose() {
	fmt.Println("Closing connection")
	cb.f.Close()
}

func (cb myCallBack) OnAudioStream(b []byte) {
	bytes, err := ioutil.ReadAll(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	_, err = cb.f.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func (cb myCallBack) OnData(response *core.DetailedResponse) {}

func (cb myCallBack) OnError(err error) {
	fmt.Println("Received error")
	panic(err)
}

func (cb myCallBack) OnTimingInformation(timings texttospeechv1.Timings) {
	core.PrettyPrint(timings, "Timing information: ")
}

func (cb myCallBack) OnMarks(marks texttospeechv1.Marks) {
	core.PrettyPrint(marks, "Mark timings: ")
}

func (cb myCallBack) OnContentType(contentType string) {
	fmt.Println("The content type identified is:", contentType)
}
