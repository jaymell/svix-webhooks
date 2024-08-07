/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EventOut struct for EventOut
type EventOut struct {
	Payload string `json:"payload"`
}

// NewEventOut instantiates a new EventOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventOut(payload string) *EventOut {
	this := EventOut{}
	this.Payload = payload
	return &this
}

// NewEventOutWithDefaults instantiates a new EventOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventOutWithDefaults() *EventOut {
	this := EventOut{}
	return &this
}

// GetPayload returns the Payload field value
func (o *EventOut) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *EventOut) GetPayloadOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *EventOut) SetPayload(v string) {
	o.Payload = v
}

func (o EventOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableEventOut struct {
	value *EventOut
	isSet bool
}

func (v NullableEventOut) Get() *EventOut {
	return v.value
}

func (v *NullableEventOut) Set(val *EventOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEventOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEventOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventOut(val *EventOut) *NullableEventOut {
	return &NullableEventOut{value: val, isSet: true}
}

func (v NullableEventOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


