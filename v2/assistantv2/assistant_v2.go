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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.22.0-937b9a1c-20201211-223043
 */

// Package assistantv2 : Operations and models for the AssistantV2 service
package assistantv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/watson-developer-cloud/go-sdk/v2/common"
)

// AssistantV2 : The IBM Watson&trade; Assistant service combines machine learning, natural language understanding, and
// an integrated dialog editor to create conversation flows between your apps and your users.
//
// The Assistant v2 API provides runtime methods your client application can use to send user input to an assistant and
// receive a response.
//
// Version: 2.0
// See: https://cloud.ibm.com/docs/assistant
type AssistantV2 struct {
	Service *core.BaseService

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2020-04-01`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.assistant.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "conversation"

// AssistantV2Options : Service options
type AssistantV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the API version you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2020-04-01`.
	Version *string `validate:"required"`
}

// NewAssistantV2 : constructs an instance of AssistantV2 with passed in options.
func NewAssistantV2(options *AssistantV2Options) (service *AssistantV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	err = baseService.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &AssistantV2{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "assistant" suitable for processing requests.
func (assistant *AssistantV2) Clone() *AssistantV2 {
	if core.IsNil(assistant) {
		return nil
	}
	clone := *assistant
	clone.Service = assistant.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (assistant *AssistantV2) SetServiceURL(url string) error {
	return assistant.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (assistant *AssistantV2) GetServiceURL() string {
	return assistant.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (assistant *AssistantV2) SetDefaultHeaders(headers http.Header) {
	assistant.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (assistant *AssistantV2) SetEnableGzipCompression(enableGzip bool) {
	assistant.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (assistant *AssistantV2) GetEnableGzipCompression() bool {
	return assistant.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (assistant *AssistantV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	assistant.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (assistant *AssistantV2) DisableRetries() {
	assistant.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (assistant *AssistantV2) DisableSSLVerification() {
	assistant.Service.DisableSSLVerification()
}

// CreateSession : Create a session
// Create a new session. A session is used to send user input to a skill and receive responses. It also maintains the
// state of the conversation. A session persists until it is deleted, or until it times out because of inactivity. (For
// more information, see the [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-settings).
func (assistant *AssistantV2) CreateSession(createSessionOptions *CreateSessionOptions) (result *SessionResponse, response *core.DetailedResponse, err error) {
	return assistant.CreateSessionWithContext(context.Background(), createSessionOptions)
}

// CreateSessionWithContext is an alternate form of the CreateSession method which supports a Context parameter
func (assistant *AssistantV2) CreateSessionWithContext(ctx context.Context, createSessionOptions *CreateSessionOptions) (result *SessionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSessionOptions, "createSessionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSessionOptions, "createSessionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assistant_id": *createSessionOptions.AssistantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/assistants/{assistant_id}/sessions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSessionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "CreateSession")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSessionResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteSession : Delete session
// Deletes a session explicitly before it times out. (For more information about the session inactivity timeout, see the
// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-settings)).
func (assistant *AssistantV2) DeleteSession(deleteSessionOptions *DeleteSessionOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteSessionWithContext(context.Background(), deleteSessionOptions)
}

// DeleteSessionWithContext is an alternate form of the DeleteSession method which supports a Context parameter
func (assistant *AssistantV2) DeleteSessionWithContext(ctx context.Context, deleteSessionOptions *DeleteSessionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSessionOptions, "deleteSessionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSessionOptions, "deleteSessionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assistant_id": *deleteSessionOptions.AssistantID,
		"session_id":   *deleteSessionOptions.SessionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/assistants/{assistant_id}/sessions/{session_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSessionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "DeleteSession")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// Message : Send user input to assistant (stateful)
// Send user input to an assistant and receive a response, with conversation state (including context data) stored by
// Watson Assistant for the duration of the session.
func (assistant *AssistantV2) Message(messageOptions *MessageOptions) (result *MessageResponse, response *core.DetailedResponse, err error) {
	return assistant.MessageWithContext(context.Background(), messageOptions)
}

// MessageWithContext is an alternate form of the Message method which supports a Context parameter
func (assistant *AssistantV2) MessageWithContext(ctx context.Context, messageOptions *MessageOptions) (result *MessageResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(messageOptions, "messageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(messageOptions, "messageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assistant_id": *messageOptions.AssistantID,
		"session_id":   *messageOptions.SessionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/assistants/{assistant_id}/sessions/{session_id}/message`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range messageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "Message")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	body := make(map[string]interface{})
	if messageOptions.Input != nil {
		body["input"] = messageOptions.Input
	}
	if messageOptions.Context != nil {
		body["context"] = messageOptions.Context
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMessageResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// MessageStateless : Send user input to assistant (stateless)
// Send user input to an assistant and receive a response, with conversation state (including context data) managed by
// your application.
func (assistant *AssistantV2) MessageStateless(messageStatelessOptions *MessageStatelessOptions) (result *MessageResponseStateless, response *core.DetailedResponse, err error) {
	return assistant.MessageStatelessWithContext(context.Background(), messageStatelessOptions)
}

// MessageStatelessWithContext is an alternate form of the MessageStateless method which supports a Context parameter
func (assistant *AssistantV2) MessageStatelessWithContext(ctx context.Context, messageStatelessOptions *MessageStatelessOptions) (result *MessageResponseStateless, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(messageStatelessOptions, "messageStatelessOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(messageStatelessOptions, "messageStatelessOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assistant_id": *messageStatelessOptions.AssistantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/assistants/{assistant_id}/message`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range messageStatelessOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "MessageStateless")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	body := make(map[string]interface{})
	if messageStatelessOptions.Input != nil {
		body["input"] = messageStatelessOptions.Input
	}
	if messageStatelessOptions.Context != nil {
		body["context"] = messageStatelessOptions.Context
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMessageResponseStateless)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// BulkClassify : Identify intents and entities in multiple user utterances
// Send multiple user inputs to a dialog skill in a single request and receive information about the intents and
// entities recognized in each input. This method is useful for testing and comparing the performance of different
// skills or skill versions.
//
// This method is available only with Premium plans.
func (assistant *AssistantV2) BulkClassify(bulkClassifyOptions *BulkClassifyOptions) (result *BulkClassifyResponse, response *core.DetailedResponse, err error) {
	return assistant.BulkClassifyWithContext(context.Background(), bulkClassifyOptions)
}

// BulkClassifyWithContext is an alternate form of the BulkClassify method which supports a Context parameter
func (assistant *AssistantV2) BulkClassifyWithContext(ctx context.Context, bulkClassifyOptions *BulkClassifyOptions) (result *BulkClassifyResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(bulkClassifyOptions, "bulkClassifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(bulkClassifyOptions, "bulkClassifyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"skill_id": *bulkClassifyOptions.SkillID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/skills/{skill_id}/workspace/bulk_classify`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range bulkClassifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "BulkClassify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))

	body := make(map[string]interface{})
	if bulkClassifyOptions.Input != nil {
		body["input"] = bulkClassifyOptions.Input
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBulkClassifyResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListLogs : List log events for an assistant
// List the events from the log of an assistant.
//
// This method is available only with Premium plans.
func (assistant *AssistantV2) ListLogs(listLogsOptions *ListLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	return assistant.ListLogsWithContext(context.Background(), listLogsOptions)
}

// ListLogsWithContext is an alternate form of the ListLogs method which supports a Context parameter
func (assistant *AssistantV2) ListLogsWithContext(ctx context.Context, listLogsOptions *ListLogsOptions) (result *LogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLogsOptions, "listLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLogsOptions, "listLogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assistant_id": *listLogsOptions.AssistantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/assistants/{assistant_id}/logs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "ListLogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	if listLogsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listLogsOptions.Sort))
	}
	if listLogsOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*listLogsOptions.Filter))
	}
	if listLogsOptions.PageLimit != nil {
		builder.AddQuery("page_limit", fmt.Sprint(*listLogsOptions.PageLimit))
	}
	if listLogsOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*listLogsOptions.Cursor))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = assistant.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogCollection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data associated with a specified customer ID. The method has no effect if no data is associated with the
// customer ID.
//
// You associate a customer ID with data by passing the `X-Watson-Metadata` header with a request that passes data. For
// more information about personal data and customer IDs, see [Information
// security](https://cloud.ibm.com/docs/assistant?topic=assistant-information-security#information-security).
//
// This operation is limited to 4 requests per minute. For more information, see **Rate limiting**.
func (assistant *AssistantV2) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return assistant.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (assistant *AssistantV2) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = assistant.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(assistant.Service.Options.URL, `/v2/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("conversation", "V2", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*assistant.Version))
	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, nil)

	return
}

// AgentAvailabilityMessage : AgentAvailabilityMessage struct
type AgentAvailabilityMessage struct {
	// The text of the message.
	Message *string `json:"message,omitempty"`
}

// UnmarshalAgentAvailabilityMessage unmarshals an instance of AgentAvailabilityMessage from the specified map of raw messages.
func UnmarshalAgentAvailabilityMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AgentAvailabilityMessage)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BulkClassifyOptions : The BulkClassify options.
type BulkClassifyOptions struct {
	// Unique identifier of the skill. To find the skill ID in the Watson Assistant user interface, open the skill settings
	// and click **API Details**.
	SkillID *string `json:"skill_id" validate:"required,ne="`

	// An array of input utterances to classify.
	Input []BulkClassifyUtterance `json:"input,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBulkClassifyOptions : Instantiate BulkClassifyOptions
func (*AssistantV2) NewBulkClassifyOptions(skillID string) *BulkClassifyOptions {
	return &BulkClassifyOptions{
		SkillID: core.StringPtr(skillID),
	}
}

// SetSkillID : Allow user to set SkillID
func (options *BulkClassifyOptions) SetSkillID(skillID string) *BulkClassifyOptions {
	options.SkillID = core.StringPtr(skillID)
	return options
}

// SetInput : Allow user to set Input
func (options *BulkClassifyOptions) SetInput(input []BulkClassifyUtterance) *BulkClassifyOptions {
	options.Input = input
	return options
}

// SetHeaders : Allow user to set Headers
func (options *BulkClassifyOptions) SetHeaders(param map[string]string) *BulkClassifyOptions {
	options.Headers = param
	return options
}

// BulkClassifyOutput : BulkClassifyOutput struct
type BulkClassifyOutput struct {
	// The user input utterance to classify.
	Input *BulkClassifyUtterance `json:"input,omitempty"`

	// An array of entities identified in the utterance.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// An array of intents recognized in the utterance.
	Intents []RuntimeIntent `json:"intents,omitempty"`
}

// UnmarshalBulkClassifyOutput unmarshals an instance of BulkClassifyOutput from the specified map of raw messages.
func UnmarshalBulkClassifyOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BulkClassifyOutput)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalBulkClassifyUtterance)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BulkClassifyResponse : BulkClassifyResponse struct
type BulkClassifyResponse struct {
	// An array of objects that contain classification information for the submitted input utterances.
	Output []BulkClassifyOutput `json:"output,omitempty"`
}

// UnmarshalBulkClassifyResponse unmarshals an instance of BulkClassifyResponse from the specified map of raw messages.
func UnmarshalBulkClassifyResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BulkClassifyResponse)
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalBulkClassifyOutput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BulkClassifyUtterance : The user input utterance to classify.
type BulkClassifyUtterance struct {
	// The text of the input utterance.
	Text *string `json:"text" validate:"required"`
}

// NewBulkClassifyUtterance : Instantiate BulkClassifyUtterance (Generic Model Constructor)
func (*AssistantV2) NewBulkClassifyUtterance(text string) (model *BulkClassifyUtterance, err error) {
	model = &BulkClassifyUtterance{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalBulkClassifyUtterance unmarshals an instance of BulkClassifyUtterance from the specified map of raw messages.
func UnmarshalBulkClassifyUtterance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BulkClassifyUtterance)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaptureGroup : CaptureGroup struct
type CaptureGroup struct {
	// A recognized capture group for the entity.
	Group *string `json:"group" validate:"required"`

	// Zero-based character offsets that indicate where the entity value begins and ends in the input text.
	Location []int64 `json:"location,omitempty"`
}

// NewCaptureGroup : Instantiate CaptureGroup (Generic Model Constructor)
func (*AssistantV2) NewCaptureGroup(group string) (model *CaptureGroup, err error) {
	model = &CaptureGroup{
		Group: core.StringPtr(group),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCaptureGroup unmarshals an instance of CaptureGroup from the specified map of raw messages.
func UnmarshalCaptureGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaptureGroup)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateSessionOptions : The CreateSession options.
type CreateSessionOptions struct {
	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSessionOptions : Instantiate CreateSessionOptions
func (*AssistantV2) NewCreateSessionOptions(assistantID string) *CreateSessionOptions {
	return &CreateSessionOptions{
		AssistantID: core.StringPtr(assistantID),
	}
}

// SetAssistantID : Allow user to set AssistantID
func (options *CreateSessionOptions) SetAssistantID(assistantID string) *CreateSessionOptions {
	options.AssistantID = core.StringPtr(assistantID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSessionOptions) SetHeaders(param map[string]string) *CreateSessionOptions {
	options.Headers = param
	return options
}

// DeleteSessionOptions : The DeleteSession options.
type DeleteSessionOptions struct {
	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required,ne="`

	// Unique identifier of the session.
	SessionID *string `json:"session_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSessionOptions : Instantiate DeleteSessionOptions
func (*AssistantV2) NewDeleteSessionOptions(assistantID string, sessionID string) *DeleteSessionOptions {
	return &DeleteSessionOptions{
		AssistantID: core.StringPtr(assistantID),
		SessionID:   core.StringPtr(sessionID),
	}
}

// SetAssistantID : Allow user to set AssistantID
func (options *DeleteSessionOptions) SetAssistantID(assistantID string) *DeleteSessionOptions {
	options.AssistantID = core.StringPtr(assistantID)
	return options
}

// SetSessionID : Allow user to set SessionID
func (options *DeleteSessionOptions) SetSessionID(sessionID string) *DeleteSessionOptions {
	options.SessionID = core.StringPtr(sessionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSessionOptions) SetHeaders(param map[string]string) *DeleteSessionOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {
	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (*AssistantV2) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(customerID string) *DeleteUserDataOptions {
	options.CustomerID = core.StringPtr(customerID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// DialogLogMessage : Dialog log message details.
type DialogLogMessage struct {
	// The severity of the log message.
	Level *string `json:"level" validate:"required"`

	// The text of the log message.
	Message *string `json:"message" validate:"required"`
}

// Constants associated with the DialogLogMessage.Level property.
// The severity of the log message.
const (
	DialogLogMessageLevelErrorConst = "error"
	DialogLogMessageLevelInfoConst  = "info"
	DialogLogMessageLevelWarnConst  = "warn"
)

// UnmarshalDialogLogMessage unmarshals an instance of DialogLogMessage from the specified map of raw messages.
func UnmarshalDialogLogMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogLogMessage)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeAction : DialogNodeAction struct
type DialogNodeAction struct {
	// The name of the action.
	Name *string `json:"name" validate:"required"`

	// The type of action to invoke.
	Type *string `json:"type,omitempty"`

	// A map of key/value pairs to be provided to the action.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The location in the dialog context where the result of the action is stored.
	ResultVariable *string `json:"result_variable" validate:"required"`

	// The name of the context variable that the client application will use to pass in credentials for the action.
	Credentials *string `json:"credentials,omitempty"`
}

// Constants associated with the DialogNodeAction.Type property.
// The type of action to invoke.
const (
	DialogNodeActionTypeClientConst        = "client"
	DialogNodeActionTypeCloudFunctionConst = "cloud-function"
	DialogNodeActionTypeServerConst        = "server"
	DialogNodeActionTypeWebActionConst     = "web-action"
)

// UnmarshalDialogNodeAction unmarshals an instance of DialogNodeAction from the specified map of raw messages.
func UnmarshalDialogNodeAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeAction)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result_variable", &obj.ResultVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credentials", &obj.Credentials)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputConnectToAgentTransferInfo : Routing or other contextual information to be used by target service desk systems.
type DialogNodeOutputConnectToAgentTransferInfo struct {
	Target map[string]map[string]interface{} `json:"target,omitempty"`
}

// UnmarshalDialogNodeOutputConnectToAgentTransferInfo unmarshals an instance of DialogNodeOutputConnectToAgentTransferInfo from the specified map of raw messages.
func UnmarshalDialogNodeOutputConnectToAgentTransferInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputConnectToAgentTransferInfo)
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputOptionsElement : DialogNodeOutputOptionsElement struct
type DialogNodeOutputOptionsElement struct {
	// The user-facing label for the option.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input to be sent to the assistant if the user selects the corresponding option.
	Value *DialogNodeOutputOptionsElementValue `json:"value" validate:"required"`
}

// UnmarshalDialogNodeOutputOptionsElement unmarshals an instance of DialogNodeOutputOptionsElement from the specified map of raw messages.
func UnmarshalDialogNodeOutputOptionsElement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputOptionsElement)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalDialogNodeOutputOptionsElementValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodeOutputOptionsElementValue : An object defining the message input to be sent to the assistant if the user selects the corresponding option.
type DialogNodeOutputOptionsElementValue struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`
}

// UnmarshalDialogNodeOutputOptionsElementValue unmarshals an instance of DialogNodeOutputOptionsElementValue from the specified map of raw messages.
func UnmarshalDialogNodeOutputOptionsElementValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodeOutputOptionsElementValue)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogNodesVisited : DialogNodesVisited struct
type DialogNodesVisited struct {
	// A dialog node that was triggered during processing of the input message.
	DialogNode *string `json:"dialog_node,omitempty"`

	// The title of the dialog node.
	Title *string `json:"title,omitempty"`

	// The conditions that trigger the dialog node.
	Conditions *string `json:"conditions,omitempty"`
}

// UnmarshalDialogNodesVisited unmarshals an instance of DialogNodesVisited from the specified map of raw messages.
func UnmarshalDialogNodesVisited(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogNodesVisited)
	err = core.UnmarshalPrimitive(m, "dialog_node", &obj.DialogNode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "conditions", &obj.Conditions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogSuggestion : DialogSuggestion struct
type DialogSuggestion struct {
	// The user-facing label for the suggestion. This label is taken from the **title** or **user_label** property of the
	// corresponding dialog node, depending on the disambiguation options.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input to be sent to the assistant if the user selects the corresponding
	// disambiguation option.
	Value *DialogSuggestionValue `json:"value" validate:"required"`

	// The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding
	// option.
	Output map[string]interface{} `json:"output,omitempty"`
}

// UnmarshalDialogSuggestion unmarshals an instance of DialogSuggestion from the specified map of raw messages.
func UnmarshalDialogSuggestion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogSuggestion)
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "value", &obj.Value, UnmarshalDialogSuggestionValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "output", &obj.Output)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DialogSuggestionValue : An object defining the message input to be sent to the assistant if the user selects the corresponding disambiguation
// option.
type DialogSuggestionValue struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`
}

// UnmarshalDialogSuggestionValue unmarshals an instance of DialogSuggestionValue from the specified map of raw messages.
func UnmarshalDialogSuggestionValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DialogSuggestionValue)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListLogsOptions : The ListLogs options.
type ListLogsOptions struct {
	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required,ne="`

	// How to sort the returned log events. You can sort by **request_timestamp**. To reverse the sort order, prefix the
	// parameter value with a minus sign (`-`).
	Sort *string `json:"sort,omitempty"`

	// A cacheable parameter that limits the results to those matching the specified filter. For more information, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-filter-reference#filter-reference).
	Filter *string `json:"filter,omitempty"`

	// The number of records to return in each page of results.
	PageLimit *int64 `json:"page_limit,omitempty"`

	// A token identifying the page of results to retrieve.
	Cursor *string `json:"cursor,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLogsOptions : Instantiate ListLogsOptions
func (*AssistantV2) NewListLogsOptions(assistantID string) *ListLogsOptions {
	return &ListLogsOptions{
		AssistantID: core.StringPtr(assistantID),
	}
}

// SetAssistantID : Allow user to set AssistantID
func (options *ListLogsOptions) SetAssistantID(assistantID string) *ListLogsOptions {
	options.AssistantID = core.StringPtr(assistantID)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListLogsOptions) SetSort(sort string) *ListLogsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetFilter : Allow user to set Filter
func (options *ListLogsOptions) SetFilter(filter string) *ListLogsOptions {
	options.Filter = core.StringPtr(filter)
	return options
}

// SetPageLimit : Allow user to set PageLimit
func (options *ListLogsOptions) SetPageLimit(pageLimit int64) *ListLogsOptions {
	options.PageLimit = core.Int64Ptr(pageLimit)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *ListLogsOptions) SetCursor(cursor string) *ListLogsOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLogsOptions) SetHeaders(param map[string]string) *ListLogsOptions {
	options.Headers = param
	return options
}

// Log : Log struct
type Log struct {
	// A unique identifier for the logged event.
	LogID *string `json:"log_id" validate:"required"`

	// A stateful message request formatted for the Watson Assistant service.
	Request *MessageRequest `json:"request" validate:"required"`

	// A response from the Watson Assistant service.
	Response *MessageResponse `json:"response" validate:"required"`

	// Unique identifier of the assistant.
	AssistantID *string `json:"assistant_id" validate:"required"`

	// The ID of the session the message was part of.
	SessionID *string `json:"session_id" validate:"required"`

	// The unique identifier of the skill that responded to the message.
	SkillID *string `json:"skill_id" validate:"required"`

	// The name of the snapshot (dialog skill version) that responded to the message (for example, `draft`).
	Snapshot *string `json:"snapshot" validate:"required"`

	// The timestamp for receipt of the message.
	RequestTimestamp *string `json:"request_timestamp" validate:"required"`

	// The timestamp for the system response to the message.
	ResponseTimestamp *string `json:"response_timestamp" validate:"required"`

	// The language of the assistant to which the message request was made.
	Language *string `json:"language" validate:"required"`

	// The customer ID specified for the message, if any.
	CustomerID *string `json:"customer_id,omitempty"`
}

// UnmarshalLog unmarshals an instance of Log from the specified map of raw messages.
func UnmarshalLog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Log)
	err = core.UnmarshalPrimitive(m, "log_id", &obj.LogID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "request", &obj.Request, UnmarshalMessageRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalMessageResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "assistant_id", &obj.AssistantID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "session_id", &obj.SessionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "skill_id", &obj.SkillID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "snapshot", &obj.Snapshot)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request_timestamp", &obj.RequestTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "response_timestamp", &obj.ResponseTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_id", &obj.CustomerID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogCollection : LogCollection struct
type LogCollection struct {
	// An array of objects describing log events.
	Logs []Log `json:"logs" validate:"required"`

	// The pagination data for the returned objects.
	Pagination *LogPagination `json:"pagination" validate:"required"`
}

// UnmarshalLogCollection unmarshals an instance of LogCollection from the specified map of raw messages.
func UnmarshalLogCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogCollection)
	err = core.UnmarshalModel(m, "logs", &obj.Logs, UnmarshalLog)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pagination", &obj.Pagination, UnmarshalLogPagination)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogPagination : The pagination data for the returned objects.
type LogPagination struct {
	// The URL that will return the next page of results, if any.
	NextURL *string `json:"next_url,omitempty"`

	// Reserved for future use.
	Matched *int64 `json:"matched,omitempty"`

	// A token identifying the next page of results.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// UnmarshalLogPagination unmarshals an instance of LogPagination from the specified map of raw messages.
func UnmarshalLogPagination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogPagination)
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "matched", &obj.Matched)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_cursor", &obj.NextCursor)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContext : MessageContext struct
type MessageContext struct {
	// Session context data that is shared by all skills used by the Assistant.
	Global *MessageContextGlobal `json:"global,omitempty"`

	// Information specific to particular skills used by the assistant.
	//
	// **Note:** Currently, only a single child property is supported, containing variables that apply to the dialog skill
	// used by the assistant.
	Skills map[string]MessageContextSkill `json:"skills,omitempty"`
}

// UnmarshalMessageContext unmarshals an instance of MessageContext from the specified map of raw messages.
func UnmarshalMessageContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContext)
	err = core.UnmarshalModel(m, "global", &obj.Global, UnmarshalMessageContextGlobal)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "skills", &obj.Skills, UnmarshalMessageContextSkill)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContextGlobal : Session context data that is shared by all skills used by the Assistant.
type MessageContextGlobal struct {
	// Built-in system properties that apply to all skills used by the assistant.
	System *MessageContextGlobalSystem `json:"system,omitempty"`

	// The session ID.
	SessionID *string `json:"session_id,omitempty"`
}

// UnmarshalMessageContextGlobal unmarshals an instance of MessageContextGlobal from the specified map of raw messages.
func UnmarshalMessageContextGlobal(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextGlobal)
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalMessageContextGlobalSystem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "session_id", &obj.SessionID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContextGlobalStateless : Session context data that is shared by all skills used by the Assistant.
type MessageContextGlobalStateless struct {
	// Built-in system properties that apply to all skills used by the assistant.
	System *MessageContextGlobalSystem `json:"system,omitempty"`

	// The unique identifier of the session.
	SessionID *string `json:"session_id,omitempty"`
}

// UnmarshalMessageContextGlobalStateless unmarshals an instance of MessageContextGlobalStateless from the specified map of raw messages.
func UnmarshalMessageContextGlobalStateless(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextGlobalStateless)
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalMessageContextGlobalSystem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "session_id", &obj.SessionID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContextGlobalSystem : Built-in system properties that apply to all skills used by the assistant.
type MessageContextGlobalSystem struct {
	// The user time zone. The assistant uses the time zone to correctly resolve relative time references.
	Timezone *string `json:"timezone,omitempty"`

	// A string value that identifies the user who is interacting with the assistant. The client must provide a unique
	// identifier for each individual end user who accesses the application. For Plus and Premium plans, this user ID is
	// used to identify unique users for billing purposes. This string cannot contain carriage return, newline, or tab
	// characters.
	UserID *string `json:"user_id,omitempty"`

	// A counter that is automatically incremented with each turn of the conversation. A value of 1 indicates that this is
	// the the first turn of a new conversation, which can affect the behavior of some skills (for example, triggering the
	// start node of a dialog).
	TurnCount *int64 `json:"turn_count,omitempty"`

	// The language code for localization in the user input. The specified locale overrides the default for the assistant,
	// and is used for interpreting entity values in user input such as date values. For example, `04/03/2018` might be
	// interpreted either as April 3 or March 4, depending on the locale.
	//
	//  This property is included only if the new system entities are enabled for the skill.
	Locale *string `json:"locale,omitempty"`

	// The base time for interpreting any relative time mentions in the user input. The specified time overrides the
	// current server time, and is used to calculate times mentioned in relative terms such as `now` or `tomorrow`. This
	// can be useful for simulating past or future times for testing purposes, or when analyzing documents such as news
	// articles.
	//
	// This value must be a UTC time value formatted according to ISO 8601 (for example, `2019-06-26T12:00:00Z` for noon on
	// 26 June 2019.
	//
	// This property is included only if the new system entities are enabled for the skill.
	ReferenceTime *string `json:"reference_time,omitempty"`
}

// Constants associated with the MessageContextGlobalSystem.Locale property.
// The language code for localization in the user input. The specified locale overrides the default for the assistant,
// and is used for interpreting entity values in user input such as date values. For example, `04/03/2018` might be
// interpreted either as April 3 or March 4, depending on the locale.
//
//  This property is included only if the new system entities are enabled for the skill.
const (
	MessageContextGlobalSystemLocaleArArConst = "ar-ar"
	MessageContextGlobalSystemLocaleCsCzConst = "cs-cz"
	MessageContextGlobalSystemLocaleDeDeConst = "de-de"
	MessageContextGlobalSystemLocaleEnCaConst = "en-ca"
	MessageContextGlobalSystemLocaleEnGbConst = "en-gb"
	MessageContextGlobalSystemLocaleEnUsConst = "en-us"
	MessageContextGlobalSystemLocaleEsEsConst = "es-es"
	MessageContextGlobalSystemLocaleFrFrConst = "fr-fr"
	MessageContextGlobalSystemLocaleItItConst = "it-it"
	MessageContextGlobalSystemLocaleJaJpConst = "ja-jp"
	MessageContextGlobalSystemLocaleKoKrConst = "ko-kr"
	MessageContextGlobalSystemLocaleNlNlConst = "nl-nl"
	MessageContextGlobalSystemLocalePtBrConst = "pt-br"
	MessageContextGlobalSystemLocaleZhCnConst = "zh-cn"
	MessageContextGlobalSystemLocaleZhTwConst = "zh-tw"
)

// UnmarshalMessageContextGlobalSystem unmarshals an instance of MessageContextGlobalSystem from the specified map of raw messages.
func UnmarshalMessageContextGlobalSystem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextGlobalSystem)
	err = core.UnmarshalPrimitive(m, "timezone", &obj.Timezone)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "turn_count", &obj.TurnCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locale", &obj.Locale)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reference_time", &obj.ReferenceTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContextSkill : Contains information specific to a particular skill used by the Assistant. The property name must be the same as the
// name of the skill (for example, `main skill`).
type MessageContextSkill struct {
	// Arbitrary variables that can be read and written by a particular skill.
	UserDefined map[string]interface{} `json:"user_defined,omitempty"`

	// System context data used by the skill.
	System *MessageContextSkillSystem `json:"system,omitempty"`
}

// UnmarshalMessageContextSkill unmarshals an instance of MessageContextSkill from the specified map of raw messages.
func UnmarshalMessageContextSkill(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextSkill)
	err = core.UnmarshalPrimitive(m, "user_defined", &obj.UserDefined)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalMessageContextSkillSystem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContextSkillSystem : System context data used by the skill.
type MessageContextSkillSystem struct {
	// An encoded string that represents the current conversation state. By saving this value and then sending it in the
	// context of a subsequent message request, you can return to an earlier point in the conversation. If you are using
	// stateful sessions, you can also use a stored state value to restore a paused conversation whose session is expired.
	State *string `json:"state,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of MessageContextSkillSystem
func (o *MessageContextSkillSystem) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of MessageContextSkillSystem
func (o *MessageContextSkillSystem) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of MessageContextSkillSystem
func (o *MessageContextSkillSystem) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of MessageContextSkillSystem
func (o *MessageContextSkillSystem) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.State != nil {
		m["state"] = o.State
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalMessageContextSkillSystem unmarshals an instance of MessageContextSkillSystem from the specified map of raw messages.
func UnmarshalMessageContextSkillSystem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextSkillSystem)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	delete(m, "state")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageContextStateless : MessageContextStateless struct
type MessageContextStateless struct {
	// Session context data that is shared by all skills used by the Assistant.
	Global *MessageContextGlobalStateless `json:"global,omitempty"`

	// Information specific to particular skills used by the assistant.
	//
	// **Note:** Currently, only a single child property is supported, containing variables that apply to the dialog skill
	// used by the assistant.
	Skills map[string]MessageContextSkill `json:"skills,omitempty"`
}

// UnmarshalMessageContextStateless unmarshals an instance of MessageContextStateless from the specified map of raw messages.
func UnmarshalMessageContextStateless(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageContextStateless)
	err = core.UnmarshalModel(m, "global", &obj.Global, UnmarshalMessageContextGlobalStateless)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "skills", &obj.Skills, UnmarshalMessageContextSkill)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageInput : An input object that includes the input text.
type MessageInput struct {
	// The type of user input. Currently, only text input is supported.
	MessageType *string `json:"message_type,omitempty"`

	// The text of the user input. This string cannot contain carriage return, newline, or tab characters.
	Text *string `json:"text,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those
	// intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those
	// entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// For internal use only.
	SuggestionID *string `json:"suggestion_id,omitempty"`

	// Optional properties that control how the assistant responds.
	Options *MessageInputOptions `json:"options,omitempty"`
}

// Constants associated with the MessageInput.MessageType property.
// The type of user input. Currently, only text input is supported.
const (
	MessageInputMessageTypeTextConst = "text"
)

// UnmarshalMessageInput unmarshals an instance of MessageInput from the specified map of raw messages.
func UnmarshalMessageInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageInput)
	err = core.UnmarshalPrimitive(m, "message_type", &obj.MessageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "suggestion_id", &obj.SuggestionID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalMessageInputOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageInputOptions : Optional properties that control how the assistant responds.
type MessageInputOptions struct {
	// Whether to restart dialog processing at the root of the dialog, regardless of any previously visited nodes.
	// **Note:** This does not affect `turn_count` or any other context variables.
	Restart *bool `json:"restart,omitempty"`

	// Whether to return more than one intent. Set to `true` to return all matching intents.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// Spelling correction options for the message. Any options specified on an individual message override the settings
	// configured for the skill.
	Spelling *MessageInputOptionsSpelling `json:"spelling,omitempty"`

	// Whether to return additional diagnostic information. Set to `true` to return additional information in the
	// `output.debug` property. If you also specify **return_context**=`true`, the returned skill context includes the
	// `system.state` property.
	Debug *bool `json:"debug,omitempty"`

	// Whether to return session context with the response. If you specify `true`, the response includes the `context`
	// property. If you also specify **debug**=`true`, the returned skill context includes the `system.state` property.
	ReturnContext *bool `json:"return_context,omitempty"`

	// Whether to return session context, including full conversation state. If you specify `true`, the response includes
	// the `context` property, and the skill context includes the `system.state` property.
	//
	// **Note:** If **export**=`true`, the context is returned regardless of the value of **return_context**.
	Export *bool `json:"export,omitempty"`
}

// UnmarshalMessageInputOptions unmarshals an instance of MessageInputOptions from the specified map of raw messages.
func UnmarshalMessageInputOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageInputOptions)
	err = core.UnmarshalPrimitive(m, "restart", &obj.Restart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternate_intents", &obj.AlternateIntents)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spelling", &obj.Spelling, UnmarshalMessageInputOptionsSpelling)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "debug", &obj.Debug)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "return_context", &obj.ReturnContext)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "export", &obj.Export)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageInputOptionsSpelling : Spelling correction options for the message. Any options specified on an individual message override the settings
// configured for the skill.
type MessageInputOptionsSpelling struct {
	// Whether to use spelling correction when processing the input. If spelling correction is used and **auto_correct** is
	// `true`, any spelling corrections are automatically applied to the user input. If **auto_correct** is `false`, any
	// suggested corrections are returned in the **output.spelling** property.
	//
	// This property overrides the value of the **spelling_suggestions** property in the workspace settings for the skill.
	Suggestions *bool `json:"suggestions,omitempty"`

	// Whether to use autocorrection when processing the input. If this property is `true`, any corrections are
	// automatically applied to the user input, and the original text is returned in the **output.spelling** property of
	// the message response. This property overrides the value of the **spelling_auto_correct** property in the workspace
	// settings for the skill.
	AutoCorrect *bool `json:"auto_correct,omitempty"`
}

// UnmarshalMessageInputOptionsSpelling unmarshals an instance of MessageInputOptionsSpelling from the specified map of raw messages.
func UnmarshalMessageInputOptionsSpelling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageInputOptionsSpelling)
	err = core.UnmarshalPrimitive(m, "suggestions", &obj.Suggestions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_correct", &obj.AutoCorrect)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageInputOptionsStateless : Optional properties that control how the assistant responds.
type MessageInputOptionsStateless struct {
	// Whether to restart dialog processing at the root of the dialog, regardless of any previously visited nodes.
	// **Note:** This does not affect `turn_count` or any other context variables.
	Restart *bool `json:"restart,omitempty"`

	// Whether to return more than one intent. Set to `true` to return all matching intents.
	AlternateIntents *bool `json:"alternate_intents,omitempty"`

	// Spelling correction options for the message. Any options specified on an individual message override the settings
	// configured for the skill.
	Spelling *MessageInputOptionsSpelling `json:"spelling,omitempty"`

	// Whether to return additional diagnostic information. Set to `true` to return additional information in the
	// `output.debug` property.
	Debug *bool `json:"debug,omitempty"`
}

// UnmarshalMessageInputOptionsStateless unmarshals an instance of MessageInputOptionsStateless from the specified map of raw messages.
func UnmarshalMessageInputOptionsStateless(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageInputOptionsStateless)
	err = core.UnmarshalPrimitive(m, "restart", &obj.Restart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternate_intents", &obj.AlternateIntents)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spelling", &obj.Spelling, UnmarshalMessageInputOptionsSpelling)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "debug", &obj.Debug)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageInputStateless : An input object that includes the input text.
type MessageInputStateless struct {
	// The type of user input. Currently, only text input is supported.
	MessageType *string `json:"message_type,omitempty"`

	// The text of the user input. This string cannot contain carriage return, newline, or tab characters.
	Text *string `json:"text,omitempty"`

	// Intents to use when evaluating the user input. Include intents from the previous response to continue using those
	// intents rather than trying to recognize intents in the new input.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// Entities to use when evaluating the message. Include entities from the previous response to continue using those
	// entities rather than detecting entities in the new input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// For internal use only.
	SuggestionID *string `json:"suggestion_id,omitempty"`

	// Optional properties that control how the assistant responds.
	Options *MessageInputOptionsStateless `json:"options,omitempty"`
}

// Constants associated with the MessageInputStateless.MessageType property.
// The type of user input. Currently, only text input is supported.
const (
	MessageInputStatelessMessageTypeTextConst = "text"
)

// UnmarshalMessageInputStateless unmarshals an instance of MessageInputStateless from the specified map of raw messages.
func UnmarshalMessageInputStateless(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageInputStateless)
	err = core.UnmarshalPrimitive(m, "message_type", &obj.MessageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "suggestion_id", &obj.SuggestionID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalMessageInputOptionsStateless)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageOptions : The Message options.
type MessageOptions struct {
	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required,ne="`

	// Unique identifier of the session.
	SessionID *string `json:"session_id" validate:"required,ne="`

	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// Context data for the conversation. You can use this property to set or modify context variables, which can also be
	// accessed by dialog nodes. The context is stored by the assistant on a per-session basis.
	//
	// **Note:** The total size of the context data stored for a stateful session cannot exceed 100KB.
	Context *MessageContext `json:"context,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMessageOptions : Instantiate MessageOptions
func (*AssistantV2) NewMessageOptions(assistantID string, sessionID string) *MessageOptions {
	return &MessageOptions{
		AssistantID: core.StringPtr(assistantID),
		SessionID:   core.StringPtr(sessionID),
	}
}

// SetAssistantID : Allow user to set AssistantID
func (options *MessageOptions) SetAssistantID(assistantID string) *MessageOptions {
	options.AssistantID = core.StringPtr(assistantID)
	return options
}

// SetSessionID : Allow user to set SessionID
func (options *MessageOptions) SetSessionID(sessionID string) *MessageOptions {
	options.SessionID = core.StringPtr(sessionID)
	return options
}

// SetInput : Allow user to set Input
func (options *MessageOptions) SetInput(input *MessageInput) *MessageOptions {
	options.Input = input
	return options
}

// SetContext : Allow user to set Context
func (options *MessageOptions) SetContext(context *MessageContext) *MessageOptions {
	options.Context = context
	return options
}

// SetHeaders : Allow user to set Headers
func (options *MessageOptions) SetHeaders(param map[string]string) *MessageOptions {
	options.Headers = param
	return options
}

// MessageOutput : Assistant output to be rendered or processed by the client.
type MessageOutput struct {
	// Output intended for any channel. It is the responsibility of the client application to implement the supported
	// response types.
	Generic []RuntimeResponseGenericIntf `json:"generic,omitempty"`

	// An array of intents recognized in the user input, sorted in descending order of confidence.
	Intents []RuntimeIntent `json:"intents,omitempty"`

	// An array of entities identified in the user input.
	Entities []RuntimeEntity `json:"entities,omitempty"`

	// An array of objects describing any actions requested by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// Additional detailed information about a message response and how it was generated.
	Debug *MessageOutputDebug `json:"debug,omitempty"`

	// An object containing any custom properties included in the response. This object includes any arbitrary properties
	// defined in the dialog JSON editor as part of the dialog node output.
	UserDefined map[string]interface{} `json:"user_defined,omitempty"`

	// Properties describing any spelling corrections in the user input that was received.
	Spelling *MessageOutputSpelling `json:"spelling,omitempty"`
}

// UnmarshalMessageOutput unmarshals an instance of MessageOutput from the specified map of raw messages.
func UnmarshalMessageOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageOutput)
	err = core.UnmarshalModel(m, "generic", &obj.Generic, UnmarshalRuntimeResponseGeneric)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intents", &obj.Intents, UnmarshalRuntimeIntent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entities", &obj.Entities, UnmarshalRuntimeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "actions", &obj.Actions, UnmarshalDialogNodeAction)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "debug", &obj.Debug, UnmarshalMessageOutputDebug)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_defined", &obj.UserDefined)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spelling", &obj.Spelling, UnmarshalMessageOutputSpelling)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageOutputDebug : Additional detailed information about a message response and how it was generated.
type MessageOutputDebug struct {
	// An array of objects containing detailed diagnostic information about the nodes that were triggered during processing
	// of the input message.
	NodesVisited []DialogNodesVisited `json:"nodes_visited,omitempty"`

	// An array of up to 50 messages logged with the request.
	LogMessages []DialogLogMessage `json:"log_messages,omitempty"`

	// Assistant sets this to true when this message response concludes or interrupts a dialog.
	BranchExited *bool `json:"branch_exited,omitempty"`

	// When `branch_exited` is set to `true` by the Assistant, the `branch_exited_reason` specifies whether the dialog
	// completed by itself or got interrupted.
	BranchExitedReason *string `json:"branch_exited_reason,omitempty"`
}

// Constants associated with the MessageOutputDebug.BranchExitedReason property.
// When `branch_exited` is set to `true` by the Assistant, the `branch_exited_reason` specifies whether the dialog
// completed by itself or got interrupted.
const (
	MessageOutputDebugBranchExitedReasonCompletedConst = "completed"
	MessageOutputDebugBranchExitedReasonFallbackConst  = "fallback"
)

// UnmarshalMessageOutputDebug unmarshals an instance of MessageOutputDebug from the specified map of raw messages.
func UnmarshalMessageOutputDebug(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageOutputDebug)
	err = core.UnmarshalModel(m, "nodes_visited", &obj.NodesVisited, UnmarshalDialogNodesVisited)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_messages", &obj.LogMessages, UnmarshalDialogLogMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "branch_exited", &obj.BranchExited)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "branch_exited_reason", &obj.BranchExitedReason)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageOutputSpelling : Properties describing any spelling corrections in the user input that was received.
type MessageOutputSpelling struct {
	// The user input text that was used to generate the response. If spelling autocorrection is enabled, this text
	// reflects any spelling corrections that were applied.
	Text *string `json:"text,omitempty"`

	// The original user input text. This property is returned only if autocorrection is enabled and the user input was
	// corrected.
	OriginalText *string `json:"original_text,omitempty"`

	// Any suggested corrections of the input text. This property is returned only if spelling correction is enabled and
	// autocorrection is disabled.
	SuggestedText *string `json:"suggested_text,omitempty"`
}

// UnmarshalMessageOutputSpelling unmarshals an instance of MessageOutputSpelling from the specified map of raw messages.
func UnmarshalMessageOutputSpelling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageOutputSpelling)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "original_text", &obj.OriginalText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "suggested_text", &obj.SuggestedText)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageRequest : A stateful message request formatted for the Watson Assistant service.
type MessageRequest struct {
	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// Context data for the conversation. You can use this property to set or modify context variables, which can also be
	// accessed by dialog nodes. The context is stored by the assistant on a per-session basis.
	//
	// **Note:** The total size of the context data stored for a stateful session cannot exceed 100KB.
	Context *MessageContext `json:"context,omitempty"`
}

// UnmarshalMessageRequest unmarshals an instance of MessageRequest from the specified map of raw messages.
func UnmarshalMessageRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageRequest)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalMessageInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalMessageContext)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageResponse : A response from the Watson Assistant service.
type MessageResponse struct {
	// Assistant output to be rendered or processed by the client.
	Output *MessageOutput `json:"output" validate:"required"`

	// Context data for the conversation. You can use this property to access context variables. The context is stored by
	// the assistant on a per-session basis.
	//
	// **Note:** The context is included in message responses only if **return_context**=`true` in the message request.
	// Full context is always included in logs.
	Context *MessageContext `json:"context,omitempty"`
}

// UnmarshalMessageResponse unmarshals an instance of MessageResponse from the specified map of raw messages.
func UnmarshalMessageResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageResponse)
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalMessageOutput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalMessageContext)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageResponseStateless : A stateless response from the Watson Assistant service.
type MessageResponseStateless struct {
	// Assistant output to be rendered or processed by the client.
	Output *MessageOutput `json:"output" validate:"required"`

	// Context data for the conversation. You can use this property to access context variables. The context is not stored
	// by the assistant; to maintain session state, include the context from the response in the next message.
	Context *MessageContextStateless `json:"context" validate:"required"`
}

// UnmarshalMessageResponseStateless unmarshals an instance of MessageResponseStateless from the specified map of raw messages.
func UnmarshalMessageResponseStateless(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MessageResponseStateless)
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalMessageOutput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalMessageContextStateless)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MessageStatelessOptions : The MessageStateless options.
type MessageStatelessOptions struct {
	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required,ne="`

	// An input object that includes the input text.
	Input *MessageInputStateless `json:"input,omitempty"`

	// Context data for the conversation. You can use this property to set or modify context variables, which can also be
	// accessed by dialog nodes. The context is not stored by the assistant. To maintain session state, include the context
	// from the previous response.
	//
	// **Note:** The total size of the context data for a stateless session cannot exceed 250KB.
	Context *MessageContextStateless `json:"context,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMessageStatelessOptions : Instantiate MessageStatelessOptions
func (*AssistantV2) NewMessageStatelessOptions(assistantID string) *MessageStatelessOptions {
	return &MessageStatelessOptions{
		AssistantID: core.StringPtr(assistantID),
	}
}

// SetAssistantID : Allow user to set AssistantID
func (options *MessageStatelessOptions) SetAssistantID(assistantID string) *MessageStatelessOptions {
	options.AssistantID = core.StringPtr(assistantID)
	return options
}

// SetInput : Allow user to set Input
func (options *MessageStatelessOptions) SetInput(input *MessageInputStateless) *MessageStatelessOptions {
	options.Input = input
	return options
}

// SetContext : Allow user to set Context
func (options *MessageStatelessOptions) SetContext(context *MessageContextStateless) *MessageStatelessOptions {
	options.Context = context
	return options
}

// SetHeaders : Allow user to set Headers
func (options *MessageStatelessOptions) SetHeaders(param map[string]string) *MessageStatelessOptions {
	options.Headers = param
	return options
}

// RuntimeEntity : The entity value that was recognized in the user input.
type RuntimeEntity struct {
	// An entity detected in the input.
	Entity *string `json:"entity" validate:"required"`

	// An array of zero-based character offsets that indicate where the detected entity values begin and end in the input
	// text.
	Location []int64 `json:"location" validate:"required"`

	// The term in the input text that was recognized as an entity value.
	Value *string `json:"value" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`

	// Any metadata for the entity.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The recognized capture groups for the entity, as defined by the entity pattern.
	Groups []CaptureGroup `json:"groups,omitempty"`

	// An object containing detailed information about the entity recognized in the user input. This property is included
	// only if the new system entities are enabled for the skill.
	//
	// For more information about how the new system entities are interpreted, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-beta-system-entities).
	Interpretation *RuntimeEntityInterpretation `json:"interpretation,omitempty"`

	// An array of possible alternative values that the user might have intended instead of the value returned in the
	// **value** property. This property is returned only for `@sys-time` and `@sys-date` entities when the user's input is
	// ambiguous.
	//
	// This property is included only if the new system entities are enabled for the skill.
	Alternatives []RuntimeEntityAlternative `json:"alternatives,omitempty"`

	// An object describing the role played by a system entity that is specifies the beginning or end of a range recognized
	// in the user input. This property is included only if the new system entities are enabled for the skill.
	Role *RuntimeEntityRole `json:"role,omitempty"`
}

// NewRuntimeEntity : Instantiate RuntimeEntity (Generic Model Constructor)
func (*AssistantV2) NewRuntimeEntity(entity string, location []int64, value string) (model *RuntimeEntity, err error) {
	model = &RuntimeEntity{
		Entity:   core.StringPtr(entity),
		Location: location,
		Value:    core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuntimeEntity unmarshals an instance of RuntimeEntity from the specified map of raw messages.
func UnmarshalRuntimeEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntity)
	err = core.UnmarshalPrimitive(m, "entity", &obj.Entity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalCaptureGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "interpretation", &obj.Interpretation, UnmarshalRuntimeEntityInterpretation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "alternatives", &obj.Alternatives, UnmarshalRuntimeEntityAlternative)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "role", &obj.Role, UnmarshalRuntimeEntityRole)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeEntityAlternative : An alternative value for the recognized entity.
type RuntimeEntityAlternative struct {
	// The entity value that was recognized in the user input.
	Value *string `json:"value,omitempty"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`
}

// UnmarshalRuntimeEntityAlternative unmarshals an instance of RuntimeEntityAlternative from the specified map of raw messages.
func UnmarshalRuntimeEntityAlternative(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntityAlternative)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeEntityInterpretation : RuntimeEntityInterpretation struct
type RuntimeEntityInterpretation struct {
	// The calendar used to represent a recognized date (for example, `Gregorian`).
	CalendarType *string `json:"calendar_type,omitempty"`

	// A unique identifier used to associate a recognized time and date. If the user input contains a date and time that
	// are mentioned together (for example, `Today at 5`, the same **datetime_link** value is returned for both the
	// `@sys-date` and `@sys-time` entities).
	DatetimeLink *string `json:"datetime_link,omitempty"`

	// A locale-specific holiday name (such as `thanksgiving` or `christmas`). This property is included when a `@sys-date`
	// entity is recognized based on a holiday name in the user input.
	Festival *string `json:"festival,omitempty"`

	// The precision or duration of a time range specified by a recognized `@sys-time` or `@sys-date` entity.
	Granularity *string `json:"granularity,omitempty"`

	// A unique identifier used to associate multiple recognized `@sys-date`, `@sys-time`, or `@sys-number` entities that
	// are recognized as a range of values in the user's input (for example, `from July 4 until July 14` or `from 20 to
	// 25`).
	RangeLink *string `json:"range_link,omitempty"`

	// The word in the user input that indicates that a `sys-date` or `sys-time` entity is part of an implied range where
	// only one date or time is specified (for example, `since` or `until`).
	RangeModifier *string `json:"range_modifier,omitempty"`

	// A recognized mention of a relative day, represented numerically as an offset from the current date (for example,
	// `-1` for `yesterday` or `10` for `in ten days`).
	RelativeDay *float64 `json:"relative_day,omitempty"`

	// A recognized mention of a relative month, represented numerically as an offset from the current month (for example,
	// `1` for `next month` or `-3` for `three months ago`).
	RelativeMonth *float64 `json:"relative_month,omitempty"`

	// A recognized mention of a relative week, represented numerically as an offset from the current week (for example,
	// `2` for `in two weeks` or `-1` for `last week).
	RelativeWeek *float64 `json:"relative_week,omitempty"`

	// A recognized mention of a relative date range for a weekend, represented numerically as an offset from the current
	// weekend (for example, `0` for `this weekend` or `-1` for `last weekend`).
	RelativeWeekend *float64 `json:"relative_weekend,omitempty"`

	// A recognized mention of a relative year, represented numerically as an offset from the current year (for example,
	// `1` for `next year` or `-5` for `five years ago`).
	RelativeYear *float64 `json:"relative_year,omitempty"`

	// A recognized mention of a specific date, represented numerically as the date within the month (for example, `30` for
	// `June 30`.).
	SpecificDay *float64 `json:"specific_day,omitempty"`

	// A recognized mention of a specific day of the week as a lowercase string (for example, `monday`).
	SpecificDayOfWeek *string `json:"specific_day_of_week,omitempty"`

	// A recognized mention of a specific month, represented numerically (for example, `7` for `July`).
	SpecificMonth *float64 `json:"specific_month,omitempty"`

	// A recognized mention of a specific quarter, represented numerically (for example, `3` for `the third quarter`).
	SpecificQuarter *float64 `json:"specific_quarter,omitempty"`

	// A recognized mention of a specific year (for example, `2016`).
	SpecificYear *float64 `json:"specific_year,omitempty"`

	// A recognized numeric value, represented as an integer or double.
	NumericValue *float64 `json:"numeric_value,omitempty"`

	// The type of numeric value recognized in the user input (`integer` or `rational`).
	Subtype *string `json:"subtype,omitempty"`

	// A recognized term for a time that was mentioned as a part of the day in the user's input (for example, `morning` or
	// `afternoon`).
	PartOfDay *string `json:"part_of_day,omitempty"`

	// A recognized mention of a relative hour, represented numerically as an offset from the current hour (for example,
	// `3` for `in three hours` or `-1` for `an hour ago`).
	RelativeHour *float64 `json:"relative_hour,omitempty"`

	// A recognized mention of a relative time, represented numerically as an offset in minutes from the current time (for
	// example, `5` for `in five minutes` or `-15` for `fifteen minutes ago`).
	RelativeMinute *float64 `json:"relative_minute,omitempty"`

	// A recognized mention of a relative time, represented numerically as an offset in seconds from the current time (for
	// example, `10` for `in ten seconds` or `-30` for `thirty seconds ago`).
	RelativeSecond *float64 `json:"relative_second,omitempty"`

	// A recognized specific hour mentioned as part of a time value (for example, `10` for `10:15 AM`.).
	SpecificHour *float64 `json:"specific_hour,omitempty"`

	// A recognized specific minute mentioned as part of a time value (for example, `15` for `10:15 AM`.).
	SpecificMinute *float64 `json:"specific_minute,omitempty"`

	// A recognized specific second mentioned as part of a time value (for example, `30` for `10:15:30 AM`.).
	SpecificSecond *float64 `json:"specific_second,omitempty"`

	// A recognized time zone mentioned as part of a time value (for example, `EST`).
	Timezone *string `json:"timezone,omitempty"`
}

// Constants associated with the RuntimeEntityInterpretation.Granularity property.
// The precision or duration of a time range specified by a recognized `@sys-time` or `@sys-date` entity.
const (
	RuntimeEntityInterpretationGranularityDayConst       = "day"
	RuntimeEntityInterpretationGranularityFortnightConst = "fortnight"
	RuntimeEntityInterpretationGranularityHourConst      = "hour"
	RuntimeEntityInterpretationGranularityInstantConst   = "instant"
	RuntimeEntityInterpretationGranularityMinuteConst    = "minute"
	RuntimeEntityInterpretationGranularityMonthConst     = "month"
	RuntimeEntityInterpretationGranularityQuarterConst   = "quarter"
	RuntimeEntityInterpretationGranularitySecondConst    = "second"
	RuntimeEntityInterpretationGranularityWeekConst      = "week"
	RuntimeEntityInterpretationGranularityWeekendConst   = "weekend"
	RuntimeEntityInterpretationGranularityYearConst      = "year"
)

// UnmarshalRuntimeEntityInterpretation unmarshals an instance of RuntimeEntityInterpretation from the specified map of raw messages.
func UnmarshalRuntimeEntityInterpretation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntityInterpretation)
	err = core.UnmarshalPrimitive(m, "calendar_type", &obj.CalendarType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "datetime_link", &obj.DatetimeLink)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "festival", &obj.Festival)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "granularity", &obj.Granularity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "range_link", &obj.RangeLink)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "range_modifier", &obj.RangeModifier)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_day", &obj.RelativeDay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_month", &obj.RelativeMonth)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_week", &obj.RelativeWeek)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_weekend", &obj.RelativeWeekend)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_year", &obj.RelativeYear)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_day", &obj.SpecificDay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_day_of_week", &obj.SpecificDayOfWeek)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_month", &obj.SpecificMonth)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_quarter", &obj.SpecificQuarter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_year", &obj.SpecificYear)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "numeric_value", &obj.NumericValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subtype", &obj.Subtype)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_day", &obj.PartOfDay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_hour", &obj.RelativeHour)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_minute", &obj.RelativeMinute)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relative_second", &obj.RelativeSecond)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_hour", &obj.SpecificHour)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_minute", &obj.SpecificMinute)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specific_second", &obj.SpecificSecond)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timezone", &obj.Timezone)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeEntityRole : An object describing the role played by a system entity that is specifies the beginning or end of a range recognized
// in the user input. This property is included only if the new system entities are enabled for the skill.
type RuntimeEntityRole struct {
	// The relationship of the entity to the range.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the RuntimeEntityRole.Type property.
// The relationship of the entity to the range.
const (
	RuntimeEntityRoleTypeDateFromConst   = "date_from"
	RuntimeEntityRoleTypeDateToConst     = "date_to"
	RuntimeEntityRoleTypeNumberFromConst = "number_from"
	RuntimeEntityRoleTypeNumberToConst   = "number_to"
	RuntimeEntityRoleTypeTimeFromConst   = "time_from"
	RuntimeEntityRoleTypeTimeToConst     = "time_to"
)

// UnmarshalRuntimeEntityRole unmarshals an instance of RuntimeEntityRole from the specified map of raw messages.
func UnmarshalRuntimeEntityRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeEntityRole)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeIntent : An intent identified in the user input.
type RuntimeIntent struct {
	// The name of the recognized intent.
	Intent *string `json:"intent" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the intent.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// NewRuntimeIntent : Instantiate RuntimeIntent (Generic Model Constructor)
func (*AssistantV2) NewRuntimeIntent(intent string, confidence float64) (model *RuntimeIntent, err error) {
	model = &RuntimeIntent{
		Intent:     core.StringPtr(intent),
		Confidence: core.Float64Ptr(confidence),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuntimeIntent unmarshals an instance of RuntimeIntent from the specified map of raw messages.
func UnmarshalRuntimeIntent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeIntent)
	err = core.UnmarshalPrimitive(m, "intent", &obj.Intent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGeneric : RuntimeResponseGeneric struct
// Models which "extend" this model:
// - RuntimeResponseGenericRuntimeResponseTypeText
// - RuntimeResponseGenericRuntimeResponseTypePause
// - RuntimeResponseGenericRuntimeResponseTypeImage
// - RuntimeResponseGenericRuntimeResponseTypeOption
// - RuntimeResponseGenericRuntimeResponseTypeConnectToAgent
// - RuntimeResponseGenericRuntimeResponseTypeSuggestion
// - RuntimeResponseGenericRuntimeResponseTypeSearch
type RuntimeResponseGeneric struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type,omitempty"`

	// The text of the response.
	Text *string `json:"text,omitempty"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`

	// The URL of the image.
	Source *string `json:"source,omitempty"`

	// The title to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// An optional message to be displayed to the user to indicate that the conversation will be transferred to the next
	// available agent.
	AgentAvailable *AgentAvailabilityMessage `json:"agent_available,omitempty"`

	// An optional message to be displayed to the user to indicate that no online agent is available to take over the
	// conversation.
	AgentUnavailable *AgentAvailabilityMessage `json:"agent_unavailable,omitempty"`

	// Routing or other contextual information to be used by target service desk systems.
	TransferInfo *DialogNodeOutputConnectToAgentTransferInfo `json:"transfer_info,omitempty"`

	// A label identifying the topic of the conversation, derived from the **title** property of the relevant node or the
	// **topic** property of the dialog node response.
	Topic *string `json:"topic,omitempty"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	Suggestions []DialogSuggestion `json:"suggestions,omitempty"`

	// The title or introductory text to show before the response. This text is defined in the search skill configuration.
	Header *string `json:"header,omitempty"`

	// An array of objects that contains the search results to be displayed in the initial response to the user.
	PrimaryResults []SearchResult `json:"primary_results,omitempty"`

	// An array of objects that contains additional search results that can be displayed to the user upon request.
	AdditionalResults []SearchResult `json:"additional_results,omitempty"`
}

// Constants associated with the RuntimeResponseGeneric.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericResponseTypeTextConst = "text"
)

// Constants associated with the RuntimeResponseGeneric.Preference property.
// The preferred type of control to display.
const (
	RuntimeResponseGenericPreferenceButtonConst   = "button"
	RuntimeResponseGenericPreferenceDropdownConst = "dropdown"
)

func (*RuntimeResponseGeneric) isaRuntimeResponseGeneric() bool {
	return true
}

type RuntimeResponseGenericIntf interface {
	isaRuntimeResponseGeneric() bool
}

// UnmarshalRuntimeResponseGeneric unmarshals an instance of RuntimeResponseGeneric from the specified map of raw messages.
func UnmarshalRuntimeResponseGeneric(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "response_type", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'response_type': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'response_type' not found in JSON object")
		return
	}
	if discValue == "connect_to_agent" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeConnectToAgent)
	} else if discValue == "image" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeImage)
	} else if discValue == "option" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeOption)
	} else if discValue == "suggestion" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeSuggestion)
	} else if discValue == "pause" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypePause)
	} else if discValue == "search" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeSearch)
	} else if discValue == "text" {
		err = core.UnmarshalModel(m, "", result, UnmarshalRuntimeResponseGenericRuntimeResponseTypeText)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'response_type': %s", discValue)
	}
	return
}

// SearchResult : SearchResult struct
type SearchResult struct {
	// The unique identifier of the document in the Discovery service collection.
	//
	// This property is included in responses from search skills, which are available only to Plus or Premium plan users.
	ID *string `json:"id" validate:"required"`

	// An object containing search result metadata from the Discovery service.
	ResultMetadata *SearchResultMetadata `json:"result_metadata" validate:"required"`

	// A description of the search result. This is taken from an abstract, summary, or highlight field in the Discovery
	// service response, as specified in the search skill configuration.
	Body *string `json:"body,omitempty"`

	// The title of the search result. This is taken from a title or name field in the Discovery service response, as
	// specified in the search skill configuration.
	Title *string `json:"title,omitempty"`

	// The URL of the original data object in its native data source.
	URL *string `json:"url,omitempty"`

	// An object containing segments of text from search results with query-matching text highlighted using HTML `<em>`
	// tags.
	Highlight *SearchResultHighlight `json:"highlight,omitempty"`
}

// UnmarshalSearchResult unmarshals an instance of SearchResult from the specified map of raw messages.
func UnmarshalSearchResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SearchResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_metadata", &obj.ResultMetadata, UnmarshalSearchResultMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "body", &obj.Body)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "highlight", &obj.Highlight, UnmarshalSearchResultHighlight)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchResultHighlight : An object containing segments of text from search results with query-matching text highlighted using HTML `<em>`
// tags.
type SearchResultHighlight struct {
	// An array of strings containing segments taken from body text in the search results, with query-matching substrings
	// highlighted.
	Body []string `json:"body,omitempty"`

	// An array of strings containing segments taken from title text in the search results, with query-matching substrings
	// highlighted.
	Title []string `json:"title,omitempty"`

	// An array of strings containing segments taken from URLs in the search results, with query-matching substrings
	// highlighted.
	URL []string `json:"url,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string][]string
}

// SetProperty allows the user to set an arbitrary property on an instance of SearchResultHighlight
func (o *SearchResultHighlight) SetProperty(key string, value []string) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string][]string)
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of SearchResultHighlight
func (o *SearchResultHighlight) GetProperty(key string) []string {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of SearchResultHighlight
func (o *SearchResultHighlight) GetProperties() map[string][]string {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of SearchResultHighlight
func (o *SearchResultHighlight) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Body != nil {
		m["body"] = o.Body
	}
	if o.Title != nil {
		m["title"] = o.Title
	}
	if o.URL != nil {
		m["url"] = o.URL
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalSearchResultHighlight unmarshals an instance of SearchResultHighlight from the specified map of raw messages.
func UnmarshalSearchResultHighlight(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SearchResultHighlight)
	err = core.UnmarshalPrimitive(m, "body", &obj.Body)
	if err != nil {
		return
	}
	delete(m, "body")
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	delete(m, "title")
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	delete(m, "url")
	for k := range m {
		var v []string
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchResultMetadata : An object containing search result metadata from the Discovery service.
type SearchResultMetadata struct {
	// The confidence score for the given result. For more information about how the confidence is calculated, see the
	// Discovery service [documentation](../discovery#query-your-collection).
	Confidence *float64 `json:"confidence,omitempty"`

	// An unbounded measure of the relevance of a particular result, dependent on the query and matching document. A higher
	// score indicates a greater match to the query parameters.
	Score *float64 `json:"score,omitempty"`
}

// UnmarshalSearchResultMetadata unmarshals an instance of SearchResultMetadata from the specified map of raw messages.
func UnmarshalSearchResultMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SearchResultMetadata)
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SessionResponse : SessionResponse struct
type SessionResponse struct {
	// The session ID.
	SessionID *string `json:"session_id" validate:"required"`
}

// UnmarshalSessionResponse unmarshals an instance of SessionResponse from the specified map of raw messages.
func UnmarshalSessionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SessionResponse)
	err = core.UnmarshalPrimitive(m, "session_id", &obj.SessionID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeConnectToAgent : An object that describes a response with response type `connect_to_agent`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeConnectToAgent struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// An optional message to be displayed to the user to indicate that the conversation will be transferred to the next
	// available agent.
	AgentAvailable *AgentAvailabilityMessage `json:"agent_available,omitempty"`

	// An optional message to be displayed to the user to indicate that no online agent is available to take over the
	// conversation.
	AgentUnavailable *AgentAvailabilityMessage `json:"agent_unavailable,omitempty"`

	// Routing or other contextual information to be used by target service desk systems.
	TransferInfo *DialogNodeOutputConnectToAgentTransferInfo `json:"transfer_info,omitempty"`

	// A label identifying the topic of the conversation, derived from the **title** property of the relevant node or the
	// **topic** property of the dialog node response.
	Topic *string `json:"topic,omitempty"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeConnectToAgent.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypeConnectToAgentResponseTypeConnectToAgentConst = "connect_to_agent"
)

func (*RuntimeResponseGenericRuntimeResponseTypeConnectToAgent) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeConnectToAgent unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeConnectToAgent from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeConnectToAgent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeConnectToAgent)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_to_human_agent", &obj.MessageToHumanAgent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "agent_available", &obj.AgentAvailable, UnmarshalAgentAvailabilityMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "agent_unavailable", &obj.AgentUnavailable, UnmarshalAgentAvailabilityMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transfer_info", &obj.TransferInfo, UnmarshalDialogNodeOutputConnectToAgentTransferInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "topic", &obj.Topic)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeImage : An object that describes a response with response type `image`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeImage struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The URL of the image.
	Source *string `json:"source" validate:"required"`

	// The title to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeImage.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypeImageResponseTypeImageConst = "image"
)

func (*RuntimeResponseGenericRuntimeResponseTypeImage) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeImage unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeImage from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeImage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeImage)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeOption : An object that describes a response with response type `option`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeOption struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The title or introductory text to show before the response.
	Title *string `json:"title" validate:"required"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options" validate:"required"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeOption.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypeOptionResponseTypeOptionConst = "option"
)

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeOption.Preference property.
// The preferred type of control to display.
const (
	RuntimeResponseGenericRuntimeResponseTypeOptionPreferenceButtonConst   = "button"
	RuntimeResponseGenericRuntimeResponseTypeOptionPreferenceDropdownConst = "dropdown"
)

func (*RuntimeResponseGenericRuntimeResponseTypeOption) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeOption unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeOption from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeOption)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "options", &obj.Options, UnmarshalDialogNodeOutputOptionsElement)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypePause : An object that describes a response with response type `pause`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypePause struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time" validate:"required"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypePause.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypePauseResponseTypePauseConst = "pause"
)

func (*RuntimeResponseGenericRuntimeResponseTypePause) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypePause unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypePause from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypePause(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypePause)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time", &obj.Time)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "typing", &obj.Typing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeSearch : An object that describes a response with response type `search`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeSearch struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The title or introductory text to show before the response. This text is defined in the search skill configuration.
	Header *string `json:"header" validate:"required"`

	// An array of objects that contains the search results to be displayed in the initial response to the user.
	PrimaryResults []SearchResult `json:"primary_results" validate:"required"`

	// An array of objects that contains additional search results that can be displayed to the user upon request.
	AdditionalResults []SearchResult `json:"additional_results" validate:"required"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeSearch.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypeSearchResponseTypeSearchConst = "search"
)

func (*RuntimeResponseGenericRuntimeResponseTypeSearch) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeSearch unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeSearch from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeSearch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeSearch)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "header", &obj.Header)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "primary_results", &obj.PrimaryResults, UnmarshalSearchResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "additional_results", &obj.AdditionalResults, UnmarshalSearchResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeSuggestion : An object that describes a response with response type `suggestion`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeSuggestion struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The title or introductory text to show before the response.
	Title *string `json:"title" validate:"required"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	Suggestions []DialogSuggestion `json:"suggestions" validate:"required"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeSuggestion.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypeSuggestionResponseTypeSuggestionConst = "suggestion"
)

func (*RuntimeResponseGenericRuntimeResponseTypeSuggestion) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeSuggestion unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeSuggestion from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeSuggestion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeSuggestion)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "suggestions", &obj.Suggestions, UnmarshalDialogSuggestion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuntimeResponseGenericRuntimeResponseTypeText : An object that describes a response with response type `text`.
// This model "extends" RuntimeResponseGeneric
type RuntimeResponseGenericRuntimeResponseTypeText struct {
	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	ResponseType *string `json:"response_type" validate:"required"`

	// The text of the response.
	Text *string `json:"text" validate:"required"`
}

// Constants associated with the RuntimeResponseGenericRuntimeResponseTypeText.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
const (
	RuntimeResponseGenericRuntimeResponseTypeTextResponseTypeTextConst = "text"
)

func (*RuntimeResponseGenericRuntimeResponseTypeText) isaRuntimeResponseGeneric() bool {
	return true
}

// UnmarshalRuntimeResponseGenericRuntimeResponseTypeText unmarshals an instance of RuntimeResponseGenericRuntimeResponseTypeText from the specified map of raw messages.
func UnmarshalRuntimeResponseGenericRuntimeResponseTypeText(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuntimeResponseGenericRuntimeResponseTypeText)
	err = core.UnmarshalPrimitive(m, "response_type", &obj.ResponseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
