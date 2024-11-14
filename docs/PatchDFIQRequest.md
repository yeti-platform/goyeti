# PatchDFIQRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfiqYaml** | **string** |  | 
**DfiqType** | [**DFIQType**](DFIQType.md) |  | 
**UpdateIndicators** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPatchDFIQRequest

`func NewPatchDFIQRequest(dfiqYaml string, dfiqType DFIQType, ) *PatchDFIQRequest`

NewPatchDFIQRequest instantiates a new PatchDFIQRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDFIQRequestWithDefaults

`func NewPatchDFIQRequestWithDefaults() *PatchDFIQRequest`

NewPatchDFIQRequestWithDefaults instantiates a new PatchDFIQRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfiqYaml

`func (o *PatchDFIQRequest) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *PatchDFIQRequest) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *PatchDFIQRequest) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetDfiqType

`func (o *PatchDFIQRequest) GetDfiqType() DFIQType`

GetDfiqType returns the DfiqType field if non-nil, zero value otherwise.

### GetDfiqTypeOk

`func (o *PatchDFIQRequest) GetDfiqTypeOk() (*DFIQType, bool)`

GetDfiqTypeOk returns a tuple with the DfiqType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqType

`func (o *PatchDFIQRequest) SetDfiqType(v DFIQType)`

SetDfiqType sets DfiqType field to given value.


### GetUpdateIndicators

`func (o *PatchDFIQRequest) GetUpdateIndicators() bool`

GetUpdateIndicators returns the UpdateIndicators field if non-nil, zero value otherwise.

### GetUpdateIndicatorsOk

`func (o *PatchDFIQRequest) GetUpdateIndicatorsOk() (*bool, bool)`

GetUpdateIndicatorsOk returns a tuple with the UpdateIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateIndicators

`func (o *PatchDFIQRequest) SetUpdateIndicators(v bool)`

SetUpdateIndicators sets UpdateIndicators field to given value.

### HasUpdateIndicators

`func (o *PatchDFIQRequest) HasUpdateIndicators() bool`

HasUpdateIndicators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


