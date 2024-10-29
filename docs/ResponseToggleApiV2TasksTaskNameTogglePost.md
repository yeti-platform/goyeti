# ResponseToggleApiV2TasksTaskNameTogglePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Status** | Pointer to [**TaskStatus**](TaskStatus.md) |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] [default to ""]
**LastRun** | Pointer to **time.Time** |  | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "feed"]
**Id** | **string** |  | [readonly] 
**ActsOn** | Pointer to **[]string** |  | [optional] [default to []]
**IncludeTags** | Pointer to **[]string** |  | [optional] [default to []]
**ExcludeTags** | Pointer to **[]string** |  | [optional] [default to []]
**IgnoreTags** | Pointer to **[]string** |  | [optional] [default to []]
**FreshTags** | Pointer to **bool** |  | [optional] [default to true]
**TemplateName** | **string** |  | 
**Sha256** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseToggleApiV2TasksTaskNameTogglePost

`func NewResponseToggleApiV2TasksTaskNameTogglePost(name string, id string, templateName string, ) *ResponseToggleApiV2TasksTaskNameTogglePost`

NewResponseToggleApiV2TasksTaskNameTogglePost instantiates a new ResponseToggleApiV2TasksTaskNameTogglePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseToggleApiV2TasksTaskNameTogglePostWithDefaults

`func NewResponseToggleApiV2TasksTaskNameTogglePostWithDefaults() *ResponseToggleApiV2TasksTaskNameTogglePost`

NewResponseToggleApiV2TasksTaskNameTogglePostWithDefaults instantiates a new ResponseToggleApiV2TasksTaskNameTogglePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastRun

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetFrequency

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetType

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetId(v string)`

SetId sets Id field to given value.


### GetActsOn

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetActsOn() []string`

GetActsOn returns the ActsOn field if non-nil, zero value otherwise.

### GetActsOnOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetActsOnOk() (*[]string, bool)`

GetActsOnOk returns a tuple with the ActsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActsOn

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetActsOn(v []string)`

SetActsOn sets ActsOn field to given value.

### HasActsOn

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasActsOn() bool`

HasActsOn returns a boolean if a field has been set.

### GetIncludeTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetExcludeTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetExcludeTags() []string`

GetExcludeTags returns the ExcludeTags field if non-nil, zero value otherwise.

### GetExcludeTagsOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetExcludeTagsOk() (*[]string, bool)`

GetExcludeTagsOk returns a tuple with the ExcludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetExcludeTags(v []string)`

SetExcludeTags sets ExcludeTags field to given value.

### HasExcludeTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasExcludeTags() bool`

HasExcludeTags returns a boolean if a field has been set.

### GetIgnoreTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetIgnoreTags() []string`

GetIgnoreTags returns the IgnoreTags field if non-nil, zero value otherwise.

### GetIgnoreTagsOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetIgnoreTagsOk() (*[]string, bool)`

GetIgnoreTagsOk returns a tuple with the IgnoreTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetIgnoreTags(v []string)`

SetIgnoreTags sets IgnoreTags field to given value.

### HasIgnoreTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasIgnoreTags() bool`

HasIgnoreTags returns a boolean if a field has been set.

### GetFreshTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetFreshTags() bool`

GetFreshTags returns the FreshTags field if non-nil, zero value otherwise.

### GetFreshTagsOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetFreshTagsOk() (*bool, bool)`

GetFreshTagsOk returns a tuple with the FreshTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetFreshTags(v bool)`

SetFreshTags sets FreshTags field to given value.

### HasFreshTags

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasFreshTags() bool`

HasFreshTags returns a boolean if a field has been set.

### GetTemplateName

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetSha256

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResponseToggleApiV2TasksTaskNameTogglePost) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


