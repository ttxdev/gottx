/*
TTX.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gottx

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreatorRarityDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatorRarityDto{}

// CreatorRarityDto struct for CreatorRarityDto
type CreatorRarityDto struct {
	Creator CreatorPartialDto `json:"creator"`
	Rarity Rarity `json:"rarity"`
}

type _CreatorRarityDto CreatorRarityDto

// NewCreatorRarityDto instantiates a new CreatorRarityDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatorRarityDto(creator CreatorPartialDto, rarity Rarity) *CreatorRarityDto {
	this := CreatorRarityDto{}
	this.Creator = creator
	this.Rarity = rarity
	return &this
}

// NewCreatorRarityDtoWithDefaults instantiates a new CreatorRarityDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatorRarityDtoWithDefaults() *CreatorRarityDto {
	this := CreatorRarityDto{}
	return &this
}

// GetCreator returns the Creator field value
func (o *CreatorRarityDto) GetCreator() CreatorPartialDto {
	if o == nil {
		var ret CreatorPartialDto
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *CreatorRarityDto) GetCreatorOk() (*CreatorPartialDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *CreatorRarityDto) SetCreator(v CreatorPartialDto) {
	o.Creator = v
}

// GetRarity returns the Rarity field value
func (o *CreatorRarityDto) GetRarity() Rarity {
	if o == nil {
		var ret Rarity
		return ret
	}

	return o.Rarity
}

// GetRarityOk returns a tuple with the Rarity field value
// and a boolean to check if the value has been set.
func (o *CreatorRarityDto) GetRarityOk() (*Rarity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rarity, true
}

// SetRarity sets field value
func (o *CreatorRarityDto) SetRarity(v Rarity) {
	o.Rarity = v
}

func (o CreatorRarityDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatorRarityDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creator"] = o.Creator
	toSerialize["rarity"] = o.Rarity
	return toSerialize, nil
}

func (o *CreatorRarityDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"creator",
		"rarity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreatorRarityDto := _CreatorRarityDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatorRarityDto)

	if err != nil {
		return err
	}

	*o = CreatorRarityDto(varCreatorRarityDto)

	return err
}

type NullableCreatorRarityDto struct {
	value *CreatorRarityDto
	isSet bool
}

func (v NullableCreatorRarityDto) Get() *CreatorRarityDto {
	return v.value
}

func (v *NullableCreatorRarityDto) Set(val *CreatorRarityDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatorRarityDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatorRarityDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatorRarityDto(val *CreatorRarityDto) *NullableCreatorRarityDto {
	return &NullableCreatorRarityDto{value: val, isSet: true}
}

func (v NullableCreatorRarityDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatorRarityDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


