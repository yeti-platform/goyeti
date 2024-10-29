# DFIQValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfiqYaml** | **string** |  | 
**DfiqType** | [**DFIQType**](DFIQType.md) |  | 
**UpdateIndicators** | Pointer to **bool** |  | [optional] [default to false]
**CheckId** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDFIQValidateRequest

`func NewDFIQValidateRequest(dfiqYaml string, dfiqType DFIQType, ) *DFIQValidateRequest`

NewDFIQValidateRequest instantiates a new DFIQValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQValidateRequestWithDefaults

`func NewDFIQValidateRequestWithDefaults() *DFIQValidateRequest`

NewDFIQValidateRequestWithDefaults instantiates a new DFIQValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfiqYaml

`func (o *DFIQValidateRequest) GetDfiqYaml() string`

GetDfiqYaml returns the DfiqYaml field if non-nil, zero value otherwise.

### GetDfiqYamlOk

`func (o *DFIQValidateRequest) GetDfiqYamlOk() (*string, bool)`

GetDfiqYamlOk returns a tuple with the DfiqYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqYaml

`func (o *DFIQValidateRequest) SetDfiqYaml(v string)`

SetDfiqYaml sets DfiqYaml field to given value.


### GetDfiqType

`func (o *DFIQValidateRequest) GetDfiqType() DFIQType`

GetDfiqType returns the DfiqType field if non-nil, zero value otherwise.

### GetDfiqTypeOk

`func (o *DFIQValidateRequest) GetDfiqTypeOk() (*DFIQType, bool)`

GetDfiqTypeOk returns a tuple with the DfiqType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfiqType

`func (o *DFIQValidateRequest) SetDfiqType(v DFIQType)`

SetDfiqType sets DfiqType field to given value.


### GetUpdateIndicators

`func (o *DFIQValidateRequest) GetUpdateIndicators() bool`

GetUpdateIndicators returns the UpdateIndicators field if non-nil, zero value otherwise.

### GetUpdateIndicatorsOk

`func (o *DFIQValidateRequest) GetUpdateIndicatorsOk() (*bool, bool)`

GetUpdateIndicatorsOk returns a tuple with the UpdateIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateIndicators

`func (o *DFIQValidateRequest) SetUpdateIndicators(v bool)`

SetUpdateIndicators sets UpdateIndicators field to given value.

### HasUpdateIndicators

`func (o *DFIQValidateRequest) HasUpdateIndicators() bool`

HasUpdateIndicators returns a boolean if a field has been set.

### GetCheckId

`func (o *DFIQValidateRequest) GetCheckId() bool`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *DFIQValidateRequest) GetCheckIdOk() (*bool, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *DFIQValidateRequest) SetCheckId(v bool)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *DFIQValidateRequest) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


