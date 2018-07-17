# Watson Developer Cloud Golang SDK

Golang client library to quickly get started with the various [Watson APIs](https://www.ibm.com/watson/developercloud/) services.

![gopher](https://www.spreadshirt.com/image-server/v1/mp/designs/1005862415,width=178,height=178/golang-gopher.png)

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
* [Contributing](#contributing)
* [License](#license)

</details>

## Before you begin

* You need an [IBM Cloud](https://console.bluemix.net/developer/watson/dashboard) account.

## Installation

Get SDK package:
```bash
go get github.com/watson-developer-cloud/golang-sdk
```

Import SDK package in your Go program:
```go
import (
  watson "github.com/watson-developer-cloud/golang-sdk"
)
```

Import the specific service package you want:
```go
import (
  "github.com/watson-developer-cloud/golang-sdk/discoveryV1"
)
```

## Examples

The examples folder has basic and advanced examples. The examples within each service assume that you already have service credentials.

## Running in IBM Cloud

If you run your app in IBM Cloud, the SDK gets credentials from the ```VCAP_SERVICES``` environment variable.

## Authentication

Watson services are migrating to token-based Identity and Access Management (IAM) authentication.

* With some service instances, you authenticate to the API by using [IAM](#iam).
* In other instances, you authenticate by providing the [username and password](#username-and-password) for the service instance.

### Getting credentials

To find out which authentication to use, view the service credentials. You find the service credentials for authentication the same way for all Watson services:

1. Go to the IBM Cloud [Dashboard](https://console.bluemix.net/dashboard/apps?category=watson) page.
2. Either click an existing Watson service instance or click **Create**.
3. Click **Show** to view your service credentials.
4. Copy the ```url``` and either ```apikey``` or ```username``` and ```password```.

### IAM
IBM Cloud is migrating to token-based Identity and Access Management (IAM) authentication. IAM authentication uses a service API key to get an access token that is passed with the call. Access tokens are valid for approximately one hour and must be regenerated.

You supply either an IAM service **API key** or an **access token**:

* Use the API key to have the SDK manage the lifecycle of the access token. The SDK requests an access token, ensures that the access token is valid, and refreshes it if necessary.

* Use the access token if you want to manage the lifecycle yourself. For details, see [Authenticating with IAM tokens](https://console.bluemix.net/docs/services/watson/getting-started-iam.html#iam).

**Supplying the IAM API key**

```go
// In the constructor, letting the SDK manage the IAM token
discovery, discoveryErr := discoveryV1.NewDiscoveryV1(watson.Credentials{
  ServiceURL: "<service_url>",
  Version: "2018-02-16",
  APIkey: "<api_key>",
})
```

**Supplying the access token**

```go
// In the constructor, assuming control of managing IAM token
discovery, discoveryErr := discoveryV1.NewDiscoveryV1(watson.Credentials{
  ServiceURL: "<service_url>",
  Version: "2018-02-16",
  IAMtoken: "<iam_token>",
})
```

### Username and password

```go
// In the constructor
discovery, discoveryErr := discoveryV1.NewDiscoveryV1(watson.Credentials{
  ServiceURL: "<service_url>",
  Version: "2018-02-16",
  Username: "<username>",
  Password: "<password>",
})
```

## Contributing

See [Contributing.md](https://github.com/watson-developer-cloud/python-sdk/blob/master/CONTRIBUTING.md).

## License

This library is licensed under the [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0).
