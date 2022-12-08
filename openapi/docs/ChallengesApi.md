# \ChallengesApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminChallengeTokens**](ChallengesApi.md#AdminChallengeTokens) | **Post** /api/token/admin-challenge | Admin challenge tokens
[**ChallengeAccept**](ChallengesApi.md#ChallengeAccept) | **Post** /api/challenge/{challengeId}/accept | Accept a challenge
[**ChallengeAi**](ChallengesApi.md#ChallengeAi) | **Post** /api/challenge/ai | Challenge the AI
[**ChallengeCancel**](ChallengesApi.md#ChallengeCancel) | **Post** /api/challenge/{challengeId}/cancel | Cancel a challenge
[**ChallengeCreate**](ChallengesApi.md#ChallengeCreate) | **Post** /api/challenge/{username} | Create a challenge
[**ChallengeDecline**](ChallengesApi.md#ChallengeDecline) | **Post** /api/challenge/{challengeId}/decline | Decline a challenge
[**ChallengeList**](ChallengesApi.md#ChallengeList) | **Get** /api/challenge | List your challenges
[**ChallengeOpen**](ChallengesApi.md#ChallengeOpen) | **Post** /api/challenge/open | Open-ended challenge
[**ChallengeStartClocks**](ChallengesApi.md#ChallengeStartClocks) | **Post** /api/challenge/{gameId}/start-clocks | Start clocks of a game
[**RoundAddTime**](ChallengesApi.md#RoundAddTime) | **Post** /api/round/{gameId}/add-time/{seconds} | Add time to the opponent clock



## AdminChallengeTokens

> AdminChallengeTokens(ctx).Users(users).Description(description).Execute()

Admin challenge tokens



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
    users := TODO // interface{} | Usernames separated with commas
    description := TODO // interface{} | User visible description of the token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.AdminChallengeTokens(context.Background()).Users(users).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.AdminChallengeTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminChallengeTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | [**interface{}**](interface{}.md) | Usernames separated with commas | 
 **description** | [**interface{}**](interface{}.md) | User visible description of the token | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeAccept

> Ok ChallengeAccept(ctx, challengeId).Execute()

Accept a challenge



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
    challengeId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeAccept(context.Background(), challengeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeAccept`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeAcceptRequest struct via the builder pattern


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


## ChallengeAi

> GameJson ChallengeAi(ctx).Level(level).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Color(color).Variant(variant).Fen(fen).Execute()

Challenge the AI



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
    level := TODO // interface{} | AI strength (optional)
    clockLimit := TODO // interface{} | Clock initial time in seconds. If empty, a correspondence game is created. (optional)
    clockIncrement := TODO // interface{} | Clock increment in seconds. If empty, a correspondence game is created. (optional)
    days := TODO // interface{} | Days per move, for correspondence games. Clock settings must be omitted. (optional)
    color := TODO // interface{} | Which color you get to play (optional) (default to random)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    fen := TODO // interface{} | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeAi(context.Background()).Level(level).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Color(color).Variant(variant).Fen(fen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeAi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeAi`: GameJson
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeAi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChallengeAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | [**interface{}**](interface{}.md) | AI strength | 
 **clockLimit** | [**interface{}**](interface{}.md) | Clock initial time in seconds. If empty, a correspondence game is created. | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds. If empty, a correspondence game is created. | 
 **days** | [**interface{}**](interface{}.md) | Days per move, for correspondence games. Clock settings must be omitted. | 
 **color** | [**interface{}**](interface{}.md) | Which color you get to play | [default to random]
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1]

### Return type

[**GameJson**](GameJson.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeCancel

> Ok ChallengeCancel(ctx, challengeId).OpponentToken(opponentToken).Execute()

Cancel a challenge



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
    challengeId := TODO // interface{} | 
    opponentToken := TODO // interface{} | Optional `challenge:write` token of the opponent. If set, the game can be canceled even if both players have moved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeCancel(context.Background(), challengeId).OpponentToken(opponentToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeCancel`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opponentToken** | [**interface{}**](interface{}.md) | Optional &#x60;challenge:write&#x60; token of the opponent. If set, the game can be canceled even if both players have moved. | 

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


## ChallengeCreate

> ChallengeJson ChallengeCreate(ctx, username).Rated(rated).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Color(color).Variant(variant).Fen(fen).KeepAliveStream(keepAliveStream).AcceptByToken(acceptByToken).Message(message).Rules(rules).Execute()

Create a challenge



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
    rated := TODO // interface{} | Game is rated and impacts players ratings (optional) (default to false)
    clockLimit := TODO // interface{} | Clock initial time in seconds. If empty, a correspondence game is created. (optional)
    clockIncrement := TODO // interface{} | Clock increment in seconds. If empty, a correspondence game is created. (optional)
    days := TODO // interface{} | Days per move, for correspondence games. Clock settings must be omitted. (optional)
    color := TODO // interface{} | Which color you get to play (optional) (default to random)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    fen := TODO // interface{} | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. Castling moves will use UCI_Chess960 notation, for example e1h1 instead of e1g1. (optional) (default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1)
    keepAliveStream := TODO // interface{} | If set, the response is streamed as [ndjson](#section/Introduction/Streaming-with-ND-JSON). The challenge is kept alive until the connection is closed by the client. When the challenge is accepted, declined or canceled, a message of the form `{\\\"done\\\":\\\"accepted\\\"}` is sent, then the connection is closed by the server. If not set, the response is not streamed, and the challenge expires after 20s if not accepted.  (optional)
    acceptByToken := TODO // interface{} | Immediately accept the challenge and create the game. Pass in an OAuth token (with the `challenge:write` scope) for the receiving user. On success, the response will contain a `game` field instead of a `challenge` field.  Alternatively, consider the [bulk pairing API](#operation/bulkPairingCreate).  (optional)
    message := TODO // interface{} | **Only if `acceptByToken` is set.**  Message that is sent to each player, when the game is created. It is sent from your user account.  `{opponent}`, `{player}` and `{game}` are placeholders that will be replaced with the opponent name, player name, and the game URLs.  You can omit this field to send the default message, but if you set your own message, it must at least contain the `{game}` placeholder.  (optional) (default to Your game with {opponent} is ready: {game}.)
    rules := TODO // interface{} | Extra game rules separated by commas. Example: `noAbort,noRematch`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeCreate(context.Background(), username).Rated(rated).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Color(color).Variant(variant).Fen(fen).KeepAliveStream(keepAliveStream).AcceptByToken(acceptByToken).Message(message).Rules(rules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeCreate`: ChallengeJson
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rated** | [**interface{}**](interface{}.md) | Game is rated and impacts players ratings | [default to false]
 **clockLimit** | [**interface{}**](interface{}.md) | Clock initial time in seconds. If empty, a correspondence game is created. | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds. If empty, a correspondence game is created. | 
 **days** | [**interface{}**](interface{}.md) | Days per move, for correspondence games. Clock settings must be omitted. | 
 **color** | [**interface{}**](interface{}.md) | Which color you get to play | [default to random]
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. Castling moves will use UCI_Chess960 notation, for example e1h1 instead of e1g1. | [default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1]
 **keepAliveStream** | [**interface{}**](interface{}.md) | If set, the response is streamed as [ndjson](#section/Introduction/Streaming-with-ND-JSON). The challenge is kept alive until the connection is closed by the client. When the challenge is accepted, declined or canceled, a message of the form &#x60;{\\\&quot;done\\\&quot;:\\\&quot;accepted\\\&quot;}&#x60; is sent, then the connection is closed by the server. If not set, the response is not streamed, and the challenge expires after 20s if not accepted.  | 
 **acceptByToken** | [**interface{}**](interface{}.md) | Immediately accept the challenge and create the game. Pass in an OAuth token (with the &#x60;challenge:write&#x60; scope) for the receiving user. On success, the response will contain a &#x60;game&#x60; field instead of a &#x60;challenge&#x60; field.  Alternatively, consider the [bulk pairing API](#operation/bulkPairingCreate).  | 
 **message** | [**interface{}**](interface{}.md) | **Only if &#x60;acceptByToken&#x60; is set.**  Message that is sent to each player, when the game is created. It is sent from your user account.  &#x60;{opponent}&#x60;, &#x60;{player}&#x60; and &#x60;{game}&#x60; are placeholders that will be replaced with the opponent name, player name, and the game URLs.  You can omit this field to send the default message, but if you set your own message, it must at least contain the &#x60;{game}&#x60; placeholder.  | [default to Your game with {opponent} is ready: {game}.]
 **rules** | [**interface{}**](interface{}.md) | Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60;  | 

### Return type

[**ChallengeJson**](ChallengeJson.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeDecline

> Ok ChallengeDecline(ctx, challengeId).Reason(reason).Execute()

Decline a challenge



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
    challengeId := TODO // interface{} | 
    reason := TODO // interface{} | Reason challenge was declined. It will be translated to the player's language. See [the full list in the translation file](https://github.com/ornicar/lila/blob/master/translation/source/challenge.xml#L14). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeDecline(context.Background(), challengeId).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeDecline`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | [**interface{}**](interface{}.md) | Reason challenge was declined. It will be translated to the player&#39;s language. See [the full list in the translation file](https://github.com/ornicar/lila/blob/master/translation/source/challenge.xml#L14). | 

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


## ChallengeList

> ChallengeList200Response ChallengeList(ctx).Execute()

List your challenges



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
    resp, r, err := apiClient.ChallengesApi.ChallengeList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeList`: ChallengeList200Response
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeListRequest struct via the builder pattern


### Return type

[**ChallengeList200Response**](ChallengeList200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeOpen

> interface{} ChallengeOpen(ctx).Rated(rated).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Variant(variant).Fen(fen).Name(name).Rules(rules).Users(users).Execute()

Open-ended challenge



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
    rated := TODO // interface{} | Game is rated and impacts players ratings (optional) (default to false)
    clockLimit := TODO // interface{} | Clock initial time in seconds. If empty, a correspondence game is created. (optional)
    clockIncrement := TODO // interface{} | Clock increment in seconds. If empty, a correspondence game is created. (optional)
    days := TODO // interface{} | Days per turn. For correspondence challenges. (optional)
    variant := *openapiclient.NewVariantKey() // VariantKey |  (optional) (default to standard)
    fen := TODO // interface{} | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1)
    name := TODO // interface{} | Optional name for the challenge, that players will see on the challenge page. (optional)
    rules := TODO // interface{} | Extra game rules separated by commas. Example: `noAbort,noRematch`  (optional)
    users := TODO // interface{} | Optional pair of usernames, separated by a comma. If set, only these users will be allowed to join the game. The first username gets the white pieces. Example: `Username1,Username2`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeOpen(context.Background()).Rated(rated).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Variant(variant).Fen(fen).Name(name).Rules(rules).Users(users).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeOpen`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeOpen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChallengeOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rated** | [**interface{}**](interface{}.md) | Game is rated and impacts players ratings | [default to false]
 **clockLimit** | [**interface{}**](interface{}.md) | Clock initial time in seconds. If empty, a correspondence game is created. | 
 **clockIncrement** | [**interface{}**](interface{}.md) | Clock increment in seconds. If empty, a correspondence game is created. | 
 **days** | [**interface{}**](interface{}.md) | Days per turn. For correspondence challenges. | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
 **fen** | [**interface{}**](interface{}.md) | Custom initial position (in FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1]
 **name** | [**interface{}**](interface{}.md) | Optional name for the challenge, that players will see on the challenge page. | 
 **rules** | [**interface{}**](interface{}.md) | Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60;  | 
 **users** | [**interface{}**](interface{}.md) | Optional pair of usernames, separated by a comma. If set, only these users will be allowed to join the game. The first username gets the white pieces. Example: &#x60;Username1,Username2&#x60;  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeStartClocks

> Ok ChallengeStartClocks(ctx, gameId).Token1(token1).Token2(token2).Execute()

Start clocks of a game



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
    token1 := TODO // interface{} | OAuth token of a player (optional)
    token2 := TODO // interface{} | OAuth token of the other player (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.ChallengeStartClocks(context.Background(), gameId).Token1(token1).Token2(token2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.ChallengeStartClocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChallengeStartClocks`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.ChallengeStartClocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeStartClocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token1** | [**interface{}**](interface{}.md) | OAuth token of a player | 
 **token2** | [**interface{}**](interface{}.md) | OAuth token of the other player | 

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


## RoundAddTime

> Ok RoundAddTime(ctx, gameId, seconds).Execute()

Add time to the opponent clock



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
    seconds := TODO // interface{} | How many seconds to give

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChallengesApi.RoundAddTime(context.Background(), gameId, seconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChallengesApi.RoundAddTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoundAddTime`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ChallengesApi.RoundAddTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | [**interface{}**](.md) |  | 
**seconds** | [**interface{}**](.md) | How many seconds to give | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoundAddTimeRequest struct via the builder pattern


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

