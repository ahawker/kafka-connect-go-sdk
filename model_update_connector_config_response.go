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

// UpdateConnectorConfigResponse Connector info.
type UpdateConnectorConfigResponse struct {
	// Name of the connector.
	Name *string `json:"name,omitempty"`
	// Configuration of a connector/task/worker. All keys/values should be strings.
	Config *map[string]string `json:"config,omitempty"`
	// List of active tasks generated by the connector.
	Tasks []CreateConnectorResponseTasksInner `json:"tasks,omitempty"`
}

// NewUpdateConnectorConfigResponse instantiates a new UpdateConnectorConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnectorConfigResponse() *UpdateConnectorConfigResponse {
	this := UpdateConnectorConfigResponse{}
	return &this
}

// NewUpdateConnectorConfigResponseWithDefaults instantiates a new UpdateConnectorConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnectorConfigResponseWithDefaults() *UpdateConnectorConfigResponse {
	this := UpdateConnectorConfigResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateConnectorConfigResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateConnectorConfigResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateConnectorConfigResponse) SetName(v string) {
	o.Name = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateConnectorConfigResponse) GetConfig() map[string]string {
	if o == nil || o.Config == nil {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigResponse) GetConfigOk() (*map[string]string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateConnectorConfigResponse) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *UpdateConnectorConfigResponse) SetConfig(v map[string]string) {
	o.Config = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *UpdateConnectorConfigResponse) GetTasks() []CreateConnectorResponseTasksInner {
	if o == nil || o.Tasks == nil {
		var ret []CreateConnectorResponseTasksInner
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigResponse) GetTasksOk() ([]CreateConnectorResponseTasksInner, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *UpdateConnectorConfigResponse) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []CreateConnectorResponseTasksInner and assigns it to the Tasks field.
func (o *UpdateConnectorConfigResponse) SetTasks(v []CreateConnectorResponseTasksInner) {
	o.Tasks = v
}

func (o UpdateConnectorConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConnectorConfigResponse struct {
	value *UpdateConnectorConfigResponse
	isSet bool
}

func (v NullableUpdateConnectorConfigResponse) Get() *UpdateConnectorConfigResponse {
	return v.value
}

func (v *NullableUpdateConnectorConfigResponse) Set(val *UpdateConnectorConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnectorConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnectorConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnectorConfigResponse(val *UpdateConnectorConfigResponse) *NullableUpdateConnectorConfigResponse {
	return &NullableUpdateConnectorConfigResponse{value: val, isSet: true}
}

func (v NullableUpdateConnectorConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnectorConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


