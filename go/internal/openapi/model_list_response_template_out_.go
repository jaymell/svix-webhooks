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

// ListResponseTemplateOut struct for ListResponseTemplateOut
type ListResponseTemplateOut struct {
	Data []TemplateOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator,omitempty"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

// NewListResponseTemplateOut instantiates a new ListResponseTemplateOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseTemplateOut(data []TemplateOut, done bool) *ListResponseTemplateOut {
	this := ListResponseTemplateOut{}
	this.Data = data
	this.Done = done
	return &this
}

// NewListResponseTemplateOutWithDefaults instantiates a new ListResponseTemplateOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseTemplateOutWithDefaults() *ListResponseTemplateOut {
	this := ListResponseTemplateOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseTemplateOut) GetData() []TemplateOut {
	if o == nil {
		var ret []TemplateOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseTemplateOut) GetDataOk() (*[]TemplateOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListResponseTemplateOut) SetData(v []TemplateOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseTemplateOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseTemplateOut) GetDoneOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseTemplateOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseTemplateOut) GetIterator() string {
	if o == nil || o.Iterator.Get() == nil {
		var ret string
		return ret
	}
	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseTemplateOut) GetIteratorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// HasIterator returns a boolean if a field has been set.
func (o *ListResponseTemplateOut) HasIterator() bool {
	if o != nil && o.Iterator.IsSet() {
		return true
	}

	return false
}

// SetIterator gets a reference to the given NullableString and assigns it to the Iterator field.
func (o *ListResponseTemplateOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}
// SetIteratorNil sets the value for Iterator to be an explicit nil
func (o *ListResponseTemplateOut) SetIteratorNil() {
	o.Iterator.Set(nil)
}

// UnsetIterator ensures that no value is present for Iterator, not even an explicit nil
func (o *ListResponseTemplateOut) UnsetIterator() {
	o.Iterator.Unset()
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseTemplateOut) GetPrevIterator() string {
	if o == nil || o.PrevIterator.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseTemplateOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseTemplateOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseTemplateOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseTemplateOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseTemplateOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseTemplateOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["done"] = o.Done
	}
	if o.Iterator.IsSet() {
		toSerialize["iterator"] = o.Iterator.Get()
	}
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListResponseTemplateOut struct {
	value *ListResponseTemplateOut
	isSet bool
}

func (v NullableListResponseTemplateOut) Get() *ListResponseTemplateOut {
	return v.value
}

func (v *NullableListResponseTemplateOut) Set(val *ListResponseTemplateOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseTemplateOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseTemplateOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseTemplateOut(val *ListResponseTemplateOut) *NullableListResponseTemplateOut {
	return &NullableListResponseTemplateOut{value: val, isSet: true}
}

func (v NullableListResponseTemplateOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseTemplateOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

