/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ListAccountsResponse Successful response to get all accounts. This returns a paginated list of accounts, which can be scrolled by following the `prev` and `next` links if present. 
type ListAccountsResponse struct {
	// The list of accounts returned in this response. 
	Data []AccountResource `json:"data"`
	Links ListAccountsResponseLinks `json:"links"`
}

// NewListAccountsResponse instantiates a new ListAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccountsResponse(data []AccountResource, links ListAccountsResponseLinks) *ListAccountsResponse {
	this := ListAccountsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewListAccountsResponseWithDefaults instantiates a new ListAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccountsResponseWithDefaults() *ListAccountsResponse {
	this := ListAccountsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListAccountsResponse) GetData() []AccountResource {
	if o == nil {
		var ret []AccountResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetDataOk() (*[]AccountResource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListAccountsResponse) SetData(v []AccountResource) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ListAccountsResponse) GetLinks() ListAccountsResponseLinks {
	if o == nil {
		var ret ListAccountsResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ListAccountsResponse) SetLinks(v ListAccountsResponseLinks) {
	o.Links = v
}

func (o ListAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableListAccountsResponse struct {
	value *ListAccountsResponse
	isSet bool
}

func (v NullableListAccountsResponse) Get() *ListAccountsResponse {
	return v.value
}

func (v *NullableListAccountsResponse) Set(val *ListAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccountsResponse(val *ListAccountsResponse) *NullableListAccountsResponse {
	return &NullableListAccountsResponse{value: val, isSet: true}
}

func (v NullableListAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


