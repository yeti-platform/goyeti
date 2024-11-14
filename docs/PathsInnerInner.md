# PathsInnerInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Target** | **string** |  | 
**Type** | **string** |  | 
**Count** | Pointer to **int32** |  | [optional] [default to 1]
**Description** | **string** |  | 
**Created** | **time.Time** |  | 
**Modified** | **time.Time** |  | 
**Id** | **string** |  | [readonly] 
**LastSeen** | **time.Time** |  | 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Fresh** | **bool** |  | 

## Methods

### NewPathsInnerInner

`func NewPathsInnerInner(source string, target string, type_ string, description string, created time.Time, modified time.Time, id string, lastSeen time.Time, fresh bool, ) *PathsInnerInner`

NewPathsInnerInner instantiates a new PathsInnerInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathsInnerInnerWithDefaults

`func NewPathsInnerInnerWithDefaults() *PathsInnerInner`

NewPathsInnerInnerWithDefaults instantiates a new PathsInnerInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PathsInnerInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PathsInnerInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PathsInnerInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *PathsInnerInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PathsInnerInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PathsInnerInner) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *PathsInnerInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PathsInnerInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PathsInnerInner) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *PathsInnerInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PathsInnerInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PathsInnerInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PathsInnerInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDescription

`func (o *PathsInnerInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PathsInnerInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PathsInnerInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *PathsInnerInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PathsInnerInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PathsInnerInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *PathsInnerInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PathsInnerInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PathsInnerInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetId

`func (o *PathsInnerInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PathsInnerInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PathsInnerInner) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeen

`func (o *PathsInnerInner) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *PathsInnerInner) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *PathsInnerInner) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetExpires

`func (o *PathsInnerInner) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PathsInnerInner) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PathsInnerInner) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PathsInnerInner) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFresh

`func (o *PathsInnerInner) GetFresh() bool`

GetFresh returns the Fresh field if non-nil, zero value otherwise.

### GetFreshOk

`func (o *PathsInnerInner) GetFreshOk() (*bool, bool)`

GetFreshOk returns a tuple with the Fresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFresh

`func (o *PathsInnerInner) SetFresh(v bool)`

SetFresh sets Fresh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


