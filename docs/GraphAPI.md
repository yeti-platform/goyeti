# \GraphAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApiV2GraphAddPost**](GraphAPI.md#AddApiV2GraphAddPost) | **Post** /api/v2/graph/add | Add
[**DeleteApiV2GraphRelationshipIdDelete**](GraphAPI.md#DeleteApiV2GraphRelationshipIdDelete) | **Delete** /api/v2/graph/{relationship_id} | Delete
[**EditApiV2GraphRelationshipIdPatch**](GraphAPI.md#EditApiV2GraphRelationshipIdPatch) | **Patch** /api/v2/graph/{relationship_id} | Edit
[**MatchApiV2GraphMatchPost**](GraphAPI.md#MatchApiV2GraphMatchPost) | **Post** /api/v2/graph/match | Match
[**SearchApiV2GraphSearchPost**](GraphAPI.md#SearchApiV2GraphSearchPost) | **Post** /api/v2/graph/search | Search
[**SwapApiV2GraphRelationshipIdSwapPost**](GraphAPI.md#SwapApiV2GraphRelationshipIdSwapPost) | **Post** /api/v2/graph/{relationship_id}/swap | Swap



## AddApiV2GraphAddPost

> Relationship AddApiV2GraphAddPost(ctx).GraphAddRequest(graphAddRequest).Execute()

Add



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
	graphAddRequest := *openapiclient.NewGraphAddRequest("Source_example", "Target_example", "LinkType_example", "Description_example") // GraphAddRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.AddApiV2GraphAddPost(context.Background()).GraphAddRequest(graphAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.AddApiV2GraphAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddApiV2GraphAddPost`: Relationship
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.AddApiV2GraphAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApiV2GraphAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphAddRequest** | [**GraphAddRequest**](GraphAddRequest.md) |  | 

### Return type

[**Relationship**](Relationship.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV2GraphRelationshipIdDelete

> interface{} DeleteApiV2GraphRelationshipIdDelete(ctx, relationshipId).Execute()

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
	relationshipId := "relationshipId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.DeleteApiV2GraphRelationshipIdDelete(context.Background(), relationshipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.DeleteApiV2GraphRelationshipIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2GraphRelationshipIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.DeleteApiV2GraphRelationshipIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2GraphRelationshipIdDeleteRequest struct via the builder pattern


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


## EditApiV2GraphRelationshipIdPatch

> Relationship EditApiV2GraphRelationshipIdPatch(ctx, relationshipId).GraphPatchRequest(graphPatchRequest).Execute()

Edit



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
	relationshipId := "relationshipId_example" // string | 
	graphPatchRequest := *openapiclient.NewGraphPatchRequest("LinkType_example", "Description_example") // GraphPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.EditApiV2GraphRelationshipIdPatch(context.Background(), relationshipId).GraphPatchRequest(graphPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.EditApiV2GraphRelationshipIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditApiV2GraphRelationshipIdPatch`: Relationship
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.EditApiV2GraphRelationshipIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApiV2GraphRelationshipIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **graphPatchRequest** | [**GraphPatchRequest**](GraphPatchRequest.md) |  | 

### Return type

[**Relationship**](Relationship.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MatchApiV2GraphMatchPost

> AnalysisResponse MatchApiV2GraphMatchPost(ctx).AnalysisRequest(analysisRequest).Execute()

Match



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
	analysisRequest := *openapiclient.NewAnalysisRequest([]string{"Observables_example"}) // AnalysisRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.MatchApiV2GraphMatchPost(context.Background()).AnalysisRequest(analysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.MatchApiV2GraphMatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MatchApiV2GraphMatchPost`: AnalysisResponse
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.MatchApiV2GraphMatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMatchApiV2GraphMatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analysisRequest** | [**AnalysisRequest**](AnalysisRequest.md) |  | 

### Return type

[**AnalysisResponse**](AnalysisResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2GraphSearchPost

> GraphSearchResponse SearchApiV2GraphSearchPost(ctx).GraphSearchRequest(graphSearchRequest).Execute()

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
	graphSearchRequest := *openapiclient.NewGraphSearchRequest("Source_example", "Graph_example", openapiclient.GraphDirection("outbound"), false) // GraphSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.SearchApiV2GraphSearchPost(context.Background()).GraphSearchRequest(graphSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.SearchApiV2GraphSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2GraphSearchPost`: GraphSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.SearchApiV2GraphSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2GraphSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphSearchRequest** | [**GraphSearchRequest**](GraphSearchRequest.md) |  | 

### Return type

[**GraphSearchResponse**](GraphSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapApiV2GraphRelationshipIdSwapPost

> Relationship SwapApiV2GraphRelationshipIdSwapPost(ctx, relationshipId).Execute()

Swap



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
	relationshipId := "relationshipId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GraphAPI.SwapApiV2GraphRelationshipIdSwapPost(context.Background(), relationshipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphAPI.SwapApiV2GraphRelationshipIdSwapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwapApiV2GraphRelationshipIdSwapPost`: Relationship
	fmt.Fprintf(os.Stdout, "Response from `GraphAPI.SwapApiV2GraphRelationshipIdSwapPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapApiV2GraphRelationshipIdSwapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Relationship**](Relationship.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

