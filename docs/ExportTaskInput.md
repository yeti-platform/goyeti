# ExportTaskInput

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
**Type** | Pointer to **string** |  | [optional] [default to "export"]
**IncludeTags** | Pointer to **[]string** |  | [optional] [default to []]
**ExcludeTags** | Pointer to **[]string** |  | [optional] [default to []]
**IgnoreTags** | Pointer to **[]string** |  | [optional] [default to []]
**FreshTags** | Pointer to **bool** |  | [optional] [default to true]
**ActsOn** | Pointer to [**[]ObservableTypeInput**](ObservableTypeInput.md) |  | [optional] [default to []]
**TemplateName** | **string** |  | 
**Sha256** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExportTaskInput

`func NewExportTaskInput(name string, templateName string, ) *ExportTaskInput`

NewExportTaskInput instantiates a new ExportTaskInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTaskInputWithDefaults

`func NewExportTaskInputWithDefaults() *ExportTaskInput`

NewExportTaskInputWithDefaults instantiates a new ExportTaskInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExportTaskInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportTaskInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportTaskInput) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ExportTaskInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExportTaskInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExportTaskInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExportTaskInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ExportTaskInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportTaskInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportTaskInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportTaskInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ExportTaskInput) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportTaskInput) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportTaskInput) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExportTaskInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ExportTaskInput) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ExportTaskInput) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ExportTaskInput) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ExportTaskInput) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastRun

`func (o *ExportTaskInput) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ExportTaskInput) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ExportTaskInput) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ExportTaskInput) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *ExportTaskInput) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *ExportTaskInput) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetFrequency

`func (o *ExportTaskInput) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ExportTaskInput) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ExportTaskInput) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ExportTaskInput) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *ExportTaskInput) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *ExportTaskInput) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetType

`func (o *ExportTaskInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportTaskInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportTaskInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExportTaskInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIncludeTags

`func (o *ExportTaskInput) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ExportTaskInput) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ExportTaskInput) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ExportTaskInput) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetExcludeTags

`func (o *ExportTaskInput) GetExcludeTags() []string`

GetExcludeTags returns the ExcludeTags field if non-nil, zero value otherwise.

### GetExcludeTagsOk

`func (o *ExportTaskInput) GetExcludeTagsOk() (*[]string, bool)`

GetExcludeTagsOk returns a tuple with the ExcludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTags

`func (o *ExportTaskInput) SetExcludeTags(v []string)`

SetExcludeTags sets ExcludeTags field to given value.

### HasExcludeTags

`func (o *ExportTaskInput) HasExcludeTags() bool`

HasExcludeTags returns a boolean if a field has been set.

### GetIgnoreTags

`func (o *ExportTaskInput) GetIgnoreTags() []string`

GetIgnoreTags returns the IgnoreTags field if non-nil, zero value otherwise.

### GetIgnoreTagsOk

`func (o *ExportTaskInput) GetIgnoreTagsOk() (*[]string, bool)`

GetIgnoreTagsOk returns a tuple with the IgnoreTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTags

`func (o *ExportTaskInput) SetIgnoreTags(v []string)`

SetIgnoreTags sets IgnoreTags field to given value.

### HasIgnoreTags

`func (o *ExportTaskInput) HasIgnoreTags() bool`

HasIgnoreTags returns a boolean if a field has been set.

### GetFreshTags

`func (o *ExportTaskInput) GetFreshTags() bool`

GetFreshTags returns the FreshTags field if non-nil, zero value otherwise.

### GetFreshTagsOk

`func (o *ExportTaskInput) GetFreshTagsOk() (*bool, bool)`

GetFreshTagsOk returns a tuple with the FreshTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshTags

`func (o *ExportTaskInput) SetFreshTags(v bool)`

SetFreshTags sets FreshTags field to given value.

### HasFreshTags

`func (o *ExportTaskInput) HasFreshTags() bool`

HasFreshTags returns a boolean if a field has been set.

### GetActsOn

`func (o *ExportTaskInput) GetActsOn() []ObservableTypeInput`

GetActsOn returns the ActsOn field if non-nil, zero value otherwise.

### GetActsOnOk

`func (o *ExportTaskInput) GetActsOnOk() (*[]ObservableTypeInput, bool)`

GetActsOnOk returns a tuple with the ActsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActsOn

`func (o *ExportTaskInput) SetActsOn(v []ObservableTypeInput)`

SetActsOn sets ActsOn field to given value.

### HasActsOn

`func (o *ExportTaskInput) HasActsOn() bool`

HasActsOn returns a boolean if a field has been set.

### GetTemplateName

`func (o *ExportTaskInput) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ExportTaskInput) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ExportTaskInput) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetSha256

`func (o *ExportTaskInput) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ExportTaskInput) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ExportTaskInput) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ExportTaskInput) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *ExportTaskInput) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *ExportTaskInput) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


