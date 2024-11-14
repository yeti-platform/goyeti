# IntrusionSetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "intrusion-set"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] [default to []]
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIntrusionSetInput

`func NewIntrusionSetInput(name string, ) *IntrusionSetInput`

NewIntrusionSetInput instantiates a new IntrusionSetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntrusionSetInputWithDefaults

`func NewIntrusionSetInputWithDefaults() *IntrusionSetInput`

NewIntrusionSetInputWithDefaults instantiates a new IntrusionSetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntrusionSetInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntrusionSetInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntrusionSetInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntrusionSetInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IntrusionSetInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntrusionSetInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntrusionSetInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IntrusionSetInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntrusionSetInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntrusionSetInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntrusionSetInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *IntrusionSetInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *IntrusionSetInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *IntrusionSetInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *IntrusionSetInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *IntrusionSetInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IntrusionSetInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IntrusionSetInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IntrusionSetInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *IntrusionSetInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IntrusionSetInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IntrusionSetInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IntrusionSetInput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAliases

`func (o *IntrusionSetInput) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *IntrusionSetInput) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *IntrusionSetInput) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *IntrusionSetInput) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFirstSeen

`func (o *IntrusionSetInput) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *IntrusionSetInput) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *IntrusionSetInput) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *IntrusionSetInput) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *IntrusionSetInput) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *IntrusionSetInput) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *IntrusionSetInput) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *IntrusionSetInput) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


