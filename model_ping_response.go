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

// PingResponse Basic ping response to verify authentication. 
type PingResponse struct {
	Meta PingResponseMeta `json:"meta"`
}

// NewPingResponse instantiates a new PingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResponse(meta PingResponseMeta) *PingResponse {
	this := PingResponse{}
	this.Meta = meta
	return &this
}

// NewPingResponseWithDefaults instantiates a new PingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResponseWithDefaults() *PingResponse {
	this := PingResponse{}
	return &this
}

// GetMeta returns the Meta field value
func (o *PingResponse) GetMeta() PingResponseMeta {
	if o == nil {
		var ret PingResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *PingResponse) GetMetaOk() (*PingResponseMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *PingResponse) SetMeta(v PingResponseMeta) {
	o.Meta = v
}

func (o PingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullablePingResponse struct {
	value *PingResponse
	isSet bool
}

func (v NullablePingResponse) Get() *PingResponse {
	return v.value
}

func (v *NullablePingResponse) Set(val *PingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResponse(val *PingResponse) *NullablePingResponse {
	return &NullablePingResponse{value: val, isSet: true}
}

func (v NullablePingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


