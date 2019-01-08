# Watson Developer Cloud Go SDK
[![Build Status](https://travis-ci.org/watson-developer-cloud/go-sdk.svg?branch=master)](https://travis-ci.org/watson-developer-cloud/go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)


Go client library to quickly get started with the various [Watson APIs](https://www.ibm.com/watson/developercloud/) services.

<details>
<summary>Table of Contents</summary>

* [Before you begin](#before-you-begin)
* [Installation](#installation)
* [Running in IBM Cloud](#running-in-ibm-cloud)
* [Authentication](#authentication)
	* [Getting-credentials](#getting-credentials)
	* [IAM](#iam)
	* [Username-and-password](#username-and-password)
* [Use](#use)
* [Examples](#examples)
* [AssistantV2](#assistantv2)
* [Tests](#tests)
* [Contributing](#contributing)
* [License](#license)

</details>

## Before you begin

* You need an [IBM Cloud][ibm-cloud-onboarding] account.

## Installation

Get SDK package:
```bash
go get -u github.com/watson-developer-cloud/go-sdk/...
```

## Running in IBM Cloud

If you run your app in IBM Cloud, the SDK gets credentials from the ```VCAP_SERVICES``` environment variable.

## Authentication

Watson services are migrating to token-based Identity and Access Management (IAM) authentication.

* With some service instances, you authenticate to the API by using [IAM](#iam).
* In other instances, you authenticate by providing the [username and password](#username-and-password) for the service instance.

### Getting credentials
To find out which authentication to use, view the service credentials. You find the service credentials for authentication the same way for all Watson services:

1. Go to the IBM Cloud [Dashboard](https://console.bluemix.net/dashboard/apps?category=ai) page.
1. Either click an existing Watson service instance or click [**Create resource > AI**](https://console.bluemix.net/catalog/?category=ai) and create a service instance.
1. Copy the `url` and either `apikey` or `username` and `password`. Click **Show** if the credentials are masked.

### IAM
IBM Cloud is migrating to token-based Identity and Access Management (IAM) authentication. IAM authentication uses a service API key to get an access token that is passed with the call. Access tokens are valid for approximately one hour and must be regenerated.

You supply either an IAM service **API key** or an **access token**:

* Use the API key to have the SDK manage the lifecycle of the access token. The SDK requests an access token, ensures that the access token is valid, and refreshes it if necessary.

* Use the access token if you want to manage the lifecycle yourself. For details, see [Authenticating with IAM tokens](https://console.bluemix.net/docs/services/watson/getting-started-iam.html#iam).

**Supplying the IAM API key**

```go
// In the constructor, letting the SDK manage the IAM token
discovery, discoveryErr := NewDiscoveryV1(&DiscoveryV1Options{
		URL:       "<service_url>",
		Version:   "2018-03-05",
		IAMApiKey: "<apikey>",
	})
```

**Supplying the access token**

```go
// In the constructor, assuming control of managing IAM token
discovery, discoveryErr := NewDiscoveryV1(&DiscoveryV1Options{
		URL:            "<service_url>",
		Version:        "2018-03-05",
		IAMAccessToken: "<IAM_access_token>",
	})
```

### Username and password

```go
// In the constructor
discovery, discoveryErr := NewDiscoveryV1(&DiscoveryV1Options{
		URL:      "<service_url>",
		Version:  "2018-03-05",
		Username: "<username>",
		Password: "<password>",
	})
```

## Use
Apply these general steps for services present in various packages
1. Import the service package
2. Create a new service instance and pass in credentials using either of [authentication](#authentication) methods
3. Invoke API methods using the service instance. For a successful response, it will contain the HTTP status code, response headers and API result
4. Handle responses and errors

```go

package main

import (
"fmt"
"github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

// Creates a Discovery service instance and does a list of environments
func main() {
// Instantiate the Watson Discovery service
service, serviceErr := discoveryv1.
  NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
    URL:       "YOUR SERVICE URL",
    Version:   "2018-03-05",
    IAMApiKey: "YOUR APIKEY",
  })

// Check successful instantiation
if serviceErr != nil {
  panic(serviceErr)
}

// Create a new ListEnvironmentsOptions, these are helper methods.
listEnvironmentsOptions := service.NewListEnvironmentsOptions()

// Call the discovery ListEnvironments method
response, responseErr := service.ListEnvironments(listEnvironmentsOptions)
// Or you can directly pass in the model options
// response, responseErr := service.ListEnvironments(&discoveryv1.ListEnvironmentsOptions{})

// Check successful call
if responseErr != nil {
  panic(responseErr)
}

// This will return the `DetailedResponse`
fmt.Println(response)
// Get the HTTP status code
response.GetStatusCode()
// Get the response headers
response.GetHeaders()
// Get the API response
response.GetResult()

// Cast response to the specific dataType returned by ListEnvironments
// NOTE: most methods have a corresponding Get<methodName>Result() function
listEnvironmentResult := service.GetListEnvironmentsResult(response)

if listEnvironmentResult != nil {
  fmt.Println(listEnvironmentResult.Environments[0])
  }
}
```

## Examples

The [examples][examples] folder has basic and advanced examples. The examples within each service assume that you already have [service credentials](#getting-credentials).

## Tests

Run all test suites:
```bash
go test ./...
```

Get code coverage for each test suite:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Run a specific test suite:
```bash
go test ./assistantv1
```

## Contributing

See [CONTRIBUTING][CONTRIBUTING].

## License

This library is licensed under the [Apache 2.0 license][license].

[ibm-cloud-onboarding]: http://console.bluemix.net/registration?target=%2Fdeveloper%2Fwatson&cm_sp=WatsonPlatform-WatsonServices-_-OnPageNavLink-IBMWatson_SDKs-_-Go
[examples]: https://github.com/watson-developer-cloud/go-sdk/tree/master/examples
[CONTRIBUTING]: https://github.com/watson-developer-cloud/go-sdk/blob/master/.github/CONTRIBUTING.md
[license]: http://www.apache.org/licenses/LICENSE-2.0
