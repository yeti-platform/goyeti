# ToolInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]
**ToolVersion** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewToolInput

`func NewToolInput(name string, ) *ToolInput`

NewToolInput instantiates a new ToolInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolInputWithDefaults

`func NewToolInputWithDefaults() *ToolInput`

NewToolInputWithDefaults instantiates a new ToolInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ToolInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *ToolInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ToolInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ToolInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ToolInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *ToolInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ToolInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ToolInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ToolInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ToolInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ToolInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ToolInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ToolInput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAliases

`func (o *ToolInput) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ToolInput) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ToolInput) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ToolInput) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *ToolInput) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ToolInput) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ToolInput) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ToolInput) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetToolVersion

`func (o *ToolInput) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *ToolInput) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *ToolInput) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *ToolInput) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


