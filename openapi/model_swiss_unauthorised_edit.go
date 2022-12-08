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

// SwissUnauthorisedEdit struct for SwissUnauthorisedEdit
type SwissUnauthorisedEdit struct {
	Error interface{} `json:"error,omitempty"`
}

// NewSwissUnauthorisedEdit instantiates a new SwissUnauthorisedEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwissUnauthorisedEdit() *SwissUnauthorisedEdit {
	this := SwissUnauthorisedEdit{}
	return &this
}

// NewSwissUnauthorisedEditWithDefaults instantiates a new SwissUnauthorisedEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwissUnauthorisedEditWithDefaults() *SwissUnauthorisedEdit {
	this := SwissUnauthorisedEdit{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwissUnauthorisedEdit) GetError() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwissUnauthorisedEdit) GetErrorOk() (*interface{}, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SwissUnauthorisedEdit) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given interface{} and assigns it to the Error field.
func (o *SwissUnauthorisedEdit) SetError(v interface{}) {
	o.Error = v
}

func (o SwissUnauthorisedEdit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableSwissUnauthorisedEdit struct {
	value *SwissUnauthorisedEdit
	isSet bool
}

func (v NullableSwissUnauthorisedEdit) Get() *SwissUnauthorisedEdit {
	return v.value
}

func (v *NullableSwissUnauthorisedEdit) Set(val *SwissUnauthorisedEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableSwissUnauthorisedEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableSwissUnauthorisedEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwissUnauthorisedEdit(val *SwissUnauthorisedEdit) *NullableSwissUnauthorisedEdit {
	return &NullableSwissUnauthorisedEdit{value: val, isSet: true}
}

func (v NullableSwissUnauthorisedEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwissUnauthorisedEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


