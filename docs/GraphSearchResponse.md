# GraphSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vertices** | [**map[string]VerticesValue**](VerticesValue.md) |  | 
**Paths** | [**[][]PathsInnerInner**]([]PathsInnerInner.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewGraphSearchResponse

`func NewGraphSearchResponse(vertices map[string]VerticesValue, paths [][]PathsInnerInner, total int32, ) *GraphSearchResponse`

NewGraphSearchResponse instantiates a new GraphSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphSearchResponseWithDefaults

`func NewGraphSearchResponseWithDefaults() *GraphSearchResponse`

NewGraphSearchResponseWithDefaults instantiates a new GraphSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVertices

`func (o *GraphSearchResponse) GetVertices() map[string]VerticesValue`

GetVertices returns the Vertices field if non-nil, zero value otherwise.

### GetVerticesOk

`func (o *GraphSearchResponse) GetVerticesOk() (*map[string]VerticesValue, bool)`

GetVerticesOk returns a tuple with the Vertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertices

`func (o *GraphSearchResponse) SetVertices(v map[string]VerticesValue)`

SetVertices sets Vertices field to given value.


### GetPaths

`func (o *GraphSearchResponse) GetPaths() [][]PathsInnerInner`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GraphSearchResponse) GetPathsOk() (*[][]PathsInnerInner, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GraphSearchResponse) SetPaths(v [][]PathsInnerInner)`

SetPaths sets Paths field to given value.


### GetTotal

`func (o *GraphSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GraphSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GraphSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


