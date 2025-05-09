/*
TTX.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gottx

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CreatorPartialDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatorPartialDto{}

// CreatorPartialDto struct for CreatorPartialDto
type CreatorPartialDto struct {
	Id int32 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	TwitchId string `json:"twitch_id"`
	Url string `json:"url"`
	AvatarUrl string `json:"avatar_url"`
	Ticker string `json:"ticker"`
	Value int64 `json:"value"`
	StreamStatus StreamStatusDto `json:"stream_status"`
	History []VoteDto `json:"history"`
}

type _CreatorPartialDto CreatorPartialDto

// NewCreatorPartialDto instantiates a new CreatorPartialDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatorPartialDto(id int32, createdAt time.Time, updatedAt time.Time, name string, slug string, twitchId string, url string, avatarUrl string, ticker string, value int64, streamStatus StreamStatusDto, history []VoteDto) *CreatorPartialDto {
	this := CreatorPartialDto{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Slug = slug
	this.TwitchId = twitchId
	this.Url = url
	this.AvatarUrl = avatarUrl
	this.Ticker = ticker
	this.Value = value
	this.StreamStatus = streamStatus
	this.History = history
	return &this
}

// NewCreatorPartialDtoWithDefaults instantiates a new CreatorPartialDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatorPartialDtoWithDefaults() *CreatorPartialDto {
	this := CreatorPartialDto{}
	return &this
}

// GetId returns the Id field value
func (o *CreatorPartialDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreatorPartialDto) SetId(v int32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreatorPartialDto) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreatorPartialDto) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CreatorPartialDto) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CreatorPartialDto) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *CreatorPartialDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatorPartialDto) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *CreatorPartialDto) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *CreatorPartialDto) SetSlug(v string) {
	o.Slug = v
}

// GetTwitchId returns the TwitchId field value
func (o *CreatorPartialDto) GetTwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TwitchId
}

// GetTwitchIdOk returns a tuple with the TwitchId field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetTwitchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TwitchId, true
}

// SetTwitchId sets field value
func (o *CreatorPartialDto) SetTwitchId(v string) {
	o.TwitchId = v
}

// GetUrl returns the Url field value
func (o *CreatorPartialDto) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreatorPartialDto) SetUrl(v string) {
	o.Url = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *CreatorPartialDto) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *CreatorPartialDto) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetTicker returns the Ticker field value
func (o *CreatorPartialDto) GetTicker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetTickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticker, true
}

// SetTicker sets field value
func (o *CreatorPartialDto) SetTicker(v string) {
	o.Ticker = v
}

// GetValue returns the Value field value
func (o *CreatorPartialDto) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreatorPartialDto) SetValue(v int64) {
	o.Value = v
}

// GetStreamStatus returns the StreamStatus field value
func (o *CreatorPartialDto) GetStreamStatus() StreamStatusDto {
	if o == nil {
		var ret StreamStatusDto
		return ret
	}

	return o.StreamStatus
}

// GetStreamStatusOk returns a tuple with the StreamStatus field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetStreamStatusOk() (*StreamStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamStatus, true
}

// SetStreamStatus sets field value
func (o *CreatorPartialDto) SetStreamStatus(v StreamStatusDto) {
	o.StreamStatus = v
}

// GetHistory returns the History field value
func (o *CreatorPartialDto) GetHistory() []VoteDto {
	if o == nil {
		var ret []VoteDto
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *CreatorPartialDto) GetHistoryOk() ([]VoteDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.History, true
}

// SetHistory sets field value
func (o *CreatorPartialDto) SetHistory(v []VoteDto) {
	o.History = v
}

func (o CreatorPartialDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatorPartialDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["twitch_id"] = o.TwitchId
	toSerialize["url"] = o.Url
	toSerialize["avatar_url"] = o.AvatarUrl
	toSerialize["ticker"] = o.Ticker
	toSerialize["value"] = o.Value
	toSerialize["stream_status"] = o.StreamStatus
	toSerialize["history"] = o.History
	return toSerialize, nil
}

func (o *CreatorPartialDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"slug",
		"twitch_id",
		"url",
		"avatar_url",
		"ticker",
		"value",
		"stream_status",
		"history",
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

	varCreatorPartialDto := _CreatorPartialDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatorPartialDto)

	if err != nil {
		return err
	}

	*o = CreatorPartialDto(varCreatorPartialDto)

	return err
}

type NullableCreatorPartialDto struct {
	value *CreatorPartialDto
	isSet bool
}

func (v NullableCreatorPartialDto) Get() *CreatorPartialDto {
	return v.value
}

func (v *NullableCreatorPartialDto) Set(val *CreatorPartialDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatorPartialDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatorPartialDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatorPartialDto(val *CreatorPartialDto) *NullableCreatorPartialDto {
	return &NullableCreatorPartialDto{value: val, isSet: true}
}

func (v NullableCreatorPartialDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatorPartialDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


