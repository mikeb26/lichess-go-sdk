# ApiExternalEngineAnalyseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **interface{}** |  | 
**Work** | [**ExternalEngineWork**](ExternalEngineWork.md) |  | 

## Methods

### NewApiExternalEngineAnalyseRequest

`func NewApiExternalEngineAnalyseRequest(clientSecret interface{}, work ExternalEngineWork, ) *ApiExternalEngineAnalyseRequest`

NewApiExternalEngineAnalyseRequest instantiates a new ApiExternalEngineAnalyseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyseRequestWithDefaults

`func NewApiExternalEngineAnalyseRequestWithDefaults() *ApiExternalEngineAnalyseRequest`

NewApiExternalEngineAnalyseRequestWithDefaults instantiates a new ApiExternalEngineAnalyseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *ApiExternalEngineAnalyseRequest) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApiExternalEngineAnalyseRequest) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApiExternalEngineAnalyseRequest) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.


### SetClientSecretNil

`func (o *ApiExternalEngineAnalyseRequest) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *ApiExternalEngineAnalyseRequest) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetWork

`func (o *ApiExternalEngineAnalyseRequest) GetWork() ExternalEngineWork`

GetWork returns the Work field if non-nil, zero value otherwise.

### GetWorkOk

`func (o *ApiExternalEngineAnalyseRequest) GetWorkOk() (*ExternalEngineWork, bool)`

GetWorkOk returns a tuple with the Work field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWork

`func (o *ApiExternalEngineAnalyseRequest) SetWork(v ExternalEngineWork)`

SetWork sets Work field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


