/*
TTX.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gottx

import (
	"encoding/json"
	"fmt"
)

// CreatorOrderBy the model 'CreatorOrderBy'
type CreatorOrderBy string

// List of CreatorOrderBy
const (
	CREATORORDERBY_NAME CreatorOrderBy = "Name"
	CREATORORDERBY_VALUE CreatorOrderBy = "Value"
	CREATORORDERBY_IS_LIVE CreatorOrderBy = "IsLive"
)

// All allowed values of CreatorOrderBy enum
var AllowedCreatorOrderByEnumValues = []CreatorOrderBy{
	"Name",
	"Value",
	"IsLive",
}

func (v *CreatorOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreatorOrderBy(value)
	for _, existing := range AllowedCreatorOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreatorOrderBy", value)
}

// NewCreatorOrderByFromValue returns a pointer to a valid CreatorOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreatorOrderByFromValue(v string) (*CreatorOrderBy, error) {
	ev := CreatorOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreatorOrderBy: valid values are %v", v, AllowedCreatorOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreatorOrderBy) IsValid() bool {
	for _, existing := range AllowedCreatorOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreatorOrderBy value
func (v CreatorOrderBy) Ptr() *CreatorOrderBy {
	return &v
}

type NullableCreatorOrderBy struct {
	value *CreatorOrderBy
	isSet bool
}

func (v NullableCreatorOrderBy) Get() *CreatorOrderBy {
	return v.value
}

func (v *NullableCreatorOrderBy) Set(val *CreatorOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatorOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatorOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatorOrderBy(val *CreatorOrderBy) *NullableCreatorOrderBy {
	return &NullableCreatorOrderBy{value: val, isSet: true}
}

func (v NullableCreatorOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatorOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

