# GetConnectorStatusResponseTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | ID of the task. | [optional] 
**Trace** | Pointer to **string** | Stack trace information if the task has failed. | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker. | [optional] 

## Methods

### NewGetConnectorStatusResponseTasks

`func NewGetConnectorStatusResponseTasks() *GetConnectorStatusResponseTasks`

NewGetConnectorStatusResponseTasks instantiates a new GetConnectorStatusResponseTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectorStatusResponseTasksWithDefaults

`func NewGetConnectorStatusResponseTasksWithDefaults() *GetConnectorStatusResponseTasks`

NewGetConnectorStatusResponseTasksWithDefaults instantiates a new GetConnectorStatusResponseTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetConnectorStatusResponseTasks) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConnectorStatusResponseTasks) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConnectorStatusResponseTasks) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetConnectorStatusResponseTasks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrace

`func (o *GetConnectorStatusResponseTasks) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *GetConnectorStatusResponseTasks) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *GetConnectorStatusResponseTasks) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *GetConnectorStatusResponseTasks) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetState

`func (o *GetConnectorStatusResponseTasks) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetConnectorStatusResponseTasks) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetConnectorStatusResponseTasks) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *GetConnectorStatusResponseTasks) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkerId

`func (o *GetConnectorStatusResponseTasks) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *GetConnectorStatusResponseTasks) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *GetConnectorStatusResponseTasks) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *GetConnectorStatusResponseTasks) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


