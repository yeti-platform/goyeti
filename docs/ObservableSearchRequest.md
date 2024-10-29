# ObservableSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**map[string]QueryValue**](QueryValue.md) |  | [optional] [default to {}]
**Type** | Pointer to [**ObservableSearchRequestType**](ObservableSearchRequestType.md) |  | [optional] 
**Sorting** | Pointer to **[][]interface{}** |  | [optional] [default to []]
**Count** | Pointer to **int32** |  | [optional] [default to 50]
**Page** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewObservableSearchRequest

`func NewObservableSearchRequest() *ObservableSearchRequest`

NewObservableSearchRequest instantiates a new ObservableSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservableSearchRequestWithDefaults

`func NewObservableSearchRequestWithDefaults() *ObservableSearchRequest`

NewObservableSearchRequestWithDefaults instantiates a new ObservableSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ObservableSearchRequest) GetQuery() map[string]QueryValue`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ObservableSearchRequest) GetQueryOk() (*map[string]QueryValue, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ObservableSearchRequest) SetQuery(v map[string]QueryValue)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ObservableSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *ObservableSearchRequest) GetType() ObservableSearchRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObservableSearchRequest) GetTypeOk() (*ObservableSearchRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObservableSearchRequest) SetType(v ObservableSearchRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *ObservableSearchRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSorting

`func (o *ObservableSearchRequest) GetSorting() [][]interface{}`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *ObservableSearchRequest) GetSortingOk() (*[][]interface{}, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *ObservableSearchRequest) SetSorting(v [][]interface{})`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *ObservableSearchRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### GetCount

`func (o *ObservableSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ObservableSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ObservableSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ObservableSearchRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *ObservableSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ObservableSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ObservableSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ObservableSearchRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


