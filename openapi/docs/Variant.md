# Variant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**VariantKey**](VariantKey.md) |  | [optional] [default to standard]
**Name** | Pointer to **interface{}** |  | [optional] 
**Short** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewVariant

`func NewVariant() *Variant`

NewVariant instantiates a new Variant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantWithDefaults

`func NewVariantWithDefaults() *Variant`

NewVariantWithDefaults instantiates a new Variant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Variant) GetKey() VariantKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Variant) GetKeyOk() (*VariantKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Variant) SetKey(v VariantKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *Variant) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *Variant) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variant) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variant) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *Variant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Variant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Variant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShort

`func (o *Variant) GetShort() interface{}`

GetShort returns the Short field if non-nil, zero value otherwise.

### GetShortOk

`func (o *Variant) GetShortOk() (*interface{}, bool)`

GetShortOk returns a tuple with the Short field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShort

`func (o *Variant) SetShort(v interface{})`

SetShort sets Short field to given value.

### HasShort

`func (o *Variant) HasShort() bool`

HasShort returns a boolean if a field has been set.

### SetShortNil

`func (o *Variant) SetShortNil(b bool)`

 SetShortNil sets the value for Short to be an explicit nil

### UnsetShort
`func (o *Variant) UnsetShort()`

UnsetShort ensures that no value is present for Short, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


