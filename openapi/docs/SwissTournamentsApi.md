# \SwissTournamentsApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSwissJoin**](SwissTournamentsApi.md#ApiSwissJoin) | **Post** /api/swiss/{id}/join | Join a Swiss tournament
[**ApiSwissNew**](SwissTournamentsApi.md#ApiSwissNew) | **Post** /api/swiss/new/{teamId} | Create a new Swiss tournament
[**ApiSwissTerminate**](SwissTournamentsApi.md#ApiSwissTerminate) | **Post** /api/swiss/{id}/terminate | Terminate a Swiss tournament
[**ApiSwissUpdate**](SwissTournamentsApi.md#ApiSwissUpdate) | **Post** /api/swiss/{id}/edit | Update a Swiss tournament.
[**ApiSwissWithdraw**](SwissTournamentsApi.md#ApiSwissWithdraw) | **Post** /api/swiss/{id}/withdraw | Pause or leave a swiss tournament
[**ApiTeamSwiss**](SwissTournamentsApi.md#ApiTeamSwiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
[**GamesBySwiss**](SwissTournamentsApi.md#GamesBySwiss) | **Get** /api/swiss/{id}/games | Export games of a Swiss tournament
[**ResultsBySwiss**](SwissTournamentsApi.md#ResultsBySwiss) | **Get** /api/swiss/{id}/results | Get results of a swiss tournament
[**Swiss**](SwissTournamentsApi.md#Swiss) | **Get** /api/swiss/{id} | Get info about a Swiss tournament
[**SwissTrf**](SwissTournamentsApi.md#SwissTrf) | **Get** /swiss/{id}.trf | Export TRF of a Swiss tournament



## ApiSwissJoin

> Ok ApiSwissJoin(ctx, id).Password(password).Execute()

Join a Swiss tournament



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
    id := TODO // interface{} | The tournament ID.
    password := TODO // interface{} | The tournament password, if one is required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.ApiSwissJoin(context.Background(), id).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ApiSwissJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSwissJoin`: Ok
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ApiSwissJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | [**interface{}**](interface{}.md) | The tournament password, if one is required | 

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


## ApiSwissNew

> interface{} ApiSwissNew(ctx, teamId).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()

Create a new Swiss tournament



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
    clockLimit := TODO // interface{} | Clock initial time in seconds
    clockIncrement := TODO // interface{} | Clock increment in seconds
    nbRounds := TODO // interface{} | Maximum number of rounds to play
    name := TODO // interface{} | The tournament name. Leave empty to get a random Grandmaster name (optional)
    startsAt := TODO // interface{} | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. (optional)
    roundInterval := TODO // interface{} | How long to wait between each round, in seconds.  Set to 99999999 to manually schedule each round from the tournament UI.  If empty or -1, a sensible value is picked automatically.  (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    description := TODO // interface{} | Anything you want to tell players about the tournament (optional)
    rated := TODO // interface{} | Games are rated and impact players ratings (optional) (default to true)
    password := TODO // interface{} | Make the tournament private and restrict access with a password. (optional)
    forbiddenPairings := TODO // interface{} | Usernames of players that must not play together.  Two usernames per line, separated by a space.  (optional)
    manualPairings := TODO // interface{} | Manual pairings for the next round.  Two usernames per line, separated by a space. Present players without a valid pairing will be given a bye, which is worth 1 point. Forfeited players will get 0 points.  (optional)
    chatFor := TODO // interface{} | Who can read and write in the chat. - 0  = No-one - 10 = Only team leaders - 20 = Only team members - 30 = All Lichess players  (optional) (default to 20)
    conditionsMinRatingRating := TODO // interface{} | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
    conditionsMaxRatingRating := TODO // interface{} | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
    conditionsNbRatedGameNb := TODO // interface{} | Minimum number of rated games required to join. (optional)
    conditionsAllowList := TODO // interface{} | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.ApiSwissNew(context.Background(), teamId).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ApiSwissNew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSwissNew`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ApiSwissNew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | [**interface{}**](.md) | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clockLimit** | [**interface{}**](interface{}.md) | Clock initial time in seconds | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds | 
 **nbRounds** | [**interface{}**](interface{}.md) | Maximum number of rounds to play | 
 **name** | [**interface{}**](interface{}.md) | The tournament name. Leave empty to get a random Grandmaster name | 
 **startsAt** | [**interface{}**](interface{}.md) | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. | 
 **roundInterval** | [**interface{}**](interface{}.md) | How long to wait between each round, in seconds.  Set to 99999999 to manually schedule each round from the tournament UI.  If empty or -1, a sensible value is picked automatically.  | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **description** | [**interface{}**](interface{}.md) | Anything you want to tell players about the tournament | 
 **rated** | [**interface{}**](interface{}.md) | Games are rated and impact players ratings | [default to true]
 **password** | [**interface{}**](interface{}.md) | Make the tournament private and restrict access with a password. | 
 **forbiddenPairings** | [**interface{}**](interface{}.md) | Usernames of players that must not play together.  Two usernames per line, separated by a space.  | 
 **manualPairings** | [**interface{}**](interface{}.md) | Manual pairings for the next round.  Two usernames per line, separated by a space. Present players without a valid pairing will be given a bye, which is worth 1 point. Forfeited players will get 0 points.  | 
 **chatFor** | [**interface{}**](interface{}.md) | Who can read and write in the chat. - 0  &#x3D; No-one - 10 &#x3D; Only team leaders - 20 &#x3D; Only team members - 30 &#x3D; All Lichess players  | [default to 20]
 **conditionsMinRatingRating** | [**interface{}**](interface{}.md) | Minimum rating to join. Leave empty to let everyone join the tournament. | 
 **conditionsMaxRatingRating** | [**interface{}**](interface{}.md) | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. | 
 **conditionsNbRatedGameNb** | [**interface{}**](interface{}.md) | Minimum number of rated games required to join. | 
 **conditionsAllowList** | [**interface{}**](interface{}.md) | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60;  | 

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


## ApiSwissTerminate

> Ok ApiSwissTerminate(ctx, id).Execute()

Terminate a Swiss tournament



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
    id := TODO // interface{} | The Swiss tournament ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.ApiSwissTerminate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ApiSwissTerminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSwissTerminate`: Ok
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ApiSwissTerminate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The Swiss tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissTerminateRequest struct via the builder pattern


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


## ApiSwissUpdate

> interface{} ApiSwissUpdate(ctx, id).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()

Update a Swiss tournament.



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
    id := TODO // interface{} | The tournament ID.
    clockLimit := TODO // interface{} | Clock initial time in seconds
    clockIncrement := TODO // interface{} | Clock increment in seconds
    nbRounds := TODO // interface{} | Maximum number of rounds to play
    name := TODO // interface{} | The tournament name. Leave empty to get a random Grandmaster name (optional)
    startsAt := TODO // interface{} | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. (optional)
    roundInterval := TODO // interface{} | How long to wait between each round, in seconds.  Set to 99999999 to manually schedule each round from the tournament UI.  If empty or -1, a sensible value is picked automatically.  (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    description := TODO // interface{} | Anything you want to tell players about the tournament (optional)
    rated := TODO // interface{} | Games are rated and impact players ratings (optional) (default to true)
    password := TODO // interface{} | Make the tournament private and restrict access with a password. (optional)
    forbiddenPairings := TODO // interface{} | Usernames of players that must not play together.  Two usernames per line, separated by a space.  (optional)
    manualPairings := TODO // interface{} | Manual pairings for the next round.  Two usernames per line, separated by a space. Present players without a valid pairing will be given a bye, which is worth 1 point. Forfeited players will get 0 points.  (optional)
    chatFor := TODO // interface{} | Who can read and write in the chat. - 0  = No-one - 10 = Only team leaders - 20 = Only team members - 30 = All Lichess players  (optional) (default to 20)
    conditionsMinRatingRating := TODO // interface{} | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
    conditionsMaxRatingRating := TODO // interface{} | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
    conditionsNbRatedGameNb := TODO // interface{} | Minimum number of rated games required to join. (optional)
    conditionsAllowList := TODO // interface{} | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.ApiSwissUpdate(context.Background(), id).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ApiSwissUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSwissUpdate`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ApiSwissUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clockLimit** | [**interface{}**](interface{}.md) | Clock initial time in seconds | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds | 
 **nbRounds** | [**interface{}**](interface{}.md) | Maximum number of rounds to play | 
 **name** | [**interface{}**](interface{}.md) | The tournament name. Leave empty to get a random Grandmaster name | 
 **startsAt** | [**interface{}**](interface{}.md) | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. | 
 **roundInterval** | [**interface{}**](interface{}.md) | How long to wait between each round, in seconds.  Set to 99999999 to manually schedule each round from the tournament UI.  If empty or -1, a sensible value is picked automatically.  | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **description** | [**interface{}**](interface{}.md) | Anything you want to tell players about the tournament | 
 **rated** | [**interface{}**](interface{}.md) | Games are rated and impact players ratings | [default to true]
 **password** | [**interface{}**](interface{}.md) | Make the tournament private and restrict access with a password. | 
 **forbiddenPairings** | [**interface{}**](interface{}.md) | Usernames of players that must not play together.  Two usernames per line, separated by a space.  | 
 **manualPairings** | [**interface{}**](interface{}.md) | Manual pairings for the next round.  Two usernames per line, separated by a space. Present players without a valid pairing will be given a bye, which is worth 1 point. Forfeited players will get 0 points.  | 
 **chatFor** | [**interface{}**](interface{}.md) | Who can read and write in the chat. - 0  &#x3D; No-one - 10 &#x3D; Only team leaders - 20 &#x3D; Only team members - 30 &#x3D; All Lichess players  | [default to 20]
 **conditionsMinRatingRating** | [**interface{}**](interface{}.md) | Minimum rating to join. Leave empty to let everyone join the tournament. | 
 **conditionsMaxRatingRating** | [**interface{}**](interface{}.md) | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. | 
 **conditionsNbRatedGameNb** | [**interface{}**](interface{}.md) | Minimum number of rated games required to join. | 
 **conditionsAllowList** | [**interface{}**](interface{}.md) | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60;  | 

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


## ApiSwissWithdraw

> Ok ApiSwissWithdraw(ctx, id).Execute()

Pause or leave a swiss tournament



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
    id := TODO // interface{} | The tournament ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.ApiSwissWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ApiSwissWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSwissWithdraw`: Ok
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ApiSwissWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissWithdrawRequest struct via the builder pattern


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
    resp, r, err := apiClient.SwissTournamentsApi.ApiTeamSwiss(context.Background(), teamId).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ApiTeamSwiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamSwiss`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ApiTeamSwiss`: %v\n", resp)
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


## GamesBySwiss

> interface{} GamesBySwiss(ctx, id).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Execute()

Export games of a Swiss tournament



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
    id := TODO // interface{} | The tournament ID.
    moves := TODO // interface{} | Include the PGN moves. (optional) (default to true)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to false)
    evals := TODO // interface{} | Include analysis evaluation comments in the PGN, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`  (optional) (default to false)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.GamesBySwiss(context.Background(), id).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.GamesBySwiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamesBySwiss`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.GamesBySwiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesBySwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moves** | [**interface{}**](interface{}.md) | Include the PGN moves. | [default to true]
 **pgnInJson** | [**interface{}**](interface{}.md) | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | [**interface{}**](interface{}.md) | Include the PGN tags. | [default to true]
 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to false]
 **evals** | [**interface{}**](interface{}.md) | Include analysis evaluation comments in the PGN, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60;  | [default to false]
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


## ResultsBySwiss

> interface{} ResultsBySwiss(ctx, id).Nb(nb).Execute()

Get results of a swiss tournament



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
    id := TODO // interface{} | The tournament ID.
    nb := TODO // interface{} | Max number of players to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.ResultsBySwiss(context.Background(), id).Nb(nb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.ResultsBySwiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResultsBySwiss`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.ResultsBySwiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResultsBySwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nb** | [**interface{}**](interface{}.md) | Max number of players to fetch | 

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


## Swiss

> interface{} Swiss(ctx, id).Execute()

Get info about a Swiss tournament



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
    id := TODO // interface{} | The Swiss tournament ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.Swiss(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.Swiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Swiss`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.Swiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The Swiss tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwissRequest struct via the builder pattern


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


## SwissTrf

> interface{} SwissTrf(ctx, id).Execute()

Export TRF of a Swiss tournament



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
    id := TODO // interface{} | The tournament ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwissTournamentsApi.SwissTrf(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwissTournamentsApi.SwissTrf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwissTrf`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SwissTournamentsApi.SwissTrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwissTrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

