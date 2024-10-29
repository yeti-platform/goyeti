# ResponsePatchApiV2IndicatorsIndicatorIdPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "sigma"]
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
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**Sid** | Pointer to **int32** |  | [optional] [default to 0]
**Metadata** | Pointer to **[]string** |  | [optional] [default to []]
**References** | Pointer to **[]string** |  | [optional] [default to []]
**QueryType** | **string** |  | 
**TargetSystems** | Pointer to **[]string** |  | [optional] [default to []]
**Sources** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**SupportedOs** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewResponsePatchApiV2IndicatorsIndicatorIdPatch

`func NewResponsePatchApiV2IndicatorsIndicatorIdPatch(name string, pattern string, diamond DiamondModel, id string, tags map[string]TagRelationshipOutput, rootType string, queryType string, ) *ResponsePatchApiV2IndicatorsIndicatorIdPatch`

NewResponsePatchApiV2IndicatorsIndicatorIdPatch instantiates a new ResponsePatchApiV2IndicatorsIndicatorIdPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePatchApiV2IndicatorsIndicatorIdPatchWithDefaults

`func NewResponsePatchApiV2IndicatorsIndicatorIdPatchWithDefaults() *ResponsePatchApiV2IndicatorsIndicatorIdPatch`

NewResponsePatchApiV2IndicatorsIndicatorIdPatchWithDefaults instantiates a new ResponsePatchApiV2IndicatorsIndicatorIdPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetId

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetSid

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetQueryType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.

### GetSources

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAliases

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSupportedOs

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetSupportedOs() []string`

GetSupportedOs returns the SupportedOs field if non-nil, zero value otherwise.

### GetSupportedOsOk

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) GetSupportedOsOk() (*[]string, bool)`

GetSupportedOsOk returns a tuple with the SupportedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOs

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) SetSupportedOs(v []string)`

SetSupportedOs sets SupportedOs field to given value.

### HasSupportedOs

`func (o *ResponsePatchApiV2IndicatorsIndicatorIdPatch) HasSupportedOs() bool`

HasSupportedOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


