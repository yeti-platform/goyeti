# \TagsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV2TagsTagIdDelete**](TagsAPI.md#DeleteApiV2TagsTagIdDelete) | **Delete** /api/v2/tags/{tag_id} | Delete
[**DetailsApiV2TagsTagIdGet**](TagsAPI.md#DetailsApiV2TagsTagIdGet) | **Get** /api/v2/tags/{tag_id} | Details
[**MergeApiV2TagsMergePost**](TagsAPI.md#MergeApiV2TagsMergePost) | **Post** /api/v2/tags/merge | Merge
[**NewApiV2TagsPost**](TagsAPI.md#NewApiV2TagsPost) | **Post** /api/v2/tags/ | New
[**SearchApiV2TagsSearchPost**](TagsAPI.md#SearchApiV2TagsSearchPost) | **Post** /api/v2/tags/search | Search
[**UpdateApiV2TagsTagIdPut**](TagsAPI.md#UpdateApiV2TagsTagIdPut) | **Put** /api/v2/tags/{tag_id} | Update



## DeleteApiV2TagsTagIdDelete

> interface{} DeleteApiV2TagsTagIdDelete(ctx, tagId).Execute()

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
	tagId := "tagId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.DeleteApiV2TagsTagIdDelete(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteApiV2TagsTagIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2TagsTagIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.DeleteApiV2TagsTagIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2TagsTagIdDeleteRequest struct via the builder pattern


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


## DetailsApiV2TagsTagIdGet

> Tag DetailsApiV2TagsTagIdGet(ctx, tagId).Execute()

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
	tagId := "tagId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.DetailsApiV2TagsTagIdGet(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DetailsApiV2TagsTagIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetailsApiV2TagsTagIdGet`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.DetailsApiV2TagsTagIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetailsApiV2TagsTagIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeApiV2TagsMergePost

> MergeTagResult MergeApiV2TagsMergePost(ctx).MergeTagRequest(mergeTagRequest).Execute()

Merge

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
	mergeTagRequest := *openapiclient.NewMergeTagRequest([]string{"Merge_example"}, "MergeInto_example") // MergeTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.MergeApiV2TagsMergePost(context.Background()).MergeTagRequest(mergeTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.MergeApiV2TagsMergePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeApiV2TagsMergePost`: MergeTagResult
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.MergeApiV2TagsMergePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeApiV2TagsMergePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mergeTagRequest** | [**MergeTagRequest**](MergeTagRequest.md) |  | 

### Return type

[**MergeTagResult**](MergeTagResult.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewApiV2TagsPost

> Tag NewApiV2TagsPost(ctx).NewRequest(newRequest).Execute()

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
	newRequest := *openapiclient.NewNewRequest("Name_example") // NewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.NewApiV2TagsPost(context.Background()).NewRequest(newRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.NewApiV2TagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewApiV2TagsPost`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.NewApiV2TagsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewApiV2TagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newRequest** | [**NewRequest**](NewRequest.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2TagsSearchPost

> TagSearchResponse SearchApiV2TagsSearchPost(ctx).TagSearchRequest(tagSearchRequest).Execute()

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
	tagSearchRequest := *openapiclient.NewTagSearchRequest(int32(123), int32(123)) // TagSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.SearchApiV2TagsSearchPost(context.Background()).TagSearchRequest(tagSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.SearchApiV2TagsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2TagsSearchPost`: TagSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.SearchApiV2TagsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2TagsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagSearchRequest** | [**TagSearchRequest**](TagSearchRequest.md) |  | 

### Return type

[**TagSearchResponse**](TagSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV2TagsTagIdPut

> Tag UpdateApiV2TagsTagIdPut(ctx, tagId).UpdateRequest(updateRequest).Execute()

Update



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
	tagId := "tagId_example" // string | 
	updateRequest := *openapiclient.NewUpdateRequest("Name_example") // UpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.UpdateApiV2TagsTagIdPut(context.Background(), tagId).UpdateRequest(updateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateApiV2TagsTagIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiV2TagsTagIdPut`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.UpdateApiV2TagsTagIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV2TagsTagIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

