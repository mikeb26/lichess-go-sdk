/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login, gameplay, and more: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/ZackClements/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/ShailChoksi/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), please wait a full minute before resuming API usage.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](http://ndjson.org/), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses. 

API version: 2.0.0
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// StudiesApiService StudiesApi service
type StudiesApiService service

type ApiStudyAllChaptersPgnRequest struct {
	ctx context.Context
	ApiService *StudiesApiService
	studyId interface{}
	clocks *interface{}
	comments *interface{}
	variations *interface{}
}

// Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r ApiStudyAllChaptersPgnRequest) Clocks(clocks interface{}) ApiStudyAllChaptersPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis and annotator comments in the PGN moves, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60; 
func (r ApiStudyAllChaptersPgnRequest) Comments(comments interface{}) ApiStudyAllChaptersPgnRequest {
	r.comments = &comments
	return r
}

// Include non-mainline moves, when available.  Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60; 
func (r ApiStudyAllChaptersPgnRequest) Variations(variations interface{}) ApiStudyAllChaptersPgnRequest {
	r.variations = &variations
	return r
}

func (r ApiStudyAllChaptersPgnRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.StudyAllChaptersPgnExecute(r)
}

/*
StudyAllChaptersPgn Export all chapters

Download all chapters of a study in PGN format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID (8 characters).
 @return ApiStudyAllChaptersPgnRequest
*/
func (a *StudiesApiService) StudyAllChaptersPgn(ctx context.Context, studyId interface{}) ApiStudyAllChaptersPgnRequest {
	return ApiStudyAllChaptersPgnRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
	}
}

// Execute executes the request
//  @return interface{}
func (a *StudiesApiService) StudyAllChaptersPgnExecute(r ApiStudyAllChaptersPgnRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesApiService.StudyAllChaptersPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterToString(r.studyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clocks != nil {
		localVarQueryParams.Add("clocks", parameterToString(*r.clocks, ""))
	}
	if r.comments != nil {
		localVarQueryParams.Add("comments", parameterToString(*r.comments, ""))
	}
	if r.variations != nil {
		localVarQueryParams.Add("variations", parameterToString(*r.variations, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStudyChapterPgnRequest struct {
	ctx context.Context
	ApiService *StudiesApiService
	studyId interface{}
	chapterId interface{}
	clocks *interface{}
	comments *interface{}
	variations *interface{}
}

// Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r ApiStudyChapterPgnRequest) Clocks(clocks interface{}) ApiStudyChapterPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis and annotator comments in the PGN moves, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60; 
func (r ApiStudyChapterPgnRequest) Comments(comments interface{}) ApiStudyChapterPgnRequest {
	r.comments = &comments
	return r
}

// Include non-mainline moves, when available.  Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60; 
func (r ApiStudyChapterPgnRequest) Variations(variations interface{}) ApiStudyChapterPgnRequest {
	r.variations = &variations
	return r
}

func (r ApiStudyChapterPgnRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.StudyChapterPgnExecute(r)
}

/*
StudyChapterPgn Export one study chapter

Download one study chapter in PGN format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID (8 characters).
 @param chapterId The chapter ID (8 characters).
 @return ApiStudyChapterPgnRequest
*/
func (a *StudiesApiService) StudyChapterPgn(ctx context.Context, studyId interface{}, chapterId interface{}) ApiStudyChapterPgnRequest {
	return ApiStudyChapterPgnRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
		chapterId: chapterId,
	}
}

// Execute executes the request
//  @return interface{}
func (a *StudiesApiService) StudyChapterPgnExecute(r ApiStudyChapterPgnRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesApiService.StudyChapterPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/study/{studyId}/{chapterId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterToString(r.studyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chapterId"+"}", url.PathEscape(parameterToString(r.chapterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clocks != nil {
		localVarQueryParams.Add("clocks", parameterToString(*r.clocks, ""))
	}
	if r.comments != nil {
		localVarQueryParams.Add("comments", parameterToString(*r.comments, ""))
	}
	if r.variations != nil {
		localVarQueryParams.Add("variations", parameterToString(*r.variations, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStudyExportAllPgnRequest struct {
	ctx context.Context
	ApiService *StudiesApiService
	username interface{}
	clocks *interface{}
	comments *interface{}
	variations *interface{}
}

// Include clock comments in the PGN moves, when available.  Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r ApiStudyExportAllPgnRequest) Clocks(clocks interface{}) ApiStudyExportAllPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis and annotator comments in the PGN moves, when available.  Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60; 
func (r ApiStudyExportAllPgnRequest) Comments(comments interface{}) ApiStudyExportAllPgnRequest {
	r.comments = &comments
	return r
}

// Include non-mainline moves, when available.  Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60; 
func (r ApiStudyExportAllPgnRequest) Variations(variations interface{}) ApiStudyExportAllPgnRequest {
	r.variations = &variations
	return r
}

func (r ApiStudyExportAllPgnRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.StudyExportAllPgnExecute(r)
}

/*
StudyExportAllPgn Export all studies of a user

Download all chapters of all studies of a user in PGN format.

If authenticated, then all public, unlisted, and private studies are included.

If not, only public, listed studies are included.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user whose studies we export
 @return ApiStudyExportAllPgnRequest
*/
func (a *StudiesApiService) StudyExportAllPgn(ctx context.Context, username interface{}) ApiStudyExportAllPgnRequest {
	return ApiStudyExportAllPgnRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return interface{}
func (a *StudiesApiService) StudyExportAllPgnExecute(r ApiStudyExportAllPgnRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesApiService.StudyExportAllPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/study/by/{username}/export.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterToString(r.username, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clocks != nil {
		localVarQueryParams.Add("clocks", parameterToString(*r.clocks, ""))
	}
	if r.comments != nil {
		localVarQueryParams.Add("comments", parameterToString(*r.comments, ""))
	}
	if r.variations != nil {
		localVarQueryParams.Add("variations", parameterToString(*r.variations, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
