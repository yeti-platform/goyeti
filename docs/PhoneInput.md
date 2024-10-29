# PhoneInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "phone"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPhoneInput

`func NewPhoneInput(name string, ) *PhoneInput`

NewPhoneInput instantiates a new PhoneInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneInputWithDefaults

`func NewPhoneInputWithDefaults() *PhoneInput`

NewPhoneInputWithDefaults instantiates a new PhoneInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PhoneInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhoneInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhoneInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PhoneInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PhoneInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhoneInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhoneInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PhoneInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhoneInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhoneInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhoneInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *PhoneInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PhoneInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PhoneInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *PhoneInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreated

`func (o *PhoneInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PhoneInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PhoneInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PhoneInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *PhoneInput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PhoneInput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PhoneInput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *PhoneInput) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


