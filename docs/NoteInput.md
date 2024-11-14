# NoteInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "note"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNoteInput

`func NewNoteInput(name string, ) *NoteInput`

NewNoteInput instantiates a new NoteInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteInputWithDefaults

`func NewNoteInputWithDefaults() *NoteInput`

NewNoteInputWithDefaults instantiates a new NoteInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NoteInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NoteInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NoteInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NoteInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoteInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoteInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NoteInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *NoteInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NoteInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NoteInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *NoteInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *NoteInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NoteInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NoteInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NoteInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *NoteInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NoteInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NoteInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NoteInput) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


