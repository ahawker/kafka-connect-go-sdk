/*
Kafka Connect REST API

Kafka Connect REST API https://docs.confluent.io/platform/current/connect/references/restapi.html

API version: 0.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetConnectorTasksResponseId Task ID.
type GetConnectorTasksResponseId struct {
	// ID of the task.
	Task *int32 `json:"task,omitempty"`
	// Name of the connector.
	Connector *string `json:"connector,omitempty"`
}

// NewGetConnectorTasksResponseId instantiates a new GetConnectorTasksResponseId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectorTasksResponseId() *GetConnectorTasksResponseId {
	this := GetConnectorTasksResponseId{}
	return &this
}

// NewGetConnectorTasksResponseIdWithDefaults instantiates a new GetConnectorTasksResponseId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectorTasksResponseIdWithDefaults() *GetConnectorTasksResponseId {
	this := GetConnectorTasksResponseId{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *GetConnectorTasksResponseId) GetTask() int32 {
	if o == nil || o.Task == nil {
		var ret int32
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorTasksResponseId) GetTaskOk() (*int32, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *GetConnectorTasksResponseId) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given int32 and assigns it to the Task field.
func (o *GetConnectorTasksResponseId) SetTask(v int32) {
	o.Task = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *GetConnectorTasksResponseId) GetConnector() string {
	if o == nil || o.Connector == nil {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorTasksResponseId) GetConnectorOk() (*string, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *GetConnectorTasksResponseId) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *GetConnectorTasksResponseId) SetConnector(v string) {
	o.Connector = &v
}

func (o GetConnectorTasksResponseId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnectorTasksResponseId struct {
	value *GetConnectorTasksResponseId
	isSet bool
}

func (v NullableGetConnectorTasksResponseId) Get() *GetConnectorTasksResponseId {
	return v.value
}

func (v *NullableGetConnectorTasksResponseId) Set(val *GetConnectorTasksResponseId) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectorTasksResponseId) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectorTasksResponseId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectorTasksResponseId(val *GetConnectorTasksResponseId) *NullableGetConnectorTasksResponseId {
	return &NullableGetConnectorTasksResponseId{value: val, isSet: true}
}

func (v NullableGetConnectorTasksResponseId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectorTasksResponseId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


