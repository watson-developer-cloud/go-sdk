package common

import (
	"fmt"
	"runtime"
)

const (
	HEADER_SDK_ANALYTICS = "X-IBMCloud-SDK-Analytics"
	HEADER_USER_AGENT    = "User-Agent"

	SDK_NAME = "watson-apis-go-sdk"
)

// GetSdkHeaders - returns the set of SDK-specific headers to be included in an outgoing request.
func GetSdkHeaders(serviceName string, serviceVersion string, operationId string) map[string]string {
	sdkHeaders := make(map[string]string)

	sdkHeaders[HEADER_SDK_ANALYTICS] = fmt.Sprintf("service_name=%s;service_version=%s;operation_id=%s",
		serviceName, serviceVersion, operationId)

	sdkHeaders[HEADER_USER_AGENT] = GetUserAgentInfo()

	return sdkHeaders
}

var userAgent string = fmt.Sprintf("%s-%s %s", SDK_NAME, Version, GetSystemInfo())

func GetUserAgentInfo() string {
	return userAgent
}

var systemInfo = fmt.Sprintf("(arch=%s; os=%s; go.version=%s)", runtime.GOARCH, runtime.GOOS, runtime.Version())

func GetSystemInfo() string {
	return systemInfo
}
