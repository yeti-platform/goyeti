# IndicatorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "query"]
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
**QueryType** | **string** |  | 
**TargetSystems** | Pointer to **[]string** |  | [optional] [default to []]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**Sources** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**SupportedOs** | Pointer to **[]string** |  | [optional] [default to []]
**Sid** | Pointer to **int32** |  | [optional] [default to 0]
**Metadata** | Pointer to **[]string** |  | [optional] [default to []]
**References** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewIndicatorsInner

`func NewIndicatorsInner(name string, pattern string, diamond DiamondModel, queryType string, id string, tags map[string]TagRelationshipOutput, rootType string, ) *IndicatorsInner`

NewIndicatorsInner instantiates a new IndicatorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorsInnerWithDefaults

`func NewIndicatorsInnerWithDefaults() *IndicatorsInner`

NewIndicatorsInnerWithDefaults instantiates a new IndicatorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IndicatorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndicatorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndicatorsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IndicatorsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndicatorsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndicatorsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IndicatorsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *IndicatorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndicatorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndicatorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IndicatorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *IndicatorsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IndicatorsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IndicatorsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IndicatorsInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *IndicatorsInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IndicatorsInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IndicatorsInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IndicatorsInner) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *IndicatorsInner) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *IndicatorsInner) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *IndicatorsInner) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *IndicatorsInner) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *IndicatorsInner) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *IndicatorsInner) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *IndicatorsInner) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *IndicatorsInner) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *IndicatorsInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *IndicatorsInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *IndicatorsInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *IndicatorsInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IndicatorsInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IndicatorsInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IndicatorsInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *IndicatorsInner) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *IndicatorsInner) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *IndicatorsInner) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *IndicatorsInner) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *IndicatorsInner) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *IndicatorsInner) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *IndicatorsInner) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *IndicatorsInner) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *IndicatorsInner) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *IndicatorsInner) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *IndicatorsInner) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetQueryType

`func (o *IndicatorsInner) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *IndicatorsInner) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *IndicatorsInner) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *IndicatorsInner) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *IndicatorsInner) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *IndicatorsInner) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *IndicatorsInner) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.

### GetId

`func (o *IndicatorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndicatorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndicatorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *IndicatorsInner) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IndicatorsInner) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IndicatorsInner) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *IndicatorsInner) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *IndicatorsInner) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *IndicatorsInner) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetSources

`func (o *IndicatorsInner) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IndicatorsInner) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IndicatorsInner) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *IndicatorsInner) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAliases

`func (o *IndicatorsInner) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *IndicatorsInner) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *IndicatorsInner) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *IndicatorsInner) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSupportedOs

`func (o *IndicatorsInner) GetSupportedOs() []string`

GetSupportedOs returns the SupportedOs field if non-nil, zero value otherwise.

### GetSupportedOsOk

`func (o *IndicatorsInner) GetSupportedOsOk() (*[]string, bool)`

GetSupportedOsOk returns a tuple with the SupportedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOs

`func (o *IndicatorsInner) SetSupportedOs(v []string)`

SetSupportedOs sets SupportedOs field to given value.

### HasSupportedOs

`func (o *IndicatorsInner) HasSupportedOs() bool`

HasSupportedOs returns a boolean if a field has been set.

### GetSid

`func (o *IndicatorsInner) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IndicatorsInner) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IndicatorsInner) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IndicatorsInner) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *IndicatorsInner) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IndicatorsInner) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IndicatorsInner) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IndicatorsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *IndicatorsInner) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *IndicatorsInner) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *IndicatorsInner) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *IndicatorsInner) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


