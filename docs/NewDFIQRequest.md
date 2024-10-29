# NewDFIQRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfiqYaml** | **string** |  | 
**DfiqType** | [**DFIQType**](DFIQType.md) |  | 
**UpdateIndicators** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNewDFIQRequest

`func NewNewDFIQRequest(dfiqYaml string, dfiqType DFIQType, ) *NewDFIQRequest`

NewNewDFIQRequest instantiates a new NewDFIQRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDFIQRequestWithDefaults

`func NewNewDFIQRequestWithDefaults() *NewDFIQRequest`

NewNewDFIQRequestWithDefaults instantiates a new NewDFIQRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfiqYaml

`func (o *NewDFIQRequest) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *NewDFIQRequest) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *NewDFIQRequest) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetDfiqType

`func (o *NewDFIQRequest) GetDfiqType() DFIQType`

GetDfiqType returns the DfiqType field if non-nil, zero value otherwise.

### GetDfiqTypeOk

`func (o *NewDFIQRequest) GetDfiqTypeOk() (*DFIQType, bool)`

GetDfiqTypeOk returns a tuple with the DfiqType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqType

`func (o *NewDFIQRequest) SetDfiqType(v DFIQType)`

SetDfiqType sets DfiqType field to given value.


### GetUpdateIndicators

`func (o *NewDFIQRequest) GetUpdateIndicators() bool`

GetUpdateIndicators returns the UpdateIndicators field if non-nil, zero value otherwise.

### GetUpdateIndicatorsOk

`func (o *NewDFIQRequest) GetUpdateIndicatorsOk() (*bool, bool)`

GetUpdateIndicatorsOk returns a tuple with the UpdateIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateIndicators

`func (o *NewDFIQRequest) SetUpdateIndicators(v bool)`

SetUpdateIndicators sets UpdateIndicators field to given value.

### HasUpdateIndicators

`func (o *NewDFIQRequest) HasUpdateIndicators() bool`

HasUpdateIndicators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


