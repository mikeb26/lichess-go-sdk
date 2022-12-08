# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** |  | [optional] 
**Username** | Pointer to **interface{}** |  | [optional] 
**Perfs** | Pointer to [**Perfs**](Perfs.md) |  | [optional] 
**CreatedAt** | Pointer to **interface{}** |  | [optional] 
**Disabled** | Pointer to **interface{}** |  | [optional] 
**TosViolation** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 
**SeenAt** | Pointer to **interface{}** |  | [optional] 
**Patron** | Pointer to **interface{}** |  | [optional] 
**Verified** | Pointer to **interface{}** |  | [optional] 
**PlayTime** | Pointer to [**PlayTime**](PlayTime.md) |  | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *User) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *User) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUsername

`func (o *User) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v interface{})`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *User) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *User) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPerfs

`func (o *User) GetPerfs() Perfs`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *User) GetPerfsOk() (*Perfs, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *User) SetPerfs(v Perfs)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *User) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *User) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *User) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDisabled

`func (o *User) GetDisabled() interface{}`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *User) GetDisabledOk() (*interface{}, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *User) SetDisabled(v interface{})`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *User) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *User) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *User) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetTosViolation

`func (o *User) GetTosViolation() interface{}`

GetTosViolation returns the TosViolation field if non-nil, zero value otherwise.

### GetTosViolationOk

`func (o *User) GetTosViolationOk() (*interface{}, bool)`

GetTosViolationOk returns a tuple with the TosViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosViolation

`func (o *User) SetTosViolation(v interface{})`

SetTosViolation sets TosViolation field to given value.

### HasTosViolation

`func (o *User) HasTosViolation() bool`

HasTosViolation returns a boolean if a field has been set.

### SetTosViolationNil

`func (o *User) SetTosViolationNil(b bool)`

 SetTosViolationNil sets the value for TosViolation to be an explicit nil

### UnsetTosViolation
`func (o *User) UnsetTosViolation()`

UnsetTosViolation ensures that no value is present for TosViolation, not even an explicit nil
### GetProfile

`func (o *User) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *User) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *User) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *User) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeenAt

`func (o *User) GetSeenAt() interface{}`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *User) GetSeenAtOk() (*interface{}, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *User) SetSeenAt(v interface{})`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *User) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### SetSeenAtNil

`func (o *User) SetSeenAtNil(b bool)`

 SetSeenAtNil sets the value for SeenAt to be an explicit nil

### UnsetSeenAt
`func (o *User) UnsetSeenAt()`

UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
### GetPatron

`func (o *User) GetPatron() interface{}`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *User) GetPatronOk() (*interface{}, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *User) SetPatron(v interface{})`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *User) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### SetPatronNil

`func (o *User) SetPatronNil(b bool)`

 SetPatronNil sets the value for Patron to be an explicit nil

### UnsetPatron
`func (o *User) UnsetPatron()`

UnsetPatron ensures that no value is present for Patron, not even an explicit nil
### GetVerified

`func (o *User) GetVerified() interface{}`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (*interface{}, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *User) SetVerified(v interface{})`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *User) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerifiedNil

`func (o *User) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *User) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil
### GetPlayTime

`func (o *User) GetPlayTime() PlayTime`

GetPlayTime returns the PlayTime field if non-nil, zero value otherwise.

### GetPlayTimeOk

`func (o *User) GetPlayTimeOk() (*PlayTime, bool)`

GetPlayTimeOk returns a tuple with the PlayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayTime

`func (o *User) SetPlayTime(v PlayTime)`

SetPlayTime sets PlayTime field to given value.

### HasPlayTime

`func (o *User) HasPlayTime() bool`

HasPlayTime returns a boolean if a field has been set.

### GetTitle

`func (o *User) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *User) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *User) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *User) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


