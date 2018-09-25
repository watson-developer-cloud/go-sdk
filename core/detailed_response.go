package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DetailedResponse : Generic response for Watson API
type DetailedResponse struct {
	StatusCode int         // HTTP status code
	Headers    http.Header // HTTP response headers
	Result     interface{} // response from service
}

// GetHeaders returns the headers
func (response *DetailedResponse) GetHeaders() http.Header {
	return response.Headers
}

// GetStatusCode returns the HTTP status code
func (response *DetailedResponse) GetStatusCode() int {
	return response.StatusCode
}

// GetResult returns the result from the service
func (response *DetailedResponse) GetResult() interface{} {
	return response.Result
}

func (response *DetailedResponse) String() string {
	output, err := json.MarshalIndent(response, "", "    ")
	if err == nil {
		return fmt.Sprintf("%+v\n", string(output))
	}
	return fmt.Sprintf("Response")
}
