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

// TransactionResourceRelationshipsAccountData struct for TransactionResourceRelationshipsAccountData
type TransactionResourceRelationshipsAccountData struct {
	// The type of this resource: `accounts`
	Type string `json:"type"`
	// The unique identifier of the resource within its type. 
	Id string `json:"id"`
}

// NewTransactionResourceRelationshipsAccountData instantiates a new TransactionResourceRelationshipsAccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResourceRelationshipsAccountData(type_ string, id string) *TransactionResourceRelationshipsAccountData {
	this := TransactionResourceRelationshipsAccountData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewTransactionResourceRelationshipsAccountDataWithDefaults instantiates a new TransactionResourceRelationshipsAccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResourceRelationshipsAccountDataWithDefaults() *TransactionResourceRelationshipsAccountData {
	this := TransactionResourceRelationshipsAccountData{}
	return &this
}

// GetType returns the Type field value
func (o *TransactionResourceRelationshipsAccountData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsAccountData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionResourceRelationshipsAccountData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TransactionResourceRelationshipsAccountData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsAccountData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionResourceRelationshipsAccountData) SetId(v string) {
	o.Id = v
}

func (o TransactionResourceRelationshipsAccountData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionResourceRelationshipsAccountData struct {
	value *TransactionResourceRelationshipsAccountData
	isSet bool
}

func (v NullableTransactionResourceRelationshipsAccountData) Get() *TransactionResourceRelationshipsAccountData {
	return v.value
}

func (v *NullableTransactionResourceRelationshipsAccountData) Set(val *TransactionResourceRelationshipsAccountData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResourceRelationshipsAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResourceRelationshipsAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResourceRelationshipsAccountData(val *TransactionResourceRelationshipsAccountData) *NullableTransactionResourceRelationshipsAccountData {
	return &NullableTransactionResourceRelationshipsAccountData{value: val, isSet: true}
}

func (v NullableTransactionResourceRelationshipsAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResourceRelationshipsAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


