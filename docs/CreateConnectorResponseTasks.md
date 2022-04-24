# CreateConnectorResponseTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Name of the connector. | [optional] 
**Task** | Pointer to **int32** | ID of the task. | [optional] 

## Methods

### NewCreateConnectorResponseTasks

`func NewCreateConnectorResponseTasks() *CreateConnectorResponseTasks`

NewCreateConnectorResponseTasks instantiates a new CreateConnectorResponseTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorResponseTasksWithDefaults

`func NewCreateConnectorResponseTasksWithDefaults() *CreateConnectorResponseTasks`

NewCreateConnectorResponseTasksWithDefaults instantiates a new CreateConnectorResponseTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *CreateConnectorResponseTasks) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CreateConnectorResponseTasks) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CreateConnectorResponseTasks) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CreateConnectorResponseTasks) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTask

`func (o *CreateConnectorResponseTasks) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *CreateConnectorResponseTasks) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *CreateConnectorResponseTasks) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *CreateConnectorResponseTasks) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


