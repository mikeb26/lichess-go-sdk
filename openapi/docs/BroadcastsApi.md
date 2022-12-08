# \BroadcastsApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastAllRoundsPgn**](BroadcastsApi.md#BroadcastAllRoundsPgn) | **Get** /api/broadcast/{broadcastTournamentId}.pgn | Export all rounds as PGN
[**BroadcastIndex**](BroadcastsApi.md#BroadcastIndex) | **Get** /api/broadcast | Get official broadcasts
[**BroadcastPush**](BroadcastsApi.md#BroadcastPush) | **Post** /broadcast/round/{broadcastRoundId}/push | Push PGN to your broadcast round
[**BroadcastRoundCreate**](BroadcastsApi.md#BroadcastRoundCreate) | **Post** /broadcast/{broadcastTournamentId}/new | Create a broadcast round
[**BroadcastRoundGet**](BroadcastsApi.md#BroadcastRoundGet) | **Get** /broadcast/{broadcastTournamentSlug}/{broadcastRoundSlug}/{broadcastRoundId} | Get your broadcast round
[**BroadcastRoundPgn**](BroadcastsApi.md#BroadcastRoundPgn) | **Get** /api/broadcast/round/{broadcastRoundId}.pgn | Export one round as PGN
[**BroadcastRoundUpdate**](BroadcastsApi.md#BroadcastRoundUpdate) | **Post** /broadcast/round/{broadcastRoundId}/edit | Update your broadcast round
[**BroadcastStreamRoundPgn**](BroadcastsApi.md#BroadcastStreamRoundPgn) | **Get** /api/stream/broadcast/round/{broadcastRoundId}.pgn | Stream an ongoing broadcast tournament as PGN
[**BroadcastTourCreate**](BroadcastsApi.md#BroadcastTourCreate) | **Post** /broadcast/new | Create a broadcast tournament
[**BroadcastTourGet**](BroadcastsApi.md#BroadcastTourGet) | **Get** /broadcast/{slug}/{broadcastTournamentId} | Get your broadcast tournament
[**BroadcastTourUpdate**](BroadcastsApi.md#BroadcastTourUpdate) | **Post** /broadcast/{broadcastTournamentId}/edit | Update your broadcast tournament



## BroadcastAllRoundsPgn

> interface{} BroadcastAllRoundsPgn(ctx, broadcastTournamentId).Execute()

Export all rounds as PGN



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
    broadcastTournamentId := TODO // interface{} | The broadcast tournament ID (8 characters).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastAllRoundsPgn(context.Background(), broadcastTournamentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastAllRoundsPgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastAllRoundsPgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastAllRoundsPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | [**interface{}**](.md) | The broadcast tournament ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastAllRoundsPgnRequest struct via the builder pattern


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


## BroadcastIndex

> []interface{} BroadcastIndex(ctx).Nb(nb).Execute()

Get official broadcasts



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
    nb := TODO // interface{} | Max number of broadcasts to fetch (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastIndex(context.Background()).Nb(nb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastIndex`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | [**interface{}**](interface{}.md) | Max number of broadcasts to fetch | [default to 20]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastPush

> []Ok BroadcastPush(ctx, broadcastRoundId).Body(body).Execute()

Push PGN to your broadcast round



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
    broadcastRoundId := TODO // interface{} | The broadcast round ID (8 characters).
    body := interface{}(987) // interface{} | The PGN. It can contain up to 64 games, separated by a double new line.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastPush(context.Background(), broadcastRoundId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastPush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastPush`: []Ok
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastPush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | [**interface{}**](.md) | The broadcast round ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | The PGN. It can contain up to 64 games, separated by a double new line. | 

### Return type

[**[]Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastRoundCreate

> interface{} BroadcastRoundCreate(ctx, broadcastTournamentId).Name(name).SyncUrl(syncUrl).StartsAt(startsAt).Execute()

Create a broadcast round



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
    broadcastTournamentId := TODO // interface{} | The broadcast tournament ID (8 characters).
    name := TODO // interface{} | Name of the broadcast round. Length must be between 3 and 80 characters.  Example: `Round 1` 
    syncUrl := TODO // interface{} | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: `https://myserver.org/myevent/round-10/games.pgn`  If the syncUrl is missing, then the broadcast needs to be fed by [pushing PGN to it](#operation/broadcastPush).  (optional)
    startsAt := TODO // interface{} | Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round.  Example: `1356998400070`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastRoundCreate(context.Background(), broadcastTournamentId).Name(name).SyncUrl(syncUrl).StartsAt(startsAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastRoundCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastRoundCreate`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastRoundCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | [**interface{}**](.md) | The broadcast tournament ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | [**interface{}**](interface{}.md) | Name of the broadcast round. Length must be between 3 and 80 characters.  Example: &#x60;Round 1&#x60;  | 
 **syncUrl** | [**interface{}**](interface{}.md) | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: &#x60;https://myserver.org/myevent/round-10/games.pgn&#x60;  If the syncUrl is missing, then the broadcast needs to be fed by [pushing PGN to it](#operation/broadcastPush).  | 
 **startsAt** | [**interface{}**](interface{}.md) | Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round.  Example: &#x60;1356998400070&#x60;  | 

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


## BroadcastRoundGet

> []interface{} BroadcastRoundGet(ctx, broadcastTournamentSlug, broadcastRoundSlug, broadcastRoundId).Execute()

Get your broadcast round



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
    broadcastTournamentSlug := TODO // interface{} | The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
    broadcastRoundSlug := TODO // interface{} | The broadcast round slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
    broadcastRoundId := TODO // interface{} | The broadcast Round ID (8 characters).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastRoundGet(context.Background(), broadcastTournamentSlug, broadcastRoundSlug, broadcastRoundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastRoundGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastRoundGet`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastRoundGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentSlug** | [**interface{}**](.md) | The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by &#x60;-&#x60;. Only the &#x60;broadcastRoundId&#x60; is actually used. | 
**broadcastRoundSlug** | [**interface{}**](.md) | The broadcast round slug. Only used for SEO, the slug can be safely replaced by &#x60;-&#x60;. Only the &#x60;broadcastRoundId&#x60; is actually used. | 
**broadcastRoundId** | [**interface{}**](.md) | The broadcast Round ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## BroadcastRoundPgn

> interface{} BroadcastRoundPgn(ctx, broadcastRoundId).Execute()

Export one round as PGN



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
    broadcastRoundId := TODO // interface{} | The round ID (8 characters).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastRoundPgn(context.Background(), broadcastRoundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastRoundPgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastRoundPgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastRoundPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | [**interface{}**](.md) | The round ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundPgnRequest struct via the builder pattern


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


## BroadcastRoundUpdate

> Ok BroadcastRoundUpdate(ctx, broadcastRoundId).Name(name).SyncUrl(syncUrl).StartsAt(startsAt).Execute()

Update your broadcast round



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
    broadcastRoundId := TODO // interface{} | The broadcast round ID (8 characters).
    name := TODO // interface{} | Name of the broadcast round. Length must be between 3 and 80 characters.  Example: `Round 10` 
    syncUrl := TODO // interface{} | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: `https://myserver.org/myevent/round-10/games.pgn`  (optional)
    startsAt := TODO // interface{} | Timestamp in milliseconds of broadcast start. Leave empty to manually start the broadcast.  Example: `1356998400070`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastRoundUpdate(context.Background(), broadcastRoundId).Name(name).SyncUrl(syncUrl).StartsAt(startsAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastRoundUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastRoundUpdate`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastRoundUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | [**interface{}**](.md) | The broadcast round ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | [**interface{}**](interface{}.md) | Name of the broadcast round. Length must be between 3 and 80 characters.  Example: &#x60;Round 10&#x60;  | 
 **syncUrl** | [**interface{}**](interface{}.md) | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: &#x60;https://myserver.org/myevent/round-10/games.pgn&#x60;  | 
 **startsAt** | [**interface{}**](interface{}.md) | Timestamp in milliseconds of broadcast start. Leave empty to manually start the broadcast.  Example: &#x60;1356998400070&#x60;  | 

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


## BroadcastStreamRoundPgn

> interface{} BroadcastStreamRoundPgn(ctx, broadcastRoundId).Execute()

Stream an ongoing broadcast tournament as PGN



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
    broadcastRoundId := TODO // interface{} | The broadcast round ID (8 characters).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastStreamRoundPgn(context.Background(), broadcastRoundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastStreamRoundPgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastStreamRoundPgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastStreamRoundPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | [**interface{}**](.md) | The broadcast round ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastStreamRoundPgnRequest struct via the builder pattern


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


## BroadcastTourCreate

> interface{} BroadcastTourCreate(ctx).Name(name).Description(description).Markdown(markdown).Official(official).Execute()

Create a broadcast tournament



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
    name := TODO // interface{} | Name of the broadcast tournament. Length must be between 3 and 80 characters.  Example: `Sinquefield Cup` 
    description := TODO // interface{} | Short description of the broadcast tournament. Length must be between 3 and 400 characters.  Example: `An 11 round classical tournament featuring the 9 highest rated players in the world. Including Carlsen, Caruana, Ding, Aronian, Nakamura and more.` 
    markdown := TODO // interface{} | Optional long description of the broadcast. Markdown is supported. Length must be less than 20,000 characters. (optional)
    official := TODO // interface{} | For Lichess internal usage only. You are not allowed to use this flag. If you do it, we will have to call the police. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastTourCreate(context.Background()).Name(name).Description(description).Markdown(markdown).Official(official).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastTourCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastTourCreate`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastTourCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTourCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | [**interface{}**](interface{}.md) | Name of the broadcast tournament. Length must be between 3 and 80 characters.  Example: &#x60;Sinquefield Cup&#x60;  | 
 **description** | [**interface{}**](interface{}.md) | Short description of the broadcast tournament. Length must be between 3 and 400 characters.  Example: &#x60;An 11 round classical tournament featuring the 9 highest rated players in the world. Including Carlsen, Caruana, Ding, Aronian, Nakamura and more.&#x60;  | 
 **markdown** | [**interface{}**](interface{}.md) | Optional long description of the broadcast. Markdown is supported. Length must be less than 20,000 characters. | 
 **official** | [**interface{}**](interface{}.md) | For Lichess internal usage only. You are not allowed to use this flag. If you do it, we will have to call the police. | 

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


## BroadcastTourGet

> []interface{} BroadcastTourGet(ctx, slug, broadcastTournamentId).Execute()

Get your broadcast tournament



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
    slug := TODO // interface{} | The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastTournamentId` is actually used.
    broadcastTournamentId := TODO // interface{} | The broadcast tournament ID (8 characters).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastTourGet(context.Background(), slug, broadcastTournamentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastTourGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastTourGet`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastTourGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | [**interface{}**](.md) | The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by &#x60;-&#x60;. Only the &#x60;broadcastTournamentId&#x60; is actually used. | 
**broadcastTournamentId** | [**interface{}**](.md) | The broadcast tournament ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTourGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## BroadcastTourUpdate

> Ok BroadcastTourUpdate(ctx, broadcastTournamentId).Name(name).Description(description).Markdown(markdown).Official(official).Execute()

Update your broadcast tournament



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
    broadcastTournamentId := TODO // interface{} | The broadcast ID (8 characters).
    name := TODO // interface{} | Name of the broadcast tournament. Length must be between 3 and 80 characters.  Example: `Sinquefield Cup` 
    description := TODO // interface{} | Short description of the broadcast tournament. Length must be between 3 and 400 characters.  Example: `An 11 round classical tournament featuring the 9 highest rated players in the world. Including Carlsen, Caruana, Ding, Aronian, Nakamura and more.` 
    markdown := TODO // interface{} | Optional long description of the broadcast tournament. Markdown is supported. Length must be less than 20,000 characters. (optional)
    official := TODO // interface{} | For Lichess internal usage only. You are not allowed to use this flag. If you do it, we will have to call the police. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BroadcastsApi.BroadcastTourUpdate(context.Background(), broadcastTournamentId).Name(name).Description(description).Markdown(markdown).Official(official).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastTourUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BroadcastTourUpdate`: Ok
    fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastTourUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | [**interface{}**](.md) | The broadcast ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTourUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | [**interface{}**](interface{}.md) | Name of the broadcast tournament. Length must be between 3 and 80 characters.  Example: &#x60;Sinquefield Cup&#x60;  | 
 **description** | [**interface{}**](interface{}.md) | Short description of the broadcast tournament. Length must be between 3 and 400 characters.  Example: &#x60;An 11 round classical tournament featuring the 9 highest rated players in the world. Including Carlsen, Caruana, Ding, Aronian, Nakamura and more.&#x60;  | 
 **markdown** | [**interface{}**](interface{}.md) | Optional long description of the broadcast tournament. Markdown is supported. Length must be less than 20,000 characters. | 
 **official** | [**interface{}**](interface{}.md) | For Lichess internal usage only. You are not allowed to use this flag. If you do it, we will have to call the police. | 

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

