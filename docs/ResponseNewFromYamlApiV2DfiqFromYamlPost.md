# ResponseNewFromYamlApiV2DfiqFromYamlPost

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

### NewResponseNewFromYamlApiV2DfiqFromYamlPost

`func NewResponseNewFromYamlApiV2DfiqFromYamlPost(name string, dfiqVersion string, dfiqYaml string, description string, id string, rootType string, parentIds []string, ) *ResponseNewFromYamlApiV2DfiqFromYamlPost`

NewResponseNewFromYamlApiV2DfiqFromYamlPost instantiates a new ResponseNewFromYamlApiV2DfiqFromYamlPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseNewFromYamlApiV2DfiqFromYamlPostWithDefaults

`func NewResponseNewFromYamlApiV2DfiqFromYamlPostWithDefaults() *ResponseNewFromYamlApiV2DfiqFromYamlPost`

NewResponseNewFromYamlApiV2DfiqFromYamlPostWithDefaults instantiates a new ResponseNewFromYamlApiV2DfiqFromYamlPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDfiqId

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqId() string`

GetDfiqId returns the DfiqId field if non-nil, zero value otherwise.

### GetDfiqIdOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqIdOk() (*string, bool)`

GetDfiqIdOk returns a tuple with the DfiqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqId

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetDfiqId(v string)`

SetDfiqId sets DfiqId field to given value.

### HasDfiqId

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasDfiqId() bool`

HasDfiqId returns a boolean if a field has been set.

### GetDfiqVersion

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqVersion() string`

GetDfiqVersion returns the DfiqVersion field if non-nil, zero value otherwise.

### GetDfiqVersionOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqVersionOk() (*string, bool)`

GetDfiqVersionOk returns a tuple with the DfiqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqVersion

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetDfiqVersion(v string)`

SetDfiqVersion sets DfiqVersion field to given value.


### GetDfiqTags

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqTags() []string`

GetDfiqTags returns the DfiqTags field if non-nil, zero value otherwise.

### GetDfiqTagsOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqTagsOk() (*[]string, bool)`

GetDfiqTagsOk returns a tuple with the DfiqTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqTags

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetDfiqTags(v []string)`

SetDfiqTags sets DfiqTags field to given value.

### HasDfiqTags

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasDfiqTags() bool`

HasDfiqTags returns a boolean if a field has been set.

### GetContributors

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetDfiqYaml

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetCreated

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetId(v string)`

SetId sets Id field to given value.


### GetRootType

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetParentIds

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.


### GetApproaches

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetApproaches() []DFIQApproach`

GetApproaches returns the Approaches field if non-nil, zero value otherwise.

### GetApproachesOk

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) GetApproachesOk() (*[]DFIQApproach, bool)`

GetApproachesOk returns a tuple with the Approaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproaches

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) SetApproaches(v []DFIQApproach)`

SetApproaches sets Approaches field to given value.

### HasApproaches

`func (o *ResponseNewFromYamlApiV2DfiqFromYamlPost) HasApproaches() bool`

HasApproaches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


