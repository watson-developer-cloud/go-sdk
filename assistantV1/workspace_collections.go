package assistantV1

import (
	"time"
	"net/http"
)

type WatsonResponse struct {
	StatusCode int
	Headers http.Header
	Result interface{}
}

type DialogNodeAction struct {

	// The name of the action.
	Name string `json:"name"`

	// The type of action to invoke.
	Type string `json:"type,omitempty"`

	// A map of key/value pairs to be provided to the action.
	Parameters interface{} `json:"parameters,omitempty"`

	// The location in the dialog context where the result of the action is stored.
	ResultVariable string `json:"result_variable"`

	// The name of the context variable that the client application will use to pass in credentials for the action.
	Credentials string `json:"credentials,omitempty"`
}

// The next step to execute following this dialog node.
type DialogNodeNextStep struct {

	// What happens after the dialog node completes. The valid values depend on the node type:  - The following values are valid for any node:    - `get_user_input`    - `skip_user_input`    - `jump_to`  - If the node is of type `event_handler` and its parent node is of type `slot` or `frame`, additional values are also valid:    - if **event_name**=`filled` and the type of the parent node is `slot`:      - `reprompt`      - `skip_all_slots`  - if **event_name**=`nomatch` and the type of the parent node is `slot`:      - `reprompt`      - `skip_slot`      - `skip_all_slots`  - if **event_name**=`generic` and the type of the parent node is `frame`:      - `reprompt`      - `skip_slot`      - `skip_all_slots`        If you specify `jump_to`, then you must also specify a value for the `dialog_node` property.
	Behavior string `json:"behavior"`

	// The ID of the dialog node to process next. This parameter is required if **behavior**=`jump_to`.
	DialogNode string `json:"dialog_node,omitempty"`

	// Which part of the dialog node to process next.
	Selector string `json:"selector,omitempty"`
}

type CreateCounterexample struct {

	// The text of a user input marked as irrelevant input. This string must conform to the following restrictions:  - It cannot contain carriage return, newline, or tab characters  - It cannot consist of only whitespace characters  - It must be no longer than 1024 characters
	Text string `json:"text"`
}

type CreateDialogNode struct {

	// The dialog node ID. This string must conform to the following restrictions:  - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.  - It must be no longer than 1024 characters.
	DialogNode string `json:"dialog_node"`

	// The description of the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// The condition that will trigger the dialog node. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 2048 characters.
	Conditions string `json:"conditions,omitempty"`

	// The ID of the parent dialog node.
	Parent string `json:"parent,omitempty"`

	// The ID of the previous dialog node.
	PreviousSibling string `json:"previous_sibling,omitempty"`

	// The output of the dialog node. For more information about how to specify dialog node output, see the [documentation](https://console.bluemix.net/docs/services/conversation/dialog-overview.html#complex).
	Output interface{} `json:"output,omitempty"`

	// The context for the dialog node.
	Context interface{} `json:"context,omitempty"`

	// The metadata for the dialog node.
	Metadata interface{} `json:"metadata,omitempty"`

	// The next step to be executed in dialog processing.
	NextStep DialogNodeNextStep `json:"next_step,omitempty"`

	// The alias used to identify the dialog node. This string must conform to the following restrictions:  - It can contain only Unicode alphanumeric, space, underscore, hyphen, and dot characters.  - It must be no longer than 64 characters.
	Title string `json:"title,omitempty"`

	// How the dialog node is processed.
	Type string `json:"type,omitempty"`

	// How an `event_handler` node is processed.
	EventName string `json:"event_name,omitempty"`

	// The location in the dialog context where output is stored.
	Variable string `json:"variable,omitempty"`

	// An array of objects describing any actions to be invoked by the dialog node.
	Actions []DialogNodeAction `json:"actions,omitempty"`

	// Whether this top-level dialog node can be digressed into.
	DigressIn string `json:"digress_in,omitempty"`

	// Whether this dialog node can be returned to after a digression.
	DigressOut string `json:"digress_out,omitempty"`

	// Whether the user can digress to top-level nodes while filling out slots.
	DigressOutSlots string `json:"digress_out_slots,omitempty"`
}

type CreateValue struct {

	// The text of the entity value. This string must conform to the following restrictions:  - It cannot contain carriage return, newline, or tab characters.  - It cannot consist of only whitespace characters.  - It must be no longer than 64 characters.
	Value string `json:"value"`

	// Any metadata related to the entity value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array containing any synonyms for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A synonym must conform to the following restrictions:  - It cannot contain carriage return, newline, or tab characters.  - It cannot consist of only whitespace characters.  - It must be no longer than 64 characters.
	Synonyms []string `json:"synonyms,omitempty"`

	// An array of patterns for the entity value. You can provide either synonyms or patterns (as indicated by **type**), but not both. A pattern is a regular expression no longer than 512 characters. For more information about how to specify a pattern, see the [documentation](https://console.bluemix.net/docs/services/conversation/entities.html#creating-entities).
	Patterns []string `json:"patterns,omitempty"`

	// Specifies the type of value.
	Type string `json:"type,omitempty"`
}

type CreateEntity struct {

	// The name of the entity. This string must conform to the following restrictions:  - It can contain only Unicode alphanumeric, underscore, and hyphen characters.  - It cannot begin with the reserved prefix `sys-`.  - It must be no longer than 64 characters.
	Entity string `json:"entity"`

	// The description of the entity. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// Any metadata related to the value.
	Metadata interface{} `json:"metadata,omitempty"`

	// An array of objects describing the entity values.
	Values []CreateValue `json:"values,omitempty"`

	// Whether to use fuzzy matching for the entity.
	FuzzyMatch bool `json:"fuzzy_match,omitempty"`
}

type CreateExample struct {

	// The text of a user input example. This string must conform to the following restrictions:  - It cannot contain carriage return, newline, or tab characters.  - It cannot consist of only whitespace characters.  - It must be no longer than 1024 characters.
	Text string `json:"text"`
}

type CreateIntent struct {

	// The name of the intent. This string must conform to the following restrictions:  - It can contain only Unicode alphanumeric, underscore, hyphen, and dot characters.  - It cannot begin with the reserved prefix `sys-`.  - It must be no longer than 128 characters.
	Intent string `json:"intent"`

	// The description of the intent. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// An array of user input examples for the intent.
	Examples []CreateExample `json:"examples,omitempty"`
}

type CreateWorkspace struct {

	// The name of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 64 characters.
	Name string `json:"name,omitempty"`

	// The description of the workspace. This string cannot contain carriage return, newline, or tab characters, and it must be no longer than 128 characters.
	Description string `json:"description,omitempty"`

	// The language of the workspace.
	Language string `json:"language,omitempty"`

	// An array of objects defining the intents for the workspace.
	Intents []CreateIntent `json:"intents,omitempty"`

	// An array of objects defining the entities for the workspace.
	Entities []CreateEntity `json:"entities,omitempty"`

	// An array of objects defining the nodes in the workspace dialog.
	DialogNodes []CreateDialogNode `json:"dialog_nodes,omitempty"`

	// An array of objects defining input examples that have been marked as irrelevant input.
	Counterexamples []CreateCounterexample `json:"counterexamples,omitempty"`

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`
}

type Workspace struct {

	// The name of the workspace.
	Name string `json:"name"`

	// The language of the workspace.
	Language string `json:"language"`

	// The timestamp for creation of the workspace.
	Created time.Time `json:"created,omitempty"`

	// The timestamp for the last update to the workspace.
	Updated time.Time `json:"updated,omitempty"`

	// The workspace ID.
	WorkspaceID string `json:"workspace_id"`

	// The description of the workspace.
	Description string `json:"description,omitempty"`

	// Any metadata related to the workspace.
	Metadata interface{} `json:"metadata,omitempty"`

	// Whether training data from the workspace (including artifacts such as intents and entities) can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out,omitempty"`
}

type Pagination struct {

	// The URL that will return the same page of results.
	RefreshURL string `json:"refresh_url"`

	// The URL that will return the next page of results.
	NextURL string `json:"next_url,omitempty"`

	// Reserved for future use.
	Total int32 `json:"total,omitempty"`

	// Reserved for future use.
	Matched int32 `json:"matched,omitempty"`

	// A token identifying the current page of results.
	RefreshCursor string `json:"refresh_cursor,omitempty"`

	// A token identifying the next page of results.
	NextCursor string `json:"next_cursor,omitempty"`
}

type ListWorkspacesRequest struct {
	Sort string `json:"sort"`
	Cursor string `json:"cursor"`
	PageLimit int `json:"page_limit"`
	IncludeCount bool `json:"include_count"`
	IncludeAudit bool `json:"include_audit"`
}

type ListWorkspacesResponse struct {

	// An array of objects describing the workspaces associated with the service instance.
	Workspaces []Workspace `json:"workspaces"`

	// An object defining the pagination data for the returned objects.
	Pagination Pagination `json:"pagination"`
}

type GetWorkspaceRequest struct {
	Export bool `json:"export"`
	IncludeAudit bool `json:"include_audit"`
}

type GetWorkspaceResponse struct {

	// The name of the workspace.
	Name string `json:"name"`

	// The description of the workspace.
	Description string `json:"description"`

	// The language of the workspace.
	Language string `json:"language"`

	// Any metadata that is required by the workspace.
	Metadata interface{} `json:"metadata"`

	// The workspace ID.
	WorkspaceID string `json:"workspace_id"`

	// The current status of the workspace.
	Status string `json:"status"`

	// Whether training data from the workspace can be used by IBM for general service improvements. `true` indicates that workspace training data is not to be used.
	LearningOptOut bool `json:"learning_opt_out"`
}

type UpdateWorkspaceRequest struct {
	Append bool `json:"append"`
}
