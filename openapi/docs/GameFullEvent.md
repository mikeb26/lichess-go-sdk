# GameFullEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Id** | **interface{}** |  | 
**Variant** | [**Variant**](Variant.md) |  | 
**Clock** | **interface{}** |  | 
**Speed** | [**Speed**](Speed.md) |  | 
**Perf** | [**GameFullEventPerf**](GameFullEventPerf.md) |  | 
**Rated** | **interface{}** |  | 
**CreatedAt** | **interface{}** |  | 
**White** | [**GameEventPlayer**](GameEventPlayer.md) |  | 
**Black** | [**GameEventPlayer**](GameEventPlayer.md) |  | 
**InitialFen** | **interface{}** |  | [default to startpos]
**State** | [**GameStateEvent**](GameStateEvent.md) |  | 
**TournamentId** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGameFullEvent

`func NewGameFullEvent(type_ interface{}, id interface{}, variant Variant, clock interface{}, speed Speed, perf GameFullEventPerf, rated interface{}, createdAt interface{}, white GameEventPlayer, black GameEventPlayer, initialFen interface{}, state GameStateEvent, ) *GameFullEvent`

NewGameFullEvent instantiates a new GameFullEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameFullEventWithDefaults

`func NewGameFullEventWithDefaults() *GameFullEvent`

NewGameFullEventWithDefaults instantiates a new GameFullEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameFullEvent) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameFullEvent) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameFullEvent) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GameFullEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GameFullEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *GameFullEvent) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameFullEvent) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameFullEvent) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *GameFullEvent) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GameFullEvent) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetVariant

`func (o *GameFullEvent) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GameFullEvent) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GameFullEvent) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetClock

`func (o *GameFullEvent) GetClock() interface{}`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GameFullEvent) GetClockOk() (*interface{}, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GameFullEvent) SetClock(v interface{})`

SetClock sets Clock field to given value.


### SetClockNil

`func (o *GameFullEvent) SetClockNil(b bool)`

 SetClockNil sets the value for Clock to be an explicit nil

### UnsetClock
`func (o *GameFullEvent) UnsetClock()`

UnsetClock ensures that no value is present for Clock, not even an explicit nil
### GetSpeed

`func (o *GameFullEvent) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GameFullEvent) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GameFullEvent) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *GameFullEvent) GetPerf() GameFullEventPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GameFullEvent) GetPerfOk() (*GameFullEventPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GameFullEvent) SetPerf(v GameFullEventPerf)`

SetPerf sets Perf field to given value.


### GetRated

`func (o *GameFullEvent) GetRated() interface{}`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GameFullEvent) GetRatedOk() (*interface{}, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GameFullEvent) SetRated(v interface{})`

SetRated sets Rated field to given value.


### SetRatedNil

`func (o *GameFullEvent) SetRatedNil(b bool)`

 SetRatedNil sets the value for Rated to be an explicit nil

### UnsetRated
`func (o *GameFullEvent) UnsetRated()`

UnsetRated ensures that no value is present for Rated, not even an explicit nil
### GetCreatedAt

`func (o *GameFullEvent) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameFullEvent) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameFullEvent) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GameFullEvent) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GameFullEvent) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetWhite

`func (o *GameFullEvent) GetWhite() GameEventPlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *GameFullEvent) GetWhiteOk() (*GameEventPlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *GameFullEvent) SetWhite(v GameEventPlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *GameFullEvent) GetBlack() GameEventPlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *GameFullEvent) GetBlackOk() (*GameEventPlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *GameFullEvent) SetBlack(v GameEventPlayer)`

SetBlack sets Black field to given value.


### GetInitialFen

`func (o *GameFullEvent) GetInitialFen() interface{}`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *GameFullEvent) GetInitialFenOk() (*interface{}, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *GameFullEvent) SetInitialFen(v interface{})`

SetInitialFen sets InitialFen field to given value.


### SetInitialFenNil

`func (o *GameFullEvent) SetInitialFenNil(b bool)`

 SetInitialFenNil sets the value for InitialFen to be an explicit nil

### UnsetInitialFen
`func (o *GameFullEvent) UnsetInitialFen()`

UnsetInitialFen ensures that no value is present for InitialFen, not even an explicit nil
### GetState

`func (o *GameFullEvent) GetState() GameStateEvent`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GameFullEvent) GetStateOk() (*GameStateEvent, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GameFullEvent) SetState(v GameStateEvent)`

SetState sets State field to given value.


### GetTournamentId

`func (o *GameFullEvent) GetTournamentId() interface{}`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *GameFullEvent) GetTournamentIdOk() (*interface{}, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *GameFullEvent) SetTournamentId(v interface{})`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *GameFullEvent) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.

### SetTournamentIdNil

`func (o *GameFullEvent) SetTournamentIdNil(b bool)`

 SetTournamentIdNil sets the value for TournamentId to be an explicit nil

### UnsetTournamentId
`func (o *GameFullEvent) UnsetTournamentId()`

UnsetTournamentId ensures that no value is present for TournamentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


