/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// WebhookEventTypeEnum Specifies the type of a webhook event. This can be used to determine what action to take in response to the event, such as which relationships to expect. 
type WebhookEventTypeEnum string

// List of WebhookEventTypeEnum
const (
	TRANSACTION_CREATED WebhookEventTypeEnum = "TRANSACTION_CREATED"
	TRANSACTION_SETTLED WebhookEventTypeEnum = "TRANSACTION_SETTLED"
	TRANSACTION_DELETED WebhookEventTypeEnum = "TRANSACTION_DELETED"
	PING WebhookEventTypeEnum = "PING"
)

var allowedWebhookEventTypeEnumEnumValues = []WebhookEventTypeEnum{
	"TRANSACTION_CREATED",
	"TRANSACTION_SETTLED",
	"TRANSACTION_DELETED",
	"PING",
}

func (v *WebhookEventTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookEventTypeEnum(value)
	for _, existing := range allowedWebhookEventTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookEventTypeEnum", value)
}

// NewWebhookEventTypeEnumFromValue returns a pointer to a valid WebhookEventTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookEventTypeEnumFromValue(v string) (*WebhookEventTypeEnum, error) {
	ev := WebhookEventTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookEventTypeEnum: valid values are %v", v, allowedWebhookEventTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookEventTypeEnum) IsValid() bool {
	for _, existing := range allowedWebhookEventTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookEventTypeEnum value
func (v WebhookEventTypeEnum) Ptr() *WebhookEventTypeEnum {
	return &v
}

type NullableWebhookEventTypeEnum struct {
	value *WebhookEventTypeEnum
	isSet bool
}

func (v NullableWebhookEventTypeEnum) Get() *WebhookEventTypeEnum {
	return v.value
}

func (v *NullableWebhookEventTypeEnum) Set(val *WebhookEventTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventTypeEnum(val *WebhookEventTypeEnum) *NullableWebhookEventTypeEnum {
	return &NullableWebhookEventTypeEnum{value: val, isSet: true}
}

func (v NullableWebhookEventTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

