# PlayerPartialDto

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
**Credits** | **int64** |  | 
**Portfolio** | **int64** |  | 
**Value** | **int64** |  | 
**Type** | [**PlayerType**](PlayerType.md) |  | 

## Methods

### NewPlayerPartialDto

`func NewPlayerPartialDto(id int32, createdAt time.Time, updatedAt time.Time, name string, slug string, twitchId string, url string, avatarUrl string, credits int64, portfolio int64, value int64, type_ PlayerType, ) *PlayerPartialDto`

NewPlayerPartialDto instantiates a new PlayerPartialDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerPartialDtoWithDefaults

`func NewPlayerPartialDtoWithDefaults() *PlayerPartialDto`

NewPlayerPartialDtoWithDefaults instantiates a new PlayerPartialDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerPartialDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerPartialDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerPartialDto) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PlayerPartialDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlayerPartialDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlayerPartialDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PlayerPartialDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlayerPartialDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlayerPartialDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *PlayerPartialDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayerPartialDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayerPartialDto) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PlayerPartialDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PlayerPartialDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PlayerPartialDto) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTwitchId

`func (o *PlayerPartialDto) GetTwitchId() string`

GetTwitchId returns the TwitchId field if non-nil, zero value otherwise.

### GetTwitchIdOk

`func (o *PlayerPartialDto) GetTwitchIdOk() (*string, bool)`

GetTwitchIdOk returns a tuple with the TwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchId

`func (o *PlayerPartialDto) SetTwitchId(v string)`

SetTwitchId sets TwitchId field to given value.


### GetUrl

`func (o *PlayerPartialDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlayerPartialDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlayerPartialDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvatarUrl

`func (o *PlayerPartialDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PlayerPartialDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PlayerPartialDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetCredits

`func (o *PlayerPartialDto) GetCredits() int64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *PlayerPartialDto) GetCreditsOk() (*int64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *PlayerPartialDto) SetCredits(v int64)`

SetCredits sets Credits field to given value.


### GetPortfolio

`func (o *PlayerPartialDto) GetPortfolio() int64`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *PlayerPartialDto) GetPortfolioOk() (*int64, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *PlayerPartialDto) SetPortfolio(v int64)`

SetPortfolio sets Portfolio field to given value.


### GetValue

`func (o *PlayerPartialDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlayerPartialDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlayerPartialDto) SetValue(v int64)`

SetValue sets Value field to given value.


### GetType

`func (o *PlayerPartialDto) GetType() PlayerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayerPartialDto) GetTypeOk() (*PlayerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayerPartialDto) SetType(v PlayerType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


