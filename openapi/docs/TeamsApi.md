# \TeamsApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamArena**](TeamsApi.md#ApiTeamArena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
[**ApiTeamSwiss**](TeamsApi.md#ApiTeamSwiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
[**TeamAll**](TeamsApi.md#TeamAll) | **Get** /api/team/all | Get popular teams
[**TeamIdJoin**](TeamsApi.md#TeamIdJoin) | **Post** /team/{teamId}/join | Join a team
[**TeamIdKickUserId**](TeamsApi.md#TeamIdKickUserId) | **Post** /team/{teamId}/kick/{userId} | Kick a user from your team
[**TeamIdPmAll**](TeamsApi.md#TeamIdPmAll) | **Post** /team/{teamId}/pm-all | Message all members
[**TeamIdQuit**](TeamsApi.md#TeamIdQuit) | **Post** /team/{teamId}/quit | Leave a team
[**TeamIdUsers**](TeamsApi.md#TeamIdUsers) | **Get** /api/team/{teamId}/users | Get members of a team
[**TeamOfUsername**](TeamsApi.md#TeamOfUsername) | **Get** /api/team/of/{username} | Teams of a player
[**TeamRequestAccept**](TeamsApi.md#TeamRequestAccept) | **Post** /api/team/{teamId}/request/{userId}/accept | Accept join request
[**TeamRequestDecline**](TeamsApi.md#TeamRequestDecline) | **Post** /api/team/{teamId}/request/{userId}/decline | Decline join request
[**TeamRequests**](TeamsApi.md#TeamRequests) | **Get** /api/team/{teamId}/requests | Get join requests
[**TeamSearch**](TeamsApi.md#TeamSearch) | **Get** /api/team/search | Search teams
[**TeamShow**](TeamsApi.md#TeamShow) | **Get** /api/team/{teamId} | Get a single team



## ApiTeamArena

> interface{} ApiTeamArena(ctx, teamId).Max(max).Execute()

Get team Arena tournaments



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
    teamId := TODO // interface{} | ID of the team
    max := TODO // interface{} | How many tournaments to download. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ApiTeamArena(context.Background(), teamId).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ApiTeamArena``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamArena`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ApiTeamArena`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamArenaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | [**interface{}**](interface{}.md) | How many tournaments to download. | [default to 100]

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


## ApiTeamSwiss

> interface{} ApiTeamSwiss(ctx, teamId).Max(max).Execute()

Get team swiss tournaments



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
    max := TODO // interface{} | How many tournaments to download. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ApiTeamSwiss(context.Background(), teamId).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ApiTeamSwiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamSwiss`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ApiTeamSwiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamSwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | [**interface{}**](interface{}.md) | How many tournaments to download. | [default to 100]

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


## TeamAll

> TeamAll200Response TeamAll(ctx).Page(page).Execute()

Get popular teams



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
    page := TODO // interface{} |  (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamAll(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamAll`: TeamAll200Response
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**interface{}**](interface{}.md) |  | [default to 1]

### Return type

[**TeamAll200Response**](TeamAll200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamIdJoin

> Ok TeamIdJoin(ctx, teamId).Message(message).Password(password).Execute()

Join a team



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
    message := TODO // interface{} | Optional request message, if the team requires one. (optional)
    password := TODO // interface{} | Optional password, if the team requires one. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamIdJoin(context.Background(), teamId).Message(message).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamIdJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamIdJoin`: Ok
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamIdJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | [**interface{}**](interface{}.md) | Optional request message, if the team requires one. | 
 **password** | [**interface{}**](interface{}.md) | Optional password, if the team requires one. | 

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


## TeamIdKickUserId

> Ok TeamIdKickUserId(ctx, teamId, userId).Execute()

Kick a user from your team



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
    userId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamIdKickUserId(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamIdKickUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamIdKickUserId`: Ok
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamIdKickUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 
**userId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdKickUserIdRequest struct via the builder pattern


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


## TeamIdPmAll

> Ok TeamIdPmAll(ctx, teamId).Message(message).Execute()

Message all members



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
    message := TODO // interface{} | The message to send to all your team members. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamIdPmAll(context.Background(), teamId).Message(message).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamIdPmAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamIdPmAll`: Ok
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamIdPmAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdPmAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | [**interface{}**](interface{}.md) | The message to send to all your team members. | 

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


## TeamIdQuit

> Ok TeamIdQuit(ctx, teamId).Execute()

Leave a team



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
    resp, r, err := apiClient.TeamsApi.TeamIdQuit(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamIdQuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamIdQuit`: Ok
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamIdQuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdQuitRequest struct via the builder pattern


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
    resp, r, err := apiClient.TeamsApi.TeamIdUsers(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamIdUsers`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamIdUsers`: %v\n", resp)
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


## TeamOfUsername

> interface{} TeamOfUsername(ctx, username).Execute()

Teams of a player



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
    resp, r, err := apiClient.TeamsApi.TeamOfUsername(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamOfUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamOfUsername`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamOfUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamOfUsernameRequest struct via the builder pattern


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


## TeamRequestAccept

> Ok TeamRequestAccept(ctx, teamId, userId).Execute()

Accept join request



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
    userId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamRequestAccept(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamRequestAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamRequestAccept`: Ok
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamRequestAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 
**userId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamRequestAcceptRequest struct via the builder pattern


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


## TeamRequestDecline

> Ok TeamRequestDecline(ctx, teamId, userId).Execute()

Decline join request



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
    userId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamRequestDecline(context.Background(), teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamRequestDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamRequestDecline`: Ok
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamRequestDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 
**userId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamRequestDeclineRequest struct via the builder pattern


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


## TeamRequests

> []TeamRequestWithUser TeamRequests(ctx, teamId).Execute()

Get join requests



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
    resp, r, err := apiClient.TeamsApi.TeamRequests(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamRequests`: []TeamRequestWithUser
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamRequestWithUser**](TeamRequestWithUser.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamSearch

> interface{} TeamSearch(ctx).Text(text).Page(page).Execute()

Search teams



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
    text := TODO // interface{} |  (optional)
    page := TODO // interface{} |  (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamSearch(context.Background()).Text(text).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamSearch`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | [**interface{}**](interface{}.md) |  | 
 **page** | [**interface{}**](interface{}.md) |  | [default to 1]

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


## TeamShow

> Team TeamShow(ctx, teamId).Execute()

Get a single team



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
    resp, r, err := apiClient.TeamsApi.TeamShow(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamShow`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

