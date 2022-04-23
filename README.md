# Go API client for openapi

Kafka Connect REST API https://docs.confluent.io/platform/current/connect/references/restapi.html

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.9
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/ahawker/kafka-connect-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateConnector**](docs/DefaultApi.md#createconnector) | **Post** /connectors | Create a new connector, returning the current connector info if successful. Return 409 (Conflict) if rebalance is in process, or if the connector already exists.
*DefaultApi* | [**DeleteConnector**](docs/DefaultApi.md#deleteconnector) | **Delete** /connectors/{name} | Delete a connector, halting all tasks and deleting its configuration.
*DefaultApi* | [**GetClusterInfo**](docs/DefaultApi.md#getclusterinfo) | **Get** / | Top-level (root) request that gets the version of the Connect worker that serves the REST request, the git commit ID of the source code, and the Kafka cluster ID that the worker is connected to.
*DefaultApi* | [**GetConnector**](docs/DefaultApi.md#getconnector) | **Get** /connectors/{name} | Get information about the connector.
*DefaultApi* | [**GetConnectorStatus**](docs/DefaultApi.md#getconnectorstatus) | **Get** /connectors/{name}/status | Gets the current status of the connector, including: * whether it is running or restarting, or if it has failed or paused * which worker it is assigned to * error information if it has failed * the state of all its tasks 
*DefaultApi* | [**ListConnectors**](docs/DefaultApi.md#listconnectors) | **Get** /connectors | Get a list of active connectors.
*DefaultApi* | [**PauseConnector**](docs/DefaultApi.md#pauseconnector) | **Put** /connectors/{name}/pause | Pause the connector and its tasks, which stops message processing until the connector is resumed. This call asynchronous and the tasks will not transition to PAUSED state at the same time.
*DefaultApi* | [**RestartConnector**](docs/DefaultApi.md#restartconnector) | **Post** /connectors/{name}/restart | Restart the connector. You may use the following query parameters to restart any combination of the Connector and/or Task instances for the connector.
*DefaultApi* | [**ResumeConnector**](docs/DefaultApi.md#resumeconnector) | **Put** /connectors/{name}/resume | Resume a paused connector or do nothing if the connector is not paused. This call asynchronous and the tasks will not transition to RUNNING state at the same time.


## Documentation For Models

 - [CreateConnectorRequest](docs/CreateConnectorRequest.md)
 - [CreateConnectorResponse](docs/CreateConnectorResponse.md)
 - [CreateConnectorResponseTasks](docs/CreateConnectorResponseTasks.md)
 - [GetClusterInfoResponse](docs/GetClusterInfoResponse.md)
 - [GetConnectorResponse](docs/GetConnectorResponse.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse200Connector](docs/InlineResponse200Connector.md)
 - [InlineResponse200Tasks](docs/InlineResponse200Tasks.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



