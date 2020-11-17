# Migrating to version 2.0.0

In order to move from versions `1.x` to version `2.0.0`, there are some breaking changes that need to be accounted for. This guide will address overall changes happening to all service packages, and then call out important changes to individual packages.

## Overall Changes to all Packages

### Breaking `go-sdk-core` version change

For all packages, we now import the `go-sdk-core/v4` module as a dependency. This replaces `go-sdk-core`. Any code that previously imported the Go Core should replace their existing imports with:

```
"github.com/IBM/go-sdk-core/v4/core"
```

### Breaking Version Date Type Change

For all packages, the `Version` property on the service struct is now a `*String`. It had previously been a `String` which was provided as a literal. The Go SDK Core has a convenience method to provide this string pointer:

```go
service, err = assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
	Version:     core.StringPtr("2020-04-01"),
	ServiceName: "assistant",
 })
```

### Breaking Changes to model setter functions

In the `1.x` version of the SDK, setter methods were provided for almost all model structs for dealing with data from the service. Version `2.0.0` makes significant breaking changes in how setters are handled for these structs.

Specifically, now the only structs that have setters for their properties are structs with the naming convention `{ServiceOperation}Options`. These are called `Options Models` and are typically the top level struct that is passed into a function. These structs have setter methods exposed that allow the individual properties within them to be set.

For lower level models underneath these options models, they had setter methods for each individual property. As this bloated the code surface and provided little value compared to setting the property directly, we have decided in version `2.0.0` to remove these setters.

If your code editor calls out that a setter method you are using has been removed, the appropriate way to update is to set the property directly.

### Providing custom `Context` to functions

Version `2.0.0` of the Go SDK now includes an alternate version of each operation belonging to the service, named `<operation>WithContext`. This alternate function allows you to pass in a `context.Context` parameter while invoking the operation.  The `context.Context` parameter can be used to specify a timeout, or cancel an in-flight request.  Details about `context.Context` can be found [here](https://golang.org/pkg/context).

Internally, the Go SDK uses these functions for all operations, but the regular function provides the background context as a default.

### Automatic Retries

All service structs can be configured to retry HTTP methods in the event of an error.

You can customize the retry logic to suit your purposes. The function `EnableRetries(maxRetries int, mayTretryInterval time.Duration)` allows you to configure the retry strategy that you would like to enable for the functions called by the service.

You can also disable retries by using the `DisableRetries()` method exposed on the service struct.

## Service specific changes

### assistantv1

- In the `Message()` function, the MessageInput{}'s `Text` property is now a `*string`.

### assistantv2

- In the `Message()` function, the MessageInput{}'s `Text` property is now a `*string`. This can be constructed via the `core.StringPtr("{message}")` convenience method.
- The `Message()` function's `MessageContext{}` struct has a property `Skills` that now accepts a `map[string]assistantv2.MessageContextSkill{}` type

### texttospeechv1

- The `VoiceModel` naming convention has been replaced by `CustomModel` for all models and functions. For example, `CreateVoiceModel` in the `1.x` branch is now `CreateCustomModel` in version `2.0.0`.
