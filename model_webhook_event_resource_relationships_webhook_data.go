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

// WebhookEventResourceRelationshipsWebhookData struct for WebhookEventResourceRelationshipsWebhookData
type WebhookEventResourceRelationshipsWebhookData struct {
	// The type of this resource: `webhooks`
	Type string `json:"type"`
	// The unique identifier of the resource within its type. 
	Id string `json:"id"`
}

// NewWebhookEventResourceRelationshipsWebhookData instantiates a new WebhookEventResourceRelationshipsWebhookData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventResourceRelationshipsWebhookData(type_ string, id string) *WebhookEventResourceRelationshipsWebhookData {
	this := WebhookEventResourceRelationshipsWebhookData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewWebhookEventResourceRelationshipsWebhookDataWithDefaults instantiates a new WebhookEventResourceRelationshipsWebhookData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventResourceRelationshipsWebhookDataWithDefaults() *WebhookEventResourceRelationshipsWebhookData {
	this := WebhookEventResourceRelationshipsWebhookData{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookEventResourceRelationshipsWebhookData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResourceRelationshipsWebhookData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookEventResourceRelationshipsWebhookData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *WebhookEventResourceRelationshipsWebhookData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResourceRelationshipsWebhookData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookEventResourceRelationshipsWebhookData) SetId(v string) {
	o.Id = v
}

func (o WebhookEventResourceRelationshipsWebhookData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookEventResourceRelationshipsWebhookData struct {
	value *WebhookEventResourceRelationshipsWebhookData
	isSet bool
}

func (v NullableWebhookEventResourceRelationshipsWebhookData) Get() *WebhookEventResourceRelationshipsWebhookData {
	return v.value
}

func (v *NullableWebhookEventResourceRelationshipsWebhookData) Set(val *WebhookEventResourceRelationshipsWebhookData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventResourceRelationshipsWebhookData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventResourceRelationshipsWebhookData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventResourceRelationshipsWebhookData(val *WebhookEventResourceRelationshipsWebhookData) *NullableWebhookEventResourceRelationshipsWebhookData {
	return &NullableWebhookEventResourceRelationshipsWebhookData{value: val, isSet: true}
}

func (v NullableWebhookEventResourceRelationshipsWebhookData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventResourceRelationshipsWebhookData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


