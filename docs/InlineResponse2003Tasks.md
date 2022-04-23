# InlineResponse2003Tasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | ID of the task. | [optional] 
**State** | Pointer to **string** | Current state of the task. | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker running the task. | [optional] 
**Trace** | Pointer to **string** | Stack trace information if the task has failed. | [optional] 

## Methods

### NewInlineResponse2003Tasks

`func NewInlineResponse2003Tasks() *InlineResponse2003Tasks`

NewInlineResponse2003Tasks instantiates a new InlineResponse2003Tasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003TasksWithDefaults

`func NewInlineResponse2003TasksWithDefaults() *InlineResponse2003Tasks`

NewInlineResponse2003TasksWithDefaults instantiates a new InlineResponse2003Tasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2003Tasks) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2003Tasks) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2003Tasks) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2003Tasks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *InlineResponse2003Tasks) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse2003Tasks) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse2003Tasks) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InlineResponse2003Tasks) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkerId

`func (o *InlineResponse2003Tasks) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *InlineResponse2003Tasks) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *InlineResponse2003Tasks) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *InlineResponse2003Tasks) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetTrace

`func (o *InlineResponse2003Tasks) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *InlineResponse2003Tasks) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *InlineResponse2003Tasks) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *InlineResponse2003Tasks) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


