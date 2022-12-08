# \TVApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TvChannelGames**](TVApi.md#TvChannelGames) | **Get** /api/tv/{channel} | Get best ongoing games of a TV channel
[**TvChannels**](TVApi.md#TvChannels) | **Get** /api/tv/channels | Get current TV games
[**TvFeed**](TVApi.md#TvFeed) | **Get** /api/tv/feed | Stream current TV game



## TvChannelGames

> interface{} TvChannelGames(ctx, channel).Nb(nb).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Opening(opening).Execute()

Get best ongoing games of a TV channel



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
    channel := TODO // interface{} | The name of the channel in camel case.
    nb := TODO // interface{} | Number of games to fetch. (optional) (default to 10)
    moves := TODO // interface{} | Include the PGN moves. (optional) (default to true)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to false)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TVApi.TvChannelGames(context.Background(), channel).Nb(nb).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Opening(opening).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TVApi.TvChannelGames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TvChannelGames`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TVApi.TvChannelGames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | [**interface{}**](.md) | The name of the channel in camel case. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTvChannelGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nb** | [**interface{}**](interface{}.md) | Number of games to fetch. | [default to 10]
 **moves** | [**interface{}**](interface{}.md) | Include the PGN moves. | [default to true]
 **pgnInJson** | [**interface{}**](interface{}.md) | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | [**interface{}**](interface{}.md) | Include the PGN tags. | [default to true]
 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to false]
 **opening** | [**interface{}**](interface{}.md) | Include the opening name.  Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TvChannels

> interface{} TvChannels(ctx).Execute()

Get current TV games



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
    resp, r, err := apiClient.TVApi.TvChannels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TVApi.TvChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TvChannels`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TVApi.TvChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTvChannelsRequest struct via the builder pattern


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


## TvFeed

> interface{} TvFeed(ctx).Execute()

Stream current TV game



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
    resp, r, err := apiClient.TVApi.TvFeed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TVApi.TvFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TvFeed`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TVApi.TvFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTvFeedRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

