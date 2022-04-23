# ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Connect worker version. | [optional] 
**Commit** | Pointer to **string** | git commit ID. | [optional] 
**KafkaClusterId** | Pointer to **string** | Kafka cluster ID. | [optional] 

## Methods

### NewClusterInfo

`func NewClusterInfo() *ClusterInfo`

NewClusterInfo instantiates a new ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoWithDefaults

`func NewClusterInfoWithDefaults() *ClusterInfo`

NewClusterInfoWithDefaults instantiates a new ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ClusterInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommit

`func (o *ClusterInfo) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ClusterInfo) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ClusterInfo) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ClusterInfo) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetKafkaClusterId

`func (o *ClusterInfo) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *ClusterInfo) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *ClusterInfo) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.

### HasKafkaClusterId

`func (o *ClusterInfo) HasKafkaClusterId() bool`

HasKafkaClusterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


