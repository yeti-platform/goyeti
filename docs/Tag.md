# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Count** | Pointer to **int32** |  | [optional] [default to 0]
**Created** | Pointer to **time.Time** |  | [optional] 
**DefaultExpiration** | Pointer to **string** |  | [optional] [default to "P90D"]
**Produces** | Pointer to **[]string** |  | [optional] [default to []]
**Replaces** | Pointer to **[]string** |  | [optional] [default to []]
**Id** | **string** |  | [readonly] 

## Methods

### NewTag

`func NewTag(name string, id string, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### GetCount

`func (o *Tag) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Tag) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Tag) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Tag) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreated

`func (o *Tag) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Tag) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Tag) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Tag) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDefaultExpiration

`func (o *Tag) GetDefaultExpiration() string`

GetDefaultExpiration returns the DefaultExpiration field if non-nil, zero value otherwise.

### GetDefaultExpirationOk

`func (o *Tag) GetDefaultExpirationOk() (*string, bool)`

GetDefaultExpirationOk returns a tuple with the DefaultExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpiration

`func (o *Tag) SetDefaultExpiration(v string)`

SetDefaultExpiration sets DefaultExpiration field to given value.

### HasDefaultExpiration

`func (o *Tag) HasDefaultExpiration() bool`

HasDefaultExpiration returns a boolean if a field has been set.

### GetProduces

`func (o *Tag) GetProduces() []string`

GetProduces returns the Produces field if non-nil, zero value otherwise.

### GetProducesOk

`func (o *Tag) GetProducesOk() (*[]string, bool)`

GetProducesOk returns a tuple with the Produces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduces

`func (o *Tag) SetProduces(v []string)`

SetProduces sets Produces field to given value.

### HasProduces

`func (o *Tag) HasProduces() bool`

HasProduces returns a boolean if a field has been set.

### GetReplaces

`func (o *Tag) GetReplaces() []string`

GetReplaces returns the Replaces field if non-nil, zero value otherwise.

### GetReplacesOk

`func (o *Tag) GetReplacesOk() (*[]string, bool)`

GetReplacesOk returns a tuple with the Replaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaces

`func (o *Tag) SetReplaces(v []string)`

SetReplaces sets Replaces field to given value.

### HasReplaces

`func (o *Tag) HasReplaces() bool`

HasReplaces returns a boolean if a field has been set.

### GetId

`func (o *Tag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


