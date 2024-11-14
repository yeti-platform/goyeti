# DFIQFacet

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
**Type** | Pointer to **string** |  | [optional] [default to "facet"]
**Id** | **string** |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewDFIQFacet

`func NewDFIQFacet(name string, dfiqVersion string, dfiqYaml string, description NullableString, parentIds []string, id string, rootType string, ) *DFIQFacet`

NewDFIQFacet instantiates a new DFIQFacet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQFacetWithDefaults

`func NewDFIQFacetWithDefaults() *DFIQFacet`

NewDFIQFacetWithDefaults instantiates a new DFIQFacet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DFIQFacet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DFIQFacet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DFIQFacet) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *DFIQFacet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DFIQFacet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DFIQFacet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DFIQFacet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *DFIQFacet) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *DFIQFacet) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetDfiqId

`func (o *DFIQFacet) GetDfiqId() string`

GetDfiqId returns the DfiqId field if non-nil, zero value otherwise.

### GetDfiqIdOk

`func (o *DFIQFacet) GetDfiqIdOk() (*string, bool)`

GetDfiqIdOk returns a tuple with the DfiqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqId

`func (o *DFIQFacet) SetDfiqId(v string)`

SetDfiqId sets DfiqId field to given value.

### HasDfiqId

`func (o *DFIQFacet) HasDfiqId() bool`

HasDfiqId returns a boolean if a field has been set.

### SetDfiqIdNil

`func (o *DFIQFacet) SetDfiqIdNil(b bool)`

 SetDfiqIdNil sets the value for DfiqId to be an explicit nil

### UnsetDfiqId
`func (o *DFIQFacet) UnsetDfiqId()`

UnsetDfiqId ensures that no value is present for DfiqId, not even an explicit nil
### GetDfiqVersion

`func (o *DFIQFacet) GetDfiqVersion() string`

GetDfiqVersion returns the DfiqVersion field if non-nil, zero value otherwise.

### GetDfiqVersionOk

`func (o *DFIQFacet) GetDfiqVersionOk() (*string, bool)`

GetDfiqVersionOk returns a tuple with the DfiqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqVersion

`func (o *DFIQFacet) SetDfiqVersion(v string)`

SetDfiqVersion sets DfiqVersion field to given value.


### GetDfiqTags

`func (o *DFIQFacet) GetDfiqTags() []string`

GetDfiqTags returns the DfiqTags field if non-nil, zero value otherwise.

### GetDfiqTagsOk

`func (o *DFIQFacet) GetDfiqTagsOk() (*[]string, bool)`

GetDfiqTagsOk returns a tuple with the DfiqTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqTags

`func (o *DFIQFacet) SetDfiqTags(v []string)`

SetDfiqTags sets DfiqTags field to given value.

### HasDfiqTags

`func (o *DFIQFacet) HasDfiqTags() bool`

HasDfiqTags returns a boolean if a field has been set.

### SetDfiqTagsNil

`func (o *DFIQFacet) SetDfiqTagsNil(b bool)`

 SetDfiqTagsNil sets the value for DfiqTags to be an explicit nil

### UnsetDfiqTags
`func (o *DFIQFacet) UnsetDfiqTags()`

UnsetDfiqTags ensures that no value is present for DfiqTags, not even an explicit nil
### GetContributors

`func (o *DFIQFacet) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DFIQFacet) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DFIQFacet) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *DFIQFacet) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### SetContributorsNil

`func (o *DFIQFacet) SetContributorsNil(b bool)`

 SetContributorsNil sets the value for Contributors to be an explicit nil

### UnsetContributors
`func (o *DFIQFacet) UnsetContributors()`

UnsetContributors ensures that no value is present for Contributors, not even an explicit nil
### GetDfiqYaml

`func (o *DFIQFacet) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *DFIQFacet) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *DFIQFacet) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetCreated

`func (o *DFIQFacet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DFIQFacet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DFIQFacet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DFIQFacet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *DFIQFacet) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DFIQFacet) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DFIQFacet) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DFIQFacet) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *DFIQFacet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DFIQFacet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DFIQFacet) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *DFIQFacet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DFIQFacet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentIds

`func (o *DFIQFacet) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *DFIQFacet) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *DFIQFacet) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.


### GetType

`func (o *DFIQFacet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DFIQFacet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DFIQFacet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DFIQFacet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *DFIQFacet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DFIQFacet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DFIQFacet) SetId(v string)`

SetId sets Id field to given value.


### GetRootType

`func (o *DFIQFacet) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *DFIQFacet) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *DFIQFacet) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


