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

// checks if the ToolInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolInput{}

// ToolInput struct for ToolInput
type ToolInput struct {
	Type            *string                  `json:"type,omitempty"`
	Name            string                   `json:"name"`
	Description     *string                  `json:"description,omitempty"`
	Context         []map[string]interface{} `json:"context,omitempty"`
	Created         *time.Time               `json:"created,omitempty"`
	Modified        *time.Time               `json:"modified,omitempty"`
	Aliases         []string                 `json:"aliases,omitempty"`
	KillChainPhases []string                 `json:"kill_chain_phases,omitempty"`
	ToolVersion     *string                  `json:"tool_version,omitempty"`
}

type _ToolInput ToolInput

// NewToolInput instantiates a new ToolInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolInput(name string) *ToolInput {
	this := ToolInput{}
	var type_ string = "tool"
	this.Type = &type_
	this.Name = name
	var description string = ""
	this.Description = &description
	var toolVersion string = ""
	this.ToolVersion = &toolVersion
	return &this
}

// NewToolInputWithDefaults instantiates a new ToolInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolInputWithDefaults() *ToolInput {
	this := ToolInput{}
	var type_ string = "tool"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	var toolVersion string = ""
	this.ToolVersion = &toolVersion
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ToolInput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ToolInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ToolInput) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *ToolInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolInput) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ToolInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ToolInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ToolInput) SetDescription(v string) {
	o.Description = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ToolInput) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ToolInput) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *ToolInput) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ToolInput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ToolInput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ToolInput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ToolInput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ToolInput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ToolInput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *ToolInput) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *ToolInput) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *ToolInput) SetAliases(v []string) {
	o.Aliases = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *ToolInput) GetKillChainPhases() []string {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []string
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetKillChainPhasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *ToolInput) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []string and assigns it to the KillChainPhases field.
func (o *ToolInput) SetKillChainPhases(v []string) {
	o.KillChainPhases = v
}

// GetToolVersion returns the ToolVersion field value if set, zero value otherwise.
func (o *ToolInput) GetToolVersion() string {
	if o == nil || IsNil(o.ToolVersion) {
		var ret string
		return ret
	}
	return *o.ToolVersion
}

// GetToolVersionOk returns a tuple with the ToolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolInput) GetToolVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ToolVersion) {
		return nil, false
	}
	return o.ToolVersion, true
}

// HasToolVersion returns a boolean if a field has been set.
func (o *ToolInput) HasToolVersion() bool {
	if o != nil && !IsNil(o.ToolVersion) {
		return true
	}

	return false
}

// SetToolVersion gets a reference to the given string and assigns it to the ToolVersion field.
func (o *ToolInput) SetToolVersion(v string) {
	o.ToolVersion = &v
}

func (o ToolInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.KillChainPhases) {
		toSerialize["kill_chain_phases"] = o.KillChainPhases
	}
	if !IsNil(o.ToolVersion) {
		toSerialize["tool_version"] = o.ToolVersion
	}
	return toSerialize, nil
}

func (o *ToolInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varToolInput := _ToolInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolInput)

	if err != nil {
		return err
	}

	*o = ToolInput(varToolInput)

	return err
}

type NullableToolInput struct {
	value *ToolInput
	isSet bool
}

func (v NullableToolInput) Get() *ToolInput {
	return v.value
}

func (v *NullableToolInput) Set(val *ToolInput) {
	v.value = val
	v.isSet = true
}

func (v NullableToolInput) IsSet() bool {
	return v.isSet
}

func (v *NullableToolInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolInput(val *ToolInput) *NullableToolInput {
	return &NullableToolInput{value: val, isSet: true}
}

func (v NullableToolInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
