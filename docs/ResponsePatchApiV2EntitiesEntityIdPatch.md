# ResponsePatchApiV2EntitiesEntityIdPatch

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
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**RelatedObservablesCount** | **int32** |  | [readonly] 
**Title** | Pointer to **string** |  | [optional] [default to ""]
**BaseScore** | Pointer to **float32** |  | [optional] [default to 0.0]
**Severity** | Pointer to [**SeverityType**](SeverityType.md) |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] [default to ""]
**Family** | Pointer to **string** |  | [optional] [default to ""]
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**ThreatActorTypes** | Pointer to **[]string** |  | [optional] [default to []]
**IdentityClass** | Pointer to **string** |  | [optional] [default to ""]
**Sectors** | Pointer to **[]string** |  | [optional] [default to []]
**ContactInformation** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewResponsePatchApiV2EntitiesEntityIdPatch

`func NewResponsePatchApiV2EntitiesEntityIdPatch(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32, ) *ResponsePatchApiV2EntitiesEntityIdPatch`

NewResponsePatchApiV2EntitiesEntityIdPatch instantiates a new ResponsePatchApiV2EntitiesEntityIdPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePatchApiV2EntitiesEntityIdPatchWithDefaults

`func NewResponsePatchApiV2EntitiesEntityIdPatchWithDefaults() *ResponsePatchApiV2EntitiesEntityIdPatch`

NewResponsePatchApiV2EntitiesEntityIdPatchWithDefaults instantiates a new ResponsePatchApiV2EntitiesEntityIdPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAliases

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetToolVersion

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetId

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetRelatedObservablesCount

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.


### GetTitle

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBaseScore

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetSeverity

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetReference

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetFamily

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetThreatActorTypes

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetThreatActorTypes() []string`

GetThreatActorTypes returns the ThreatActorTypes field if non-nil, zero value otherwise.

### GetThreatActorTypesOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetThreatActorTypesOk() (*[]string, bool)`

GetThreatActorTypesOk returns a tuple with the ThreatActorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorTypes

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetThreatActorTypes(v []string)`

SetThreatActorTypes sets ThreatActorTypes field to given value.

### HasThreatActorTypes

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasThreatActorTypes() bool`

HasThreatActorTypes returns a boolean if a field has been set.

### GetIdentityClass

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetIdentityClass() string`

GetIdentityClass returns the IdentityClass field if non-nil, zero value otherwise.

### GetIdentityClassOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetIdentityClassOk() (*string, bool)`

GetIdentityClassOk returns a tuple with the IdentityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClass

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetIdentityClass(v string)`

SetIdentityClass sets IdentityClass field to given value.

### HasIdentityClass

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasIdentityClass() bool`

HasIdentityClass returns a boolean if a field has been set.

### GetSectors

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetSectors() []string`

GetSectors returns the Sectors field if non-nil, zero value otherwise.

### GetSectorsOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetSectorsOk() (*[]string, bool)`

GetSectorsOk returns a tuple with the Sectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectors

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetSectors(v []string)`

SetSectors sets Sectors field to given value.

### HasSectors

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasSectors() bool`

HasSectors returns a boolean if a field has been set.

### GetContactInformation

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *ResponsePatchApiV2EntitiesEntityIdPatch) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


