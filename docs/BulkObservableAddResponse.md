# BulkObservableAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to [**[]AddedInner**](AddedInner.md) |  | [optional] [default to []]
**Failed** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewBulkObservableAddResponse

`func NewBulkObservableAddResponse() *BulkObservableAddResponse`

NewBulkObservableAddResponse instantiates a new BulkObservableAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkObservableAddResponseWithDefaults

`func NewBulkObservableAddResponseWithDefaults() *BulkObservableAddResponse`

NewBulkObservableAddResponseWithDefaults instantiates a new BulkObservableAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *BulkObservableAddResponse) GetAdded() []AddedInner`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *BulkObservableAddResponse) GetAddedOk() (*[]AddedInner, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *BulkObservableAddResponse) SetAdded(v []AddedInner)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *BulkObservableAddResponse) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetFailed

`func (o *BulkObservableAddResponse) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkObservableAddResponse) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkObservableAddResponse) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *BulkObservableAddResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


