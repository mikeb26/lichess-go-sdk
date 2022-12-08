# ApiExternalEngineAnalyse200ResponsePvsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | **interface{}** | Current search depth of the pv | 
**Cp** | Pointer to **interface{}** | Evaluation in centi-pawns, from White&#39;s point of view | [optional] 
**Mate** | Pointer to **interface{}** | Evaluation in signed moves to mate, from White&#39;s point of view | [optional] 
**Moves** | **[]interface{}** | Variation in UCI notation | 

## Methods

### NewApiExternalEngineAnalyse200ResponsePvsInner

`func NewApiExternalEngineAnalyse200ResponsePvsInner(depth interface{}, moves []interface{}, ) *ApiExternalEngineAnalyse200ResponsePvsInner`

NewApiExternalEngineAnalyse200ResponsePvsInner instantiates a new ApiExternalEngineAnalyse200ResponsePvsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyse200ResponsePvsInnerWithDefaults

`func NewApiExternalEngineAnalyse200ResponsePvsInnerWithDefaults() *ApiExternalEngineAnalyse200ResponsePvsInner`

NewApiExternalEngineAnalyse200ResponsePvsInnerWithDefaults instantiates a new ApiExternalEngineAnalyse200ResponsePvsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetDepth() interface{}`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetDepthOk() (*interface{}, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetDepth(v interface{})`

SetDepth sets Depth field to given value.


### SetDepthNil

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetDepthNil(b bool)`

 SetDepthNil sets the value for Depth to be an explicit nil

### UnsetDepth
`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) UnsetDepth()`

UnsetDepth ensures that no value is present for Depth, not even an explicit nil
### GetCp

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetCp() interface{}`

GetCp returns the Cp field if non-nil, zero value otherwise.

### GetCpOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetCpOk() (*interface{}, bool)`

GetCpOk returns a tuple with the Cp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCp

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetCp(v interface{})`

SetCp sets Cp field to given value.

### HasCp

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) HasCp() bool`

HasCp returns a boolean if a field has been set.

### SetCpNil

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetCpNil(b bool)`

 SetCpNil sets the value for Cp to be an explicit nil

### UnsetCp
`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) UnsetCp()`

UnsetCp ensures that no value is present for Cp, not even an explicit nil
### GetMate

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMate() interface{}`

GetMate returns the Mate field if non-nil, zero value otherwise.

### GetMateOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMateOk() (*interface{}, bool)`

GetMateOk returns a tuple with the Mate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMate

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetMate(v interface{})`

SetMate sets Mate field to given value.

### HasMate

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) HasMate() bool`

HasMate returns a boolean if a field has been set.

### SetMateNil

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetMateNil(b bool)`

 SetMateNil sets the value for Mate to be an explicit nil

### UnsetMate
`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) UnsetMate()`

UnsetMate ensures that no value is present for Mate, not even an explicit nil
### GetMoves

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMoves() []interface{}`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMovesOk() (*[]interface{}, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetMoves(v []interface{})`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


