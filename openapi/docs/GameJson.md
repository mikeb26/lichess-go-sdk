# GameJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** |  | 
**Rated** | **interface{}** |  | 
**Variant** | [**VariantKey**](VariantKey.md) |  | [default to standard]
**Speed** | [**Speed**](Speed.md) |  | 
**Perf** | **interface{}** |  | 
**CreatedAt** | **interface{}** |  | 
**LastMoveAt** | **interface{}** |  | 
**Status** | [**GameStatus**](GameStatus.md) |  | 
**Players** | [**GameJsonPlayers**](GameJsonPlayers.md) |  | 
**InitialFen** | Pointer to **interface{}** |  | [optional] 
**Winner** | Pointer to **interface{}** |  | [optional] 
**Opening** | Pointer to [**GameJsonOpening**](GameJsonOpening.md) |  | [optional] 
**Moves** | Pointer to **interface{}** |  | [optional] 
**Pgn** | Pointer to **interface{}** |  | [optional] 
**DaysPerTurn** | Pointer to **interface{}** |  | [optional] 
**Analysis** | Pointer to **interface{}** |  | [optional] 
**Tournament** | Pointer to **interface{}** |  | [optional] 
**Swiss** | Pointer to **interface{}** |  | [optional] 
**Clock** | Pointer to [**GameJsonClock**](GameJsonClock.md) |  | [optional] 

## Methods

### NewGameJson

`func NewGameJson(id interface{}, rated interface{}, variant VariantKey, speed Speed, perf interface{}, createdAt interface{}, lastMoveAt interface{}, status GameStatus, players GameJsonPlayers, ) *GameJson`

NewGameJson instantiates a new GameJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameJsonWithDefaults

`func NewGameJsonWithDefaults() *GameJson`

NewGameJsonWithDefaults instantiates a new GameJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameJson) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameJson) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameJson) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *GameJson) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GameJson) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRated

`func (o *GameJson) GetRated() interface{}`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GameJson) GetRatedOk() (*interface{}, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GameJson) SetRated(v interface{})`

SetRated sets Rated field to given value.


### SetRatedNil

`func (o *GameJson) SetRatedNil(b bool)`

 SetRatedNil sets the value for Rated to be an explicit nil

### UnsetRated
`func (o *GameJson) UnsetRated()`

UnsetRated ensures that no value is present for Rated, not even an explicit nil
### GetVariant

`func (o *GameJson) GetVariant() VariantKey`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GameJson) GetVariantOk() (*VariantKey, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GameJson) SetVariant(v VariantKey)`

SetVariant sets Variant field to given value.


### GetSpeed

`func (o *GameJson) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GameJson) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GameJson) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *GameJson) GetPerf() interface{}`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GameJson) GetPerfOk() (*interface{}, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GameJson) SetPerf(v interface{})`

SetPerf sets Perf field to given value.


### SetPerfNil

`func (o *GameJson) SetPerfNil(b bool)`

 SetPerfNil sets the value for Perf to be an explicit nil

### UnsetPerf
`func (o *GameJson) UnsetPerf()`

UnsetPerf ensures that no value is present for Perf, not even an explicit nil
### GetCreatedAt

`func (o *GameJson) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameJson) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameJson) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GameJson) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GameJson) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastMoveAt

`func (o *GameJson) GetLastMoveAt() interface{}`

GetLastMoveAt returns the LastMoveAt field if non-nil, zero value otherwise.

### GetLastMoveAtOk

`func (o *GameJson) GetLastMoveAtOk() (*interface{}, bool)`

GetLastMoveAtOk returns a tuple with the LastMoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMoveAt

`func (o *GameJson) SetLastMoveAt(v interface{})`

SetLastMoveAt sets LastMoveAt field to given value.


### SetLastMoveAtNil

`func (o *GameJson) SetLastMoveAtNil(b bool)`

 SetLastMoveAtNil sets the value for LastMoveAt to be an explicit nil

### UnsetLastMoveAt
`func (o *GameJson) UnsetLastMoveAt()`

UnsetLastMoveAt ensures that no value is present for LastMoveAt, not even an explicit nil
### GetStatus

`func (o *GameJson) GetStatus() GameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GameJson) GetStatusOk() (*GameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GameJson) SetStatus(v GameStatus)`

SetStatus sets Status field to given value.


### GetPlayers

`func (o *GameJson) GetPlayers() GameJsonPlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GameJson) GetPlayersOk() (*GameJsonPlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GameJson) SetPlayers(v GameJsonPlayers)`

SetPlayers sets Players field to given value.


### GetInitialFen

`func (o *GameJson) GetInitialFen() interface{}`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *GameJson) GetInitialFenOk() (*interface{}, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *GameJson) SetInitialFen(v interface{})`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *GameJson) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### SetInitialFenNil

`func (o *GameJson) SetInitialFenNil(b bool)`

 SetInitialFenNil sets the value for InitialFen to be an explicit nil

### UnsetInitialFen
`func (o *GameJson) UnsetInitialFen()`

UnsetInitialFen ensures that no value is present for InitialFen, not even an explicit nil
### GetWinner

`func (o *GameJson) GetWinner() interface{}`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GameJson) GetWinnerOk() (*interface{}, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GameJson) SetWinner(v interface{})`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GameJson) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### SetWinnerNil

`func (o *GameJson) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *GameJson) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetOpening

`func (o *GameJson) GetOpening() GameJsonOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *GameJson) GetOpeningOk() (*GameJsonOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *GameJson) SetOpening(v GameJsonOpening)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *GameJson) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

### GetMoves

`func (o *GameJson) GetMoves() interface{}`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *GameJson) GetMovesOk() (*interface{}, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *GameJson) SetMoves(v interface{})`

SetMoves sets Moves field to given value.

### HasMoves

`func (o *GameJson) HasMoves() bool`

HasMoves returns a boolean if a field has been set.

### SetMovesNil

`func (o *GameJson) SetMovesNil(b bool)`

 SetMovesNil sets the value for Moves to be an explicit nil

### UnsetMoves
`func (o *GameJson) UnsetMoves()`

UnsetMoves ensures that no value is present for Moves, not even an explicit nil
### GetPgn

`func (o *GameJson) GetPgn() interface{}`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *GameJson) GetPgnOk() (*interface{}, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *GameJson) SetPgn(v interface{})`

SetPgn sets Pgn field to given value.

### HasPgn

`func (o *GameJson) HasPgn() bool`

HasPgn returns a boolean if a field has been set.

### SetPgnNil

`func (o *GameJson) SetPgnNil(b bool)`

 SetPgnNil sets the value for Pgn to be an explicit nil

### UnsetPgn
`func (o *GameJson) UnsetPgn()`

UnsetPgn ensures that no value is present for Pgn, not even an explicit nil
### GetDaysPerTurn

`func (o *GameJson) GetDaysPerTurn() interface{}`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *GameJson) GetDaysPerTurnOk() (*interface{}, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *GameJson) SetDaysPerTurn(v interface{})`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *GameJson) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### SetDaysPerTurnNil

`func (o *GameJson) SetDaysPerTurnNil(b bool)`

 SetDaysPerTurnNil sets the value for DaysPerTurn to be an explicit nil

### UnsetDaysPerTurn
`func (o *GameJson) UnsetDaysPerTurn()`

UnsetDaysPerTurn ensures that no value is present for DaysPerTurn, not even an explicit nil
### GetAnalysis

`func (o *GameJson) GetAnalysis() interface{}`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GameJson) GetAnalysisOk() (*interface{}, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GameJson) SetAnalysis(v interface{})`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GameJson) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### SetAnalysisNil

`func (o *GameJson) SetAnalysisNil(b bool)`

 SetAnalysisNil sets the value for Analysis to be an explicit nil

### UnsetAnalysis
`func (o *GameJson) UnsetAnalysis()`

UnsetAnalysis ensures that no value is present for Analysis, not even an explicit nil
### GetTournament

`func (o *GameJson) GetTournament() interface{}`

GetTournament returns the Tournament field if non-nil, zero value otherwise.

### GetTournamentOk

`func (o *GameJson) GetTournamentOk() (*interface{}, bool)`

GetTournamentOk returns a tuple with the Tournament field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournament

`func (o *GameJson) SetTournament(v interface{})`

SetTournament sets Tournament field to given value.

### HasTournament

`func (o *GameJson) HasTournament() bool`

HasTournament returns a boolean if a field has been set.

### SetTournamentNil

`func (o *GameJson) SetTournamentNil(b bool)`

 SetTournamentNil sets the value for Tournament to be an explicit nil

### UnsetTournament
`func (o *GameJson) UnsetTournament()`

UnsetTournament ensures that no value is present for Tournament, not even an explicit nil
### GetSwiss

`func (o *GameJson) GetSwiss() interface{}`

GetSwiss returns the Swiss field if non-nil, zero value otherwise.

### GetSwissOk

`func (o *GameJson) GetSwissOk() (*interface{}, bool)`

GetSwissOk returns a tuple with the Swiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiss

`func (o *GameJson) SetSwiss(v interface{})`

SetSwiss sets Swiss field to given value.

### HasSwiss

`func (o *GameJson) HasSwiss() bool`

HasSwiss returns a boolean if a field has been set.

### SetSwissNil

`func (o *GameJson) SetSwissNil(b bool)`

 SetSwissNil sets the value for Swiss to be an explicit nil

### UnsetSwiss
`func (o *GameJson) UnsetSwiss()`

UnsetSwiss ensures that no value is present for Swiss, not even an explicit nil
### GetClock

`func (o *GameJson) GetClock() GameJsonClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GameJson) GetClockOk() (*GameJsonClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GameJson) SetClock(v GameJsonClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *GameJson) HasClock() bool`

HasClock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


