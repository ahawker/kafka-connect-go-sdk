# GetConnectorTasksResponseId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to **int32** | ID of the task. | [optional] 
**Connector** | Pointer to **string** | Name of the connector. | [optional] 

## Methods

### NewGetConnectorTasksResponseId

`func NewGetConnectorTasksResponseId() *GetConnectorTasksResponseId`

NewGetConnectorTasksResponseId instantiates a new GetConnectorTasksResponseId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectorTasksResponseIdWithDefaults

`func NewGetConnectorTasksResponseIdWithDefaults() *GetConnectorTasksResponseId`

NewGetConnectorTasksResponseIdWithDefaults instantiates a new GetConnectorTasksResponseId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *GetConnectorTasksResponseId) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetConnectorTasksResponseId) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetConnectorTasksResponseId) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetConnectorTasksResponseId) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetConnector

`func (o *GetConnectorTasksResponseId) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *GetConnectorTasksResponseId) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *GetConnectorTasksResponseId) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *GetConnectorTasksResponseId) HasConnector() bool`

HasConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


