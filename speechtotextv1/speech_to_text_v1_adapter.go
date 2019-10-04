package speechtotextv1

import (
	"fmt"
	"io"

	"github.com/IBM/go-sdk-core/core"

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

	// If `true`, requests processing metrics about the service's transcription of the input audio. The service returns
	// processing metrics at the interval specified by the `processing_metrics_interval` parameter. It also returns
	// processing metrics for transcription events, for example, for final and interim results. By default, the service
	// returns no processing metrics.
	ProcessingMetrics *bool `json:"processing_metrics,omitempty"`

	// Specifies the interval in real wall-clock seconds at which the service is to return processing metrics. The
	// parameter is ignored unless the `processing_metrics` parameter is set to `true`.
	//
	// The parameter accepts a minimum value of 0.1 seconds. The level of precision is not restricted, so you can specify
	// values such as 0.25 and 0.125.
	//
	// The service does not impose a maximum value. If you want to receive processing metrics only for transcription events
	// instead of at periodic intervals, set the value to a large number. If the value is larger than the duration of the
	// audio, the service returns processing metrics only for transcription events.
	ProcessingMetricsInterval *float32 `json:"processing_metrics_interval,omitempty"`
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

// SetProcessingMetrics : Allow user to set ProcessingMetrics
func (recognizeWSOptions *RecognizeUsingWebsocketOptions) SetProcessingMetrics(processingMetrics bool) *RecognizeUsingWebsocketOptions {
	recognizeWSOptions.ProcessingMetrics = core.BoolPtr(processingMetrics)
	return recognizeWSOptions
}

// SetProcessingMetricsInterval : Allow user to set ProcessingMetricsInterval
func (recognizeWSOptions *RecognizeUsingWebsocketOptions) SetProcessingMetricsInterval(processingMetricsInterval float32) *RecognizeUsingWebsocketOptions {
	recognizeWSOptions.ProcessingMetricsInterval = core.Float32Ptr(processingMetricsInterval)
	return recognizeWSOptions
}

// NewRecognizeUsingWebsocketOptions: Instantiate RecognizeOptions to enable websocket support
func (speechToText *SpeechToTextV1) NewRecognizeUsingWebsocketOptions(audio io.ReadCloser, contentType string) *RecognizeUsingWebsocketOptions {
	recognizeOptions := speechToText.NewRecognizeOptions(audio)
	recognizeOptions.SetContentType(contentType)
	recognizeWSOptions := &RecognizeUsingWebsocketOptions{*recognizeOptions, nil, nil, nil, nil}
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

// RecognizeUsingWebsocket: Recognize audio over websocket connection
func (speechToText *SpeechToTextV1) RecognizeUsingWebsocket(recognizeWSOptions *RecognizeUsingWebsocketOptions, callback RecognizeCallbackWrapper) {
	if err := core.ValidateNotNil(recognizeWSOptions, "recognizeOptions cannot be nil"); err != nil {
		panic(err)
	}
	if err := core.ValidateStruct(recognizeWSOptions, "recognizeOptions"); err != nil {
		panic(err)
	}

	// Add authentication to the outbound request.
	if speechToText.Service.Options.Authenticator == nil {
		panic(fmt.Errorf("Authentication information was not properly configured."))
	}

	// Create a dummy request for authenticate
	// Need to update design to let recognizeListener take in a request object
	req, err := http.NewRequest("POST", speechToText.Service.Options.URL, nil)
	err = speechToText.Service.Options.Authenticator.Authenticate(req)
	if err != nil {
		panic(err)
	}
	headers := req.Header

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
