# QueryInput

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

## Methods

### NewQueryInput

`func NewQueryInput(name string, pattern string, diamond DiamondModel, queryType string, ) *QueryInput`

NewQueryInput instantiates a new QueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInputWithDefaults

`func NewQueryInputWithDefaults() *QueryInput`

NewQueryInputWithDefaults instantiates a new QueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryInput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *QueryInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *QueryInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueryInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *QueryInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *QueryInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *QueryInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *QueryInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *QueryInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *QueryInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *QueryInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *QueryInput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetValidFrom

`func (o *QueryInput) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *QueryInput) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *QueryInput) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *QueryInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *QueryInput) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *QueryInput) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *QueryInput) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *QueryInput) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPattern

`func (o *QueryInput) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *QueryInput) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *QueryInput) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetLocation

`func (o *QueryInput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *QueryInput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *QueryInput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *QueryInput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDiamond

`func (o *QueryInput) GetDiamond() DiamondModel`

GetDiamond returns the Diamond field if non-nil, zero value otherwise.

### GetDiamondOk

`func (o *QueryInput) GetDiamondOk() (*DiamondModel, bool)`

GetDiamondOk returns a tuple with the Diamond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiamond

`func (o *QueryInput) SetDiamond(v DiamondModel)`

SetDiamond sets Diamond field to given value.


### GetKillChainPhases

`func (o *QueryInput) GetKillChainPhases() []string`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *QueryInput) GetKillChainPhasesOk() (*[]string, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *QueryInput) SetKillChainPhases(v []string)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *QueryInput) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.

### GetRelevantTags

`func (o *QueryInput) GetRelevantTags() []string`

GetRelevantTags returns the RelevantTags field if non-nil, zero value otherwise.

### GetRelevantTagsOk

`func (o *QueryInput) GetRelevantTagsOk() (*[]string, bool)`

GetRelevantTagsOk returns a tuple with the RelevantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTags

`func (o *QueryInput) SetRelevantTags(v []string)`

SetRelevantTags sets RelevantTags field to given value.

### HasRelevantTags

`func (o *QueryInput) HasRelevantTags() bool`

HasRelevantTags returns a boolean if a field has been set.

### GetQueryType

`func (o *QueryInput) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *QueryInput) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *QueryInput) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetTargetSystems

`func (o *QueryInput) GetTargetSystems() []string`

GetTargetSystems returns the TargetSystems field if non-nil, zero value otherwise.

### GetTargetSystemsOk

`func (o *QueryInput) GetTargetSystemsOk() (*[]string, bool)`

GetTargetSystemsOk returns a tuple with the TargetSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystems

`func (o *QueryInput) SetTargetSystems(v []string)`

SetTargetSystems sets TargetSystems field to given value.

### HasTargetSystems

`func (o *QueryInput) HasTargetSystems() bool`

HasTargetSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


