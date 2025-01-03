/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the EndpointMessageOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointMessageOut{}

// EndpointMessageOut A model containing information on a given message plus additional fields on the last attempt for that message.
type EndpointMessageOut struct {
	// List of free-form identifiers that endpoints can filter by
	Channels []string `json:"channels,omitempty"`
	// Optional unique identifier for the message
	EventId *string `json:"eventId,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The event type's name
	EventType string `json:"eventType" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The msg's ID
	Id string `json:"id"`
	NextAttempt *time.Time `json:"nextAttempt,omitempty"`
	Payload map[string]interface{} `json:"payload"`
	Status MessageStatus `json:"status"`
	Tags []string `json:"tags,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type _EndpointMessageOut EndpointMessageOut

// NewEndpointMessageOut instantiates a new EndpointMessageOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointMessageOut(eventType string, id string, payload map[string]interface{}, status MessageStatus, timestamp time.Time) *EndpointMessageOut {
	this := EndpointMessageOut{}
	this.EventType = eventType
	this.Id = id
	this.Payload = payload
	this.Status = status
	this.Timestamp = timestamp
	return &this
}

// NewEndpointMessageOutWithDefaults instantiates a new EndpointMessageOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointMessageOutWithDefaults() *EndpointMessageOut {
	this := EndpointMessageOut{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *EndpointMessageOut) GetChannels() []string {
	if o == nil || IsNil(o.Channels) {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *EndpointMessageOut) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *EndpointMessageOut) SetChannels(v []string) {
	o.Channels = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EndpointMessageOut) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EndpointMessageOut) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EndpointMessageOut) SetEventId(v string) {
	o.EventId = &v
}

// GetEventType returns the EventType field value
func (o *EndpointMessageOut) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EndpointMessageOut) SetEventType(v string) {
	o.EventType = v
}

// GetId returns the Id field value
func (o *EndpointMessageOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndpointMessageOut) SetId(v string) {
	o.Id = v
}

// GetNextAttempt returns the NextAttempt field value if set, zero value otherwise.
func (o *EndpointMessageOut) GetNextAttempt() time.Time {
	if o == nil || IsNil(o.NextAttempt) {
		var ret time.Time
		return ret
	}
	return *o.NextAttempt
}

// GetNextAttemptOk returns a tuple with the NextAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetNextAttemptOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextAttempt) {
		return nil, false
	}
	return o.NextAttempt, true
}

// HasNextAttempt returns a boolean if a field has been set.
func (o *EndpointMessageOut) HasNextAttempt() bool {
	if o != nil && !IsNil(o.NextAttempt) {
		return true
	}

	return false
}

// SetNextAttempt gets a reference to the given time.Time and assigns it to the NextAttempt field.
func (o *EndpointMessageOut) SetNextAttempt(v time.Time) {
	o.NextAttempt = &v
}

// GetPayload returns the Payload field value
func (o *EndpointMessageOut) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *EndpointMessageOut) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetStatus returns the Status field value
func (o *EndpointMessageOut) GetStatus() MessageStatus {
	if o == nil {
		var ret MessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetStatusOk() (*MessageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EndpointMessageOut) SetStatus(v MessageStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EndpointMessageOut) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EndpointMessageOut) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EndpointMessageOut) SetTags(v []string) {
	o.Tags = v
}

// GetTimestamp returns the Timestamp field value
func (o *EndpointMessageOut) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EndpointMessageOut) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o EndpointMessageOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointMessageOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	toSerialize["eventType"] = o.EventType
	toSerialize["id"] = o.Id
	if !IsNil(o.NextAttempt) {
		toSerialize["nextAttempt"] = o.NextAttempt
	}
	toSerialize["payload"] = o.Payload
	toSerialize["status"] = o.Status
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *EndpointMessageOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventType",
		"id",
		"payload",
		"status",
		"timestamp",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEndpointMessageOut := _EndpointMessageOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointMessageOut)

	if err != nil {
		return err
	}

	*o = EndpointMessageOut(varEndpointMessageOut)

	return err
}

type NullableEndpointMessageOut struct {
	value *EndpointMessageOut
	isSet bool
}

func (v NullableEndpointMessageOut) Get() *EndpointMessageOut {
	return v.value
}

func (v *NullableEndpointMessageOut) Set(val *EndpointMessageOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointMessageOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointMessageOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointMessageOut(val *EndpointMessageOut) *NullableEndpointMessageOut {
	return &NullableEndpointMessageOut{value: val, isSet: true}
}

func (v NullableEndpointMessageOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointMessageOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


