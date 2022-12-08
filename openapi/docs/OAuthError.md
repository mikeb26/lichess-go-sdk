# OAuthError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **interface{}** | The cause of the error. | [optional] 
**ErrorDescription** | Pointer to **interface{}** | The reason why the request was rejected. | [optional] 

## Methods

### NewOAuthError

`func NewOAuthError() *OAuthError`

NewOAuthError instantiates a new OAuthError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthErrorWithDefaults

`func NewOAuthErrorWithDefaults() *OAuthError`

NewOAuthErrorWithDefaults instantiates a new OAuthError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OAuthError) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OAuthError) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OAuthError) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *OAuthError) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *OAuthError) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *OAuthError) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetErrorDescription

`func (o *OAuthError) GetErrorDescription() interface{}`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *OAuthError) GetErrorDescriptionOk() (*interface{}, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *OAuthError) SetErrorDescription(v interface{})`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *OAuthError) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *OAuthError) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *OAuthError) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


