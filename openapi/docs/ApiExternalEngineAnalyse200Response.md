# ApiExternalEngineAnalyse200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **interface{}** | Number of milliseconds the search has been going on | 
**Depth** | **interface{}** | Current search depth | 
**Nodes** | **interface{}** | Number of nodes visited so far | 
**Pvs** | [**[]ApiExternalEngineAnalyse200ResponsePvsInner**](ApiExternalEngineAnalyse200ResponsePvsInner.md) | Information about up to 5 pvs, with the primary pv at index 0. | 

## Methods

### NewApiExternalEngineAnalyse200Response

`func NewApiExternalEngineAnalyse200Response(time interface{}, depth interface{}, nodes interface{}, pvs []ApiExternalEngineAnalyse200ResponsePvsInner, ) *ApiExternalEngineAnalyse200Response`

NewApiExternalEngineAnalyse200Response instantiates a new ApiExternalEngineAnalyse200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyse200ResponseWithDefaults

`func NewApiExternalEngineAnalyse200ResponseWithDefaults() *ApiExternalEngineAnalyse200Response`

NewApiExternalEngineAnalyse200ResponseWithDefaults instantiates a new ApiExternalEngineAnalyse200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ApiExternalEngineAnalyse200Response) GetTime() interface{}`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApiExternalEngineAnalyse200Response) GetTimeOk() (*interface{}, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApiExternalEngineAnalyse200Response) SetTime(v interface{})`

SetTime sets Time field to given value.


### SetTimeNil

`func (o *ApiExternalEngineAnalyse200Response) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *ApiExternalEngineAnalyse200Response) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetDepth

`func (o *ApiExternalEngineAnalyse200Response) GetDepth() interface{}`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiExternalEngineAnalyse200Response) GetDepthOk() (*interface{}, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiExternalEngineAnalyse200Response) SetDepth(v interface{})`

SetDepth sets Depth field to given value.


### SetDepthNil

`func (o *ApiExternalEngineAnalyse200Response) SetDepthNil(b bool)`

 SetDepthNil sets the value for Depth to be an explicit nil

### UnsetDepth
`func (o *ApiExternalEngineAnalyse200Response) UnsetDepth()`

UnsetDepth ensures that no value is present for Depth, not even an explicit nil
### GetNodes

`func (o *ApiExternalEngineAnalyse200Response) GetNodes() interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ApiExternalEngineAnalyse200Response) GetNodesOk() (*interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ApiExternalEngineAnalyse200Response) SetNodes(v interface{})`

SetNodes sets Nodes field to given value.


### SetNodesNil

`func (o *ApiExternalEngineAnalyse200Response) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *ApiExternalEngineAnalyse200Response) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetPvs

`func (o *ApiExternalEngineAnalyse200Response) GetPvs() []ApiExternalEngineAnalyse200ResponsePvsInner`

GetPvs returns the Pvs field if non-nil, zero value otherwise.

### GetPvsOk

`func (o *ApiExternalEngineAnalyse200Response) GetPvsOk() (*[]ApiExternalEngineAnalyse200ResponsePvsInner, bool)`

GetPvsOk returns a tuple with the Pvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvs

`func (o *ApiExternalEngineAnalyse200Response) SetPvs(v []ApiExternalEngineAnalyse200ResponsePvsInner)`

SetPvs sets Pvs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


