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

// WebhookDeliveryLogResourceRelationshipsWebhookEvent struct for WebhookDeliveryLogResourceRelationshipsWebhookEvent
type WebhookDeliveryLogResourceRelationshipsWebhookEvent struct {
	Data WebhookDeliveryLogResourceRelationshipsWebhookEventData `json:"data"`
}

// NewWebhookDeliveryLogResourceRelationshipsWebhookEvent instantiates a new WebhookDeliveryLogResourceRelationshipsWebhookEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDeliveryLogResourceRelationshipsWebhookEvent(data WebhookDeliveryLogResourceRelationshipsWebhookEventData) *WebhookDeliveryLogResourceRelationshipsWebhookEvent {
	this := WebhookDeliveryLogResourceRelationshipsWebhookEvent{}
	this.Data = data
	return &this
}

// NewWebhookDeliveryLogResourceRelationshipsWebhookEventWithDefaults instantiates a new WebhookDeliveryLogResourceRelationshipsWebhookEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryLogResourceRelationshipsWebhookEventWithDefaults() *WebhookDeliveryLogResourceRelationshipsWebhookEvent {
	this := WebhookDeliveryLogResourceRelationshipsWebhookEvent{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) GetData() WebhookDeliveryLogResourceRelationshipsWebhookEventData {
	if o == nil {
		var ret WebhookDeliveryLogResourceRelationshipsWebhookEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) GetDataOk() (*WebhookDeliveryLogResourceRelationshipsWebhookEventData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) SetData(v WebhookDeliveryLogResourceRelationshipsWebhookEventData) {
	o.Data = v
}

func (o WebhookDeliveryLogResourceRelationshipsWebhookEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent struct {
	value *WebhookDeliveryLogResourceRelationshipsWebhookEvent
	isSet bool
}

func (v NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) Get() *WebhookDeliveryLogResourceRelationshipsWebhookEvent {
	return v.value
}

func (v *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) Set(val *WebhookDeliveryLogResourceRelationshipsWebhookEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryLogResourceRelationshipsWebhookEvent(val *WebhookDeliveryLogResourceRelationshipsWebhookEvent) *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent {
	return &NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent{value: val, isSet: true}
}

func (v NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


