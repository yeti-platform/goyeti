# ResponseDetailsApiV2DfiqDfiqIdGet

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

### NewResponseDetailsApiV2DfiqDfiqIdGet

`func NewResponseDetailsApiV2DfiqDfiqIdGet(name string, dfiqVersion string, dfiqYaml string, description string, id string, rootType string, parentIds []string, ) *ResponseDetailsApiV2DfiqDfiqIdGet`

NewResponseDetailsApiV2DfiqDfiqIdGet instantiates a new ResponseDetailsApiV2DfiqDfiqIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsApiV2DfiqDfiqIdGetWithDefaults

`func NewResponseDetailsApiV2DfiqDfiqIdGetWithDefaults() *ResponseDetailsApiV2DfiqDfiqIdGet`

NewResponseDetailsApiV2DfiqDfiqIdGetWithDefaults instantiates a new ResponseDetailsApiV2DfiqDfiqIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDfiqId

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqId() string`

GetDfiqId returns the DfiqId field if non-nil, zero value otherwise.

### GetDfiqIdOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqIdOk() (*string, bool)`

GetDfiqIdOk returns a tuple with the DfiqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqId

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetDfiqId(v string)`

SetDfiqId sets DfiqId field to given value.

### HasDfiqId

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasDfiqId() bool`

HasDfiqId returns a boolean if a field has been set.

### GetDfiqVersion

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqVersion() string`

GetDfiqVersion returns the DfiqVersion field if non-nil, zero value otherwise.

### GetDfiqVersionOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqVersionOk() (*string, bool)`

GetDfiqVersionOk returns a tuple with the DfiqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqVersion

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetDfiqVersion(v string)`

SetDfiqVersion sets DfiqVersion field to given value.


### GetDfiqTags

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqTags() []string`

GetDfiqTags returns the DfiqTags field if non-nil, zero value otherwise.

### GetDfiqTagsOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqTagsOk() (*[]string, bool)`

GetDfiqTagsOk returns a tuple with the DfiqTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqTags

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetDfiqTags(v []string)`

SetDfiqTags sets DfiqTags field to given value.

### HasDfiqTags

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasDfiqTags() bool`

HasDfiqTags returns a boolean if a field has been set.

### GetContributors

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetDfiqYaml

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetCreated

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetId(v string)`

SetId sets Id field to given value.


### GetRootType

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetParentIds

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.


### GetApproaches

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetApproaches() []DFIQApproach`

GetApproaches returns the Approaches field if non-nil, zero value otherwise.

### GetApproachesOk

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) GetApproachesOk() (*[]DFIQApproach, bool)`

GetApproachesOk returns a tuple with the Approaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproaches

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) SetApproaches(v []DFIQApproach)`

SetApproaches sets Approaches field to given value.

### HasApproaches

`func (o *ResponseDetailsApiV2DfiqDfiqIdGet) HasApproaches() bool`

HasApproaches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


