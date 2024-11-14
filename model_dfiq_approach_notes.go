/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"encoding/json"
)

// checks if the DFIQApproachNotes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DFIQApproachNotes{}

// DFIQApproachNotes struct for DFIQApproachNotes
type DFIQApproachNotes struct {
	Covered    []string `json:"covered,omitempty"`
	NotCovered []string `json:"not_covered,omitempty"`
}

// NewDFIQApproachNotes instantiates a new DFIQApproachNotes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDFIQApproachNotes() *DFIQApproachNotes {
	this := DFIQApproachNotes{}
	return &this
}

// NewDFIQApproachNotesWithDefaults instantiates a new DFIQApproachNotes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDFIQApproachNotesWithDefaults() *DFIQApproachNotes {
	this := DFIQApproachNotes{}
	return &this
}

// GetCovered returns the Covered field value if set, zero value otherwise.
func (o *DFIQApproachNotes) GetCovered() []string {
	if o == nil || IsNil(o.Covered) {
		var ret []string
		return ret
	}
	return o.Covered
}

// GetCoveredOk returns a tuple with the Covered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DFIQApproachNotes) GetCoveredOk() ([]string, bool) {
	if o == nil || IsNil(o.Covered) {
		return nil, false
	}
	return o.Covered, true
}

// HasCovered returns a boolean if a field has been set.
func (o *DFIQApproachNotes) HasCovered() bool {
	if o != nil && !IsNil(o.Covered) {
		return true
	}

	return false
}

// SetCovered gets a reference to the given []string and assigns it to the Covered field.
func (o *DFIQApproachNotes) SetCovered(v []string) {
	o.Covered = v
}

// GetNotCovered returns the NotCovered field value if set, zero value otherwise.
func (o *DFIQApproachNotes) GetNotCovered() []string {
	if o == nil || IsNil(o.NotCovered) {
		var ret []string
		return ret
	}
	return o.NotCovered
}

// GetNotCoveredOk returns a tuple with the NotCovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DFIQApproachNotes) GetNotCoveredOk() ([]string, bool) {
	if o == nil || IsNil(o.NotCovered) {
		return nil, false
	}
	return o.NotCovered, true
}

// HasNotCovered returns a boolean if a field has been set.
func (o *DFIQApproachNotes) HasNotCovered() bool {
	if o != nil && !IsNil(o.NotCovered) {
		return true
	}

	return false
}

// SetNotCovered gets a reference to the given []string and assigns it to the NotCovered field.
func (o *DFIQApproachNotes) SetNotCovered(v []string) {
	o.NotCovered = v
}

func (o DFIQApproachNotes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DFIQApproachNotes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Covered) {
		toSerialize["covered"] = o.Covered
	}
	if !IsNil(o.NotCovered) {
		toSerialize["not_covered"] = o.NotCovered
	}
	return toSerialize, nil
}

type NullableDFIQApproachNotes struct {
	value *DFIQApproachNotes
	isSet bool
}

func (v NullableDFIQApproachNotes) Get() *DFIQApproachNotes {
	return v.value
}

func (v *NullableDFIQApproachNotes) Set(val *DFIQApproachNotes) {
	v.value = val
	v.isSet = true
}

func (v NullableDFIQApproachNotes) IsSet() bool {
	return v.isSet
}

func (v *NullableDFIQApproachNotes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDFIQApproachNotes(val *DFIQApproachNotes) *NullableDFIQApproachNotes {
	return &NullableDFIQApproachNotes{value: val, isSet: true}
}

func (v NullableDFIQApproachNotes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDFIQApproachNotes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
