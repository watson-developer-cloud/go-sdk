/**
 * (C) Copyright IBM Corp. 2020.
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

// Package assistantv2 : Operations and models for the AssistantV2 service
package assistantv2

import (
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// AssistantV2 : The IBM Watson&trade; Assistant service combines machine learning, natural language understanding, and
// an integrated dialog editor to create conversation flows between your apps and your users.
//
// The Assistant v2 API provides runtime methods your client application can use to send user input to an assistant and
// receive a response.
//
// Version: 2.0
// See: https://cloud.ibm.com/docs/assistant/
type AssistantV2 struct {
	Service *core.BaseService
	Version string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://gateway.watsonplatform.net/assistant/api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "conversation"

// AssistantV2Options : Service options
type AssistantV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
	Version       string
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

// SetServiceURL sets the service URL
func (assistant *AssistantV2) SetServiceURL(url string) error {
	return assistant.Service.SetServiceURL(url)
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
	err = core.ValidateNotNil(createSessionOptions, "createSessionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSessionOptions, "createSessionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/assistants", "sessions"}
	pathParameters := []string{*createSessionOptions.AssistantID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = assistant.Service.Request(request, new(SessionResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*SessionResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// DeleteSession : Delete session
// Deletes a session explicitly before it times out. (For more information about the session inactivity timeout, see the
// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-settings)).
func (assistant *AssistantV2) DeleteSession(deleteSessionOptions *DeleteSessionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSessionOptions, "deleteSessionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSessionOptions, "deleteSessionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/assistants", "sessions"}
	pathParameters := []string{*deleteSessionOptions.AssistantID, *deleteSessionOptions.SessionID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

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
//
// There is no rate limit for this operation.
func (assistant *AssistantV2) Message(messageOptions *MessageOptions) (result *MessageResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(messageOptions, "messageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(messageOptions, "messageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/assistants", "sessions", "message"}
	pathParameters := []string{*messageOptions.AssistantID, *messageOptions.SessionID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(MessageResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*MessageResponse)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

	return
}

// MessageStateless : Send user input to assistant (stateless)
// Send user input to an assistant and receive a response, with conversation state (including context data) managed by
// your application.
//
// There is no rate limit for this operation.
func (assistant *AssistantV2) MessageStateless(messageStatelessOptions *MessageStatelessOptions) (result *MessageResponseStateless, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(messageStatelessOptions, "messageStatelessOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(messageStatelessOptions, "messageStatelessOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/assistants", "message"}
	pathParameters := []string{*messageStatelessOptions.AssistantID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(assistant.Service.Options.URL, pathSegments, pathParameters)
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
	builder.AddQuery("version", assistant.Version)

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

	response, err = assistant.Service.Request(request, new(MessageResponseStateless))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*MessageResponseStateless)
		if !ok {
			err = fmt.Errorf("An error occurred while processing the operation response.")
		}
	}

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
func (assistant *AssistantV2) NewCaptureGroup(group string) (model *CaptureGroup, err error) {
	model = &CaptureGroup{
		Group: core.StringPtr(group),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// CreateSessionOptions : The CreateSession options.
type CreateSessionOptions struct {

	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateSessionOptions : Instantiate CreateSessionOptions
func (assistant *AssistantV2) NewCreateSessionOptions(assistantID string) *CreateSessionOptions {
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
	AssistantID *string `json:"assistant_id" validate:"required"`

	// Unique identifier of the session.
	SessionID *string `json:"session_id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteSessionOptions : Instantiate DeleteSessionOptions
func (assistant *AssistantV2) NewDeleteSessionOptions(assistantID string, sessionID string) *DeleteSessionOptions {
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
	DialogLogMessage_Level_Error = "error"
	DialogLogMessage_Level_Info  = "info"
	DialogLogMessage_Level_Warn  = "warn"
)

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
	DialogNodeAction_Type_Client        = "client"
	DialogNodeAction_Type_CloudFunction = "cloud-function"
	DialogNodeAction_Type_Server        = "server"
	DialogNodeAction_Type_WebAction     = "web-action"
)

// DialogNodeOutputOptionsElement : DialogNodeOutputOptionsElement struct
type DialogNodeOutputOptionsElement struct {

	// The user-facing label for the option.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input to be sent to the assistant if the user selects the corresponding option.
	Value *DialogNodeOutputOptionsElementValue `json:"value" validate:"required"`
}

// DialogNodeOutputOptionsElementValue : An object defining the message input to be sent to the assistant if the user selects the corresponding option.
type DialogNodeOutputOptionsElementValue struct {

	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`
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

// DialogSuggestion : DialogSuggestion struct
type DialogSuggestion struct {

	// The user-facing label for the disambiguation option. This label is taken from the **title** or **user_label**
	// property of the corresponding dialog node, depending on the disambiguation options.
	Label *string `json:"label" validate:"required"`

	// An object defining the message input to be sent to the assistant if the user selects the corresponding
	// disambiguation option.
	Value *DialogSuggestionValue `json:"value" validate:"required"`

	// The dialog output that will be returned from the Watson Assistant service if the user selects the corresponding
	// option.
	Output map[string]interface{} `json:"output,omitempty"`
}

// DialogSuggestionValue : An object defining the message input to be sent to the assistant if the user selects the corresponding disambiguation
// option.
type DialogSuggestionValue struct {

	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`
}

// MessageContext : MessageContext struct
type MessageContext struct {

	// Session context data that is shared by all skills used by the Assistant.
	Global *MessageContextGlobal `json:"global,omitempty"`

	// Information specific to particular skills used by the assistant.
	//
	// **Note:** Currently, only a single child property is supported, containing variables that apply to the dialog skill
	// used by the assistant.
	Skills *MessageContextSkills `json:"skills,omitempty"`
}

// MessageContextGlobal : Session context data that is shared by all skills used by the Assistant.
type MessageContextGlobal struct {

	// Built-in system properties that apply to all skills used by the assistant.
	System *MessageContextGlobalSystem `json:"system,omitempty"`

	// The session ID.
	SessionID *string `json:"session_id,omitempty"`
}

// MessageContextGlobalStateless : Session context data that is shared by all skills used by the Assistant.
type MessageContextGlobalStateless struct {

	// Built-in system properties that apply to all skills used by the assistant.
	System *MessageContextGlobalSystem `json:"system,omitempty"`

	// The unique identifier of the session.
	SessionID *string `json:"session_id,omitempty"`
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
	MessageContextGlobalSystem_Locale_ArAr = "ar-ar"
	MessageContextGlobalSystem_Locale_CsCz = "cs-cz"
	MessageContextGlobalSystem_Locale_DeDe = "de-de"
	MessageContextGlobalSystem_Locale_EnCa = "en-ca"
	MessageContextGlobalSystem_Locale_EnGb = "en-gb"
	MessageContextGlobalSystem_Locale_EnUs = "en-us"
	MessageContextGlobalSystem_Locale_EsEs = "es-es"
	MessageContextGlobalSystem_Locale_FrFr = "fr-fr"
	MessageContextGlobalSystem_Locale_ItIt = "it-it"
	MessageContextGlobalSystem_Locale_JaJp = "ja-jp"
	MessageContextGlobalSystem_Locale_KoKr = "ko-kr"
	MessageContextGlobalSystem_Locale_NlNl = "nl-nl"
	MessageContextGlobalSystem_Locale_PtBr = "pt-br"
	MessageContextGlobalSystem_Locale_ZhCn = "zh-cn"
	MessageContextGlobalSystem_Locale_ZhTw = "zh-tw"
)

// MessageContextSkill : Contains information specific to a particular skill used by the Assistant. The property name must be the same as the
// name of the skill (for example, `main skill`).
type MessageContextSkill struct {

	// Arbitrary variables that can be read and written by a particular skill.
	UserDefined map[string]interface{} `json:"user_defined,omitempty"`

	// System context data used by the skill.
	System *MessageContextSkillSystem `json:"system,omitempty"`
}

// MessageContextSkillSystem : System context data used by the skill.
type MessageContextSkillSystem map[string]interface{}

// SetState : Allow user to set State
func (this *MessageContextSkillSystem) SetState(State *string) {
	(*this)["state"] = State
}

// GetState : Allow user to get State
func (this *MessageContextSkillSystem) GetState() *string {
	return (*this)["state"].(*string)
}

// SetProperty : Allow user to set arbitrary property
func (this *MessageContextSkillSystem) SetProperty(Key string, Value *interface{}) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *MessageContextSkillSystem) GetProperty(Key string) *interface{} {
	return (*this)[Key].(*interface{})
}

// MessageContextSkills : Information specific to particular skills used by the assistant.
//
// **Note:** Currently, only a single child property is supported, containing variables that apply to the dialog skill
// used by the assistant.
type MessageContextSkills map[string]interface{}

// SetProperty : Allow user to set arbitrary property
func (this *MessageContextSkills) SetProperty(Key string, Value *MessageContextSkill) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *MessageContextSkills) GetProperty(Key string) *MessageContextSkill {
	return (*this)[Key].(*MessageContextSkill)
}

// MessageContextStateless : MessageContextStateless struct
type MessageContextStateless struct {

	// Session context data that is shared by all skills used by the Assistant.
	Global *MessageContextGlobalStateless `json:"global,omitempty"`

	// Information specific to particular skills used by the assistant.
	//
	// **Note:** Currently, only a single child property is supported, containing variables that apply to the dialog skill
	// used by the assistant.
	Skills *MessageContextSkills `json:"skills,omitempty"`
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
	MessageInput_MessageType_Text = "text"
)

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
	MessageInputStateless_MessageType_Text = "text"
)

// MessageOptions : The Message options.
type MessageOptions struct {

	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required"`

	// Unique identifier of the session.
	SessionID *string `json:"session_id" validate:"required"`

	// An input object that includes the input text.
	Input *MessageInput `json:"input,omitempty"`

	// Context data for the conversation. You can use this property to set or modify context variables, which can also be
	// accessed by dialog nodes. The context is stored by the assistant on a per-session basis.
	//
	// **Note:** The total size of the context data stored for a stateful session cannot exceed 100KB.
	Context *MessageContext `json:"context,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewMessageOptions : Instantiate MessageOptions
func (assistant *AssistantV2) NewMessageOptions(assistantID string, sessionID string) *MessageOptions {
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
	Generic []RuntimeResponseGeneric `json:"generic,omitempty"`

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
	MessageOutputDebug_BranchExitedReason_Completed = "completed"
	MessageOutputDebug_BranchExitedReason_Fallback  = "fallback"
)

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

// MessageResponse : A response from the Watson Assistant service.
type MessageResponse struct {

	// Assistant output to be rendered or processed by the client.
	Output *MessageOutput `json:"output" validate:"required"`

	// Context data for the conversation. You can use this property to access context variables. The context is stored by
	// the assistant on a per-session basis.
	//
	// **Note:** The context is included in message responses only if **return_context**=`true` in the message request.
	Context *MessageContext `json:"context,omitempty"`
}

// MessageResponseStateless : A stateless response from the Watson Assistant service.
type MessageResponseStateless struct {

	// Assistant output to be rendered or processed by the client.
	Output *MessageOutput `json:"output" validate:"required"`

	// Context data for the conversation. You can use this property to access context variables. The context is not stored
	// by the assistant; to maintain session state, include the context from the response in the next message.
	Context *MessageContextStateless `json:"context" validate:"required"`
}

// MessageStatelessOptions : The MessageStateless options.
type MessageStatelessOptions struct {

	// Unique identifier of the assistant. To find the assistant ID in the Watson Assistant user interface, open the
	// assistant settings and click **API Details**. For information about creating assistants, see the
	// [documentation](https://cloud.ibm.com/docs/assistant?topic=assistant-assistant-add#assistant-add-task).
	//
	// **Note:** Currently, the v2 API does not support creating assistants.
	AssistantID *string `json:"assistant_id" validate:"required"`

	// An input object that includes the input text.
	Input *MessageInputStateless `json:"input,omitempty"`

	// Context data for the conversation. You can use this property to set or modify context variables, which can also be
	// accessed by dialog nodes. The context is not stored by the assistant. To maintain session state, include the context
	// from the previous response.
	//
	// **Note:** The total size of the context data for a stateless session cannot exceed 250KB.
	Context *MessageContextStateless `json:"context,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewMessageStatelessOptions : Instantiate MessageStatelessOptions
func (assistant *AssistantV2) NewMessageStatelessOptions(assistantID string) *MessageStatelessOptions {
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
func (assistant *AssistantV2) NewRuntimeEntity(entity string, location []int64, value string) (model *RuntimeEntity, err error) {
	model = &RuntimeEntity{
		Entity:   core.StringPtr(entity),
		Location: location,
		Value:    core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// RuntimeEntityAlternative : An alternative value for the recognized entity.
type RuntimeEntityAlternative struct {

	// The entity value that was recognized in the user input.
	Value *string `json:"value,omitempty"`

	// A decimal percentage that represents Watson's confidence in the recognized entity.
	Confidence *float64 `json:"confidence,omitempty"`
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
	RuntimeEntityInterpretation_Granularity_Day       = "day"
	RuntimeEntityInterpretation_Granularity_Fortnight = "fortnight"
	RuntimeEntityInterpretation_Granularity_Hour      = "hour"
	RuntimeEntityInterpretation_Granularity_Instant   = "instant"
	RuntimeEntityInterpretation_Granularity_Minute    = "minute"
	RuntimeEntityInterpretation_Granularity_Month     = "month"
	RuntimeEntityInterpretation_Granularity_Quarter   = "quarter"
	RuntimeEntityInterpretation_Granularity_Second    = "second"
	RuntimeEntityInterpretation_Granularity_Week      = "week"
	RuntimeEntityInterpretation_Granularity_Weekend   = "weekend"
	RuntimeEntityInterpretation_Granularity_Year      = "year"
)

// RuntimeEntityRole : An object describing the role played by a system entity that is specifies the beginning or end of a range recognized
// in the user input. This property is included only if the new system entities are enabled for the skill.
type RuntimeEntityRole struct {

	// The relationship of the entity to the range.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the RuntimeEntityRole.Type property.
// The relationship of the entity to the range.
const (
	RuntimeEntityRole_Type_DateFrom   = "date_from"
	RuntimeEntityRole_Type_DateTo     = "date_to"
	RuntimeEntityRole_Type_NumberFrom = "number_from"
	RuntimeEntityRole_Type_NumberTo   = "number_to"
	RuntimeEntityRole_Type_TimeFrom   = "time_from"
	RuntimeEntityRole_Type_TimeTo     = "time_to"
)

// RuntimeIntent : An intent identified in the user input.
type RuntimeIntent struct {

	// The name of the recognized intent.
	Intent *string `json:"intent" validate:"required"`

	// A decimal percentage that represents Watson's confidence in the intent.
	Confidence *float64 `json:"confidence" validate:"required"`
}

// NewRuntimeIntent : Instantiate RuntimeIntent (Generic Model Constructor)
func (assistant *AssistantV2) NewRuntimeIntent(intent string, confidence float64) (model *RuntimeIntent, err error) {
	model = &RuntimeIntent{
		Intent:     core.StringPtr(intent),
		Confidence: core.Float64Ptr(confidence),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// RuntimeResponseGeneric : RuntimeResponseGeneric struct
type RuntimeResponseGeneric struct {

	// The type of response returned by the dialog node. The specified response type must be supported by the client
	// application or channel.
	//
	// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for
	// Premium users.
	ResponseType *string `json:"response_type" validate:"required"`

	// The text of the response.
	Text *string `json:"text,omitempty"`

	// How long to pause, in milliseconds.
	Time *int64 `json:"time,omitempty"`

	// Whether to send a "user is typing" event during the pause.
	Typing *bool `json:"typing,omitempty"`

	// The URL of the image.
	Source *string `json:"source,omitempty"`

	// The title or introductory text to show before the response.
	Title *string `json:"title,omitempty"`

	// The description to show with the the response.
	Description *string `json:"description,omitempty"`

	// The preferred type of control to display.
	Preference *string `json:"preference,omitempty"`

	// An array of objects describing the options from which the user can choose.
	Options []DialogNodeOutputOptionsElement `json:"options,omitempty"`

	// A message to be sent to the human agent who will be taking over the conversation.
	MessageToHumanAgent *string `json:"message_to_human_agent,omitempty"`

	// A label identifying the topic of the conversation, derived from the **user_label** property of the relevant node.
	Topic *string `json:"topic,omitempty"`

	// An array of objects describing the possible matching dialog nodes from which the user can choose.
	//
	// **Note:** The **suggestions** property is part of the disambiguation feature, which is only available for Premium
	// users.
	Suggestions []DialogSuggestion `json:"suggestions,omitempty"`

	// The title or introductory text to show before the response. This text is defined in the search skill configuration.
	Header *string `json:"header,omitempty"`

	// An array of objects containing search results.
	Results []SearchResult `json:"results,omitempty"`
}

// Constants associated with the RuntimeResponseGeneric.ResponseType property.
// The type of response returned by the dialog node. The specified response type must be supported by the client
// application or channel.
//
// **Note:** The **suggestion** response type is part of the disambiguation feature, which is only available for Premium
// users.
const (
	RuntimeResponseGeneric_ResponseType_ConnectToAgent = "connect_to_agent"
	RuntimeResponseGeneric_ResponseType_Image          = "image"
	RuntimeResponseGeneric_ResponseType_Option         = "option"
	RuntimeResponseGeneric_ResponseType_Pause          = "pause"
	RuntimeResponseGeneric_ResponseType_Search         = "search"
	RuntimeResponseGeneric_ResponseType_Suggestion     = "suggestion"
	RuntimeResponseGeneric_ResponseType_Text           = "text"
)

// Constants associated with the RuntimeResponseGeneric.Preference property.
// The preferred type of control to display.
const (
	RuntimeResponseGeneric_Preference_Button   = "button"
	RuntimeResponseGeneric_Preference_Dropdown = "dropdown"
)

// SearchResult : SearchResult struct
type SearchResult struct {

	// The unique identifier of the document in the Discovery service collection.
	//
	// This property is included in responses from search skills, which are a beta feature available only to Plus or
	// Premium plan users.
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

// SearchResultHighlight : An object containing segments of text from search results with query-matching text highlighted using HTML `<em>`
// tags.
type SearchResultHighlight map[string]interface{}

// SetBody : Allow user to set Body
func (this *SearchResultHighlight) SetBody(Body *[]string) {
	(*this)["body"] = Body
}

// GetBody : Allow user to get Body
func (this *SearchResultHighlight) GetBody() *[]string {
	return (*this)["body"].(*[]string)
}

// SetTitle : Allow user to set Title
func (this *SearchResultHighlight) SetTitle(Title *[]string) {
	(*this)["title"] = Title
}

// GetTitle : Allow user to get Title
func (this *SearchResultHighlight) GetTitle() *[]string {
	return (*this)["title"].(*[]string)
}

// SetURL : Allow user to set URL
func (this *SearchResultHighlight) SetURL(URL *[]string) {
	(*this)["url"] = URL
}

// GetURL : Allow user to get URL
func (this *SearchResultHighlight) GetURL() *[]string {
	return (*this)["url"].(*[]string)
}

// SetProperty : Allow user to set arbitrary property
func (this *SearchResultHighlight) SetProperty(Key string, Value *[]string) {
	(*this)[Key] = Value
}

// GetProperty : Allow user to get arbitrary property
func (this *SearchResultHighlight) GetProperty(Key string) *[]string {
	return (*this)[Key].(*[]string)
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

// SessionResponse : SessionResponse struct
type SessionResponse struct {

	// The session ID.
	SessionID *string `json:"session_id" validate:"required"`
}
