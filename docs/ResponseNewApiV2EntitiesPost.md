# ResponseNewApiV2EntitiesPost

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

### NewResponseNewApiV2EntitiesPost

`func NewResponseNewApiV2EntitiesPost(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32, ) *ResponseNewApiV2EntitiesPost`

NewResponseNewApiV2EntitiesPost instantiates a new ResponseNewApiV2EntitiesPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseNewApiV2EntitiesPostWithDefaults

`func NewResponseNewApiV2EntitiesPostWithDefaults() *ResponseNewApiV2EntitiesPost`

NewResponseNewApiV2EntitiesPostWithDefaults instantiates a new ResponseNewApiV2EntitiesPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseNewApiV2EntitiesPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseNewApiV2EntitiesPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseNewApiV2EntitiesPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseNewApiV2EntitiesPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponseNewApiV2EntitiesPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseNewApiV2EntitiesPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseNewApiV2EntitiesPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ResponseNewApiV2EntitiesPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseNewApiV2EntitiesPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseNewApiV2EntitiesPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseNewApiV2EntitiesPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *ResponseNewApiV2EntitiesPost) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponseNewApiV2EntitiesPost) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponseNewApiV2EntitiesPost) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponseNewApiV2EntitiesPost) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseNewApiV2EntitiesPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseNewApiV2EntitiesPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseNewApiV2EntitiesPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseNewApiV2EntitiesPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponseNewApiV2EntitiesPost) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponseNewApiV2EntitiesPost) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponseNewApiV2EntitiesPost) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponseNewApiV2EntitiesPost) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAliases

`func (o *ResponseNewApiV2EntitiesPost) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ResponseNewApiV2EntitiesPost) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ResponseNewApiV2EntitiesPost) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ResponseNewApiV2EntitiesPost) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *ResponseNewApiV2EntitiesPost) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ResponseNewApiV2EntitiesPost) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ResponseNewApiV2EntitiesPost) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ResponseNewApiV2EntitiesPost) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetToolVersion

`func (o *ResponseNewApiV2EntitiesPost) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *ResponseNewApiV2EntitiesPost) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *ResponseNewApiV2EntitiesPost) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *ResponseNewApiV2EntitiesPost) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetId

`func (o *ResponseNewApiV2EntitiesPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseNewApiV2EntitiesPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseNewApiV2EntitiesPost) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseNewApiV2EntitiesPost) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseNewApiV2EntitiesPost) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseNewApiV2EntitiesPost) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseNewApiV2EntitiesPost) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseNewApiV2EntitiesPost) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseNewApiV2EntitiesPost) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetRelatedObservablesCount

`func (o *ResponseNewApiV2EntitiesPost) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *ResponseNewApiV2EntitiesPost) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *ResponseNewApiV2EntitiesPost) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.


### GetTitle

`func (o *ResponseNewApiV2EntitiesPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResponseNewApiV2EntitiesPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResponseNewApiV2EntitiesPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResponseNewApiV2EntitiesPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBaseScore

`func (o *ResponseNewApiV2EntitiesPost) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *ResponseNewApiV2EntitiesPost) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *ResponseNewApiV2EntitiesPost) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *ResponseNewApiV2EntitiesPost) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetSeverity

`func (o *ResponseNewApiV2EntitiesPost) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ResponseNewApiV2EntitiesPost) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ResponseNewApiV2EntitiesPost) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ResponseNewApiV2EntitiesPost) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetReference

`func (o *ResponseNewApiV2EntitiesPost) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ResponseNewApiV2EntitiesPost) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ResponseNewApiV2EntitiesPost) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ResponseNewApiV2EntitiesPost) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetFamily

`func (o *ResponseNewApiV2EntitiesPost) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ResponseNewApiV2EntitiesPost) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ResponseNewApiV2EntitiesPost) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ResponseNewApiV2EntitiesPost) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponseNewApiV2EntitiesPost) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponseNewApiV2EntitiesPost) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponseNewApiV2EntitiesPost) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponseNewApiV2EntitiesPost) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseNewApiV2EntitiesPost) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseNewApiV2EntitiesPost) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseNewApiV2EntitiesPost) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseNewApiV2EntitiesPost) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetThreatActorTypes

`func (o *ResponseNewApiV2EntitiesPost) GetThreatActorTypes() []string`

GetThreatActorTypes returns the ThreatActorTypes field if non-nil, zero value otherwise.

### GetThreatActorTypesOk

`func (o *ResponseNewApiV2EntitiesPost) GetThreatActorTypesOk() (*[]string, bool)`

GetThreatActorTypesOk returns a tuple with the ThreatActorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorTypes

`func (o *ResponseNewApiV2EntitiesPost) SetThreatActorTypes(v []string)`

SetThreatActorTypes sets ThreatActorTypes field to given value.

### HasThreatActorTypes

`func (o *ResponseNewApiV2EntitiesPost) HasThreatActorTypes() bool`

HasThreatActorTypes returns a boolean if a field has been set.

### GetIdentityClass

`func (o *ResponseNewApiV2EntitiesPost) GetIdentityClass() string`

GetIdentityClass returns the IdentityClass field if non-nil, zero value otherwise.

### GetIdentityClassOk

`func (o *ResponseNewApiV2EntitiesPost) GetIdentityClassOk() (*string, bool)`

GetIdentityClassOk returns a tuple with the IdentityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClass

`func (o *ResponseNewApiV2EntitiesPost) SetIdentityClass(v string)`

SetIdentityClass sets IdentityClass field to given value.

### HasIdentityClass

`func (o *ResponseNewApiV2EntitiesPost) HasIdentityClass() bool`

HasIdentityClass returns a boolean if a field has been set.

### GetSectors

`func (o *ResponseNewApiV2EntitiesPost) GetSectors() []string`

GetSectors returns the Sectors field if non-nil, zero value otherwise.

### GetSectorsOk

`func (o *ResponseNewApiV2EntitiesPost) GetSectorsOk() (*[]string, bool)`

GetSectorsOk returns a tuple with the Sectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectors

`func (o *ResponseNewApiV2EntitiesPost) SetSectors(v []string)`

SetSectors sets Sectors field to given value.

### HasSectors

`func (o *ResponseNewApiV2EntitiesPost) HasSectors() bool`

HasSectors returns a boolean if a field has been set.

### GetContactInformation

`func (o *ResponseNewApiV2EntitiesPost) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *ResponseNewApiV2EntitiesPost) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *ResponseNewApiV2EntitiesPost) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *ResponseNewApiV2EntitiesPost) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


