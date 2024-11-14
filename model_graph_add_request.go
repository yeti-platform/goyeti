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

// checks if the GraphAddRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphAddRequest{}

// GraphAddRequest struct for GraphAddRequest
type GraphAddRequest struct {
	Source               string `json:"source"`
	Target               string `json:"target"`
	LinkType             string `json:"link_type"`
	Description          string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _GraphAddRequest GraphAddRequest

// NewGraphAddRequest instantiates a new GraphAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphAddRequest(source string, target string, linkType string, description string) *GraphAddRequest {
	this := GraphAddRequest{}
	this.Source = source
	this.Target = target
	this.LinkType = linkType
	this.Description = description
	return &this
}

// NewGraphAddRequestWithDefaults instantiates a new GraphAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphAddRequestWithDefaults() *GraphAddRequest {
	this := GraphAddRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *GraphAddRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GraphAddRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GraphAddRequest) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *GraphAddRequest) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *GraphAddRequest) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *GraphAddRequest) SetTarget(v string) {
	o.Target = v
}

// GetLinkType returns the LinkType field value
func (o *GraphAddRequest) GetLinkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value
// and a boolean to check if the value has been set.
func (o *GraphAddRequest) GetLinkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkType, true
}

// SetLinkType sets field value
func (o *GraphAddRequest) SetLinkType(v string) {
	o.LinkType = v
}

// GetDescription returns the Description field value
func (o *GraphAddRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GraphAddRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GraphAddRequest) SetDescription(v string) {
	o.Description = v
}

func (o GraphAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphAddRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["link_type"] = o.LinkType
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GraphAddRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"target",
		"link_type",
		"description",
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

	varGraphAddRequest := _GraphAddRequest{}

	err = json.Unmarshal(data, &varGraphAddRequest)

	if err != nil {
		return err
	}

	*o = GraphAddRequest(varGraphAddRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "target")
		delete(additionalProperties, "link_type")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGraphAddRequest struct {
	value *GraphAddRequest
	isSet bool
}

func (v NullableGraphAddRequest) Get() *GraphAddRequest {
	return v.value
}

func (v *NullableGraphAddRequest) Set(val *GraphAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphAddRequest(val *GraphAddRequest) *NullableGraphAddRequest {
	return &NullableGraphAddRequest{value: val, isSet: true}
}

func (v NullableGraphAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
