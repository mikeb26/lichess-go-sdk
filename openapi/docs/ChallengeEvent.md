# ChallengeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** |  | [optional] 
**Challenge** | Pointer to [**ChallengeJson**](ChallengeJson.md) |  | [optional] 

## Methods

### NewChallengeEvent

`func NewChallengeEvent() *ChallengeEvent`

NewChallengeEvent instantiates a new ChallengeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeEventWithDefaults

`func NewChallengeEventWithDefaults() *ChallengeEvent`

NewChallengeEventWithDefaults instantiates a new ChallengeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChallengeEvent) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeEvent) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeEvent) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *ChallengeEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ChallengeEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChallengeEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetChallenge

`func (o *ChallengeEvent) GetChallenge() ChallengeJson`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ChallengeEvent) GetChallengeOk() (*ChallengeJson, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ChallengeEvent) SetChallenge(v ChallengeJson)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *ChallengeEvent) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


