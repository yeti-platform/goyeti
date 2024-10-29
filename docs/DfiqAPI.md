# \DfiqAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigApiV2DfiqConfigGet**](DfiqAPI.md#ConfigApiV2DfiqConfigGet) | **Get** /api/v2/dfiq/config | Config
[**DeleteApiV2DfiqDfiqIdDelete**](DfiqAPI.md#DeleteApiV2DfiqDfiqIdDelete) | **Delete** /api/v2/dfiq/{dfiq_id} | Delete
[**DetailsApiV2DfiqDfiqIdGet**](DfiqAPI.md#DetailsApiV2DfiqDfiqIdGet) | **Get** /api/v2/dfiq/{dfiq_id} | Details
[**FromArchiveApiV2DfiqFromArchivePost**](DfiqAPI.md#FromArchiveApiV2DfiqFromArchivePost) | **Post** /api/v2/dfiq/from_archive | From Archive
[**NewFromYamlApiV2DfiqFromYamlPost**](DfiqAPI.md#NewFromYamlApiV2DfiqFromYamlPost) | **Post** /api/v2/dfiq/from_yaml | New From Yaml
[**PatchApiV2DfiqDfiqIdPatch**](DfiqAPI.md#PatchApiV2DfiqDfiqIdPatch) | **Patch** /api/v2/dfiq/{dfiq_id} | Patch
[**SearchApiV2DfiqSearchPost**](DfiqAPI.md#SearchApiV2DfiqSearchPost) | **Post** /api/v2/dfiq/search | Search
[**ToArchiveApiV2DfiqToArchivePost**](DfiqAPI.md#ToArchiveApiV2DfiqToArchivePost) | **Post** /api/v2/dfiq/to_archive | To Archive
[**ValidateDfiqYamlApiV2DfiqValidatePost**](DfiqAPI.md#ValidateDfiqYamlApiV2DfiqValidatePost) | **Post** /api/v2/dfiq/validate | Validate Dfiq Yaml



## ConfigApiV2DfiqConfigGet

> DFIQConfigResponse ConfigApiV2DfiqConfigGet(ctx).Execute()

Config

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
	resp, r, err := apiClient.DfiqAPI.ConfigApiV2DfiqConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.ConfigApiV2DfiqConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigApiV2DfiqConfigGet`: DFIQConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.ConfigApiV2DfiqConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConfigApiV2DfiqConfigGetRequest struct via the builder pattern


### Return type

[**DFIQConfigResponse**](DFIQConfigResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV2DfiqDfiqIdDelete

> interface{} DeleteApiV2DfiqDfiqIdDelete(ctx, dfiqId).Execute()

Delete



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
	dfiqId := "dfiqId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.DeleteApiV2DfiqDfiqIdDelete(context.Background(), dfiqId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.DeleteApiV2DfiqDfiqIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2DfiqDfiqIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.DeleteApiV2DfiqDfiqIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfiqId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2DfiqDfiqIdDeleteRequest struct via the builder pattern


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


## DetailsApiV2DfiqDfiqIdGet

> ResponseDetailsApiV2DfiqDfiqIdGet DetailsApiV2DfiqDfiqIdGet(ctx, dfiqId).Execute()

Details



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
	dfiqId := TODO // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.DetailsApiV2DfiqDfiqIdGet(context.Background(), dfiqId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.DetailsApiV2DfiqDfiqIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsApiV2DfiqDfiqIdGet`: ResponseDetailsApiV2DfiqDfiqIdGet
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.DetailsApiV2DfiqDfiqIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfiqId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsApiV2DfiqDfiqIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDetailsApiV2DfiqDfiqIdGet**](ResponseDetailsApiV2DfiqDfiqIdGet.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FromArchiveApiV2DfiqFromArchivePost

> map[string]int32 FromArchiveApiV2DfiqFromArchivePost(ctx).Archive(archive).Execute()

From Archive



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
	archive := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.FromArchiveApiV2DfiqFromArchivePost(context.Background()).Archive(archive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.FromArchiveApiV2DfiqFromArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FromArchiveApiV2DfiqFromArchivePost`: map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.FromArchiveApiV2DfiqFromArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFromArchiveApiV2DfiqFromArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archive** | ***os.File** |  | 

### Return type

**map[string]int32**

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewFromYamlApiV2DfiqFromYamlPost

> ResponseNewFromYamlApiV2DfiqFromYamlPost NewFromYamlApiV2DfiqFromYamlPost(ctx).NewDFIQRequest(newDFIQRequest).Execute()

New From Yaml



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
	newDFIQRequest := *openapiclient.NewNewDFIQRequest("DfiqYaml_example", openapiclient.DFIQType("scenario")) // NewDFIQRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.NewFromYamlApiV2DfiqFromYamlPost(context.Background()).NewDFIQRequest(newDFIQRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.NewFromYamlApiV2DfiqFromYamlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewFromYamlApiV2DfiqFromYamlPost`: ResponseNewFromYamlApiV2DfiqFromYamlPost
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.NewFromYamlApiV2DfiqFromYamlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewFromYamlApiV2DfiqFromYamlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDFIQRequest** | [**NewDFIQRequest**](NewDFIQRequest.md) |  | 

### Return type

[**ResponseNewFromYamlApiV2DfiqFromYamlPost**](ResponseNewFromYamlApiV2DfiqFromYamlPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiV2DfiqDfiqIdPatch

> ResponsePatchApiV2DfiqDfiqIdPatch PatchApiV2DfiqDfiqIdPatch(ctx, dfiqId).PatchDFIQRequest(patchDFIQRequest).Execute()

Patch



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
	dfiqId := TODO // interface{} | 
	patchDFIQRequest := *openapiclient.NewPatchDFIQRequest("DfiqYaml_example", openapiclient.DFIQType("scenario")) // PatchDFIQRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.PatchApiV2DfiqDfiqIdPatch(context.Background(), dfiqId).PatchDFIQRequest(patchDFIQRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.PatchApiV2DfiqDfiqIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiV2DfiqDfiqIdPatch`: ResponsePatchApiV2DfiqDfiqIdPatch
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.PatchApiV2DfiqDfiqIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfiqId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiV2DfiqDfiqIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDFIQRequest** | [**PatchDFIQRequest**](PatchDFIQRequest.md) |  | 

### Return type

[**ResponsePatchApiV2DfiqDfiqIdPatch**](ResponsePatchApiV2DfiqDfiqIdPatch.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2DfiqSearchPost

> DFIQSearchResponse SearchApiV2DfiqSearchPost(ctx).DFIQSearchRequest(dFIQSearchRequest).Execute()

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
	dFIQSearchRequest := *openapiclient.NewDFIQSearchRequest() // DFIQSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.SearchApiV2DfiqSearchPost(context.Background()).DFIQSearchRequest(dFIQSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.SearchApiV2DfiqSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2DfiqSearchPost`: DFIQSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.SearchApiV2DfiqSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2DfiqSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dFIQSearchRequest** | [**DFIQSearchRequest**](DFIQSearchRequest.md) |  | 

### Return type

[**DFIQSearchResponse**](DFIQSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToArchiveApiV2DfiqToArchivePost

> interface{} ToArchiveApiV2DfiqToArchivePost(ctx).DFIQSearchRequest(dFIQSearchRequest).Execute()

To Archive



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
	dFIQSearchRequest := *openapiclient.NewDFIQSearchRequest() // DFIQSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.ToArchiveApiV2DfiqToArchivePost(context.Background()).DFIQSearchRequest(dFIQSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.ToArchiveApiV2DfiqToArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToArchiveApiV2DfiqToArchivePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.ToArchiveApiV2DfiqToArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToArchiveApiV2DfiqToArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dFIQSearchRequest** | [**DFIQSearchRequest**](DFIQSearchRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateDfiqYamlApiV2DfiqValidatePost

> DFIQValidateResponse ValidateDfiqYamlApiV2DfiqValidatePost(ctx).DFIQValidateRequest(dFIQValidateRequest).Execute()

Validate Dfiq Yaml



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
	dFIQValidateRequest := *openapiclient.NewDFIQValidateRequest("DfiqYaml_example", openapiclient.DFIQType("scenario")) // DFIQValidateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfiqAPI.ValidateDfiqYamlApiV2DfiqValidatePost(context.Background()).DFIQValidateRequest(dFIQValidateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfiqAPI.ValidateDfiqYamlApiV2DfiqValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateDfiqYamlApiV2DfiqValidatePost`: DFIQValidateResponse
	fmt.Fprintf(os.Stdout, "Response from `DfiqAPI.ValidateDfiqYamlApiV2DfiqValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDfiqYamlApiV2DfiqValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dFIQValidateRequest** | [**DFIQValidateRequest**](DFIQValidateRequest.md) |  | 

### Return type

[**DFIQValidateResponse**](DFIQValidateResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

