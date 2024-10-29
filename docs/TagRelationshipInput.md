# TagRelationshipInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Target** | **string** |  | 
**LastSeen** | **time.Time** |  | 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Fresh** | **bool** |  | 

## Methods

### NewTagRelationshipInput

`func NewTagRelationshipInput(source string, target string, lastSeen time.Time, fresh bool, ) *TagRelationshipInput`

NewTagRelationshipInput instantiates a new TagRelationshipInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagRelationshipInputWithDefaults

`func NewTagRelationshipInputWithDefaults() *TagRelationshipInput`

NewTagRelationshipInputWithDefaults instantiates a new TagRelationshipInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TagRelationshipInput) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TagRelationshipInput) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TagRelationshipInput) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *TagRelationshipInput) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TagRelationshipInput) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TagRelationshipInput) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetLastSeen

`func (o *TagRelationshipInput) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *TagRelationshipInput) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *TagRelationshipInput) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetExpires

`func (o *TagRelationshipInput) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TagRelationshipInput) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TagRelationshipInput) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TagRelationshipInput) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *TagRelationshipInput) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *TagRelationshipInput) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFresh

`func (o *TagRelationshipInput) GetFresh() bool`

GetFresh returns the Fresh field if non-nil, zero value otherwise.

### GetFreshOk

`func (o *TagRelationshipInput) GetFreshOk() (*bool, bool)`

GetFreshOk returns a tuple with the Fresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFresh

`func (o *TagRelationshipInput) SetFresh(v bool)`

SetFresh sets Fresh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


