# TaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the task. | [optional] 
**Trace** | Pointer to **string** | Stack trace information if the task has failed. | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker. | [optional] 

## Methods

### NewTaskStatus

`func NewTaskStatus() *TaskStatus`

NewTaskStatus instantiates a new TaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStatusWithDefaults

`func NewTaskStatusWithDefaults() *TaskStatus`

NewTaskStatusWithDefaults instantiates a new TaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrace

`func (o *TaskStatus) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *TaskStatus) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *TaskStatus) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *TaskStatus) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetState

`func (o *TaskStatus) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskStatus) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskStatus) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *TaskStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkerId

`func (o *TaskStatus) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *TaskStatus) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *TaskStatus) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *TaskStatus) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


