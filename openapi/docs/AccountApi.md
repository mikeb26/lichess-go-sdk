# \AccountApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Account**](AccountApi.md#Account) | **Get** /api/account/preferences | Get my preferences
[**AccountEmail**](AccountApi.md#AccountEmail) | **Get** /api/account/email | Get my email address
[**AccountKid**](AccountApi.md#AccountKid) | **Get** /api/account/kid | Get my kid mode status
[**AccountKidPost**](AccountApi.md#AccountKidPost) | **Post** /api/account/kid | Set my kid mode status
[**AccountMe**](AccountApi.md#AccountMe) | **Get** /api/account | Get my profile



## Account

> Account200Response Account(ctx).Execute()

Get my preferences



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
    resp, r, err := apiClient.AccountApi.Account(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.Account``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Account`: Account200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.Account`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountRequest struct via the builder pattern


### Return type

[**Account200Response**](Account200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountEmail

> AccountEmail200Response AccountEmail(ctx).Execute()

Get my email address



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
    resp, r, err := apiClient.AccountApi.AccountEmail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountEmail`: AccountEmail200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountEmailRequest struct via the builder pattern


### Return type

[**AccountEmail200Response**](AccountEmail200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountKid

> AccountKid200Response AccountKid(ctx).Execute()

Get my kid mode status



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
    resp, r, err := apiClient.AccountApi.AccountKid(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountKid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountKid`: AccountKid200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountKid`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountKidRequest struct via the builder pattern


### Return type

[**AccountKid200Response**](AccountKid200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountKidPost

> Ok AccountKidPost(ctx).V(v).Execute()

Set my kid mode status



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
    v := TODO // interface{} | Kid mode status

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.AccountKidPost(context.Background()).V(v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountKidPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountKidPost`: Ok
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountKidPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountKidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v** | [**interface{}**](interface{}.md) | Kid mode status | 

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


## AccountMe

> interface{} AccountMe(ctx).Execute()

Get my profile



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
    resp, r, err := apiClient.AccountApi.AccountMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountMe`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountMeRequest struct via the builder pattern


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

