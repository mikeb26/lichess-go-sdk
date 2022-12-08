# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **interface{}** |  | [optional] 
**Open** | Pointer to **interface{}** |  | [optional] 
**Leader** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 
**Leaders** | Pointer to **interface{}** |  | [optional] 
**NbMembers** | Pointer to **interface{}** |  | [optional] 
**Location** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTeam

`func NewTeam() *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Team) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Team) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Team) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *Team) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Team) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Team) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Team) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Team) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Team) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOpen

`func (o *Team) GetOpen() interface{}`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Team) GetOpenOk() (*interface{}, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Team) SetOpen(v interface{})`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Team) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### SetOpenNil

`func (o *Team) SetOpenNil(b bool)`

 SetOpenNil sets the value for Open to be an explicit nil

### UnsetOpen
`func (o *Team) UnsetOpen()`

UnsetOpen ensures that no value is present for Open, not even an explicit nil
### GetLeader

`func (o *Team) GetLeader() LightUser`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *Team) GetLeaderOk() (*LightUser, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *Team) SetLeader(v LightUser)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *Team) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetLeaders

`func (o *Team) GetLeaders() interface{}`

GetLeaders returns the Leaders field if non-nil, zero value otherwise.

### GetLeadersOk

`func (o *Team) GetLeadersOk() (*interface{}, bool)`

GetLeadersOk returns a tuple with the Leaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaders

`func (o *Team) SetLeaders(v interface{})`

SetLeaders sets Leaders field to given value.

### HasLeaders

`func (o *Team) HasLeaders() bool`

HasLeaders returns a boolean if a field has been set.

### SetLeadersNil

`func (o *Team) SetLeadersNil(b bool)`

 SetLeadersNil sets the value for Leaders to be an explicit nil

### UnsetLeaders
`func (o *Team) UnsetLeaders()`

UnsetLeaders ensures that no value is present for Leaders, not even an explicit nil
### GetNbMembers

`func (o *Team) GetNbMembers() interface{}`

GetNbMembers returns the NbMembers field if non-nil, zero value otherwise.

### GetNbMembersOk

`func (o *Team) GetNbMembersOk() (*interface{}, bool)`

GetNbMembersOk returns a tuple with the NbMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbMembers

`func (o *Team) SetNbMembers(v interface{})`

SetNbMembers sets NbMembers field to given value.

### HasNbMembers

`func (o *Team) HasNbMembers() bool`

HasNbMembers returns a boolean if a field has been set.

### SetNbMembersNil

`func (o *Team) SetNbMembersNil(b bool)`

 SetNbMembersNil sets the value for NbMembers to be an explicit nil

### UnsetNbMembers
`func (o *Team) UnsetNbMembers()`

UnsetNbMembers ensures that no value is present for NbMembers, not even an explicit nil
### GetLocation

`func (o *Team) GetLocation() interface{}`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Team) GetLocationOk() (*interface{}, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Team) SetLocation(v interface{})`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Team) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Team) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Team) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


