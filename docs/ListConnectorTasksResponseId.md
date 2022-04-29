# ListConnectorTasksResponseId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to **int32** | ID of the task. | [optional] 
**Connector** | Pointer to **string** | Name of the connector. | [optional] 

## Methods

### NewListConnectorTasksResponseId

`func NewListConnectorTasksResponseId() *ListConnectorTasksResponseId`

NewListConnectorTasksResponseId instantiates a new ListConnectorTasksResponseId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorTasksResponseIdWithDefaults

`func NewListConnectorTasksResponseIdWithDefaults() *ListConnectorTasksResponseId`

NewListConnectorTasksResponseIdWithDefaults instantiates a new ListConnectorTasksResponseId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *ListConnectorTasksResponseId) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ListConnectorTasksResponseId) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ListConnectorTasksResponseId) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *ListConnectorTasksResponseId) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetConnector

`func (o *ListConnectorTasksResponseId) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ListConnectorTasksResponseId) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ListConnectorTasksResponseId) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ListConnectorTasksResponseId) HasConnector() bool`

HasConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


