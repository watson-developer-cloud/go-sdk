package speechtotextv1

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	core "github.com/watson-developer-cloud/go-sdk/core"
	"io"
	"time"
)

type WebsocketListener struct {
	IsClosed chan bool
	Callback RecognizeCallbackWrapper
}

/*
	OnOpen: Sends start message to server when connection created
*/
func (wsHandle WebsocketListener) OnOpen(recognizeOpt *RecognizeOptions, conn *websocket.Conn) {
	wsHandle.Callback.OnOpen()
	sendStartMessage(conn, recognizeOpt)
}

/*
	OnClose: Callback when websocket connection is closed
*/
func (wsHandle WebsocketListener) OnClose() {
	<-wsHandle.IsClosed
	wsHandle.Callback.OnClose()
}

/*
	OnData: Callback when websocket connection receives data
*/
func (wsHandle WebsocketListener) OnData(conn *websocket.Conn, recognizeOptions *RecognizeOptions) {
	isListening := false
	for {
		var websocketResponse SpeechRecognitionResults
		_, result, err := conn.ReadMessage()
		if err != nil {
			wsHandle.OnError(err)
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

func (wsHandle WebsocketListener) OnError(err error) {
	wsHandle.Callback.OnError(err)
}

/*
	sendStartMessage : Sends start message to server
*/
func sendStartMessage(conn *websocket.Conn, textParams *RecognizeOptions) {
	action := "start"
	textParams.Action = &action
	startMsgBytes, _ := json.Marshal(textParams)
	err := conn.WriteMessage(websocket.TextMessage, startMsgBytes)
	if err != nil {
		textParams.WSListener.OnError(err)
	}
}

/*
	sendCloseMessage : Sends end message to server
*/
func sendCloseMessage(conn *websocket.Conn) {
	stop := "stop"
	closeMsgBytes, _ := json.Marshal(RecognizeOptions{Action: &stop})
	conn.WriteMessage(websocket.TextMessage, closeMsgBytes)
}

/*
	sendAudio : Sends audio data to the server
*/
func sendAudio(conn *websocket.Conn, recognizeOptions *RecognizeOptions) {
	chunk := make([]byte, ONE_KB*2)
	for {
		bytesRead, err := (*recognizeOptions.Audio).Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				recognizeOptions.WSListener.OnError(err)
			}
		}
		err = conn.WriteMessage(websocket.BinaryMessage, chunk[:bytesRead])
		if err != nil {
			recognizeOptions.WSListener.OnError(err)
		}
		time.Sleep(TEN_MILLISECONDS)
	}
	sendCloseMessage(conn)
}
