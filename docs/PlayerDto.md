# PlayerDto

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
**Transactions** | [**[]PlayerTransactionDto**](PlayerTransactionDto.md) |  | 
**LootBoxes** | [**[]LootBoxDto**](LootBoxDto.md) |  | 
**Shares** | [**[]PlayerShareDto**](PlayerShareDto.md) |  | 
**History** | [**[]PortfolioSnapshotDto**](PortfolioSnapshotDto.md) |  | 

## Methods

### NewPlayerDto

`func NewPlayerDto(id int32, createdAt time.Time, updatedAt time.Time, name string, slug string, twitchId string, url string, avatarUrl string, credits int64, portfolio int64, value int64, type_ PlayerType, transactions []PlayerTransactionDto, lootBoxes []LootBoxDto, shares []PlayerShareDto, history []PortfolioSnapshotDto, ) *PlayerDto`

NewPlayerDto instantiates a new PlayerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerDtoWithDefaults

`func NewPlayerDtoWithDefaults() *PlayerDto`

NewPlayerDtoWithDefaults instantiates a new PlayerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerDto) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PlayerDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlayerDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlayerDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PlayerDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlayerDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlayerDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *PlayerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayerDto) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PlayerDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PlayerDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PlayerDto) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTwitchId

`func (o *PlayerDto) GetTwitchId() string`

GetTwitchId returns the TwitchId field if non-nil, zero value otherwise.

### GetTwitchIdOk

`func (o *PlayerDto) GetTwitchIdOk() (*string, bool)`

GetTwitchIdOk returns a tuple with the TwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchId

`func (o *PlayerDto) SetTwitchId(v string)`

SetTwitchId sets TwitchId field to given value.


### GetUrl

`func (o *PlayerDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlayerDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlayerDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvatarUrl

`func (o *PlayerDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PlayerDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PlayerDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetCredits

`func (o *PlayerDto) GetCredits() int64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *PlayerDto) GetCreditsOk() (*int64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *PlayerDto) SetCredits(v int64)`

SetCredits sets Credits field to given value.


### GetPortfolio

`func (o *PlayerDto) GetPortfolio() int64`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *PlayerDto) GetPortfolioOk() (*int64, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *PlayerDto) SetPortfolio(v int64)`

SetPortfolio sets Portfolio field to given value.


### GetValue

`func (o *PlayerDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlayerDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlayerDto) SetValue(v int64)`

SetValue sets Value field to given value.


### GetType

`func (o *PlayerDto) GetType() PlayerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayerDto) GetTypeOk() (*PlayerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayerDto) SetType(v PlayerType)`

SetType sets Type field to given value.


### GetTransactions

`func (o *PlayerDto) GetTransactions() []PlayerTransactionDto`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PlayerDto) GetTransactionsOk() (*[]PlayerTransactionDto, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PlayerDto) SetTransactions(v []PlayerTransactionDto)`

SetTransactions sets Transactions field to given value.


### GetLootBoxes

`func (o *PlayerDto) GetLootBoxes() []LootBoxDto`

GetLootBoxes returns the LootBoxes field if non-nil, zero value otherwise.

### GetLootBoxesOk

`func (o *PlayerDto) GetLootBoxesOk() (*[]LootBoxDto, bool)`

GetLootBoxesOk returns a tuple with the LootBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLootBoxes

`func (o *PlayerDto) SetLootBoxes(v []LootBoxDto)`

SetLootBoxes sets LootBoxes field to given value.


### GetShares

`func (o *PlayerDto) GetShares() []PlayerShareDto`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *PlayerDto) GetSharesOk() (*[]PlayerShareDto, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *PlayerDto) SetShares(v []PlayerShareDto)`

SetShares sets Shares field to given value.


### GetHistory

`func (o *PlayerDto) GetHistory() []PortfolioSnapshotDto`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *PlayerDto) GetHistoryOk() (*[]PortfolioSnapshotDto, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *PlayerDto) SetHistory(v []PortfolioSnapshotDto)`

SetHistory sets History field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


