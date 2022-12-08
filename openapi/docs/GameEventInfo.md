# GameEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** |  | [optional] 
**Source** | Pointer to **interface{}** |  | [optional] 
**Compat** | Pointer to [**GameEventInfoCompat**](GameEventInfoCompat.md) |  | [optional] 

## Methods

### NewGameEventInfo

`func NewGameEventInfo() *GameEventInfo`

NewGameEventInfo instantiates a new GameEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameEventInfoWithDefaults

`func NewGameEventInfoWithDefaults() *GameEventInfo`

NewGameEventInfoWithDefaults instantiates a new GameEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameEventInfo) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameEventInfo) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameEventInfo) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *GameEventInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GameEventInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GameEventInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSource

`func (o *GameEventInfo) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GameEventInfo) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GameEventInfo) SetSource(v interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *GameEventInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *GameEventInfo) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GameEventInfo) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCompat

`func (o *GameEventInfo) GetCompat() GameEventInfoCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *GameEventInfo) GetCompatOk() (*GameEventInfoCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *GameEventInfo) SetCompat(v GameEventInfoCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *GameEventInfo) HasCompat() bool`

HasCompat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


