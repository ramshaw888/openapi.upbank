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

// CashbackObject Provides information about an instant reimbursement in the form of cashback. 
type CashbackObject struct {
	// A brief description of why this cashback was paid. 
	Description string `json:"description"`
	// The total amount of cashback paid, represented as a positive value. 
	Amount MoneyObject `json:"amount"`
}

// NewCashbackObject instantiates a new CashbackObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashbackObject(description string, amount MoneyObject) *CashbackObject {
	this := CashbackObject{}
	this.Description = description
	this.Amount = amount
	return &this
}

// NewCashbackObjectWithDefaults instantiates a new CashbackObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashbackObjectWithDefaults() *CashbackObject {
	this := CashbackObject{}
	return &this
}

// GetDescription returns the Description field value
func (o *CashbackObject) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CashbackObject) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CashbackObject) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *CashbackObject) GetAmount() MoneyObject {
	if o == nil {
		var ret MoneyObject
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CashbackObject) GetAmountOk() (*MoneyObject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CashbackObject) SetAmount(v MoneyObject) {
	o.Amount = v
}

func (o CashbackObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableCashbackObject struct {
	value *CashbackObject
	isSet bool
}

func (v NullableCashbackObject) Get() *CashbackObject {
	return v.value
}

func (v *NullableCashbackObject) Set(val *CashbackObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCashbackObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCashbackObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashbackObject(val *CashbackObject) *NullableCashbackObject {
	return &NullableCashbackObject{value: val, isSet: true}
}

func (v NullableCashbackObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashbackObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


