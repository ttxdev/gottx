# CreatorPartialDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**TwitchId** | **string** |  | 
**Url** | **string** |  | [readonly] 
**AvatarUrl** | **string** |  | 
**Ticker** | **string** |  | 
**Value** | **int64** |  | 
**StreamStatus** | [**StreamStatusDto**](StreamStatusDto.md) |  | 
**History** | [**[]VoteDto**](VoteDto.md) |  | 

## Methods

### NewCreatorPartialDto

`func NewCreatorPartialDto(id int32, createdAt time.Time, updatedAt time.Time, name string, slug string, twitchId string, url string, avatarUrl string, ticker string, value int64, streamStatus StreamStatusDto, history []VoteDto, ) *CreatorPartialDto`

NewCreatorPartialDto instantiates a new CreatorPartialDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatorPartialDtoWithDefaults

`func NewCreatorPartialDtoWithDefaults() *CreatorPartialDto`

NewCreatorPartialDtoWithDefaults instantiates a new CreatorPartialDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatorPartialDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatorPartialDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatorPartialDto) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreatorPartialDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatorPartialDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatorPartialDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreatorPartialDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreatorPartialDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreatorPartialDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *CreatorPartialDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatorPartialDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatorPartialDto) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *CreatorPartialDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreatorPartialDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreatorPartialDto) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTwitchId

`func (o *CreatorPartialDto) GetTwitchId() string`

GetTwitchId returns the TwitchId field if non-nil, zero value otherwise.

### GetTwitchIdOk

`func (o *CreatorPartialDto) GetTwitchIdOk() (*string, bool)`

GetTwitchIdOk returns a tuple with the TwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchId

`func (o *CreatorPartialDto) SetTwitchId(v string)`

SetTwitchId sets TwitchId field to given value.


### GetUrl

`func (o *CreatorPartialDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreatorPartialDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreatorPartialDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvatarUrl

`func (o *CreatorPartialDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *CreatorPartialDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *CreatorPartialDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetTicker

`func (o *CreatorPartialDto) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *CreatorPartialDto) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *CreatorPartialDto) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### GetValue

`func (o *CreatorPartialDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreatorPartialDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreatorPartialDto) SetValue(v int64)`

SetValue sets Value field to given value.


### GetStreamStatus

`func (o *CreatorPartialDto) GetStreamStatus() StreamStatusDto`

GetStreamStatus returns the StreamStatus field if non-nil, zero value otherwise.

### GetStreamStatusOk

`func (o *CreatorPartialDto) GetStreamStatusOk() (*StreamStatusDto, bool)`

GetStreamStatusOk returns a tuple with the StreamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamStatus

`func (o *CreatorPartialDto) SetStreamStatus(v StreamStatusDto)`

SetStreamStatus sets StreamStatus field to given value.


### GetHistory

`func (o *CreatorPartialDto) GetHistory() []VoteDto`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CreatorPartialDto) GetHistoryOk() (*[]VoteDto, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CreatorPartialDto) SetHistory(v []VoteDto)`

SetHistory sets History field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


