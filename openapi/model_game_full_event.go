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

// GameFullEvent struct for GameFullEvent
type GameFullEvent struct {
	Type interface{} `json:"type"`
	Id interface{} `json:"id"`
	Variant Variant `json:"variant"`
	Clock interface{} `json:"clock"`
	Speed Speed `json:"speed"`
	Perf GameFullEventPerf `json:"perf"`
	Rated interface{} `json:"rated"`
	CreatedAt interface{} `json:"createdAt"`
	White GameEventPlayer `json:"white"`
	Black GameEventPlayer `json:"black"`
	InitialFen interface{} `json:"initialFen"`
	State GameStateEvent `json:"state"`
	TournamentId interface{} `json:"tournamentId,omitempty"`
}

// NewGameFullEvent instantiates a new GameFullEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameFullEvent(type_ interface{}, id interface{}, variant Variant, clock interface{}, speed Speed, perf GameFullEventPerf, rated interface{}, createdAt interface{}, white GameEventPlayer, black GameEventPlayer, initialFen interface{}, state GameStateEvent) *GameFullEvent {
	this := GameFullEvent{}
	this.Type = type_
	this.Id = id
	this.Variant = variant
	this.Clock = clock
	this.Speed = speed
	this.Perf = perf
	this.Rated = rated
	this.CreatedAt = createdAt
	this.White = white
	this.Black = black
	this.InitialFen = initialFen
	this.State = state
	return &this
}

// NewGameFullEventWithDefaults instantiates a new GameFullEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameFullEventWithDefaults() *GameFullEvent {
	this := GameFullEvent{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameFullEvent) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameFullEvent) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameFullEvent) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameFullEvent) SetId(v interface{}) {
	o.Id = v
}

// GetVariant returns the Variant field value
func (o *GameFullEvent) GetVariant() Variant {
	if o == nil {
		var ret Variant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetVariantOk() (*Variant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *GameFullEvent) SetVariant(v Variant) {
	o.Variant = v
}

// GetClock returns the Clock field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameFullEvent) GetClock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Clock
}

// GetClockOk returns a tuple with the Clock field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetClockOk() (*interface{}, bool) {
	if o == nil || o.Clock == nil {
		return nil, false
	}
	return &o.Clock, true
}

// SetClock sets field value
func (o *GameFullEvent) SetClock(v interface{}) {
	o.Clock = v
}

// GetSpeed returns the Speed field value
func (o *GameFullEvent) GetSpeed() Speed {
	if o == nil {
		var ret Speed
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetSpeedOk() (*Speed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *GameFullEvent) SetSpeed(v Speed) {
	o.Speed = v
}

// GetPerf returns the Perf field value
func (o *GameFullEvent) GetPerf() GameFullEventPerf {
	if o == nil {
		var ret GameFullEventPerf
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetPerfOk() (*GameFullEventPerf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *GameFullEvent) SetPerf(v GameFullEventPerf) {
	o.Perf = v
}

// GetRated returns the Rated field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameFullEvent) GetRated() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetRatedOk() (*interface{}, bool) {
	if o == nil || o.Rated == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *GameFullEvent) SetRated(v interface{}) {
	o.Rated = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameFullEvent) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GameFullEvent) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetWhite returns the White field value
func (o *GameFullEvent) GetWhite() GameEventPlayer {
	if o == nil {
		var ret GameEventPlayer
		return ret
	}

	return o.White
}

// GetWhiteOk returns a tuple with the White field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetWhiteOk() (*GameEventPlayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.White, true
}

// SetWhite sets field value
func (o *GameFullEvent) SetWhite(v GameEventPlayer) {
	o.White = v
}

// GetBlack returns the Black field value
func (o *GameFullEvent) GetBlack() GameEventPlayer {
	if o == nil {
		var ret GameEventPlayer
		return ret
	}

	return o.Black
}

// GetBlackOk returns a tuple with the Black field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetBlackOk() (*GameEventPlayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Black, true
}

// SetBlack sets field value
func (o *GameFullEvent) SetBlack(v GameEventPlayer) {
	o.Black = v
}

// GetInitialFen returns the InitialFen field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameFullEvent) GetInitialFen() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetInitialFenOk() (*interface{}, bool) {
	if o == nil || o.InitialFen == nil {
		return nil, false
	}
	return &o.InitialFen, true
}

// SetInitialFen sets field value
func (o *GameFullEvent) SetInitialFen(v interface{}) {
	o.InitialFen = v
}

// GetState returns the State field value
func (o *GameFullEvent) GetState() GameStateEvent {
	if o == nil {
		var ret GameStateEvent
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetStateOk() (*GameStateEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *GameFullEvent) SetState(v GameStateEvent) {
	o.State = v
}

// GetTournamentId returns the TournamentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameFullEvent) GetTournamentId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TournamentId
}

// GetTournamentIdOk returns a tuple with the TournamentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameFullEvent) GetTournamentIdOk() (*interface{}, bool) {
	if o == nil || o.TournamentId == nil {
		return nil, false
	}
	return &o.TournamentId, true
}

// HasTournamentId returns a boolean if a field has been set.
func (o *GameFullEvent) HasTournamentId() bool {
	if o != nil && o.TournamentId != nil {
		return true
	}

	return false
}

// SetTournamentId gets a reference to the given interface{} and assigns it to the TournamentId field.
func (o *GameFullEvent) SetTournamentId(v interface{}) {
	o.TournamentId = v
}

func (o GameFullEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["variant"] = o.Variant
	}
	if o.Clock != nil {
		toSerialize["clock"] = o.Clock
	}
	if true {
		toSerialize["speed"] = o.Speed
	}
	if true {
		toSerialize["perf"] = o.Perf
	}
	if o.Rated != nil {
		toSerialize["rated"] = o.Rated
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["white"] = o.White
	}
	if true {
		toSerialize["black"] = o.Black
	}
	if o.InitialFen != nil {
		toSerialize["initialFen"] = o.InitialFen
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.TournamentId != nil {
		toSerialize["tournamentId"] = o.TournamentId
	}
	return json.Marshal(toSerialize)
}

type NullableGameFullEvent struct {
	value *GameFullEvent
	isSet bool
}

func (v NullableGameFullEvent) Get() *GameFullEvent {
	return v.value
}

func (v *NullableGameFullEvent) Set(val *GameFullEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableGameFullEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableGameFullEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameFullEvent(val *GameFullEvent) *NullableGameFullEvent {
	return &NullableGameFullEvent{value: val, isSet: true}
}

func (v NullableGameFullEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameFullEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


