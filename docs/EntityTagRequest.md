# EntityTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] [default to []]
**Strict** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEntityTagRequest

`func NewEntityTagRequest(ids []string, ) *EntityTagRequest`

NewEntityTagRequest instantiates a new EntityTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTagRequestWithDefaults

`func NewEntityTagRequestWithDefaults() *EntityTagRequest`

NewEntityTagRequestWithDefaults instantiates a new EntityTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *EntityTagRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *EntityTagRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *EntityTagRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetTags

`func (o *EntityTagRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityTagRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityTagRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityTagRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStrict

`func (o *EntityTagRequest) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *EntityTagRequest) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *EntityTagRequest) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *EntityTagRequest) HasStrict() bool`

HasStrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


