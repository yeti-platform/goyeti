# GraphSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**LinkTypes** | Pointer to **[]string** |  | [optional] [default to []]
**TargetTypes** | Pointer to **[]string** |  | [optional] [default to []]
**Hops** | Pointer to [**Hops**](Hops.md) |  | [optional] 
**MinHops** | Pointer to [**MinHops**](MinHops.md) |  | [optional] 
**MaxHops** | Pointer to [**MaxHops**](MaxHops.md) |  | [optional] 
**Graph** | **string** |  | 
**Direction** | [**GraphDirection**](GraphDirection.md) |  | 
**Filter** | Pointer to [**[]GraphFilter**](GraphFilter.md) |  | [optional] [default to []]
**IncludeOriginal** | **bool** |  | 
**Count** | Pointer to **int32** |  | [optional] [default to 50]
**Page** | Pointer to **int32** |  | [optional] [default to 0]
**Sorting** | Pointer to **[][]interface{}** |  | [optional] [default to []]

## Methods

### NewGraphSearchRequest

`func NewGraphSearchRequest(source string, graph string, direction GraphDirection, includeOriginal bool, ) *GraphSearchRequest`

NewGraphSearchRequest instantiates a new GraphSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphSearchRequestWithDefaults

`func NewGraphSearchRequestWithDefaults() *GraphSearchRequest`

NewGraphSearchRequestWithDefaults instantiates a new GraphSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *GraphSearchRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphSearchRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphSearchRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetLinkTypes

`func (o *GraphSearchRequest) GetLinkTypes() []string`

GetLinkTypes returns the LinkTypes field if non-nil, zero value otherwise.

### GetLinkTypesOk

`func (o *GraphSearchRequest) GetLinkTypesOk() (*[]string, bool)`

GetLinkTypesOk returns a tuple with the LinkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypes

`func (o *GraphSearchRequest) SetLinkTypes(v []string)`

SetLinkTypes sets LinkTypes field to given value.

### HasLinkTypes

`func (o *GraphSearchRequest) HasLinkTypes() bool`

HasLinkTypes returns a boolean if a field has been set.

### GetTargetTypes

`func (o *GraphSearchRequest) GetTargetTypes() []string`

GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.

### GetTargetTypesOk

`func (o *GraphSearchRequest) GetTargetTypesOk() (*[]string, bool)`

GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTypes

`func (o *GraphSearchRequest) SetTargetTypes(v []string)`

SetTargetTypes sets TargetTypes field to given value.

### HasTargetTypes

`func (o *GraphSearchRequest) HasTargetTypes() bool`

HasTargetTypes returns a boolean if a field has been set.

### GetHops

`func (o *GraphSearchRequest) GetHops() Hops`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *GraphSearchRequest) GetHopsOk() (*Hops, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *GraphSearchRequest) SetHops(v Hops)`

SetHops sets Hops field to given value.

### HasHops

`func (o *GraphSearchRequest) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetMinHops

`func (o *GraphSearchRequest) GetMinHops() MinHops`

GetMinHops returns the MinHops field if non-nil, zero value otherwise.

### GetMinHopsOk

`func (o *GraphSearchRequest) GetMinHopsOk() (*MinHops, bool)`

GetMinHopsOk returns a tuple with the MinHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHops

`func (o *GraphSearchRequest) SetMinHops(v MinHops)`

SetMinHops sets MinHops field to given value.

### HasMinHops

`func (o *GraphSearchRequest) HasMinHops() bool`

HasMinHops returns a boolean if a field has been set.

### GetMaxHops

`func (o *GraphSearchRequest) GetMaxHops() MaxHops`

GetMaxHops returns the MaxHops field if non-nil, zero value otherwise.

### GetMaxHopsOk

`func (o *GraphSearchRequest) GetMaxHopsOk() (*MaxHops, bool)`

GetMaxHopsOk returns a tuple with the MaxHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHops

`func (o *GraphSearchRequest) SetMaxHops(v MaxHops)`

SetMaxHops sets MaxHops field to given value.

### HasMaxHops

`func (o *GraphSearchRequest) HasMaxHops() bool`

HasMaxHops returns a boolean if a field has been set.

### GetGraph

`func (o *GraphSearchRequest) GetGraph() string`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *GraphSearchRequest) GetGraphOk() (*string, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *GraphSearchRequest) SetGraph(v string)`

SetGraph sets Graph field to given value.


### GetDirection

`func (o *GraphSearchRequest) GetDirection() GraphDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GraphSearchRequest) GetDirectionOk() (*GraphDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GraphSearchRequest) SetDirection(v GraphDirection)`

SetDirection sets Direction field to given value.


### GetFilter

`func (o *GraphSearchRequest) GetFilter() []GraphFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GraphSearchRequest) GetFilterOk() (*[]GraphFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GraphSearchRequest) SetFilter(v []GraphFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GraphSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIncludeOriginal

`func (o *GraphSearchRequest) GetIncludeOriginal() bool`

GetIncludeOriginal returns the IncludeOriginal field if non-nil, zero value otherwise.

### GetIncludeOriginalOk

`func (o *GraphSearchRequest) GetIncludeOriginalOk() (*bool, bool)`

GetIncludeOriginalOk returns a tuple with the IncludeOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOriginal

`func (o *GraphSearchRequest) SetIncludeOriginal(v bool)`

SetIncludeOriginal sets IncludeOriginal field to given value.


### GetCount

`func (o *GraphSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GraphSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GraphSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GraphSearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *GraphSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GraphSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GraphSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GraphSearchRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSorting

`func (o *GraphSearchRequest) GetSorting() [][]interface{}`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *GraphSearchRequest) GetSortingOk() (*[][]interface{}, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *GraphSearchRequest) SetSorting(v [][]interface{})`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *GraphSearchRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


