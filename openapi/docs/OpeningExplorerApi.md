# \OpeningExplorerApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpeningExplorerLichess**](OpeningExplorerApi.md#OpeningExplorerLichess) | **Get** /lichess | Lichess games
[**OpeningExplorerLichessHistory**](OpeningExplorerApi.md#OpeningExplorerLichessHistory) | **Get** /lichess/history | Lichess games history
[**OpeningExplorerMaster**](OpeningExplorerApi.md#OpeningExplorerMaster) | **Get** /masters | Masters database
[**OpeningExplorerMasterGame**](OpeningExplorerApi.md#OpeningExplorerMasterGame) | **Get** /master/pgn/{gameId} | OTB master game
[**OpeningExplorerPlayer**](OpeningExplorerApi.md#OpeningExplorerPlayer) | **Get** /player | Player games



## OpeningExplorerLichess

> interface{} OpeningExplorerLichess(ctx).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Ratings(ratings).Since(since).Until(until).Moves(moves).TopGames(topGames).RecentGames(recentGames).Execute()

Lichess games



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
    variant := *openapiclient.NewVariantKey() // VariantKey | Variant (optional) (default to standard)
    fen := TODO // interface{} | FEN of the root position (optional)
    play := TODO // interface{} | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to )
    speeds := []openapiclient.Speed{*openapiclient.NewSpeed()} // []Speed | Comma separated list of game speeds to look for (optional)
    ratings := []interface{}{interface{}(123)} // []interface{} | Comma separated list of rating groups, ranging from their value to the next higher group (optional)
    since := TODO // interface{} | Include only games from this month or later (optional) (default to 1952-01)
    until := TODO // interface{} | Include only games from this month or earlier (optional) (default to 3000-12)
    moves := TODO // interface{} | Number of most common moves to display (optional) (default to 12)
    topGames := TODO // interface{} | Number of top games to display (optional) (default to 4)
    recentGames := TODO // interface{} | Number of recent games to display (optional) (default to 4)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpeningExplorerApi.OpeningExplorerLichess(context.Background()).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Ratings(ratings).Since(since).Until(until).Moves(moves).TopGames(topGames).RecentGames(recentGames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerApi.OpeningExplorerLichess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpeningExplorerLichess`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerApi.OpeningExplorerLichess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerLichessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | [**VariantKey**](VariantKey.md) | Variant | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | FEN of the root position | 
 **play** | [**interface{}**](interface{}.md) | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to ]
 **speeds** | [**[]Speed**](Speed.md) | Comma separated list of game speeds to look for | 
 **ratings** | **[]interface{}** | Comma separated list of rating groups, ranging from their value to the next higher group | 
 **since** | [**interface{}**](interface{}.md) | Include only games from this month or later | [default to 1952-01]
 **until** | [**interface{}**](interface{}.md) | Include only games from this month or earlier | [default to 3000-12]
 **moves** | [**interface{}**](interface{}.md) | Number of most common moves to display | [default to 12]
 **topGames** | [**interface{}**](interface{}.md) | Number of top games to display | [default to 4]
 **recentGames** | [**interface{}**](interface{}.md) | Number of recent games to display | [default to 4]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpeningExplorerLichessHistory

> interface{} OpeningExplorerLichessHistory(ctx).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Ratings(ratings).Since(since).Until(until).Execute()

Lichess games history



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
    variant := *openapiclient.NewVariantKey() // VariantKey | Variant (optional) (default to standard)
    fen := TODO // interface{} | FEN of the root position (optional)
    play := TODO // interface{} | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to )
    speeds := TODO // interface{} | Comma separated list of game speeds to look for (optional)
    ratings := TODO // interface{} | Comma separated list of rating groups, ranging from their value to the next higher group (optional)
    since := TODO // interface{} | Include only games from this month or later (optional)
    until := TODO // interface{} | Include only games from this month or earlier (optional) (default to 3000-12)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpeningExplorerApi.OpeningExplorerLichessHistory(context.Background()).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Ratings(ratings).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerApi.OpeningExplorerLichessHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpeningExplorerLichessHistory`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerApi.OpeningExplorerLichessHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerLichessHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | [**VariantKey**](VariantKey.md) | Variant | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | FEN of the root position | 
 **play** | [**interface{}**](interface{}.md) | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to ]
 **speeds** | [**interface{}**](interface{}.md) | Comma separated list of game speeds to look for | 
 **ratings** | [**interface{}**](interface{}.md) | Comma separated list of rating groups, ranging from their value to the next higher group | 
 **since** | [**interface{}**](interface{}.md) | Include only games from this month or later | 
 **until** | [**interface{}**](interface{}.md) | Include only games from this month or earlier | [default to 3000-12]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpeningExplorerMaster

> interface{} OpeningExplorerMaster(ctx).Fen(fen).Play(play).Since(since).Until(until).Moves(moves).TopGames(topGames).Execute()

Masters database



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
    fen := TODO // interface{} | FEN of the root position (optional)
    play := TODO // interface{} | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to )
    since := TODO // interface{} | Include only games from this year or later (optional) (default to 1952)
    until := TODO // interface{} | Include only games from this year or earlier (optional)
    moves := TODO // interface{} | Number of most common moves to display (optional) (default to 12)
    topGames := TODO // interface{} | Number of top games to display (optional) (default to 15)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpeningExplorerApi.OpeningExplorerMaster(context.Background()).Fen(fen).Play(play).Since(since).Until(until).Moves(moves).TopGames(topGames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerApi.OpeningExplorerMaster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpeningExplorerMaster`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerApi.OpeningExplorerMaster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerMasterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | [**interface{}**](interface{}.md) | FEN of the root position | 
 **play** | [**interface{}**](interface{}.md) | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to ]
 **since** | [**interface{}**](interface{}.md) | Include only games from this year or later | [default to 1952]
 **until** | [**interface{}**](interface{}.md) | Include only games from this year or earlier | 
 **moves** | [**interface{}**](interface{}.md) | Number of most common moves to display | [default to 12]
 **topGames** | [**interface{}**](interface{}.md) | Number of top games to display | [default to 15]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpeningExplorerMasterGame

> interface{} OpeningExplorerMasterGame(ctx, gameId).Execute()

OTB master game



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
    resp, r, err := apiClient.OpeningExplorerApi.OpeningExplorerMasterGame(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerApi.OpeningExplorerMasterGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpeningExplorerMasterGame`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerApi.OpeningExplorerMasterGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerMasterGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpeningExplorerPlayer

> interface{} OpeningExplorerPlayer(ctx).Player(player).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Modes(modes).Since(since).Until(until).Moves(moves).RecentGames(recentGames).Execute()

Player games



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
    player := TODO // interface{} | Username or ID of the player (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey | Variant (optional) (default to standard)
    fen := TODO // interface{} | FEN of the root position (optional)
    play := TODO // interface{} | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to )
    speeds := TODO // interface{} | Comma separated list of game speeds to look for (optional)
    modes := []interface{}{interface{}(123)} // []interface{} | Comma separated list of modes (optional)
    since := TODO // interface{} | Include only games from this month or later (optional) (default to 1952-01)
    until := TODO // interface{} | Include only games from this month or earlier (optional) (default to 3000-12)
    moves := TODO // interface{} | Number of most common moves to display (optional)
    recentGames := TODO // interface{} | Number of recent games to display (optional) (default to 8)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpeningExplorerApi.OpeningExplorerPlayer(context.Background()).Player(player).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Modes(modes).Since(since).Until(until).Moves(moves).RecentGames(recentGames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerApi.OpeningExplorerPlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpeningExplorerPlayer`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerApi.OpeningExplorerPlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **player** | [**interface{}**](interface{}.md) | Username or ID of the player | 
 **variant** | [**VariantKey**](VariantKey.md) | Variant | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | FEN of the root position | 
 **play** | [**interface{}**](interface{}.md) | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to ]
 **speeds** | [**interface{}**](interface{}.md) | Comma separated list of game speeds to look for | 
 **modes** | **[]interface{}** | Comma separated list of modes | 
 **since** | [**interface{}**](interface{}.md) | Include only games from this month or later | [default to 1952-01]
 **until** | [**interface{}**](interface{}.md) | Include only games from this month or earlier | [default to 3000-12]
 **moves** | [**interface{}**](interface{}.md) | Number of most common moves to display | 
 **recentGames** | [**interface{}**](interface{}.md) | Number of recent games to display | [default to 8]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/nd-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

