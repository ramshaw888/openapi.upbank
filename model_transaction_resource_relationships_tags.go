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

// TransactionResourceRelationshipsTags struct for TransactionResourceRelationshipsTags
type TransactionResourceRelationshipsTags struct {
	Data []TransactionResourceRelationshipsTagsData `json:"data"`
	Links *TransactionResourceRelationshipsTagsLinks `json:"links,omitempty"`
}

// NewTransactionResourceRelationshipsTags instantiates a new TransactionResourceRelationshipsTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResourceRelationshipsTags(data []TransactionResourceRelationshipsTagsData) *TransactionResourceRelationshipsTags {
	this := TransactionResourceRelationshipsTags{}
	this.Data = data
	return &this
}

// NewTransactionResourceRelationshipsTagsWithDefaults instantiates a new TransactionResourceRelationshipsTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResourceRelationshipsTagsWithDefaults() *TransactionResourceRelationshipsTags {
	this := TransactionResourceRelationshipsTags{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionResourceRelationshipsTags) GetData() []TransactionResourceRelationshipsTagsData {
	if o == nil {
		var ret []TransactionResourceRelationshipsTagsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsTags) GetDataOk() (*[]TransactionResourceRelationshipsTagsData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionResourceRelationshipsTags) SetData(v []TransactionResourceRelationshipsTagsData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TransactionResourceRelationshipsTags) GetLinks() TransactionResourceRelationshipsTagsLinks {
	if o == nil || o.Links == nil {
		var ret TransactionResourceRelationshipsTagsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsTags) GetLinksOk() (*TransactionResourceRelationshipsTagsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TransactionResourceRelationshipsTags) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TransactionResourceRelationshipsTagsLinks and assigns it to the Links field.
func (o *TransactionResourceRelationshipsTags) SetLinks(v TransactionResourceRelationshipsTagsLinks) {
	o.Links = &v
}

func (o TransactionResourceRelationshipsTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionResourceRelationshipsTags struct {
	value *TransactionResourceRelationshipsTags
	isSet bool
}

func (v NullableTransactionResourceRelationshipsTags) Get() *TransactionResourceRelationshipsTags {
	return v.value
}

func (v *NullableTransactionResourceRelationshipsTags) Set(val *TransactionResourceRelationshipsTags) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResourceRelationshipsTags) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResourceRelationshipsTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResourceRelationshipsTags(val *TransactionResourceRelationshipsTags) *NullableTransactionResourceRelationshipsTags {
	return &NullableTransactionResourceRelationshipsTags{value: val, isSet: true}
}

func (v NullableTransactionResourceRelationshipsTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResourceRelationshipsTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


