# \TablebaseApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AntichessAtomic**](TablebaseApi.md#AntichessAtomic) | **Get** /antichess | Tablebase lookup for Antichess
[**TablebaseAtomic**](TablebaseApi.md#TablebaseAtomic) | **Get** /atomic | Tablebase lookup for Atomic chess
[**TablebaseStandard**](TablebaseApi.md#TablebaseStandard) | **Get** /standard | Tablebase lookup



## AntichessAtomic

> interface{} AntichessAtomic(ctx).Execute()

Tablebase lookup for Antichess



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
    resp, r, err := apiClient.TablebaseApi.AntichessAtomic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablebaseApi.AntichessAtomic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AntichessAtomic`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TablebaseApi.AntichessAtomic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAntichessAtomicRequest struct via the builder pattern


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


## TablebaseAtomic

> interface{} TablebaseAtomic(ctx).Execute()

Tablebase lookup for Atomic chess



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
    resp, r, err := apiClient.TablebaseApi.TablebaseAtomic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablebaseApi.TablebaseAtomic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TablebaseAtomic`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `TablebaseApi.TablebaseAtomic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTablebaseAtomicRequest struct via the builder pattern


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


## TablebaseStandard

> TablebaseJson TablebaseStandard(ctx).Fen(fen).Execute()

Tablebase lookup



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
    fen := TODO // interface{} | FEN of the position. Underscores allowed.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TablebaseApi.TablebaseStandard(context.Background()).Fen(fen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TablebaseApi.TablebaseStandard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TablebaseStandard`: TablebaseJson
    fmt.Fprintf(os.Stdout, "Response from `TablebaseApi.TablebaseStandard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTablebaseStandardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | [**interface{}**](interface{}.md) | FEN of the position. Underscores allowed. | 

### Return type

[**TablebaseJson**](TablebaseJson.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

