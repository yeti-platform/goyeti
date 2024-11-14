# OneShotTask

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
**Type** | Pointer to **string** |  | [optional] [default to "oneshot"]
**ActsOn** | Pointer to **[]string** |  | [optional] [default to []]
**Id** | **string** |  | [readonly] 

## Methods

### NewOneShotTask

`func NewOneShotTask(name string, id string, ) *OneShotTask`

NewOneShotTask instantiates a new OneShotTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneShotTaskWithDefaults

`func NewOneShotTaskWithDefaults() *OneShotTask`

NewOneShotTaskWithDefaults instantiates a new OneShotTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OneShotTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OneShotTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OneShotTask) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *OneShotTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OneShotTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OneShotTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OneShotTask) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *OneShotTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OneShotTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OneShotTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OneShotTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *OneShotTask) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OneShotTask) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OneShotTask) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OneShotTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *OneShotTask) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *OneShotTask) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *OneShotTask) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *OneShotTask) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastRun

`func (o *OneShotTask) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *OneShotTask) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *OneShotTask) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *OneShotTask) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *OneShotTask) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *OneShotTask) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetFrequency

`func (o *OneShotTask) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *OneShotTask) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *OneShotTask) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *OneShotTask) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *OneShotTask) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *OneShotTask) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetType

`func (o *OneShotTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OneShotTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OneShotTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OneShotTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActsOn

`func (o *OneShotTask) GetActsOn() []string`

GetActsOn returns the ActsOn field if non-nil, zero value otherwise.

### GetActsOnOk

`func (o *OneShotTask) GetActsOnOk() (*[]string, bool)`

GetActsOnOk returns a tuple with the ActsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActsOn

`func (o *OneShotTask) SetActsOn(v []string)`

SetActsOn sets ActsOn field to given value.

### HasActsOn

`func (o *OneShotTask) HasActsOn() bool`

HasActsOn returns a boolean if a field has been set.

### GetId

`func (o *OneShotTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OneShotTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OneShotTask) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


