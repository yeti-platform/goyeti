# NoteOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "note"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | [readonly] 
**Tags** | [**map[string]TagRelationshipOutput**](TagRelationshipOutput.md) |  | [readonly] 
**RootType** | **string** |  | [readonly] 
**RelatedObservablesCount** | **int32** |  | [readonly] 

## Methods

### NewNoteOutput

`func NewNoteOutput(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32, ) *NoteOutput`

NewNoteOutput instantiates a new NoteOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteOutputWithDefaults

`func NewNoteOutputWithDefaults() *NoteOutput`

NewNoteOutputWithDefaults instantiates a new NoteOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NoteOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NoteOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NoteOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteOutput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NoteOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoteOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoteOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NoteOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *NoteOutput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NoteOutput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NoteOutput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *NoteOutput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *NoteOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NoteOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NoteOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NoteOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *NoteOutput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NoteOutput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NoteOutput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NoteOutput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetId

`func (o *NoteOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteOutput) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *NoteOutput) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NoteOutput) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NoteOutput) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *NoteOutput) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *NoteOutput) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *NoteOutput) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetRelatedObservablesCount

`func (o *NoteOutput) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *NoteOutput) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *NoteOutput) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


