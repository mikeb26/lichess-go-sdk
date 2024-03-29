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

// GameEventInfo struct for GameEventInfo
type GameEventInfo struct {
	Id interface{} `json:"id,omitempty"`
	Source interface{} `json:"source,omitempty"`
	Compat *GameEventInfoCompat `json:"compat,omitempty"`
}

// NewGameEventInfo instantiates a new GameEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameEventInfo() *GameEventInfo {
	this := GameEventInfo{}
	return &this
}

// NewGameEventInfoWithDefaults instantiates a new GameEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameEventInfoWithDefaults() *GameEventInfo {
	this := GameEventInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventInfo) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventInfo) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GameEventInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GameEventInfo) SetId(v interface{}) {
	o.Id = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventInfo) GetSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventInfo) GetSourceOk() (*interface{}, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GameEventInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given interface{} and assigns it to the Source field.
func (o *GameEventInfo) SetSource(v interface{}) {
	o.Source = v
}

// GetCompat returns the Compat field value if set, zero value otherwise.
func (o *GameEventInfo) GetCompat() GameEventInfoCompat {
	if o == nil || o.Compat == nil {
		var ret GameEventInfoCompat
		return ret
	}
	return *o.Compat
}

// GetCompatOk returns a tuple with the Compat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetCompatOk() (*GameEventInfoCompat, bool) {
	if o == nil || o.Compat == nil {
		return nil, false
	}
	return o.Compat, true
}

// HasCompat returns a boolean if a field has been set.
func (o *GameEventInfo) HasCompat() bool {
	if o != nil && o.Compat != nil {
		return true
	}

	return false
}

// SetCompat gets a reference to the given GameEventInfoCompat and assigns it to the Compat field.
func (o *GameEventInfo) SetCompat(v GameEventInfoCompat) {
	o.Compat = &v
}

func (o GameEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Compat != nil {
		toSerialize["compat"] = o.Compat
	}
	return json.Marshal(toSerialize)
}

type NullableGameEventInfo struct {
	value *GameEventInfo
	isSet bool
}

func (v NullableGameEventInfo) Get() *GameEventInfo {
	return v.value
}

func (v *NullableGameEventInfo) Set(val *GameEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGameEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGameEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameEventInfo(val *GameEventInfo) *NullableGameEventInfo {
	return &NullableGameEventInfo{value: val, isSet: true}
}

func (v NullableGameEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


