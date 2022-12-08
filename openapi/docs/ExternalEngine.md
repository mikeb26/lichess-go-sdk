# ExternalEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | Unique engine registration ID. | 
**Name** | **interface{}** | Display name of the engine. | 
**ClientSecret** | **interface{}** | A secret token that can be used to [*request* analysis](#tag/External-engine/operation/apiExternalEngineAnalyse) from this external engine.  | 
**UserId** | **interface{}** | The user this engine has been registered for. | 
**MaxThreads** | **interface{}** | Maximum number of available threads. | 
**MaxHash** | **interface{}** | Maximum available hash table size, in MiB. | 
**DefaultDepth** | **interface{}** | Estimated depth of normal search. | 
**Variants** | **interface{}** | List of supported chess variants. | 
**ProviderData** | Pointer to **interface{}** | Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the &#x60;providerSecret&#x60;.  | [optional] 

## Methods

### NewExternalEngine

`func NewExternalEngine(id interface{}, name interface{}, clientSecret interface{}, userId interface{}, maxThreads interface{}, maxHash interface{}, defaultDepth interface{}, variants interface{}, ) *ExternalEngine`

NewExternalEngine instantiates a new ExternalEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineWithDefaults

`func NewExternalEngineWithDefaults() *ExternalEngine`

NewExternalEngineWithDefaults instantiates a new ExternalEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalEngine) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalEngine) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalEngine) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExternalEngine) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalEngine) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ExternalEngine) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalEngine) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalEngine) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *ExternalEngine) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExternalEngine) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClientSecret

`func (o *ExternalEngine) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ExternalEngine) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ExternalEngine) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.


### SetClientSecretNil

`func (o *ExternalEngine) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *ExternalEngine) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetUserId

`func (o *ExternalEngine) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalEngine) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalEngine) SetUserId(v interface{})`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *ExternalEngine) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExternalEngine) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMaxThreads

`func (o *ExternalEngine) GetMaxThreads() interface{}`

GetMaxThreads returns the MaxThreads field if non-nil, zero value otherwise.

### GetMaxThreadsOk

`func (o *ExternalEngine) GetMaxThreadsOk() (*interface{}, bool)`

GetMaxThreadsOk returns a tuple with the MaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreads

`func (o *ExternalEngine) SetMaxThreads(v interface{})`

SetMaxThreads sets MaxThreads field to given value.


### SetMaxThreadsNil

`func (o *ExternalEngine) SetMaxThreadsNil(b bool)`

 SetMaxThreadsNil sets the value for MaxThreads to be an explicit nil

### UnsetMaxThreads
`func (o *ExternalEngine) UnsetMaxThreads()`

UnsetMaxThreads ensures that no value is present for MaxThreads, not even an explicit nil
### GetMaxHash

`func (o *ExternalEngine) GetMaxHash() interface{}`

GetMaxHash returns the MaxHash field if non-nil, zero value otherwise.

### GetMaxHashOk

`func (o *ExternalEngine) GetMaxHashOk() (*interface{}, bool)`

GetMaxHashOk returns a tuple with the MaxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHash

`func (o *ExternalEngine) SetMaxHash(v interface{})`

SetMaxHash sets MaxHash field to given value.


### SetMaxHashNil

`func (o *ExternalEngine) SetMaxHashNil(b bool)`

 SetMaxHashNil sets the value for MaxHash to be an explicit nil

### UnsetMaxHash
`func (o *ExternalEngine) UnsetMaxHash()`

UnsetMaxHash ensures that no value is present for MaxHash, not even an explicit nil
### GetDefaultDepth

`func (o *ExternalEngine) GetDefaultDepth() interface{}`

GetDefaultDepth returns the DefaultDepth field if non-nil, zero value otherwise.

### GetDefaultDepthOk

`func (o *ExternalEngine) GetDefaultDepthOk() (*interface{}, bool)`

GetDefaultDepthOk returns a tuple with the DefaultDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepth

`func (o *ExternalEngine) SetDefaultDepth(v interface{})`

SetDefaultDepth sets DefaultDepth field to given value.


### SetDefaultDepthNil

`func (o *ExternalEngine) SetDefaultDepthNil(b bool)`

 SetDefaultDepthNil sets the value for DefaultDepth to be an explicit nil

### UnsetDefaultDepth
`func (o *ExternalEngine) UnsetDefaultDepth()`

UnsetDefaultDepth ensures that no value is present for DefaultDepth, not even an explicit nil
### GetVariants

`func (o *ExternalEngine) GetVariants() interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExternalEngine) GetVariantsOk() (*interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExternalEngine) SetVariants(v interface{})`

SetVariants sets Variants field to given value.


### SetVariantsNil

`func (o *ExternalEngine) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *ExternalEngine) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil
### GetProviderData

`func (o *ExternalEngine) GetProviderData() interface{}`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *ExternalEngine) GetProviderDataOk() (*interface{}, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *ExternalEngine) SetProviderData(v interface{})`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *ExternalEngine) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### SetProviderDataNil

`func (o *ExternalEngine) SetProviderDataNil(b bool)`

 SetProviderDataNil sets the value for ProviderData to be an explicit nil

### UnsetProviderData
`func (o *ExternalEngine) UnsetProviderData()`

UnsetProviderData ensures that no value is present for ProviderData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


