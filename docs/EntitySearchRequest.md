# EntitySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**map[string]QueryValue**](QueryValue.md) |  | [optional] [default to {}]
**Type** | Pointer to [**EntitySearchRequestType**](EntitySearchRequestType.md) |  | [optional] 
**Sorting** | Pointer to **[][]interface{}** |  | [optional] [default to []]
**FilterAliases** | Pointer to **[][]interface{}** |  | [optional] [default to []]
**Count** | Pointer to **int32** |  | [optional] [default to 50]
**Page** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewEntitySearchRequest

`func NewEntitySearchRequest() *EntitySearchRequest`

NewEntitySearchRequest instantiates a new EntitySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySearchRequestWithDefaults

`func NewEntitySearchRequestWithDefaults() *EntitySearchRequest`

NewEntitySearchRequestWithDefaults instantiates a new EntitySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *EntitySearchRequest) GetQuery() map[string]QueryValue`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EntitySearchRequest) GetQueryOk() (*map[string]QueryValue, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EntitySearchRequest) SetQuery(v map[string]QueryValue)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EntitySearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *EntitySearchRequest) GetType() EntitySearchRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitySearchRequest) GetTypeOk() (*EntitySearchRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitySearchRequest) SetType(v EntitySearchRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *EntitySearchRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSorting

`func (o *EntitySearchRequest) GetSorting() [][]interface{}`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *EntitySearchRequest) GetSortingOk() (*[][]interface{}, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *EntitySearchRequest) SetSorting(v [][]interface{})`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *EntitySearchRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetFilterAliases

`func (o *EntitySearchRequest) GetFilterAliases() [][]interface{}`

GetFilterAliases returns the FilterAliases field if non-nil, zero value otherwise.

### GetFilterAliasesOk

`func (o *EntitySearchRequest) GetFilterAliasesOk() (*[][]interface{}, bool)`

GetFilterAliasesOk returns a tuple with the FilterAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAliases

`func (o *EntitySearchRequest) SetFilterAliases(v [][]interface{})`

SetFilterAliases sets FilterAliases field to given value.

### HasFilterAliases

`func (o *EntitySearchRequest) HasFilterAliases() bool`

HasFilterAliases returns a boolean if a field has been set.

### GetCount

`func (o *EntitySearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EntitySearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EntitySearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EntitySearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *EntitySearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EntitySearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EntitySearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *EntitySearchRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


