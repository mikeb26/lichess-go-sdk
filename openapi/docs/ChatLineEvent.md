# ChatLineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Room** | **interface{}** |  | 
**Username** | **interface{}** |  | 
**Text** | **interface{}** |  | 

## Methods

### NewChatLineEvent

`func NewChatLineEvent(type_ interface{}, room interface{}, username interface{}, text interface{}, ) *ChatLineEvent`

NewChatLineEvent instantiates a new ChatLineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatLineEventWithDefaults

`func NewChatLineEventWithDefaults() *ChatLineEvent`

NewChatLineEventWithDefaults instantiates a new ChatLineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatLineEvent) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatLineEvent) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatLineEvent) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ChatLineEvent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChatLineEvent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetRoom

`func (o *ChatLineEvent) GetRoom() interface{}`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *ChatLineEvent) GetRoomOk() (*interface{}, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *ChatLineEvent) SetRoom(v interface{})`

SetRoom sets Room field to given value.


### SetRoomNil

`func (o *ChatLineEvent) SetRoomNil(b bool)`

 SetRoomNil sets the value for Room to be an explicit nil

### UnsetRoom
`func (o *ChatLineEvent) UnsetRoom()`

UnsetRoom ensures that no value is present for Room, not even an explicit nil
### GetUsername

`func (o *ChatLineEvent) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ChatLineEvent) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ChatLineEvent) SetUsername(v interface{})`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *ChatLineEvent) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ChatLineEvent) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetText

`func (o *ChatLineEvent) GetText() interface{}`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatLineEvent) GetTextOk() (*interface{}, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatLineEvent) SetText(v interface{})`

SetText sets Text field to given value.


### SetTextNil

`func (o *ChatLineEvent) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ChatLineEvent) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


