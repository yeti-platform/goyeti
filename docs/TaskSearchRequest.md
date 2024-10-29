# TaskSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**map[string]QueryValue**](QueryValue.md) |  | [optional] [default to {}]
**Type** | Pointer to [**TaskSearchRequestType**](TaskSearchRequestType.md) |  | [optional] 
**Sorting** | Pointer to **[][]interface{}** |  | [optional] [default to []]
**Count** | Pointer to **int32** |  | [optional] [default to 100]
**Page** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewTaskSearchRequest

`func NewTaskSearchRequest() *TaskSearchRequest`

NewTaskSearchRequest instantiates a new TaskSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSearchRequestWithDefaults

`func NewTaskSearchRequestWithDefaults() *TaskSearchRequest`

NewTaskSearchRequestWithDefaults instantiates a new TaskSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *TaskSearchRequest) GetQuery() map[string]QueryValue`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TaskSearchRequest) GetQueryOk() (*map[string]QueryValue, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TaskSearchRequest) SetQuery(v map[string]QueryValue)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TaskSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *TaskSearchRequest) GetType() TaskSearchRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskSearchRequest) GetTypeOk() (*TaskSearchRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskSearchRequest) SetType(v TaskSearchRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *TaskSearchRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSorting

`func (o *TaskSearchRequest) GetSorting() [][]interface{}`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *TaskSearchRequest) GetSortingOk() (*[][]interface{}, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *TaskSearchRequest) SetSorting(v [][]interface{})`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *TaskSearchRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetCount

`func (o *TaskSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TaskSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TaskSearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *TaskSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TaskSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TaskSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *TaskSearchRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


