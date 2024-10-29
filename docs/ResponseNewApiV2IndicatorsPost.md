# ResponseNewApiV2IndicatorsPost

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

### NewResponseNewApiV2IndicatorsPost

`func NewResponseNewApiV2IndicatorsPost(name string, pattern string, diamond DiamondModel, id string, tags map[string]TagRelationshipOutput, rootType string, queryType string, ) *ResponseNewApiV2IndicatorsPost`

NewResponseNewApiV2IndicatorsPost instantiates a new ResponseNewApiV2IndicatorsPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseNewApiV2IndicatorsPostWithDefaults

`func NewResponseNewApiV2IndicatorsPostWithDefaults() *ResponseNewApiV2IndicatorsPost`

NewResponseNewApiV2IndicatorsPostWithDefaults instantiates a new ResponseNewApiV2IndicatorsPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseNewApiV2IndicatorsPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseNewApiV2IndicatorsPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseNewApiV2IndicatorsPost) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResponseNewApiV2IndicatorsPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseNewApiV2IndicatorsPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseNewApiV2IndicatorsPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseNewApiV2IndicatorsPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseNewApiV2IndicatorsPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseNewApiV2IndicatorsPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseNewApiV2IndicatorsPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseNewApiV2IndicatorsPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseNewApiV2IndicatorsPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseNewApiV2IndicatorsPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseNewApiV2IndicatorsPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseNewApiV2IndicatorsPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponseNewApiV2IndicatorsPost) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponseNewApiV2IndicatorsPost) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponseNewApiV2IndicatorsPost) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponseNewApiV2IndicatorsPost) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *ResponseNewApiV2IndicatorsPost) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *ResponseNewApiV2IndicatorsPost) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *ResponseNewApiV2IndicatorsPost) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *ResponseNewApiV2IndicatorsPost) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *ResponseNewApiV2IndicatorsPost) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ResponseNewApiV2IndicatorsPost) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ResponseNewApiV2IndicatorsPost) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *ResponseNewApiV2IndicatorsPost) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *ResponseNewApiV2IndicatorsPost) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ResponseNewApiV2IndicatorsPost) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ResponseNewApiV2IndicatorsPost) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *ResponseNewApiV2IndicatorsPost) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResponseNewApiV2IndicatorsPost) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResponseNewApiV2IndicatorsPost) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ResponseNewApiV2IndicatorsPost) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *ResponseNewApiV2IndicatorsPost) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *ResponseNewApiV2IndicatorsPost) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *ResponseNewApiV2IndicatorsPost) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *ResponseNewApiV2IndicatorsPost) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ResponseNewApiV2IndicatorsPost) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ResponseNewApiV2IndicatorsPost) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ResponseNewApiV2IndicatorsPost) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *ResponseNewApiV2IndicatorsPost) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *ResponseNewApiV2IndicatorsPost) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *ResponseNewApiV2IndicatorsPost) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *ResponseNewApiV2IndicatorsPost) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetId

`func (o *ResponseNewApiV2IndicatorsPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseNewApiV2IndicatorsPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseNewApiV2IndicatorsPost) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseNewApiV2IndicatorsPost) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseNewApiV2IndicatorsPost) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseNewApiV2IndicatorsPost) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseNewApiV2IndicatorsPost) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseNewApiV2IndicatorsPost) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseNewApiV2IndicatorsPost) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetSid

`func (o *ResponseNewApiV2IndicatorsPost) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ResponseNewApiV2IndicatorsPost) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ResponseNewApiV2IndicatorsPost) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ResponseNewApiV2IndicatorsPost) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponseNewApiV2IndicatorsPost) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponseNewApiV2IndicatorsPost) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponseNewApiV2IndicatorsPost) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponseNewApiV2IndicatorsPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *ResponseNewApiV2IndicatorsPost) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ResponseNewApiV2IndicatorsPost) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ResponseNewApiV2IndicatorsPost) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ResponseNewApiV2IndicatorsPost) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetQueryType

`func (o *ResponseNewApiV2IndicatorsPost) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *ResponseNewApiV2IndicatorsPost) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *ResponseNewApiV2IndicatorsPost) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *ResponseNewApiV2IndicatorsPost) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *ResponseNewApiV2IndicatorsPost) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *ResponseNewApiV2IndicatorsPost) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *ResponseNewApiV2IndicatorsPost) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.

### GetSources

`func (o *ResponseNewApiV2IndicatorsPost) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ResponseNewApiV2IndicatorsPost) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ResponseNewApiV2IndicatorsPost) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *ResponseNewApiV2IndicatorsPost) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAliases

`func (o *ResponseNewApiV2IndicatorsPost) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ResponseNewApiV2IndicatorsPost) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ResponseNewApiV2IndicatorsPost) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ResponseNewApiV2IndicatorsPost) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSupportedOs

`func (o *ResponseNewApiV2IndicatorsPost) GetSupportedOs() []string`

GetSupportedOs returns the SupportedOs field if non-nil, zero value otherwise.

### GetSupportedOsOk

`func (o *ResponseNewApiV2IndicatorsPost) GetSupportedOsOk() (*[]string, bool)`

GetSupportedOsOk returns a tuple with the SupportedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOs

`func (o *ResponseNewApiV2IndicatorsPost) SetSupportedOs(v []string)`

SetSupportedOs sets SupportedOs field to given value.

### HasSupportedOs

`func (o *ResponseNewApiV2IndicatorsPost) HasSupportedOs() bool`

HasSupportedOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


