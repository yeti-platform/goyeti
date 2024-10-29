# DFIQApproach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] [default to []]
**References** | Pointer to **[]string** |  | [optional] [default to []]
**Notes** | Pointer to [**NullableDFIQApproachNotes**](DFIQApproachNotes.md) |  | [optional] 
**Steps** | Pointer to [**[]DFIQApproachStep**](DFIQApproachStep.md) |  | [optional] [default to []]

## Methods

### NewDFIQApproach

`func NewDFIQApproach(name string, description string, ) *DFIQApproach`

NewDFIQApproach instantiates a new DFIQApproach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQApproachWithDefaults

`func NewDFIQApproachWithDefaults() *DFIQApproach`

NewDFIQApproachWithDefaults instantiates a new DFIQApproach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DFIQApproach) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DFIQApproach) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DFIQApproach) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DFIQApproach) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DFIQApproach) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DFIQApproach) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *DFIQApproach) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DFIQApproach) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DFIQApproach) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DFIQApproach) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetReferences

`func (o *DFIQApproach) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *DFIQApproach) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *DFIQApproach) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *DFIQApproach) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetNotes

`func (o *DFIQApproach) GetNotes() DFIQApproachNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DFIQApproach) GetNotesOk() (*DFIQApproachNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DFIQApproach) SetNotes(v DFIQApproachNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DFIQApproach) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *DFIQApproach) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *DFIQApproach) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetSteps

`func (o *DFIQApproach) GetSteps() []DFIQApproachStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DFIQApproach) GetStepsOk() (*[]DFIQApproachStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DFIQApproach) SetSteps(v []DFIQApproachStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *DFIQApproach) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


