# \AnalysisApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCloudEval**](AnalysisApi.md#ApiCloudEval) | **Get** /api/cloud-eval | Get cloud evaluation of a position.



## ApiCloudEval

> ApiCloudEval(ctx).Fen(fen).MultiPv(multiPv).Variant(variant).Execute()

Get cloud evaluation of a position.



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
    fen := TODO // interface{} | FEN of the position
    multiPv := TODO // interface{} | Number of variations (optional) (default to 1)
    variant := *openapiclient.NewVariantKey() // VariantKey | Variant (optional) (default to standard)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalysisApi.ApiCloudEval(context.Background()).Fen(fen).MultiPv(multiPv).Variant(variant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalysisApi.ApiCloudEval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCloudEvalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | [**interface{}**](interface{}.md) | FEN of the position | 
 **multiPv** | [**interface{}**](interface{}.md) | Number of variations | [default to 1]
 **variant** | [**VariantKey**](VariantKey.md) | Variant | [default to standard]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

