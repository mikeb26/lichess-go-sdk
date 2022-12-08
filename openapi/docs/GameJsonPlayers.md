# GameJsonPlayers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**White** | Pointer to [**GameUser**](GameUser.md) |  | [optional] 
**Black** | Pointer to [**GameUser**](GameUser.md) |  | [optional] 

## Methods

### NewGameJsonPlayers

`func NewGameJsonPlayers() *GameJsonPlayers`

NewGameJsonPlayers instantiates a new GameJsonPlayers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameJsonPlayersWithDefaults

`func NewGameJsonPlayersWithDefaults() *GameJsonPlayers`

NewGameJsonPlayersWithDefaults instantiates a new GameJsonPlayers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhite

`func (o *GameJsonPlayers) GetWhite() GameUser`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *GameJsonPlayers) GetWhiteOk() (*GameUser, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *GameJsonPlayers) SetWhite(v GameUser)`

SetWhite sets White field to given value.

### HasWhite

`func (o *GameJsonPlayers) HasWhite() bool`

HasWhite returns a boolean if a field has been set.

### GetBlack

`func (o *GameJsonPlayers) GetBlack() GameUser`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *GameJsonPlayers) GetBlackOk() (*GameUser, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *GameJsonPlayers) SetBlack(v GameUser)`

SetBlack sets Black field to given value.

### HasBlack

`func (o *GameJsonPlayers) HasBlack() bool`

HasBlack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


