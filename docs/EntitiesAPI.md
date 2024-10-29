# \EntitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV2EntitiesEntityIdDelete**](EntitiesAPI.md#DeleteApiV2EntitiesEntityIdDelete) | **Delete** /api/v2/entities/{entity_id} | Delete
[**DetailsApiV2EntitiesEntityIdGet**](EntitiesAPI.md#DetailsApiV2EntitiesEntityIdGet) | **Get** /api/v2/entities/{entity_id} | Details
[**NewApiV2EntitiesPost**](EntitiesAPI.md#NewApiV2EntitiesPost) | **Post** /api/v2/entities/ | New
[**PatchApiV2EntitiesEntityIdPatch**](EntitiesAPI.md#PatchApiV2EntitiesEntityIdPatch) | **Patch** /api/v2/entities/{entity_id} | Patch
[**SearchApiV2EntitiesSearchPost**](EntitiesAPI.md#SearchApiV2EntitiesSearchPost) | **Post** /api/v2/entities/search | Search
[**TagApiV2EntitiesTagPost**](EntitiesAPI.md#TagApiV2EntitiesTagPost) | **Post** /api/v2/entities/tag | Tag



## DeleteApiV2EntitiesEntityIdDelete

> interface{} DeleteApiV2EntitiesEntityIdDelete(ctx, entityId).Execute()

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
	entityId := "entityId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.DeleteApiV2EntitiesEntityIdDelete(context.Background(), entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.DeleteApiV2EntitiesEntityIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2EntitiesEntityIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.DeleteApiV2EntitiesEntityIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2EntitiesEntityIdDeleteRequest struct via the builder pattern


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


## DetailsApiV2EntitiesEntityIdGet

> ResponseDetailsApiV2EntitiesEntityIdGet DetailsApiV2EntitiesEntityIdGet(ctx, entityId).Execute()

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
	entityId := TODO // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.DetailsApiV2EntitiesEntityIdGet(context.Background(), entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.DetailsApiV2EntitiesEntityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsApiV2EntitiesEntityIdGet`: ResponseDetailsApiV2EntitiesEntityIdGet
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.DetailsApiV2EntitiesEntityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsApiV2EntitiesEntityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDetailsApiV2EntitiesEntityIdGet**](ResponseDetailsApiV2EntitiesEntityIdGet.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewApiV2EntitiesPost

> ResponseNewApiV2EntitiesPost NewApiV2EntitiesPost(ctx).NewEntityRequest(newEntityRequest).Execute()

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
	newEntityRequest := *openapiclient.NewNewEntityRequest(openapiclient.Entity_1{AttackPatternInput: openapiclient.NewAttackPatternInput("Name_example")}) // NewEntityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.NewApiV2EntitiesPost(context.Background()).NewEntityRequest(newEntityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.NewApiV2EntitiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewApiV2EntitiesPost`: ResponseNewApiV2EntitiesPost
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.NewApiV2EntitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewApiV2EntitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newEntityRequest** | [**NewEntityRequest**](NewEntityRequest.md) |  | 

### Return type

[**ResponseNewApiV2EntitiesPost**](ResponseNewApiV2EntitiesPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiV2EntitiesEntityIdPatch

> ResponsePatchApiV2EntitiesEntityIdPatch PatchApiV2EntitiesEntityIdPatch(ctx, entityId).PatchEntityRequest(patchEntityRequest).Execute()

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
	entityId := TODO // interface{} | 
	patchEntityRequest := *openapiclient.NewPatchEntityRequest(openapiclient.Entity_1{AttackPatternInput: openapiclient.NewAttackPatternInput("Name_example")}) // PatchEntityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.PatchApiV2EntitiesEntityIdPatch(context.Background(), entityId).PatchEntityRequest(patchEntityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.PatchApiV2EntitiesEntityIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiV2EntitiesEntityIdPatch`: ResponsePatchApiV2EntitiesEntityIdPatch
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.PatchApiV2EntitiesEntityIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiV2EntitiesEntityIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchEntityRequest** | [**PatchEntityRequest**](PatchEntityRequest.md) |  | 

### Return type

[**ResponsePatchApiV2EntitiesEntityIdPatch**](ResponsePatchApiV2EntitiesEntityIdPatch.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2EntitiesSearchPost

> EntitySearchResponse SearchApiV2EntitiesSearchPost(ctx).EntitySearchRequest(entitySearchRequest).Execute()

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
	entitySearchRequest := *openapiclient.NewEntitySearchRequest() // EntitySearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.SearchApiV2EntitiesSearchPost(context.Background()).EntitySearchRequest(entitySearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.SearchApiV2EntitiesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2EntitiesSearchPost`: EntitySearchResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.SearchApiV2EntitiesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2EntitiesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitySearchRequest** | [**EntitySearchRequest**](EntitySearchRequest.md) |  | 

### Return type

[**EntitySearchResponse**](EntitySearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagApiV2EntitiesTagPost

> EntityTagResponse TagApiV2EntitiesTagPost(ctx).EntityTagRequest(entityTagRequest).Execute()

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
	entityTagRequest := *openapiclient.NewEntityTagRequest([]string{"Ids_example"}) // EntityTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitiesAPI.TagApiV2EntitiesTagPost(context.Background()).EntityTagRequest(entityTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitiesAPI.TagApiV2EntitiesTagPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagApiV2EntitiesTagPost`: EntityTagResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitiesAPI.TagApiV2EntitiesTagPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagApiV2EntitiesTagPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityTagRequest** | [**EntityTagRequest**](EntityTagRequest.md) |  | 

### Return type

[**EntityTagResponse**](EntityTagResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

