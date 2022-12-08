# \MessagingApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InboxUsername**](MessagingApi.md#InboxUsername) | **Post** /inbox/{username} | Send a private message



## InboxUsername

> Ok InboxUsername(ctx, username).Text(text).Execute()

Send a private message



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
    text := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagingApi.InboxUsername(context.Background(), username).Text(text).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagingApi.InboxUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InboxUsername`: Ok
    fmt.Fprintf(os.Stdout, "Response from `MessagingApi.InboxUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboxUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **text** | [**interface{}**](interface{}.md) |  | 

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

