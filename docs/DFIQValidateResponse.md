# DFIQValidateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] [default to ]
**ErrorType** | Pointer to **string** |  | [optional] [default to "message"]

## Methods

### NewDFIQValidateResponse

`func NewDFIQValidateResponse(valid bool, ) *DFIQValidateResponse`

NewDFIQValidateResponse instantiates a new DFIQValidateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDFIQValidateResponseWithDefaults

`func NewDFIQValidateResponseWithDefaults() *DFIQValidateResponse`

NewDFIQValidateResponseWithDefaults instantiates a new DFIQValidateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *DFIQValidateResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DFIQValidateResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DFIQValidateResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetError

`func (o *DFIQValidateResponse) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DFIQValidateResponse) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DFIQValidateResponse) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *DFIQValidateResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorType

`func (o *DFIQValidateResponse) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *DFIQValidateResponse) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *DFIQValidateResponse) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *DFIQValidateResponse) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


