# DFIQSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dfiq** | [**[]DfiqInner**](DfiqInner.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewDFIQSearchResponse

`func NewDFIQSearchResponse(dfiq []DfiqInner, total int32, ) *DFIQSearchResponse`

NewDFIQSearchResponse instantiates a new DFIQSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQSearchResponseWithDefaults

`func NewDFIQSearchResponseWithDefaults() *DFIQSearchResponse`

NewDFIQSearchResponseWithDefaults instantiates a new DFIQSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfiq

`func (o *DFIQSearchResponse) GetDfiq() []DfiqInner`

GetDfiq returns the Dfiq field if non-nil, zero value otherwise.

### GetDfiqOk

`func (o *DFIQSearchResponse) GetDfiqOk() (*[]DfiqInner, bool)`

GetDfiqOk returns a tuple with the Dfiq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiq

`func (o *DFIQSearchResponse) SetDfiq(v []DfiqInner)`

SetDfiq sets Dfiq field to given value.


### GetTotal

`func (o *DFIQSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DFIQSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DFIQSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


