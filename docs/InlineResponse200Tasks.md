# InlineResponse200Tasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | The name of the connector the task belongs to. | [optional] 
**Task** | Pointer to **float32** | Task ID within the connector. | [optional] 

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

### GetConnector

`func (o *InlineResponse200Tasks) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *InlineResponse200Tasks) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *InlineResponse200Tasks) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *InlineResponse200Tasks) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTask

`func (o *InlineResponse200Tasks) GetTask() float32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *InlineResponse200Tasks) GetTaskOk() (*float32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *InlineResponse200Tasks) SetTask(v float32)`

SetTask sets Task field to given value.

### HasTask

`func (o *InlineResponse200Tasks) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


