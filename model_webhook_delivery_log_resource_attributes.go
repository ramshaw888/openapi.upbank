/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WebhookDeliveryLogResourceAttributes struct for WebhookDeliveryLogResourceAttributes
type WebhookDeliveryLogResourceAttributes struct {
	Request WebhookDeliveryLogResourceAttributesRequest `json:"request"`
	Response NullableWebhookDeliveryLogResourceAttributesResponse `json:"response"`
	// The success or failure status of this delivery attempt. 
	DeliveryStatus WebhookDeliveryStatusEnum `json:"deliveryStatus"`
	// The date-time at which this log entry was created. 
	CreatedAt time.Time `json:"createdAt"`
}

// NewWebhookDeliveryLogResourceAttributes instantiates a new WebhookDeliveryLogResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDeliveryLogResourceAttributes(request WebhookDeliveryLogResourceAttributesRequest, response NullableWebhookDeliveryLogResourceAttributesResponse, deliveryStatus WebhookDeliveryStatusEnum, createdAt time.Time) *WebhookDeliveryLogResourceAttributes {
	this := WebhookDeliveryLogResourceAttributes{}
	this.Request = request
	this.Response = response
	this.DeliveryStatus = deliveryStatus
	this.CreatedAt = createdAt
	return &this
}

// NewWebhookDeliveryLogResourceAttributesWithDefaults instantiates a new WebhookDeliveryLogResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryLogResourceAttributesWithDefaults() *WebhookDeliveryLogResourceAttributes {
	this := WebhookDeliveryLogResourceAttributes{}
	return &this
}

// GetRequest returns the Request field value
func (o *WebhookDeliveryLogResourceAttributes) GetRequest() WebhookDeliveryLogResourceAttributesRequest {
	if o == nil {
		var ret WebhookDeliveryLogResourceAttributesRequest
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResourceAttributes) GetRequestOk() (*WebhookDeliveryLogResourceAttributesRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *WebhookDeliveryLogResourceAttributes) SetRequest(v WebhookDeliveryLogResourceAttributesRequest) {
	o.Request = v
}

// GetResponse returns the Response field value
// If the value is explicit nil, the zero value for WebhookDeliveryLogResourceAttributesResponse will be returned
func (o *WebhookDeliveryLogResourceAttributes) GetResponse() WebhookDeliveryLogResourceAttributesResponse {
	if o == nil || o.Response.Get() == nil {
		var ret WebhookDeliveryLogResourceAttributesResponse
		return ret
	}

	return *o.Response.Get()
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDeliveryLogResourceAttributes) GetResponseOk() (*WebhookDeliveryLogResourceAttributesResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Response.Get(), o.Response.IsSet()
}

// SetResponse sets field value
func (o *WebhookDeliveryLogResourceAttributes) SetResponse(v WebhookDeliveryLogResourceAttributesResponse) {
	o.Response.Set(&v)
}

// GetDeliveryStatus returns the DeliveryStatus field value
// If the value is explicit nil, the zero value for WebhookDeliveryStatusEnum will be returned
func (o *WebhookDeliveryLogResourceAttributes) GetDeliveryStatus() WebhookDeliveryStatusEnum {
	if o == nil {
		var ret WebhookDeliveryStatusEnum
		return ret
	}

	return o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDeliveryLogResourceAttributes) GetDeliveryStatusOk() (*WebhookDeliveryStatusEnum, bool) {
	if o == nil || o.DeliveryStatus == nil {
		return nil, false
	}
	return &o.DeliveryStatus, true
}

// SetDeliveryStatus sets field value
func (o *WebhookDeliveryLogResourceAttributes) SetDeliveryStatus(v WebhookDeliveryStatusEnum) {
	o.DeliveryStatus = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookDeliveryLogResourceAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResourceAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookDeliveryLogResourceAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o WebhookDeliveryLogResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request"] = o.Request
	}
	if true {
		toSerialize["response"] = o.Response.Get()
	}
	if o.DeliveryStatus != nil {
		toSerialize["deliveryStatus"] = o.DeliveryStatus
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookDeliveryLogResourceAttributes struct {
	value *WebhookDeliveryLogResourceAttributes
	isSet bool
}

func (v NullableWebhookDeliveryLogResourceAttributes) Get() *WebhookDeliveryLogResourceAttributes {
	return v.value
}

func (v *NullableWebhookDeliveryLogResourceAttributes) Set(val *WebhookDeliveryLogResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryLogResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryLogResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryLogResourceAttributes(val *WebhookDeliveryLogResourceAttributes) *NullableWebhookDeliveryLogResourceAttributes {
	return &NullableWebhookDeliveryLogResourceAttributes{value: val, isSet: true}
}

func (v NullableWebhookDeliveryLogResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryLogResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


