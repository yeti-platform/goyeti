# EntityTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | **int32** |  | 
**Tags** | [**map[string]map[string]TagRelationshipOutput**](map.md) |  | 

## Methods

### NewEntityTagResponse

`func NewEntityTagResponse(tagged int32, tags map[string]map[string]TagRelationshipOutput, ) *EntityTagResponse`

NewEntityTagResponse instantiates a new EntityTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTagResponseWithDefaults

`func NewEntityTagResponseWithDefaults() *EntityTagResponse`

NewEntityTagResponseWithDefaults instantiates a new EntityTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *EntityTagResponse) GetTagged() int32`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *EntityTagResponse) GetTaggedOk() (*int32, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *EntityTagResponse) SetTagged(v int32)`

SetTagged sets Tagged field to given value.


### GetTags

`func (o *EntityTagResponse) GetTags() map[string]map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityTagResponse) GetTagsOk() (*map[string]map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityTagResponse) SetTags(v map[string]map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


