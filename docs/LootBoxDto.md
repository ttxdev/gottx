# LootBoxDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**IsOpen** | **bool** |  | 
**Result** | [**CreatorPartialDto**](CreatorPartialDto.md) |  | 
**Player** | [**PlayerPartialDto**](PlayerPartialDto.md) |  | 

## Methods

### NewLootBoxDto

`func NewLootBoxDto(id int32, createdAt time.Time, updatedAt time.Time, isOpen bool, result CreatorPartialDto, player PlayerPartialDto, ) *LootBoxDto`

NewLootBoxDto instantiates a new LootBoxDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLootBoxDtoWithDefaults

`func NewLootBoxDtoWithDefaults() *LootBoxDto`

NewLootBoxDtoWithDefaults instantiates a new LootBoxDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LootBoxDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LootBoxDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LootBoxDto) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *LootBoxDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LootBoxDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LootBoxDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LootBoxDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LootBoxDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LootBoxDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsOpen

`func (o *LootBoxDto) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *LootBoxDto) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *LootBoxDto) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.


### GetResult

`func (o *LootBoxDto) GetResult() CreatorPartialDto`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LootBoxDto) GetResultOk() (*CreatorPartialDto, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LootBoxDto) SetResult(v CreatorPartialDto)`

SetResult sets Result field to given value.


### GetPlayer

`func (o *LootBoxDto) GetPlayer() PlayerPartialDto`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *LootBoxDto) GetPlayerOk() (*PlayerPartialDto, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *LootBoxDto) SetPlayer(v PlayerPartialDto)`

SetPlayer sets Player field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


