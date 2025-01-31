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

// checks if the TaskSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskSearchRequest{}

// TaskSearchRequest struct for TaskSearchRequest
type TaskSearchRequest struct {
	Query                map[string]QueryValue  `json:"query,omitempty"`
	Type                 *TaskSearchRequestType `json:"type,omitempty"`
	Sorting              [][]interface{}        `json:"sorting,omitempty"`
	Count                *int32                 `json:"count,omitempty"`
	Page                 *int32                 `json:"page,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskSearchRequest TaskSearchRequest

// NewTaskSearchRequest instantiates a new TaskSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskSearchRequest() *TaskSearchRequest {
	this := TaskSearchRequest{}
	var count int32 = 100
	this.Count = &count
	var page int32 = 0
	this.Page = &page
	return &this
}

// NewTaskSearchRequestWithDefaults instantiates a new TaskSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSearchRequestWithDefaults() *TaskSearchRequest {
	this := TaskSearchRequest{}
	var count int32 = 100
	this.Count = &count
	var page int32 = 0
	this.Page = &page
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TaskSearchRequest) GetQuery() map[string]QueryValue {
	if o == nil || IsNil(o.Query) {
		var ret map[string]QueryValue
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSearchRequest) GetQueryOk() (map[string]QueryValue, bool) {
	if o == nil || IsNil(o.Query) {
		return map[string]QueryValue{}, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TaskSearchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]QueryValue and assigns it to the Query field.
func (o *TaskSearchRequest) SetQuery(v map[string]QueryValue) {
	o.Query = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskSearchRequest) GetType() TaskSearchRequestType {
	if o == nil || IsNil(o.Type) {
		var ret TaskSearchRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSearchRequest) GetTypeOk() (*TaskSearchRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskSearchRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskSearchRequestType and assigns it to the Type field.
func (o *TaskSearchRequest) SetType(v TaskSearchRequestType) {
	o.Type = &v
}

// GetSorting returns the Sorting field value if set, zero value otherwise.
func (o *TaskSearchRequest) GetSorting() [][]interface{} {
	if o == nil || IsNil(o.Sorting) {
		var ret [][]interface{}
		return ret
	}
	return o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSearchRequest) GetSortingOk() ([][]interface{}, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *TaskSearchRequest) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given [][]interface{} and assigns it to the Sorting field.
func (o *TaskSearchRequest) SetSorting(v [][]interface{}) {
	o.Sorting = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TaskSearchRequest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSearchRequest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TaskSearchRequest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TaskSearchRequest) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *TaskSearchRequest) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSearchRequest) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *TaskSearchRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *TaskSearchRequest) SetPage(v int32) {
	o.Page = &v
}

func (o TaskSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Sorting) {
		toSerialize["sorting"] = o.Sorting
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskSearchRequest) UnmarshalJSON(data []byte) (err error) {
	varTaskSearchRequest := _TaskSearchRequest{}

	err = json.Unmarshal(data, &varTaskSearchRequest)

	if err != nil {
		return err
	}

	*o = TaskSearchRequest(varTaskSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		delete(additionalProperties, "type")
		delete(additionalProperties, "sorting")
		delete(additionalProperties, "count")
		delete(additionalProperties, "page")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskSearchRequest struct {
	value *TaskSearchRequest
	isSet bool
}

func (v NullableTaskSearchRequest) Get() *TaskSearchRequest {
	return v.value
}

func (v *NullableTaskSearchRequest) Set(val *TaskSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSearchRequest(val *TaskSearchRequest) *NullableTaskSearchRequest {
	return &NullableTaskSearchRequest{value: val, isSet: true}
}

func (v NullableTaskSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
