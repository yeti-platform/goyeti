# DFIQSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**map[string]QueryValue**](QueryValue.md) |  | [optional] [default to {}]
**Type** | Pointer to [**DFIQSearchRequestType**](DFIQSearchRequestType.md) |  | [optional] 
**Sorting** | Pointer to **[][]interface{}** |  | [optional] [default to []]
**FilterAliases** | Pointer to **[][]interface{}** |  | [optional] [default to []]
**Count** | Pointer to **int32** |  | [optional] [default to 50]
**Page** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewDFIQSearchRequest

`func NewDFIQSearchRequest() *DFIQSearchRequest`

NewDFIQSearchRequest instantiates a new DFIQSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQSearchRequestWithDefaults

`func NewDFIQSearchRequestWithDefaults() *DFIQSearchRequest`

NewDFIQSearchRequestWithDefaults instantiates a new DFIQSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *DFIQSearchRequest) GetQuery() map[string]QueryValue`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DFIQSearchRequest) GetQueryOk() (*map[string]QueryValue, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DFIQSearchRequest) SetQuery(v map[string]QueryValue)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DFIQSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *DFIQSearchRequest) GetType() DFIQSearchRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DFIQSearchRequest) GetTypeOk() (*DFIQSearchRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DFIQSearchRequest) SetType(v DFIQSearchRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *DFIQSearchRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSorting

`func (o *DFIQSearchRequest) GetSorting() [][]interface{}`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *DFIQSearchRequest) GetSortingOk() (*[][]interface{}, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *DFIQSearchRequest) SetSorting(v [][]interface{})`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *DFIQSearchRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetFilterAliases

`func (o *DFIQSearchRequest) GetFilterAliases() [][]interface{}`

GetFilterAliases returns the FilterAliases field if non-nil, zero value otherwise.

### GetFilterAliasesOk

`func (o *DFIQSearchRequest) GetFilterAliasesOk() (*[][]interface{}, bool)`

GetFilterAliasesOk returns a tuple with the FilterAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAliases

`func (o *DFIQSearchRequest) SetFilterAliases(v [][]interface{})`

SetFilterAliases sets FilterAliases field to given value.

### HasFilterAliases

`func (o *DFIQSearchRequest) HasFilterAliases() bool`

HasFilterAliases returns a boolean if a field has been set.

### GetCount

`func (o *DFIQSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DFIQSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DFIQSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DFIQSearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *DFIQSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DFIQSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DFIQSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *DFIQSearchRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


