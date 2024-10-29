# AnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | **[][]interface{}** |  | 
**Observables** | **[][]interface{}** |  | 
**Known** | [**[]Observable**](Observable.md) |  | 
**Matches** | **[][]interface{}** |  | 
**Unknown** | **[]string** |  | 

## Methods

### NewAnalysisResponse

`func NewAnalysisResponse(entities [][]interface{}, observables [][]interface{}, known []Observable, matches [][]interface{}, unknown []string, ) *AnalysisResponse`

NewAnalysisResponse instantiates a new AnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisResponseWithDefaults

`func NewAnalysisResponseWithDefaults() *AnalysisResponse`

NewAnalysisResponseWithDefaults instantiates a new AnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AnalysisResponse) GetEntities() [][]interface{}`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AnalysisResponse) GetEntitiesOk() (*[][]interface{}, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AnalysisResponse) SetEntities(v [][]interface{})`

SetEntities sets Entities field to given value.


### GetObservables

`func (o *AnalysisResponse) GetObservables() [][]interface{}`

GetObservables returns the Observables field if non-nil, zero value otherwise.

### GetObservablesOk

`func (o *AnalysisResponse) GetObservablesOk() (*[][]interface{}, bool)`

GetObservablesOk returns a tuple with the Observables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservables

`func (o *AnalysisResponse) SetObservables(v [][]interface{})`

SetObservables sets Observables field to given value.


### GetKnown

`func (o *AnalysisResponse) GetKnown() []Observable`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *AnalysisResponse) GetKnownOk() (*[]Observable, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *AnalysisResponse) SetKnown(v []Observable)`

SetKnown sets Known field to given value.


### GetMatches

`func (o *AnalysisResponse) GetMatches() [][]interface{}`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *AnalysisResponse) GetMatchesOk() (*[][]interface{}, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *AnalysisResponse) SetMatches(v [][]interface{})`

SetMatches sets Matches field to given value.


### GetUnknown

`func (o *AnalysisResponse) GetUnknown() []string`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *AnalysisResponse) GetUnknownOk() (*[]string, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *AnalysisResponse) SetUnknown(v []string)`

SetUnknown sets Unknown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


