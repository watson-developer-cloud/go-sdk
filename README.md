# Watson Developer Cloud Go SDK v2.0.0
[![Build Status](https://travis-ci.org/watson-developer-cloud/go-sdk.svg?branch=master)](https://travis-ci.org/watson-developer-cloud/go-sdk)
[![Release](https://img.shields.io/github/v/release/watson-developer-cloud/go-sdk)](https://github.com/watson-developer-cloud/go-sdk/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/watson-developer-cloud/go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![wdc-community.slack.com](https://wdc-slack-inviter.mybluemix.net/badge.svg)](http://wdc-slack-inviter.mybluemix.net/)


Go client library to quickly get started with the various [Watson APIs](https://www.ibm.com/watson/developercloud/) services.

<details>
<summary>Table of Contents</summary>

* [Before you begin](#before-you-begin)
* [Installation](#installation)
* [Preview of 2.0.0 release](#preview-release-for-version-2.0.0)
* [Personality Insights Deprecation](#personality-insights-deprecation)
* [Discovery v2 only on CP4D](#discovery-v2-only-on-cp4d)
* [Running in IBM Cloud](#running-in-ibm-cloud)
* [Authentication](#authentication)
	* [Getting-credentials](#getting-credentials)
	* [IAM](#iam)
	* [Username-and-password](#username-and-password)
* [Use](#use)
* [Configuring the HTTP Client](#configuring-the-http-client)
* [Disable SSL certificate verification](#disable-ssl-certificate-verification)
* [Set Service URL](#set-service-url)
* [Using the global transaction ID header](#using-the-global-transaction-id)
* [IBM Cloud Pak for Data(ICP4D)](#ibm-cloud-pak-for-data(icp4d))
* [Examples](#examples)
* [Tests](#tests)
* [Contributing](#contributing)
* [Migrating](#migrating)
* [License](#license)
* [Featured Projects](#featured-projects)

</details>

## ANNOUNCEMENTS!
### Personality Insights deprecation
IBM Watson™ Personality Insights is discontinued. For a period of one year from 1 December 2020, you will still be able to use Watson Personality Insights. However, as of 1 December 2021, the offering will no longer be available.

As an alternative, we encourage you to consider migrating to IBM Watson™ [Natural Language Understanding](https://cloud.ibm.com/docs/natural-language-understanding), a service on IBM Cloud® that uses deep learning to extract data and insights from text such as keywords, categories, sentiment, emotion, and syntax to provide insights for your business or industry. For more information, see About Natural Language Understanding.

### Visual Recognition deprecation
IBM Watson™ Visual Recognition is discontinued. Existing instances are supported until 1 December 2021, but as of 7 January 2021, you can't create instances. Any instance that is provisioned on 1 December 2021 will be deleted.

## Before you begin

* You need an [IBM Cloud][ibm-cloud-onboarding] account.

## Installation

Get SDK package:
```bash
go get -u github.com/watson-developer-cloud/go-sdk@v2.0.0
```

Note: For the latest tag release, look into [examples][examples] folder for basic and advanced examples.

## Preview release for version `2.0.0`

We have a preview release available for the upcoming version `2.0.0`. This introduces a number of breaking changes to the SDK, and adds some convenient new features for all packages in the SDK. See `Migrating-V2` at the root of the SDK for details on what is changing.

If you want to try out the new version before it is released, you can add it as a dependency by targeting the branch directly in your `go.mod`. The easier way to do this is by running this command:

```sh
go get github.com/watson-developer-cloud/go-sdk@c2fbbe350edbc8f9a6ce6325d4c7880c3c1c538d
```

If you decide to preview the new release, and have any feedback or comments they are greatly appreciated!

## Personality Insights Deprecation

IBM® will begin sunsetting IBM Watson™ Personality Insights on 1 December 2020. For a period of one year from this date, you will still be able to use Watson Personality Insights. However, as of 1 December 2021, the offering will no longer be available.

As an alternative, we encourage you to consider migrating to IBM Watson™ Natural Language Understanding, a service on IBM Cloud® that uses deep learning to extract data and insights from text such as keywords, categories, sentiment, emotion, and syntax to provide insights for your business or industry.

## Discovery v2 only on CP4D
Discovery v2 is only available on Cloud Pak for Data.

## Running in IBM Cloud

If you run your app in IBM Cloud, the SDK gets credentials from the ```VCAP_SERVICES``` environment variable.

## Authentication

Watson services are migrating to token-based Identity and Access Management (IAM) authentication.

* With some service instances, you authenticate to the API by using [IAM](#iam).
* In other instances, you authenticate by providing the [username and password](#username-and-password) for the service instance.

### Getting credentials
To find out which authentication to use, view the service credentials. You find the service credentials for authentication the same way for all Watson services:

1. Go to the IBM Cloud [Dashboard](https://cloud.ibm.com/) page.
1. Either click an existing Watson service instance in your [resource list](https://cloud.ibm.com/resources) or click [**Create resource > AI**](https://cloud.ibm.com/catalog?category=ai) and create a service instance.
1. Click on the **Manage** item in the left nav bar of your service instance.

On this page, you should be able to see your credentials for accessing your service instance.

### Supplying credentials

There are two ways to supply the credentials you found above to the SDK for authentication.

#### Credential file

With a credential file, you just need to put the file in the right place and the SDK will do the work of parsing and authenticating. You can get this file by clicking the **Download** button for the credentials in the **Manage** tab of your service instance.

The file downloaded will be called `ibm-credentials.env`. This is the name the SDK will search for and **must** be preserved unless you want to configure the file path (more on that later). The SDK will look for your `ibm-credentials.env` file in the following places (in order):

- The top-level directory of the project you're using the SDK in
- Your system's home directory

As long as you set that up correctly, you don't have to worry about setting any authentication options in your code. So, for example, if you created and downloaded the credential file for your Discovery instance, you just need to do the following:

```go
import "github.com/watson-developer-cloud/go-sdk/servicev1"

service, serviceErr := servicev1.NewServiceV1(&servicev1.ServiceV1Options{
	Version:   "2018-03-05",
})
```

And that's it!

If you're using more than one service at a time in your code and get two different `ibm-credentials.env` files, just put the contents together in one `ibm-credentials.env` file and the SDK will handle assigning credentials to their appropriate services.

If you would like to configure the location/name of your credential file, you can set an environment variable called `IBM_CREDENTIALS_FILE`. **This will take precedence over the locations specified above.** Here's how you can do that:

```bash
export IBM_CREDENTIALS_FILE="<path>"
```

where `<path>` is something like `/home/user/Downloads/<file_name>.env`.

#### Manually
If you'd prefer to set authentication values manually in your code, the SDK supports that as well. The way you'll do this depends on what type of credentials your service instance gives you.

### IAM
IBM Cloud is migrating to token-based Identity and Access Management (IAM) authentication. IAM authentication uses a service API key to get an access token that is passed with the call. Access tokens are valid for approximately one hour and must be regenerated.

You supply either an IAM service **API key** or an **access token**:

* Use the API key to have the SDK manage the lifecycle of the access token. The SDK requests an access token, ensures that the access token is valid, and refreshes it if necessary.

* Use the access token if you want to manage the lifecycle yourself. For details, see [Authenticating with IAM tokens](https://cloud.ibm.com/docs/watson?topic=watson-iam#iam).

**Supplying the IAM API key**

```go
// In the constructor, letting the SDK manage the IAM token
import "github.com/watson-developer-cloud/go-sdk/servicev1"
authenticator := &core.IamAuthenticator{
	ApiKey: "<apikey>",
}
service, serviceErr := servicev1.NewServiceV1(&servicev1.ServiceV1Options{
		URL:       "<service_url>",
		Authenticator: authenticator,
	})
```

### Username and password

```go
// In the constructor
import "github.com/watson-developer-cloud/go-sdk/servicev1"

service, serviceErr := servicev1.NewServiceV1(&servicev1.ServiceV1Options{
		URL:      "<service_url>",
		Authenticator: &core.BasicAuthenticator{
			Username: "<username>",
			Password: "<password>",
		},
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
    Version:   "2018-03-05",
	Authenticator: &core.IamAuthenticator{
		Apikey: "YOUR APIKEY",
	},
  })
service.SetServiceURL("YOUR SERVICE URL")

// Check successful instantiation
if serviceErr != nil {
  panic(serviceErr)
}

// Create a new ListEnvironmentsOptions, these are helper methods.
listEnvironmentsOptions := service.NewListEnvironmentsOptions()

// Call the discovery ListEnvironments method
listEnvironmentResult, response, responseErr := service.ListEnvironments(listEnvironmentsOptions)
// Or you can directly pass in the model options
// listEnvironmentResult, response, responseErr := service.ListEnvironments(&discoveryv1.ListEnvironmentsOptions{})

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

// Directly access the result
if listEnvironmentResult != nil {
  fmt.Println(listEnvironmentResult.Environments[0])
  }
}
```

## Configuring the HTTP Client

To change client configs like timeout, setting proxy, etc, pass in your own client using the `SetHTTPClient()` method. Below is an example to pass a proxy

```go
package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"

  discovery "github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func main() {
  // Instantiate the Watson Discovery service
  discoveryService, discoveryServiceErr := discovery.NewDiscoveryV1(&discovery.DiscoveryV1Options{
    URL:       "YOUR SERVICE URL",
    Version:   "2018-03-05",
	Authenticator: &core.IamAuthenticator{
		Apikey: "YOUR APIKEY",
	},
  })
  // Check successful instantiation
  if discoveryServiceErr != nil {
    fmt.Println(discoveryServiceErr)
    return
  }

  //creating the proxyURL
  proxyStr := "http://<proxy host>:<port>"
  proxyURL, err := url.Parse(proxyStr)
  if err != nil {
    log.Println(err)
  }

  //adding the proxy settings to the Transport object
  transport := &http.Transport{
    Proxy: http.ProxyURL(proxyURL),
  }

  //adding the Transport object to the http Client
  client := &http.Client{
    Transport: transport,
  }

  discoveryService.Service.SetHTTPClient(client)
}
```

## Disable SSL certificate verification
Disable the SSL verification using `DisableSSLVerification()` method

```go
package main

import (
  "fmt"
  discovery "github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func main() {
  // Instantiate the Watson Discovery service
  discoveryService, discoveryServiceErr := discovery.NewDiscoveryV1(&discovery.DiscoveryV1Options{
    URL:       "YOUR SERVICE URL",
    Version:   "2018-03-05",
	Authenticator: &core.IamAuthenticator{
		Apikey: "YOUR APIKEY",
	},
  })
  // Check successful instantiation
  if discoveryServiceErr != nil {
    fmt.Println(discoveryServiceErr)
    return
  }

  discoveryService.Service.DisableSSLVerification()
}
```

Or can set it from external sources. For example, environment variables:

```bash
export <YOUR SERVICE NAME>_DISABLE_SSL=True

## Set Service URL
To set the service URL,

```go
service.SetServiceURL("my new url")
```

Or can set it from external sources. For example, environment variables:

```bash
export <YOUR SERVICE NAME>_URL="my new url"
```

## Using the global transaction ID
Every SDK call returns a response with a transaction ID in the x-global-transaction-id header. This transaction ID is useful for troubleshooting and accessing relevant logs from your service instance.

```go
result, response, responseErr := service.MyServiceCall(myServiceOptions)
fmt.Println(response.GetHeaders().Get("X-Global-Transaction-Id"))
```

Additionally, you can set a custom `X-Global-Transaction-Id` header in your request. This can make it useful if you need to provide a transaction ID to IBM support to help troubleshoot an issue.

```go
service, err = assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
	Version:     "2020-04-01",
	ServiceName: "assistant",
})

customHeaders := http.Header{}
customHeaders.Add("X-Global-Transaction-Id", "my-custom-value")
service.Service.SetDefaultHeaders(customHeaders)
```

## Cloud Pak for Data(CP4D)
If your service instance is of ICP4D, below are two ways of initializing the assistant service.

1) Supplying the `Username`, `Password`, `ICP4DURL` and `AuthenticationType`

```go
package main

import (
  "fmt"
  discovery "github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func main() {
  // Instantiate the Watson Discovery service
	authenticator := &core.CloudPakForDataAuthenticator{
		URL:                    "CP4D URL FOR AUTHENTICATION", // should be of the form https://{icp_cluster_host}
		Username:               "YOUR USERNAME", 
		Password:               "YOUR PASSWORD",
		DisableSSLVerification: true,
	}
  discoveryService, discoveryServiceErr := discovery.NewDiscoveryV1(&discovery.DiscoveryV1Options{
    URL:                "YOUR SERVICE URL", // should be of the form https://{icp_cluster_host}/{deployment}/discovery/{instance-id}/api
    Version:            "2018-03-05",
	Authenticator: authenticator,
  })

  // Check successful instantiation
  if discoveryServiceErr != nil {
    fmt.Println(discoveryServiceErr)
    return
  }

  discoveryService.DisableSSLVerification() // MAKE SURE SSL VERIFICATION IS DISABLED
}
```

2) Supplying the access token
```go
package main

import (
  "fmt"
  discovery "github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func main() {
	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: "YOUR BEARER TOKEN",
	}

  // Instantiate the Watson Discovery service
  discoveryService, discoveryServiceErr := discovery.NewDiscoveryV1(&discovery.DiscoveryV1Options{
    URL:              "YOUR SERVICE URL", // should be of the form https://{icp_cluster_host}/{deployment}/discovery/{instance-id}/api
    Version:          "2018-03-05",
    Authenticator: authenticator,
  })
  // Check successful instantiation
  if discoveryServiceErr != nil {
    fmt.Println(discoveryServiceErr)
    return
  }

  discoveryService.DisableSSLVerification() // MAKE SURE SSL VERIFICATION IS DISABLED
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

## Questions

If you have issues with the APIs or have a question about the Watson services, see [Stack Overflow](https://stackoverflow.com/questions/tagged/ibm-watson+go).

## Migrating
* To move from `v0.11.0` to `v1.0.0`, refer to the [MIGRATION-V1](https://github.com/watson-developer-cloud/go-sdk/blob/master/MIGRATION-V1.md)

## Contributing

See [CONTRIBUTING][CONTRIBUTING].

## Featured Projects

We'd love to highlight cool open-source projects that use this SDK! If you'd like to get your project added to the list, feel free to make an issue linking us to it.

## License

This library is licensed under the [Apache 2.0 license][license].

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration?target=%2Fdeveloper%2Fwatson&cm_sp=WatsonPlatform-WatsonServices-_-OnPageNavLink-IBMWatson_SDKs-_-Go
[examples]: https://github.com/watson-developer-cloud/go-sdk/tree/master/examples
[CONTRIBUTING]: https://github.com/watson-developer-cloud/go-sdk/blob/master/.github/CONTRIBUTING.md
[license]: http://www.apache.org/licenses/LICENSE-2.0
