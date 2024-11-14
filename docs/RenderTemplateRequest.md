# RenderTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateName** | **string** |  | 
**ObservableIds** | Pointer to **[]string** |  | [optional] [default to []]
**SearchQuery** | Pointer to [**SearchQuery**](SearchQuery.md) |  | [optional] 

## Methods

### NewRenderTemplateRequest

`func NewRenderTemplateRequest(templateName string, ) *RenderTemplateRequest`

NewRenderTemplateRequest instantiates a new RenderTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderTemplateRequestWithDefaults

`func NewRenderTemplateRequestWithDefaults() *RenderTemplateRequest`

NewRenderTemplateRequestWithDefaults instantiates a new RenderTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateName

`func (o *RenderTemplateRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *RenderTemplateRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *RenderTemplateRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetObservableIds

`func (o *RenderTemplateRequest) GetObservableIds() []string`

GetObservableIds returns the ObservableIds field if non-nil, zero value otherwise.

### GetObservableIdsOk

`func (o *RenderTemplateRequest) GetObservableIdsOk() (*[]string, bool)`

GetObservableIdsOk returns a tuple with the ObservableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservableIds

`func (o *RenderTemplateRequest) SetObservableIds(v []string)`

SetObservableIds sets ObservableIds field to given value.

### HasObservableIds

`func (o *RenderTemplateRequest) HasObservableIds() bool`

HasObservableIds returns a boolean if a field has been set.

### GetSearchQuery

`func (o *RenderTemplateRequest) GetSearchQuery() SearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *RenderTemplateRequest) GetSearchQueryOk() (*SearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *RenderTemplateRequest) SetSearchQuery(v SearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *RenderTemplateRequest) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


