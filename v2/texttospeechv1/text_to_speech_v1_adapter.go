/**
 * (C) Copyright IBM Corp. 2019, 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package texttospeechv1

import (
	"fmt"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
)

// Timings : An array of words and their start and end times in seconds from the beginning of the synthesized audio.
type Timings struct {
	Words [][]interface{} `json:"words,omitempty"`
}

// Timings : An array of mark times
type Marks struct {
	Marks [][]interface{} `json:"marks"`
}

// AudioContentTypeWrapper : The service sends this message to confirm the audio format
type AudioContentTypeWrapper struct {
	BinaryStreams []struct {
		ContentType string `json:"content_type"`
	} `json:"binary_streams"`
}

// SynthesizeCallbackWrapper : callback for synthesize using websocket
type SynthesizeCallbackWrapper interface {
	OnOpen()
	OnError(error)
	OnContentType(string)
	OnTimingInformation(Timings)
	OnMarks(Marks)
	OnAudioStream([]byte)
	OnData(*core.DetailedResponse)
	OnClose()
}

// SynthesizeOptions : The SynthesizeUsingWebsocket options
type SynthesizeUsingWebsocketOptions struct {
	SynthesizeOptions

	// Callback to listen to events
	Callback SynthesizeCallbackWrapper `json:"callback" validate:"required"`

	// Timings specifies that the service is to return word timing information for all strings of the
	// input text. The service returns the start and end time of each string of the input. Specify words as the lone element
	// of the array to request word timings. Specify an empty array or omit the parameter to receive no word timings. For
	// more information, see [Obtaining word timings](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-timing#timing).
	// Not supported for Japanese input text.
	Timings []string `json:"action,omitempty"`
}

// NewSynthesizeUsingWebsocketOptions: Instantiate SynthesizeOptions to enable websocket support
func (textToSpeech *TextToSpeechV1) NewSynthesizeUsingWebsocketOptions(text string, callback SynthesizeCallbackWrapper) *SynthesizeUsingWebsocketOptions {
	synthesizeOptions := textToSpeech.NewSynthesizeOptions(text)
	synthesizeWSOptions := &SynthesizeUsingWebsocketOptions{*synthesizeOptions, callback, nil}
	return synthesizeWSOptions
}

// SetCallback: Allows user to set the Callback
func (options *SynthesizeUsingWebsocketOptions) SetCallback(callback SynthesizeCallbackWrapper) *SynthesizeUsingWebsocketOptions {
	options.Callback = callback
	return options
}

// SetTimings: Allows user to set the Timings
func (options *SynthesizeUsingWebsocketOptions) SetTimings(timings []string) *SynthesizeUsingWebsocketOptions {
	options.Timings = timings
	return options
}

// SynthesizeUsingWebsocket: Synthesize text over websocket connection
func (textToSpeech *TextToSpeechV1) SynthesizeUsingWebsocket(synthesizeOptions *SynthesizeUsingWebsocketOptions) error {
	if err := core.ValidateNotNil(synthesizeOptions, "synthesizeOptions cannot be nil"); err != nil {
		return err
	}
	if err := core.ValidateStruct(synthesizeOptions, "synthesizeOptions"); err != nil {
		return err
	}

	// Add authentication to the outbound request.
	if textToSpeech.Service.Options.Authenticator == nil {
		return fmt.Errorf("Authentication information was not properly configured.")
	}

	pathSegments := []string{"v1/synthesize"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	dialURL := strings.Replace(textToSpeech.Service.Options.URL, "https", "wss", 1)
	_, err := builder.ConstructHTTPURL(dialURL, pathSegments, pathParameters)
	if err != nil {
		return err
	}

	for headerName, headerValue := range synthesizeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "Synthesize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Content-Type", "application/json")

	if synthesizeOptions.Voice != nil {
		builder.AddQuery("voice", fmt.Sprint(*synthesizeOptions.Voice))
	}
	if synthesizeOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*synthesizeOptions.CustomizationID))
	}

	body := make(map[string]interface{})
	if synthesizeOptions.Text != nil {
		body["text"] = synthesizeOptions.Text
	}
	if synthesizeOptions.Accept != nil {
		body["accept"] = synthesizeOptions.Accept
	}
	if synthesizeOptions.Timings != nil {
		body["timings"] = synthesizeOptions.Timings
	}

	if _, err := builder.SetBodyContentJSON(body); err != nil {
		return err
	}

	request, err := builder.Build()
	if err != nil {
		return err
	}

	// Add the authentication header
	err = textToSpeech.Service.Options.Authenticator.Authenticate(request)
	if err != nil {
		return err
	}

	textToSpeech.NewSynthesizeListener(synthesizeOptions.Callback, request)
	return nil
}
