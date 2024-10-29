# AttackPatternInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "attack-pattern"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewAttackPatternInput

`func NewAttackPatternInput(name string, ) *AttackPatternInput`

NewAttackPatternInput instantiates a new AttackPatternInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackPatternInputWithDefaults

`func NewAttackPatternInputWithDefaults() *AttackPatternInput`

NewAttackPatternInputWithDefaults instantiates a new AttackPatternInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttackPatternInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttackPatternInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttackPatternInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttackPatternInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AttackPatternInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttackPatternInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttackPatternInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AttackPatternInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttackPatternInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttackPatternInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttackPatternInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *AttackPatternInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AttackPatternInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AttackPatternInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *AttackPatternInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *AttackPatternInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AttackPatternInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AttackPatternInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AttackPatternInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *AttackPatternInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AttackPatternInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AttackPatternInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AttackPatternInput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAliases

`func (o *AttackPatternInput) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AttackPatternInput) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AttackPatternInput) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AttackPatternInput) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *AttackPatternInput) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *AttackPatternInput) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *AttackPatternInput) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *AttackPatternInput) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


