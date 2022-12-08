# \UsersApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCrosstable**](UsersApi.md#ApiCrosstable) | **Get** /api/crosstable/{user1}/{user2} | Get crosstable
[**ApiUser**](UsersApi.md#ApiUser) | **Get** /api/user/{username} | Get user public data
[**ApiUserActivity**](UsersApi.md#ApiUserActivity) | **Get** /api/user/{username}/activity | Get user activity
[**ApiUserPerf**](UsersApi.md#ApiUserPerf) | **Get** /api/user/{username}/perf/{perf} | Get performance statistics of a user
[**ApiUserRatingHistory**](UsersApi.md#ApiUserRatingHistory) | **Get** /api/user/{username}/rating-history | Get rating history of a user
[**ApiUsers**](UsersApi.md#ApiUsers) | **Post** /api/users | Get users by ID
[**ApiUsersStatus**](UsersApi.md#ApiUsersStatus) | **Get** /api/users/status | Get real-time users status
[**Player**](UsersApi.md#Player) | **Get** /api/player | Get all top 10
[**PlayerTopNbPerfType**](UsersApi.md#PlayerTopNbPerfType) | **Get** /api/player/top/{nb}/{perfType} | Get one leaderboard
[**StreamerLive**](UsersApi.md#StreamerLive) | **Get** /api/streamer/live | Get live streamers
[**TeamIdUsers**](UsersApi.md#TeamIdUsers) | **Get** /api/team/{teamId}/users | Get members of a team



## ApiCrosstable

> interface{} ApiCrosstable(ctx, user1, user2).Matchup(matchup).Execute()

Get crosstable



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
    user1 := TODO // interface{} | 
    user2 := TODO // interface{} | 
    matchup := TODO // interface{} | Whether to get the current match data, if any (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiCrosstable(context.Background(), user1, user2).Matchup(matchup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiCrosstable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiCrosstable`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiCrosstable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user1** | [**interface{}**](.md) |  | 
**user2** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCrosstableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **matchup** | [**interface{}**](interface{}.md) | Whether to get the current match data, if any | 

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


## ApiUser

> interface{} ApiUser(ctx, username).Trophies(trophies).Execute()

Get user public data



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
    username := TODO // interface{} | 
    trophies := TODO // interface{} | Include user trophies (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiUser(context.Background(), username).Trophies(trophies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUser`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trophies** | [**interface{}**](interface{}.md) | Include user trophies | [default to false]

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


## ApiUserActivity

> ApiUserActivity(ctx, username).Execute()

Get user activity



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
    username := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiUserActivity(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUserActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserPerf

> interface{} ApiUserPerf(ctx, username, perf).Execute()

Get performance statistics of a user



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
    username := TODO // interface{} | 
    perf := *openapiclient.NewPerfType() // PerfType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiUserPerf(context.Background(), username, perf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUserPerf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserPerf`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUserPerf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 
**perf** | [**PerfType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserPerfRequest struct via the builder pattern


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


## ApiUserRatingHistory

> interface{} ApiUserRatingHistory(ctx, username).Execute()

Get rating history of a user



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
    username := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiUserRatingHistory(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUserRatingHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserRatingHistory`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUserRatingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserRatingHistoryRequest struct via the builder pattern


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


## ApiUsers

> []User ApiUsers(ctx).Body(body).Execute()

Get users by ID



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
    body := interface{}(aliquantus,chess-network,lovlas) // interface{} | User IDs separated by commas.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiUsers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | User IDs separated by commas. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersStatus

> []ApiUsersStatus200ResponseInner ApiUsersStatus(ctx).Ids(ids).WithGameIds(withGameIds).Execute()

Get real-time users status



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
    ids := TODO // interface{} | User IDs separated by commas. Up to 100 IDs.
    withGameIds := TODO // interface{} | Also return the ID of the game being played, if any, for each player, in a `playingId` field. Defaults to `false` to preserve server resources.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ApiUsersStatus(context.Background()).Ids(ids).WithGameIds(withGameIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUsersStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersStatus`: []ApiUsersStatus200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUsersStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**interface{}**](interface{}.md) | User IDs separated by commas. Up to 100 IDs. | 
 **withGameIds** | [**interface{}**](interface{}.md) | Also return the ID of the game being played, if any, for each player, in a &#x60;playingId&#x60; field. Defaults to &#x60;false&#x60; to preserve server resources.  | 

### Return type

[**[]ApiUsersStatus200ResponseInner**](ApiUsersStatus200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Player

> interface{} Player(ctx).Execute()

Get all top 10



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
    resp, r, err := apiClient.UsersApi.Player(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.Player``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Player`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.Player`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerRequest struct via the builder pattern


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


## PlayerTopNbPerfType

> interface{} PlayerTopNbPerfType(ctx, nb, perfType).Execute()

Get one leaderboard



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
    nb := TODO // interface{} | How many users to fetch
    perfType := TODO // interface{} | The speed or variant

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.PlayerTopNbPerfType(context.Background(), nb, perfType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.PlayerTopNbPerfType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerTopNbPerfType`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.PlayerTopNbPerfType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nb** | [**interface{}**](.md) | How many users to fetch | 
**perfType** | [**interface{}**](.md) | The speed or variant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerTopNbPerfTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.lichess.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamerLive

> []LightUser StreamerLive(ctx).Execute()

Get live streamers



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
    resp, r, err := apiClient.UsersApi.StreamerLive(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.StreamerLive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StreamerLive`: []LightUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.StreamerLive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamerLiveRequest struct via the builder pattern


### Return type

[**[]LightUser**](LightUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamIdUsers

> interface{} TeamIdUsers(ctx, teamId).Execute()

Get members of a team



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
    teamId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.TeamIdUsers(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.TeamIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamIdUsers`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.TeamIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdUsersRequest struct via the builder pattern


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

