# TagSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**Produces** | Pointer to **[]string** |  | [optional] [default to []]
**Replaces** | Pointer to **[]string** |  | [optional] [default to []]
**Count** | **int32** |  | 
**Page** | **int32** |  | 

## Methods

### NewTagSearchRequest

`func NewTagSearchRequest(count int32, page int32, ) *TagSearchRequest`

NewTagSearchRequest instantiates a new TagSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSearchRequestWithDefaults

`func NewTagSearchRequestWithDefaults() *TagSearchRequest`

NewTagSearchRequestWithDefaults instantiates a new TagSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagSearchRequest) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagSearchRequest) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagSearchRequest) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *TagSearchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProduces

`func (o *TagSearchRequest) GetProduces() []string`

GetProduces returns the Produces field if non-nil, zero value otherwise.

### GetProducesOk

`func (o *TagSearchRequest) GetProducesOk() (*[]string, bool)`

GetProducesOk returns a tuple with the Produces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduces

`func (o *TagSearchRequest) SetProduces(v []string)`

SetProduces sets Produces field to given value.

### HasProduces

`func (o *TagSearchRequest) HasProduces() bool`

HasProduces returns a boolean if a field has been set.

### GetReplaces

`func (o *TagSearchRequest) GetReplaces() []string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *TagSearchRequest) GetReplacesOk() (*[]string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *TagSearchRequest) SetReplaces(v []string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *TagSearchRequest) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.

### GetCount

`func (o *TagSearchRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TagSearchRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TagSearchRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TagSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TagSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TagSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


