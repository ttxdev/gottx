# LootBoxResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LootboxId** | [**ModelId**](ModelId.md) |  | 
**Player** | [**PlayerPartialDto**](PlayerPartialDto.md) |  | 
**Result** | [**CreatorRarityDto**](CreatorRarityDto.md) |  | 
**Rarities** | [**[]CreatorRarityDto**](CreatorRarityDto.md) |  | 

## Methods

### NewLootBoxResultDto

`func NewLootBoxResultDto(lootboxId ModelId, player PlayerPartialDto, result CreatorRarityDto, rarities []CreatorRarityDto, ) *LootBoxResultDto`

NewLootBoxResultDto instantiates a new LootBoxResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLootBoxResultDtoWithDefaults

`func NewLootBoxResultDtoWithDefaults() *LootBoxResultDto`

NewLootBoxResultDtoWithDefaults instantiates a new LootBoxResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLootboxId

`func (o *LootBoxResultDto) GetLootboxId() ModelId`

GetLootboxId returns the LootboxId field if non-nil, zero value otherwise.

### GetLootboxIdOk

`func (o *LootBoxResultDto) GetLootboxIdOk() (*ModelId, bool)`

GetLootboxIdOk returns a tuple with the LootboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLootboxId

`func (o *LootBoxResultDto) SetLootboxId(v ModelId)`

SetLootboxId sets LootboxId field to given value.


### GetPlayer

`func (o *LootBoxResultDto) GetPlayer() PlayerPartialDto`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *LootBoxResultDto) GetPlayerOk() (*PlayerPartialDto, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *LootBoxResultDto) SetPlayer(v PlayerPartialDto)`

SetPlayer sets Player field to given value.


### GetResult

`func (o *LootBoxResultDto) GetResult() CreatorRarityDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LootBoxResultDto) GetResultOk() (*CreatorRarityDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LootBoxResultDto) SetResult(v CreatorRarityDto)`

SetResult sets Result field to given value.


### GetRarities

`func (o *LootBoxResultDto) GetRarities() []CreatorRarityDto`

GetRarities returns the Rarities field if non-nil, zero value otherwise.

### GetRaritiesOk

`func (o *LootBoxResultDto) GetRaritiesOk() (*[]CreatorRarityDto, bool)`

GetRaritiesOk returns a tuple with the Rarities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarities

`func (o *LootBoxResultDto) SetRarities(v []CreatorRarityDto)`

SetRarities sets Rarities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


