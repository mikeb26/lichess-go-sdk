# TeamRequestWithUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**TeamRequest**](TeamRequest.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewTeamRequestWithUser

`func NewTeamRequestWithUser() *TeamRequestWithUser`

NewTeamRequestWithUser instantiates a new TeamRequestWithUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRequestWithUserWithDefaults

`func NewTeamRequestWithUserWithDefaults() *TeamRequestWithUser`

NewTeamRequestWithUserWithDefaults instantiates a new TeamRequestWithUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *TeamRequestWithUser) GetRequest() TeamRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *TeamRequestWithUser) GetRequestOk() (*TeamRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *TeamRequestWithUser) SetRequest(v TeamRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *TeamRequestWithUser) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetUser

`func (o *TeamRequestWithUser) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamRequestWithUser) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamRequestWithUser) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *TeamRequestWithUser) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


