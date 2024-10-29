# ObservableSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Observables** | [**[]AddedInner**](AddedInner.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewObservableSearchResponse

`func NewObservableSearchResponse(observables []AddedInner, total int32, ) *ObservableSearchResponse`

NewObservableSearchResponse instantiates a new ObservableSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservableSearchResponseWithDefaults

`func NewObservableSearchResponseWithDefaults() *ObservableSearchResponse`

NewObservableSearchResponseWithDefaults instantiates a new ObservableSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservables

`func (o *ObservableSearchResponse) GetObservables() []AddedInner`

GetObservables returns the Observables field if non-nil, zero value otherwise.

### GetObservablesOk

`func (o *ObservableSearchResponse) GetObservablesOk() (*[]AddedInner, bool)`

GetObservablesOk returns a tuple with the Observables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservables

`func (o *ObservableSearchResponse) SetObservables(v []AddedInner)`

SetObservables sets Observables field to given value.


### GetTotal

`func (o *ObservableSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ObservableSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ObservableSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


