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

// WebhookDeliveryLogResource Provides historical webhook event delivery information for analysis and debugging purposes. 
type WebhookDeliveryLogResource struct {
	// The type of this resource: `webhook-delivery-logs`
	Type string `json:"type"`
	// The unique identifier for this log entry. 
	Id string `json:"id"`
	Attributes WebhookDeliveryLogResourceAttributes `json:"attributes"`
	Relationships WebhookDeliveryLogResourceRelationships `json:"relationships"`
}

// NewWebhookDeliveryLogResource instantiates a new WebhookDeliveryLogResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDeliveryLogResource(type_ string, id string, attributes WebhookDeliveryLogResourceAttributes, relationships WebhookDeliveryLogResourceRelationships) *WebhookDeliveryLogResource {
	this := WebhookDeliveryLogResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewWebhookDeliveryLogResourceWithDefaults instantiates a new WebhookDeliveryLogResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryLogResourceWithDefaults() *WebhookDeliveryLogResource {
	this := WebhookDeliveryLogResource{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookDeliveryLogResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookDeliveryLogResource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *WebhookDeliveryLogResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResource) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookDeliveryLogResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *WebhookDeliveryLogResource) GetAttributes() WebhookDeliveryLogResourceAttributes {
	if o == nil {
		var ret WebhookDeliveryLogResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResource) GetAttributesOk() (*WebhookDeliveryLogResourceAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookDeliveryLogResource) SetAttributes(v WebhookDeliveryLogResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *WebhookDeliveryLogResource) GetRelationships() WebhookDeliveryLogResourceRelationships {
	if o == nil {
		var ret WebhookDeliveryLogResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResource) GetRelationshipsOk() (*WebhookDeliveryLogResourceRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *WebhookDeliveryLogResource) SetRelationships(v WebhookDeliveryLogResourceRelationships) {
	o.Relationships = v
}

func (o WebhookDeliveryLogResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookDeliveryLogResource struct {
	value *WebhookDeliveryLogResource
	isSet bool
}

func (v NullableWebhookDeliveryLogResource) Get() *WebhookDeliveryLogResource {
	return v.value
}

func (v *NullableWebhookDeliveryLogResource) Set(val *WebhookDeliveryLogResource) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryLogResource) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryLogResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryLogResource(val *WebhookDeliveryLogResource) *NullableWebhookDeliveryLogResource {
	return &NullableWebhookDeliveryLogResource{value: val, isSet: true}
}

func (v NullableWebhookDeliveryLogResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryLogResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


