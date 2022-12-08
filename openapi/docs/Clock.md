# Clock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **interface{}** |  | [optional] 
**Increment** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewClock

`func NewClock() *Clock`

NewClock instantiates a new Clock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClockWithDefaults

`func NewClockWithDefaults() *Clock`

NewClockWithDefaults instantiates a new Clock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Clock) GetLimit() interface{}`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Clock) GetLimitOk() (*interface{}, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Clock) SetLimit(v interface{})`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Clock) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Clock) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Clock) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetIncrement

`func (o *Clock) GetIncrement() interface{}`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *Clock) GetIncrementOk() (*interface{}, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *Clock) SetIncrement(v interface{})`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *Clock) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.

### SetIncrementNil

`func (o *Clock) SetIncrementNil(b bool)`

 SetIncrementNil sets the value for Increment to be an explicit nil

### UnsetIncrement
`func (o *Clock) UnsetIncrement()`

UnsetIncrement ensures that no value is present for Increment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


