# SearchUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Count** | Pointer to **int32** |  | [optional] [default to 50]
**Page** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewSearchUserRequest

`func NewSearchUserRequest(username string, ) *SearchUserRequest`

NewSearchUserRequest instantiates a new SearchUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUserRequestWithDefaults

`func NewSearchUserRequestWithDefaults() *SearchUserRequest`

NewSearchUserRequestWithDefaults instantiates a new SearchUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SearchUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SearchUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SearchUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCount

`func (o *SearchUserRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchUserRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchUserRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchUserRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *SearchUserRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchUserRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchUserRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchUserRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


