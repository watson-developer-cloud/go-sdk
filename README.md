# Watson Developer Cloud Go SDK

Go client library to quickly get started with the various [Watson APIs](https://www.ibm.com/watson/developercloud/) services.

<details>
<summary>Table of Contents</summary>

* [Before you begin](#before-you-begin)
* [Installation](#installation)
* [Examples](#examples)
* [Running in IBM Cloud](#running-in-ibm-cloud)
* [Authentication](#authentication)
	* [Getting-credentials](#getting-credentials)
	* [IAM](#iam)
	* [Username-and-password](#username-and-password)
* [Tests](#tests)
* [Contributing](#contributing)
* [License](#license)

</details>

## Before you begin

* You need an [IBM Cloud](https://console.bluemix.net/developer/watson/dashboard) account.

## Installation

Get SDK package:
```bash
go get github.com/watson-developer-cloud/go-sdk
```

Import the specific service package you want to use in your Go program:
```go
import (
  . "github.com/watson-developer-cloud/go-sdk/discoveryV1"
)
```
(Learn about Go [import types](https://medium.com/golangspec/import-declarations-in-go-8de0fd3ae8ff))

## Examples

The [examples](https://github.ibm.com/arf/go-sdk/tree/master/examples) folder has basic and advanced examples. The examples within each service assume that you already have [service credentials](#getting-credentials).

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
discovery, discoveryErr := NewDiscoveryV1(&ServiceCredentials{
  ServiceURL: "<service_url>",
  Version: "2018-02-16",
  APIkey: "<api_key>",
})
```

**Supplying the access token**

```go
// In the constructor, assuming control of managing IAM token
discovery, discoveryErr := NewDiscoveryV1(&ServiceCredentials{
  ServiceURL: "<service_url>",
  Version: "2018-02-16",
  IAMtoken: "<iam_token>",
})
```

### Username and password

```go
// In the constructor
discovery, discoveryErr := NewDiscoveryV1(&ServiceCredentials{
  ServiceURL: "<service_url>",
  Version: "2018-02-16",
  Username: "<username>",
  Password: "<password>",
})
```

## Tests
Testing is implemented using the [Ginkgo](https://onsi.github.io/ginkgo/) framework.

Run all test suites:
```bash
ginkgo -r
```

Get code coverage for each test suite:
```bash
ginkgo -r -cover
```

Run a specific test suite:
```bash
ginkgo -cover discoveryV1/
```

## Contributing

See [CONTRIBUTING.md](https://github.ibm.com/arf/go-sdk/tree/master/CONTRIBUTING.md).

## License

This library is licensed under the [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0).
