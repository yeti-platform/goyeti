# CourseOfActionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "course-of-action"]
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

### NewCourseOfActionOutput

`func NewCourseOfActionOutput(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32, ) *CourseOfActionOutput`

NewCourseOfActionOutput instantiates a new CourseOfActionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseOfActionOutputWithDefaults

`func NewCourseOfActionOutputWithDefaults() *CourseOfActionOutput`

NewCourseOfActionOutputWithDefaults instantiates a new CourseOfActionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CourseOfActionOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CourseOfActionOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CourseOfActionOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CourseOfActionOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CourseOfActionOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CourseOfActionOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CourseOfActionOutput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CourseOfActionOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseOfActionOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseOfActionOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseOfActionOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *CourseOfActionOutput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CourseOfActionOutput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CourseOfActionOutput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *CourseOfActionOutput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *CourseOfActionOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CourseOfActionOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CourseOfActionOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CourseOfActionOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *CourseOfActionOutput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CourseOfActionOutput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CourseOfActionOutput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CourseOfActionOutput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetId

`func (o *CourseOfActionOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseOfActionOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseOfActionOutput) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *CourseOfActionOutput) GetTags() map[string]TagRelationshipOutput`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CourseOfActionOutput) GetTagsOk() (*map[string]TagRelationshipOutput, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CourseOfActionOutput) SetTags(v map[string]TagRelationshipOutput)`

SetTags sets Tags field to given value.


### GetRootType

`func (o *CourseOfActionOutput) GetRootType() string`

GetRootType returns the RootType field if non-nil, zero value otherwise.

### GetRootTypeOk

`func (o *CourseOfActionOutput) GetRootTypeOk() (*string, bool)`

GetRootTypeOk returns a tuple with the RootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootType

`func (o *CourseOfActionOutput) SetRootType(v string)`

SetRootType sets RootType field to given value.


### GetRelatedObservablesCount

`func (o *CourseOfActionOutput) GetRelatedObservablesCount() int32`

GetRelatedObservablesCount returns the RelatedObservablesCount field if non-nil, zero value otherwise.

### GetRelatedObservablesCountOk

`func (o *CourseOfActionOutput) GetRelatedObservablesCountOk() (*int32, bool)`

GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObservablesCount

`func (o *CourseOfActionOutput) SetRelatedObservablesCount(v int32)`

SetRelatedObservablesCount sets RelatedObservablesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


