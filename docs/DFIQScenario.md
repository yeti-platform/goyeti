# DFIQScenario

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
**Description** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "scenario"]
**Id** | **string** |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewDFIQScenario

`func NewDFIQScenario(name string, dfiqVersion string, dfiqYaml string, description string, id string, rootType string, ) *DFIQScenario`

NewDFIQScenario instantiates a new DFIQScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQScenarioWithDefaults

`func NewDFIQScenarioWithDefaults() *DFIQScenario`

NewDFIQScenarioWithDefaults instantiates a new DFIQScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DFIQScenario) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DFIQScenario) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DFIQScenario) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *DFIQScenario) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DFIQScenario) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DFIQScenario) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DFIQScenario) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *DFIQScenario) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *DFIQScenario) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetDfiqId

`func (o *DFIQScenario) GetDfiqId() string`

GetDfiqId returns the DfiqId field if non-nil, zero value otherwise.

### GetDfiqIdOk

`func (o *DFIQScenario) GetDfiqIdOk() (*string, bool)`

GetDfiqIdOk returns a tuple with the DfiqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqId

`func (o *DFIQScenario) SetDfiqId(v string)`

SetDfiqId sets DfiqId field to given value.

### HasDfiqId

`func (o *DFIQScenario) HasDfiqId() bool`

HasDfiqId returns a boolean if a field has been set.

### SetDfiqIdNil

`func (o *DFIQScenario) SetDfiqIdNil(b bool)`

 SetDfiqIdNil sets the value for DfiqId to be an explicit nil

### UnsetDfiqId
`func (o *DFIQScenario) UnsetDfiqId()`

UnsetDfiqId ensures that no value is present for DfiqId, not even an explicit nil
### GetDfiqVersion

`func (o *DFIQScenario) GetDfiqVersion() string`

GetDfiqVersion returns the DfiqVersion field if non-nil, zero value otherwise.

### GetDfiqVersionOk

`func (o *DFIQScenario) GetDfiqVersionOk() (*string, bool)`

GetDfiqVersionOk returns a tuple with the DfiqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqVersion

`func (o *DFIQScenario) SetDfiqVersion(v string)`

SetDfiqVersion sets DfiqVersion field to given value.


### GetDfiqTags

`func (o *DFIQScenario) GetDfiqTags() []string`

GetDfiqTags returns the DfiqTags field if non-nil, zero value otherwise.

### GetDfiqTagsOk

`func (o *DFIQScenario) GetDfiqTagsOk() (*[]string, bool)`

GetDfiqTagsOk returns a tuple with the DfiqTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqTags

`func (o *DFIQScenario) SetDfiqTags(v []string)`

SetDfiqTags sets DfiqTags field to given value.

### HasDfiqTags

`func (o *DFIQScenario) HasDfiqTags() bool`

HasDfiqTags returns a boolean if a field has been set.

### SetDfiqTagsNil

`func (o *DFIQScenario) SetDfiqTagsNil(b bool)`

 SetDfiqTagsNil sets the value for DfiqTags to be an explicit nil

### UnsetDfiqTags
`func (o *DFIQScenario) UnsetDfiqTags()`

UnsetDfiqTags ensures that no value is present for DfiqTags, not even an explicit nil
### GetContributors

`func (o *DFIQScenario) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DFIQScenario) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DFIQScenario) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *DFIQScenario) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### SetContributorsNil

`func (o *DFIQScenario) SetContributorsNil(b bool)`

 SetContributorsNil sets the value for Contributors to be an explicit nil

### UnsetContributors
`func (o *DFIQScenario) UnsetContributors()`

UnsetContributors ensures that no value is present for Contributors, not even an explicit nil
### GetDfiqYaml

`func (o *DFIQScenario) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *DFIQScenario) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *DFIQScenario) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetCreated

`func (o *DFIQScenario) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DFIQScenario) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DFIQScenario) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DFIQScenario) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *DFIQScenario) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DFIQScenario) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DFIQScenario) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DFIQScenario) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *DFIQScenario) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DFIQScenario) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DFIQScenario) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *DFIQScenario) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DFIQScenario) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DFIQScenario) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DFIQScenario) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *DFIQScenario) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DFIQScenario) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DFIQScenario) SetId(v string)`

SetId sets Id field to given value.


### GetRootType

`func (o *DFIQScenario) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *DFIQScenario) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *DFIQScenario) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


