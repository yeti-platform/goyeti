/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"encoding/json"
	"fmt"
)

// checks if the TemplateSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateSearchResponse{}

// TemplateSearchResponse struct for TemplateSearchResponse
type TemplateSearchResponse struct {
	Templates            []Template `json:"templates"`
	Total                int32      `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _TemplateSearchResponse TemplateSearchResponse

// NewTemplateSearchResponse instantiates a new TemplateSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSearchResponse(templates []Template, total int32) *TemplateSearchResponse {
	this := TemplateSearchResponse{}
	this.Templates = templates
	this.Total = total
	return &this
}

// NewTemplateSearchResponseWithDefaults instantiates a new TemplateSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSearchResponseWithDefaults() *TemplateSearchResponse {
	this := TemplateSearchResponse{}
	return &this
}

// GetTemplates returns the Templates field value
func (o *TemplateSearchResponse) GetTemplates() []Template {
	if o == nil {
		var ret []Template
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *TemplateSearchResponse) GetTemplatesOk() ([]Template, bool) {
	if o == nil {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *TemplateSearchResponse) SetTemplates(v []Template) {
	o.Templates = v
}

// GetTotal returns the Total field value
func (o *TemplateSearchResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TemplateSearchResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *TemplateSearchResponse) SetTotal(v int32) {
	o.Total = v
}

func (o TemplateSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templates"] = o.Templates
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateSearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templates",
		"total",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplateSearchResponse := _TemplateSearchResponse{}

	err = json.Unmarshal(data, &varTemplateSearchResponse)

	if err != nil {
		return err
	}

	*o = TemplateSearchResponse(varTemplateSearchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "templates")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateSearchResponse struct {
	value *TemplateSearchResponse
	isSet bool
}

func (v NullableTemplateSearchResponse) Get() *TemplateSearchResponse {
	return v.value
}

func (v *NullableTemplateSearchResponse) Set(val *TemplateSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSearchResponse(val *TemplateSearchResponse) *NullableTemplateSearchResponse {
	return &NullableTemplateSearchResponse{value: val, isSet: true}
}

func (v NullableTemplateSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
