# AnalyticsTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Status** | Pointer to [**TaskStatus**](TaskStatus.md) |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] [default to ""]
**LastRun** | Pointer to **NullableTime** |  | [optional] 
**Frequency** | Pointer to **NullableString** |  | [optional] 
**ActsOn** | Pointer to **[]string** |  | [optional] [default to []]
**Type** | Pointer to **string** |  | [optional] [default to "analytics"]
**Id** | **string** |  | [readonly] 

## Methods

### NewAnalyticsTask

`func NewAnalyticsTask(name string, id string, ) *AnalyticsTask`

NewAnalyticsTask instantiates a new AnalyticsTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsTaskWithDefaults

`func NewAnalyticsTaskWithDefaults() *AnalyticsTask`

NewAnalyticsTaskWithDefaults instantiates a new AnalyticsTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AnalyticsTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyticsTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyticsTask) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *AnalyticsTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AnalyticsTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AnalyticsTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AnalyticsTask) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *AnalyticsTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnalyticsTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnalyticsTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnalyticsTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *AnalyticsTask) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalyticsTask) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalyticsTask) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnalyticsTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AnalyticsTask) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AnalyticsTask) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AnalyticsTask) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AnalyticsTask) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastRun

`func (o *AnalyticsTask) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *AnalyticsTask) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *AnalyticsTask) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *AnalyticsTask) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *AnalyticsTask) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *AnalyticsTask) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetFrequency

`func (o *AnalyticsTask) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AnalyticsTask) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AnalyticsTask) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *AnalyticsTask) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *AnalyticsTask) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *AnalyticsTask) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetActsOn

`func (o *AnalyticsTask) GetActsOn() []string`

GetActsOn returns the ActsOn field if non-nil, zero value otherwise.

### GetActsOnOk

`func (o *AnalyticsTask) GetActsOnOk() (*[]string, bool)`

GetActsOnOk returns a tuple with the ActsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActsOn

`func (o *AnalyticsTask) SetActsOn(v []string)`

SetActsOn sets ActsOn field to given value.

### HasActsOn

`func (o *AnalyticsTask) HasActsOn() bool`

HasActsOn returns a boolean if a field has been set.

### GetType

`func (o *AnalyticsTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalyticsTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalyticsTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnalyticsTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AnalyticsTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsTask) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


