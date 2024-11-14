# \SystemAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigApiV2SystemConfigGet**](SystemAPI.md#GetConfigApiV2SystemConfigGet) | **Get** /api/v2/system/config | Get Config
[**GetWorkerStatusApiV2SystemWorkersGet**](SystemAPI.md#GetWorkerStatusApiV2SystemWorkersGet) | **Get** /api/v2/system/workers | Get Worker Status
[**RestartWorkerApiV2SystemRestartworkerWorkerNamePost**](SystemAPI.md#RestartWorkerApiV2SystemRestartworkerWorkerNamePost) | **Post** /api/v2/system/restartworker/{worker_name} | Restart Worker



## GetConfigApiV2SystemConfigGet

> SystemConfigResponse GetConfigApiV2SystemConfigGet(ctx).Execute()

Get Config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yeti-platform/goyeti"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetConfigApiV2SystemConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetConfigApiV2SystemConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigApiV2SystemConfigGet`: SystemConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetConfigApiV2SystemConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigApiV2SystemConfigGetRequest struct via the builder pattern


### Return type

[**SystemConfigResponse**](SystemConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkerStatusApiV2SystemWorkersGet

> WorkerStatusResponse GetWorkerStatusApiV2SystemWorkersGet(ctx).Execute()

Get Worker Status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yeti-platform/goyeti"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetWorkerStatusApiV2SystemWorkersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetWorkerStatusApiV2SystemWorkersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkerStatusApiV2SystemWorkersGet`: WorkerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetWorkerStatusApiV2SystemWorkersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkerStatusApiV2SystemWorkersGetRequest struct via the builder pattern


### Return type

[**WorkerStatusResponse**](WorkerStatusResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartWorkerApiV2SystemRestartworkerWorkerNamePost

> WorkerRestartResponse RestartWorkerApiV2SystemRestartworkerWorkerNamePost(ctx, workerName).Execute()

Restart Worker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yeti-platform/goyeti"
)

func main() {
	workerName := "workerName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.RestartWorkerApiV2SystemRestartworkerWorkerNamePost(context.Background(), workerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RestartWorkerApiV2SystemRestartworkerWorkerNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartWorkerApiV2SystemRestartworkerWorkerNamePost`: WorkerRestartResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.RestartWorkerApiV2SystemRestartworkerWorkerNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartWorkerApiV2SystemRestartworkerWorkerNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkerRestartResponse**](WorkerRestartResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

