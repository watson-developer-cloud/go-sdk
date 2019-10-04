Here are simple steps to move from `v0.11.0` to `v1.0.0`.

## AUTHENTICATION MECHANISM

The constructor no longer accepts individual credentials like `IAMApiKey`, etc. We initialize authenticators from the [core](https://github.com/IBM/go-sdk-core). The core supports various authentication mechanisms, choose the one appropriate to your instance and use case.

For example, to pass a IAM apikey:
#### Before
```go
service, serviceErr := servicev1.
NewServiceV1(&servicev1.ServiceV1Options{
    URL: "{url}",
    IAMApiKey: "{apikey}",
})
```

#### After(V1.0)
```go
import "github.com/IBM/go-sdk-core/core"

authenticator := &core.IamAuthenticator{
    ApiKey:     "{apikey}",
}

service, serviceErr := servicev1.
NewServiceV1(&servicev1.ServiceV1Options{
    URL: "{url}",
    Authenticator: authenticator,
})
```

There are 5 authentication variants supplied in the SDK (shown below):

#### BasicAuthenticator
```go
import (
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/myservicev1"
)

authenticator := &core.BasicAuthenticator{
    Username: "user1", 
    Password: "password1",
}
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}
service := myservicev1.NewMyServiceV1(options)
```

#### BearerTokenAuthenticator
```go
import (
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/myservicev1"
)

authenticator := &core.BearerTokenAuthenticator{
    BearerToken: token,
}
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}
service := myservicev1.NewMyServiceV1(options)
```

#### CloudPakForDataAuthenticator
```go
import (
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/myservicev1"
)

authenticator := &core.CloudPakForDataAuthenticator{
    URL:                    "https://my-cp4d-url"
    Username:               "user1", 
    Password:               "password1",
    DisableSSLVerification: true,
}
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}
service := myservicev1.NewMyServiceV1(options)
```

#### IAMAuthenticator
```python
import (
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/myservicev1"
)
...
authenticator := &core.IamAuthenticator{
    ApiKey:     "my-iam-apikey",
}
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}
service := myservicev1.NewMyServiceV1(options)
```

#### NoAuthAuthenticator
```go
import (
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/myservicev1"
)
...
options := &myservicev1.MyServiceV1Options{
    Authenticator: &core.NoAuthAuthenticator{},
}
service := myservicev1.NewMyServiceV1(options)
```

#### Creating an Authenticator from Environmental Configuration
External env variables can be obtained from CREDENTIALS, Environment variables or VCAP_SERVICES

```go
import (
    "github.com/IBM/go-sdk-core/core"
    "<appropriate-git-repo-url>/myservicev1"
)
...
authenticator := core.GetAuthenticatorFromEnvironment("my_service")
options := &myservicev1.MyServiceV1Options{
    Authenticator: authenticator,
}
service := myservicev1.NewMyServiceV1(options)
```

## METHOD CALLS RETURN FIRST PARAM AS RESULT
For functions which have a result, method calls will return `result` as the first param

#### Before
```go
response, responseErr := service.MethodCall(&servicev1.MethodCallOptions{})

result := discovery.GetListEnvironmentsResult(response)
```

#### After(V1.0)
Result would be returned directly

```go
result, response, responseErr := service.MethodCall(&servicev1.MethodCallOptions{})
```

## SERVICE CHANGES
#### AssistantV1
* `IncludeCount` is no longer a parameter of the ListWorkspacesOptions
* `IncludeCount` is no longer a parameter of the ListIntentsOptions
* `IncludeCount` is no longer a parameter of the ListExamplesOptions
* `IncludeCount` is no longer a parameter of the ListCounterexamplesOptions
* `IncludeCount` is no longer a parameter of the ListEntitiesOptions
* `IncludeCount` is no longer a parameter of the ListValuesOptions
* `IncludeCount` is no longer a parameter of the ListSynonymsOptions
* `IncludeCount` is no longer a parameter of the ListDialogNodes
* `ValueType` was renamed to `Type` in the CreateValueOptions method
* `ValueType` was renamed to `NewType` in the UpdateValueOptions method
* `NodeType` was renamed to `Type` in the CreateDialogNodeOptions
* `NewNodeType` was renamed to `NewType` in the UpdateDialogNodeOptions
* `ValueType` was renamed to `Type` in the CreateValue model
* `NodeType` was renamed to `Type` in the DialogNode model
* `ActionType` was renamed to `Type` in the DialogNodeAction model
* `QueryType` property was added to the DialogNodeOutputGeneric model
* `Query` property was added to the DialogNodeOutputGeneric model
* `Filter` property was added to the DialogNodeOutputGeneric model
* `DiscoveryVersion` property was added to the DialogNodeOutputGeneric model
* LogMessage model no longer has allows additonal properties
* `DialogRuntimeResponseGeneric` was renamed to `RuntimeResponseGeneric`
* RuntimeEntity model no longer has allows additonal properties
* RuntimeIntent model no longer has allows additonal properties
* `ValueType` was renamed to `Type` in the Value model

#### AssistantV2
* `ActionType` was renamed to `Type` in the DialogNodeAction model
* DialogRuntimeResponseGeneric was renamed to RuntimeResponseGeneric

#### Compare and Comply
* `convertToHTMLOptions` does not require a filename parameter

#### DiscoveryV1
* `ReturnFields` was renamed to `Return` in the QueryOptions
* `LoggingOptOut` was renamed to `XWatsonLoggingOptOut` in the QueryOptions
* `SpellingSuggestions` was added to the QueryOptions
* `collectionIds` is no longer a parameter of the QueryOptions
* `ReturnFields` was renamed to `Return` in the QueryNoticesOptions
* `LoggingOptOut` was renamed to `XWatsonLoggingOptOut` in the FederatedQueryOptions
* `ReturnFields` was renamed to `Return` in the FederatedQueryOptions
* `ReturnFields` was renamed to `Return` in the FederatedQueryNoticesOptions
* `EnrichmentName` was renamed to `Enrichment` in the Enrichment model
* `FieldType` was renamed to `Type` in the Field model
* `FieldName` was renamed to `Field` in the Field model
* TestConfigurationInEnvironment() method was removed
* QueryEntities() method was removed
* QueryRelations() method was removed

#### Language Translator V3
* `DefaultModels` was renamed to `Default` in the ListModelsOptions
* `TranslationOutput` was renamed to `Translation` in the Translation model

#### Natural Language Classifier V1
* `Metadata` was renamed to `TrainingMetadata` in the CreateClassifierOptions

#### Speech to Text V1
* `FinalResults` was renamed to `Final` in the SpeakerLabelsResult model
* `FinalResults` was renamed to `Final` in the SpeechRecognitionResult model
* `customization_id` no longer a param in `recognize_using_websocket()` method

#### Visual Recognition V3
* `DetectFaces()` method was removed
* `ClassName` was renamed to `Class` in the ClassResult model
* `ClassName` was renamed to `Class` in the ModelClass model

#### Visual Recognition V4
* New Service!


