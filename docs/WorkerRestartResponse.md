# WorkerRestartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successes** | **[]string** |  | 
**Failures** | **[]string** |  | 

## Methods

### NewWorkerRestartResponse

`func NewWorkerRestartResponse(successes []string, failures []string, ) *WorkerRestartResponse`

NewWorkerRestartResponse instantiates a new WorkerRestartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerRestartResponseWithDefaults

`func NewWorkerRestartResponseWithDefaults() *WorkerRestartResponse`

NewWorkerRestartResponseWithDefaults instantiates a new WorkerRestartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccesses

`func (o *WorkerRestartResponse) GetSuccesses() []string`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *WorkerRestartResponse) GetSuccessesOk() (*[]string, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *WorkerRestartResponse) SetSuccesses(v []string)`

SetSuccesses sets Successes field to given value.


### GetFailures

`func (o *WorkerRestartResponse) GetFailures() []string`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *WorkerRestartResponse) GetFailuresOk() (*[]string, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *WorkerRestartResponse) SetFailures(v []string)`

SetFailures sets Failures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


