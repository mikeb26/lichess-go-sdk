# TablebaseJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **interface{}** | &#x60;cursed-win&#x60; and &#x60;blessed-loss&#x60; means the 50-move rule prevents the decisive result.  &#x60;maybe-win&#x60; and &#x60;maybe-loss&#x60; means exact result is unknown due to [DTZ rounding](https://syzygy-tables.info/metrics#dtz), i.e., the win or loss could also be prevented by the 50-move rule if the user has deviated from the tablebase recommendation since the last pawn move or capture.  | [optional] 
**Dtz** | Pointer to **interface{}** | [DTZ50&#39;&#39; with rounding](https://syzygy-tables.info/metrics#dtz) or null if unknown  | [optional] 
**PreciseDtz** | Pointer to **interface{}** | DTZ50&#39;&#39; (only if guaranteed to be not rounded) or null if unknown  | [optional] 
**Dtm** | Pointer to **interface{}** | Distance to mate (only for positions with not more than 5 pieces) | [optional] 
**Checkmate** | Pointer to **interface{}** |  | [optional] 
**Stalemate** | Pointer to **interface{}** |  | [optional] 
**VariantWin** | Pointer to **interface{}** | Only in chess variants | [optional] 
**VariantLoss** | Pointer to **interface{}** | Only in chess variants | [optional] 
**InsufficientMaterial** | Pointer to **interface{}** |  | [optional] 
**Moves** | Pointer to **interface{}** | Information about legal moves, best first | [optional] 

## Methods

### NewTablebaseJson

`func NewTablebaseJson() *TablebaseJson`

NewTablebaseJson instantiates a new TablebaseJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablebaseJsonWithDefaults

`func NewTablebaseJsonWithDefaults() *TablebaseJson`

NewTablebaseJsonWithDefaults instantiates a new TablebaseJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TablebaseJson) GetCategory() interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TablebaseJson) GetCategoryOk() (*interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TablebaseJson) SetCategory(v interface{})`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TablebaseJson) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TablebaseJson) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TablebaseJson) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDtz

`func (o *TablebaseJson) GetDtz() interface{}`

GetDtz returns the Dtz field if non-nil, zero value otherwise.

### GetDtzOk

`func (o *TablebaseJson) GetDtzOk() (*interface{}, bool)`

GetDtzOk returns a tuple with the Dtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtz

`func (o *TablebaseJson) SetDtz(v interface{})`

SetDtz sets Dtz field to given value.

### HasDtz

`func (o *TablebaseJson) HasDtz() bool`

HasDtz returns a boolean if a field has been set.

### SetDtzNil

`func (o *TablebaseJson) SetDtzNil(b bool)`

 SetDtzNil sets the value for Dtz to be an explicit nil

### UnsetDtz
`func (o *TablebaseJson) UnsetDtz()`

UnsetDtz ensures that no value is present for Dtz, not even an explicit nil
### GetPreciseDtz

`func (o *TablebaseJson) GetPreciseDtz() interface{}`

GetPreciseDtz returns the PreciseDtz field if non-nil, zero value otherwise.

### GetPreciseDtzOk

`func (o *TablebaseJson) GetPreciseDtzOk() (*interface{}, bool)`

GetPreciseDtzOk returns a tuple with the PreciseDtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseDtz

`func (o *TablebaseJson) SetPreciseDtz(v interface{})`

SetPreciseDtz sets PreciseDtz field to given value.

### HasPreciseDtz

`func (o *TablebaseJson) HasPreciseDtz() bool`

HasPreciseDtz returns a boolean if a field has been set.

### SetPreciseDtzNil

`func (o *TablebaseJson) SetPreciseDtzNil(b bool)`

 SetPreciseDtzNil sets the value for PreciseDtz to be an explicit nil

### UnsetPreciseDtz
`func (o *TablebaseJson) UnsetPreciseDtz()`

UnsetPreciseDtz ensures that no value is present for PreciseDtz, not even an explicit nil
### GetDtm

`func (o *TablebaseJson) GetDtm() interface{}`

GetDtm returns the Dtm field if non-nil, zero value otherwise.

### GetDtmOk

`func (o *TablebaseJson) GetDtmOk() (*interface{}, bool)`

GetDtmOk returns a tuple with the Dtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtm

`func (o *TablebaseJson) SetDtm(v interface{})`

SetDtm sets Dtm field to given value.

### HasDtm

`func (o *TablebaseJson) HasDtm() bool`

HasDtm returns a boolean if a field has been set.

### SetDtmNil

`func (o *TablebaseJson) SetDtmNil(b bool)`

 SetDtmNil sets the value for Dtm to be an explicit nil

### UnsetDtm
`func (o *TablebaseJson) UnsetDtm()`

UnsetDtm ensures that no value is present for Dtm, not even an explicit nil
### GetCheckmate

`func (o *TablebaseJson) GetCheckmate() interface{}`

GetCheckmate returns the Checkmate field if non-nil, zero value otherwise.

### GetCheckmateOk

`func (o *TablebaseJson) GetCheckmateOk() (*interface{}, bool)`

GetCheckmateOk returns a tuple with the Checkmate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmate

`func (o *TablebaseJson) SetCheckmate(v interface{})`

SetCheckmate sets Checkmate field to given value.

### HasCheckmate

`func (o *TablebaseJson) HasCheckmate() bool`

HasCheckmate returns a boolean if a field has been set.

### SetCheckmateNil

`func (o *TablebaseJson) SetCheckmateNil(b bool)`

 SetCheckmateNil sets the value for Checkmate to be an explicit nil

### UnsetCheckmate
`func (o *TablebaseJson) UnsetCheckmate()`

UnsetCheckmate ensures that no value is present for Checkmate, not even an explicit nil
### GetStalemate

`func (o *TablebaseJson) GetStalemate() interface{}`

GetStalemate returns the Stalemate field if non-nil, zero value otherwise.

### GetStalemateOk

`func (o *TablebaseJson) GetStalemateOk() (*interface{}, bool)`

GetStalemateOk returns a tuple with the Stalemate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalemate

`func (o *TablebaseJson) SetStalemate(v interface{})`

SetStalemate sets Stalemate field to given value.

### HasStalemate

`func (o *TablebaseJson) HasStalemate() bool`

HasStalemate returns a boolean if a field has been set.

### SetStalemateNil

`func (o *TablebaseJson) SetStalemateNil(b bool)`

 SetStalemateNil sets the value for Stalemate to be an explicit nil

### UnsetStalemate
`func (o *TablebaseJson) UnsetStalemate()`

UnsetStalemate ensures that no value is present for Stalemate, not even an explicit nil
### GetVariantWin

`func (o *TablebaseJson) GetVariantWin() interface{}`

GetVariantWin returns the VariantWin field if non-nil, zero value otherwise.

### GetVariantWinOk

`func (o *TablebaseJson) GetVariantWinOk() (*interface{}, bool)`

GetVariantWinOk returns a tuple with the VariantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantWin

`func (o *TablebaseJson) SetVariantWin(v interface{})`

SetVariantWin sets VariantWin field to given value.

### HasVariantWin

`func (o *TablebaseJson) HasVariantWin() bool`

HasVariantWin returns a boolean if a field has been set.

### SetVariantWinNil

`func (o *TablebaseJson) SetVariantWinNil(b bool)`

 SetVariantWinNil sets the value for VariantWin to be an explicit nil

### UnsetVariantWin
`func (o *TablebaseJson) UnsetVariantWin()`

UnsetVariantWin ensures that no value is present for VariantWin, not even an explicit nil
### GetVariantLoss

`func (o *TablebaseJson) GetVariantLoss() interface{}`

GetVariantLoss returns the VariantLoss field if non-nil, zero value otherwise.

### GetVariantLossOk

`func (o *TablebaseJson) GetVariantLossOk() (*interface{}, bool)`

GetVariantLossOk returns a tuple with the VariantLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantLoss

`func (o *TablebaseJson) SetVariantLoss(v interface{})`

SetVariantLoss sets VariantLoss field to given value.

### HasVariantLoss

`func (o *TablebaseJson) HasVariantLoss() bool`

HasVariantLoss returns a boolean if a field has been set.

### SetVariantLossNil

`func (o *TablebaseJson) SetVariantLossNil(b bool)`

 SetVariantLossNil sets the value for VariantLoss to be an explicit nil

### UnsetVariantLoss
`func (o *TablebaseJson) UnsetVariantLoss()`

UnsetVariantLoss ensures that no value is present for VariantLoss, not even an explicit nil
### GetInsufficientMaterial

`func (o *TablebaseJson) GetInsufficientMaterial() interface{}`

GetInsufficientMaterial returns the InsufficientMaterial field if non-nil, zero value otherwise.

### GetInsufficientMaterialOk

`func (o *TablebaseJson) GetInsufficientMaterialOk() (*interface{}, bool)`

GetInsufficientMaterialOk returns a tuple with the InsufficientMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientMaterial

`func (o *TablebaseJson) SetInsufficientMaterial(v interface{})`

SetInsufficientMaterial sets InsufficientMaterial field to given value.

### HasInsufficientMaterial

`func (o *TablebaseJson) HasInsufficientMaterial() bool`

HasInsufficientMaterial returns a boolean if a field has been set.

### SetInsufficientMaterialNil

`func (o *TablebaseJson) SetInsufficientMaterialNil(b bool)`

 SetInsufficientMaterialNil sets the value for InsufficientMaterial to be an explicit nil

### UnsetInsufficientMaterial
`func (o *TablebaseJson) UnsetInsufficientMaterial()`

UnsetInsufficientMaterial ensures that no value is present for InsufficientMaterial, not even an explicit nil
### GetMoves

`func (o *TablebaseJson) GetMoves() interface{}`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *TablebaseJson) GetMovesOk() (*interface{}, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *TablebaseJson) SetMoves(v interface{})`

SetMoves sets Moves field to given value.

### HasMoves

`func (o *TablebaseJson) HasMoves() bool`

HasMoves returns a boolean if a field has been set.

### SetMovesNil

`func (o *TablebaseJson) SetMovesNil(b bool)`

 SetMovesNil sets the value for Moves to be an explicit nil

### UnsetMoves
`func (o *TablebaseJson) UnsetMoves()`

UnsetMoves ensures that no value is present for Moves, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


