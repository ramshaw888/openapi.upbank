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

// WebhookResourceRelationships struct for WebhookResourceRelationships
type WebhookResourceRelationships struct {
	Logs AccountResourceRelationshipsTransactions `json:"logs"`
}

// NewWebhookResourceRelationships instantiates a new WebhookResourceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResourceRelationships(logs AccountResourceRelationshipsTransactions) *WebhookResourceRelationships {
	this := WebhookResourceRelationships{}
	this.Logs = logs
	return &this
}

// NewWebhookResourceRelationshipsWithDefaults instantiates a new WebhookResourceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResourceRelationshipsWithDefaults() *WebhookResourceRelationships {
	this := WebhookResourceRelationships{}
	return &this
}

// GetLogs returns the Logs field value
func (o *WebhookResourceRelationships) GetLogs() AccountResourceRelationshipsTransactions {
	if o == nil {
		var ret AccountResourceRelationshipsTransactions
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *WebhookResourceRelationships) GetLogsOk() (*AccountResourceRelationshipsTransactions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Logs, true
}

// SetLogs sets field value
func (o *WebhookResourceRelationships) SetLogs(v AccountResourceRelationshipsTransactions) {
	o.Logs = v
}

func (o WebhookResourceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookResourceRelationships struct {
	value *WebhookResourceRelationships
	isSet bool
}

func (v NullableWebhookResourceRelationships) Get() *WebhookResourceRelationships {
	return v.value
}

func (v *NullableWebhookResourceRelationships) Set(val *WebhookResourceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResourceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResourceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResourceRelationships(val *WebhookResourceRelationships) *NullableWebhookResourceRelationships {
	return &NullableWebhookResourceRelationships{value: val, isSet: true}
}

func (v NullableWebhookResourceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResourceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


