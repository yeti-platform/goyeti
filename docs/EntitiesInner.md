# EntitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "investigation"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] [default to ""]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**RelatedObservablesCount** | **int32** |  | [readonly] 
**IdentityClass** | Pointer to **string** |  | [optional] [default to ""]
**Sectors** | Pointer to **[]string** |  | [optional] [default to []]
**ContactInformation** | Pointer to **string** |  | [optional] [default to ""]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]
**ToolVersion** | Pointer to **string** |  | [optional] [default to ""]
**ThreatActorTypes** | Pointer to **[]string** |  | [optional] [default to []]
**Family** | Pointer to **string** |  | [optional] [default to ""]
**Title** | Pointer to **string** |  | [optional] [default to ""]
**BaseScore** | Pointer to **float32** |  | [optional] [default to 0.0]
**Severity** | Pointer to [**SeverityType**](SeverityType.md) |  | [optional] 

## Methods

### NewEntitiesInner

`func NewEntitiesInner(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32, ) *EntitiesInner`

NewEntitiesInner instantiates a new EntitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitiesInnerWithDefaults

`func NewEntitiesInnerWithDefaults() *EntitiesInner`

NewEntitiesInnerWithDefaults instantiates a new EntitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntitiesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitiesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitiesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntitiesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EntitiesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitiesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitiesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntitiesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitiesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitiesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitiesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *EntitiesInner) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EntitiesInner) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EntitiesInner) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *EntitiesInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *EntitiesInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitiesInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitiesInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntitiesInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *EntitiesInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *EntitiesInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *EntitiesInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *EntitiesInner) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetReference

`func (o *EntitiesInner) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *EntitiesInner) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *EntitiesInner) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *EntitiesInner) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetId

`func (o *EntitiesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitiesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitiesInner) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *EntitiesInner) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntitiesInner) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntitiesInner) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *EntitiesInner) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *EntitiesInner) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *EntitiesInner) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetRelatedObservablesCount

`func (o *EntitiesInner) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *EntitiesInner) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *EntitiesInner) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.


### GetIdentityClass

`func (o *EntitiesInner) GetIdentityClass() string`

GetIdentityClass returns the IdentityClass field if non-nil, zero value otherwise.

### GetIdentityClassOk

`func (o *EntitiesInner) GetIdentityClassOk() (*string, bool)`

GetIdentityClassOk returns a tuple with the IdentityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClass

`func (o *EntitiesInner) SetIdentityClass(v string)`

SetIdentityClass sets IdentityClass field to given value.

### HasIdentityClass

`func (o *EntitiesInner) HasIdentityClass() bool`

HasIdentityClass returns a boolean if a field has been set.

### GetSectors

`func (o *EntitiesInner) GetSectors() []string`

GetSectors returns the Sectors field if non-nil, zero value otherwise.

### GetSectorsOk

`func (o *EntitiesInner) GetSectorsOk() (*[]string, bool)`

GetSectorsOk returns a tuple with the Sectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectors

`func (o *EntitiesInner) SetSectors(v []string)`

SetSectors sets Sectors field to given value.

### HasSectors

`func (o *EntitiesInner) HasSectors() bool`

HasSectors returns a boolean if a field has been set.

### GetContactInformation

`func (o *EntitiesInner) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *EntitiesInner) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *EntitiesInner) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *EntitiesInner) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.

### GetAliases

`func (o *EntitiesInner) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *EntitiesInner) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *EntitiesInner) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *EntitiesInner) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFirstSeen

`func (o *EntitiesInner) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *EntitiesInner) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *EntitiesInner) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *EntitiesInner) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *EntitiesInner) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *EntitiesInner) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *EntitiesInner) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *EntitiesInner) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *EntitiesInner) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *EntitiesInner) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *EntitiesInner) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *EntitiesInner) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetToolVersion

`func (o *EntitiesInner) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *EntitiesInner) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *EntitiesInner) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *EntitiesInner) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetThreatActorTypes

`func (o *EntitiesInner) GetThreatActorTypes() []string`

GetThreatActorTypes returns the ThreatActorTypes field if non-nil, zero value otherwise.

### GetThreatActorTypesOk

`func (o *EntitiesInner) GetThreatActorTypesOk() (*[]string, bool)`

GetThreatActorTypesOk returns a tuple with the ThreatActorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorTypes

`func (o *EntitiesInner) SetThreatActorTypes(v []string)`

SetThreatActorTypes sets ThreatActorTypes field to given value.

### HasThreatActorTypes

`func (o *EntitiesInner) HasThreatActorTypes() bool`

HasThreatActorTypes returns a boolean if a field has been set.

### GetFamily

`func (o *EntitiesInner) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *EntitiesInner) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *EntitiesInner) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *EntitiesInner) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetTitle

`func (o *EntitiesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EntitiesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EntitiesInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EntitiesInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBaseScore

`func (o *EntitiesInner) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *EntitiesInner) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *EntitiesInner) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *EntitiesInner) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetSeverity

`func (o *EntitiesInner) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EntitiesInner) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EntitiesInner) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EntitiesInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


