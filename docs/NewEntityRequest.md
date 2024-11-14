# NewEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | [**Entity1**](Entity1.md) |  | 
**Tags** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewNewEntityRequest

`func NewNewEntityRequest(entity Entity1, ) *NewEntityRequest`

NewNewEntityRequest instantiates a new NewEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewEntityRequestWithDefaults

`func NewNewEntityRequestWithDefaults() *NewEntityRequest`

NewNewEntityRequestWithDefaults instantiates a new NewEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *NewEntityRequest) GetEntity() Entity1`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *NewEntityRequest) GetEntityOk() (*Entity1, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *NewEntityRequest) SetEntity(v Entity1)`

SetEntity sets Entity field to given value.


### GetTags

`func (o *NewEntityRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NewEntityRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NewEntityRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NewEntityRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


