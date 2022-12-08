# \PuzzlesApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPuzzleActivity**](PuzzlesApi.md#ApiPuzzleActivity) | **Get** /api/puzzle/activity | Get your puzzle activity
[**ApiPuzzleDaily**](PuzzlesApi.md#ApiPuzzleDaily) | **Get** /api/puzzle/daily | Get the daily puzzle
[**ApiPuzzleDashboard**](PuzzlesApi.md#ApiPuzzleDashboard) | **Get** /api/puzzle/dashboard/{days} | Get your puzzle dashboard
[**ApiPuzzleId**](PuzzlesApi.md#ApiPuzzleId) | **Get** /api/puzzle/{id} | Get a puzzle by its ID
[**ApiStormDashboard**](PuzzlesApi.md#ApiStormDashboard) | **Get** /api/storm/dashboard/{username} | Get the storm dashboard of a player
[**RacerPost**](PuzzlesApi.md#RacerPost) | **Post** /api/racer | Create and join a puzzle race



## ApiPuzzleActivity

> PuzzleRoundJson ApiPuzzleActivity(ctx).Max(max).Execute()

Get your puzzle activity



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
    max := TODO // interface{} | How many entries to download. Leave empty to download all activity. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PuzzlesApi.ApiPuzzleActivity(context.Background()).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesApi.ApiPuzzleActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPuzzleActivity`: PuzzleRoundJson
    fmt.Fprintf(os.Stdout, "Response from `PuzzlesApi.ApiPuzzleActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | [**interface{}**](interface{}.md) | How many entries to download. Leave empty to download all activity. | 

### Return type

[**PuzzleRoundJson**](PuzzleRoundJson.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleDaily

> interface{} ApiPuzzleDaily(ctx).Execute()

Get the daily puzzle



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
    resp, r, err := apiClient.PuzzlesApi.ApiPuzzleDaily(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesApi.ApiPuzzleDaily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPuzzleDaily`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PuzzlesApi.ApiPuzzleDaily`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleDailyRequest struct via the builder pattern


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


## ApiPuzzleDashboard

> interface{} ApiPuzzleDashboard(ctx, days).Execute()

Get your puzzle dashboard



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
    days := TODO // interface{} | How many days to look back when aggregating puzzle results. 30 is sensible.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PuzzlesApi.ApiPuzzleDashboard(context.Background(), days).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesApi.ApiPuzzleDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPuzzleDashboard`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PuzzlesApi.ApiPuzzleDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**days** | [**interface{}**](.md) | How many days to look back when aggregating puzzle results. 30 is sensible. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleId

> interface{} ApiPuzzleId(ctx, id).Execute()

Get a puzzle by its ID



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
    id := TODO // interface{} | The puzzle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PuzzlesApi.ApiPuzzleId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesApi.ApiPuzzleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPuzzleId`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PuzzlesApi.ApiPuzzleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The puzzle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiStormDashboard

> interface{} ApiStormDashboard(ctx, username).Days(days).Execute()

Get the storm dashboard of a player



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
    username := TODO // interface{} | Username of the player
    days := TODO // interface{} | How many days of history to return (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PuzzlesApi.ApiStormDashboard(context.Background(), username).Days(days).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesApi.ApiStormDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStormDashboard`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PuzzlesApi.ApiStormDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) | Username of the player | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStormDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | [**interface{}**](interface{}.md) | How many days of history to return | [default to 30]

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


## RacerPost

> interface{} RacerPost(ctx).Execute()

Create and join a puzzle race



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
    resp, r, err := apiClient.PuzzlesApi.RacerPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesApi.RacerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacerPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PuzzlesApi.RacerPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRacerPostRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

