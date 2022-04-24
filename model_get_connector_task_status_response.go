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

// GetConnectorTaskStatusResponse Task status.
type GetConnectorTaskStatusResponse struct {
	State *State `json:"state,omitempty"`
	// ID of the worker.
	WorkerId *string `json:"worker_id,omitempty"`
}

// NewGetConnectorTaskStatusResponse instantiates a new GetConnectorTaskStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectorTaskStatusResponse() *GetConnectorTaskStatusResponse {
	this := GetConnectorTaskStatusResponse{}
	return &this
}

// NewGetConnectorTaskStatusResponseWithDefaults instantiates a new GetConnectorTaskStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectorTaskStatusResponseWithDefaults() *GetConnectorTaskStatusResponse {
	this := GetConnectorTaskStatusResponse{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetConnectorTaskStatusResponse) GetState() State {
	if o == nil || o.State == nil {
		var ret State
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorTaskStatusResponse) GetStateOk() (*State, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetConnectorTaskStatusResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given State and assigns it to the State field.
func (o *GetConnectorTaskStatusResponse) SetState(v State) {
	o.State = &v
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise.
func (o *GetConnectorTaskStatusResponse) GetWorkerId() string {
	if o == nil || o.WorkerId == nil {
		var ret string
		return ret
	}
	return *o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectorTaskStatusResponse) GetWorkerIdOk() (*string, bool) {
	if o == nil || o.WorkerId == nil {
		return nil, false
	}
	return o.WorkerId, true
}

// HasWorkerId returns a boolean if a field has been set.
func (o *GetConnectorTaskStatusResponse) HasWorkerId() bool {
	if o != nil && o.WorkerId != nil {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given string and assigns it to the WorkerId field.
func (o *GetConnectorTaskStatusResponse) SetWorkerId(v string) {
	o.WorkerId = &v
}

func (o GetConnectorTaskStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.WorkerId != nil {
		toSerialize["worker_id"] = o.WorkerId
	}
	return json.Marshal(toSerialize)
}

type NullableGetConnectorTaskStatusResponse struct {
	value *GetConnectorTaskStatusResponse
	isSet bool
}

func (v NullableGetConnectorTaskStatusResponse) Get() *GetConnectorTaskStatusResponse {
	return v.value
}

func (v *NullableGetConnectorTaskStatusResponse) Set(val *GetConnectorTaskStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectorTaskStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectorTaskStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectorTaskStatusResponse(val *GetConnectorTaskStatusResponse) *NullableGetConnectorTaskStatusResponse {
	return &NullableGetConnectorTaskStatusResponse{value: val, isSet: true}
}

func (v NullableGetConnectorTaskStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectorTaskStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


