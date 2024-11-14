# CommandLineInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "command_line"]
**Created** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**LastAnalysis** | Pointer to [**map[string]time.Time**](time.Time.md) |  | [optional] [default to {}]

## Methods

### NewCommandLineInput

`func NewCommandLineInput(value string, ) *CommandLineInput`

NewCommandLineInput instantiates a new CommandLineInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandLineInputWithDefaults

`func NewCommandLineInputWithDefaults() *CommandLineInput`

NewCommandLineInputWithDefaults instantiates a new CommandLineInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CommandLineInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CommandLineInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CommandLineInput) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *CommandLineInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommandLineInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommandLineInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommandLineInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *CommandLineInput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CommandLineInput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CommandLineInput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CommandLineInput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetContext

`func (o *CommandLineInput) GetContext() []map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CommandLineInput) GetContextOk() (*[]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CommandLineInput) SetContext(v []map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *CommandLineInput) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastAnalysis

`func (o *CommandLineInput) GetLastAnalysis() map[string]time.Time`

GetLastAnalysis returns the LastAnalysis field if non-nil, zero value otherwise.

### GetLastAnalysisOk

`func (o *CommandLineInput) GetLastAnalysisOk() (*map[string]time.Time, bool)`

GetLastAnalysisOk returns a tuple with the LastAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysis

`func (o *CommandLineInput) SetLastAnalysis(v map[string]time.Time)`

SetLastAnalysis sets LastAnalysis field to given value.

### HasLastAnalysis

`func (o *CommandLineInput) HasLastAnalysis() bool`

HasLastAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


