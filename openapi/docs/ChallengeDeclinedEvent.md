# ChallengeDeclinedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** |  | [optional] 
**Challenge** | Pointer to [**ChallengeCanceledJson**](ChallengeCanceledJson.md) |  | [optional] 

## Methods

### NewChallengeDeclinedEvent

`func NewChallengeDeclinedEvent() *ChallengeDeclinedEvent`

NewChallengeDeclinedEvent instantiates a new ChallengeDeclinedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeDeclinedEventWithDefaults

`func NewChallengeDeclinedEventWithDefaults() *ChallengeDeclinedEvent`

NewChallengeDeclinedEventWithDefaults instantiates a new ChallengeDeclinedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChallengeDeclinedEvent) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeDeclinedEvent) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeDeclinedEvent) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *ChallengeDeclinedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ChallengeDeclinedEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChallengeDeclinedEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetChallenge

`func (o *ChallengeDeclinedEvent) GetChallenge() ChallengeCanceledJson`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ChallengeDeclinedEvent) GetChallengeOk() (*ChallengeCanceledJson, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ChallengeDeclinedEvent) SetChallenge(v ChallengeCanceledJson)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *ChallengeDeclinedEvent) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


