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

// ApiExternalEngineAnalyse200Response struct for ApiExternalEngineAnalyse200Response
type ApiExternalEngineAnalyse200Response struct {
	// Number of milliseconds the search has been going on
	Time interface{} `json:"time"`
	// Current search depth
	Depth interface{} `json:"depth"`
	// Number of nodes visited so far
	Nodes interface{} `json:"nodes"`
	// Information about up to 5 pvs, with the primary pv at index 0.
	Pvs []ApiExternalEngineAnalyse200ResponsePvsInner `json:"pvs"`
}

// NewApiExternalEngineAnalyse200Response instantiates a new ApiExternalEngineAnalyse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiExternalEngineAnalyse200Response(time interface{}, depth interface{}, nodes interface{}, pvs []ApiExternalEngineAnalyse200ResponsePvsInner) *ApiExternalEngineAnalyse200Response {
	this := ApiExternalEngineAnalyse200Response{}
	this.Time = time
	this.Depth = depth
	this.Nodes = nodes
	this.Pvs = pvs
	return &this
}

// NewApiExternalEngineAnalyse200ResponseWithDefaults instantiates a new ApiExternalEngineAnalyse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiExternalEngineAnalyse200ResponseWithDefaults() *ApiExternalEngineAnalyse200Response {
	this := ApiExternalEngineAnalyse200Response{}
	return &this
}

// GetTime returns the Time field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApiExternalEngineAnalyse200Response) GetTime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiExternalEngineAnalyse200Response) GetTimeOk() (*interface{}, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ApiExternalEngineAnalyse200Response) SetTime(v interface{}) {
	o.Time = v
}

// GetDepth returns the Depth field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApiExternalEngineAnalyse200Response) GetDepth() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiExternalEngineAnalyse200Response) GetDepthOk() (*interface{}, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *ApiExternalEngineAnalyse200Response) SetDepth(v interface{}) {
	o.Depth = v
}

// GetNodes returns the Nodes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApiExternalEngineAnalyse200Response) GetNodes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiExternalEngineAnalyse200Response) GetNodesOk() (*interface{}, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *ApiExternalEngineAnalyse200Response) SetNodes(v interface{}) {
	o.Nodes = v
}

// GetPvs returns the Pvs field value
func (o *ApiExternalEngineAnalyse200Response) GetPvs() []ApiExternalEngineAnalyse200ResponsePvsInner {
	if o == nil {
		var ret []ApiExternalEngineAnalyse200ResponsePvsInner
		return ret
	}

	return o.Pvs
}

// GetPvsOk returns a tuple with the Pvs field value
// and a boolean to check if the value has been set.
func (o *ApiExternalEngineAnalyse200Response) GetPvsOk() ([]ApiExternalEngineAnalyse200ResponsePvsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pvs, true
}

// SetPvs sets field value
func (o *ApiExternalEngineAnalyse200Response) SetPvs(v []ApiExternalEngineAnalyse200ResponsePvsInner) {
	o.Pvs = v
}

func (o ApiExternalEngineAnalyse200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Depth != nil {
		toSerialize["depth"] = o.Depth
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if true {
		toSerialize["pvs"] = o.Pvs
	}
	return json.Marshal(toSerialize)
}

type NullableApiExternalEngineAnalyse200Response struct {
	value *ApiExternalEngineAnalyse200Response
	isSet bool
}

func (v NullableApiExternalEngineAnalyse200Response) Get() *ApiExternalEngineAnalyse200Response {
	return v.value
}

func (v *NullableApiExternalEngineAnalyse200Response) Set(val *ApiExternalEngineAnalyse200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiExternalEngineAnalyse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiExternalEngineAnalyse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiExternalEngineAnalyse200Response(val *ApiExternalEngineAnalyse200Response) *NullableApiExternalEngineAnalyse200Response {
	return &NullableApiExternalEngineAnalyse200Response{value: val, isSet: true}
}

func (v NullableApiExternalEngineAnalyse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiExternalEngineAnalyse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


