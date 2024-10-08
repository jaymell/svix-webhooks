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

// StreamSinkInOneOf2 struct for StreamSinkInOneOf2
type StreamSinkInOneOf2 struct {
	Config SinkHttpConfig `json:"config"`
	Type string `json:"type"`
}

// NewStreamSinkInOneOf2 instantiates a new StreamSinkInOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamSinkInOneOf2(config SinkHttpConfig, type_ string) *StreamSinkInOneOf2 {
	this := StreamSinkInOneOf2{}
	this.Config = config
	this.Type = type_
	return &this
}

// NewStreamSinkInOneOf2WithDefaults instantiates a new StreamSinkInOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamSinkInOneOf2WithDefaults() *StreamSinkInOneOf2 {
	this := StreamSinkInOneOf2{}
	return &this
}

// GetConfig returns the Config field value
func (o *StreamSinkInOneOf2) GetConfig() SinkHttpConfig {
	if o == nil {
		var ret SinkHttpConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *StreamSinkInOneOf2) GetConfigOk() (*SinkHttpConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *StreamSinkInOneOf2) SetConfig(v SinkHttpConfig) {
	o.Config = v
}

// GetType returns the Type field value
func (o *StreamSinkInOneOf2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StreamSinkInOneOf2) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StreamSinkInOneOf2) SetType(v string) {
	o.Type = v
}

func (o StreamSinkInOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableStreamSinkInOneOf2 struct {
	value *StreamSinkInOneOf2
	isSet bool
}

func (v NullableStreamSinkInOneOf2) Get() *StreamSinkInOneOf2 {
	return v.value
}

func (v *NullableStreamSinkInOneOf2) Set(val *StreamSinkInOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSinkInOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSinkInOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSinkInOneOf2(val *StreamSinkInOneOf2) *NullableStreamSinkInOneOf2 {
	return &NullableStreamSinkInOneOf2{value: val, isSet: true}
}

func (v NullableStreamSinkInOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamSinkInOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


