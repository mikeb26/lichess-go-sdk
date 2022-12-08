/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login, gameplay, and more: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/ZackClements/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/ShailChoksi/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), please wait a full minute before resuming API usage.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](http://ndjson.org/), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses. 

API version: 2.0.0
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiUsersStatus200ResponseInner struct for ApiUsersStatus200ResponseInner
type ApiUsersStatus200ResponseInner struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Title interface{} `json:"title,omitempty"`
	Online interface{} `json:"online,omitempty"`
	Playing interface{} `json:"playing,omitempty"`
	Streaming interface{} `json:"streaming,omitempty"`
	Patron interface{} `json:"patron,omitempty"`
}

// NewApiUsersStatus200ResponseInner instantiates a new ApiUsersStatus200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUsersStatus200ResponseInner(id interface{}, name interface{}) *ApiUsersStatus200ResponseInner {
	this := ApiUsersStatus200ResponseInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewApiUsersStatus200ResponseInnerWithDefaults instantiates a new ApiUsersStatus200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUsersStatus200ResponseInnerWithDefaults() *ApiUsersStatus200ResponseInner {
	this := ApiUsersStatus200ResponseInner{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApiUsersStatus200ResponseInner) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiUsersStatus200ResponseInner) SetId(v interface{}) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApiUsersStatus200ResponseInner) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiUsersStatus200ResponseInner) SetName(v interface{}) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiUsersStatus200ResponseInner) GetTitle() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetTitleOk() (*interface{}, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given interface{} and assigns it to the Title field.
func (o *ApiUsersStatus200ResponseInner) SetTitle(v interface{}) {
	o.Title = v
}

// GetOnline returns the Online field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiUsersStatus200ResponseInner) GetOnline() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetOnlineOk() (*interface{}, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return &o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given interface{} and assigns it to the Online field.
func (o *ApiUsersStatus200ResponseInner) SetOnline(v interface{}) {
	o.Online = v
}

// GetPlaying returns the Playing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiUsersStatus200ResponseInner) GetPlaying() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Playing
}

// GetPlayingOk returns a tuple with the Playing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetPlayingOk() (*interface{}, bool) {
	if o == nil || o.Playing == nil {
		return nil, false
	}
	return &o.Playing, true
}

// HasPlaying returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasPlaying() bool {
	if o != nil && o.Playing != nil {
		return true
	}

	return false
}

// SetPlaying gets a reference to the given interface{} and assigns it to the Playing field.
func (o *ApiUsersStatus200ResponseInner) SetPlaying(v interface{}) {
	o.Playing = v
}

// GetStreaming returns the Streaming field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiUsersStatus200ResponseInner) GetStreaming() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Streaming
}

// GetStreamingOk returns a tuple with the Streaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetStreamingOk() (*interface{}, bool) {
	if o == nil || o.Streaming == nil {
		return nil, false
	}
	return &o.Streaming, true
}

// HasStreaming returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasStreaming() bool {
	if o != nil && o.Streaming != nil {
		return true
	}

	return false
}

// SetStreaming gets a reference to the given interface{} and assigns it to the Streaming field.
func (o *ApiUsersStatus200ResponseInner) SetStreaming(v interface{}) {
	o.Streaming = v
}

// GetPatron returns the Patron field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiUsersStatus200ResponseInner) GetPatron() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Patron
}

// GetPatronOk returns a tuple with the Patron field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiUsersStatus200ResponseInner) GetPatronOk() (*interface{}, bool) {
	if o == nil || o.Patron == nil {
		return nil, false
	}
	return &o.Patron, true
}

// HasPatron returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasPatron() bool {
	if o != nil && o.Patron != nil {
		return true
	}

	return false
}

// SetPatron gets a reference to the given interface{} and assigns it to the Patron field.
func (o *ApiUsersStatus200ResponseInner) SetPatron(v interface{}) {
	o.Patron = v
}

func (o ApiUsersStatus200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.Playing != nil {
		toSerialize["playing"] = o.Playing
	}
	if o.Streaming != nil {
		toSerialize["streaming"] = o.Streaming
	}
	if o.Patron != nil {
		toSerialize["patron"] = o.Patron
	}
	return json.Marshal(toSerialize)
}

type NullableApiUsersStatus200ResponseInner struct {
	value *ApiUsersStatus200ResponseInner
	isSet bool
}

func (v NullableApiUsersStatus200ResponseInner) Get() *ApiUsersStatus200ResponseInner {
	return v.value
}

func (v *NullableApiUsersStatus200ResponseInner) Set(val *ApiUsersStatus200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUsersStatus200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUsersStatus200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUsersStatus200ResponseInner(val *ApiUsersStatus200ResponseInner) *NullableApiUsersStatus200ResponseInner {
	return &NullableApiUsersStatus200ResponseInner{value: val, isSet: true}
}

func (v NullableApiUsersStatus200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUsersStatus200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

