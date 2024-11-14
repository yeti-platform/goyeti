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

// checks if the NewBulkObservableAddRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewBulkObservableAddRequest{}

// NewBulkObservableAddRequest struct for NewBulkObservableAddRequest
type NewBulkObservableAddRequest struct {
	Observables          []NewObservableRequest `json:"observables"`
	AdditionalProperties map[string]interface{}
}

type _NewBulkObservableAddRequest NewBulkObservableAddRequest

// NewNewBulkObservableAddRequest instantiates a new NewBulkObservableAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewBulkObservableAddRequest(observables []NewObservableRequest) *NewBulkObservableAddRequest {
	this := NewBulkObservableAddRequest{}
	this.Observables = observables
	return &this
}

// NewNewBulkObservableAddRequestWithDefaults instantiates a new NewBulkObservableAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewBulkObservableAddRequestWithDefaults() *NewBulkObservableAddRequest {
	this := NewBulkObservableAddRequest{}
	return &this
}

// GetObservables returns the Observables field value
func (o *NewBulkObservableAddRequest) GetObservables() []NewObservableRequest {
	if o == nil {
		var ret []NewObservableRequest
		return ret
	}

	return o.Observables
}

// GetObservablesOk returns a tuple with the Observables field value
// and a boolean to check if the value has been set.
func (o *NewBulkObservableAddRequest) GetObservablesOk() ([]NewObservableRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Observables, true
}

// SetObservables sets field value
func (o *NewBulkObservableAddRequest) SetObservables(v []NewObservableRequest) {
	o.Observables = v
}

func (o NewBulkObservableAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewBulkObservableAddRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["observables"] = o.Observables

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewBulkObservableAddRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"observables",
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

	varNewBulkObservableAddRequest := _NewBulkObservableAddRequest{}

	err = json.Unmarshal(data, &varNewBulkObservableAddRequest)

	if err != nil {
		return err
	}

	*o = NewBulkObservableAddRequest(varNewBulkObservableAddRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "observables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewBulkObservableAddRequest struct {
	value *NewBulkObservableAddRequest
	isSet bool
}

func (v NullableNewBulkObservableAddRequest) Get() *NewBulkObservableAddRequest {
	return v.value
}

func (v *NullableNewBulkObservableAddRequest) Set(val *NewBulkObservableAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBulkObservableAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBulkObservableAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBulkObservableAddRequest(val *NewBulkObservableAddRequest) *NullableNewBulkObservableAddRequest {
	return &NullableNewBulkObservableAddRequest{value: val, isSet: true}
}

func (v NullableNewBulkObservableAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBulkObservableAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
