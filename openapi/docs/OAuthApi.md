# \OAuthApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiToken**](OAuthApi.md#ApiToken) | **Post** /api/token | Obtain access token
[**ApiTokenDelete**](OAuthApi.md#ApiTokenDelete) | **Delete** /api/token | Revoke access token
[**Oauth**](OAuthApi.md#Oauth) | **Get** /oauth | Request authorization code
[**TokenTest**](OAuthApi.md#TokenTest) | **Post** /api/token/test | Test multiple OAuth tokens



## ApiToken

> ApiToken(ctx).GrantType(grantType).Code(code).CodeVerifier(codeVerifier).RedirectUri(redirectUri).ClientId(clientId).Execute()

Obtain access token



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
    grantType := TODO // interface{} | Must be `authorization_code`. (optional)
    code := TODO // interface{} | The authorization code that was sent in the `code` parameter to your `redirect_uri`. (optional)
    codeVerifier := TODO // interface{} | A `code_challenge` was used to request the authorization code. This must be the `code_verifier` it was derived from. (optional)
    redirectUri := TODO // interface{} | Must match the `redirect_uri` used to request the authorization code. (optional)
    clientId := TODO // interface{} | Must match the `client_id` used to request the authorization code. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthApi.ApiToken(context.Background()).GrantType(grantType).Code(code).CodeVerifier(codeVerifier).RedirectUri(redirectUri).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.ApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | [**interface{}**](interface{}.md) | Must be &#x60;authorization_code&#x60;. | 
 **code** | [**interface{}**](interface{}.md) | The authorization code that was sent in the &#x60;code&#x60; parameter to your &#x60;redirect_uri&#x60;. | 
 **codeVerifier** | [**interface{}**](interface{}.md) | A &#x60;code_challenge&#x60; was used to request the authorization code. This must be the &#x60;code_verifier&#x60; it was derived from. | 
 **redirectUri** | [**interface{}**](interface{}.md) | Must match the &#x60;redirect_uri&#x60; used to request the authorization code. | 
 **clientId** | [**interface{}**](interface{}.md) | Must match the &#x60;client_id&#x60; used to request the authorization code. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTokenDelete

> ApiTokenDelete(ctx).Execute()

Revoke access token



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
    resp, r, err := apiClient.OAuthApi.ApiTokenDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.ApiTokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiTokenDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth

> Oauth(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).CodeChallengeMethod(codeChallengeMethod).CodeChallenge(codeChallenge).Scope(scope).Username(username).State(state).Execute()

Request authorization code



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
    responseType := TODO // interface{} | Must be `code`.
    clientId := TODO // interface{} | Arbitrary identifier that uniquely identifies your application.
    redirectUri := TODO // interface{} | The absolute URL that the user should be redirected to with the authorization result.
    codeChallengeMethod := TODO // interface{} | Must be `S256`.
    codeChallenge := TODO // interface{} | Compute `BASE64URL(SHA256(code_verifier))`.
    scope := TODO // interface{} | Space separated list of requested OAuth scopes, if any. (optional)
    username := TODO // interface{} | Hint that you want the user to log in with a specific Lichess username. (optional)
    state := TODO // interface{} | Arbitrary state that will be returned verbatim with the authorization result. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthApi.Oauth(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).CodeChallengeMethod(codeChallengeMethod).CodeChallenge(codeChallenge).Scope(scope).Username(username).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.Oauth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | [**interface{}**](interface{}.md) | Must be &#x60;code&#x60;. | 
 **clientId** | [**interface{}**](interface{}.md) | Arbitrary identifier that uniquely identifies your application. | 
 **redirectUri** | [**interface{}**](interface{}.md) | The absolute URL that the user should be redirected to with the authorization result. | 
 **codeChallengeMethod** | [**interface{}**](interface{}.md) | Must be &#x60;S256&#x60;. | 
 **codeChallenge** | [**interface{}**](interface{}.md) | Compute &#x60;BASE64URL(SHA256(code_verifier))&#x60;. | 
 **scope** | [**interface{}**](interface{}.md) | Space separated list of requested OAuth scopes, if any. | 
 **username** | [**interface{}**](interface{}.md) | Hint that you want the user to log in with a specific Lichess username. | 
 **state** | [**interface{}**](interface{}.md) | Arbitrary state that will be returned verbatim with the authorization result. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenTest

> []map[string]interface{} TokenTest(ctx).Body(body).Execute()

Test multiple OAuth tokens



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
    body := interface{}(lip_AvsS88TozFeSMEaoLN5c,lip_badToken) // interface{} | OAuth tokens separated by commas. Up to 1000.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthApi.TokenTest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.TokenTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenTest`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.TokenTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | OAuth tokens separated by commas. Up to 1000. | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

