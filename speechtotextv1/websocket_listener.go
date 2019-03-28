package speechtotextv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RecognizeListener struct {
	IsClosed chan bool
	Callback RecognizeCallbackWrapper
}

/*
	OnOpen: Sends start message to server when connection created
*/
func (wsHandle RecognizeListener) OnOpen(recognizeOpt *RecognizeUsingWebsocketOptions, conn *websocket.Conn) {
	wsHandle.Callback.OnOpen()
	sendStartMessage(conn, recognizeOpt, &wsHandle)
}

/*
	OnClose: Callback when websocket connection is closed
*/
func (wsHandle RecognizeListener) OnClose() {
	<-wsHandle.IsClosed
	wsHandle.Callback.OnClose()
}

/*
	OnData: Callback when websocket connection receives data
*/
func (wsHandle RecognizeListener) OnData(conn *websocket.Conn, recognizeOptions *RecognizeUsingWebsocketOptions) {
	isListening := false
	for {
		var websocketResponse WebsocketRecognitionResults
		_, result, err := conn.ReadMessage()
		if err != nil {
			wsHandle.OnError(err)
			break
		}
		json.Unmarshal(result, &websocketResponse)

		if websocketResponse.State == "listening" {
			if !isListening {
				isListening = true
				continue
			} else {
				break
			}
		}
		detailResp := core.DetailedResponse{}
		detailResp.Result = result
		detailResp.StatusCode = SUCCESS
		wsHandle.Callback.OnData(&detailResp)
	}
	conn.Close()
	wsHandle.IsClosed <- true
}

/*
	OnError: Callback when error encountered
*/
func (wsHandle RecognizeListener) OnError(err error) {
	wsHandle.Callback.OnError(err)
}

/*
	sendStartMessage : Sends start message to server
*/
func sendStartMessage(conn *websocket.Conn, textParams *RecognizeUsingWebsocketOptions, recognizeListener *RecognizeListener) {
	action := "start"
	textParams.Action = &action
	startMsgBytes, _ := json.Marshal(textParams)
	err := conn.WriteMessage(websocket.TextMessage, startMsgBytes)
	if err != nil {
		recognizeListener.OnError(err)
	}
}

/*
	sendCloseMessage : Sends end message to server
*/
func sendCloseMessage(conn *websocket.Conn) {
	stop := "stop"
	closeMsgBytes, _ := json.Marshal(RecognizeUsingWebsocketOptions{Action: &stop})
	conn.WriteMessage(websocket.TextMessage, closeMsgBytes)
}

/*
	sendAudio : Sends audio data to the server
*/
func sendAudio(conn *websocket.Conn, recognizeOptions *RecognizeUsingWebsocketOptions, recognizeListener *RecognizeListener) {
	chunk := make([]byte, ONE_KB*2)
	for {
		bytesRead, err := (*recognizeOptions.Audio).Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				recognizeListener.OnError(err)
			}
		}
		err = conn.WriteMessage(websocket.BinaryMessage, chunk[:bytesRead])
		if err != nil {
			recognizeListener.OnError(err)
		}
		time.Sleep(TEN_MILLISECONDS)
	}
	sendCloseMessage(conn)
}

/*
	NewRecognizeListener : Instantiates a listener instance to control the sending/receiving of audio/text
*/
func (speechToText *SpeechToTextV1) NewRecognizeListener(callback RecognizeCallbackWrapper, recognizeWSOptions *RecognizeUsingWebsocketOptions, dialURL string, param url.Values, headers http.Header) {
	recognizeListener := RecognizeListener{Callback: callback, IsClosed: make(chan bool, 1)}
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("%s%s?%s", dialURL, RECOGNIZE_ENDPOINT, param.Encode()), headers)
	if err != nil {
		recognizeListener.OnError(err)
	}
	recognizeListener.OnOpen(recognizeWSOptions, conn)
	go recognizeListener.OnData(conn, recognizeWSOptions)
	go sendAudio(conn, recognizeWSOptions, &recognizeListener)
	recognizeListener.OnClose()
}
