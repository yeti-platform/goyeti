# SystemConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **map[string]interface{}** |  | 
**System** | **map[string]interface{}** |  | 

## Methods

### NewSystemConfigResponse

`func NewSystemConfigResponse(auth map[string]interface{}, system map[string]interface{}, ) *SystemConfigResponse`

NewSystemConfigResponse instantiates a new SystemConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigResponseWithDefaults

`func NewSystemConfigResponseWithDefaults() *SystemConfigResponse`

NewSystemConfigResponseWithDefaults instantiates a new SystemConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *SystemConfigResponse) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SystemConfigResponse) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SystemConfigResponse) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.


### GetSystem

`func (o *SystemConfigResponse) GetSystem() map[string]interface{}`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SystemConfigResponse) GetSystemOk() (*map[string]interface{}, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SystemConfigResponse) SetSystem(v map[string]interface{})`

SetSystem sets System field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


