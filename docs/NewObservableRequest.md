# NewObservableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] [default to []]
**Value** | **string** |  | 
**Type** | [**ObservableTypeInput**](ObservableTypeInput.md) |  | 

## Methods

### NewNewObservableRequest

`func NewNewObservableRequest(value string, type_ ObservableTypeInput, ) *NewObservableRequest`

NewNewObservableRequest instantiates a new NewObservableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewObservableRequestWithDefaults

`func NewNewObservableRequestWithDefaults() *NewObservableRequest`

NewNewObservableRequestWithDefaults instantiates a new NewObservableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *NewObservableRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NewObservableRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NewObservableRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NewObservableRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValue

`func (o *NewObservableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewObservableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewObservableRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *NewObservableRequest) GetType() ObservableTypeInput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewObservableRequest) GetTypeOk() (*ObservableTypeInput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewObservableRequest) SetType(v ObservableTypeInput)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


