/**
 * (C) Copyright IBM Corp. 2019.
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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/gorilla/websocket"
)

const (
	SUCCESS = 200
)

type SynthesizeListener struct {
	IsClosed chan bool
	Callback SynthesizeCallbackWrapper
}

func decode(b []byte, target interface{}) {
	_ = json.NewDecoder(bytes.NewReader(b)).Decode(&target)
}

// OnError: Callback when error encountered
func (listener SynthesizeListener) OnError(err error) {
	listener.Callback.OnError(err)
}

// SendText: Sends the text message
// Note: The service handles one request per connection
func (listener SynthesizeListener) SendText(conn *websocket.Conn, req *http.Request) {
	listener.OnOpen(conn)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		listener.OnError(err)
	}
	err = conn.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		listener.OnError(err)
	}
}

// OnOpen: Sends start message to server when connection created
func (listener SynthesizeListener) OnOpen(conn *websocket.Conn) {
	listener.Callback.OnOpen()
}

// OnClose: Callback when websocket connection is closed
func (listener SynthesizeListener) OnClose() {
	<-listener.IsClosed
	listener.Callback.OnClose()
}

// OnData: Callback when websocket connection receives data
func (listener SynthesizeListener) OnData(conn *websocket.Conn) {
	for {
		messageType, result, err := conn.ReadMessage()

		// The service will close the connection. We need to decipher
		// if the error is a normal close signal
		if err != nil {
			listener.IsClosed <- true

			if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				listener.OnError(err)
				conn.Close()
			}
			break
		}

		if messageType == websocket.TextMessage {
			var r map[string]interface{}
			err = json.NewDecoder(bytes.NewReader(result)).Decode(&r)
			if err != nil {
				listener.OnError(fmt.Errorf(err.Error()))
				listener.IsClosed <- true
				break
			}
			if err, ok := r["error"]; ok {
				listener.OnError(fmt.Errorf(err.(string)))
				listener.IsClosed <- true
				break
			}

			if _, ok := r["binary_streams"]; ok {
				audioContentTypeWrapper := new(AudioContentTypeWrapper)
				decode(result, audioContentTypeWrapper)
				listener.Callback.OnContentType(audioContentTypeWrapper.BinaryStreams[0].ContentType)
			} else if _, ok := r["words"]; ok {
				timings := new(Timings)
				decode(result, timings)
				listener.Callback.OnTimingInformation(*timings)
			} else if _, ok := r["marks"]; ok {
				marks := new(Marks)
				decode(result, marks)
				listener.Callback.OnMarks(*marks)
			}
		} else if messageType == websocket.BinaryMessage {
			listener.Callback.OnAudioStream(result)
		}

		detailResponse := core.DetailedResponse{}
		detailResponse.Result = result
		detailResponse.StatusCode = SUCCESS
		listener.Callback.OnData(&detailResponse)
	}
}

func (textToSpeechV1 *TextToSpeechV1) NewSynthesizeListener(callback SynthesizeCallbackWrapper, req *http.Request) {
	synthesizeListener := SynthesizeListener{Callback: callback, IsClosed: make(chan bool, 1)}
	conn, _, err := websocket.DefaultDialer.Dial(req.URL.String(), req.Header)
	if err != nil {
		synthesizeListener.OnError(err)
	}

	go synthesizeListener.OnData(conn)
	go synthesizeListener.SendText(conn, req)
	synthesizeListener.OnClose()
}
