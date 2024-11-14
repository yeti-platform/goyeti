# \ImportDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportMispJsonApiV2ImportDataImportMispJsonPost**](ImportDataAPI.md#ImportMispJsonApiV2ImportDataImportMispJsonPost) | **Post** /api/v2/import_data/import_misp_json | Import Misp Json



## ImportMispJsonApiV2ImportDataImportMispJsonPost

> interface{} ImportMispJsonApiV2ImportDataImportMispJsonPost(ctx).MispFileJson(mispFileJson).Execute()

Import Misp Json

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
	mispFileJson := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportDataAPI.ImportMispJsonApiV2ImportDataImportMispJsonPost(context.Background()).MispFileJson(mispFileJson).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportDataAPI.ImportMispJsonApiV2ImportDataImportMispJsonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportMispJsonApiV2ImportDataImportMispJsonPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImportDataAPI.ImportMispJsonApiV2ImportDataImportMispJsonPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportMispJsonApiV2ImportDataImportMispJsonPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mispFileJson** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

