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

// ResponsePatchApiV2EntitiesEntityIdPatch struct for ResponsePatchApiV2EntitiesEntityIdPatch
type ResponsePatchApiV2EntitiesEntityIdPatch struct {
	AttackPatternOutput  *AttackPatternOutput
	CampaignOutput       *CampaignOutput
	CompanyOutput        *CompanyOutput
	CourseOfActionOutput *CourseOfActionOutput
	IdentityOutput       *IdentityOutput
	IntrusionSetOutput   *IntrusionSetOutput
	InvestigationOutput  *InvestigationOutput
	MalwareOutput        *MalwareOutput
	NoteOutput           *NoteOutput
	PhoneOutput          *PhoneOutput
	ThreatActorOutput    *ThreatActorOutput
	ToolOutput           *ToolOutput
	VulnerabilityOutput  *VulnerabilityOutput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResponsePatchApiV2EntitiesEntityIdPatch) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AttackPatternOutput
	err = json.Unmarshal(data, &dst.AttackPatternOutput)
	if err == nil {
		jsonAttackPatternOutput, _ := json.Marshal(dst.AttackPatternOutput)
		if string(jsonAttackPatternOutput) == "{}" { // empty struct
			dst.AttackPatternOutput = nil
		} else {
			return nil // data stored in dst.AttackPatternOutput, return on the first match
		}
	} else {
		dst.AttackPatternOutput = nil
	}

	// try to unmarshal JSON data into CampaignOutput
	err = json.Unmarshal(data, &dst.CampaignOutput)
	if err == nil {
		jsonCampaignOutput, _ := json.Marshal(dst.CampaignOutput)
		if string(jsonCampaignOutput) == "{}" { // empty struct
			dst.CampaignOutput = nil
		} else {
			return nil // data stored in dst.CampaignOutput, return on the first match
		}
	} else {
		dst.CampaignOutput = nil
	}

	// try to unmarshal JSON data into CompanyOutput
	err = json.Unmarshal(data, &dst.CompanyOutput)
	if err == nil {
		jsonCompanyOutput, _ := json.Marshal(dst.CompanyOutput)
		if string(jsonCompanyOutput) == "{}" { // empty struct
			dst.CompanyOutput = nil
		} else {
			return nil // data stored in dst.CompanyOutput, return on the first match
		}
	} else {
		dst.CompanyOutput = nil
	}

	// try to unmarshal JSON data into CourseOfActionOutput
	err = json.Unmarshal(data, &dst.CourseOfActionOutput)
	if err == nil {
		jsonCourseOfActionOutput, _ := json.Marshal(dst.CourseOfActionOutput)
		if string(jsonCourseOfActionOutput) == "{}" { // empty struct
			dst.CourseOfActionOutput = nil
		} else {
			return nil // data stored in dst.CourseOfActionOutput, return on the first match
		}
	} else {
		dst.CourseOfActionOutput = nil
	}

	// try to unmarshal JSON data into IdentityOutput
	err = json.Unmarshal(data, &dst.IdentityOutput)
	if err == nil {
		jsonIdentityOutput, _ := json.Marshal(dst.IdentityOutput)
		if string(jsonIdentityOutput) == "{}" { // empty struct
			dst.IdentityOutput = nil
		} else {
			return nil // data stored in dst.IdentityOutput, return on the first match
		}
	} else {
		dst.IdentityOutput = nil
	}

	// try to unmarshal JSON data into IntrusionSetOutput
	err = json.Unmarshal(data, &dst.IntrusionSetOutput)
	if err == nil {
		jsonIntrusionSetOutput, _ := json.Marshal(dst.IntrusionSetOutput)
		if string(jsonIntrusionSetOutput) == "{}" { // empty struct
			dst.IntrusionSetOutput = nil
		} else {
			return nil // data stored in dst.IntrusionSetOutput, return on the first match
		}
	} else {
		dst.IntrusionSetOutput = nil
	}

	// try to unmarshal JSON data into InvestigationOutput
	err = json.Unmarshal(data, &dst.InvestigationOutput)
	if err == nil {
		jsonInvestigationOutput, _ := json.Marshal(dst.InvestigationOutput)
		if string(jsonInvestigationOutput) == "{}" { // empty struct
			dst.InvestigationOutput = nil
		} else {
			return nil // data stored in dst.InvestigationOutput, return on the first match
		}
	} else {
		dst.InvestigationOutput = nil
	}

	// try to unmarshal JSON data into MalwareOutput
	err = json.Unmarshal(data, &dst.MalwareOutput)
	if err == nil {
		jsonMalwareOutput, _ := json.Marshal(dst.MalwareOutput)
		if string(jsonMalwareOutput) == "{}" { // empty struct
			dst.MalwareOutput = nil
		} else {
			return nil // data stored in dst.MalwareOutput, return on the first match
		}
	} else {
		dst.MalwareOutput = nil
	}

	// try to unmarshal JSON data into NoteOutput
	err = json.Unmarshal(data, &dst.NoteOutput)
	if err == nil {
		jsonNoteOutput, _ := json.Marshal(dst.NoteOutput)
		if string(jsonNoteOutput) == "{}" { // empty struct
			dst.NoteOutput = nil
		} else {
			return nil // data stored in dst.NoteOutput, return on the first match
		}
	} else {
		dst.NoteOutput = nil
	}

	// try to unmarshal JSON data into PhoneOutput
	err = json.Unmarshal(data, &dst.PhoneOutput)
	if err == nil {
		jsonPhoneOutput, _ := json.Marshal(dst.PhoneOutput)
		if string(jsonPhoneOutput) == "{}" { // empty struct
			dst.PhoneOutput = nil
		} else {
			return nil // data stored in dst.PhoneOutput, return on the first match
		}
	} else {
		dst.PhoneOutput = nil
	}

	// try to unmarshal JSON data into ThreatActorOutput
	err = json.Unmarshal(data, &dst.ThreatActorOutput)
	if err == nil {
		jsonThreatActorOutput, _ := json.Marshal(dst.ThreatActorOutput)
		if string(jsonThreatActorOutput) == "{}" { // empty struct
			dst.ThreatActorOutput = nil
		} else {
			return nil // data stored in dst.ThreatActorOutput, return on the first match
		}
	} else {
		dst.ThreatActorOutput = nil
	}

	// try to unmarshal JSON data into ToolOutput
	err = json.Unmarshal(data, &dst.ToolOutput)
	if err == nil {
		jsonToolOutput, _ := json.Marshal(dst.ToolOutput)
		if string(jsonToolOutput) == "{}" { // empty struct
			dst.ToolOutput = nil
		} else {
			return nil // data stored in dst.ToolOutput, return on the first match
		}
	} else {
		dst.ToolOutput = nil
	}

	// try to unmarshal JSON data into VulnerabilityOutput
	err = json.Unmarshal(data, &dst.VulnerabilityOutput)
	if err == nil {
		jsonVulnerabilityOutput, _ := json.Marshal(dst.VulnerabilityOutput)
		if string(jsonVulnerabilityOutput) == "{}" { // empty struct
			dst.VulnerabilityOutput = nil
		} else {
			return nil // data stored in dst.VulnerabilityOutput, return on the first match
		}
	} else {
		dst.VulnerabilityOutput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResponsePatchApiV2EntitiesEntityIdPatch)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResponsePatchApiV2EntitiesEntityIdPatch) MarshalJSON() ([]byte, error) {
	if src.AttackPatternOutput != nil {
		return json.Marshal(&src.AttackPatternOutput)
	}

	if src.CampaignOutput != nil {
		return json.Marshal(&src.CampaignOutput)
	}

	if src.CompanyOutput != nil {
		return json.Marshal(&src.CompanyOutput)
	}

	if src.CourseOfActionOutput != nil {
		return json.Marshal(&src.CourseOfActionOutput)
	}

	if src.IdentityOutput != nil {
		return json.Marshal(&src.IdentityOutput)
	}

	if src.IntrusionSetOutput != nil {
		return json.Marshal(&src.IntrusionSetOutput)
	}

	if src.InvestigationOutput != nil {
		return json.Marshal(&src.InvestigationOutput)
	}

	if src.MalwareOutput != nil {
		return json.Marshal(&src.MalwareOutput)
	}

	if src.NoteOutput != nil {
		return json.Marshal(&src.NoteOutput)
	}

	if src.PhoneOutput != nil {
		return json.Marshal(&src.PhoneOutput)
	}

	if src.ThreatActorOutput != nil {
		return json.Marshal(&src.ThreatActorOutput)
	}

	if src.ToolOutput != nil {
		return json.Marshal(&src.ToolOutput)
	}

	if src.VulnerabilityOutput != nil {
		return json.Marshal(&src.VulnerabilityOutput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResponsePatchApiV2EntitiesEntityIdPatch struct {
	value *ResponsePatchApiV2EntitiesEntityIdPatch
	isSet bool
}

func (v NullableResponsePatchApiV2EntitiesEntityIdPatch) Get() *ResponsePatchApiV2EntitiesEntityIdPatch {
	return v.value
}

func (v *NullableResponsePatchApiV2EntitiesEntityIdPatch) Set(val *ResponsePatchApiV2EntitiesEntityIdPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePatchApiV2EntitiesEntityIdPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePatchApiV2EntitiesEntityIdPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePatchApiV2EntitiesEntityIdPatch(val *ResponsePatchApiV2EntitiesEntityIdPatch) *NullableResponsePatchApiV2EntitiesEntityIdPatch {
	return &NullableResponsePatchApiV2EntitiesEntityIdPatch{value: val, isSet: true}
}

func (v NullableResponsePatchApiV2EntitiesEntityIdPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePatchApiV2EntitiesEntityIdPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
