# IndicatorSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indicators** | [**[]IndicatorsInner**](IndicatorsInner.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewIndicatorSearchResponse

`func NewIndicatorSearchResponse(indicators []IndicatorsInner, total int32, ) *IndicatorSearchResponse`

NewIndicatorSearchResponse instantiates a new IndicatorSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorSearchResponseWithDefaults

`func NewIndicatorSearchResponseWithDefaults() *IndicatorSearchResponse`

NewIndicatorSearchResponseWithDefaults instantiates a new IndicatorSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicators

`func (o *IndicatorSearchResponse) GetIndicators() []IndicatorsInner`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *IndicatorSearchResponse) GetIndicatorsOk() (*[]IndicatorsInner, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *IndicatorSearchResponse) SetIndicators(v []IndicatorsInner)`

SetIndicators sets Indicators field to given value.


### GetTotal

`func (o *IndicatorSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IndicatorSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IndicatorSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


