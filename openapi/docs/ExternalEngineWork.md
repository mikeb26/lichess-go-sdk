# ExternalEngineWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **interface{}** | Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.  | 
**Threads** | **interface{}** | Number of threads to use for analysis. | 
**Hash** | **interface{}** | Hash table size to use for analysis, in MiB. | 
**Infinite** | Pointer to **interface{}** | Request an infinite search (rather than roughly aiming for &#x60;defaultDepth&#x60;).  | [optional] 
**MultiPv** | **interface{}** | Requested number of principal variations. | 
**Variant** | [**UciVariant**](UciVariant.md) |  | [default to chess]
**InitialFen** | **interface{}** | Initial position of the game. | 
**Moves** | **interface{}** | List of moves played from the initial position, in UCI notation. | 

## Methods

### NewExternalEngineWork

`func NewExternalEngineWork(sessionId interface{}, threads interface{}, hash interface{}, multiPv interface{}, variant UciVariant, initialFen interface{}, moves interface{}, ) *ExternalEngineWork`

NewExternalEngineWork instantiates a new ExternalEngineWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineWorkWithDefaults

`func NewExternalEngineWorkWithDefaults() *ExternalEngineWork`

NewExternalEngineWorkWithDefaults instantiates a new ExternalEngineWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *ExternalEngineWork) GetSessionId() interface{}`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExternalEngineWork) GetSessionIdOk() (*interface{}, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExternalEngineWork) SetSessionId(v interface{})`

SetSessionId sets SessionId field to given value.


### SetSessionIdNil

`func (o *ExternalEngineWork) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *ExternalEngineWork) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetThreads

`func (o *ExternalEngineWork) GetThreads() interface{}`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ExternalEngineWork) GetThreadsOk() (*interface{}, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ExternalEngineWork) SetThreads(v interface{})`

SetThreads sets Threads field to given value.


### SetThreadsNil

`func (o *ExternalEngineWork) SetThreadsNil(b bool)`

 SetThreadsNil sets the value for Threads to be an explicit nil

### UnsetThreads
`func (o *ExternalEngineWork) UnsetThreads()`

UnsetThreads ensures that no value is present for Threads, not even an explicit nil
### GetHash

`func (o *ExternalEngineWork) GetHash() interface{}`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ExternalEngineWork) GetHashOk() (*interface{}, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ExternalEngineWork) SetHash(v interface{})`

SetHash sets Hash field to given value.


### SetHashNil

`func (o *ExternalEngineWork) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *ExternalEngineWork) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetInfinite

`func (o *ExternalEngineWork) GetInfinite() interface{}`

GetInfinite returns the Infinite field if non-nil, zero value otherwise.

### GetInfiniteOk

`func (o *ExternalEngineWork) GetInfiniteOk() (*interface{}, bool)`

GetInfiniteOk returns a tuple with the Infinite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfinite

`func (o *ExternalEngineWork) SetInfinite(v interface{})`

SetInfinite sets Infinite field to given value.

### HasInfinite

`func (o *ExternalEngineWork) HasInfinite() bool`

HasInfinite returns a boolean if a field has been set.

### SetInfiniteNil

`func (o *ExternalEngineWork) SetInfiniteNil(b bool)`

 SetInfiniteNil sets the value for Infinite to be an explicit nil

### UnsetInfinite
`func (o *ExternalEngineWork) UnsetInfinite()`

UnsetInfinite ensures that no value is present for Infinite, not even an explicit nil
### GetMultiPv

`func (o *ExternalEngineWork) GetMultiPv() interface{}`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ExternalEngineWork) GetMultiPvOk() (*interface{}, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ExternalEngineWork) SetMultiPv(v interface{})`

SetMultiPv sets MultiPv field to given value.


### SetMultiPvNil

`func (o *ExternalEngineWork) SetMultiPvNil(b bool)`

 SetMultiPvNil sets the value for MultiPv to be an explicit nil

### UnsetMultiPv
`func (o *ExternalEngineWork) UnsetMultiPv()`

UnsetMultiPv ensures that no value is present for MultiPv, not even an explicit nil
### GetVariant

`func (o *ExternalEngineWork) GetVariant() UciVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ExternalEngineWork) GetVariantOk() (*UciVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ExternalEngineWork) SetVariant(v UciVariant)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ExternalEngineWork) GetInitialFen() interface{}`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ExternalEngineWork) GetInitialFenOk() (*interface{}, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ExternalEngineWork) SetInitialFen(v interface{})`

SetInitialFen sets InitialFen field to given value.


### SetInitialFenNil

`func (o *ExternalEngineWork) SetInitialFenNil(b bool)`

 SetInitialFenNil sets the value for InitialFen to be an explicit nil

### UnsetInitialFen
`func (o *ExternalEngineWork) UnsetInitialFen()`

UnsetInitialFen ensures that no value is present for InitialFen, not even an explicit nil
### GetMoves

`func (o *ExternalEngineWork) GetMoves() interface{}`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ExternalEngineWork) GetMovesOk() (*interface{}, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ExternalEngineWork) SetMoves(v interface{})`

SetMoves sets Moves field to given value.


### SetMovesNil

`func (o *ExternalEngineWork) SetMovesNil(b bool)`

 SetMovesNil sets the value for Moves to be an explicit nil

### UnsetMoves
`func (o *ExternalEngineWork) UnsetMoves()`

UnsetMoves ensures that no value is present for Moves, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


