# \BulkPairingsApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkPairingCreate**](BulkPairingsApi.md#BulkPairingCreate) | **Post** /api/bulk-pairing | Create a bulk pairing
[**BulkPairingDelete**](BulkPairingsApi.md#BulkPairingDelete) | **Delete** /api/bulk-pairing/{id} | Cancel a bulk pairing
[**BulkPairingGet**](BulkPairingsApi.md#BulkPairingGet) | **Get** /api/bulk-pairing | View upcoming bulk pairings
[**BulkPairingStartClocks**](BulkPairingsApi.md#BulkPairingStartClocks) | **Post** /api/bulk-pairing/{id}/start-clocks | Manually start clocks



## BulkPairingCreate

> interface{} BulkPairingCreate(ctx).Players(players).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).PairAt(pairAt).StartClocksAt(startClocksAt).Rated(rated).Variant(variant).Fen(fen).Message(message).Rules(rules).Execute()

Create a bulk pairing



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
    players := TODO // interface{} | OAuth tokens of all the players to pair, with the syntax `tokenOfWhitePlayerInGame1:tokenOfBlackPlayerInGame1,tokenOfWhitePlayerInGame2:tokenOfBlackPlayerInGame2,...`.  The 2 tokens of the players of a game are separated with `:`. The first token gets the white pieces. Games are separated with `,`.  Up to 1000 tokens can be sent, for a max of 500 games.  Each token must be included at most once.  Example: `token1:token2,token3:token4,token5:token6`  (optional)
    clockLimit := TODO // interface{} | Clock initial time in seconds. Example: `600`  (optional)
    clockIncrement := TODO // interface{} | Clock increment in seconds. Example: `2`  (optional)
    days := TODO // interface{} | Days per turn. For correspondence games only. (optional)
    pairAt := TODO // interface{} | Date at which the games will be created as a Unix timestamp in milliseconds. Up to 24h in the future. Omit, or set to current date and time, to start the games immediately. Example: `1612289869919`  (optional)
    startClocksAt := TODO // interface{} | Date at which the clocks will be automatically started as a Unix timestamp in milliseconds. Up to 24h in the future. Note that the clocks can start earlier than specified, if players start making moves in the game. If omitted, the clocks will not start automatically. Example: `1612289869919`  (optional)
    rated := TODO // interface{} | Game is rated and impacts players ratings (optional) (default to false)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    fen := TODO // interface{} | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1)
    message := TODO // interface{} | Message that will be sent to each player, when the game is created.  It is sent from your user account.  `{opponent}` and `{game}` are placeholders that will be replaced with the opponent and the game URLs.  You can omit this field to send the default message, but if you set your own message, it must at least contain the `{game}` placeholder.  (optional) (default to Your game with {opponent} is ready: {game}.)
    rules := TODO // interface{} | Extra game rules separated by commas. Example: `noAbort,noRematch`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkPairingsApi.BulkPairingCreate(context.Background()).Players(players).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).PairAt(pairAt).StartClocksAt(startClocksAt).Rated(rated).Variant(variant).Fen(fen).Message(message).Rules(rules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsApi.BulkPairingCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkPairingCreate`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BulkPairingsApi.BulkPairingCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **players** | [**interface{}**](interface{}.md) | OAuth tokens of all the players to pair, with the syntax &#x60;tokenOfWhitePlayerInGame1:tokenOfBlackPlayerInGame1,tokenOfWhitePlayerInGame2:tokenOfBlackPlayerInGame2,...&#x60;.  The 2 tokens of the players of a game are separated with &#x60;:&#x60;. The first token gets the white pieces. Games are separated with &#x60;,&#x60;.  Up to 1000 tokens can be sent, for a max of 500 games.  Each token must be included at most once.  Example: &#x60;token1:token2,token3:token4,token5:token6&#x60;  | 
 **clockLimit** | [**interface{}**](interface{}.md) | Clock initial time in seconds. Example: &#x60;600&#x60;  | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds. Example: &#x60;2&#x60;  | 
 **days** | [**interface{}**](interface{}.md) | Days per turn. For correspondence games only. | 
 **pairAt** | [**interface{}**](interface{}.md) | Date at which the games will be created as a Unix timestamp in milliseconds. Up to 24h in the future. Omit, or set to current date and time, to start the games immediately. Example: &#x60;1612289869919&#x60;  | 
 **startClocksAt** | [**interface{}**](interface{}.md) | Date at which the clocks will be automatically started as a Unix timestamp in milliseconds. Up to 24h in the future. Note that the clocks can start earlier than specified, if players start making moves in the game. If omitted, the clocks will not start automatically. Example: &#x60;1612289869919&#x60;  | 
 **rated** | [**interface{}**](interface{}.md) | Game is rated and impacts players ratings | [default to false]
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1]
 **message** | [**interface{}**](interface{}.md) | Message that will be sent to each player, when the game is created.  It is sent from your user account.  &#x60;{opponent}&#x60; and &#x60;{game}&#x60; are placeholders that will be replaced with the opponent and the game URLs.  You can omit this field to send the default message, but if you set your own message, it must at least contain the &#x60;{game}&#x60; placeholder.  | [default to Your game with {opponent} is ready: {game}.]
 **rules** | [**interface{}**](interface{}.md) | Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60;  | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPairingDelete

> Ok BulkPairingDelete(ctx, id).Execute()

Cancel a bulk pairing



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkPairingsApi.BulkPairingDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsApi.BulkPairingDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkPairingDelete`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BulkPairingsApi.BulkPairingDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingDeleteRequest struct via the builder pattern


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


## BulkPairingGet

> []interface{} BulkPairingGet(ctx).Execute()

View upcoming bulk pairings



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
    resp, r, err := apiClient.BulkPairingsApi.BulkPairingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsApi.BulkPairingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkPairingGet`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `BulkPairingsApi.BulkPairingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingGetRequest struct via the builder pattern


### Return type

**[]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPairingStartClocks

> Ok BulkPairingStartClocks(ctx, id).Execute()

Manually start clocks



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkPairingsApi.BulkPairingStartClocks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsApi.BulkPairingStartClocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkPairingStartClocks`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BulkPairingsApi.BulkPairingStartClocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingStartClocksRequest struct via the builder pattern


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

