# GameStateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Moves** | **interface{}** | Current moves in UCI format | 
**Wtime** | **interface{}** | Integer of milliseconds White has left on the clock | 
**Btime** | **interface{}** | Integer of milliseconds Black has left on the clock | 
**Winc** | **interface{}** | Integer of White Fisher increment. | 
**Binc** | **interface{}** | Integer of Black Fisher increment. | 
**Status** | [**GameStatus**](GameStatus.md) |  | 
**Winner** | Pointer to **interface{}** | Color of the winner, if any | [optional] 
**Wdraw** | Pointer to **interface{}** | true if white is offering draw, else omitted | [optional] 
**Bdraw** | Pointer to **interface{}** | true if black is offering draw, else omitted | [optional] 
**Wtakeback** | Pointer to **interface{}** | true if white is proposing takeback, else omitted | [optional] 
**Btakeback** | Pointer to **interface{}** | true if black is proposing takeback, else omitted | [optional] 

## Methods

### NewGameStateEvent

`func NewGameStateEvent(type_ interface{}, moves interface{}, wtime interface{}, btime interface{}, winc interface{}, binc interface{}, status GameStatus, ) *GameStateEvent`

NewGameStateEvent instantiates a new GameStateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameStateEventWithDefaults

`func NewGameStateEventWithDefaults() *GameStateEvent`

NewGameStateEventWithDefaults instantiates a new GameStateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameStateEvent) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameStateEvent) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameStateEvent) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GameStateEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GameStateEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMoves

`func (o *GameStateEvent) GetMoves() interface{}`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *GameStateEvent) GetMovesOk() (*interface{}, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *GameStateEvent) SetMoves(v interface{})`

SetMoves sets Moves field to given value.


### SetMovesNil

`func (o *GameStateEvent) SetMovesNil(b bool)`

 SetMovesNil sets the value for Moves to be an explicit nil

### UnsetMoves
`func (o *GameStateEvent) UnsetMoves()`

UnsetMoves ensures that no value is present for Moves, not even an explicit nil
### GetWtime

`func (o *GameStateEvent) GetWtime() interface{}`

GetWtime returns the Wtime field if non-nil, zero value otherwise.

### GetWtimeOk

`func (o *GameStateEvent) GetWtimeOk() (*interface{}, bool)`

GetWtimeOk returns a tuple with the Wtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtime

`func (o *GameStateEvent) SetWtime(v interface{})`

SetWtime sets Wtime field to given value.


### SetWtimeNil

`func (o *GameStateEvent) SetWtimeNil(b bool)`

 SetWtimeNil sets the value for Wtime to be an explicit nil

### UnsetWtime
`func (o *GameStateEvent) UnsetWtime()`

UnsetWtime ensures that no value is present for Wtime, not even an explicit nil
### GetBtime

`func (o *GameStateEvent) GetBtime() interface{}`

GetBtime returns the Btime field if non-nil, zero value otherwise.

### GetBtimeOk

`func (o *GameStateEvent) GetBtimeOk() (*interface{}, bool)`

GetBtimeOk returns a tuple with the Btime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtime

`func (o *GameStateEvent) SetBtime(v interface{})`

SetBtime sets Btime field to given value.


### SetBtimeNil

`func (o *GameStateEvent) SetBtimeNil(b bool)`

 SetBtimeNil sets the value for Btime to be an explicit nil

### UnsetBtime
`func (o *GameStateEvent) UnsetBtime()`

UnsetBtime ensures that no value is present for Btime, not even an explicit nil
### GetWinc

`func (o *GameStateEvent) GetWinc() interface{}`

GetWinc returns the Winc field if non-nil, zero value otherwise.

### GetWincOk

`func (o *GameStateEvent) GetWincOk() (*interface{}, bool)`

GetWincOk returns a tuple with the Winc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinc

`func (o *GameStateEvent) SetWinc(v interface{})`

SetWinc sets Winc field to given value.


### SetWincNil

`func (o *GameStateEvent) SetWincNil(b bool)`

 SetWincNil sets the value for Winc to be an explicit nil

### UnsetWinc
`func (o *GameStateEvent) UnsetWinc()`

UnsetWinc ensures that no value is present for Winc, not even an explicit nil
### GetBinc

`func (o *GameStateEvent) GetBinc() interface{}`

GetBinc returns the Binc field if non-nil, zero value otherwise.

### GetBincOk

`func (o *GameStateEvent) GetBincOk() (*interface{}, bool)`

GetBincOk returns a tuple with the Binc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinc

`func (o *GameStateEvent) SetBinc(v interface{})`

SetBinc sets Binc field to given value.


### SetBincNil

`func (o *GameStateEvent) SetBincNil(b bool)`

 SetBincNil sets the value for Binc to be an explicit nil

### UnsetBinc
`func (o *GameStateEvent) UnsetBinc()`

UnsetBinc ensures that no value is present for Binc, not even an explicit nil
### GetStatus

`func (o *GameStateEvent) GetStatus() GameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GameStateEvent) GetStatusOk() (*GameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GameStateEvent) SetStatus(v GameStatus)`

SetStatus sets Status field to given value.


### GetWinner

`func (o *GameStateEvent) GetWinner() interface{}`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GameStateEvent) GetWinnerOk() (*interface{}, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GameStateEvent) SetWinner(v interface{})`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GameStateEvent) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### SetWinnerNil

`func (o *GameStateEvent) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *GameStateEvent) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetWdraw

`func (o *GameStateEvent) GetWdraw() interface{}`

GetWdraw returns the Wdraw field if non-nil, zero value otherwise.

### GetWdrawOk

`func (o *GameStateEvent) GetWdrawOk() (*interface{}, bool)`

GetWdrawOk returns a tuple with the Wdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdraw

`func (o *GameStateEvent) SetWdraw(v interface{})`

SetWdraw sets Wdraw field to given value.

### HasWdraw

`func (o *GameStateEvent) HasWdraw() bool`

HasWdraw returns a boolean if a field has been set.

### SetWdrawNil

`func (o *GameStateEvent) SetWdrawNil(b bool)`

 SetWdrawNil sets the value for Wdraw to be an explicit nil

### UnsetWdraw
`func (o *GameStateEvent) UnsetWdraw()`

UnsetWdraw ensures that no value is present for Wdraw, not even an explicit nil
### GetBdraw

`func (o *GameStateEvent) GetBdraw() interface{}`

GetBdraw returns the Bdraw field if non-nil, zero value otherwise.

### GetBdrawOk

`func (o *GameStateEvent) GetBdrawOk() (*interface{}, bool)`

GetBdrawOk returns a tuple with the Bdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdraw

`func (o *GameStateEvent) SetBdraw(v interface{})`

SetBdraw sets Bdraw field to given value.

### HasBdraw

`func (o *GameStateEvent) HasBdraw() bool`

HasBdraw returns a boolean if a field has been set.

### SetBdrawNil

`func (o *GameStateEvent) SetBdrawNil(b bool)`

 SetBdrawNil sets the value for Bdraw to be an explicit nil

### UnsetBdraw
`func (o *GameStateEvent) UnsetBdraw()`

UnsetBdraw ensures that no value is present for Bdraw, not even an explicit nil
### GetWtakeback

`func (o *GameStateEvent) GetWtakeback() interface{}`

GetWtakeback returns the Wtakeback field if non-nil, zero value otherwise.

### GetWtakebackOk

`func (o *GameStateEvent) GetWtakebackOk() (*interface{}, bool)`

GetWtakebackOk returns a tuple with the Wtakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtakeback

`func (o *GameStateEvent) SetWtakeback(v interface{})`

SetWtakeback sets Wtakeback field to given value.

### HasWtakeback

`func (o *GameStateEvent) HasWtakeback() bool`

HasWtakeback returns a boolean if a field has been set.

### SetWtakebackNil

`func (o *GameStateEvent) SetWtakebackNil(b bool)`

 SetWtakebackNil sets the value for Wtakeback to be an explicit nil

### UnsetWtakeback
`func (o *GameStateEvent) UnsetWtakeback()`

UnsetWtakeback ensures that no value is present for Wtakeback, not even an explicit nil
### GetBtakeback

`func (o *GameStateEvent) GetBtakeback() interface{}`

GetBtakeback returns the Btakeback field if non-nil, zero value otherwise.

### GetBtakebackOk

`func (o *GameStateEvent) GetBtakebackOk() (*interface{}, bool)`

GetBtakebackOk returns a tuple with the Btakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtakeback

`func (o *GameStateEvent) SetBtakeback(v interface{})`

SetBtakeback sets Btakeback field to given value.

### HasBtakeback

`func (o *GameStateEvent) HasBtakeback() bool`

HasBtakeback returns a boolean if a field has been set.

### SetBtakebackNil

`func (o *GameStateEvent) SetBtakebackNil(b bool)`

 SetBtakebackNil sets the value for Btakeback to be an explicit nil

### UnsetBtakeback
`func (o *GameStateEvent) UnsetBtakeback()`

UnsetBtakeback ensures that no value is present for Btakeback, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


