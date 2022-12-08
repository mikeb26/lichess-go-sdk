# \StudiesApi

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StudyAllChaptersPgn**](StudiesApi.md#StudyAllChaptersPgn) | **Get** /api/study/{studyId}.pgn | Export all chapters
[**StudyChapterPgn**](StudiesApi.md#StudyChapterPgn) | **Get** /study/{studyId}/{chapterId}.pgn | Export one study chapter
[**StudyExportAllPgn**](StudiesApi.md#StudyExportAllPgn) | **Get** /study/by/{username}/export.pgn | Export all studies of a user



## StudyAllChaptersPgn

> interface{} StudyAllChaptersPgn(ctx, studyId).Clocks(clocks).Comments(comments).Variations(variations).Execute()

Export all chapters



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
    studyId := TODO // interface{} | The study ID (8 characters).
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
    comments := TODO // interface{} | Include analysis and annotator comments in the PGN moves, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`  (optional) (default to true)
    variations := TODO // interface{} | Include non-mainline moves, when available.  Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StudiesApi.StudyAllChaptersPgn(context.Background(), studyId).Clocks(clocks).Comments(comments).Variations(variations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StudiesApi.StudyAllChaptersPgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StudyAllChaptersPgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StudiesApi.StudyAllChaptersPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | [**interface{}**](.md) | The study ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyAllChaptersPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | [**interface{}**](interface{}.md) | Include analysis and annotator comments in the PGN moves, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60;  | [default to true]
 **variations** | [**interface{}**](interface{}.md) | Include non-mainline moves, when available.  Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60;  | [default to true]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StudyChapterPgn

> interface{} StudyChapterPgn(ctx, studyId, chapterId).Clocks(clocks).Comments(comments).Variations(variations).Execute()

Export one study chapter



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
    studyId := TODO // interface{} | The study ID (8 characters).
    chapterId := TODO // interface{} | The chapter ID (8 characters).
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
    comments := TODO // interface{} | Include analysis and annotator comments in the PGN moves, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`  (optional) (default to true)
    variations := TODO // interface{} | Include non-mainline moves, when available.  Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StudiesApi.StudyChapterPgn(context.Background(), studyId, chapterId).Clocks(clocks).Comments(comments).Variations(variations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StudiesApi.StudyChapterPgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StudyChapterPgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StudiesApi.StudyChapterPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | [**interface{}**](.md) | The study ID (8 characters). | 
**chapterId** | [**interface{}**](.md) | The chapter ID (8 characters). | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyChapterPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | [**interface{}**](interface{}.md) | Include analysis and annotator comments in the PGN moves, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60;  | [default to true]
 **variations** | [**interface{}**](interface{}.md) | Include non-mainline moves, when available.  Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60;  | [default to true]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StudyExportAllPgn

> interface{} StudyExportAllPgn(ctx, username).Clocks(clocks).Comments(comments).Variations(variations).Execute()

Export all studies of a user



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
    username := TODO // interface{} | The user whose studies we export
    clocks := TODO // interface{} | Include clock comments in the PGN moves, when available.  Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
    comments := TODO // interface{} | Include analysis and annotator comments in the PGN moves, when available.  Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`  (optional) (default to true)
    variations := TODO // interface{} | Include non-mainline moves, when available.  Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StudiesApi.StudyExportAllPgn(context.Background(), username).Clocks(clocks).Comments(comments).Variations(variations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StudiesApi.StudyExportAllPgn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StudyExportAllPgn`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StudiesApi.StudyExportAllPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | [**interface{}**](.md) | The user whose studies we export | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyExportAllPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | [**interface{}**](interface{}.md) | Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | [**interface{}**](interface{}.md) | Include analysis and annotator comments in the PGN moves, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60;  | [default to true]
 **variations** | [**interface{}**](interface{}.md) | Include non-mainline moves, when available.  Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60;  | [default to true]

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

