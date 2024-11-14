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

// checks if the AttackPatternOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackPatternOutput{}

// AttackPatternOutput struct for AttackPatternOutput
type AttackPatternOutput struct {
	Type                    *string                          `json:"type,omitempty"`
	Name                    string                           `json:"name"`
	Description             *string                          `json:"description,omitempty"`
	Context                 []map[string]interface{}         `json:"context,omitempty"`
	Created                 *time.Time                       `json:"created,omitempty"`
	Modified                *time.Time                       `json:"modified,omitempty"`
	Aliases                 []string                         `json:"aliases,omitempty"`
	KillChainPhases         []string                         `json:"kill_chain_phases,omitempty"`
	Id                      string                           `json:"id"`
	Tags                    map[string]TagRelationshipOutput `json:"tags"`
	RootType                string                           `json:"root_type"`
	RelatedObservablesCount int32                            `json:"related_observables_count"`
}

type _AttackPatternOutput AttackPatternOutput

// NewAttackPatternOutput instantiates a new AttackPatternOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackPatternOutput(name string, id string, tags map[string]TagRelationshipOutput, rootType string, relatedObservablesCount int32) *AttackPatternOutput {
	this := AttackPatternOutput{}
	var type_ string = "attack-pattern"
	this.Type = &type_
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Id = id
	this.Tags = tags
	this.RootType = rootType
	this.RelatedObservablesCount = relatedObservablesCount
	return &this
}

// NewAttackPatternOutputWithDefaults instantiates a new AttackPatternOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackPatternOutputWithDefaults() *AttackPatternOutput {
	this := AttackPatternOutput{}
	var type_ string = "attack-pattern"
	this.Type = &type_
	var description string = ""
	this.Description = &description
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AttackPatternOutput) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *AttackPatternOutput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttackPatternOutput) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttackPatternOutput) SetDescription(v string) {
	o.Description = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetContext() []map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetContextOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given []map[string]interface{} and assigns it to the Context field.
func (o *AttackPatternOutput) SetContext(v []map[string]interface{}) {
	o.Context = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AttackPatternOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *AttackPatternOutput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *AttackPatternOutput) SetAliases(v []string) {
	o.Aliases = v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *AttackPatternOutput) GetKillChainPhases() []string {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []string
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetKillChainPhasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *AttackPatternOutput) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []string and assigns it to the KillChainPhases field.
func (o *AttackPatternOutput) SetKillChainPhases(v []string) {
	o.KillChainPhases = v
}

// GetId returns the Id field value
func (o *AttackPatternOutput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttackPatternOutput) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *AttackPatternOutput) GetTags() map[string]TagRelationshipOutput {
	if o == nil {
		var ret map[string]TagRelationshipOutput
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetTagsOk() (map[string]TagRelationshipOutput, bool) {
	if o == nil {
		return map[string]TagRelationshipOutput{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *AttackPatternOutput) SetTags(v map[string]TagRelationshipOutput) {
	o.Tags = v
}

// GetRootType returns the RootType field value
func (o *AttackPatternOutput) GetRootType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootType
}

// GetRootTypeOk returns a tuple with the RootType field value
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetRootTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootType, true
}

// SetRootType sets field value
func (o *AttackPatternOutput) SetRootType(v string) {
	o.RootType = v
}

// GetRelatedObservablesCount returns the RelatedObservablesCount field value
func (o *AttackPatternOutput) GetRelatedObservablesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelatedObservablesCount
}

// GetRelatedObservablesCountOk returns a tuple with the RelatedObservablesCount field value
// and a boolean to check if the value has been set.
func (o *AttackPatternOutput) GetRelatedObservablesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedObservablesCount, true
}

// SetRelatedObservablesCount sets field value
func (o *AttackPatternOutput) SetRelatedObservablesCount(v int32) {
	o.RelatedObservablesCount = v
}

func (o AttackPatternOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackPatternOutput) ToMap() (map[string]interface{}, error) {
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
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["root_type"] = o.RootType
	toSerialize["related_observables_count"] = o.RelatedObservablesCount
	return toSerialize, nil
}

func (o *AttackPatternOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
		"tags",
		"root_type",
		"related_observables_count",
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

	varAttackPatternOutput := _AttackPatternOutput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttackPatternOutput)

	if err != nil {
		return err
	}

	*o = AttackPatternOutput(varAttackPatternOutput)

	return err
}

type NullableAttackPatternOutput struct {
	value *AttackPatternOutput
	isSet bool
}

func (v NullableAttackPatternOutput) Get() *AttackPatternOutput {
	return v.value
}

func (v *NullableAttackPatternOutput) Set(val *AttackPatternOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackPatternOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackPatternOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackPatternOutput(val *AttackPatternOutput) *NullableAttackPatternOutput {
	return &NullableAttackPatternOutput{value: val, isSet: true}
}

func (v NullableAttackPatternOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackPatternOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
