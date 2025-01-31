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

// TaskStatus the model 'TaskStatus'
type TaskStatus string

// List of TaskStatus
const (
	IDLE      TaskStatus = "idle"
	RUNNING   TaskStatus = "running"
	COMPLETED TaskStatus = "completed"
	FAILED    TaskStatus = "failed"
)

// All allowed values of TaskStatus enum
var AllowedTaskStatusEnumValues = []TaskStatus{
	"idle",
	"running",
	"completed",
	"failed",
}

func (v *TaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskStatus(value)
	for _, existing := range AllowedTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskStatus", value)
}

// NewTaskStatusFromValue returns a pointer to a valid TaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskStatusFromValue(v string) (*TaskStatus, error) {
	ev := TaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskStatus: valid values are %v", v, AllowedTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskStatus) IsValid() bool {
	for _, existing := range AllowedTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskStatus value
func (v TaskStatus) Ptr() *TaskStatus {
	return &v
}

type NullableTaskStatus struct {
	value *TaskStatus
	isSet bool
}

func (v NullableTaskStatus) Get() *TaskStatus {
	return v.value
}

func (v *NullableTaskStatus) Set(val *TaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatus(val *TaskStatus) *NullableTaskStatus {
	return &NullableTaskStatus{value: val, isSet: true}
}

func (v NullableTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
