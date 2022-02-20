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

// CategoryResourceRelationshipsChildren struct for CategoryResourceRelationshipsChildren
type CategoryResourceRelationshipsChildren struct {
	Data []CategoryResourceRelationshipsChildrenData `json:"data"`
	Links *AccountResourceRelationshipsTransactionsLinks `json:"links,omitempty"`
}

// NewCategoryResourceRelationshipsChildren instantiates a new CategoryResourceRelationshipsChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResourceRelationshipsChildren(data []CategoryResourceRelationshipsChildrenData) *CategoryResourceRelationshipsChildren {
	this := CategoryResourceRelationshipsChildren{}
	this.Data = data
	return &this
}

// NewCategoryResourceRelationshipsChildrenWithDefaults instantiates a new CategoryResourceRelationshipsChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResourceRelationshipsChildrenWithDefaults() *CategoryResourceRelationshipsChildren {
	this := CategoryResourceRelationshipsChildren{}
	return &this
}

// GetData returns the Data field value
func (o *CategoryResourceRelationshipsChildren) GetData() []CategoryResourceRelationshipsChildrenData {
	if o == nil {
		var ret []CategoryResourceRelationshipsChildrenData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsChildren) GetDataOk() (*[]CategoryResourceRelationshipsChildrenData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CategoryResourceRelationshipsChildren) SetData(v []CategoryResourceRelationshipsChildrenData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CategoryResourceRelationshipsChildren) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || o.Links == nil {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsChildren) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CategoryResourceRelationshipsChildren) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *CategoryResourceRelationshipsChildren) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o CategoryResourceRelationshipsChildren) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryResourceRelationshipsChildren struct {
	value *CategoryResourceRelationshipsChildren
	isSet bool
}

func (v NullableCategoryResourceRelationshipsChildren) Get() *CategoryResourceRelationshipsChildren {
	return v.value
}

func (v *NullableCategoryResourceRelationshipsChildren) Set(val *CategoryResourceRelationshipsChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResourceRelationshipsChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResourceRelationshipsChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResourceRelationshipsChildren(val *CategoryResourceRelationshipsChildren) *NullableCategoryResourceRelationshipsChildren {
	return &NullableCategoryResourceRelationshipsChildren{value: val, isSet: true}
}

func (v NullableCategoryResourceRelationshipsChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResourceRelationshipsChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


