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

// ListTransactionsResponse Successful response to get all transactions. This returns a paginated list of transactions, which can be scrolled by following the `prev` and `next` links if present. 
type ListTransactionsResponse struct {
	// The list of transactions returned in this response. 
	Data []TransactionResource `json:"data"`
	Links ListAccountsResponseLinks `json:"links"`
}

// NewListTransactionsResponse instantiates a new ListTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsResponse(data []TransactionResource, links ListAccountsResponseLinks) *ListTransactionsResponse {
	this := ListTransactionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewListTransactionsResponseWithDefaults instantiates a new ListTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsResponseWithDefaults() *ListTransactionsResponse {
	this := ListTransactionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListTransactionsResponse) GetData() []TransactionResource {
	if o == nil {
		var ret []TransactionResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsResponse) GetDataOk() ([]TransactionResource, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListTransactionsResponse) SetData(v []TransactionResource) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ListTransactionsResponse) GetLinks() ListAccountsResponseLinks {
	if o == nil {
		var ret ListAccountsResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ListTransactionsResponse) SetLinks(v ListAccountsResponseLinks) {
	o.Links = v
}

func (o ListTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsResponse struct {
	value *ListTransactionsResponse
	isSet bool
}

func (v NullableListTransactionsResponse) Get() *ListTransactionsResponse {
	return v.value
}

func (v *NullableListTransactionsResponse) Set(val *ListTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsResponse(val *ListTransactionsResponse) *NullableListTransactionsResponse {
	return &NullableListTransactionsResponse{value: val, isSet: true}
}

func (v NullableListTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


