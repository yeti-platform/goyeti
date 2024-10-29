# UpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DefaultExpiration** | Pointer to **string** |  | [optional] [default to "P90D"]
**Produces** | Pointer to **[]string** |  | [optional] [default to []]
**Replaces** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewUpdateRequest

`func NewUpdateRequest(name string, ) *UpdateRequest`

NewUpdateRequest instantiates a new UpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestWithDefaults

`func NewUpdateRequestWithDefaults() *UpdateRequest`

NewUpdateRequestWithDefaults instantiates a new UpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultExpiration

`func (o *UpdateRequest) GetDefaultExpiration() string`

GetDefaultExpiration returns the DefaultExpiration field if non-nil, zero value otherwise.

### GetDefaultExpirationOk

`func (o *UpdateRequest) GetDefaultExpirationOk() (*string, bool)`

GetDefaultExpirationOk returns a tuple with the DefaultExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpiration

`func (o *UpdateRequest) SetDefaultExpiration(v string)`

SetDefaultExpiration sets DefaultExpiration field to given value.

### HasDefaultExpiration

`func (o *UpdateRequest) HasDefaultExpiration() bool`

HasDefaultExpiration returns a boolean if a field has been set.

### GetProduces

`func (o *UpdateRequest) GetProduces() []string`

GetProduces returns the Produces field if non-nil, zero value otherwise.

### GetProducesOk

`func (o *UpdateRequest) GetProducesOk() (*[]string, bool)`

GetProducesOk returns a tuple with the Produces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduces

`func (o *UpdateRequest) SetProduces(v []string)`

SetProduces sets Produces field to given value.

### HasProduces

`func (o *UpdateRequest) HasProduces() bool`

HasProduces returns a boolean if a field has been set.

### GetReplaces

`func (o *UpdateRequest) GetReplaces() []string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *UpdateRequest) GetReplacesOk() (*[]string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *UpdateRequest) SetReplaces(v []string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *UpdateRequest) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


