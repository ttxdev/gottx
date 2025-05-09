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

// checks if the CreateTransactionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransactionDto{}

// CreateTransactionDto struct for CreateTransactionDto
type CreateTransactionDto struct {
	Creator string `json:"creator"`
	Action TransactionAction `json:"action"`
	Amount int32 `json:"amount"`
}

type _CreateTransactionDto CreateTransactionDto

// NewCreateTransactionDto instantiates a new CreateTransactionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransactionDto(creator string, action TransactionAction, amount int32) *CreateTransactionDto {
	this := CreateTransactionDto{}
	this.Creator = creator
	this.Action = action
	this.Amount = amount
	return &this
}

// NewCreateTransactionDtoWithDefaults instantiates a new CreateTransactionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransactionDtoWithDefaults() *CreateTransactionDto {
	this := CreateTransactionDto{}
	return &this
}

// GetCreator returns the Creator field value
func (o *CreateTransactionDto) GetCreator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionDto) GetCreatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *CreateTransactionDto) SetCreator(v string) {
	o.Creator = v
}

// GetAction returns the Action field value
func (o *CreateTransactionDto) GetAction() TransactionAction {
	if o == nil {
		var ret TransactionAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionDto) GetActionOk() (*TransactionAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CreateTransactionDto) SetAction(v TransactionAction) {
	o.Action = v
}

// GetAmount returns the Amount field value
func (o *CreateTransactionDto) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionDto) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateTransactionDto) SetAmount(v int32) {
	o.Amount = v
}

func (o CreateTransactionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTransactionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creator"] = o.Creator
	toSerialize["action"] = o.Action
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *CreateTransactionDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"creator",
		"action",
		"amount",
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

	varCreateTransactionDto := _CreateTransactionDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTransactionDto)

	if err != nil {
		return err
	}

	*o = CreateTransactionDto(varCreateTransactionDto)

	return err
}

type NullableCreateTransactionDto struct {
	value *CreateTransactionDto
	isSet bool
}

func (v NullableCreateTransactionDto) Get() *CreateTransactionDto {
	return v.value
}

func (v *NullableCreateTransactionDto) Set(val *CreateTransactionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransactionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransactionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransactionDto(val *CreateTransactionDto) *NullableCreateTransactionDto {
	return &NullableCreateTransactionDto{value: val, isSet: true}
}

func (v NullableCreateTransactionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransactionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


