# TagRelationshipOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Target** | **string** |  | 
**LastSeen** | **time.Time** |  | 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Fresh** | **bool** |  | 
**Id** | **string** |  | [readonly] 

## Methods

### NewTagRelationshipOutput

`func NewTagRelationshipOutput(source string, target string, lastSeen time.Time, fresh bool, id string, ) *TagRelationshipOutput`

NewTagRelationshipOutput instantiates a new TagRelationshipOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagRelationshipOutputWithDefaults

`func NewTagRelationshipOutputWithDefaults() *TagRelationshipOutput`

NewTagRelationshipOutputWithDefaults instantiates a new TagRelationshipOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TagRelationshipOutput) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TagRelationshipOutput) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TagRelationshipOutput) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *TagRelationshipOutput) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TagRelationshipOutput) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TagRelationshipOutput) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetLastSeen

`func (o *TagRelationshipOutput) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *TagRelationshipOutput) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *TagRelationshipOutput) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetExpires

`func (o *TagRelationshipOutput) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TagRelationshipOutput) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TagRelationshipOutput) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TagRelationshipOutput) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *TagRelationshipOutput) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *TagRelationshipOutput) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFresh

`func (o *TagRelationshipOutput) GetFresh() bool`

GetFresh returns the Fresh field if non-nil, zero value otherwise.

### GetFreshOk

`func (o *TagRelationshipOutput) GetFreshOk() (*bool, bool)`

GetFreshOk returns a tuple with the Fresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFresh

`func (o *TagRelationshipOutput) SetFresh(v bool)`

SetFresh sets Fresh field to given value.


### GetId

`func (o *TagRelationshipOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagRelationshipOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagRelationshipOutput) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


