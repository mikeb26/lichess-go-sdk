# \ExternalEngineApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiExternalEngineAcquire**](ExternalEngineApi.md#ApiExternalEngineAcquire) | **Post** /api/external-engine/work | Acquire analysis request
[**ApiExternalEngineAnalyse**](ExternalEngineApi.md#ApiExternalEngineAnalyse) | **Post** /api/external-engine/{id}/analyse | Analyse with external engine
[**ApiExternalEngineCreate**](ExternalEngineApi.md#ApiExternalEngineCreate) | **Post** /api/external-engine | Create external engine
[**ApiExternalEngineDelete**](ExternalEngineApi.md#ApiExternalEngineDelete) | **Delete** /api/external-engine/{id} | Delete external engine
[**ApiExternalEngineGet**](ExternalEngineApi.md#ApiExternalEngineGet) | **Get** /api/external-engine/{id} | Get external engine
[**ApiExternalEngineList**](ExternalEngineApi.md#ApiExternalEngineList) | **Get** /api/external-engine | List external engines
[**ApiExternalEnginePut**](ExternalEngineApi.md#ApiExternalEnginePut) | **Put** /api/external-engine/{id} | Update external engine
[**ApiExternalEngineSubmit**](ExternalEngineApi.md#ApiExternalEngineSubmit) | **Post** /api/external-engine/work/{id} | Answer analysis request



## ApiExternalEngineAcquire

> ApiExternalEngineAcquire200Response ApiExternalEngineAcquire(ctx).ApiExternalEngineAcquireRequest(apiExternalEngineAcquireRequest).Execute()

Acquire analysis request



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
    apiExternalEngineAcquireRequest := *openapiclient.NewApiExternalEngineAcquireRequest() // ApiExternalEngineAcquireRequest | Provider credentials.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineAcquire(context.Background()).ApiExternalEngineAcquireRequest(apiExternalEngineAcquireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineAcquire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEngineAcquire`: ApiExternalEngineAcquire200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEngineAcquire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineAcquireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiExternalEngineAcquireRequest** | [**ApiExternalEngineAcquireRequest**](ApiExternalEngineAcquireRequest.md) | Provider credentials. | 

### Return type

[**ApiExternalEngineAcquire200Response**](ApiExternalEngineAcquire200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExternalEngineAnalyse

> ApiExternalEngineAnalyse200Response ApiExternalEngineAnalyse(ctx, id).ApiExternalEngineAnalyseRequest(apiExternalEngineAnalyseRequest).Execute()

Analyse with external engine



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
    id := TODO // interface{} | The external engine id.
    apiExternalEngineAnalyseRequest := *openapiclient.NewApiExternalEngineAnalyseRequest(interface{}(ees_mdF2hK0hlKGSPeC6), *openapiclient.NewExternalEngineWork(interface{}(abcd1234), interface{}(4), interface{}(128), interface{}(1), *openapiclient.NewUciVariant(), interface{}(rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1), interface{}([e2e4, g8f6]))) // ApiExternalEngineAnalyseRequest | Engine credentials and analysis request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineAnalyse(context.Background(), id).ApiExternalEngineAnalyseRequest(apiExternalEngineAnalyseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineAnalyse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEngineAnalyse`: ApiExternalEngineAnalyse200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEngineAnalyse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The external engine id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineAnalyseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiExternalEngineAnalyseRequest** | [**ApiExternalEngineAnalyseRequest**](ApiExternalEngineAnalyseRequest.md) | Engine credentials and analysis request. | 

### Return type

[**ApiExternalEngineAnalyse200Response**](ApiExternalEngineAnalyse200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExternalEngineCreate

> ExternalEngine ApiExternalEngineCreate(ctx).ExternalEngineRegistration(externalEngineRegistration).Execute()

Create external engine



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
    externalEngineRegistration := *openapiclient.NewExternalEngineRegistration(interface{}(Stockfish 15), interface{}(8), interface{}(2048), interface{}(24), interface{}(Dee3uwieZei9ahpaici9bee2yahsai0K)) // ExternalEngineRegistration | A new external engine registration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineCreate(context.Background()).ExternalEngineRegistration(externalEngineRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEngineCreate`: ExternalEngine
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEngineCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalEngineRegistration** | [**ExternalEngineRegistration**](ExternalEngineRegistration.md) | A new external engine registration. | 

### Return type

[**ExternalEngine**](ExternalEngine.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExternalEngineDelete

> Ok ApiExternalEngineDelete(ctx, id).Execute()

Delete external engine



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
    id := TODO // interface{} | The external engine id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEngineDelete`: Ok
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEngineDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The external engine id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineDeleteRequest struct via the builder pattern


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


## ApiExternalEngineGet

> ExternalEngine ApiExternalEngineGet(ctx, id).Execute()

Get external engine



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
    id := TODO // interface{} | The external engine id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEngineGet`: ExternalEngine
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEngineGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The external engine id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEngine**](ExternalEngine.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExternalEngineList

> []ExternalEngine ApiExternalEngineList(ctx).Execute()

List external engines



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
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEngineList`: []ExternalEngine
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEngineList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineListRequest struct via the builder pattern


### Return type

[**[]ExternalEngine**](ExternalEngine.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExternalEnginePut

> ExternalEngine ApiExternalEnginePut(ctx, id).ExternalEngineRegistration(externalEngineRegistration).Execute()

Update external engine



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
    id := TODO // interface{} | The external engine id.
    externalEngineRegistration := *openapiclient.NewExternalEngineRegistration(interface{}(Stockfish 15), interface{}(8), interface{}(2048), interface{}(24), interface{}(Dee3uwieZei9ahpaici9bee2yahsai0K)) // ExternalEngineRegistration | A modified engine registration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEnginePut(context.Background(), id).ExternalEngineRegistration(externalEngineRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEnginePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExternalEnginePut`: ExternalEngine
    fmt.Fprintf(os.Stdout, "Response from `ExternalEngineApi.ApiExternalEnginePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The external engine id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEnginePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalEngineRegistration** | [**ExternalEngineRegistration**](ExternalEngineRegistration.md) | A modified engine registration. | 

### Return type

[**ExternalEngine**](ExternalEngine.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExternalEngineSubmit

> ApiExternalEngineSubmit(ctx, id).Body(body).Execute()

Answer analysis request



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
    body := interface{}(info multipv 1 depth 20 seldepth 30 time 1373 nodes 1494341 score cp 47 hashfull 594 nps 1088376 tbhits 0 pv d2d4 d7d5 c2c4 e7e6 b1c3 f8b4 c4d5 e6d5 g1f3 g8f6 c1g5 h7h6 g5f6 d8f6 d1b3 c7c5 e2e3 b8c6 d4c5 e8g8 f1d3) // interface{} | Analysis results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalEngineApi.ApiExternalEngineSubmit(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalEngineApi.ApiExternalEngineSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExternalEngineSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | Analysis results | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

