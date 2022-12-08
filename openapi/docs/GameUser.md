# GameUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 
**Rating** | Pointer to **interface{}** |  | [optional] 
**RatingDiff** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **interface{}** |  | [optional] 
**Provisional** | Pointer to **interface{}** |  | [optional] 
**AiLevel** | Pointer to **interface{}** |  | [optional] 
**Analysis** | Pointer to [**GameUserAnalysis**](GameUserAnalysis.md) |  | [optional] 
**Team** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGameUser

`func NewGameUser() *GameUser`

NewGameUser instantiates a new GameUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameUserWithDefaults

`func NewGameUserWithDefaults() *GameUser`

NewGameUserWithDefaults instantiates a new GameUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GameUser) GetUser() LightUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GameUser) GetUserOk() (*LightUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GameUser) SetUser(v LightUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GameUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRating

`func (o *GameUser) GetRating() interface{}`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GameUser) GetRatingOk() (*interface{}, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GameUser) SetRating(v interface{})`

SetRating sets Rating field to given value.

### HasRating

`func (o *GameUser) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *GameUser) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *GameUser) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetRatingDiff

`func (o *GameUser) GetRatingDiff() interface{}`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *GameUser) GetRatingDiffOk() (*interface{}, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *GameUser) SetRatingDiff(v interface{})`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *GameUser) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### SetRatingDiffNil

`func (o *GameUser) SetRatingDiffNil(b bool)`

 SetRatingDiffNil sets the value for RatingDiff to be an explicit nil

### UnsetRatingDiff
`func (o *GameUser) UnsetRatingDiff()`

UnsetRatingDiff ensures that no value is present for RatingDiff, not even an explicit nil
### GetName

`func (o *GameUser) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GameUser) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GameUser) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GameUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GameUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GameUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvisional

`func (o *GameUser) GetProvisional() interface{}`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *GameUser) GetProvisionalOk() (*interface{}, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *GameUser) SetProvisional(v interface{})`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *GameUser) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### SetProvisionalNil

`func (o *GameUser) SetProvisionalNil(b bool)`

 SetProvisionalNil sets the value for Provisional to be an explicit nil

### UnsetProvisional
`func (o *GameUser) UnsetProvisional()`

UnsetProvisional ensures that no value is present for Provisional, not even an explicit nil
### GetAiLevel

`func (o *GameUser) GetAiLevel() interface{}`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *GameUser) GetAiLevelOk() (*interface{}, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *GameUser) SetAiLevel(v interface{})`

SetAiLevel sets AiLevel field to given value.

### HasAiLevel

`func (o *GameUser) HasAiLevel() bool`

HasAiLevel returns a boolean if a field has been set.

### SetAiLevelNil

`func (o *GameUser) SetAiLevelNil(b bool)`

 SetAiLevelNil sets the value for AiLevel to be an explicit nil

### UnsetAiLevel
`func (o *GameUser) UnsetAiLevel()`

UnsetAiLevel ensures that no value is present for AiLevel, not even an explicit nil
### GetAnalysis

`func (o *GameUser) GetAnalysis() GameUserAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GameUser) GetAnalysisOk() (*GameUserAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GameUser) SetAnalysis(v GameUserAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GameUser) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetTeam

`func (o *GameUser) GetTeam() interface{}`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GameUser) GetTeamOk() (*interface{}, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GameUser) SetTeam(v interface{})`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GameUser) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *GameUser) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *GameUser) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


