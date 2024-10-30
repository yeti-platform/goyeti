# \ObservablesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContextApiV2ObservablesObservableIdContextPost**](ObservablesAPI.md#AddContextApiV2ObservablesObservableIdContextPost) | **Post** /api/v2/observables/{observable_id}/context | Add Context
[**AddTextApiV2ObservablesAddTextPost**](ObservablesAPI.md#AddTextApiV2ObservablesAddTextPost) | **Post** /api/v2/observables/add_text | Add Text
[**BulkAddApiV2ObservablesBulkPost**](ObservablesAPI.md#BulkAddApiV2ObservablesBulkPost) | **Post** /api/v2/observables/bulk | Bulk Add
[**DeleteApiV2ObservablesObservableIdDelete**](ObservablesAPI.md#DeleteApiV2ObservablesObservableIdDelete) | **Delete** /api/v2/observables/{observable_id} | Delete
[**DeleteContextApiV2ObservablesObservableIdContextDeletePost**](ObservablesAPI.md#DeleteContextApiV2ObservablesObservableIdContextDeletePost) | **Post** /api/v2/observables/{observable_id}/context/delete | Delete Context
[**DetailsApiV2ObservablesObservableIdGet**](ObservablesAPI.md#DetailsApiV2ObservablesObservableIdGet) | **Get** /api/v2/observables/{observable_id} | Details
[**NewApiV2ObservablesPost**](ObservablesAPI.md#NewApiV2ObservablesPost) | **Post** /api/v2/observables/ | New
[**NewExtendedApiV2ObservablesExtendedPost**](ObservablesAPI.md#NewExtendedApiV2ObservablesExtendedPost) | **Post** /api/v2/observables/extended | New Extended
[**ObservablesRootApiV2ObservablesGet**](ObservablesAPI.md#ObservablesRootApiV2ObservablesGet) | **Get** /api/v2/observables/ | Observables Root
[**PatchApiV2ObservablesObservableIdPatch**](ObservablesAPI.md#PatchApiV2ObservablesObservableIdPatch) | **Patch** /api/v2/observables/{observable_id} | Patch
[**SearchApiV2ObservablesSearchPost**](ObservablesAPI.md#SearchApiV2ObservablesSearchPost) | **Post** /api/v2/observables/search | Search
[**TagObservableApiV2ObservablesTagPost**](ObservablesAPI.md#TagObservableApiV2ObservablesTagPost) | **Post** /api/v2/observables/tag | Tag Observable



## AddContextApiV2ObservablesObservableIdContextPost

> ResponseAddContextApiV2ObservablesObservableIdContextPost AddContextApiV2ObservablesObservableIdContextPost(ctx, observableId).AddContextRequest(addContextRequest).Execute()

Add Context



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
	observableId := TODO // interface{} | 
	addContextRequest := *openapiclient.NewAddContextRequest("Source_example", map[string]interface{}(123)) // AddContextRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.AddContextApiV2ObservablesObservableIdContextPost(context.Background(), observableId).AddContextRequest(addContextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.AddContextApiV2ObservablesObservableIdContextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContextApiV2ObservablesObservableIdContextPost`: ResponseAddContextApiV2ObservablesObservableIdContextPost
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.AddContextApiV2ObservablesObservableIdContextPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContextApiV2ObservablesObservableIdContextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addContextRequest** | [**AddContextRequest**](AddContextRequest.md) |  | 

### Return type

[**ResponseAddContextApiV2ObservablesObservableIdContextPost**](ResponseAddContextApiV2ObservablesObservableIdContextPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTextApiV2ObservablesAddTextPost

> ResponseAddTextApiV2ObservablesAddTextPost AddTextApiV2ObservablesAddTextPost(ctx).AddTextRequest(addTextRequest).Execute()

Add Text



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
	addTextRequest := *openapiclient.NewAddTextRequest("Text_example") // AddTextRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.AddTextApiV2ObservablesAddTextPost(context.Background()).AddTextRequest(addTextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.AddTextApiV2ObservablesAddTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTextApiV2ObservablesAddTextPost`: ResponseAddTextApiV2ObservablesAddTextPost
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.AddTextApiV2ObservablesAddTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTextApiV2ObservablesAddTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTextRequest** | [**AddTextRequest**](AddTextRequest.md) |  | 

### Return type

[**ResponseAddTextApiV2ObservablesAddTextPost**](ResponseAddTextApiV2ObservablesAddTextPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAddApiV2ObservablesBulkPost

> BulkObservableAddResponse BulkAddApiV2ObservablesBulkPost(ctx).NewBulkObservableAddRequest(newBulkObservableAddRequest).Execute()

Bulk Add



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
	newBulkObservableAddRequest := *openapiclient.NewNewBulkObservableAddRequest([]openapiclient.NewObservableRequest{*openapiclient.NewNewObservableRequest("Value_example", openapiclient.ObservableType-Input("mutex"))}) // NewBulkObservableAddRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.BulkAddApiV2ObservablesBulkPost(context.Background()).NewBulkObservableAddRequest(newBulkObservableAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.BulkAddApiV2ObservablesBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAddApiV2ObservablesBulkPost`: BulkObservableAddResponse
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.BulkAddApiV2ObservablesBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddApiV2ObservablesBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBulkObservableAddRequest** | [**NewBulkObservableAddRequest**](NewBulkObservableAddRequest.md) |  | 

### Return type

[**BulkObservableAddResponse**](BulkObservableAddResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV2ObservablesObservableIdDelete

> interface{} DeleteApiV2ObservablesObservableIdDelete(ctx, observableId).Execute()

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
	observableId := "observableId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.DeleteApiV2ObservablesObservableIdDelete(context.Background(), observableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.DeleteApiV2ObservablesObservableIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2ObservablesObservableIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.DeleteApiV2ObservablesObservableIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2ObservablesObservableIdDeleteRequest struct via the builder pattern


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


## DeleteContextApiV2ObservablesObservableIdContextDeletePost

> ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost DeleteContextApiV2ObservablesObservableIdContextDeletePost(ctx, observableId).DeleteContextRequest(deleteContextRequest).Execute()

Delete Context



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
	observableId := TODO // interface{} | 
	deleteContextRequest := *openapiclient.NewDeleteContextRequest("Source_example", map[string]interface{}(123)) // DeleteContextRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.DeleteContextApiV2ObservablesObservableIdContextDeletePost(context.Background(), observableId).DeleteContextRequest(deleteContextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.DeleteContextApiV2ObservablesObservableIdContextDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContextApiV2ObservablesObservableIdContextDeletePost`: ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.DeleteContextApiV2ObservablesObservableIdContextDeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextApiV2ObservablesObservableIdContextDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteContextRequest** | [**DeleteContextRequest**](DeleteContextRequest.md) |  | 

### Return type

[**ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost**](ResponseDeleteContextApiV2ObservablesObservableIdContextDeletePost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailsApiV2ObservablesObservableIdGet

> ResponseDetailsApiV2ObservablesObservableIdGet DetailsApiV2ObservablesObservableIdGet(ctx, observableId).Execute()

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
	observableId := TODO // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.DetailsApiV2ObservablesObservableIdGet(context.Background(), observableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.DetailsApiV2ObservablesObservableIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsApiV2ObservablesObservableIdGet`: ResponseDetailsApiV2ObservablesObservableIdGet
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.DetailsApiV2ObservablesObservableIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsApiV2ObservablesObservableIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDetailsApiV2ObservablesObservableIdGet**](ResponseDetailsApiV2ObservablesObservableIdGet.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewApiV2ObservablesPost

> ResponseNewApiV2ObservablesPost NewApiV2ObservablesPost(ctx).NewObservableRequest(newObservableRequest).Execute()

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
	newObservableRequest := *openapiclient.NewNewObservableRequest("Value_example", openapiclient.ObservableType-Input("mutex")) // NewObservableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.NewApiV2ObservablesPost(context.Background()).NewObservableRequest(newObservableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.NewApiV2ObservablesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewApiV2ObservablesPost`: ResponseNewApiV2ObservablesPost
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.NewApiV2ObservablesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewApiV2ObservablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newObservableRequest** | [**NewObservableRequest**](NewObservableRequest.md) |  | 

### Return type

[**ResponseNewApiV2ObservablesPost**](ResponseNewApiV2ObservablesPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewExtendedApiV2ObservablesExtendedPost

> ResponseNewExtendedApiV2ObservablesExtendedPost NewExtendedApiV2ObservablesExtendedPost(ctx).NewExtendedObservableRequest(newExtendedObservableRequest).Execute()

New Extended



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
	newExtendedObservableRequest := *openapiclient.NewNewExtendedObservableRequest(*openapiclient.NewObservable1("Value_example", "Key_example", "TODO", openapiclient.RegistryHive("HKEY_CURRENT_CONFIG"))) // NewExtendedObservableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.NewExtendedApiV2ObservablesExtendedPost(context.Background()).NewExtendedObservableRequest(newExtendedObservableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.NewExtendedApiV2ObservablesExtendedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewExtendedApiV2ObservablesExtendedPost`: ResponseNewExtendedApiV2ObservablesExtendedPost
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.NewExtendedApiV2ObservablesExtendedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewExtendedApiV2ObservablesExtendedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newExtendedObservableRequest** | [**NewExtendedObservableRequest**](NewExtendedObservableRequest.md) |  | 

### Return type

[**ResponseNewExtendedApiV2ObservablesExtendedPost**](ResponseNewExtendedApiV2ObservablesExtendedPost.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObservablesRootApiV2ObservablesGet

> []Observable ObservablesRootApiV2ObservablesGet(ctx).Execute()

Observables Root

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
	resp, r, err := apiClient.ObservablesAPI.ObservablesRootApiV2ObservablesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.ObservablesRootApiV2ObservablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObservablesRootApiV2ObservablesGet`: []Observable
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.ObservablesRootApiV2ObservablesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiObservablesRootApiV2ObservablesGetRequest struct via the builder pattern


### Return type

[**[]Observable**](Observable.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiV2ObservablesObservableIdPatch

> ResponsePatchApiV2ObservablesObservableIdPatch PatchApiV2ObservablesObservableIdPatch(ctx, observableId).PatchObservableRequest(patchObservableRequest).Execute()

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
	observableId := TODO // interface{} | 
	patchObservableRequest := *openapiclient.NewPatchObservableRequest(*openapiclient.NewObservable1("Value_example", "Key_example", "TODO", openapiclient.RegistryHive("HKEY_CURRENT_CONFIG"))) // PatchObservableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.PatchApiV2ObservablesObservableIdPatch(context.Background(), observableId).PatchObservableRequest(patchObservableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.PatchApiV2ObservablesObservableIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiV2ObservablesObservableIdPatch`: ResponsePatchApiV2ObservablesObservableIdPatch
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.PatchApiV2ObservablesObservableIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**observableId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiV2ObservablesObservableIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchObservableRequest** | [**PatchObservableRequest**](PatchObservableRequest.md) |  | 

### Return type

[**ResponsePatchApiV2ObservablesObservableIdPatch**](ResponsePatchApiV2ObservablesObservableIdPatch.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2ObservablesSearchPost

> ObservableSearchResponse SearchApiV2ObservablesSearchPost(ctx).ObservableSearchRequest(observableSearchRequest).Execute()

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
	observableSearchRequest := *openapiclient.NewObservableSearchRequest() // ObservableSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.SearchApiV2ObservablesSearchPost(context.Background()).ObservableSearchRequest(observableSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.SearchApiV2ObservablesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2ObservablesSearchPost`: ObservableSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.SearchApiV2ObservablesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2ObservablesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **observableSearchRequest** | [**ObservableSearchRequest**](ObservableSearchRequest.md) |  | 

### Return type

[**ObservableSearchResponse**](ObservableSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagObservableApiV2ObservablesTagPost

> ObservableTagResponse TagObservableApiV2ObservablesTagPost(ctx).ObservableTagRequest(observableTagRequest).Execute()

Tag Observable



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
	observableTagRequest := *openapiclient.NewObservableTagRequest([]string{"Ids_example"}) // ObservableTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservablesAPI.TagObservableApiV2ObservablesTagPost(context.Background()).ObservableTagRequest(observableTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservablesAPI.TagObservableApiV2ObservablesTagPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagObservableApiV2ObservablesTagPost`: ObservableTagResponse
	fmt.Fprintf(os.Stdout, "Response from `ObservablesAPI.TagObservableApiV2ObservablesTagPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagObservableApiV2ObservablesTagPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **observableTagRequest** | [**ObservableTagRequest**](ObservableTagRequest.md) |  | 

### Return type

[**ObservableTagResponse**](ObservableTagResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

