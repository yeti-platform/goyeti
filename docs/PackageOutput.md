# PackageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "package"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Version** | Pointer to **string** |  | [optional] 
**RegitryType** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 

## Methods

### NewPackageOutput

`func NewPackageOutput(value string, id string, tags map[string]TagRelationshipOutput, rootType string, ) *PackageOutput`

NewPackageOutput instantiates a new PackageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageOutputWithDefaults

`func NewPackageOutputWithDefaults() *PackageOutput`

NewPackageOutputWithDefaults instantiates a new PackageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PackageOutput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PackageOutput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PackageOutput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *PackageOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PackageOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *PackageOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PackageOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PackageOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PackageOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *PackageOutput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PackageOutput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PackageOutput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *PackageOutput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *PackageOutput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *PackageOutput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *PackageOutput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *PackageOutput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetVersion

`func (o *PackageOutput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageOutput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageOutput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageOutput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegitryType

`func (o *PackageOutput) GetRegitryType() string`

GetRegitryType returns the RegitryType field if non-nil, zero value otherwise.

### GetRegitryTypeOk

`func (o *PackageOutput) GetRegitryTypeOk() (*string, bool)`

GetRegitryTypeOk returns a tuple with the RegitryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegitryType

`func (o *PackageOutput) SetRegitryType(v string)`

SetRegitryType sets RegitryType field to given value.

### HasRegitryType

`func (o *PackageOutput) HasRegitryType() bool`

HasRegitryType returns a boolean if a field has been set.

### GetId

`func (o *PackageOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageOutput) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *PackageOutput) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackageOutput) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackageOutput) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *PackageOutput) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *PackageOutput) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *PackageOutput) SetRootType(v string)`

SetRootType sets RootType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


