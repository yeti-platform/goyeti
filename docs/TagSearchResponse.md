# TagSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**[]Tag**](Tag.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewTagSearchResponse

`func NewTagSearchResponse(tags []Tag, total int32, ) *TagSearchResponse`

NewTagSearchResponse instantiates a new TagSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSearchResponseWithDefaults

`func NewTagSearchResponseWithDefaults() *TagSearchResponse`

NewTagSearchResponseWithDefaults instantiates a new TagSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TagSearchResponse) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagSearchResponse) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagSearchResponse) SetTags(v []Tag)`

SetTags sets Tags field to given value.


### GetTotal

`func (o *TagSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TagSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TagSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


