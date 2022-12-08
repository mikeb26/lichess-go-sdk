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

// TeamAll200Response struct for TeamAll200Response
type TeamAll200Response struct {
	CurrentPage interface{} `json:"currentPage,omitempty"`
	MaxPerPage interface{} `json:"maxPerPage,omitempty"`
	CurrentPageResults []Team `json:"currentPageResults,omitempty"`
	NbResults interface{} `json:"nbResults,omitempty"`
	PreviousPage interface{} `json:"previousPage,omitempty"`
	NextPage interface{} `json:"nextPage,omitempty"`
	NbPages interface{} `json:"nbPages,omitempty"`
}

// NewTeamAll200Response instantiates a new TeamAll200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamAll200Response() *TeamAll200Response {
	this := TeamAll200Response{}
	return &this
}

// NewTeamAll200ResponseWithDefaults instantiates a new TeamAll200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamAll200ResponseWithDefaults() *TeamAll200Response {
	this := TeamAll200Response{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAll200Response) GetCurrentPage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamAll200Response) GetCurrentPageOk() (*interface{}, bool) {
	if o == nil || o.CurrentPage == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *TeamAll200Response) HasCurrentPage() bool {
	if o != nil && o.CurrentPage != nil {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given interface{} and assigns it to the CurrentPage field.
func (o *TeamAll200Response) SetCurrentPage(v interface{}) {
	o.CurrentPage = v
}

// GetMaxPerPage returns the MaxPerPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAll200Response) GetMaxPerPage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MaxPerPage
}

// GetMaxPerPageOk returns a tuple with the MaxPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamAll200Response) GetMaxPerPageOk() (*interface{}, bool) {
	if o == nil || o.MaxPerPage == nil {
		return nil, false
	}
	return &o.MaxPerPage, true
}

// HasMaxPerPage returns a boolean if a field has been set.
func (o *TeamAll200Response) HasMaxPerPage() bool {
	if o != nil && o.MaxPerPage != nil {
		return true
	}

	return false
}

// SetMaxPerPage gets a reference to the given interface{} and assigns it to the MaxPerPage field.
func (o *TeamAll200Response) SetMaxPerPage(v interface{}) {
	o.MaxPerPage = v
}

// GetCurrentPageResults returns the CurrentPageResults field value if set, zero value otherwise.
func (o *TeamAll200Response) GetCurrentPageResults() []Team {
	if o == nil || o.CurrentPageResults == nil {
		var ret []Team
		return ret
	}
	return o.CurrentPageResults
}

// GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAll200Response) GetCurrentPageResultsOk() ([]Team, bool) {
	if o == nil || o.CurrentPageResults == nil {
		return nil, false
	}
	return o.CurrentPageResults, true
}

// HasCurrentPageResults returns a boolean if a field has been set.
func (o *TeamAll200Response) HasCurrentPageResults() bool {
	if o != nil && o.CurrentPageResults != nil {
		return true
	}

	return false
}

// SetCurrentPageResults gets a reference to the given []Team and assigns it to the CurrentPageResults field.
func (o *TeamAll200Response) SetCurrentPageResults(v []Team) {
	o.CurrentPageResults = v
}

// GetNbResults returns the NbResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAll200Response) GetNbResults() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NbResults
}

// GetNbResultsOk returns a tuple with the NbResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamAll200Response) GetNbResultsOk() (*interface{}, bool) {
	if o == nil || o.NbResults == nil {
		return nil, false
	}
	return &o.NbResults, true
}

// HasNbResults returns a boolean if a field has been set.
func (o *TeamAll200Response) HasNbResults() bool {
	if o != nil && o.NbResults != nil {
		return true
	}

	return false
}

// SetNbResults gets a reference to the given interface{} and assigns it to the NbResults field.
func (o *TeamAll200Response) SetNbResults(v interface{}) {
	o.NbResults = v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAll200Response) GetPreviousPage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamAll200Response) GetPreviousPageOk() (*interface{}, bool) {
	if o == nil || o.PreviousPage == nil {
		return nil, false
	}
	return &o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *TeamAll200Response) HasPreviousPage() bool {
	if o != nil && o.PreviousPage != nil {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given interface{} and assigns it to the PreviousPage field.
func (o *TeamAll200Response) SetPreviousPage(v interface{}) {
	o.PreviousPage = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAll200Response) GetNextPage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamAll200Response) GetNextPageOk() (*interface{}, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *TeamAll200Response) HasNextPage() bool {
	if o != nil && o.NextPage != nil {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given interface{} and assigns it to the NextPage field.
func (o *TeamAll200Response) SetNextPage(v interface{}) {
	o.NextPage = v
}

// GetNbPages returns the NbPages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamAll200Response) GetNbPages() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NbPages
}

// GetNbPagesOk returns a tuple with the NbPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamAll200Response) GetNbPagesOk() (*interface{}, bool) {
	if o == nil || o.NbPages == nil {
		return nil, false
	}
	return &o.NbPages, true
}

// HasNbPages returns a boolean if a field has been set.
func (o *TeamAll200Response) HasNbPages() bool {
	if o != nil && o.NbPages != nil {
		return true
	}

	return false
}

// SetNbPages gets a reference to the given interface{} and assigns it to the NbPages field.
func (o *TeamAll200Response) SetNbPages(v interface{}) {
	o.NbPages = v
}

func (o TeamAll200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPage != nil {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if o.MaxPerPage != nil {
		toSerialize["maxPerPage"] = o.MaxPerPage
	}
	if o.CurrentPageResults != nil {
		toSerialize["currentPageResults"] = o.CurrentPageResults
	}
	if o.NbResults != nil {
		toSerialize["nbResults"] = o.NbResults
	}
	if o.PreviousPage != nil {
		toSerialize["previousPage"] = o.PreviousPage
	}
	if o.NextPage != nil {
		toSerialize["nextPage"] = o.NextPage
	}
	if o.NbPages != nil {
		toSerialize["nbPages"] = o.NbPages
	}
	return json.Marshal(toSerialize)
}

type NullableTeamAll200Response struct {
	value *TeamAll200Response
	isSet bool
}

func (v NullableTeamAll200Response) Get() *TeamAll200Response {
	return v.value
}

func (v *NullableTeamAll200Response) Set(val *TeamAll200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamAll200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamAll200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamAll200Response(val *TeamAll200Response) *NullableTeamAll200Response {
	return &NullableTeamAll200Response{value: val, isSet: true}
}

func (v NullableTeamAll200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamAll200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


