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

// CreateConnectorResponseTasks struct for CreateConnectorResponseTasks
type CreateConnectorResponseTasks struct {
	// Name of the connector.
	Connector *string `json:"connector,omitempty"`
	// ID of the task.
	Task *float32 `json:"task,omitempty"`
}

// NewCreateConnectorResponseTasks instantiates a new CreateConnectorResponseTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectorResponseTasks() *CreateConnectorResponseTasks {
	this := CreateConnectorResponseTasks{}
	return &this
}

// NewCreateConnectorResponseTasksWithDefaults instantiates a new CreateConnectorResponseTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectorResponseTasksWithDefaults() *CreateConnectorResponseTasks {
	this := CreateConnectorResponseTasks{}
	return &this
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *CreateConnectorResponseTasks) GetConnector() string {
	if o == nil || o.Connector == nil {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorResponseTasks) GetConnectorOk() (*string, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *CreateConnectorResponseTasks) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *CreateConnectorResponseTasks) SetConnector(v string) {
	o.Connector = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *CreateConnectorResponseTasks) GetTask() float32 {
	if o == nil || o.Task == nil {
		var ret float32
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorResponseTasks) GetTaskOk() (*float32, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *CreateConnectorResponseTasks) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given float32 and assigns it to the Task field.
func (o *CreateConnectorResponseTasks) SetTask(v float32) {
	o.Task = &v
}

func (o CreateConnectorResponseTasks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	return json.Marshal(toSerialize)
}

type NullableCreateConnectorResponseTasks struct {
	value *CreateConnectorResponseTasks
	isSet bool
}

func (v NullableCreateConnectorResponseTasks) Get() *CreateConnectorResponseTasks {
	return v.value
}

func (v *NullableCreateConnectorResponseTasks) Set(val *CreateConnectorResponseTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectorResponseTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectorResponseTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectorResponseTasks(val *CreateConnectorResponseTasks) *NullableCreateConnectorResponseTasks {
	return &NullableCreateConnectorResponseTasks{value: val, isSet: true}
}

func (v NullableCreateConnectorResponseTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectorResponseTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


