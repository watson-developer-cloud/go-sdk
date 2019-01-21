package speechtotextv1

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/websocket"
	core "github.com/watson-developer-cloud/go-sdk/core"

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

type RecognizeCallbackWrapper interface {
	OnOpen()
	OnClose()
	OnData(*core.DetailedResponse)
	OnError(error)
}

// RecognizeUsingWebsockets: Recognize audio over websocket connection
func (speechToText *SpeechToTextV1) RecognizeUsingWebsockets(recognizeOptions *RecognizeOptions, wsListener websocketListener) {
	var token string
	headers := http.Header{}
	finish := make(chan bool, 1)

	if err := core.ValidateNotNil(recognizeOptions, "recognizeOptions cannot be nil"); err != nil {
		wsListener.Callback.OnError(err)
	}
	if err := core.ValidateNotNil(recognizeOptions.AudioMetaData, "Audio metadata cannot be nil"); err != nil {
		wsListener.Callback.OnError(err)
	}
	if err := core.ValidateStruct(recognizeOptions, "recognizeOptions"); err != nil {
		wsListener.Callback.OnError(err)
	}

	if speechToText.Service.Options.IAMApiKey != "" || speechToText.Service.TokenManager != nil {
		token = speechToText.Service.TokenManager.GetToken()
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	} else {
		auth := []byte(speechToText.Service.Options.Username + ":" + speechToText.Service.Options.Password)
		headers.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(auth)))
	}

	headers.Set("Content-Type", *recognizeOptions.ContentType)

	dialURL := strings.Replace(speechToText.Service.Options.URL, "https", "wss", 1)
	param := url.Values{}

	if recognizeOptions.Model != nil {
		param.Set("model", *recognizeOptions.Model)
	}
	if recognizeOptions.LanguageCustomizationID != nil {
		param.Set("language_customization_id", *recognizeOptions.LanguageCustomizationID)
	}
	if recognizeOptions.AcousticCustomizationID != nil {
		param.Set("acoustic_customization_id", *recognizeOptions.AcousticCustomizationID)
	}
	if recognizeOptions.BaseModelVersion != nil {
		param.Set("base_model_version", *recognizeOptions.BaseModelVersion)
	}

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("%s%s?%s", dialURL, RECOGNIZE_ENDPOINT, param.Encode()), headers)
	if err != nil {
		panic(err)
	}

	go wsListener.OnOpen(recognizeOptions)
	go wsListener.OnClose(finish)
	go wsListener.OnData(conn, recognizeOptions)
	wsListener.WSOpenChan <- conn
	<-finish
}
