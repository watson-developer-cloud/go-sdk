package speechtotextv1

import (
	"encoding/base64"
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"io"

	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ONE_KB             = 1024
	TEN_MILLISECONDS   = 10 * time.Millisecond
	SUCCESS            = 200
	RECOGNIZE_ENDPOINT = "/v1/recognize"
)

type RecognizeUsingWebsocketOptions struct {
	RecognizeOptions

	// Action that is to be performed. Allowable values: start, stop
	Action *string `json:"action,omitempty"`

	// If true, the service returns interim results as a stream of JSON SpeechRecognitionResults objects.
	// If false, the service returns a single SpeechRecognitionResults object with final results only.
	InterimResults *bool `json:"interim_results,omitempty"`
}

// SetAction: Allows user to set the Action
func (recognizeWSOptions *RecognizeUsingWebsocketOptions) SetAction(action string) *RecognizeUsingWebsocketOptions {
	recognizeWSOptions.Action = core.StringPtr(action)
	return recognizeWSOptions
}

// SetInterimResults: Allows user to set InterimResults
func (recognizeWSOptions *RecognizeUsingWebsocketOptions) SetInterimResults(interimResults bool) *RecognizeUsingWebsocketOptions {
	recognizeWSOptions.InterimResults = core.BoolPtr(interimResults)
	return recognizeWSOptions
}

// NewRecognizeUsingWebsocketOptions: Instantiate RecognizeOptions to enable websocket support
func (speechToText *SpeechToTextV1) NewRecognizeUsingWebsocketOptions(audio io.ReadCloser, contentType string) *RecognizeUsingWebsocketOptions {
	recognizeOptions := speechToText.NewRecognizeOptions(audio)
	recognizeOptions.SetContentType(contentType)
	recognizeWSOptions := &RecognizeUsingWebsocketOptions{*recognizeOptions, nil, nil}
	return recognizeWSOptions
}

type WebsocketRecognitionResults struct {
	// Acknowledges that a start/end message was received, and indicates
	// the start/end of the audio data
	State string `json:"state,omitempty"`

	SpeechRecognitionResults
}

type RecognizeCallbackWrapper interface {
	OnOpen()
	OnClose()
	OnData(*core.DetailedResponse)
	OnError(error)
}

// RecognizeUsingWebsockets: Recognize audio over websocket connection
func (speechToText *SpeechToTextV1) RecognizeUsingWebsockets(recognizeWSOptions *RecognizeUsingWebsocketOptions, callback RecognizeCallbackWrapper) {
	headers := http.Header{}

	if err := core.ValidateNotNil(recognizeWSOptions, "recognizeOptions cannot be nil"); err != nil {
		panic(err)
	}
	if err := core.ValidateStruct(recognizeWSOptions, "recognizeOptions"); err != nil {
		panic(err)
	}

	if speechToText.Service.Options.IAMApiKey != "" || speechToText.Service.TokenManager != nil || speechToText.Service.Options.IAMAccessToken != "" {
		token, err := speechToText.Service.TokenManager.GetToken()
		if err != nil {
			panic(err)
		}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	} else {
		auth := []byte(speechToText.Service.Options.Username + ":" + speechToText.Service.Options.Password)
		headers.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(auth)))
	}

	headers.Set("Content-Type", *recognizeWSOptions.ContentType)

	dialURL := strings.Replace(speechToText.Service.Options.URL, "https", "wss", 1)
	param := url.Values{}

	if recognizeWSOptions.Model != nil {
		param.Set("model", *recognizeWSOptions.Model)
	}
	if recognizeWSOptions.LanguageCustomizationID != nil {
		param.Set("language_customization_id", *recognizeWSOptions.LanguageCustomizationID)
	}
	if recognizeWSOptions.AcousticCustomizationID != nil {
		param.Set("acoustic_customization_id", *recognizeWSOptions.AcousticCustomizationID)
	}
	if recognizeWSOptions.BaseModelVersion != nil {
		param.Set("base_model_version", *recognizeWSOptions.BaseModelVersion)
	}

	speechToText.NewRecognizeListener(callback, recognizeWSOptions, dialURL, param, headers)
}
