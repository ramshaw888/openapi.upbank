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

// UpdateTransactionCategoryRequest Request to update the category associated with a transaction. 
type UpdateTransactionCategoryRequest struct {
	// The category to set on the transaction. Set this entire key to `null` de-categorize a transaction. 
	Data NullableCategoryInputResourceIdentifier `json:"data"`
}

// NewUpdateTransactionCategoryRequest instantiates a new UpdateTransactionCategoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTransactionCategoryRequest(data NullableCategoryInputResourceIdentifier) *UpdateTransactionCategoryRequest {
	this := UpdateTransactionCategoryRequest{}
	this.Data = data
	return &this
}

// NewUpdateTransactionCategoryRequestWithDefaults instantiates a new UpdateTransactionCategoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTransactionCategoryRequestWithDefaults() *UpdateTransactionCategoryRequest {
	this := UpdateTransactionCategoryRequest{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for CategoryInputResourceIdentifier will be returned
func (o *UpdateTransactionCategoryRequest) GetData() CategoryInputResourceIdentifier {
	if o == nil || o.Data.Get() == nil {
		var ret CategoryInputResourceIdentifier
		return ret
	}

	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTransactionCategoryRequest) GetDataOk() (*CategoryInputResourceIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value
func (o *UpdateTransactionCategoryRequest) SetData(v CategoryInputResourceIdentifier) {
	o.Data.Set(&v)
}

func (o UpdateTransactionCategoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTransactionCategoryRequest struct {
	value *UpdateTransactionCategoryRequest
	isSet bool
}

func (v NullableUpdateTransactionCategoryRequest) Get() *UpdateTransactionCategoryRequest {
	return v.value
}

func (v *NullableUpdateTransactionCategoryRequest) Set(val *UpdateTransactionCategoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTransactionCategoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTransactionCategoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTransactionCategoryRequest(val *UpdateTransactionCategoryRequest) *NullableUpdateTransactionCategoryRequest {
	return &NullableUpdateTransactionCategoryRequest{value: val, isSet: true}
}

func (v NullableUpdateTransactionCategoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTransactionCategoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


