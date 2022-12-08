# \ArenaTournamentsApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamArena**](ArenaTournamentsApi.md#ApiTeamArena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
[**ApiTournament**](ArenaTournamentsApi.md#ApiTournament) | **Get** /api/tournament | Get current tournaments
[**ApiTournamentJoin**](ArenaTournamentsApi.md#ApiTournamentJoin) | **Post** /api/tournament/{id}/join | Join an Arena tournament
[**ApiTournamentPost**](ArenaTournamentsApi.md#ApiTournamentPost) | **Post** /api/tournament | Create a new Arena tournament
[**ApiTournamentTeamBattlePost**](ArenaTournamentsApi.md#ApiTournamentTeamBattlePost) | **Post** /api/tournament/team-battle/{id} | Update a team battle
[**ApiTournamentTerminate**](ArenaTournamentsApi.md#ApiTournamentTerminate) | **Post** /api/tournament/{id}/terminate | Terminate an Arena tournament
[**ApiTournamentUpdate**](ArenaTournamentsApi.md#ApiTournamentUpdate) | **Post** /api/tournament/{id} | Update an Arena tournament
[**ApiTournamentWithdraw**](ArenaTournamentsApi.md#ApiTournamentWithdraw) | **Post** /api/tournament/{id}/withdraw | Pause or leave an Arena tournament
[**ApiUserNameTournamentCreated**](ArenaTournamentsApi.md#ApiUserNameTournamentCreated) | **Get** /api/user/{username}/tournament/created | Get tournaments created by a user
[**GamesByTournament**](ArenaTournamentsApi.md#GamesByTournament) | **Get** /api/tournament/{id}/games | Export games of an Arena tournament
[**ResultsByTournament**](ArenaTournamentsApi.md#ResultsByTournament) | **Get** /api/tournament/{id}/results | Get results of an Arena tournament
[**TeamsByTournament**](ArenaTournamentsApi.md#TeamsByTournament) | **Get** /api/tournament/{id}/teams | Get team standing of a team battle
[**Tournament**](ArenaTournamentsApi.md#Tournament) | **Get** /api/tournament/{id} | Get info about an Arena tournament



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
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTeamArena(context.Background(), teamId).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTeamArena``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTeamArena`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTeamArena`: %v\n", resp)
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


## ApiTournament

> ArenaTournaments ApiTournament(ctx).Execute()

Get current tournaments



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
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournament(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournament`: ArenaTournaments
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournament`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentRequest struct via the builder pattern


### Return type

[**ArenaTournaments**](ArenaTournaments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTournamentJoin

> Ok ApiTournamentJoin(ctx, id).Password(password).Team(team).PairMeAsap(pairMeAsap).Execute()

Join an Arena tournament



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
    password := TODO // interface{} | The tournament password, if one is required. Can also be a [user-specific entry code](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) generated and shared by the organizer.  (optional)
    team := TODO // interface{} | The team to join the tournament with, for team battle tournaments (optional)
    pairMeAsap := TODO // interface{} | If the tournament is started, attempt to pair the user, even if they are not connected to the tournament page. This expires after one minute, to avoid pairing a user who is long gone. You may call \\\"join\\\" again to extend the waiting.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournamentJoin(context.Background(), id).Password(password).Team(team).PairMeAsap(pairMeAsap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournamentJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournamentJoin`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournamentJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | [**interface{}**](interface{}.md) | The tournament password, if one is required. Can also be a [user-specific entry code](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) generated and shared by the organizer.  | 
 **team** | [**interface{}**](interface{}.md) | The team to join the tournament with, for team battle tournaments | 
 **pairMeAsap** | [**interface{}**](interface{}.md) | If the tournament is started, attempt to pair the user, even if they are not connected to the tournament page. This expires after one minute, to avoid pairing a user who is long gone. You may call \\\&quot;join\\\&quot; again to extend the waiting.  | [default to false]

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


## ApiTournamentPost

> interface{} ApiTournamentPost(ctx).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).TeamBattleByTeam(teamBattleByTeam).ConditionsTeamMemberTeamId(conditionsTeamMemberTeamId).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()

Create a new Arena tournament



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
    clockTime := TODO // interface{} | Clock initial time in minutes
    clockIncrement := TODO // interface{} | Clock increment in seconds
    minutes := TODO // interface{} | How long the tournament lasts, in minutes
    name := TODO // interface{} | The tournament name. Leave empty to get a random Grandmaster name (optional)
    waitMinutes := TODO // interface{} | How long to wait before starting the tournament, from now, in minutes (optional) (default to 5)
    startDate := TODO // interface{} | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the `waitMinutes` setting (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    rated := TODO // interface{} | Games are rated and impact players ratings (optional) (default to true)
    position := TODO // interface{} | Custom initial position (in FEN) for all games of the tournament. Must be a legal chess position. Only works with standard chess, not variants (except Chess960). (optional)
    berserkable := TODO // interface{} | Whether the players can use berserk. Only allowed if clockIncrement <= clockTime * 2 (optional) (default to true)
    streakable := TODO // interface{} | After 2 wins, consecutive wins grant 4 points instead of 2. (optional) (default to true)
    hasChat := TODO // interface{} | Whether the players can discuss in a chat (optional) (default to true)
    description := TODO // interface{} | Anything you want to tell players about the tournament (optional)
    password := TODO // interface{} | Make the tournament private, and restrict access with a password. You can also [generate user-specific entry codes](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) based on this password.  (optional)
    teamBattleByTeam := TODO // interface{} | Set the ID of a team you lead to create a team battle. The other teams can be added using the [team battle edit endpoint](#operation/apiTournamentTeamBattlePost).  (optional)
    conditionsTeamMemberTeamId := TODO // interface{} | Restrict entry to members of a team.  The teamId is the last part of a team URL, e.g. `https://lichess.org/team/coders` has teamId = `coders`.  Leave empty to let everyone join the tournament.  Do not use this to create team battles, use `teamBattleByTeam` instead.  (optional)
    conditionsMinRatingRating := TODO // interface{} | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
    conditionsMaxRatingRating := TODO // interface{} | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
    conditionsNbRatedGameNb := TODO // interface{} | Minimum number of rated games required to join. (optional)
    conditionsAllowList := TODO // interface{} | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournamentPost(context.Background()).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).TeamBattleByTeam(teamBattleByTeam).ConditionsTeamMemberTeamId(conditionsTeamMemberTeamId).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournamentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournamentPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournamentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clockTime** | [**interface{}**](interface{}.md) | Clock initial time in minutes | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds | 
 **minutes** | [**interface{}**](interface{}.md) | How long the tournament lasts, in minutes | 
 **name** | [**interface{}**](interface{}.md) | The tournament name. Leave empty to get a random Grandmaster name | 
 **waitMinutes** | [**interface{}**](interface{}.md) | How long to wait before starting the tournament, from now, in minutes | [default to 5]
 **startDate** | [**interface{}**](interface{}.md) | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the &#x60;waitMinutes&#x60; setting | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **rated** | [**interface{}**](interface{}.md) | Games are rated and impact players ratings | [default to true]
 **position** | [**interface{}**](interface{}.md) | Custom initial position (in FEN) for all games of the tournament. Must be a legal chess position. Only works with standard chess, not variants (except Chess960). | 
 **berserkable** | [**interface{}**](interface{}.md) | Whether the players can use berserk. Only allowed if clockIncrement &lt;&#x3D; clockTime * 2 | [default to true]
 **streakable** | [**interface{}**](interface{}.md) | After 2 wins, consecutive wins grant 4 points instead of 2. | [default to true]
 **hasChat** | [**interface{}**](interface{}.md) | Whether the players can discuss in a chat | [default to true]
 **description** | [**interface{}**](interface{}.md) | Anything you want to tell players about the tournament | 
 **password** | [**interface{}**](interface{}.md) | Make the tournament private, and restrict access with a password. You can also [generate user-specific entry codes](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) based on this password.  | 
 **teamBattleByTeam** | [**interface{}**](interface{}.md) | Set the ID of a team you lead to create a team battle. The other teams can be added using the [team battle edit endpoint](#operation/apiTournamentTeamBattlePost).  | 
 **conditionsTeamMemberTeamId** | [**interface{}**](interface{}.md) | Restrict entry to members of a team.  The teamId is the last part of a team URL, e.g. &#x60;https://lichess.org/team/coders&#x60; has teamId &#x3D; &#x60;coders&#x60;.  Leave empty to let everyone join the tournament.  Do not use this to create team battles, use &#x60;teamBattleByTeam&#x60; instead.  | 
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


## ApiTournamentTeamBattlePost

> interface{} ApiTournamentTeamBattlePost(ctx, id).Teams(teams).NbLeaders(nbLeaders).Execute()

Update a team battle



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
    id := TODO // interface{} | The tournament ID (8 characters)..
    teams := TODO // interface{} | All team IDs of the team battle, separated by commas. Make sure to always send the full list. Teams that are not in the list will be removed from the team battle.  Example: `coders,zhigalko_sergei-fan-club,hhSwTKZv` 
    nbLeaders := TODO // interface{} | Number team leaders per team.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournamentTeamBattlePost(context.Background(), id).Teams(teams).NbLeaders(nbLeaders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournamentTeamBattlePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournamentTeamBattlePost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournamentTeamBattlePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID (8 characters).. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentTeamBattlePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teams** | [**interface{}**](interface{}.md) | All team IDs of the team battle, separated by commas. Make sure to always send the full list. Teams that are not in the list will be removed from the team battle.  Example: &#x60;coders,zhigalko_sergei-fan-club,hhSwTKZv&#x60;  | 
 **nbLeaders** | [**interface{}**](interface{}.md) | Number team leaders per team. | 

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


## ApiTournamentTerminate

> Ok ApiTournamentTerminate(ctx, id).Execute()

Terminate an Arena tournament



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
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournamentTerminate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournamentTerminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournamentTerminate`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournamentTerminate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentTerminateRequest struct via the builder pattern


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


## ApiTournamentUpdate

> interface{} ApiTournamentUpdate(ctx, id).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()

Update an Arena tournament



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
    clockTime := TODO // interface{} | Clock initial time in minutes
    clockIncrement := TODO // interface{} | Clock increment in seconds
    minutes := TODO // interface{} | How long the tournament lasts, in minutes
    name := TODO // interface{} | The tournament name. Leave empty to get a random Grandmaster name (optional)
    waitMinutes := TODO // interface{} | How long to wait before starting the tournament, from now, in minutes (optional) (default to 5)
    startDate := TODO // interface{} | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the `waitMinutes` setting (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    rated := TODO // interface{} | Games are rated and impact players ratings (optional) (default to true)
    position := TODO // interface{} | Custom initial position (in FEN) for all games of the tournament. Must be a legal chess position. Only works with standard chess, not variants (except Chess960). (optional)
    berserkable := TODO // interface{} | Whether the players can use berserk. Only allowed if clockIncrement <= clockTime * 2 (optional) (default to true)
    streakable := TODO // interface{} | After 2 wins, consecutive wins grant 4 points instead of 2. (optional) (default to true)
    hasChat := TODO // interface{} | Whether the players can discuss in a chat (optional) (default to true)
    description := TODO // interface{} | Anything you want to tell players about the tournament (optional)
    password := TODO // interface{} | Make the tournament private, and restrict access with a password (optional)
    conditionsMinRatingRating := TODO // interface{} | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
    conditionsMaxRatingRating := TODO // interface{} | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
    conditionsNbRatedGameNb := TODO // interface{} | Minimum number of rated games required to join. (optional)
    conditionsAllowList := TODO // interface{} | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournamentUpdate(context.Background(), id).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournamentUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournamentUpdate`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournamentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clockTime** | [**interface{}**](interface{}.md) | Clock initial time in minutes | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds | 
 **minutes** | [**interface{}**](interface{}.md) | How long the tournament lasts, in minutes | 
 **name** | [**interface{}**](interface{}.md) | The tournament name. Leave empty to get a random Grandmaster name | 
 **waitMinutes** | [**interface{}**](interface{}.md) | How long to wait before starting the tournament, from now, in minutes | [default to 5]
 **startDate** | [**interface{}**](interface{}.md) | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the &#x60;waitMinutes&#x60; setting | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **rated** | [**interface{}**](interface{}.md) | Games are rated and impact players ratings | [default to true]
 **position** | [**interface{}**](interface{}.md) | Custom initial position (in FEN) for all games of the tournament. Must be a legal chess position. Only works with standard chess, not variants (except Chess960). | 
 **berserkable** | [**interface{}**](interface{}.md) | Whether the players can use berserk. Only allowed if clockIncrement &lt;&#x3D; clockTime * 2 | [default to true]
 **streakable** | [**interface{}**](interface{}.md) | After 2 wins, consecutive wins grant 4 points instead of 2. | [default to true]
 **hasChat** | [**interface{}**](interface{}.md) | Whether the players can discuss in a chat | [default to true]
 **description** | [**interface{}**](interface{}.md) | Anything you want to tell players about the tournament | 
 **password** | [**interface{}**](interface{}.md) | Make the tournament private, and restrict access with a password | 
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


## ApiTournamentWithdraw

> Ok ApiTournamentWithdraw(ctx, id).Execute()

Pause or leave an Arena tournament



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
    resp, r, err := apiClient.ArenaTournamentsApi.ApiTournamentWithdraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiTournamentWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTournamentWithdraw`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiTournamentWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentWithdrawRequest struct via the builder pattern


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


## ApiUserNameTournamentCreated

> interface{} ApiUserNameTournamentCreated(ctx, username).Status(status).Execute()

Get tournaments created by a user



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
    username := TODO // interface{} | The user whose created tournaments to fetch
    status := TODO // interface{} | Include tournaments in the given status: \"Created\" (10), \"Started\" (20), \"Finished\" (30)  You can add this parameter more than once to include tournaments in different statuses.   Example: `?status=10&status=20`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.ApiUserNameTournamentCreated(context.Background(), username).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ApiUserNameTournamentCreated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserNameTournamentCreated`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ApiUserNameTournamentCreated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) | The user whose created tournaments to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserNameTournamentCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**interface{}**](interface{}.md) | Include tournaments in the given status: \&quot;Created\&quot; (10), \&quot;Started\&quot; (20), \&quot;Finished\&quot; (30)  You can add this parameter more than once to include tournaments in different statuses.   Example: &#x60;?status&#x3D;10&amp;status&#x3D;20&#x60;  | 

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


## GamesByTournament

> interface{} GamesByTournament(ctx, id).Player(player).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Execute()

Export games of an Arena tournament



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
    player := TODO // interface{} | Only games of a particular player. Leave empty to fetch games of all players. (optional)
    moves := TODO // interface{} | Include the PGN moves. (optional) (default to true)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to false)
    evals := TODO // interface{} | Include analysis evaluation comments in the PGN, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`  (optional) (default to false)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.GamesByTournament(context.Background(), id).Player(player).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.GamesByTournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamesByTournament`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.GamesByTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesByTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **player** | [**interface{}**](interface{}.md) | Only games of a particular player. Leave empty to fetch games of all players. | 
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


## ResultsByTournament

> interface{} ResultsByTournament(ctx, id).Nb(nb).Execute()

Get results of an Arena tournament



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
    resp, r, err := apiClient.ArenaTournamentsApi.ResultsByTournament(context.Background(), id).Nb(nb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.ResultsByTournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResultsByTournament`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.ResultsByTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResultsByTournamentRequest struct via the builder pattern


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


## TeamsByTournament

> interface{} TeamsByTournament(ctx, id).Execute()

Get team standing of a team battle



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
    resp, r, err := apiClient.ArenaTournamentsApi.TeamsByTournament(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.TeamsByTournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsByTournament`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.TeamsByTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsByTournamentRequest struct via the builder pattern


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


## Tournament

> interface{} Tournament(ctx, id).Page(page).Execute()

Get info about an Arena tournament



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
    page := TODO // interface{} | Specify which page of player standings to view. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArenaTournamentsApi.Tournament(context.Background(), id).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArenaTournamentsApi.Tournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Tournament`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArenaTournamentsApi.Tournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**interface{}**](interface{}.md) | Specify which page of player standings to view. | [default to 1]

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

