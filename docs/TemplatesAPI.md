# \TemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RenderApiV2TemplatesRenderPost**](TemplatesAPI.md#RenderApiV2TemplatesRenderPost) | **Post** /api/v2/templates/render | Render
[**SearchApiV2TemplatesSearchPost**](TemplatesAPI.md#SearchApiV2TemplatesSearchPost) | **Post** /api/v2/templates/search | Search



## RenderApiV2TemplatesRenderPost

> interface{} RenderApiV2TemplatesRenderPost(ctx).RenderTemplateRequest(renderTemplateRequest).Execute()

Render



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
	renderTemplateRequest := *openapiclient.NewRenderTemplateRequest("TemplateName_example") // RenderTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.RenderApiV2TemplatesRenderPost(context.Background()).RenderTemplateRequest(renderTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.RenderApiV2TemplatesRenderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderApiV2TemplatesRenderPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.RenderApiV2TemplatesRenderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderApiV2TemplatesRenderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renderTemplateRequest** | [**RenderTemplateRequest**](RenderTemplateRequest.md) |  | 

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


## SearchApiV2TemplatesSearchPost

> TemplateSearchResponse SearchApiV2TemplatesSearchPost(ctx).TemplateSearchRequest(templateSearchRequest).Execute()

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
	templateSearchRequest := *openapiclient.NewTemplateSearchRequest() // TemplateSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.SearchApiV2TemplatesSearchPost(context.Background()).TemplateSearchRequest(templateSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.SearchApiV2TemplatesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2TemplatesSearchPost`: TemplateSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.SearchApiV2TemplatesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2TemplatesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateSearchRequest** | [**TemplateSearchRequest**](TemplateSearchRequest.md) |  | 

### Return type

[**TemplateSearchResponse**](TemplateSearchResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

