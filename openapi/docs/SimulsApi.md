# \SimulsApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSimul**](SimulsApi.md#ApiSimul) | **Get** /api/simul | Get current simuls



## ApiSimul

> []interface{} ApiSimul(ctx).Execute()

Get current simuls



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
    resp, r, err := apiClient.SimulsApi.ApiSimul(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulsApi.ApiSimul``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSimul`: []interface{}
    fmt.Fprintf(os.Stdout, "Response from `SimulsApi.ApiSimul`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiSimulRequest struct via the builder pattern


### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

