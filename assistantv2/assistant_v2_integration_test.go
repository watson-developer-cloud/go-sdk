// +build integration

package assistantv2_test

/**
 * (C) Copyright IBM Corp. 2018, 2020.
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

import (
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *assistantv2.AssistantV2
var assistantId string

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load(configFile)
	if err != nil {
		t.Skip(skipMessage)
	}

	assistantId = os.Getenv("ASSISTANT_ID")
	assert.NotEmpty(t, assistantId)
	if assistantId != "" {
		configLoaded = true
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
		Version:     "2020-04-01",
		ServiceName: "assistant",
	})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestSession(t *testing.T) {
	shouldSkipTest(t)

	// Create session
	createSession, _, responseErr := service.CreateSession(
		&assistantv2.CreateSessionOptions{
			AssistantID: core.StringPtr(assistantId),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createSession)

	// Message
	message, _, responseErr := service.Message(
		&assistantv2.MessageOptions{
			AssistantID: core.StringPtr(assistantId),
			SessionID:   createSession.SessionID,
			Input: &assistantv2.MessageInput{
				Text: core.StringPtr("Whats the weather like?"),
			},
			Context: &assistantv2.MessageContext{
				Global: &assistantv2.MessageContextGlobal{
					System: &assistantv2.MessageContextGlobalSystem{
						UserID: core.StringPtr("dummy"),
					},
				},
				Skills: &assistantv2.MessageContextSkills{
					"main_skill": map[string]interface{}{
						"user_defined": map[string]string{
							"account_number": "12345",
						},
					},
				},
			},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, message)

	// Delete session
	_, responseErr = service.DeleteSession(
		&assistantv2.DeleteSessionOptions{
			AssistantID: core.StringPtr(assistantId),
			SessionID:   createSession.SessionID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestMessageStateless(t *testing.T) {
	shouldSkipTest(t)

	message, _, responseErr := service.MessageStateless(
		&assistantv2.MessageStatelessOptions{
			AssistantID: core.StringPtr(assistantId),
			Input: &assistantv2.MessageInputStateless{
				MessageType: core.StringPtr("text"),
				Text:        core.StringPtr("Hello, this is test."),
			},
		},
	)

	assert.Nil(t, responseErr)
	assert.NotNil(t, message)
}

func TestListLogs(t *testing.T) {
	t.Skip("endpoint only available on premium instance")

	options := service.NewListLogsOptions(assistantId)
	logs, _, responseErr := service.ListLogs(options)

	assert.Nil(t, responseErr)
	assert.NotNil(t, logs)
}
