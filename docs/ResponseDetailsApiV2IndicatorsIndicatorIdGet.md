# ResponseDetailsApiV2IndicatorsIndicatorIdGet

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

### NewResponseDetailsApiV2IndicatorsIndicatorIdGet

`func NewResponseDetailsApiV2IndicatorsIndicatorIdGet(name string, pattern string, diamond DiamondModel, queryType string, id string, tags map[string]TagRelationshipOutput, rootType string, ) *ResponseDetailsApiV2IndicatorsIndicatorIdGet`

NewResponseDetailsApiV2IndicatorsIndicatorIdGet instantiates a new ResponseDetailsApiV2IndicatorsIndicatorIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsApiV2IndicatorsIndicatorIdGetWithDefaults

`func NewResponseDetailsApiV2IndicatorsIndicatorIdGetWithDefaults() *ResponseDetailsApiV2IndicatorsIndicatorIdGet`

NewResponseDetailsApiV2IndicatorsIndicatorIdGetWithDefaults instantiates a new ResponseDetailsApiV2IndicatorsIndicatorIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetQueryType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.

### GetId

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetSources

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAliases

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSupportedOs

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetSupportedOs() []string`

GetSupportedOs returns the SupportedOs field if non-nil, zero value otherwise.

### GetSupportedOsOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetSupportedOsOk() (*[]string, bool)`

GetSupportedOsOk returns a tuple with the SupportedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOs

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetSupportedOs(v []string)`

SetSupportedOs sets SupportedOs field to given value.

### HasSupportedOs

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasSupportedOs() bool`

HasSupportedOs returns a boolean if a field has been set.

### GetSid

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ResponseDetailsApiV2IndicatorsIndicatorIdGet) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


