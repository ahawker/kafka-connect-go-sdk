# ListConnectorTasksResponseTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**ListConnectorTasksResponseId**](ListConnectorTasksResponseId.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration of a connector/task/worker. All keys/values should be strings. | [optional] 

## Methods

### NewListConnectorTasksResponseTasks

`func NewListConnectorTasksResponseTasks() *ListConnectorTasksResponseTasks`

NewListConnectorTasksResponseTasks instantiates a new ListConnectorTasksResponseTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorTasksResponseTasksWithDefaults

`func NewListConnectorTasksResponseTasksWithDefaults() *ListConnectorTasksResponseTasks`

NewListConnectorTasksResponseTasksWithDefaults instantiates a new ListConnectorTasksResponseTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConnectorTasksResponseTasks) GetId() ListConnectorTasksResponseId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectorTasksResponseTasks) GetIdOk() (*ListConnectorTasksResponseId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectorTasksResponseTasks) SetId(v ListConnectorTasksResponseId)`

SetId sets Id field to given value.

### HasId

`func (o *ListConnectorTasksResponseTasks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfig

`func (o *ListConnectorTasksResponseTasks) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListConnectorTasksResponseTasks) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListConnectorTasksResponseTasks) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListConnectorTasksResponseTasks) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


