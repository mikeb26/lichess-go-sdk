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

// GameJsonOpening struct for GameJsonOpening
type GameJsonOpening struct {
	Eco interface{} `json:"eco,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Ply interface{} `json:"ply,omitempty"`
}

// NewGameJsonOpening instantiates a new GameJsonOpening object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameJsonOpening() *GameJsonOpening {
	this := GameJsonOpening{}
	return &this
}

// NewGameJsonOpeningWithDefaults instantiates a new GameJsonOpening object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameJsonOpeningWithDefaults() *GameJsonOpening {
	this := GameJsonOpening{}
	return &this
}

// GetEco returns the Eco field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameJsonOpening) GetEco() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Eco
}

// GetEcoOk returns a tuple with the Eco field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameJsonOpening) GetEcoOk() (*interface{}, bool) {
	if o == nil || o.Eco == nil {
		return nil, false
	}
	return &o.Eco, true
}

// HasEco returns a boolean if a field has been set.
func (o *GameJsonOpening) HasEco() bool {
	if o != nil && o.Eco != nil {
		return true
	}

	return false
}

// SetEco gets a reference to the given interface{} and assigns it to the Eco field.
func (o *GameJsonOpening) SetEco(v interface{}) {
	o.Eco = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameJsonOpening) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameJsonOpening) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GameJsonOpening) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GameJsonOpening) SetName(v interface{}) {
	o.Name = v
}

// GetPly returns the Ply field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameJsonOpening) GetPly() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Ply
}

// GetPlyOk returns a tuple with the Ply field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameJsonOpening) GetPlyOk() (*interface{}, bool) {
	if o == nil || o.Ply == nil {
		return nil, false
	}
	return &o.Ply, true
}

// HasPly returns a boolean if a field has been set.
func (o *GameJsonOpening) HasPly() bool {
	if o != nil && o.Ply != nil {
		return true
	}

	return false
}

// SetPly gets a reference to the given interface{} and assigns it to the Ply field.
func (o *GameJsonOpening) SetPly(v interface{}) {
	o.Ply = v
}

func (o GameJsonOpening) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Eco != nil {
		toSerialize["eco"] = o.Eco
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ply != nil {
		toSerialize["ply"] = o.Ply
	}
	return json.Marshal(toSerialize)
}

type NullableGameJsonOpening struct {
	value *GameJsonOpening
	isSet bool
}

func (v NullableGameJsonOpening) Get() *GameJsonOpening {
	return v.value
}

func (v *NullableGameJsonOpening) Set(val *GameJsonOpening) {
	v.value = val
	v.isSet = true
}

func (v NullableGameJsonOpening) IsSet() bool {
	return v.isSet
}

func (v *NullableGameJsonOpening) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameJsonOpening(val *GameJsonOpening) *NullableGameJsonOpening {
	return &NullableGameJsonOpening{value: val, isSet: true}
}

func (v NullableGameJsonOpening) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameJsonOpening) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


