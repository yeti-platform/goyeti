# RegistryKeyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "registry_key"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]
**Key** | **string** |  | 
**Data** | ***os.File** |  | 
**Hive** | [**RegistryHive**](RegistryHive.md) |  | 
**PathFile** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRegistryKeyInput

`func NewRegistryKeyInput(value string, key string, data *os.File, hive RegistryHive, ) *RegistryKeyInput`

NewRegistryKeyInput instantiates a new RegistryKeyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryKeyInputWithDefaults

`func NewRegistryKeyInputWithDefaults() *RegistryKeyInput`

NewRegistryKeyInputWithDefaults instantiates a new RegistryKeyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RegistryKeyInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RegistryKeyInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RegistryKeyInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *RegistryKeyInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryKeyInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryKeyInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistryKeyInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *RegistryKeyInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RegistryKeyInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RegistryKeyInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RegistryKeyInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *RegistryKeyInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RegistryKeyInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RegistryKeyInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *RegistryKeyInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *RegistryKeyInput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *RegistryKeyInput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *RegistryKeyInput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *RegistryKeyInput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.

### GetKey

`func (o *RegistryKeyInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RegistryKeyInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RegistryKeyInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetData

`func (o *RegistryKeyInput) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegistryKeyInput) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegistryKeyInput) SetData(v *os.File)`

SetData sets Data field to given value.


### GetHive

`func (o *RegistryKeyInput) GetHive() RegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *RegistryKeyInput) GetHiveOk() (*RegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *RegistryKeyInput) SetHive(v RegistryHive)`

SetHive sets Hive field to given value.


### GetPathFile

`func (o *RegistryKeyInput) GetPathFile() string`

GetPathFile returns the PathFile field if non-nil, zero value otherwise.

### GetPathFileOk

`func (o *RegistryKeyInput) GetPathFileOk() (*string, bool)`

GetPathFileOk returns a tuple with the PathFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFile

`func (o *RegistryKeyInput) SetPathFile(v string)`

SetPathFile sets PathFile field to given value.

### HasPathFile

`func (o *RegistryKeyInput) HasPathFile() bool`

HasPathFile returns a boolean if a field has been set.

### SetPathFileNil

`func (o *RegistryKeyInput) SetPathFileNil(b bool)`

 SetPathFileNil sets the value for PathFile to be an explicit nil

### UnsetPathFile
`func (o *RegistryKeyInput) UnsetPathFile()`

UnsetPathFile ensures that no value is present for PathFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


