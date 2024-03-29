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

// TransactionResourceRelationshipsAccount struct for TransactionResourceRelationshipsAccount
type TransactionResourceRelationshipsAccount struct {
	Data TransactionResourceRelationshipsAccountData `json:"data"`
	Links *AccountResourceRelationshipsTransactionsLinks `json:"links,omitempty"`
}

// NewTransactionResourceRelationshipsAccount instantiates a new TransactionResourceRelationshipsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResourceRelationshipsAccount(data TransactionResourceRelationshipsAccountData) *TransactionResourceRelationshipsAccount {
	this := TransactionResourceRelationshipsAccount{}
	this.Data = data
	return &this
}

// NewTransactionResourceRelationshipsAccountWithDefaults instantiates a new TransactionResourceRelationshipsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResourceRelationshipsAccountWithDefaults() *TransactionResourceRelationshipsAccount {
	this := TransactionResourceRelationshipsAccount{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionResourceRelationshipsAccount) GetData() TransactionResourceRelationshipsAccountData {
	if o == nil {
		var ret TransactionResourceRelationshipsAccountData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsAccount) GetDataOk() (*TransactionResourceRelationshipsAccountData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionResourceRelationshipsAccount) SetData(v TransactionResourceRelationshipsAccountData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TransactionResourceRelationshipsAccount) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || o.Links == nil {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsAccount) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TransactionResourceRelationshipsAccount) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *TransactionResourceRelationshipsAccount) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o TransactionResourceRelationshipsAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionResourceRelationshipsAccount struct {
	value *TransactionResourceRelationshipsAccount
	isSet bool
}

func (v NullableTransactionResourceRelationshipsAccount) Get() *TransactionResourceRelationshipsAccount {
	return v.value
}

func (v *NullableTransactionResourceRelationshipsAccount) Set(val *TransactionResourceRelationshipsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResourceRelationshipsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResourceRelationshipsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResourceRelationshipsAccount(val *TransactionResourceRelationshipsAccount) *NullableTransactionResourceRelationshipsAccount {
	return &NullableTransactionResourceRelationshipsAccount{value: val, isSet: true}
}

func (v NullableTransactionResourceRelationshipsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResourceRelationshipsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


