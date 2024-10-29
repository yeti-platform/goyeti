# WorkerStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registered** | **map[string][]string** |  | 
**Active** | **[][]interface{}** |  | 

## Methods

### NewWorkerStatusResponse

`func NewWorkerStatusResponse(registered map[string][]string, active [][]interface{}, ) *WorkerStatusResponse`

NewWorkerStatusResponse instantiates a new WorkerStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerStatusResponseWithDefaults

`func NewWorkerStatusResponseWithDefaults() *WorkerStatusResponse`

NewWorkerStatusResponseWithDefaults instantiates a new WorkerStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistered

`func (o *WorkerStatusResponse) GetRegistered() map[string][]string`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *WorkerStatusResponse) GetRegisteredOk() (*map[string][]string, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *WorkerStatusResponse) SetRegistered(v map[string][]string)`

SetRegistered sets Registered field to given value.


### GetActive

`func (o *WorkerStatusResponse) GetActive() [][]interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkerStatusResponse) GetActiveOk() (*[][]interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkerStatusResponse) SetActive(v [][]interface{})`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


