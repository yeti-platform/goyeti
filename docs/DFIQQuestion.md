# DFIQQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**DfiqId** | Pointer to **NullableString** |  | [optional] 
**DfiqVersion** | **string** |  | 
**DfiqTags** | Pointer to **[]string** |  | [optional] 
**Contributors** | Pointer to **[]string** |  | [optional] 
**DfiqYaml** | **string** |  | 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Description** | **NullableString** |  | 
**ParentIds** | **[]string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "question"]
**Approaches** | Pointer to [**[]DFIQApproach**](DFIQApproach.md) |  | [optional] [default to []]
**Id** | **string** |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewDFIQQuestion

`func NewDFIQQuestion(name string, dfiqVersion string, dfiqYaml string, description NullableString, parentIds []string, id string, rootType string, ) *DFIQQuestion`

NewDFIQQuestion instantiates a new DFIQQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQQuestionWithDefaults

`func NewDFIQQuestionWithDefaults() *DFIQQuestion`

NewDFIQQuestionWithDefaults instantiates a new DFIQQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DFIQQuestion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DFIQQuestion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DFIQQuestion) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *DFIQQuestion) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DFIQQuestion) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DFIQQuestion) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DFIQQuestion) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *DFIQQuestion) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *DFIQQuestion) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetDfiqId

`func (o *DFIQQuestion) GetDfiqId() string`

GetDfiqId returns the DfiqId field if non-nil, zero value otherwise.

### GetDfiqIdOk

`func (o *DFIQQuestion) GetDfiqIdOk() (*string, bool)`

GetDfiqIdOk returns a tuple with the DfiqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqId

`func (o *DFIQQuestion) SetDfiqId(v string)`

SetDfiqId sets DfiqId field to given value.

### HasDfiqId

`func (o *DFIQQuestion) HasDfiqId() bool`

HasDfiqId returns a boolean if a field has been set.

### SetDfiqIdNil

`func (o *DFIQQuestion) SetDfiqIdNil(b bool)`

 SetDfiqIdNil sets the value for DfiqId to be an explicit nil

### UnsetDfiqId
`func (o *DFIQQuestion) UnsetDfiqId()`

UnsetDfiqId ensures that no value is present for DfiqId, not even an explicit nil
### GetDfiqVersion

`func (o *DFIQQuestion) GetDfiqVersion() string`

GetDfiqVersion returns the DfiqVersion field if non-nil, zero value otherwise.

### GetDfiqVersionOk

`func (o *DFIQQuestion) GetDfiqVersionOk() (*string, bool)`

GetDfiqVersionOk returns a tuple with the DfiqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqVersion

`func (o *DFIQQuestion) SetDfiqVersion(v string)`

SetDfiqVersion sets DfiqVersion field to given value.


### GetDfiqTags

`func (o *DFIQQuestion) GetDfiqTags() []string`

GetDfiqTags returns the DfiqTags field if non-nil, zero value otherwise.

### GetDfiqTagsOk

`func (o *DFIQQuestion) GetDfiqTagsOk() (*[]string, bool)`

GetDfiqTagsOk returns a tuple with the DfiqTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqTags

`func (o *DFIQQuestion) SetDfiqTags(v []string)`

SetDfiqTags sets DfiqTags field to given value.

### HasDfiqTags

`func (o *DFIQQuestion) HasDfiqTags() bool`

HasDfiqTags returns a boolean if a field has been set.

### SetDfiqTagsNil

`func (o *DFIQQuestion) SetDfiqTagsNil(b bool)`

 SetDfiqTagsNil sets the value for DfiqTags to be an explicit nil

### UnsetDfiqTags
`func (o *DFIQQuestion) UnsetDfiqTags()`

UnsetDfiqTags ensures that no value is present for DfiqTags, not even an explicit nil
### GetContributors

`func (o *DFIQQuestion) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DFIQQuestion) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DFIQQuestion) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *DFIQQuestion) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### SetContributorsNil

`func (o *DFIQQuestion) SetContributorsNil(b bool)`

 SetContributorsNil sets the value for Contributors to be an explicit nil

### UnsetContributors
`func (o *DFIQQuestion) UnsetContributors()`

UnsetContributors ensures that no value is present for Contributors, not even an explicit nil
### GetDfiqYaml

`func (o *DFIQQuestion) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *DFIQQuestion) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *DFIQQuestion) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetCreated

`func (o *DFIQQuestion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DFIQQuestion) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DFIQQuestion) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DFIQQuestion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *DFIQQuestion) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DFIQQuestion) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DFIQQuestion) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DFIQQuestion) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *DFIQQuestion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DFIQQuestion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DFIQQuestion) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *DFIQQuestion) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DFIQQuestion) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentIds

`func (o *DFIQQuestion) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *DFIQQuestion) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *DFIQQuestion) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.


### GetType

`func (o *DFIQQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DFIQQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DFIQQuestion) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DFIQQuestion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApproaches

`func (o *DFIQQuestion) GetApproaches() []DFIQApproach`

GetApproaches returns the Approaches field if non-nil, zero value otherwise.

### GetApproachesOk

`func (o *DFIQQuestion) GetApproachesOk() (*[]DFIQApproach, bool)`

GetApproachesOk returns a tuple with the Approaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproaches

`func (o *DFIQQuestion) SetApproaches(v []DFIQApproach)`

SetApproaches sets Approaches field to given value.

### HasApproaches

`func (o *DFIQQuestion) HasApproaches() bool`

HasApproaches returns a boolean if a field has been set.

### GetId

`func (o *DFIQQuestion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DFIQQuestion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DFIQQuestion) SetId(v string)`

SetId sets Id field to given value.


### GetRootType

`func (o *DFIQQuestion) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *DFIQQuestion) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *DFIQQuestion) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


