# \IndicatorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV2IndicatorsIndicatorIdDelete**](IndicatorsAPI.md#DeleteApiV2IndicatorsIndicatorIdDelete) | **Delete** /api/v2/indicators/{indicator_id} | Delete
[**DetailsApiV2IndicatorsIndicatorIdGet**](IndicatorsAPI.md#DetailsApiV2IndicatorsIndicatorIdGet) | **Get** /api/v2/indicators/{indicator_id} | Details
[**NewApiV2IndicatorsPost**](IndicatorsAPI.md#NewApiV2IndicatorsPost) | **Post** /api/v2/indicators/ | New
[**PatchApiV2IndicatorsIndicatorIdPatch**](IndicatorsAPI.md#PatchApiV2IndicatorsIndicatorIdPatch) | **Patch** /api/v2/indicators/{indicator_id} | Patch
[**SearchApiV2IndicatorsSearchPost**](IndicatorsAPI.md#SearchApiV2IndicatorsSearchPost) | **Post** /api/v2/indicators/search | Search
[**TagApiV2IndicatorsTagPost**](IndicatorsAPI.md#TagApiV2IndicatorsTagPost) | **Post** /api/v2/indicators/tag | Tag



## DeleteApiV2IndicatorsIndicatorIdDelete

> interface{} DeleteApiV2IndicatorsIndicatorIdDelete(ctx, indicatorId).Execute()

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
	indicatorId := "indicatorId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicatorsAPI.DeleteApiV2IndicatorsIndicatorIdDelete(context.Background(), indicatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicatorsAPI.DeleteApiV2IndicatorsIndicatorIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2IndicatorsIndicatorIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IndicatorsAPI.DeleteApiV2IndicatorsIndicatorIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indicatorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2IndicatorsIndicatorIdDeleteRequest struct via the builder pattern


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


## DetailsApiV2IndicatorsIndicatorIdGet

> ResponseDetailsApiV2IndicatorsIndicatorIdGet DetailsApiV2IndicatorsIndicatorIdGet(ctx, indicatorId).Execute()

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
	indicatorId := TODO // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicatorsAPI.DetailsApiV2IndicatorsIndicatorIdGet(context.Background(), indicatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicatorsAPI.DetailsApiV2IndicatorsIndicatorIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsApiV2IndicatorsIndicatorIdGet`: ResponseDetailsApiV2IndicatorsIndicatorIdGet
	fmt.Fprintf(os.Stdout, "Response from `IndicatorsAPI.DetailsApiV2IndicatorsIndicatorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indicatorId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsApiV2IndicatorsIndicatorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDetailsApiV2IndicatorsIndicatorIdGet**](ResponseDetailsApiV2IndicatorsIndicatorIdGet.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewApiV2IndicatorsPost

> ResponseNewApiV2IndicatorsPost NewApiV2IndicatorsPost(ctx).NewIndicatorRequest(newIndicatorRequest).Execute()

New



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
	newIndicatorRequest := *openapiclient.NewNewIndicatorRequest(openapiclient.Indicator_1{ForensicArtifactInput: openapiclient.NewForensicArtifactInput("Name_example", "Pattern_example", openapiclient.DiamondModel("adversary"))}) // NewIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicatorsAPI.NewApiV2IndicatorsPost(context.Background()).NewIndicatorRequest(newIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicatorsAPI.NewApiV2IndicatorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewApiV2IndicatorsPost`: ResponseNewApiV2IndicatorsPost
	fmt.Fprintf(os.Stdout, "Response from `IndicatorsAPI.NewApiV2IndicatorsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewApiV2IndicatorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newIndicatorRequest** | [**NewIndicatorRequest**](NewIndicatorRequest.md) |  | 

### Return type

[**ResponseNewApiV2IndicatorsPost**](ResponseNewApiV2IndicatorsPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiV2IndicatorsIndicatorIdPatch

> ResponsePatchApiV2IndicatorsIndicatorIdPatch PatchApiV2IndicatorsIndicatorIdPatch(ctx, indicatorId).PatchIndicatorRequest(patchIndicatorRequest).Execute()

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
	indicatorId := TODO // interface{} | 
	patchIndicatorRequest := *openapiclient.NewPatchIndicatorRequest(openapiclient.Indicator_1{ForensicArtifactInput: openapiclient.NewForensicArtifactInput("Name_example", "Pattern_example", openapiclient.DiamondModel("adversary"))}) // PatchIndicatorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicatorsAPI.PatchApiV2IndicatorsIndicatorIdPatch(context.Background(), indicatorId).PatchIndicatorRequest(patchIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicatorsAPI.PatchApiV2IndicatorsIndicatorIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiV2IndicatorsIndicatorIdPatch`: ResponsePatchApiV2IndicatorsIndicatorIdPatch
	fmt.Fprintf(os.Stdout, "Response from `IndicatorsAPI.PatchApiV2IndicatorsIndicatorIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indicatorId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiV2IndicatorsIndicatorIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchIndicatorRequest** | [**PatchIndicatorRequest**](PatchIndicatorRequest.md) |  | 

### Return type

[**ResponsePatchApiV2IndicatorsIndicatorIdPatch**](ResponsePatchApiV2IndicatorsIndicatorIdPatch.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2IndicatorsSearchPost

> IndicatorSearchResponse SearchApiV2IndicatorsSearchPost(ctx).IndicatorSearchRequest(indicatorSearchRequest).Execute()

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
	indicatorSearchRequest := *openapiclient.NewIndicatorSearchRequest() // IndicatorSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicatorsAPI.SearchApiV2IndicatorsSearchPost(context.Background()).IndicatorSearchRequest(indicatorSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicatorsAPI.SearchApiV2IndicatorsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2IndicatorsSearchPost`: IndicatorSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `IndicatorsAPI.SearchApiV2IndicatorsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2IndicatorsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indicatorSearchRequest** | [**IndicatorSearchRequest**](IndicatorSearchRequest.md) |  | 

### Return type

[**IndicatorSearchResponse**](IndicatorSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagApiV2IndicatorsTagPost

> IndicatorTagResponse TagApiV2IndicatorsTagPost(ctx).IndicatorTagRequest(indicatorTagRequest).Execute()

Tag



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
	indicatorTagRequest := *openapiclient.NewIndicatorTagRequest([]string{"Ids_example"}) // IndicatorTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndicatorsAPI.TagApiV2IndicatorsTagPost(context.Background()).IndicatorTagRequest(indicatorTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndicatorsAPI.TagApiV2IndicatorsTagPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagApiV2IndicatorsTagPost`: IndicatorTagResponse
	fmt.Fprintf(os.Stdout, "Response from `IndicatorsAPI.TagApiV2IndicatorsTagPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagApiV2IndicatorsTagPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indicatorTagRequest** | [**IndicatorTagRequest**](IndicatorTagRequest.md) |  | 

### Return type

[**IndicatorTagResponse**](IndicatorTagResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

