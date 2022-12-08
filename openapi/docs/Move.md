# Move

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | Pointer to **interface{}** |  | [optional] 
**San** | Pointer to **interface{}** |  | [optional] 
**Category** | Pointer to **interface{}** |  | [optional] 
**Dtz** | Pointer to **interface{}** | DTZ50&#39;&#39; with rounding or null if unknown | [optional] 
**PreciseDtz** | Pointer to **interface{}** | DTZ50&#39;&#39; (only if guaranteed to be not rounded) or null if unknown  | [optional] 
**Dtm** | Pointer to **interface{}** | Distance to mate (only for positions with not more than 5 pieces) | [optional] 
**Zeroing** | Pointer to **interface{}** |  | [optional] 
**Checkmate** | Pointer to **interface{}** |  | [optional] 
**Stalemate** | Pointer to **interface{}** |  | [optional] 
**VariantWin** | Pointer to **interface{}** |  | [optional] 
**VariantLoss** | Pointer to **interface{}** |  | [optional] 
**InsufficientMaterial** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewMove

`func NewMove() *Move`

NewMove instantiates a new Move object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveWithDefaults

`func NewMoveWithDefaults() *Move`

NewMoveWithDefaults instantiates a new Move object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *Move) GetUci() interface{}`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *Move) GetUciOk() (*interface{}, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *Move) SetUci(v interface{})`

SetUci sets Uci field to given value.

### HasUci

`func (o *Move) HasUci() bool`

HasUci returns a boolean if a field has been set.

### SetUciNil

`func (o *Move) SetUciNil(b bool)`

 SetUciNil sets the value for Uci to be an explicit nil

### UnsetUci
`func (o *Move) UnsetUci()`

UnsetUci ensures that no value is present for Uci, not even an explicit nil
### GetSan

`func (o *Move) GetSan() interface{}`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *Move) GetSanOk() (*interface{}, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *Move) SetSan(v interface{})`

SetSan sets San field to given value.

### HasSan

`func (o *Move) HasSan() bool`

HasSan returns a boolean if a field has been set.

### SetSanNil

`func (o *Move) SetSanNil(b bool)`

 SetSanNil sets the value for San to be an explicit nil

### UnsetSan
`func (o *Move) UnsetSan()`

UnsetSan ensures that no value is present for San, not even an explicit nil
### GetCategory

`func (o *Move) GetCategory() interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Move) GetCategoryOk() (*interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Move) SetCategory(v interface{})`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Move) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Move) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Move) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDtz

`func (o *Move) GetDtz() interface{}`

GetDtz returns the Dtz field if non-nil, zero value otherwise.

### GetDtzOk

`func (o *Move) GetDtzOk() (*interface{}, bool)`

GetDtzOk returns a tuple with the Dtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtz

`func (o *Move) SetDtz(v interface{})`

SetDtz sets Dtz field to given value.

### HasDtz

`func (o *Move) HasDtz() bool`

HasDtz returns a boolean if a field has been set.

### SetDtzNil

`func (o *Move) SetDtzNil(b bool)`

 SetDtzNil sets the value for Dtz to be an explicit nil

### UnsetDtz
`func (o *Move) UnsetDtz()`

UnsetDtz ensures that no value is present for Dtz, not even an explicit nil
### GetPreciseDtz

`func (o *Move) GetPreciseDtz() interface{}`

GetPreciseDtz returns the PreciseDtz field if non-nil, zero value otherwise.

### GetPreciseDtzOk

`func (o *Move) GetPreciseDtzOk() (*interface{}, bool)`

GetPreciseDtzOk returns a tuple with the PreciseDtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseDtz

`func (o *Move) SetPreciseDtz(v interface{})`

SetPreciseDtz sets PreciseDtz field to given value.

### HasPreciseDtz

`func (o *Move) HasPreciseDtz() bool`

HasPreciseDtz returns a boolean if a field has been set.

### SetPreciseDtzNil

`func (o *Move) SetPreciseDtzNil(b bool)`

 SetPreciseDtzNil sets the value for PreciseDtz to be an explicit nil

### UnsetPreciseDtz
`func (o *Move) UnsetPreciseDtz()`

UnsetPreciseDtz ensures that no value is present for PreciseDtz, not even an explicit nil
### GetDtm

`func (o *Move) GetDtm() interface{}`

GetDtm returns the Dtm field if non-nil, zero value otherwise.

### GetDtmOk

`func (o *Move) GetDtmOk() (*interface{}, bool)`

GetDtmOk returns a tuple with the Dtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtm

`func (o *Move) SetDtm(v interface{})`

SetDtm sets Dtm field to given value.

### HasDtm

`func (o *Move) HasDtm() bool`

HasDtm returns a boolean if a field has been set.

### SetDtmNil

`func (o *Move) SetDtmNil(b bool)`

 SetDtmNil sets the value for Dtm to be an explicit nil

### UnsetDtm
`func (o *Move) UnsetDtm()`

UnsetDtm ensures that no value is present for Dtm, not even an explicit nil
### GetZeroing

`func (o *Move) GetZeroing() interface{}`

GetZeroing returns the Zeroing field if non-nil, zero value otherwise.

### GetZeroingOk

`func (o *Move) GetZeroingOk() (*interface{}, bool)`

GetZeroingOk returns a tuple with the Zeroing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroing

`func (o *Move) SetZeroing(v interface{})`

SetZeroing sets Zeroing field to given value.

### HasZeroing

`func (o *Move) HasZeroing() bool`

HasZeroing returns a boolean if a field has been set.

### SetZeroingNil

`func (o *Move) SetZeroingNil(b bool)`

 SetZeroingNil sets the value for Zeroing to be an explicit nil

### UnsetZeroing
`func (o *Move) UnsetZeroing()`

UnsetZeroing ensures that no value is present for Zeroing, not even an explicit nil
### GetCheckmate

`func (o *Move) GetCheckmate() interface{}`

GetCheckmate returns the Checkmate field if non-nil, zero value otherwise.

### GetCheckmateOk

`func (o *Move) GetCheckmateOk() (*interface{}, bool)`

GetCheckmateOk returns a tuple with the Checkmate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmate

`func (o *Move) SetCheckmate(v interface{})`

SetCheckmate sets Checkmate field to given value.

### HasCheckmate

`func (o *Move) HasCheckmate() bool`

HasCheckmate returns a boolean if a field has been set.

### SetCheckmateNil

`func (o *Move) SetCheckmateNil(b bool)`

 SetCheckmateNil sets the value for Checkmate to be an explicit nil

### UnsetCheckmate
`func (o *Move) UnsetCheckmate()`

UnsetCheckmate ensures that no value is present for Checkmate, not even an explicit nil
### GetStalemate

`func (o *Move) GetStalemate() interface{}`

GetStalemate returns the Stalemate field if non-nil, zero value otherwise.

### GetStalemateOk

`func (o *Move) GetStalemateOk() (*interface{}, bool)`

GetStalemateOk returns a tuple with the Stalemate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalemate

`func (o *Move) SetStalemate(v interface{})`

SetStalemate sets Stalemate field to given value.

### HasStalemate

`func (o *Move) HasStalemate() bool`

HasStalemate returns a boolean if a field has been set.

### SetStalemateNil

`func (o *Move) SetStalemateNil(b bool)`

 SetStalemateNil sets the value for Stalemate to be an explicit nil

### UnsetStalemate
`func (o *Move) UnsetStalemate()`

UnsetStalemate ensures that no value is present for Stalemate, not even an explicit nil
### GetVariantWin

`func (o *Move) GetVariantWin() interface{}`

GetVariantWin returns the VariantWin field if non-nil, zero value otherwise.

### GetVariantWinOk

`func (o *Move) GetVariantWinOk() (*interface{}, bool)`

GetVariantWinOk returns a tuple with the VariantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantWin

`func (o *Move) SetVariantWin(v interface{})`

SetVariantWin sets VariantWin field to given value.

### HasVariantWin

`func (o *Move) HasVariantWin() bool`

HasVariantWin returns a boolean if a field has been set.

### SetVariantWinNil

`func (o *Move) SetVariantWinNil(b bool)`

 SetVariantWinNil sets the value for VariantWin to be an explicit nil

### UnsetVariantWin
`func (o *Move) UnsetVariantWin()`

UnsetVariantWin ensures that no value is present for VariantWin, not even an explicit nil
### GetVariantLoss

`func (o *Move) GetVariantLoss() interface{}`

GetVariantLoss returns the VariantLoss field if non-nil, zero value otherwise.

### GetVariantLossOk

`func (o *Move) GetVariantLossOk() (*interface{}, bool)`

GetVariantLossOk returns a tuple with the VariantLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantLoss

`func (o *Move) SetVariantLoss(v interface{})`

SetVariantLoss sets VariantLoss field to given value.

### HasVariantLoss

`func (o *Move) HasVariantLoss() bool`

HasVariantLoss returns a boolean if a field has been set.

### SetVariantLossNil

`func (o *Move) SetVariantLossNil(b bool)`

 SetVariantLossNil sets the value for VariantLoss to be an explicit nil

### UnsetVariantLoss
`func (o *Move) UnsetVariantLoss()`

UnsetVariantLoss ensures that no value is present for VariantLoss, not even an explicit nil
### GetInsufficientMaterial

`func (o *Move) GetInsufficientMaterial() interface{}`

GetInsufficientMaterial returns the InsufficientMaterial field if non-nil, zero value otherwise.

### GetInsufficientMaterialOk

`func (o *Move) GetInsufficientMaterialOk() (*interface{}, bool)`

GetInsufficientMaterialOk returns a tuple with the InsufficientMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientMaterial

`func (o *Move) SetInsufficientMaterial(v interface{})`

SetInsufficientMaterial sets InsufficientMaterial field to given value.

### HasInsufficientMaterial

`func (o *Move) HasInsufficientMaterial() bool`

HasInsufficientMaterial returns a boolean if a field has been set.

### SetInsufficientMaterialNil

`func (o *Move) SetInsufficientMaterialNil(b bool)`

 SetInsufficientMaterialNil sets the value for InsufficientMaterial to be an explicit nil

### UnsetInsufficientMaterial
`func (o *Move) UnsetInsufficientMaterial()`

UnsetInsufficientMaterial ensures that no value is present for InsufficientMaterial, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


