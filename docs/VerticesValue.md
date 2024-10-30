# VerticesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | **string** |  | 
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Modified** | Pointer to **time.Time** |  | [optional] 
**IdentityClass** | Pointer to **string** |  | [optional] [default to ""]
**Sectors** | Pointer to **[]string** |  | [optional] [default to []]
**ContactInformation** | Pointer to **string** |  | [optional] [default to ""]
**RelatedObservablesCount** | **int32** |  | [readonly] 
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**Family** | Pointer to **string** |  | [optional] [default to ""]
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**ThreatActorTypes** | Pointer to **[]string** |  | [optional] [default to []]
**Title** | Pointer to **string** |  | [optional] [default to ""]
**BaseScore** | Pointer to **float32** |  | [optional] [default to 0.0]
**Severity** | Pointer to [**SeverityType**](SeverityType.md) |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] [default to ""]
**ToolVersion** | Pointer to **string** |  | [optional] [default to ""]
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidUntil** | Pointer to **time.Time** |  | [optional] 
**Pattern** | **string** |  | 
**Location** | Pointer to **string** |  | [optional] [default to ""]
**Diamond** | [**DiamondModel**](DiamondModel.md) |  | 
**RelevantTags** | Pointer to **[]string** |  | [optional] [default to []]
**Sources** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**SupportedOs** | Pointer to **[]string** |  | [optional] [default to []]
**QueryType** | **string** |  | 
**TargetSystems** | Pointer to **[]string** |  | [optional] [default to []]
**Sid** | Pointer to **int32** |  | [optional] [default to 0]
**Metadata** | Pointer to **[]string** |  | [optional] [default to []]
**References** | Pointer to **[]string** |  | [optional] [default to []]
**Count** | Pointer to **int32** |  | [optional] [default to 0]
**DefaultExpiration** | Pointer to **string** |  | [optional] [default to "P90D"]
**Produces** | Pointer to **[]string** |  | [optional] [default to []]
**Replaces** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewVerticesValue

`func NewVerticesValue(value string, type_ string, id string, tags map[string]TagRelationshipOutput, rootType string, name string, relatedObservablesCount int32, pattern string, diamond DiamondModel, queryType string, ) *VerticesValue`

NewVerticesValue instantiates a new VerticesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticesValueWithDefaults

`func NewVerticesValueWithDefaults() *VerticesValue`

NewVerticesValueWithDefaults instantiates a new VerticesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VerticesValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VerticesValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VerticesValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *VerticesValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VerticesValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VerticesValue) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *VerticesValue) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VerticesValue) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VerticesValue) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VerticesValue) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *VerticesValue) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *VerticesValue) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *VerticesValue) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *VerticesValue) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *VerticesValue) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *VerticesValue) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *VerticesValue) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *VerticesValue) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetId

`func (o *VerticesValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerticesValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerticesValue) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *VerticesValue) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VerticesValue) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VerticesValue) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *VerticesValue) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *VerticesValue) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *VerticesValue) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetName

`func (o *VerticesValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerticesValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerticesValue) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VerticesValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VerticesValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VerticesValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VerticesValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModified

`func (o *VerticesValue) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *VerticesValue) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *VerticesValue) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *VerticesValue) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetIdentityClass

`func (o *VerticesValue) GetIdentityClass() string`

GetIdentityClass returns the IdentityClass field if non-nil, zero value otherwise.

### GetIdentityClassOk

`func (o *VerticesValue) GetIdentityClassOk() (*string, bool)`

GetIdentityClassOk returns a tuple with the IdentityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClass

`func (o *VerticesValue) SetIdentityClass(v string)`

SetIdentityClass sets IdentityClass field to given value.

### HasIdentityClass

`func (o *VerticesValue) HasIdentityClass() bool`

HasIdentityClass returns a boolean if a field has been set.

### GetSectors

`func (o *VerticesValue) GetSectors() []string`

GetSectors returns the Sectors field if non-nil, zero value otherwise.

### GetSectorsOk

`func (o *VerticesValue) GetSectorsOk() (*[]string, bool)`

GetSectorsOk returns a tuple with the Sectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectors

`func (o *VerticesValue) SetSectors(v []string)`

SetSectors sets Sectors field to given value.

### HasSectors

`func (o *VerticesValue) HasSectors() bool`

HasSectors returns a boolean if a field has been set.

### GetContactInformation

`func (o *VerticesValue) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *VerticesValue) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *VerticesValue) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *VerticesValue) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.

### GetRelatedObservablesCount

`func (o *VerticesValue) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *VerticesValue) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *VerticesValue) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.


### GetKillChainPhases

`func (o *VerticesValue) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *VerticesValue) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *VerticesValue) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *VerticesValue) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetAliases

`func (o *VerticesValue) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *VerticesValue) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *VerticesValue) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *VerticesValue) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFamily

`func (o *VerticesValue) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *VerticesValue) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *VerticesValue) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *VerticesValue) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetFirstSeen

`func (o *VerticesValue) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *VerticesValue) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *VerticesValue) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *VerticesValue) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *VerticesValue) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *VerticesValue) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *VerticesValue) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *VerticesValue) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetThreatActorTypes

`func (o *VerticesValue) GetThreatActorTypes() []string`

GetThreatActorTypes returns the ThreatActorTypes field if non-nil, zero value otherwise.

### GetThreatActorTypesOk

`func (o *VerticesValue) GetThreatActorTypesOk() (*[]string, bool)`

GetThreatActorTypesOk returns a tuple with the ThreatActorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorTypes

`func (o *VerticesValue) SetThreatActorTypes(v []string)`

SetThreatActorTypes sets ThreatActorTypes field to given value.

### HasThreatActorTypes

`func (o *VerticesValue) HasThreatActorTypes() bool`

HasThreatActorTypes returns a boolean if a field has been set.

### GetTitle

`func (o *VerticesValue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VerticesValue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VerticesValue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VerticesValue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBaseScore

`func (o *VerticesValue) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *VerticesValue) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *VerticesValue) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *VerticesValue) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetSeverity

`func (o *VerticesValue) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *VerticesValue) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *VerticesValue) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *VerticesValue) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetReference

`func (o *VerticesValue) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *VerticesValue) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *VerticesValue) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *VerticesValue) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetToolVersion

`func (o *VerticesValue) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *VerticesValue) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *VerticesValue) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *VerticesValue) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetValidFrom

`func (o *VerticesValue) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *VerticesValue) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *VerticesValue) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *VerticesValue) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *VerticesValue) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *VerticesValue) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *VerticesValue) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *VerticesValue) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *VerticesValue) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *VerticesValue) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *VerticesValue) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *VerticesValue) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VerticesValue) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VerticesValue) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VerticesValue) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *VerticesValue) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *VerticesValue) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *VerticesValue) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetRelevantTags

`func (o *VerticesValue) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *VerticesValue) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *VerticesValue) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *VerticesValue) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetSources

`func (o *VerticesValue) GetSources() []map[string]interface{}`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *VerticesValue) GetSourcesOk() (*[]map[string]interface{}, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *VerticesValue) SetSources(v []map[string]interface{})`

SetSources sets Sources field to given value.

### HasSources

`func (o *VerticesValue) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSupportedOs

`func (o *VerticesValue) GetSupportedOs() []string`

GetSupportedOs returns the SupportedOs field if non-nil, zero value otherwise.

### GetSupportedOsOk

`func (o *VerticesValue) GetSupportedOsOk() (*[]string, bool)`

GetSupportedOsOk returns a tuple with the SupportedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOs

`func (o *VerticesValue) SetSupportedOs(v []string)`

SetSupportedOs sets SupportedOs field to given value.

### HasSupportedOs

`func (o *VerticesValue) HasSupportedOs() bool`

HasSupportedOs returns a boolean if a field has been set.

### GetQueryType

`func (o *VerticesValue) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *VerticesValue) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *VerticesValue) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *VerticesValue) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *VerticesValue) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *VerticesValue) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *VerticesValue) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.

### GetSid

`func (o *VerticesValue) GetSid() int32`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerticesValue) GetSidOk() (*int32, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerticesValue) SetSid(v int32)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerticesValue) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetMetadata

`func (o *VerticesValue) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VerticesValue) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VerticesValue) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VerticesValue) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferences

`func (o *VerticesValue) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *VerticesValue) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *VerticesValue) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *VerticesValue) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetCount

`func (o *VerticesValue) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VerticesValue) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VerticesValue) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VerticesValue) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDefaultExpiration

`func (o *VerticesValue) GetDefaultExpiration() string`

GetDefaultExpiration returns the DefaultExpiration field if non-nil, zero value otherwise.

### GetDefaultExpirationOk

`func (o *VerticesValue) GetDefaultExpirationOk() (*string, bool)`

GetDefaultExpirationOk returns a tuple with the DefaultExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpiration

`func (o *VerticesValue) SetDefaultExpiration(v string)`

SetDefaultExpiration sets DefaultExpiration field to given value.

### HasDefaultExpiration

`func (o *VerticesValue) HasDefaultExpiration() bool`

HasDefaultExpiration returns a boolean if a field has been set.

### GetProduces

`func (o *VerticesValue) GetProduces() []string`

GetProduces returns the Produces field if non-nil, zero value otherwise.

### GetProducesOk

`func (o *VerticesValue) GetProducesOk() (*[]string, bool)`

GetProducesOk returns a tuple with the Produces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduces

`func (o *VerticesValue) SetProduces(v []string)`

SetProduces sets Produces field to given value.

### HasProduces

`func (o *VerticesValue) HasProduces() bool`

HasProduces returns a boolean if a field has been set.

### GetReplaces

`func (o *VerticesValue) GetReplaces() []string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *VerticesValue) GetReplacesOk() (*[]string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *VerticesValue) SetReplaces(v []string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *VerticesValue) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


