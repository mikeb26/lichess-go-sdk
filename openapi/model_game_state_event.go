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

// GameStateEvent struct for GameStateEvent
type GameStateEvent struct {
	Type interface{} `json:"type"`
	// Current moves in UCI format
	Moves interface{} `json:"moves"`
	// Integer of milliseconds White has left on the clock
	Wtime interface{} `json:"wtime"`
	// Integer of milliseconds Black has left on the clock
	Btime interface{} `json:"btime"`
	// Integer of White Fisher increment.
	Winc interface{} `json:"winc"`
	// Integer of Black Fisher increment.
	Binc interface{} `json:"binc"`
	Status GameStatus `json:"status"`
	// Color of the winner, if any
	Winner interface{} `json:"winner,omitempty"`
	// true if white is offering draw, else omitted
	Wdraw interface{} `json:"wdraw,omitempty"`
	// true if black is offering draw, else omitted
	Bdraw interface{} `json:"bdraw,omitempty"`
	// true if white is proposing takeback, else omitted
	Wtakeback interface{} `json:"wtakeback,omitempty"`
	// true if black is proposing takeback, else omitted
	Btakeback interface{} `json:"btakeback,omitempty"`
}

// NewGameStateEvent instantiates a new GameStateEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameStateEvent(type_ interface{}, moves interface{}, wtime interface{}, btime interface{}, winc interface{}, binc interface{}, status GameStatus) *GameStateEvent {
	this := GameStateEvent{}
	this.Type = type_
	this.Moves = moves
	this.Wtime = wtime
	this.Btime = btime
	this.Winc = winc
	this.Binc = binc
	this.Status = status
	return &this
}

// NewGameStateEventWithDefaults instantiates a new GameStateEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameStateEventWithDefaults() *GameStateEvent {
	this := GameStateEvent{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameStateEvent) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameStateEvent) SetType(v interface{}) {
	o.Type = v
}

// GetMoves returns the Moves field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameStateEvent) GetMoves() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Moves
}

// GetMovesOk returns a tuple with the Moves field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetMovesOk() (*interface{}, bool) {
	if o == nil || o.Moves == nil {
		return nil, false
	}
	return &o.Moves, true
}

// SetMoves sets field value
func (o *GameStateEvent) SetMoves(v interface{}) {
	o.Moves = v
}

// GetWtime returns the Wtime field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameStateEvent) GetWtime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Wtime
}

// GetWtimeOk returns a tuple with the Wtime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetWtimeOk() (*interface{}, bool) {
	if o == nil || o.Wtime == nil {
		return nil, false
	}
	return &o.Wtime, true
}

// SetWtime sets field value
func (o *GameStateEvent) SetWtime(v interface{}) {
	o.Wtime = v
}

// GetBtime returns the Btime field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameStateEvent) GetBtime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Btime
}

// GetBtimeOk returns a tuple with the Btime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetBtimeOk() (*interface{}, bool) {
	if o == nil || o.Btime == nil {
		return nil, false
	}
	return &o.Btime, true
}

// SetBtime sets field value
func (o *GameStateEvent) SetBtime(v interface{}) {
	o.Btime = v
}

// GetWinc returns the Winc field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameStateEvent) GetWinc() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Winc
}

// GetWincOk returns a tuple with the Winc field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetWincOk() (*interface{}, bool) {
	if o == nil || o.Winc == nil {
		return nil, false
	}
	return &o.Winc, true
}

// SetWinc sets field value
func (o *GameStateEvent) SetWinc(v interface{}) {
	o.Winc = v
}

// GetBinc returns the Binc field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GameStateEvent) GetBinc() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Binc
}

// GetBincOk returns a tuple with the Binc field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetBincOk() (*interface{}, bool) {
	if o == nil || o.Binc == nil {
		return nil, false
	}
	return &o.Binc, true
}

// SetBinc sets field value
func (o *GameStateEvent) SetBinc(v interface{}) {
	o.Binc = v
}

// GetStatus returns the Status field value
func (o *GameStateEvent) GetStatus() GameStatus {
	if o == nil {
		var ret GameStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GameStateEvent) GetStatusOk() (*GameStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GameStateEvent) SetStatus(v GameStatus) {
	o.Status = v
}

// GetWinner returns the Winner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameStateEvent) GetWinner() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetWinnerOk() (*interface{}, bool) {
	if o == nil || o.Winner == nil {
		return nil, false
	}
	return &o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *GameStateEvent) HasWinner() bool {
	if o != nil && o.Winner != nil {
		return true
	}

	return false
}

// SetWinner gets a reference to the given interface{} and assigns it to the Winner field.
func (o *GameStateEvent) SetWinner(v interface{}) {
	o.Winner = v
}

// GetWdraw returns the Wdraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameStateEvent) GetWdraw() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Wdraw
}

// GetWdrawOk returns a tuple with the Wdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetWdrawOk() (*interface{}, bool) {
	if o == nil || o.Wdraw == nil {
		return nil, false
	}
	return &o.Wdraw, true
}

// HasWdraw returns a boolean if a field has been set.
func (o *GameStateEvent) HasWdraw() bool {
	if o != nil && o.Wdraw != nil {
		return true
	}

	return false
}

// SetWdraw gets a reference to the given interface{} and assigns it to the Wdraw field.
func (o *GameStateEvent) SetWdraw(v interface{}) {
	o.Wdraw = v
}

// GetBdraw returns the Bdraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameStateEvent) GetBdraw() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Bdraw
}

// GetBdrawOk returns a tuple with the Bdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetBdrawOk() (*interface{}, bool) {
	if o == nil || o.Bdraw == nil {
		return nil, false
	}
	return &o.Bdraw, true
}

// HasBdraw returns a boolean if a field has been set.
func (o *GameStateEvent) HasBdraw() bool {
	if o != nil && o.Bdraw != nil {
		return true
	}

	return false
}

// SetBdraw gets a reference to the given interface{} and assigns it to the Bdraw field.
func (o *GameStateEvent) SetBdraw(v interface{}) {
	o.Bdraw = v
}

// GetWtakeback returns the Wtakeback field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameStateEvent) GetWtakeback() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Wtakeback
}

// GetWtakebackOk returns a tuple with the Wtakeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetWtakebackOk() (*interface{}, bool) {
	if o == nil || o.Wtakeback == nil {
		return nil, false
	}
	return &o.Wtakeback, true
}

// HasWtakeback returns a boolean if a field has been set.
func (o *GameStateEvent) HasWtakeback() bool {
	if o != nil && o.Wtakeback != nil {
		return true
	}

	return false
}

// SetWtakeback gets a reference to the given interface{} and assigns it to the Wtakeback field.
func (o *GameStateEvent) SetWtakeback(v interface{}) {
	o.Wtakeback = v
}

// GetBtakeback returns the Btakeback field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GameStateEvent) GetBtakeback() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Btakeback
}

// GetBtakebackOk returns a tuple with the Btakeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GameStateEvent) GetBtakebackOk() (*interface{}, bool) {
	if o == nil || o.Btakeback == nil {
		return nil, false
	}
	return &o.Btakeback, true
}

// HasBtakeback returns a boolean if a field has been set.
func (o *GameStateEvent) HasBtakeback() bool {
	if o != nil && o.Btakeback != nil {
		return true
	}

	return false
}

// SetBtakeback gets a reference to the given interface{} and assigns it to the Btakeback field.
func (o *GameStateEvent) SetBtakeback(v interface{}) {
	o.Btakeback = v
}

func (o GameStateEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Moves != nil {
		toSerialize["moves"] = o.Moves
	}
	if o.Wtime != nil {
		toSerialize["wtime"] = o.Wtime
	}
	if o.Btime != nil {
		toSerialize["btime"] = o.Btime
	}
	if o.Winc != nil {
		toSerialize["winc"] = o.Winc
	}
	if o.Binc != nil {
		toSerialize["binc"] = o.Binc
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Winner != nil {
		toSerialize["winner"] = o.Winner
	}
	if o.Wdraw != nil {
		toSerialize["wdraw"] = o.Wdraw
	}
	if o.Bdraw != nil {
		toSerialize["bdraw"] = o.Bdraw
	}
	if o.Wtakeback != nil {
		toSerialize["wtakeback"] = o.Wtakeback
	}
	if o.Btakeback != nil {
		toSerialize["btakeback"] = o.Btakeback
	}
	return json.Marshal(toSerialize)
}

type NullableGameStateEvent struct {
	value *GameStateEvent
	isSet bool
}

func (v NullableGameStateEvent) Get() *GameStateEvent {
	return v.value
}

func (v *NullableGameStateEvent) Set(val *GameStateEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableGameStateEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableGameStateEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameStateEvent(val *GameStateEvent) *NullableGameStateEvent {
	return &NullableGameStateEvent{value: val, isSet: true}
}

func (v NullableGameStateEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameStateEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


