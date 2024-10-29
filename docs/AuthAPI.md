# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginApiApiV2AuthApiTokenPost**](AuthAPI.md#LoginApiApiV2AuthApiTokenPost) | **Post** /api/v2/auth/api-token | Login Api
[**LoginApiV2AuthTokenPost**](AuthAPI.md#LoginApiV2AuthTokenPost) | **Post** /api/v2/auth/token | Login
[**LogoutApiV2AuthLogoutPost**](AuthAPI.md#LogoutApiV2AuthLogoutPost) | **Post** /api/v2/auth/logout | Logout
[**MeApiV2AuthMeGet**](AuthAPI.md#MeApiV2AuthMeGet) | **Get** /api/v2/auth/me | Me



## LoginApiApiV2AuthApiTokenPost

> interface{} LoginApiApiV2AuthApiTokenPost(ctx).Execute()

Login Api

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
	resp, r, err := apiClient.AuthAPI.LoginApiApiV2AuthApiTokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.LoginApiApiV2AuthApiTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginApiApiV2AuthApiTokenPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.LoginApiApiV2AuthApiTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginApiApiV2AuthApiTokenPostRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginApiV2AuthTokenPost

> interface{} LoginApiV2AuthTokenPost(ctx).Username(username).Password(password).GrantType(grantType).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).Execute()

Login

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
	username := "username_example" // string | 
	password := "password_example" // string | 
	grantType := "grantType_example" // string |  (optional)
	scope := "scope_example" // string |  (optional) (default to "")
	clientId := "clientId_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.LoginApiV2AuthTokenPost(context.Background()).Username(username).Password(password).GrantType(grantType).Scope(scope).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.LoginApiV2AuthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginApiV2AuthTokenPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.LoginApiV2AuthTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginApiV2AuthTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **password** | **string** |  | 
 **grantType** | **string** |  | 
 **scope** | **string** |  | [default to &quot;&quot;]
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutApiV2AuthLogoutPost

> interface{} LogoutApiV2AuthLogoutPost(ctx).Execute()

Logout

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
	resp, r, err := apiClient.AuthAPI.LogoutApiV2AuthLogoutPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.LogoutApiV2AuthLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogoutApiV2AuthLogoutPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.LogoutApiV2AuthLogoutPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutApiV2AuthLogoutPostRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeApiV2AuthMeGet

> User MeApiV2AuthMeGet(ctx).Execute()

Me

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
	resp, r, err := apiClient.AuthAPI.MeApiV2AuthMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.MeApiV2AuthMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeApiV2AuthMeGet`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.MeApiV2AuthMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMeApiV2AuthMeGetRequest struct via the builder pattern


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

