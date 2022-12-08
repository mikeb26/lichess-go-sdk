# GameFinishEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** |  | [optional] 
**Game** | Pointer to [**GameEventInfo**](GameEventInfo.md) |  | [optional] 

## Methods

### NewGameFinishEvent

`func NewGameFinishEvent() *GameFinishEvent`

NewGameFinishEvent instantiates a new GameFinishEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameFinishEventWithDefaults

`func NewGameFinishEventWithDefaults() *GameFinishEvent`

NewGameFinishEventWithDefaults instantiates a new GameFinishEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameFinishEvent) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameFinishEvent) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameFinishEvent) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GameFinishEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GameFinishEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GameFinishEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetGame

`func (o *GameFinishEvent) GetGame() GameEventInfo`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *GameFinishEvent) GetGameOk() (*GameEventInfo, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *GameFinishEvent) SetGame(v GameEventInfo)`

SetGame sets Game field to given value.

### HasGame

`func (o *GameFinishEvent) HasGame() bool`

HasGame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


