# ObservableTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] [default to []]
**Ids** | **[]string** |  | 
**Strict** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewObservableTagRequest

`func NewObservableTagRequest(ids []string, ) *ObservableTagRequest`

NewObservableTagRequest instantiates a new ObservableTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservableTagRequestWithDefaults

`func NewObservableTagRequestWithDefaults() *ObservableTagRequest`

NewObservableTagRequestWithDefaults instantiates a new ObservableTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ObservableTagRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ObservableTagRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ObservableTagRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ObservableTagRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIds

`func (o *ObservableTagRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ObservableTagRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ObservableTagRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetStrict

`func (o *ObservableTagRequest) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *ObservableTagRequest) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *ObservableTagRequest) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *ObservableTagRequest) HasStrict() bool`

HasStrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


