# ChallengeJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** |  | 
**Url** | **interface{}** |  | 
**Status** | **interface{}** |  | 
**Challenger** | [**ChallengeUser**](ChallengeUser.md) |  | 
**DestUser** | **interface{}** |  | 
**Variant** | [**Variant**](Variant.md) |  | 
**Rated** | **interface{}** |  | 
**Speed** | [**Speed**](Speed.md) |  | 
**TimeControl** | **interface{}** |  | 
**Color** | **interface{}** |  | 
**Perf** | [**ChallengeJsonPerf**](ChallengeJsonPerf.md) |  | 
**Direction** | Pointer to **interface{}** |  | [optional] 
**InitialFen** | Pointer to **interface{}** |  | [optional] 
**DeclineReason** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewChallengeJson

`func NewChallengeJson(id interface{}, url interface{}, status interface{}, challenger ChallengeUser, destUser interface{}, variant Variant, rated interface{}, speed Speed, timeControl interface{}, color interface{}, perf ChallengeJsonPerf, ) *ChallengeJson`

NewChallengeJson instantiates a new ChallengeJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeJsonWithDefaults

`func NewChallengeJsonWithDefaults() *ChallengeJson`

NewChallengeJsonWithDefaults instantiates a new ChallengeJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeJson) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeJson) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeJson) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ChallengeJson) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ChallengeJson) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUrl

`func (o *ChallengeJson) GetUrl() interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeJson) GetUrlOk() (*interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeJson) SetUrl(v interface{})`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *ChallengeJson) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChallengeJson) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetStatus

`func (o *ChallengeJson) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeJson) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeJson) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ChallengeJson) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ChallengeJson) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetChallenger

`func (o *ChallengeJson) GetChallenger() ChallengeUser`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ChallengeJson) GetChallengerOk() (*ChallengeUser, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ChallengeJson) SetChallenger(v ChallengeUser)`

SetChallenger sets Challenger field to given value.


### GetDestUser

`func (o *ChallengeJson) GetDestUser() interface{}`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ChallengeJson) GetDestUserOk() (*interface{}, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ChallengeJson) SetDestUser(v interface{})`

SetDestUser sets DestUser field to given value.


### SetDestUserNil

`func (o *ChallengeJson) SetDestUserNil(b bool)`

 SetDestUserNil sets the value for DestUser to be an explicit nil

### UnsetDestUser
`func (o *ChallengeJson) UnsetDestUser()`

UnsetDestUser ensures that no value is present for DestUser, not even an explicit nil
### GetVariant

`func (o *ChallengeJson) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeJson) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeJson) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ChallengeJson) GetRated() interface{}`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeJson) GetRatedOk() (*interface{}, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeJson) SetRated(v interface{})`

SetRated sets Rated field to given value.


### SetRatedNil

`func (o *ChallengeJson) SetRatedNil(b bool)`

 SetRatedNil sets the value for Rated to be an explicit nil

### UnsetRated
`func (o *ChallengeJson) UnsetRated()`

UnsetRated ensures that no value is present for Rated, not even an explicit nil
### GetSpeed

`func (o *ChallengeJson) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeJson) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeJson) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ChallengeJson) GetTimeControl() interface{}`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ChallengeJson) GetTimeControlOk() (*interface{}, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ChallengeJson) SetTimeControl(v interface{})`

SetTimeControl sets TimeControl field to given value.


### SetTimeControlNil

`func (o *ChallengeJson) SetTimeControlNil(b bool)`

 SetTimeControlNil sets the value for TimeControl to be an explicit nil

### UnsetTimeControl
`func (o *ChallengeJson) UnsetTimeControl()`

UnsetTimeControl ensures that no value is present for TimeControl, not even an explicit nil
### GetColor

`func (o *ChallengeJson) GetColor() interface{}`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChallengeJson) GetColorOk() (*interface{}, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChallengeJson) SetColor(v interface{})`

SetColor sets Color field to given value.


### SetColorNil

`func (o *ChallengeJson) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ChallengeJson) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetPerf

`func (o *ChallengeJson) GetPerf() ChallengeJsonPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeJson) GetPerfOk() (*ChallengeJsonPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeJson) SetPerf(v ChallengeJsonPerf)`

SetPerf sets Perf field to given value.


### GetDirection

`func (o *ChallengeJson) GetDirection() interface{}`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ChallengeJson) GetDirectionOk() (*interface{}, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ChallengeJson) SetDirection(v interface{})`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ChallengeJson) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *ChallengeJson) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *ChallengeJson) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetInitialFen

`func (o *ChallengeJson) GetInitialFen() interface{}`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ChallengeJson) GetInitialFenOk() (*interface{}, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ChallengeJson) SetInitialFen(v interface{})`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ChallengeJson) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### SetInitialFenNil

`func (o *ChallengeJson) SetInitialFenNil(b bool)`

 SetInitialFenNil sets the value for InitialFen to be an explicit nil

### UnsetInitialFen
`func (o *ChallengeJson) UnsetInitialFen()`

UnsetInitialFen ensures that no value is present for InitialFen, not even an explicit nil
### GetDeclineReason

`func (o *ChallengeJson) GetDeclineReason() interface{}`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *ChallengeJson) GetDeclineReasonOk() (*interface{}, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *ChallengeJson) SetDeclineReason(v interface{})`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *ChallengeJson) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### SetDeclineReasonNil

`func (o *ChallengeJson) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *ChallengeJson) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


