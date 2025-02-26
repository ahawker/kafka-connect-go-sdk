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

// GetConnectorStatusResponse Connector status.
type GetConnectorStatusResponse struct {
	// Name of the connector.
	Name *string `json:"name,omitempty"`
	Connector *ConnectorStatus `json:"connector,omitempty"`
	// List of task status for the connector.
	Tasks []TaskStatus `json:"tasks,omitempty"`
}

// NewGetConnectorStatusResponse instantiates a new GetConnectorStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectorStatusResponse() *GetConnectorStatusResponse {
	this := GetConnectorStatusResponse{}
	return &this
}

// NewGetConnectorStatusResponseWithDefaults instantiates a new GetConnectorStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectorStatusResponseWithDefaults() *GetConnectorStatusResponse {
	this := GetConnectorStatusResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetConnectorStatusResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorStatusResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetConnectorStatusResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetConnectorStatusResponse) SetName(v string) {
	o.Name = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *GetConnectorStatusResponse) GetConnector() ConnectorStatus {
	if o == nil || o.Connector == nil {
		var ret ConnectorStatus
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorStatusResponse) GetConnectorOk() (*ConnectorStatus, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *GetConnectorStatusResponse) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given ConnectorStatus and assigns it to the Connector field.
func (o *GetConnectorStatusResponse) SetConnector(v ConnectorStatus) {
	o.Connector = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *GetConnectorStatusResponse) GetTasks() []TaskStatus {
	if o == nil || o.Tasks == nil {
		var ret []TaskStatus
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorStatusResponse) GetTasksOk() ([]TaskStatus, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *GetConnectorStatusResponse) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []TaskStatus and assigns it to the Tasks field.
func (o *GetConnectorStatusResponse) SetTasks(v []TaskStatus) {
	o.Tasks = v
}

func (o GetConnectorStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnectorStatusResponse struct {
	value *GetConnectorStatusResponse
	isSet bool
}

func (v NullableGetConnectorStatusResponse) Get() *GetConnectorStatusResponse {
	return v.value
}

func (v *NullableGetConnectorStatusResponse) Set(val *GetConnectorStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectorStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectorStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectorStatusResponse(val *GetConnectorStatusResponse) *NullableGetConnectorStatusResponse {
	return &NullableGetConnectorStatusResponse{value: val, isSet: true}
}

func (v NullableGetConnectorStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectorStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


