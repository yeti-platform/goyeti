# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV2UsersPost**](UsersAPI.md#CreateApiV2UsersPost) | **Post** /api/v2/users/ | Create
[**DeleteApiV2UsersUserIdDelete**](UsersAPI.md#DeleteApiV2UsersUserIdDelete) | **Delete** /api/v2/users/{user_id} | Delete
[**GetApiV2UsersUserIdGet**](UsersAPI.md#GetApiV2UsersUserIdGet) | **Get** /api/v2/users/{user_id} | Get
[**ResetApiKeyApiV2UsersResetApiKeyPost**](UsersAPI.md#ResetApiKeyApiV2UsersResetApiKeyPost) | **Post** /api/v2/users/reset-api-key | Reset Api Key
[**ResetPasswordApiV2UsersResetPasswordPost**](UsersAPI.md#ResetPasswordApiV2UsersResetPasswordPost) | **Post** /api/v2/users/reset-password | Reset Password
[**SearchApiV2UsersSearchPost**](UsersAPI.md#SearchApiV2UsersSearchPost) | **Post** /api/v2/users/search | Search
[**ToggleApiV2UsersTogglePost**](UsersAPI.md#ToggleApiV2UsersTogglePost) | **Post** /api/v2/users/toggle | Toggle



## CreateApiV2UsersPost

> User CreateApiV2UsersPost(ctx).NewUserRequest(newUserRequest).Execute()

Create



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
	newUserRequest := *openapiclient.NewNewUserRequest("Username_example", "Password_example", false) // NewUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateApiV2UsersPost(context.Background()).NewUserRequest(newUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateApiV2UsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiV2UsersPost`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateApiV2UsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV2UsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newUserRequest** | [**NewUserRequest**](NewUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV2UsersUserIdDelete

> interface{} DeleteApiV2UsersUserIdDelete(ctx, userId).Execute()

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteApiV2UsersUserIdDelete(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteApiV2UsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV2UsersUserIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteApiV2UsersUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV2UsersUserIdDeleteRequest struct via the builder pattern


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


## GetApiV2UsersUserIdGet

> User GetApiV2UsersUserIdGet(ctx, userId).Execute()

Get



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetApiV2UsersUserIdGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetApiV2UsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV2UsersUserIdGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetApiV2UsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV2UsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetApiKeyApiV2UsersResetApiKeyPost

> User ResetApiKeyApiV2UsersResetApiKeyPost(ctx).ResetApiKeyRequest(resetApiKeyRequest).Execute()

Reset Api Key



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
	resetApiKeyRequest := *openapiclient.NewResetApiKeyRequest("UserId_example") // ResetApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ResetApiKeyApiV2UsersResetApiKeyPost(context.Background()).ResetApiKeyRequest(resetApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResetApiKeyApiV2UsersResetApiKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetApiKeyApiV2UsersResetApiKeyPost`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ResetApiKeyApiV2UsersResetApiKeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetApiKeyApiV2UsersResetApiKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetApiKeyRequest** | [**ResetApiKeyRequest**](ResetApiKeyRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordApiV2UsersResetPasswordPost

> User ResetPasswordApiV2UsersResetPasswordPost(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Reset Password



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
	resetPasswordRequest := *openapiclient.NewResetPasswordRequest("UserId_example", "NewPassword_example") // ResetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ResetPasswordApiV2UsersResetPasswordPost(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResetPasswordApiV2UsersResetPasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetPasswordApiV2UsersResetPasswordPost`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ResetPasswordApiV2UsersResetPasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordApiV2UsersResetPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApiV2UsersSearchPost

> SearchUserResponse SearchApiV2UsersSearchPost(ctx).SearchUserRequest(searchUserRequest).Execute()

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
	searchUserRequest := *openapiclient.NewSearchUserRequest("Username_example") // SearchUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SearchApiV2UsersSearchPost(context.Background()).SearchUserRequest(searchUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SearchApiV2UsersSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApiV2UsersSearchPost`: SearchUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SearchApiV2UsersSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApiV2UsersSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchUserRequest** | [**SearchUserRequest**](SearchUserRequest.md) |  | 

### Return type

[**SearchUserResponse**](SearchUserResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleApiV2UsersTogglePost

> User ToggleApiV2UsersTogglePost(ctx).ToggleUserRequest(toggleUserRequest).Execute()

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
	toggleUserRequest := *openapiclient.NewToggleUserRequest("UserId_example", openapiclient.ToggleableField("enabled")) // ToggleUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ToggleApiV2UsersTogglePost(context.Background()).ToggleUserRequest(toggleUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ToggleApiV2UsersTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleApiV2UsersTogglePost`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ToggleApiV2UsersTogglePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToggleApiV2UsersTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toggleUserRequest** | [**ToggleUserRequest**](ToggleUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer), [APIKeyCookie](../README.md#APIKeyCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

