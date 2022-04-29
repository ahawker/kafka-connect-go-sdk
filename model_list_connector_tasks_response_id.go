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

// ListConnectorTasksResponseId Task ID.
type ListConnectorTasksResponseId struct {
	// ID of the task.
	Task *int32 `json:"task,omitempty"`
	// Name of the connector.
	Connector *string `json:"connector,omitempty"`
}

// NewListConnectorTasksResponseId instantiates a new ListConnectorTasksResponseId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorTasksResponseId() *ListConnectorTasksResponseId {
	this := ListConnectorTasksResponseId{}
	return &this
}

// NewListConnectorTasksResponseIdWithDefaults instantiates a new ListConnectorTasksResponseId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorTasksResponseIdWithDefaults() *ListConnectorTasksResponseId {
	this := ListConnectorTasksResponseId{}
	return &this
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *ListConnectorTasksResponseId) GetTask() int32 {
	if o == nil || o.Task == nil {
		var ret int32
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorTasksResponseId) GetTaskOk() (*int32, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *ListConnectorTasksResponseId) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given int32 and assigns it to the Task field.
func (o *ListConnectorTasksResponseId) SetTask(v int32) {
	o.Task = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *ListConnectorTasksResponseId) GetConnector() string {
	if o == nil || o.Connector == nil {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorTasksResponseId) GetConnectorOk() (*string, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *ListConnectorTasksResponseId) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *ListConnectorTasksResponseId) SetConnector(v string) {
	o.Connector = &v
}

func (o ListConnectorTasksResponseId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectorTasksResponseId struct {
	value *ListConnectorTasksResponseId
	isSet bool
}

func (v NullableListConnectorTasksResponseId) Get() *ListConnectorTasksResponseId {
	return v.value
}

func (v *NullableListConnectorTasksResponseId) Set(val *ListConnectorTasksResponseId) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorTasksResponseId) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorTasksResponseId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorTasksResponseId(val *ListConnectorTasksResponseId) *NullableListConnectorTasksResponseId {
	return &NullableListConnectorTasksResponseId{value: val, isSet: true}
}

func (v NullableListConnectorTasksResponseId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorTasksResponseId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


