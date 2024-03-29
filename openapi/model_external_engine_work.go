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

// ExternalEngineWork struct for ExternalEngineWork
type ExternalEngineWork struct {
	// Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.
	SessionId interface{} `json:"sessionId"`
	// Number of threads to use for analysis.
	Threads interface{} `json:"threads"`
	// Hash table size to use for analysis, in MiB.
	Hash interface{} `json:"hash"`
	// Request an infinite search (rather than roughly aiming for `defaultDepth`).
	Infinite interface{} `json:"infinite,omitempty"`
	// Requested number of principal variations.
	MultiPv interface{} `json:"multiPv"`
	Variant UciVariant  `json:"variant"`
	// Initial position of the game.
	InitialFen interface{} `json:"initialFen"`
	// List of moves played from the initial position, in UCI notation.
	Moves interface{} `json:"moves"`
}

// NewExternalEngineWork instantiates a new ExternalEngineWork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEngineWork(sessionId interface{}, threads interface{}, hash interface{}, multiPv interface{}, variant UciVariant, initialFen interface{}, moves interface{}) *ExternalEngineWork {
	this := ExternalEngineWork{}
	this.SessionId = sessionId
	this.Threads = threads
	this.Hash = hash
	this.MultiPv = multiPv
	this.Variant = variant
	this.InitialFen = initialFen
	this.Moves = moves
	return &this
}

// NewExternalEngineWorkWithDefaults instantiates a new ExternalEngineWork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEngineWorkWithDefaults() *ExternalEngineWork {
	this := ExternalEngineWork{}
	var variant UciVariant
	this.Variant = variant
	return &this
}

// GetSessionId returns the SessionId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngineWork) GetSessionId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetSessionIdOk() (*interface{}, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ExternalEngineWork) SetSessionId(v interface{}) {
	o.SessionId = v
}

// GetThreads returns the Threads field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngineWork) GetThreads() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetThreadsOk() (*interface{}, bool) {
	if o == nil || o.Threads == nil {
		return nil, false
	}
	return &o.Threads, true
}

// SetThreads sets field value
func (o *ExternalEngineWork) SetThreads(v interface{}) {
	o.Threads = v
}

// GetHash returns the Hash field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngineWork) GetHash() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetHashOk() (*interface{}, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *ExternalEngineWork) SetHash(v interface{}) {
	o.Hash = v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalEngineWork) GetInfinite() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetInfiniteOk() (*interface{}, bool) {
	if o == nil || o.Infinite == nil {
		return nil, false
	}
	return &o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *ExternalEngineWork) HasInfinite() bool {
	if o != nil && o.Infinite != nil {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given interface{} and assigns it to the Infinite field.
func (o *ExternalEngineWork) SetInfinite(v interface{}) {
	o.Infinite = v
}

// GetMultiPv returns the MultiPv field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngineWork) GetMultiPv() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MultiPv
}

// GetMultiPvOk returns a tuple with the MultiPv field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetMultiPvOk() (*interface{}, bool) {
	if o == nil || o.MultiPv == nil {
		return nil, false
	}
	return &o.MultiPv, true
}

// SetMultiPv sets field value
func (o *ExternalEngineWork) SetMultiPv(v interface{}) {
	o.MultiPv = v
}

// GetVariant returns the Variant field value
func (o *ExternalEngineWork) GetVariant() UciVariant {
	if o == nil {
		var ret UciVariant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWork) GetVariantOk() (*UciVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ExternalEngineWork) SetVariant(v UciVariant) {
	o.Variant = v
}

// GetInitialFen returns the InitialFen field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngineWork) GetInitialFen() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetInitialFenOk() (*interface{}, bool) {
	if o == nil || o.InitialFen == nil {
		return nil, false
	}
	return &o.InitialFen, true
}

// SetInitialFen sets field value
func (o *ExternalEngineWork) SetInitialFen(v interface{}) {
	o.InitialFen = v
}

// GetMoves returns the Moves field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngineWork) GetMoves() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Moves
}

// GetMovesOk returns a tuple with the Moves field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngineWork) GetMovesOk() (*interface{}, bool) {
	if o == nil || o.Moves == nil {
		return nil, false
	}
	return &o.Moves, true
}

// SetMoves sets field value
func (o *ExternalEngineWork) SetMoves(v interface{}) {
	o.Moves = v
}

func (o ExternalEngineWork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.Threads != nil {
		toSerialize["threads"] = o.Threads
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Infinite != nil {
		toSerialize["infinite"] = o.Infinite
	}
	if o.MultiPv != nil {
		toSerialize["multiPv"] = o.MultiPv
	}
	if true {
		toSerialize["variant"] = o.Variant
	}
	if o.InitialFen != nil {
		toSerialize["initialFen"] = o.InitialFen
	}
	if o.Moves != nil {
		toSerialize["moves"] = o.Moves
	}
	return json.Marshal(toSerialize)
}

type NullableExternalEngineWork struct {
	value *ExternalEngineWork
	isSet bool
}

func (v NullableExternalEngineWork) Get() *ExternalEngineWork {
	return v.value
}

func (v *NullableExternalEngineWork) Set(val *ExternalEngineWork) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEngineWork) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEngineWork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEngineWork(val *ExternalEngineWork) *NullableExternalEngineWork {
	return &NullableExternalEngineWork{value: val, isSet: true}
}

func (v NullableExternalEngineWork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEngineWork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
