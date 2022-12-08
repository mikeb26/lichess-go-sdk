# ExternalEngineRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | Display name of the engine. | 
**MaxThreads** | **interface{}** | Maximum number of available threads. | 
**MaxHash** | **interface{}** | Maximum available hash table size, in MiB. | 
**DefaultDepth** | **interface{}** | Estimated depth of normal search. | 
**Variants** | Pointer to **interface{}** | Optional list of supported chess variants. | [optional] 
**ProviderSecret** | **interface{}** | A random token that can be used to [wait for analysis requests](#tag/External-engine/operation/apiExternalEngineAcquire) and provide analysis.  The engine provider should securely generate a random string.  The token will not be readable again, even by the user.  The analysis provider can register multiple engines with the same token, even for different users, and wait for analysis requests from any of them. In this case, the request must not be made via CORS, so that the token is not revealed to any of the users.  | 
**ProviderData** | Pointer to **interface{}** | Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the &#x60;providerSecret&#x60;.  | [optional] 

## Methods

### NewExternalEngineRegistration

`func NewExternalEngineRegistration(name interface{}, maxThreads interface{}, maxHash interface{}, defaultDepth interface{}, providerSecret interface{}, ) *ExternalEngineRegistration`

NewExternalEngineRegistration instantiates a new ExternalEngineRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineRegistrationWithDefaults

`func NewExternalEngineRegistrationWithDefaults() *ExternalEngineRegistration`

NewExternalEngineRegistrationWithDefaults instantiates a new ExternalEngineRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalEngineRegistration) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalEngineRegistration) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalEngineRegistration) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *ExternalEngineRegistration) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExternalEngineRegistration) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMaxThreads

`func (o *ExternalEngineRegistration) GetMaxThreads() interface{}`

GetMaxThreads returns the MaxThreads field if non-nil, zero value otherwise.

### GetMaxThreadsOk

`func (o *ExternalEngineRegistration) GetMaxThreadsOk() (*interface{}, bool)`

GetMaxThreadsOk returns a tuple with the MaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreads

`func (o *ExternalEngineRegistration) SetMaxThreads(v interface{})`

SetMaxThreads sets MaxThreads field to given value.


### SetMaxThreadsNil

`func (o *ExternalEngineRegistration) SetMaxThreadsNil(b bool)`

 SetMaxThreadsNil sets the value for MaxThreads to be an explicit nil

### UnsetMaxThreads
`func (o *ExternalEngineRegistration) UnsetMaxThreads()`

UnsetMaxThreads ensures that no value is present for MaxThreads, not even an explicit nil
### GetMaxHash

`func (o *ExternalEngineRegistration) GetMaxHash() interface{}`

GetMaxHash returns the MaxHash field if non-nil, zero value otherwise.

### GetMaxHashOk

`func (o *ExternalEngineRegistration) GetMaxHashOk() (*interface{}, bool)`

GetMaxHashOk returns a tuple with the MaxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHash

`func (o *ExternalEngineRegistration) SetMaxHash(v interface{})`

SetMaxHash sets MaxHash field to given value.


### SetMaxHashNil

`func (o *ExternalEngineRegistration) SetMaxHashNil(b bool)`

 SetMaxHashNil sets the value for MaxHash to be an explicit nil

### UnsetMaxHash
`func (o *ExternalEngineRegistration) UnsetMaxHash()`

UnsetMaxHash ensures that no value is present for MaxHash, not even an explicit nil
### GetDefaultDepth

`func (o *ExternalEngineRegistration) GetDefaultDepth() interface{}`

GetDefaultDepth returns the DefaultDepth field if non-nil, zero value otherwise.

### GetDefaultDepthOk

`func (o *ExternalEngineRegistration) GetDefaultDepthOk() (*interface{}, bool)`

GetDefaultDepthOk returns a tuple with the DefaultDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepth

`func (o *ExternalEngineRegistration) SetDefaultDepth(v interface{})`

SetDefaultDepth sets DefaultDepth field to given value.


### SetDefaultDepthNil

`func (o *ExternalEngineRegistration) SetDefaultDepthNil(b bool)`

 SetDefaultDepthNil sets the value for DefaultDepth to be an explicit nil

### UnsetDefaultDepth
`func (o *ExternalEngineRegistration) UnsetDefaultDepth()`

UnsetDefaultDepth ensures that no value is present for DefaultDepth, not even an explicit nil
### GetVariants

`func (o *ExternalEngineRegistration) GetVariants() interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExternalEngineRegistration) GetVariantsOk() (*interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExternalEngineRegistration) SetVariants(v interface{})`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ExternalEngineRegistration) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *ExternalEngineRegistration) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *ExternalEngineRegistration) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil
### GetProviderSecret

`func (o *ExternalEngineRegistration) GetProviderSecret() interface{}`

GetProviderSecret returns the ProviderSecret field if non-nil, zero value otherwise.

### GetProviderSecretOk

`func (o *ExternalEngineRegistration) GetProviderSecretOk() (*interface{}, bool)`

GetProviderSecretOk returns a tuple with the ProviderSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSecret

`func (o *ExternalEngineRegistration) SetProviderSecret(v interface{})`

SetProviderSecret sets ProviderSecret field to given value.


### SetProviderSecretNil

`func (o *ExternalEngineRegistration) SetProviderSecretNil(b bool)`

 SetProviderSecretNil sets the value for ProviderSecret to be an explicit nil

### UnsetProviderSecret
`func (o *ExternalEngineRegistration) UnsetProviderSecret()`

UnsetProviderSecret ensures that no value is present for ProviderSecret, not even an explicit nil
### GetProviderData

`func (o *ExternalEngineRegistration) GetProviderData() interface{}`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *ExternalEngineRegistration) GetProviderDataOk() (*interface{}, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *ExternalEngineRegistration) SetProviderData(v interface{})`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *ExternalEngineRegistration) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### SetProviderDataNil

`func (o *ExternalEngineRegistration) SetProviderDataNil(b bool)`

 SetProviderDataNil sets the value for ProviderData to be an explicit nil

### UnsetProviderData
`func (o *ExternalEngineRegistration) UnsetProviderData()`

UnsetProviderData ensures that no value is present for ProviderData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


