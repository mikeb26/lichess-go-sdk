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

// ExternalEngine struct for ExternalEngine
type ExternalEngine struct {
	// Unique engine registration ID.
	Id interface{} `json:"id"`
	// Display name of the engine.
	Name interface{} `json:"name"`
	// A secret token that can be used to [*request* analysis](#tag/External-engine/operation/apiExternalEngineAnalyse) from this external engine. 
	ClientSecret interface{} `json:"clientSecret"`
	// The user this engine has been registered for.
	UserId interface{} `json:"userId"`
	// Maximum number of available threads.
	MaxThreads interface{} `json:"maxThreads"`
	// Maximum available hash table size, in MiB.
	MaxHash interface{} `json:"maxHash"`
	// Estimated depth of normal search.
	DefaultDepth interface{} `json:"defaultDepth"`
	// List of supported chess variants.
	Variants interface{} `json:"variants"`
	// Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the `providerSecret`. 
	ProviderData interface{} `json:"providerData,omitempty"`
}

// NewExternalEngine instantiates a new ExternalEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEngine(id interface{}, name interface{}, clientSecret interface{}, userId interface{}, maxThreads interface{}, maxHash interface{}, defaultDepth interface{}, variants interface{}) *ExternalEngine {
	this := ExternalEngine{}
	this.Id = id
	this.Name = name
	this.ClientSecret = clientSecret
	this.UserId = userId
	this.MaxThreads = maxThreads
	this.MaxHash = maxHash
	this.DefaultDepth = defaultDepth
	this.Variants = variants
	return &this
}

// NewExternalEngineWithDefaults instantiates a new ExternalEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEngineWithDefaults() *ExternalEngine {
	this := ExternalEngine{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalEngine) SetId(v interface{}) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalEngine) SetName(v interface{}) {
	o.Name = v
}

// GetClientSecret returns the ClientSecret field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetClientSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetClientSecretOk() (*interface{}, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ExternalEngine) SetClientSecret(v interface{}) {
	o.ClientSecret = v
}

// GetUserId returns the UserId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetUserId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetUserIdOk() (*interface{}, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ExternalEngine) SetUserId(v interface{}) {
	o.UserId = v
}

// GetMaxThreads returns the MaxThreads field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetMaxThreads() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MaxThreads
}

// GetMaxThreadsOk returns a tuple with the MaxThreads field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetMaxThreadsOk() (*interface{}, bool) {
	if o == nil || o.MaxThreads == nil {
		return nil, false
	}
	return &o.MaxThreads, true
}

// SetMaxThreads sets field value
func (o *ExternalEngine) SetMaxThreads(v interface{}) {
	o.MaxThreads = v
}

// GetMaxHash returns the MaxHash field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetMaxHash() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MaxHash
}

// GetMaxHashOk returns a tuple with the MaxHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetMaxHashOk() (*interface{}, bool) {
	if o == nil || o.MaxHash == nil {
		return nil, false
	}
	return &o.MaxHash, true
}

// SetMaxHash sets field value
func (o *ExternalEngine) SetMaxHash(v interface{}) {
	o.MaxHash = v
}

// GetDefaultDepth returns the DefaultDepth field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetDefaultDepth() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DefaultDepth
}

// GetDefaultDepthOk returns a tuple with the DefaultDepth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetDefaultDepthOk() (*interface{}, bool) {
	if o == nil || o.DefaultDepth == nil {
		return nil, false
	}
	return &o.DefaultDepth, true
}

// SetDefaultDepth sets field value
func (o *ExternalEngine) SetDefaultDepth(v interface{}) {
	o.DefaultDepth = v
}

// GetVariants returns the Variants field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalEngine) GetVariants() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetVariantsOk() (*interface{}, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value
func (o *ExternalEngine) SetVariants(v interface{}) {
	o.Variants = v
}

// GetProviderData returns the ProviderData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalEngine) GetProviderData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ProviderData
}

// GetProviderDataOk returns a tuple with the ProviderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetProviderDataOk() (*interface{}, bool) {
	if o == nil || o.ProviderData == nil {
		return nil, false
	}
	return &o.ProviderData, true
}

// HasProviderData returns a boolean if a field has been set.
func (o *ExternalEngine) HasProviderData() bool {
	if o != nil && o.ProviderData != nil {
		return true
	}

	return false
}

// SetProviderData gets a reference to the given interface{} and assigns it to the ProviderData field.
func (o *ExternalEngine) SetProviderData(v interface{}) {
	o.ProviderData = v
}

func (o ExternalEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.MaxThreads != nil {
		toSerialize["maxThreads"] = o.MaxThreads
	}
	if o.MaxHash != nil {
		toSerialize["maxHash"] = o.MaxHash
	}
	if o.DefaultDepth != nil {
		toSerialize["defaultDepth"] = o.DefaultDepth
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.ProviderData != nil {
		toSerialize["providerData"] = o.ProviderData
	}
	return json.Marshal(toSerialize)
}

type NullableExternalEngine struct {
	value *ExternalEngine
	isSet bool
}

func (v NullableExternalEngine) Get() *ExternalEngine {
	return v.value
}

func (v *NullableExternalEngine) Set(val *ExternalEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEngine(val *ExternalEngine) *NullableExternalEngine {
	return &NullableExternalEngine{value: val, isSet: true}
}

func (v NullableExternalEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


