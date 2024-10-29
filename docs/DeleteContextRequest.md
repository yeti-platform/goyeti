# DeleteContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Context** | **map[string]interface{}** |  | 
**SkipCompare** | Pointer to **[]interface{}** |  | [optional] [default to []]

## Methods

### NewDeleteContextRequest

`func NewDeleteContextRequest(source string, context map[string]interface{}, ) *DeleteContextRequest`

NewDeleteContextRequest instantiates a new DeleteContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteContextRequestWithDefaults

`func NewDeleteContextRequestWithDefaults() *DeleteContextRequest`

NewDeleteContextRequestWithDefaults instantiates a new DeleteContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DeleteContextRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeleteContextRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeleteContextRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetContext

`func (o *DeleteContextRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DeleteContextRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DeleteContextRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.


### GetSkipCompare

`func (o *DeleteContextRequest) GetSkipCompare() []interface{}`

GetSkipCompare returns the SkipCompare field if non-nil, zero value otherwise.

### GetSkipCompareOk

`func (o *DeleteContextRequest) GetSkipCompareOk() (*[]interface{}, bool)`

GetSkipCompareOk returns a tuple with the SkipCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCompare

`func (o *DeleteContextRequest) SetSkipCompare(v []interface{})`

SetSkipCompare sets SkipCompare field to given value.

### HasSkipCompare

`func (o *DeleteContextRequest) HasSkipCompare() bool`

HasSkipCompare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


