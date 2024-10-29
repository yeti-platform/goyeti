# \TasksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExportApiV2TasksExportExportNameDelete**](TasksAPI.md#DeleteExportApiV2TasksExportExportNameDelete) | **Delete** /api/v2/tasks/export/{export_name} | Delete Export
[**ExportContentApiV2TasksExportExportIdContentGet**](TasksAPI.md#ExportContentApiV2TasksExportExportIdContentGet) | **Get** /api/v2/tasks/export/{export_id}/content | Export Content
[**NewExportApiV2TasksExportNewPost**](TasksAPI.md#NewExportApiV2TasksExportNewPost) | **Post** /api/v2/tasks/export/new | New Export
[**PatchExportApiV2TasksExportExportNamePatch**](TasksAPI.md#PatchExportApiV2TasksExportExportNamePatch) | **Patch** /api/v2/tasks/export/{export_name} | Patch Export
[**RunApiV2TasksTaskNameRunPost**](TasksAPI.md#RunApiV2TasksTaskNameRunPost) | **Post** /api/v2/tasks/{task_name}/run | Run
[**SearchApiV2TasksSearchPost**](TasksAPI.md#SearchApiV2TasksSearchPost) | **Post** /api/v2/tasks/search | Search
[**ToggleApiV2TasksTaskNameTogglePost**](TasksAPI.md#ToggleApiV2TasksTaskNameTogglePost) | **Post** /api/v2/tasks/{task_name}/toggle | Toggle



## DeleteExportApiV2TasksExportExportNameDelete

> interface{} DeleteExportApiV2TasksExportExportNameDelete(ctx, exportName).Execute()

Delete Export



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
	exportName := "exportName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.DeleteExportApiV2TasksExportExportNameDelete(context.Background(), exportName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.DeleteExportApiV2TasksExportExportNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExportApiV2TasksExportExportNameDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.DeleteExportApiV2TasksExportExportNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExportApiV2TasksExportExportNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportContentApiV2TasksExportExportIdContentGet

> interface{} ExportContentApiV2TasksExportExportIdContentGet(ctx, exportId).Execute()

Export Content



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
	exportId := "exportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ExportContentApiV2TasksExportExportIdContentGet(context.Background(), exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ExportContentApiV2TasksExportExportIdContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportContentApiV2TasksExportExportIdContentGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ExportContentApiV2TasksExportExportIdContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportContentApiV2TasksExportExportIdContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewExportApiV2TasksExportNewPost

> ExportTaskOutput NewExportApiV2TasksExportNewPost(ctx).NewExportRequest(newExportRequest).Execute()

New Export



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
	newExportRequest := *openapiclient.NewNewExportRequest(*openapiclient.NewExportTaskInput("Name_example", "TemplateName_example")) // NewExportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.NewExportApiV2TasksExportNewPost(context.Background()).NewExportRequest(newExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.NewExportApiV2TasksExportNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewExportApiV2TasksExportNewPost`: ExportTaskOutput
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.NewExportApiV2TasksExportNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewExportApiV2TasksExportNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newExportRequest** | [**NewExportRequest**](NewExportRequest.md) |  | 

### Return type

[**ExportTaskOutput**](ExportTaskOutput.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExportApiV2TasksExportExportNamePatch

> ExportTaskOutput PatchExportApiV2TasksExportExportNamePatch(ctx, exportName).PatchExportRequest(patchExportRequest).Execute()

Patch Export



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
	exportName := "exportName_example" // string | 
	patchExportRequest := *openapiclient.NewPatchExportRequest(*openapiclient.NewExportTaskInput("Name_example", "TemplateName_example")) // PatchExportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.PatchExportApiV2TasksExportExportNamePatch(context.Background(), exportName).PatchExportRequest(patchExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.PatchExportApiV2TasksExportExportNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExportApiV2TasksExportExportNamePatch`: ExportTaskOutput
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.PatchExportApiV2TasksExportExportNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExportApiV2TasksExportExportNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchExportRequest** | [**PatchExportRequest**](PatchExportRequest.md) |  | 

### Return type

[**ExportTaskOutput**](ExportTaskOutput.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunApiV2TasksTaskNameRunPost

> map[string]interface{} RunApiV2TasksTaskNameRunPost(ctx, taskName).TaskParams(taskParams).Execute()

Run



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
	taskName := TODO // interface{} | 
	taskParams := *openapiclient.NewTaskParams() // TaskParams |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.RunApiV2TasksTaskNameRunPost(context.Background(), taskName).TaskParams(taskParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.RunApiV2TasksTaskNameRunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunApiV2TasksTaskNameRunPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.RunApiV2TasksTaskNameRunPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunApiV2TasksTaskNameRunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskParams** | [**TaskParams**](TaskParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2TasksSearchPost

> TaskSearchResponse SearchApiV2TasksSearchPost(ctx).TaskSearchRequest(taskSearchRequest).Execute()

Search



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
	taskSearchRequest := *openapiclient.NewTaskSearchRequest() // TaskSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.SearchApiV2TasksSearchPost(context.Background()).TaskSearchRequest(taskSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.SearchApiV2TasksSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2TasksSearchPost`: TaskSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.SearchApiV2TasksSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2TasksSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskSearchRequest** | [**TaskSearchRequest**](TaskSearchRequest.md) |  | 

### Return type

[**TaskSearchResponse**](TaskSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleApiV2TasksTaskNameTogglePost

> ResponseToggleApiV2TasksTaskNameTogglePost ToggleApiV2TasksTaskNameTogglePost(ctx, taskName).Execute()

Toggle



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
	taskName := TODO // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ToggleApiV2TasksTaskNameTogglePost(context.Background(), taskName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ToggleApiV2TasksTaskNameTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleApiV2TasksTaskNameTogglePost`: ResponseToggleApiV2TasksTaskNameTogglePost
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ToggleApiV2TasksTaskNameTogglePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleApiV2TasksTaskNameTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseToggleApiV2TasksTaskNameTogglePost**](ResponseToggleApiV2TasksTaskNameTogglePost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

