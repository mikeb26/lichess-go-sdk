# \BotApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBotOnline**](BotApi.md#ApiBotOnline) | **Get** /api/bot/online | Get online bots
[**ApiStreamEvent**](BotApi.md#ApiStreamEvent) | **Get** /api/stream/event | Stream incoming events
[**BotAccountUpgrade**](BotApi.md#BotAccountUpgrade) | **Post** /api/bot/account/upgrade | Upgrade to Bot account
[**BotGameAbort**](BotApi.md#BotGameAbort) | **Post** /api/bot/game/{gameId}/abort | Abort a game
[**BotGameChat**](BotApi.md#BotGameChat) | **Post** /api/bot/game/{gameId}/chat | Write in the chat
[**BotGameChatGet**](BotApi.md#BotGameChatGet) | **Get** /api/bot/game/{gameId}/chat | Fetch the game chat
[**BotGameMove**](BotApi.md#BotGameMove) | **Post** /api/bot/game/{gameId}/move/{move} | Make a Bot move
[**BotGameResign**](BotApi.md#BotGameResign) | **Post** /api/bot/game/{gameId}/resign | Resign a game
[**BotGameStream**](BotApi.md#BotGameStream) | **Get** /api/bot/game/stream/{gameId} | Stream Bot game state



## ApiBotOnline

> User ApiBotOnline(ctx).Nb(nb).Execute()

Get online bots



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nb := TODO // interface{} | How many bot users to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.ApiBotOnline(context.Background()).Nb(nb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.ApiBotOnline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiBotOnline`: User
    fmt.Fprintf(os.Stdout, "Response from `BotApi.ApiBotOnline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBotOnlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | [**interface{}**](interface{}.md) | How many bot users to fetch | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStreamEvent

> interface{} ApiStreamEvent(ctx).Execute()

Stream incoming events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.ApiStreamEvent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.ApiStreamEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStreamEvent`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BotApi.ApiStreamEvent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiStreamEventRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotAccountUpgrade

> Ok BotAccountUpgrade(ctx).Execute()

Upgrade to Bot account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotAccountUpgrade(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotAccountUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotAccountUpgrade`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotAccountUpgrade`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotAccountUpgradeRequest struct via the builder pattern


### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameAbort

> Ok BotGameAbort(ctx, gameId).Execute()

Abort a game



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotGameAbort(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotGameAbort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotGameAbort`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotGameAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameChat

> Ok BotGameChat(ctx, gameId).Room(room).Text(text).Execute()

Write in the chat



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := TODO // interface{} | 
    room := TODO // interface{} | 
    text := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotGameChat(context.Background(), gameId).Room(room).Text(text).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotGameChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotGameChat`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotGameChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **room** | [**interface{}**](interface{}.md) |  | 
 **text** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameChatGet

> interface{} BotGameChatGet(ctx, gameId).Execute()

Fetch the game chat



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotGameChatGet(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotGameChatGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotGameChatGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotGameChatGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameChatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameMove

> Ok BotGameMove(ctx, gameId, move).OfferingDraw(offeringDraw).Execute()

Make a Bot move



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := TODO // interface{} | 
    move := TODO // interface{} | The move to play, in UCI format
    offeringDraw := TODO // interface{} | Whether to offer (or agree to) a draw (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotGameMove(context.Background(), gameId, move).OfferingDraw(offeringDraw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotGameMove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotGameMove`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotGameMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 
**move** | [**interface{}**](.md) | The move to play, in UCI format | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offeringDraw** | [**interface{}**](interface{}.md) | Whether to offer (or agree to) a draw | 

### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameResign

> Ok BotGameResign(ctx, gameId).Execute()

Resign a game



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotGameResign(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotGameResign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotGameResign`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotGameResign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameResignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameStream

> interface{} BotGameStream(ctx, gameId).Execute()

Stream Bot game state



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gameId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BotApi.BotGameStream(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BotApi.BotGameStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BotGameStream`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BotApi.BotGameStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

