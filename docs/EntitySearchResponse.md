# EntitySearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]EntitiesInner**](EntitiesInner.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewEntitySearchResponse

`func NewEntitySearchResponse(entities []EntitiesInner, total int32, ) *EntitySearchResponse`

NewEntitySearchResponse instantiates a new EntitySearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySearchResponseWithDefaults

`func NewEntitySearchResponseWithDefaults() *EntitySearchResponse`

NewEntitySearchResponseWithDefaults instantiates a new EntitySearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *EntitySearchResponse) GetEntities() []EntitiesInner`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EntitySearchResponse) GetEntitiesOk() (*[]EntitiesInner, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EntitySearchResponse) SetEntities(v []EntitiesInner)`

SetEntities sets Entities field to given value.


### GetTotal

`func (o *EntitySearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EntitySearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EntitySearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


