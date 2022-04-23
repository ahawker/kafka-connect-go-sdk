# InlineResponse200Tasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | ID of the task. | [optional] 
**State** | Pointer to **string** | Current state of the task. | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker running the task. | [optional] 
**Trace** | Pointer to **string** | Stack trace information if the task has failed. | [optional] 

## Methods

### NewInlineResponse200Tasks

`func NewInlineResponse200Tasks() *InlineResponse200Tasks`

NewInlineResponse200Tasks instantiates a new InlineResponse200Tasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200TasksWithDefaults

`func NewInlineResponse200TasksWithDefaults() *InlineResponse200Tasks`

NewInlineResponse200TasksWithDefaults instantiates a new InlineResponse200Tasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200Tasks) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200Tasks) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200Tasks) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200Tasks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *InlineResponse200Tasks) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse200Tasks) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse200Tasks) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InlineResponse200Tasks) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkerId

`func (o *InlineResponse200Tasks) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *InlineResponse200Tasks) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *InlineResponse200Tasks) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *InlineResponse200Tasks) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetTrace

`func (o *InlineResponse200Tasks) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *InlineResponse200Tasks) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *InlineResponse200Tasks) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *InlineResponse200Tasks) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


