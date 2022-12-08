# \GamesApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountPlaying**](GamesApi.md#ApiAccountPlaying) | **Get** /api/account/playing | Get my ongoing games
[**ApiGamesUser**](GamesApi.md#ApiGamesUser) | **Get** /api/games/user/{username} | Export games of a user
[**ApiUserCurrentGame**](GamesApi.md#ApiUserCurrentGame) | **Get** /api/user/{username}/current-game | Export ongoing game of a user
[**GameImport**](GamesApi.md#GameImport) | **Post** /api/import | Import one game
[**GamePgn**](GamesApi.md#GamePgn) | **Get** /game/export/{gameId} | Export one game
[**GamesByIds**](GamesApi.md#GamesByIds) | **Post** /api/stream/games/{streamId} | Stream games by IDs
[**GamesByIdsAdd**](GamesApi.md#GamesByIdsAdd) | **Post** /api/stream/games/{streamId}/add | Add game IDs to stream
[**GamesByUsers**](GamesApi.md#GamesByUsers) | **Post** /api/stream/games-by-users | Stream games of users
[**GamesExportIds**](GamesApi.md#GamesExportIds) | **Post** /api/games/export/_ids | Export games by IDs
[**StreamGame**](GamesApi.md#StreamGame) | **Get** /api/stream/game/{id} | Stream moves of a game



## ApiAccountPlaying

> interface{} ApiAccountPlaying(ctx).Nb(nb).Execute()

Get my ongoing games



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
    nb := TODO // interface{} | Max number of games to fetch (optional) (default to 9)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.ApiAccountPlaying(context.Background()).Nb(nb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.ApiAccountPlaying``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountPlaying`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.ApiAccountPlaying`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountPlayingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | [**interface{}**](interface{}.md) | Max number of games to fetch | [default to 9]

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


## ApiGamesUser

> interface{} ApiGamesUser(ctx, username).Since(since).Until(until).Max(max).Vs(vs).Rated(rated).PerfType(perfType).Color(color).Analysed(analysed).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Ongoing(ongoing).Finished(finished).Literate(literate).Players(players).Sort(sort).Execute()

Export games of a user



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
    username := TODO // interface{} | The user name.
    since := TODO // interface{} | Download games played since this timestamp. Defaults to account creation date. (optional)
    until := TODO // interface{} | Download games played until this timestamp. Defaults to now. (optional)
    max := TODO // interface{} | How many games to download. Leave empty to download all games. (optional)
    vs := TODO // interface{} | [Filter] Only games played against this opponent (optional)
    rated := TODO // interface{} | [Filter] Only rated (`true`) or casual (`false`) games (optional)
    perfType := TODO // interface{} | [Filter] Only games in these speeds or variants.   Multiple perf types can be specified, separated by a comma.   Example: blitz,rapid,classical (optional)
    color := TODO // interface{} | [Filter] Only games played as this color. (optional)
    analysed := TODO // interface{} | [Filter] Only games with or without a computer analysis available (optional)
    moves := TODO // interface{} | Include the PGN moves. (optional)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. The response type must be set to `application/x-ndjson` by the request `Accept` header. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to false)
    evals := TODO // interface{} | Include analysis evaluation comments in the PGN, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`  (optional) (default to false)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)
    ongoing := TODO // interface{} | Include ongoing games. The last 3 moves will be omitted. (optional) (default to false)
    finished := TODO // interface{} | Include finished games. Set to `false` to only get ongoing games. (optional) (default to true)
    literate := TODO // interface{} | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
    players := TODO // interface{} | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>  (optional)
    sort := TODO // interface{} | Sort order of the games. (optional) (default to dateDesc)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.ApiGamesUser(context.Background(), username).Since(since).Until(until).Max(max).Vs(vs).Rated(rated).PerfType(perfType).Color(color).Analysed(analysed).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Ongoing(ongoing).Finished(finished).Literate(literate).Players(players).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.ApiGamesUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiGamesUser`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.ApiGamesUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) | The user name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGamesUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | [**interface{}**](interface{}.md) | Download games played since this timestamp. Defaults to account creation date. | 
 **until** | [**interface{}**](interface{}.md) | Download games played until this timestamp. Defaults to now. | 
 **max** | [**interface{}**](interface{}.md) | How many games to download. Leave empty to download all games. | 
 **vs** | [**interface{}**](interface{}.md) | [Filter] Only games played against this opponent | 
 **rated** | [**interface{}**](interface{}.md) | [Filter] Only rated (&#x60;true&#x60;) or casual (&#x60;false&#x60;) games | 
 **perfType** | [**interface{}**](interface{}.md) | [Filter] Only games in these speeds or variants.   Multiple perf types can be specified, separated by a comma.   Example: blitz,rapid,classical | 
 **color** | [**interface{}**](interface{}.md) | [Filter] Only games played as this color. | 
 **analysed** | [**interface{}**](interface{}.md) | [Filter] Only games with or without a computer analysis available | 
 **moves** | [**interface{}**](interface{}.md) | Include the PGN moves. | 
 **pgnInJson** | [**interface{}**](interface{}.md) | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. | [default to false]
 **tags** | [**interface{}**](interface{}.md) | Include the PGN tags. | [default to true]
 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to false]
 **evals** | [**interface{}**](interface{}.md) | Include analysis evaluation comments in the PGN, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60;  | [default to false]
 **opening** | [**interface{}**](interface{}.md) | Include the opening name.  Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]
 **ongoing** | [**interface{}**](interface{}.md) | Include ongoing games. The last 3 moves will be omitted. | [default to false]
 **finished** | [**interface{}**](interface{}.md) | Include finished games. Set to &#x60;false&#x60; to only get ongoing games. | [default to true]
 **literate** | [**interface{}**](interface{}.md) | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **players** | [**interface{}**](interface{}.md) | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: &lt;https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt&gt;  | 
 **sort** | [**interface{}**](interface{}.md) | Sort order of the games. | [default to dateDesc]

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserCurrentGame

> interface{} ApiUserCurrentGame(ctx, username).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Literate(literate).Players(players).Execute()

Export ongoing game of a user



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
    moves := TODO // interface{} | Include the PGN moves. (optional) (default to true)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
    evals := TODO // interface{} | Include analysis evaluation comments in the PGN, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`  (optional) (default to true)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to true)
    literate := TODO // interface{} | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
    players := TODO // interface{} | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.ApiUserCurrentGame(context.Background(), username).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Literate(literate).Players(players).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.ApiUserCurrentGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserCurrentGame`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.ApiUserCurrentGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserCurrentGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moves** | [**interface{}**](interface{}.md) | Include the PGN moves. | [default to true]
 **pgnInJson** | [**interface{}**](interface{}.md) | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | [**interface{}**](interface{}.md) | Include the PGN tags. | [default to true]
 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **evals** | [**interface{}**](interface{}.md) | Include analysis evaluation comments in the PGN, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60;  | [default to true]
 **opening** | [**interface{}**](interface{}.md) | Include the opening name.  Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to true]
 **literate** | [**interface{}**](interface{}.md) | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **players** | [**interface{}**](interface{}.md) | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: &lt;https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt&gt;  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameImport

> interface{} GameImport(ctx).Pgn(pgn).Execute()

Import one game



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
    pgn := TODO // interface{} | The PGN. It can contain only one game. Most standard tags are supported. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GameImport(context.Background()).Pgn(pgn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GameImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameImport`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GameImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pgn** | [**interface{}**](interface{}.md) | The PGN. It can contain only one game. Most standard tags are supported. | 

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


## GamePgn

> interface{} GamePgn(ctx, gameId).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Literate(literate).Players(players).Execute()

Export one game



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
    gameId := TODO // interface{} | The game ID (8 characters).
    moves := TODO // interface{} | Include the PGN moves. (optional) (default to true)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
    evals := TODO // interface{} | Include analysis evaluation comments in the PGN, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`  (optional) (default to true)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to true)
    literate := TODO // interface{} | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
    players := TODO // interface{} | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GamePgn(context.Background(), gameId).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Literate(literate).Players(players).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GamePgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamePgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GamePgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) | The game ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamePgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moves** | [**interface{}**](interface{}.md) | Include the PGN moves. | [default to true]
 **pgnInJson** | [**interface{}**](interface{}.md) | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | [**interface{}**](interface{}.md) | Include the PGN tags. | [default to true]
 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **evals** | [**interface{}**](interface{}.md) | Include analysis evaluation comments in the PGN, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60;  | [default to true]
 **opening** | [**interface{}**](interface{}.md) | Include the opening name.  Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to true]
 **literate** | [**interface{}**](interface{}.md) | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **players** | [**interface{}**](interface{}.md) | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: &lt;https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt&gt;  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByIds

> interface{} GamesByIds(ctx, streamId).Body(body).Execute()

Stream games by IDs



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
    streamId := TODO // interface{} | 
    body := interface{}(987) // interface{} | Up to 500 or 1000 game IDs separated by commas. Example: `gameId01,gameId02,gameId03` 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GamesByIds(context.Background(), streamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GamesByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamesByIds`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GamesByIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | Up to 500 or 1000 game IDs separated by commas. Example: &#x60;gameId01,gameId02,gameId03&#x60;  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByIdsAdd

> Ok GamesByIdsAdd(ctx, streamId).Body(body).Execute()

Add game IDs to stream



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
    streamId := TODO // interface{} | 
    body := interface{}(987) // interface{} | Up to 500 or 1000 game IDs separated by commas. Example: `gameId04,gameId05,gameId06` 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GamesByIdsAdd(context.Background(), streamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GamesByIdsAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamesByIdsAdd`: Ok
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GamesByIdsAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesByIdsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | Up to 500 or 1000 game IDs separated by commas. Example: &#x60;gameId04,gameId05,gameId06&#x60;  | 

### Return type

[**Ok**](Ok.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByUsers

> interface{} GamesByUsers(ctx).Body(body).WithCurrentGames(withCurrentGames).Execute()

Stream games of users



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
    body := interface{}(987) // interface{} | Up to 300 user IDs separated by commas. Example: `aliquantus,chess-network,lovlas` 
    withCurrentGames := TODO // interface{} | Include the already started games at the beginning of the stream. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GamesByUsers(context.Background()).Body(body).WithCurrentGames(withCurrentGames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GamesByUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamesByUsers`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GamesByUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGamesByUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | Up to 300 user IDs separated by commas. Example: &#x60;aliquantus,chess-network,lovlas&#x60;  | 
 **withCurrentGames** | [**interface{}**](interface{}.md) | Include the already started games at the beginning of the stream. | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesExportIds

> interface{} GamesExportIds(ctx).Body(body).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Literate(literate).Players(players).Execute()

Export games by IDs



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
    body := interface{}(TJxUmbWK,4OtIh2oh,ILwozzRZ) // interface{} | Game IDs separated by commas. Up to 300.
    moves := TODO // interface{} | Include the PGN moves. (optional) (default to true)
    pgnInJson := TODO // interface{} | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
    tags := TODO // interface{} | Include the PGN tags. (optional) (default to true)
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to false)
    evals := TODO // interface{} | Include analysis evaluation comments in the PGN, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`  (optional) (default to false)
    opening := TODO // interface{} | Include the opening name.  Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)
    literate := TODO // interface{} | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
    players := TODO // interface{} | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GamesExportIds(context.Background()).Body(body).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Opening(opening).Literate(literate).Players(players).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GamesExportIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GamesExportIds`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GamesExportIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGamesExportIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | Game IDs separated by commas. Up to 300. | 
 **moves** | [**interface{}**](interface{}.md) | Include the PGN moves. | [default to true]
 **pgnInJson** | [**interface{}**](interface{}.md) | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | [**interface{}**](interface{}.md) | Include the PGN tags. | [default to true]
 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to false]
 **evals** | [**interface{}**](interface{}.md) | Include analysis evaluation comments in the PGN, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60;  | [default to false]
 **opening** | [**interface{}**](interface{}.md) | Include the opening name.  Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]
 **literate** | [**interface{}**](interface{}.md) | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.  Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **players** | [**interface{}**](interface{}.md) | URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN. Example: &lt;https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt&gt;  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/x-chess-pgn, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamGame

> interface{} StreamGame(ctx, id).Execute()

Stream moves of a game



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
    resp, r, err := apiClient.GamesApi.StreamGame(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.StreamGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StreamGame`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.StreamGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

