/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EventTypeSchemaIn struct for EventTypeSchemaIn
type EventTypeSchemaIn struct {
	// The schema for an event type
	Schema map[string]interface{} `json:"schema"`
}

// NewEventTypeSchemaIn instantiates a new EventTypeSchemaIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeSchemaIn(schema map[string]interface{}) *EventTypeSchemaIn {
	this := EventTypeSchemaIn{}
	this.Schema = schema
	return &this
}

// NewEventTypeSchemaInWithDefaults instantiates a new EventTypeSchemaIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeSchemaInWithDefaults() *EventTypeSchemaIn {
	this := EventTypeSchemaIn{}
	return &this
}

// GetSchema returns the Schema field value
func (o *EventTypeSchemaIn) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *EventTypeSchemaIn) GetSchemaOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *EventTypeSchemaIn) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

func (o EventTypeSchemaIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schema"] = o.Schema
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypeSchemaIn struct {
	value *EventTypeSchemaIn
	isSet bool
}

func (v NullableEventTypeSchemaIn) Get() *EventTypeSchemaIn {
	return v.value
}

func (v *NullableEventTypeSchemaIn) Set(val *EventTypeSchemaIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeSchemaIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeSchemaIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeSchemaIn(val *EventTypeSchemaIn) *NullableEventTypeSchemaIn {
	return &NullableEventTypeSchemaIn{value: val, isSet: true}
}

func (v NullableEventTypeSchemaIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeSchemaIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

