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

// GameFinishEvent struct for GameFinishEvent
type GameFinishEvent struct {
	Type interface{} `json:"type,omitempty"`
	Game *GameEventInfo `json:"game,omitempty"`
}

// NewGameFinishEvent instantiates a new GameFinishEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameFinishEvent() *GameFinishEvent {
	this := GameFinishEvent{}
	return &this
}

// NewGameFinishEventWithDefaults instantiates a new GameFinishEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameFinishEventWithDefaults() *GameFinishEvent {
	this := GameFinishEvent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameFinishEvent) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFinishEvent) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GameFinishEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GameFinishEvent) SetType(v interface{}) {
	o.Type = v
}

// GetGame returns the Game field value if set, zero value otherwise.
func (o *GameFinishEvent) GetGame() GameEventInfo {
	if o == nil || o.Game == nil {
		var ret GameEventInfo
		return ret
	}
	return *o.Game
}

// GetGameOk returns a tuple with the Game field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameFinishEvent) GetGameOk() (*GameEventInfo, bool) {
	if o == nil || o.Game == nil {
		return nil, false
	}
	return o.Game, true
}

// HasGame returns a boolean if a field has been set.
func (o *GameFinishEvent) HasGame() bool {
	if o != nil && o.Game != nil {
		return true
	}

	return false
}

// SetGame gets a reference to the given GameEventInfo and assigns it to the Game field.
func (o *GameFinishEvent) SetGame(v GameEventInfo) {
	o.Game = &v
}

func (o GameFinishEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Game != nil {
		toSerialize["game"] = o.Game
	}
	return json.Marshal(toSerialize)
}

type NullableGameFinishEvent struct {
	value *GameFinishEvent
	isSet bool
}

func (v NullableGameFinishEvent) Get() *GameFinishEvent {
	return v.value
}

func (v *NullableGameFinishEvent) Set(val *GameFinishEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableGameFinishEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableGameFinishEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameFinishEvent(val *GameFinishEvent) *NullableGameFinishEvent {
	return &NullableGameFinishEvent{value: val, isSet: true}
}

func (v NullableGameFinishEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameFinishEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

