/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the MacAddressInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacAddressInput{}

// MacAddressInput struct for MacAddressInput
type MacAddressInput struct {
	Value        string                   `json:"value"`
	Type         *string                  `json:"type,omitempty"`
	Created      *time.Time               `json:"created,omitempty"`
	Context      []map[string]interface{} `json:"context,omitempty"`
	LastAnalysis map[string]time.Time     `json:"last_analysis,omitempty"`
}

type _MacAddressInput MacAddressInput

// NewMacAddressInput instantiates a new MacAddressInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacAddressInput(value string) *MacAddressInput {
	this := MacAddressInput{}
	this.Value = value
	var type_ string = "mac_address"
	this.Type = &type_
	return &this
}

// NewMacAddressInputWithDefaults instantiates a new MacAddressInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacAddressInputWithDefaults() *MacAddressInput {
	this := MacAddressInput{}
	var type_ string = "mac_address"
	this.Type = &type_
	return &this
}

// GetValue returns the Value field value
func (o *MacAddressInput) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MacAddressInput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MacAddressInput) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MacAddressInput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacAddressInput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MacAddressInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MacAddressInput) SetType(v string) {
	o.Type = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MacAddressInput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacAddressInput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MacAddressInput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *MacAddressInput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *MacAddressInput) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacAddressInput) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MacAddressInput) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *MacAddressInput) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetLastAnalysis returns the LastAnalysis field value if set, zero value otherwise.
func (o *MacAddressInput) GetLastAnalysis() map[string]time.Time {
	if o == nil || IsNil(o.LastAnalysis) {
		var ret map[string]time.Time
		return ret
	}
	return o.LastAnalysis
}

// GetLastAnalysisOk returns a tuple with the LastAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacAddressInput) GetLastAnalysisOk() (map[string]time.Time, bool) {
	if o == nil || IsNil(o.LastAnalysis) {
		return map[string]time.Time{}, false
	}
	return o.LastAnalysis, true
}

// HasLastAnalysis returns a boolean if a field has been set.
func (o *MacAddressInput) HasLastAnalysis() bool {
	if o != nil && !IsNil(o.LastAnalysis) {
		return true
	}

	return false
}

// SetLastAnalysis gets a reference to the given map[string]time.Time and assigns it to the LastAnalysis field.
func (o *MacAddressInput) SetLastAnalysis(v map[string]time.Time) {
	o.LastAnalysis = v
}

func (o MacAddressInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacAddressInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.LastAnalysis) {
		toSerialize["last_analysis"] = o.LastAnalysis
	}
	return toSerialize, nil
}

func (o *MacAddressInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varMacAddressInput := _MacAddressInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMacAddressInput)

	if err != nil {
		return err
	}

	*o = MacAddressInput(varMacAddressInput)

	return err
}

type NullableMacAddressInput struct {
	value *MacAddressInput
	isSet bool
}

func (v NullableMacAddressInput) Get() *MacAddressInput {
	return v.value
}

func (v *NullableMacAddressInput) Set(val *MacAddressInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMacAddressInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMacAddressInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacAddressInput(val *MacAddressInput) *NullableMacAddressInput {
	return &NullableMacAddressInput{value: val, isSet: true}
}

func (v NullableMacAddressInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacAddressInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
