# \BoardApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBoardSeek**](BoardApi.md#ApiBoardSeek) | **Post** /api/board/seek | Create a seek
[**ApiStreamEvent**](BoardApi.md#ApiStreamEvent) | **Get** /api/stream/event | Stream incoming events
[**BoardGameAbort**](BoardApi.md#BoardGameAbort) | **Post** /api/board/game/{gameId}/abort | Abort a game
[**BoardGameBerserk**](BoardApi.md#BoardGameBerserk) | **Post** /api/board/game/{gameId}/berserk | Berserk a tournament game
[**BoardGameChatGet**](BoardApi.md#BoardGameChatGet) | **Get** /api/board/game/{gameId}/chat | Fetch the game chat
[**BoardGameChatPost**](BoardApi.md#BoardGameChatPost) | **Post** /api/board/game/{gameId}/chat | Write in the chat
[**BoardGameClaimVictory**](BoardApi.md#BoardGameClaimVictory) | **Post** /api/board/game/{gameId}/claim-victory | Claim victory of a game
[**BoardGameDraw**](BoardApi.md#BoardGameDraw) | **Post** /api/board/game/{gameId}/draw/{accept} | Handle draw offers
[**BoardGameMove**](BoardApi.md#BoardGameMove) | **Post** /api/board/game/{gameId}/move/{move} | Make a Board move
[**BoardGameResign**](BoardApi.md#BoardGameResign) | **Post** /api/board/game/{gameId}/resign | Resign a game
[**BoardGameStream**](BoardApi.md#BoardGameStream) | **Get** /api/board/game/stream/{gameId} | Stream Board game state
[**BoardGameTakeback**](BoardApi.md#BoardGameTakeback) | **Post** /api/board/game/{gameId}/takeback/{accept} | Handle takeback offers



## ApiBoardSeek

> ApiBoardSeek(ctx).Rated(rated).Time(time).Increment(increment).Days(days).Variant(variant).Color(color).RatingRange(ratingRange).Execute()

Create a seek



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
    rated := TODO // interface{} | Whether the game is rated and impacts players ratings. (optional) (default to false)
    time := TODO // interface{} | Clock initial time in minutes. Required for real-time seeks. (optional)
    increment := TODO // interface{} | Clock increment in seconds. Required for real-time seeks. (optional)
    days := TODO // interface{} | Days per turn. Required for correspondence seeks. (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    color := TODO // interface{} | The color to play. Better left empty to automatically get 50% white. (optional) (default to random)
    ratingRange := TODO // interface{} | The rating range of potential opponents. Better left empty. Example: 1500-1800  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoardApi.ApiBoardSeek(context.Background()).Rated(rated).Time(time).Increment(increment).Days(days).Variant(variant).Color(color).RatingRange(ratingRange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.ApiBoardSeek``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBoardSeekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rated** | [**interface{}**](interface{}.md) | Whether the game is rated and impacts players ratings. | [default to false]
 **time** | [**interface{}**](interface{}.md) | Clock initial time in minutes. Required for real-time seeks. | 
 **increment** | [**interface{}**](interface{}.md) | Clock increment in seconds. Required for real-time seeks. | 
 **days** | [**interface{}**](interface{}.md) | Days per turn. Required for correspondence seeks. | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **color** | [**interface{}**](interface{}.md) | The color to play. Better left empty to automatically get 50% white. | [default to random]
 **ratingRange** | [**interface{}**](interface{}.md) | The rating range of potential opponents. Better left empty. Example: 1500-1800  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain, application/json

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
    resp, r, err := apiClient.BoardApi.ApiStreamEvent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.ApiStreamEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStreamEvent`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.ApiStreamEvent`: %v\n", resp)
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


## BoardGameAbort

> Ok BoardGameAbort(ctx, gameId).Execute()

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
    resp, r, err := apiClient.BoardApi.BoardGameAbort(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameAbort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameAbort`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameAbortRequest struct via the builder pattern


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


## BoardGameBerserk

> Ok BoardGameBerserk(ctx, gameId).Execute()

Berserk a tournament game



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
    resp, r, err := apiClient.BoardApi.BoardGameBerserk(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameBerserk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameBerserk`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameBerserk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameBerserkRequest struct via the builder pattern


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


## BoardGameChatGet

> interface{} BoardGameChatGet(ctx, gameId).Execute()

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
    resp, r, err := apiClient.BoardApi.BoardGameChatGet(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameChatGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameChatGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameChatGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameChatGetRequest struct via the builder pattern


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


## BoardGameChatPost

> Ok BoardGameChatPost(ctx, gameId).Room(room).Text(text).Execute()

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
    resp, r, err := apiClient.BoardApi.BoardGameChatPost(context.Background(), gameId).Room(room).Text(text).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameChatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameChatPost`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameChatPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameChatPostRequest struct via the builder pattern


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


## BoardGameClaimVictory

> Ok BoardGameClaimVictory(ctx, gameId).Execute()

Claim victory of a game



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
    resp, r, err := apiClient.BoardApi.BoardGameClaimVictory(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameClaimVictory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameClaimVictory`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameClaimVictory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameClaimVictoryRequest struct via the builder pattern


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


## BoardGameDraw

> Ok BoardGameDraw(ctx, gameId, accept).Execute()

Handle draw offers



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
    accept := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoardApi.BoardGameDraw(context.Background(), gameId, accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameDraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameDraw`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameDraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 
**accept** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameDrawRequest struct via the builder pattern


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


## BoardGameMove

> Ok BoardGameMove(ctx, gameId, move).OfferingDraw(offeringDraw).Execute()

Make a Board move



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
    resp, r, err := apiClient.BoardApi.BoardGameMove(context.Background(), gameId, move).OfferingDraw(offeringDraw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameMove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameMove`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 
**move** | [**interface{}**](.md) | The move to play, in UCI format | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameMoveRequest struct via the builder pattern


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


## BoardGameResign

> Ok BoardGameResign(ctx, gameId).Execute()

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
    resp, r, err := apiClient.BoardApi.BoardGameResign(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameResign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameResign`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameResign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameResignRequest struct via the builder pattern


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


## BoardGameStream

> interface{} BoardGameStream(ctx, gameId).Execute()

Stream Board game state



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
    resp, r, err := apiClient.BoardApi.BoardGameStream(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameStream`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameStreamRequest struct via the builder pattern


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


## BoardGameTakeback

> Ok BoardGameTakeback(ctx, gameId, accept).Execute()

Handle takeback offers



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
    accept := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoardApi.BoardGameTakeback(context.Background(), gameId, accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoardApi.BoardGameTakeback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BoardGameTakeback`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BoardApi.BoardGameTakeback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 
**accept** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameTakebackRequest struct via the builder pattern


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

