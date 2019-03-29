// +build integration

package assistantv2_test

/**
 * Copyright 2018 IBM All Rights Reserved.
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
	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
	"os"
	"testing"
)

var service *assistantv2.AssistantV2
var serviceErr error

func init() {
	err := godotenv.Load("../.env")

	if err == nil {
		service, serviceErr = assistantv2.
			NewAssistantV2(&assistantv2.AssistantV2Options{
				URL:      os.Getenv("ASSISTANT_URL"),
				Version:  "2017-04-21",
				Username: os.Getenv("ASSISTANT_USERNAME"),
				Password: os.Getenv("ASSISTANT_PASSWORD"),
			})
	}
}

func shouldSkipTest(t *testing.T) {
	if service == nil {
		t.Skip("Skipping test as service credentials are missing")
	}
}

func TestSession(t *testing.T) {
	shouldSkipTest(t)

	// Create session
	response, responseErr := service.CreateSession(
		&assistantv2.CreateSessionOptions{
			AssistantID: core.StringPtr(os.Getenv("ASSISTANT_ID")),
		},
	)
	assert.Nil(t, responseErr)

	createSession := service.GetCreateSessionResult(response)
	assert.NotNil(t, createSession)

	// Message
	response, responseErr = service.Message(
		&assistantv2.MessageOptions{
			AssistantID: core.StringPtr(os.Getenv("ASSISTANT_ID")),
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

	message := service.GetMessageResult(response)
	assert.NotNil(t, message)

	// Delete session
	response, responseErr = service.DeleteSession(
		&assistantv2.DeleteSessionOptions{
			AssistantID: core.StringPtr(os.Getenv("ASSISTANT_ID")),
			SessionID:   createSession.SessionID,
		},
	)
	assert.Nil(t, responseErr)
}
