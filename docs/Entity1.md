# Entity1

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

### NewEntity1

`func NewEntity1(name string, ) *Entity1`

NewEntity1 instantiates a new Entity1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntity1WithDefaults

`func NewEntity1WithDefaults() *Entity1`

NewEntity1WithDefaults instantiates a new Entity1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Entity1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entity1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entity1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Entity1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Entity1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entity1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entity1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Entity1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Entity1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Entity1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Entity1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *Entity1) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Entity1) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Entity1) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Entity1) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *Entity1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Entity1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Entity1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Entity1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Entity1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Entity1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Entity1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Entity1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetReference

`func (o *Entity1) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Entity1) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Entity1) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Entity1) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetIdentityClass

`func (o *Entity1) GetIdentityClass() string`

GetIdentityClass returns the IdentityClass field if non-nil, zero value otherwise.

### GetIdentityClassOk

`func (o *Entity1) GetIdentityClassOk() (*string, bool)`

GetIdentityClassOk returns a tuple with the IdentityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClass

`func (o *Entity1) SetIdentityClass(v string)`

SetIdentityClass sets IdentityClass field to given value.

### HasIdentityClass

`func (o *Entity1) HasIdentityClass() bool`

HasIdentityClass returns a boolean if a field has been set.

### GetSectors

`func (o *Entity1) GetSectors() []string`

GetSectors returns the Sectors field if non-nil, zero value otherwise.

### GetSectorsOk

`func (o *Entity1) GetSectorsOk() (*[]string, bool)`

GetSectorsOk returns a tuple with the Sectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectors

`func (o *Entity1) SetSectors(v []string)`

SetSectors sets Sectors field to given value.

### HasSectors

`func (o *Entity1) HasSectors() bool`

HasSectors returns a boolean if a field has been set.

### GetContactInformation

`func (o *Entity1) GetContactInformation() string`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *Entity1) GetContactInformationOk() (*string, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *Entity1) SetContactInformation(v string)`

SetContactInformation sets ContactInformation field to given value.

### HasContactInformation

`func (o *Entity1) HasContactInformation() bool`

HasContactInformation returns a boolean if a field has been set.

### GetAliases

`func (o *Entity1) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Entity1) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Entity1) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Entity1) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFirstSeen

`func (o *Entity1) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *Entity1) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *Entity1) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *Entity1) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *Entity1) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Entity1) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Entity1) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Entity1) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *Entity1) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *Entity1) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *Entity1) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *Entity1) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetToolVersion

`func (o *Entity1) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *Entity1) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *Entity1) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *Entity1) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetThreatActorTypes

`func (o *Entity1) GetThreatActorTypes() []string`

GetThreatActorTypes returns the ThreatActorTypes field if non-nil, zero value otherwise.

### GetThreatActorTypesOk

`func (o *Entity1) GetThreatActorTypesOk() (*[]string, bool)`

GetThreatActorTypesOk returns a tuple with the ThreatActorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorTypes

`func (o *Entity1) SetThreatActorTypes(v []string)`

SetThreatActorTypes sets ThreatActorTypes field to given value.

### HasThreatActorTypes

`func (o *Entity1) HasThreatActorTypes() bool`

HasThreatActorTypes returns a boolean if a field has been set.

### GetFamily

`func (o *Entity1) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Entity1) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Entity1) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Entity1) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetTitle

`func (o *Entity1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Entity1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Entity1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Entity1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBaseScore

`func (o *Entity1) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *Entity1) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *Entity1) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *Entity1) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetSeverity

`func (o *Entity1) GetSeverity() SeverityType`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Entity1) GetSeverityOk() (*SeverityType, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Entity1) SetSeverity(v SeverityType)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Entity1) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


