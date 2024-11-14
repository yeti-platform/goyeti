# TemplateSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | [**[]Template**](Template.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewTemplateSearchResponse

`func NewTemplateSearchResponse(templates []Template, total int32, ) *TemplateSearchResponse`

NewTemplateSearchResponse instantiates a new TemplateSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateSearchResponseWithDefaults

`func NewTemplateSearchResponseWithDefaults() *TemplateSearchResponse`

NewTemplateSearchResponseWithDefaults instantiates a new TemplateSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *TemplateSearchResponse) GetTemplates() []Template`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *TemplateSearchResponse) GetTemplatesOk() (*[]Template, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *TemplateSearchResponse) SetTemplates(v []Template)`

SetTemplates sets Templates field to given value.


### GetTotal

`func (o *TemplateSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TemplateSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TemplateSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


