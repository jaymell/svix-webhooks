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

// SinkInOneOf1 struct for SinkInOneOf1
type SinkInOneOf1 struct {
	AccessKey string `json:"accessKey"`
	QueueDsn string `json:"queueDsn"`
	Region string `json:"region"`
	SecretKey string `json:"secretKey"`
	Type string `json:"type"`
}

// NewSinkInOneOf1 instantiates a new SinkInOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinkInOneOf1(accessKey string, queueDsn string, region string, secretKey string, type_ string) *SinkInOneOf1 {
	this := SinkInOneOf1{}
	this.AccessKey = accessKey
	this.QueueDsn = queueDsn
	this.Region = region
	this.SecretKey = secretKey
	this.Type = type_
	return &this
}

// NewSinkInOneOf1WithDefaults instantiates a new SinkInOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinkInOneOf1WithDefaults() *SinkInOneOf1 {
	this := SinkInOneOf1{}
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *SinkInOneOf1) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *SinkInOneOf1) GetAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *SinkInOneOf1) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetQueueDsn returns the QueueDsn field value
func (o *SinkInOneOf1) GetQueueDsn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueueDsn
}

// GetQueueDsnOk returns a tuple with the QueueDsn field value
// and a boolean to check if the value has been set.
func (o *SinkInOneOf1) GetQueueDsnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QueueDsn, true
}

// SetQueueDsn sets field value
func (o *SinkInOneOf1) SetQueueDsn(v string) {
	o.QueueDsn = v
}

// GetRegion returns the Region field value
func (o *SinkInOneOf1) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *SinkInOneOf1) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *SinkInOneOf1) SetRegion(v string) {
	o.Region = v
}

// GetSecretKey returns the SecretKey field value
func (o *SinkInOneOf1) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *SinkInOneOf1) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *SinkInOneOf1) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetType returns the Type field value
func (o *SinkInOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SinkInOneOf1) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SinkInOneOf1) SetType(v string) {
	o.Type = v
}

func (o SinkInOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessKey"] = o.AccessKey
	}
	if true {
		toSerialize["queueDsn"] = o.QueueDsn
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["secretKey"] = o.SecretKey
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSinkInOneOf1 struct {
	value *SinkInOneOf1
	isSet bool
}

func (v NullableSinkInOneOf1) Get() *SinkInOneOf1 {
	return v.value
}

func (v *NullableSinkInOneOf1) Set(val *SinkInOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSinkInOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSinkInOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinkInOneOf1(val *SinkInOneOf1) *NullableSinkInOneOf1 {
	return &NullableSinkInOneOf1{value: val, isSet: true}
}

func (v NullableSinkInOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinkInOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

