# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Connector** | Pointer to [**InlineResponse200Connector**](InlineResponse200Connector.md) |  | [optional] 
**Tasks** | Pointer to [**[]InlineResponse200Tasks**](InlineResponse200Tasks.md) | List of task status for the connector. | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnector

`func (o *InlineResponse200) GetConnector() InlineResponse200Connector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *InlineResponse200) GetConnectorOk() (*InlineResponse200Connector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *InlineResponse200) SetConnector(v InlineResponse200Connector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *InlineResponse200) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTasks

`func (o *InlineResponse200) GetTasks() []InlineResponse200Tasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InlineResponse200) GetTasksOk() (*[]InlineResponse200Tasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InlineResponse200) SetTasks(v []InlineResponse200Tasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InlineResponse200) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


