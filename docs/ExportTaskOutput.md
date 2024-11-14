# ExportTaskOutput

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
**ActsOn** | Pointer to **[]string** |  | [optional] [default to []]
**TemplateName** | **string** |  | 
**Sha256** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | [readonly] 

## Methods

### NewExportTaskOutput

`func NewExportTaskOutput(name string, templateName string, id string, ) *ExportTaskOutput`

NewExportTaskOutput instantiates a new ExportTaskOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTaskOutputWithDefaults

`func NewExportTaskOutputWithDefaults() *ExportTaskOutput`

NewExportTaskOutputWithDefaults instantiates a new ExportTaskOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExportTaskOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportTaskOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportTaskOutput) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ExportTaskOutput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExportTaskOutput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExportTaskOutput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExportTaskOutput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ExportTaskOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportTaskOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportTaskOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportTaskOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ExportTaskOutput) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportTaskOutput) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportTaskOutput) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExportTaskOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ExportTaskOutput) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ExportTaskOutput) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ExportTaskOutput) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ExportTaskOutput) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastRun

`func (o *ExportTaskOutput) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ExportTaskOutput) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ExportTaskOutput) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ExportTaskOutput) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *ExportTaskOutput) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *ExportTaskOutput) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetFrequency

`func (o *ExportTaskOutput) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ExportTaskOutput) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ExportTaskOutput) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ExportTaskOutput) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *ExportTaskOutput) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *ExportTaskOutput) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetType

`func (o *ExportTaskOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportTaskOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportTaskOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExportTaskOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIncludeTags

`func (o *ExportTaskOutput) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ExportTaskOutput) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ExportTaskOutput) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ExportTaskOutput) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetExcludeTags

`func (o *ExportTaskOutput) GetExcludeTags() []string`

GetExcludeTags returns the ExcludeTags field if non-nil, zero value otherwise.

### GetExcludeTagsOk

`func (o *ExportTaskOutput) GetExcludeTagsOk() (*[]string, bool)`

GetExcludeTagsOk returns a tuple with the ExcludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTags

`func (o *ExportTaskOutput) SetExcludeTags(v []string)`

SetExcludeTags sets ExcludeTags field to given value.

### HasExcludeTags

`func (o *ExportTaskOutput) HasExcludeTags() bool`

HasExcludeTags returns a boolean if a field has been set.

### GetIgnoreTags

`func (o *ExportTaskOutput) GetIgnoreTags() []string`

GetIgnoreTags returns the IgnoreTags field if non-nil, zero value otherwise.

### GetIgnoreTagsOk

`func (o *ExportTaskOutput) GetIgnoreTagsOk() (*[]string, bool)`

GetIgnoreTagsOk returns a tuple with the IgnoreTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTags

`func (o *ExportTaskOutput) SetIgnoreTags(v []string)`

SetIgnoreTags sets IgnoreTags field to given value.

### HasIgnoreTags

`func (o *ExportTaskOutput) HasIgnoreTags() bool`

HasIgnoreTags returns a boolean if a field has been set.

### GetFreshTags

`func (o *ExportTaskOutput) GetFreshTags() bool`

GetFreshTags returns the FreshTags field if non-nil, zero value otherwise.

### GetFreshTagsOk

`func (o *ExportTaskOutput) GetFreshTagsOk() (*bool, bool)`

GetFreshTagsOk returns a tuple with the FreshTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshTags

`func (o *ExportTaskOutput) SetFreshTags(v bool)`

SetFreshTags sets FreshTags field to given value.

### HasFreshTags

`func (o *ExportTaskOutput) HasFreshTags() bool`

HasFreshTags returns a boolean if a field has been set.

### GetActsOn

`func (o *ExportTaskOutput) GetActsOn() []string`

GetActsOn returns the ActsOn field if non-nil, zero value otherwise.

### GetActsOnOk

`func (o *ExportTaskOutput) GetActsOnOk() (*[]string, bool)`

GetActsOnOk returns a tuple with the ActsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActsOn

`func (o *ExportTaskOutput) SetActsOn(v []string)`

SetActsOn sets ActsOn field to given value.

### HasActsOn

`func (o *ExportTaskOutput) HasActsOn() bool`

HasActsOn returns a boolean if a field has been set.

### GetTemplateName

`func (o *ExportTaskOutput) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ExportTaskOutput) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ExportTaskOutput) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetSha256

`func (o *ExportTaskOutput) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ExportTaskOutput) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ExportTaskOutput) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ExportTaskOutput) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *ExportTaskOutput) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *ExportTaskOutput) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetId

`func (o *ExportTaskOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportTaskOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportTaskOutput) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


