# MergeTagResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merged** | **int32** |  | 
**Into** | [**Tag**](Tag.md) |  | 

## Methods

### NewMergeTagResult

`func NewMergeTagResult(merged int32, into Tag, ) *MergeTagResult`

NewMergeTagResult instantiates a new MergeTagResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeTagResultWithDefaults

`func NewMergeTagResultWithDefaults() *MergeTagResult`

NewMergeTagResultWithDefaults instantiates a new MergeTagResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerged

`func (o *MergeTagResult) GetMerged() int32`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *MergeTagResult) GetMergedOk() (*int32, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *MergeTagResult) SetMerged(v int32)`

SetMerged sets Merged field to given value.


### GetInto

`func (o *MergeTagResult) GetInto() Tag`

GetInto returns the Into field if non-nil, zero value otherwise.

### GetIntoOk

`func (o *MergeTagResult) GetIntoOk() (*Tag, bool)`

GetIntoOk returns a tuple with the Into field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInto

`func (o *MergeTagResult) SetInto(v Tag)`

SetInto sets Into field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


