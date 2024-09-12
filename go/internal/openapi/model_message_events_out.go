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

// MessageEventsOut struct for MessageEventsOut
type MessageEventsOut struct {
	Data []MessageOut `json:"data"`
	Done bool `json:"done"`
	Iterator string `json:"iterator"`
}

// NewMessageEventsOut instantiates a new MessageEventsOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEventsOut(data []MessageOut, done bool, iterator string) *MessageEventsOut {
	this := MessageEventsOut{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewMessageEventsOutWithDefaults instantiates a new MessageEventsOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEventsOutWithDefaults() *MessageEventsOut {
	this := MessageEventsOut{}
	return &this
}

// GetData returns the Data field value
func (o *MessageEventsOut) GetData() []MessageOut {
	if o == nil {
		var ret []MessageOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MessageEventsOut) GetDataOk() (*[]MessageOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MessageEventsOut) SetData(v []MessageOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *MessageEventsOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *MessageEventsOut) GetDoneOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *MessageEventsOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
func (o *MessageEventsOut) GetIterator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iterator
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
func (o *MessageEventsOut) GetIteratorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Iterator, true
}

// SetIterator sets field value
func (o *MessageEventsOut) SetIterator(v string) {
	o.Iterator = v
}

func (o MessageEventsOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["done"] = o.Done
	}
	if true {
		toSerialize["iterator"] = o.Iterator
	}
	return json.Marshal(toSerialize)
}

type NullableMessageEventsOut struct {
	value *MessageEventsOut
	isSet bool
}

func (v NullableMessageEventsOut) Get() *MessageEventsOut {
	return v.value
}

func (v *NullableMessageEventsOut) Set(val *MessageEventsOut) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEventsOut) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEventsOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEventsOut(val *MessageEventsOut) *NullableMessageEventsOut {
	return &NullableMessageEventsOut{value: val, isSet: true}
}

func (v NullableMessageEventsOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEventsOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

