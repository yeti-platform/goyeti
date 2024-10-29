# AnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Observables** | **[]string** |  | 
**AddTags** | Pointer to **[]string** |  | [optional] [default to []]
**RegexMatch** | Pointer to **bool** |  | [optional] [default to false]
**AddType** | Pointer to [**NullableObservableTypeInput**](ObservableTypeInput.md) |  | [optional] 
**FetchNeighbors** | Pointer to **bool** |  | [optional] [default to true]
**AddUnknown** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAnalysisRequest

`func NewAnalysisRequest(observables []string, ) *AnalysisRequest`

NewAnalysisRequest instantiates a new AnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisRequestWithDefaults

`func NewAnalysisRequestWithDefaults() *AnalysisRequest`

NewAnalysisRequestWithDefaults instantiates a new AnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservables

`func (o *AnalysisRequest) GetObservables() []string`

GetObservables returns the Observables field if non-nil, zero value otherwise.

### GetObservablesOk

`func (o *AnalysisRequest) GetObservablesOk() (*[]string, bool)`

GetObservablesOk returns a tuple with the Observables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservables

`func (o *AnalysisRequest) SetObservables(v []string)`

SetObservables sets Observables field to given value.


### GetAddTags

`func (o *AnalysisRequest) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *AnalysisRequest) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *AnalysisRequest) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *AnalysisRequest) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRegexMatch

`func (o *AnalysisRequest) GetRegexMatch() bool`

GetRegexMatch returns the RegexMatch field if non-nil, zero value otherwise.

### GetRegexMatchOk

`func (o *AnalysisRequest) GetRegexMatchOk() (*bool, bool)`

GetRegexMatchOk returns a tuple with the RegexMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexMatch

`func (o *AnalysisRequest) SetRegexMatch(v bool)`

SetRegexMatch sets RegexMatch field to given value.

### HasRegexMatch

`func (o *AnalysisRequest) HasRegexMatch() bool`

HasRegexMatch returns a boolean if a field has been set.

### GetAddType

`func (o *AnalysisRequest) GetAddType() ObservableTypeInput`

GetAddType returns the AddType field if non-nil, zero value otherwise.

### GetAddTypeOk

`func (o *AnalysisRequest) GetAddTypeOk() (*ObservableTypeInput, bool)`

GetAddTypeOk returns a tuple with the AddType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddType

`func (o *AnalysisRequest) SetAddType(v ObservableTypeInput)`

SetAddType sets AddType field to given value.

### HasAddType

`func (o *AnalysisRequest) HasAddType() bool`

HasAddType returns a boolean if a field has been set.

### SetAddTypeNil

`func (o *AnalysisRequest) SetAddTypeNil(b bool)`

 SetAddTypeNil sets the value for AddType to be an explicit nil

### UnsetAddType
`func (o *AnalysisRequest) UnsetAddType()`

UnsetAddType ensures that no value is present for AddType, not even an explicit nil
### GetFetchNeighbors

`func (o *AnalysisRequest) GetFetchNeighbors() bool`

GetFetchNeighbors returns the FetchNeighbors field if non-nil, zero value otherwise.

### GetFetchNeighborsOk

`func (o *AnalysisRequest) GetFetchNeighborsOk() (*bool, bool)`

GetFetchNeighborsOk returns a tuple with the FetchNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchNeighbors

`func (o *AnalysisRequest) SetFetchNeighbors(v bool)`

SetFetchNeighbors sets FetchNeighbors field to given value.

### HasFetchNeighbors

`func (o *AnalysisRequest) HasFetchNeighbors() bool`

HasFetchNeighbors returns a boolean if a field has been set.

### GetAddUnknown

`func (o *AnalysisRequest) GetAddUnknown() bool`

GetAddUnknown returns the AddUnknown field if non-nil, zero value otherwise.

### GetAddUnknownOk

`func (o *AnalysisRequest) GetAddUnknownOk() (*bool, bool)`

GetAddUnknownOk returns a tuple with the AddUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUnknown

`func (o *AnalysisRequest) SetAddUnknown(v bool)`

SetAddUnknown sets AddUnknown field to given value.

### HasAddUnknown

`func (o *AnalysisRequest) HasAddUnknown() bool`

HasAddUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


