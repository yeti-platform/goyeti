# TaskSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | [**[]TasksInner**](TasksInner.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewTaskSearchResponse

`func NewTaskSearchResponse(tasks []TasksInner, total int32, ) *TaskSearchResponse`

NewTaskSearchResponse instantiates a new TaskSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSearchResponseWithDefaults

`func NewTaskSearchResponseWithDefaults() *TaskSearchResponse`

NewTaskSearchResponseWithDefaults instantiates a new TaskSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *TaskSearchResponse) GetTasks() []TasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskSearchResponse) GetTasksOk() (*[]TasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TaskSearchResponse) SetTasks(v []TasksInner)`

SetTasks sets Tasks field to given value.


### GetTotal

`func (o *TaskSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TaskSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TaskSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


