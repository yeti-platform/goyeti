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

// checks if the SystemConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemConfigResponse{}

// SystemConfigResponse System config template.
type SystemConfigResponse struct {
	Auth                 map[string]interface{} `json:"auth"`
	System               map[string]interface{} `json:"system"`
	AdditionalProperties map[string]interface{}
}

type _SystemConfigResponse SystemConfigResponse

// NewSystemConfigResponse instantiates a new SystemConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemConfigResponse(auth map[string]interface{}, system map[string]interface{}) *SystemConfigResponse {
	this := SystemConfigResponse{}
	this.Auth = auth
	this.System = system
	return &this
}

// NewSystemConfigResponseWithDefaults instantiates a new SystemConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemConfigResponseWithDefaults() *SystemConfigResponse {
	this := SystemConfigResponse{}
	return &this
}

// GetAuth returns the Auth field value
func (o *SystemConfigResponse) GetAuth() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *SystemConfigResponse) GetAuthOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Auth, true
}

// SetAuth sets field value
func (o *SystemConfigResponse) SetAuth(v map[string]interface{}) {
	o.Auth = v
}

// GetSystem returns the System field value
func (o *SystemConfigResponse) GetSystem() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *SystemConfigResponse) GetSystemOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.System, true
}

// SetSystem sets field value
func (o *SystemConfigResponse) SetSystem(v map[string]interface{}) {
	o.System = v
}

func (o SystemConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth"] = o.Auth
	toSerialize["system"] = o.System

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemConfigResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth",
		"system",
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

	varSystemConfigResponse := _SystemConfigResponse{}

	err = json.Unmarshal(data, &varSystemConfigResponse)

	if err != nil {
		return err
	}

	*o = SystemConfigResponse(varSystemConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth")
		delete(additionalProperties, "system")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemConfigResponse struct {
	value *SystemConfigResponse
	isSet bool
}

func (v NullableSystemConfigResponse) Get() *SystemConfigResponse {
	return v.value
}

func (v *NullableSystemConfigResponse) Set(val *SystemConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigResponse(val *SystemConfigResponse) *NullableSystemConfigResponse {
	return &NullableSystemConfigResponse{value: val, isSet: true}
}

func (v NullableSystemConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
