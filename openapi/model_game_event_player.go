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

// GameEventPlayer struct for GameEventPlayer
type GameEventPlayer struct {
	AiLevel interface{} `json:"aiLevel,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Title interface{} `json:"title,omitempty"`
	Rating interface{} `json:"rating,omitempty"`
	Provisional interface{} `json:"provisional,omitempty"`
}

// NewGameEventPlayer instantiates a new GameEventPlayer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameEventPlayer() *GameEventPlayer {
	this := GameEventPlayer{}
	return &this
}

// NewGameEventPlayerWithDefaults instantiates a new GameEventPlayer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameEventPlayerWithDefaults() *GameEventPlayer {
	this := GameEventPlayer{}
	return &this
}

// GetAiLevel returns the AiLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventPlayer) GetAiLevel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AiLevel
}

// GetAiLevelOk returns a tuple with the AiLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventPlayer) GetAiLevelOk() (*interface{}, bool) {
	if o == nil || o.AiLevel == nil {
		return nil, false
	}
	return &o.AiLevel, true
}

// HasAiLevel returns a boolean if a field has been set.
func (o *GameEventPlayer) HasAiLevel() bool {
	if o != nil && o.AiLevel != nil {
		return true
	}

	return false
}

// SetAiLevel gets a reference to the given interface{} and assigns it to the AiLevel field.
func (o *GameEventPlayer) SetAiLevel(v interface{}) {
	o.AiLevel = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventPlayer) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventPlayer) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GameEventPlayer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GameEventPlayer) SetId(v interface{}) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventPlayer) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventPlayer) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GameEventPlayer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GameEventPlayer) SetName(v interface{}) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventPlayer) GetTitle() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventPlayer) GetTitleOk() (*interface{}, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GameEventPlayer) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given interface{} and assigns it to the Title field.
func (o *GameEventPlayer) SetTitle(v interface{}) {
	o.Title = v
}

// GetRating returns the Rating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventPlayer) GetRating() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventPlayer) GetRatingOk() (*interface{}, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return &o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *GameEventPlayer) HasRating() bool {
	if o != nil && o.Rating != nil {
		return true
	}

	return false
}

// SetRating gets a reference to the given interface{} and assigns it to the Rating field.
func (o *GameEventPlayer) SetRating(v interface{}) {
	o.Rating = v
}

// GetProvisional returns the Provisional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameEventPlayer) GetProvisional() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Provisional
}

// GetProvisionalOk returns a tuple with the Provisional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameEventPlayer) GetProvisionalOk() (*interface{}, bool) {
	if o == nil || o.Provisional == nil {
		return nil, false
	}
	return &o.Provisional, true
}

// HasProvisional returns a boolean if a field has been set.
func (o *GameEventPlayer) HasProvisional() bool {
	if o != nil && o.Provisional != nil {
		return true
	}

	return false
}

// SetProvisional gets a reference to the given interface{} and assigns it to the Provisional field.
func (o *GameEventPlayer) SetProvisional(v interface{}) {
	o.Provisional = v
}

func (o GameEventPlayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AiLevel != nil {
		toSerialize["aiLevel"] = o.AiLevel
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	if o.Provisional != nil {
		toSerialize["provisional"] = o.Provisional
	}
	return json.Marshal(toSerialize)
}

type NullableGameEventPlayer struct {
	value *GameEventPlayer
	isSet bool
}

func (v NullableGameEventPlayer) Get() *GameEventPlayer {
	return v.value
}

func (v *NullableGameEventPlayer) Set(val *GameEventPlayer) {
	v.value = val
	v.isSet = true
}

func (v NullableGameEventPlayer) IsSet() bool {
	return v.isSet
}

func (v *NullableGameEventPlayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameEventPlayer(val *GameEventPlayer) *NullableGameEventPlayer {
	return &NullableGameEventPlayer{value: val, isSet: true}
}

func (v NullableGameEventPlayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameEventPlayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


