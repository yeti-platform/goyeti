# ResponseDetailsApiV2EntitiesEntityIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "identity"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**IdentityClass** | Pointer to **string** |  | [optional] [default to ""]
**Sectors** | Pointer to **[]string** |  | [optional] [default to []]
**ContactInformation** | Pointer to **string** |  | [optional] [default to ""]
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**RelatedObservablesCount** | **int32** |  | [readonly] 
**Reference** | Pointer to **string** |  | [optional] [default to ""]
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**KillChainPhases** | Pointer to **[]string** |  | [optional] [default to []]
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**ToolVersion** | Pointer to **string** |  | [optional] [default to ""]
**Family** | Pointer to **string** |  | [optional] [default to ""]
**ThreatActorTypes** | Pointer to **[]string** |  | [optional] [default to []]
**Title** | Pointer to **string** |  | [optional] [default to ""]
**BaseScore** | Pointer to **float32** |  | [optional] [default to 0.0]
**Severity** | Pointer to [**SeverityType**](SeverityType.md) |  | [optional] 

## Methods

### NewResponseDetailsApiV2EntitiesEntityIdGet

`func NewResponseDetailsApiV2EntitiesEntityIdGet(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32, ) *ResponseDetailsApiV2EntitiesEntityIdGet`

NewResponseDetailsApiV2EntitiesEntityIdGet instantiates a new ResponseDetailsApiV2EntitiesEntityIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsApiV2EntitiesEntityIdGetWithDefaults

`func NewResponseDetailsApiV2EntitiesEntityIdGetWithDefaults() *ResponseDetailsApiV2EntitiesEntityIdGet`

NewResponseDetailsApiV2EntitiesEntityIdGetWithDefaults instantiates a new ResponseDetailsApiV2EntitiesEntityIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetIdentityClass

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetIdentityClass() string`

GetIdentityClass returns the IdentityClass field if non-nil, zero value otherwise.

### GetIdentityClassOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetIdentityClassOk() (*string, bool)`

GetIdentityClassOk returns a tuple with the IdentityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClass

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetIdentityClass(v string)`

SetIdentityClass sets IdentityClass field to given value.

### HasIdentityClass

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasIdentityClass() bool`

HasIdentityClass returns a boolean if a field has been set.

### GetSectors

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetSectors() []string`

GetSectors returns the Sectors field if non-nil, zero value otherwise.

### GetSectorsOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetSectorsOk() (*[]string, bool)`

GetSectorsOk returns a tuple with the Sectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectors

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetSectors(v []string)`

SetSectors sets Sectors field to given value.

### HasSectors

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasSectors() bool`

HasSectors returns a boolean if a field has been set.

### GetContactInformation

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.

### GetId

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetRelatedObservablesCount

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.


### GetReference

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetAliases

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetToolVersion

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetFamily

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetThreatActorTypes

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetThreatActorTypes() []string`

GetThreatActorTypes returns the ThreatActorTypes field if non-nil, zero value otherwise.

### GetThreatActorTypesOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetThreatActorTypesOk() (*[]string, bool)`

GetThreatActorTypesOk returns a tuple with the ThreatActorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorTypes

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetThreatActorTypes(v []string)`

SetThreatActorTypes sets ThreatActorTypes field to given value.

### HasThreatActorTypes

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasThreatActorTypes() bool`

HasThreatActorTypes returns a boolean if a field has been set.

### GetTitle

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBaseScore

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetSeverity

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ResponseDetailsApiV2EntitiesEntityIdGet) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


