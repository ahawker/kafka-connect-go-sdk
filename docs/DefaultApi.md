# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](DefaultApi.md#CreateConnector) | **Post** /connectors | Create a new connector, returning the current connector info if successful. Return 409 (Conflict) if rebalance is in process, or if the connector already exists.
[**DeleteConnector**](DefaultApi.md#DeleteConnector) | **Delete** /connectors/{name} | Delete a connector, halting all tasks and deleting its configuration.
[**GetClusterInfo**](DefaultApi.md#GetClusterInfo) | **Get** / | Top-level (root) request that gets the version of the Connect worker that serves the REST request, the git commit ID of the source code, and the Kafka cluster ID that the worker is connected to.
[**GetConnector**](DefaultApi.md#GetConnector) | **Get** /connectors/{name} | Get information about the connector.
[**GetConnectorStatus**](DefaultApi.md#GetConnectorStatus) | **Get** /connectors/{name}/status | Gets the current status of the connector, including: * whether it is running or restarting, or if it has failed or paused * which worker it is assigned to * error information if it has failed * the state of all its tasks 
[**ListConnectors**](DefaultApi.md#ListConnectors) | **Get** /connectors | Get a list of active connectors.
[**PauseConnector**](DefaultApi.md#PauseConnector) | **Put** /connectors/{name}/pause | Pause the connector and its tasks, which stops message processing until the connector is resumed. This call asynchronous and the tasks will not transition to PAUSED state at the same time.
[**RestartConnector**](DefaultApi.md#RestartConnector) | **Post** /connectors/{name}/restart | Restart the connector. You may use the following query parameters to restart any combination of the Connector and/or Task instances for the connector.
[**ResumeConnector**](DefaultApi.md#ResumeConnector) | **Put** /connectors/{name}/resume | Resume a paused connector or do nothing if the connector is not paused. This call asynchronous and the tasks will not transition to RUNNING state at the same time.



## CreateConnector

> InlineResponse200 CreateConnector(ctx).InlineObject(inlineObject).Execute()

Create a new connector, returning the current connector info if successful. Return 409 (Conflict) if rebalance is in process, or if the connector already exists.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject := *openapiclient.NewInlineObject() // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateConnector(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnector`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnector

> DeleteConnector(ctx, name).Execute()

Delete a connector, halting all tasks and deleting its configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the created connector.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteConnector(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the created connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterInfo

> ClusterInfo GetClusterInfo(ctx).Execute()

Top-level (root) request that gets the version of the Connect worker that serves the REST request, the git commit ID of the source code, and the Kafka cluster ID that the worker is connected to.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetClusterInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterInfo`: ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusterInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterInfoRequest struct via the builder pattern


### Return type

[**ClusterInfo**](ClusterInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnector

> InlineResponse2001 GetConnector(ctx, name).Execute()

Get information about the connector.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the created connector.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConnector(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnector`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the created connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorStatus

> InlineResponse2002 GetConnectorStatus(ctx, name).Execute()

Gets the current status of the connector, including: * whether it is running or restarting, or if it has failed or paused * which worker it is assigned to * error information if it has failed * the state of all its tasks 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the created connector.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetConnectorStatus(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnectorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorStatus`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnectorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the created connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> map[string]map[string]interface{} ListConnectors(ctx).Expand(expand).Execute()

Get a list of active connectors.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    expand := []string{"Inner_example"} // []string | Retrieves additional state/configuration information for each of the connectors returned in the API call. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConnectors(context.Background()).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectors`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Retrieves additional state/configuration information for each of the connectors returned in the API call. | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseConnector

> map[string]interface{} PauseConnector(ctx, name).Execute()

Pause the connector and its tasks, which stops message processing until the connector is resumed. This call asynchronous and the tasks will not transition to PAUSED state at the same time.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the connector to pause.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PauseConnector(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PauseConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseConnector`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PauseConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the connector to pause. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartConnector

> RestartConnector(ctx, name).IncludeTasks(includeTasks).OnlyFailed(onlyFailed).Execute()

Restart the connector. You may use the following query parameters to restart any combination of the Connector and/or Task instances for the connector.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the connector.
    includeTasks := true // bool | Specifies whether to restart the connector instance and task instances (includeTasks=true) or just the connector instance (includeTasks=false). (optional)
    onlyFailed := true // bool | Specifies whether to restart just the instances with a FAILED status (onlyFailed=true) or all instances (onlyFailed=false). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RestartConnector(context.Background(), name).IncludeTasks(includeTasks).OnlyFailed(onlyFailed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestartConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTasks** | **bool** | Specifies whether to restart the connector instance and task instances (includeTasks&#x3D;true) or just the connector instance (includeTasks&#x3D;false). | 
 **onlyFailed** | **bool** | Specifies whether to restart just the instances with a FAILED status (onlyFailed&#x3D;true) or all instances (onlyFailed&#x3D;false). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeConnector

> map[string]interface{} ResumeConnector(ctx, name).Execute()

Resume a paused connector or do nothing if the connector is not paused. This call asynchronous and the tasks will not transition to RUNNING state at the same time.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the connector to resume.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ResumeConnector(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResumeConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeConnector`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResumeConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the connector to resume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

