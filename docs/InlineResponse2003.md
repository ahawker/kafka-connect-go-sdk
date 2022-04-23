# InlineResponse2003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Connector** | Pointer to [**InlineResponse2003Connector**](InlineResponse2003Connector.md) |  | [optional] 
**Tasks** | Pointer to [**[]InlineResponse2003Tasks**](InlineResponse2003Tasks.md) | List of task status for the connector. | [optional] 

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003() *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2003) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2003) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2003) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2003) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnector

`func (o *InlineResponse2003) GetConnector() InlineResponse2003Connector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *InlineResponse2003) GetConnectorOk() (*InlineResponse2003Connector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *InlineResponse2003) SetConnector(v InlineResponse2003Connector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *InlineResponse2003) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTasks

`func (o *InlineResponse2003) GetTasks() []InlineResponse2003Tasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InlineResponse2003) GetTasksOk() (*[]InlineResponse2003Tasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InlineResponse2003) SetTasks(v []InlineResponse2003Tasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InlineResponse2003) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


