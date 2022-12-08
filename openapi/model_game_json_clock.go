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

// GameJsonClock struct for GameJsonClock
type GameJsonClock struct {
	Initial interface{} `json:"initial,omitempty"`
	Increment interface{} `json:"increment,omitempty"`
	TotalTime interface{} `json:"totalTime,omitempty"`
}

// NewGameJsonClock instantiates a new GameJsonClock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameJsonClock() *GameJsonClock {
	this := GameJsonClock{}
	return &this
}

// NewGameJsonClockWithDefaults instantiates a new GameJsonClock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameJsonClockWithDefaults() *GameJsonClock {
	this := GameJsonClock{}
	return &this
}

// GetInitial returns the Initial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameJsonClock) GetInitial() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Initial
}

// GetInitialOk returns a tuple with the Initial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameJsonClock) GetInitialOk() (*interface{}, bool) {
	if o == nil || o.Initial == nil {
		return nil, false
	}
	return &o.Initial, true
}

// HasInitial returns a boolean if a field has been set.
func (o *GameJsonClock) HasInitial() bool {
	if o != nil && o.Initial != nil {
		return true
	}

	return false
}

// SetInitial gets a reference to the given interface{} and assigns it to the Initial field.
func (o *GameJsonClock) SetInitial(v interface{}) {
	o.Initial = v
}

// GetIncrement returns the Increment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameJsonClock) GetIncrement() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameJsonClock) GetIncrementOk() (*interface{}, bool) {
	if o == nil || o.Increment == nil {
		return nil, false
	}
	return &o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *GameJsonClock) HasIncrement() bool {
	if o != nil && o.Increment != nil {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given interface{} and assigns it to the Increment field.
func (o *GameJsonClock) SetIncrement(v interface{}) {
	o.Increment = v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameJsonClock) GetTotalTime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameJsonClock) GetTotalTimeOk() (*interface{}, bool) {
	if o == nil || o.TotalTime == nil {
		return nil, false
	}
	return &o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *GameJsonClock) HasTotalTime() bool {
	if o != nil && o.TotalTime != nil {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given interface{} and assigns it to the TotalTime field.
func (o *GameJsonClock) SetTotalTime(v interface{}) {
	o.TotalTime = v
}

func (o GameJsonClock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Initial != nil {
		toSerialize["initial"] = o.Initial
	}
	if o.Increment != nil {
		toSerialize["increment"] = o.Increment
	}
	if o.TotalTime != nil {
		toSerialize["totalTime"] = o.TotalTime
	}
	return json.Marshal(toSerialize)
}

type NullableGameJsonClock struct {
	value *GameJsonClock
	isSet bool
}

func (v NullableGameJsonClock) Get() *GameJsonClock {
	return v.value
}

func (v *NullableGameJsonClock) Set(val *GameJsonClock) {
	v.value = val
	v.isSet = true
}

func (v NullableGameJsonClock) IsSet() bool {
	return v.isSet
}

func (v *NullableGameJsonClock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameJsonClock(val *GameJsonClock) *NullableGameJsonClock {
	return &NullableGameJsonClock{value: val, isSet: true}
}

func (v NullableGameJsonClock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameJsonClock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


