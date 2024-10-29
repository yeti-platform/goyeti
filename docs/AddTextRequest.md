# AddTextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] [default to []]
**Text** | **string** |  | 

## Methods

### NewAddTextRequest

`func NewAddTextRequest(text string, ) *AddTextRequest`

NewAddTextRequest instantiates a new AddTextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTextRequestWithDefaults

`func NewAddTextRequestWithDefaults() *AddTextRequest`

NewAddTextRequestWithDefaults instantiates a new AddTextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AddTextRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddTextRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddTextRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddTextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetText

`func (o *AddTextRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AddTextRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AddTextRequest) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


