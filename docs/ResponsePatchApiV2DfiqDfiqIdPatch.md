# ResponsePatchApiV2DfiqDfiqIdPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Uuid** | Pointer to **string** |  | [optional] 
**DfiqId** | Pointer to **string** |  | [optional] 
**DfiqVersion** | **string** |  | 
**DfiqTags** | Pointer to **[]string** |  | [optional] 
**Contributors** | Pointer to **[]string** |  | [optional] 
**DfiqYaml** | **string** |  | 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "scenario"]
**Id** | **string** |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**ParentIds** | **[]string** |  | 
**Approaches** | Pointer to [**[]DFIQApproach**](DFIQApproach.md) |  | [optional] [default to []]

## Methods

### NewResponsePatchApiV2DfiqDfiqIdPatch

`func NewResponsePatchApiV2DfiqDfiqIdPatch(name string, dfiqVersion string, dfiqYaml string, description string, id string, rootType string, parentIds []string, ) *ResponsePatchApiV2DfiqDfiqIdPatch`

NewResponsePatchApiV2DfiqDfiqIdPatch instantiates a new ResponsePatchApiV2DfiqDfiqIdPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePatchApiV2DfiqDfiqIdPatchWithDefaults

`func NewResponsePatchApiV2DfiqDfiqIdPatchWithDefaults() *ResponsePatchApiV2DfiqDfiqIdPatch`

NewResponsePatchApiV2DfiqDfiqIdPatchWithDefaults instantiates a new ResponsePatchApiV2DfiqDfiqIdPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDfiqId

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqId() string`

GetDfiqId returns the DfiqId field if non-nil, zero value otherwise.

### GetDfiqIdOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqIdOk() (*string, bool)`

GetDfiqIdOk returns a tuple with the DfiqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqId

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetDfiqId(v string)`

SetDfiqId sets DfiqId field to given value.

### HasDfiqId

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasDfiqId() bool`

HasDfiqId returns a boolean if a field has been set.

### GetDfiqVersion

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqVersion() string`

GetDfiqVersion returns the DfiqVersion field if non-nil, zero value otherwise.

### GetDfiqVersionOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqVersionOk() (*string, bool)`

GetDfiqVersionOk returns a tuple with the DfiqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqVersion

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetDfiqVersion(v string)`

SetDfiqVersion sets DfiqVersion field to given value.


### GetDfiqTags

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqTags() []string`

GetDfiqTags returns the DfiqTags field if non-nil, zero value otherwise.

### GetDfiqTagsOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqTagsOk() (*[]string, bool)`

GetDfiqTagsOk returns a tuple with the DfiqTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqTags

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetDfiqTags(v []string)`

SetDfiqTags sets DfiqTags field to given value.

### HasDfiqTags

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasDfiqTags() bool`

HasDfiqTags returns a boolean if a field has been set.

### GetContributors

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetDfiqYaml

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetCreated

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetId(v string)`

SetId sets Id field to given value.


### GetRootType

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetParentIds

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.


### GetApproaches

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetApproaches() []DFIQApproach`

GetApproaches returns the Approaches field if non-nil, zero value otherwise.

### GetApproachesOk

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) GetApproachesOk() (*[]DFIQApproach, bool)`

GetApproachesOk returns a tuple with the Approaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproaches

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) SetApproaches(v []DFIQApproach)`

SetApproaches sets Approaches field to given value.

### HasApproaches

`func (o *ResponsePatchApiV2DfiqDfiqIdPatch) HasApproaches() bool`

HasApproaches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


