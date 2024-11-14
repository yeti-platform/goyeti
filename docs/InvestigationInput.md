# InvestigationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "investigation"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewInvestigationInput

`func NewInvestigationInput(name string, ) *InvestigationInput`

NewInvestigationInput instantiates a new InvestigationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationInputWithDefaults

`func NewInvestigationInputWithDefaults() *InvestigationInput`

NewInvestigationInputWithDefaults instantiates a new InvestigationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvestigationInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvestigationInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvestigationInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvestigationInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InvestigationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestigationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestigationInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvestigationInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvestigationInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvestigationInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvestigationInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *InvestigationInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *InvestigationInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *InvestigationInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *InvestigationInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *InvestigationInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InvestigationInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InvestigationInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InvestigationInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *InvestigationInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InvestigationInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InvestigationInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *InvestigationInput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetReference

`func (o *InvestigationInput) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InvestigationInput) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InvestigationInput) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *InvestigationInput) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


