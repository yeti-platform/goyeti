# TemplateSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [default to ""]
**Count** | Pointer to **int32** |  | [optional] [default to 50]
**Page** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewTemplateSearchRequest

`func NewTemplateSearchRequest() *TemplateSearchRequest`

NewTemplateSearchRequest instantiates a new TemplateSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSearchRequestWithDefaults

`func NewTemplateSearchRequestWithDefaults() *TemplateSearchRequest`

NewTemplateSearchRequestWithDefaults instantiates a new TemplateSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateSearchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateSearchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateSearchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateSearchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCount

`func (o *TemplateSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TemplateSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TemplateSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TemplateSearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *TemplateSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TemplateSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TemplateSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *TemplateSearchRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


