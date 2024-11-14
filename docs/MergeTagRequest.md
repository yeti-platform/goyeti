# MergeTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merge** | **[]string** |  | 
**MergeInto** | **string** |  | 
**Permanent** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMergeTagRequest

`func NewMergeTagRequest(merge []string, mergeInto string, ) *MergeTagRequest`

NewMergeTagRequest instantiates a new MergeTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeTagRequestWithDefaults

`func NewMergeTagRequestWithDefaults() *MergeTagRequest`

NewMergeTagRequestWithDefaults instantiates a new MergeTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerge

`func (o *MergeTagRequest) GetMerge() []string`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *MergeTagRequest) GetMergeOk() (*[]string, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *MergeTagRequest) SetMerge(v []string)`

SetMerge sets Merge field to given value.


### GetMergeInto

`func (o *MergeTagRequest) GetMergeInto() string`

GetMergeInto returns the MergeInto field if non-nil, zero value otherwise.

### GetMergeIntoOk

`func (o *MergeTagRequest) GetMergeIntoOk() (*string, bool)`

GetMergeIntoOk returns a tuple with the MergeInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeInto

`func (o *MergeTagRequest) SetMergeInto(v string)`

SetMergeInto sets MergeInto field to given value.


### GetPermanent

`func (o *MergeTagRequest) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *MergeTagRequest) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *MergeTagRequest) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *MergeTagRequest) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


