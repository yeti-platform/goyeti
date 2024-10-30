# Indicator1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "regex"]
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidUntil** | Pointer to **time.Time** |  | [optional] 
**Pattern** | **string** |  | 
**Location** | Pointer to **string** |  | [optional] [default to ""]
**Diamond** | [**DiamondModel**](DiamondModel.md) |  | 
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]
**RelevantTags** | Pointer to **[]string** |  | [optional] [default to []]
**Sources** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**SupportedOs** | Pointer to **[]string** |  | [optional] [default to []]
**QueryType** | **string** |  | 
**TargetSystems** | Pointer to **[]string** |  | [optional] [default to []]
**Sid** | Pointer to **int32** |  | [optional] [default to 0]
**Metadata** | Pointer to **[]string** |  | [optional] [default to []]
**References** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewIndicator1

`func NewIndicator1(name string, pattern string, diamond DiamondModel, queryType string, ) *Indicator1`

NewIndicator1 instantiates a new Indicator1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicator1WithDefaults

`func NewIndicator1WithDefaults() *Indicator1`

NewIndicator1WithDefaults instantiates a new Indicator1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Indicator1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Indicator1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Indicator1) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Indicator1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Indicator1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Indicator1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Indicator1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Indicator1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Indicator1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Indicator1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Indicator1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *Indicator1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Indicator1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Indicator1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Indicator1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Indicator1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Indicator1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Indicator1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Indicator1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *Indicator1) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *Indicator1) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *Indicator1) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *Indicator1) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *Indicator1) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *Indicator1) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *Indicator1) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *Indicator1) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *Indicator1) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Indicator1) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Indicator1) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *Indicator1) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Indicator1) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Indicator1) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Indicator1) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *Indicator1) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *Indicator1) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *Indicator1) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *Indicator1) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *Indicator1) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *Indicator1) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *Indicator1) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *Indicator1) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *Indicator1) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *Indicator1) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *Indicator1) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetSources

`func (o *Indicator1) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Indicator1) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Indicator1) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *Indicator1) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAliases

`func (o *Indicator1) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Indicator1) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Indicator1) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Indicator1) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSupportedOs

`func (o *Indicator1) GetSupportedOs() []string`

GetSupportedOs returns the SupportedOs field if non-nil, zero value otherwise.

### GetSupportedOsOk

`func (o *Indicator1) GetSupportedOsOk() (*[]string, bool)`

GetSupportedOsOk returns a tuple with the SupportedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOs

`func (o *Indicator1) SetSupportedOs(v []string)`

SetSupportedOs sets SupportedOs field to given value.

### HasSupportedOs

`func (o *Indicator1) HasSupportedOs() bool`

HasSupportedOs returns a boolean if a field has been set.

### GetQueryType

`func (o *Indicator1) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *Indicator1) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *Indicator1) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *Indicator1) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *Indicator1) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *Indicator1) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *Indicator1) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.

### GetSid

`func (o *Indicator1) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *Indicator1) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *Indicator1) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *Indicator1) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *Indicator1) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Indicator1) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Indicator1) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Indicator1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *Indicator1) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Indicator1) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Indicator1) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Indicator1) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


