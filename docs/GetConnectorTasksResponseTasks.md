# GetConnectorTasksResponseTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**GetConnectorTasksResponseId**](GetConnectorTasksResponseId.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration of a connector/task/worker. All keys/values should be strings. | [optional] 

## Methods

### NewGetConnectorTasksResponseTasks

`func NewGetConnectorTasksResponseTasks() *GetConnectorTasksResponseTasks`

NewGetConnectorTasksResponseTasks instantiates a new GetConnectorTasksResponseTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectorTasksResponseTasksWithDefaults

`func NewGetConnectorTasksResponseTasksWithDefaults() *GetConnectorTasksResponseTasks`

NewGetConnectorTasksResponseTasksWithDefaults instantiates a new GetConnectorTasksResponseTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetConnectorTasksResponseTasks) GetId() GetConnectorTasksResponseId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConnectorTasksResponseTasks) GetIdOk() (*GetConnectorTasksResponseId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConnectorTasksResponseTasks) SetId(v GetConnectorTasksResponseId)`

SetId sets Id field to given value.

### HasId

`func (o *GetConnectorTasksResponseTasks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfig

`func (o *GetConnectorTasksResponseTasks) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetConnectorTasksResponseTasks) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetConnectorTasksResponseTasks) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetConnectorTasksResponseTasks) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


