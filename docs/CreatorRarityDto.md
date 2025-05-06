# CreatorRarityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | [**CreatorPartialDto**](CreatorPartialDto.md) |  | 
**Rarity** | [**Rarity**](Rarity.md) |  | 

## Methods

### NewCreatorRarityDto

`func NewCreatorRarityDto(creator CreatorPartialDto, rarity Rarity, ) *CreatorRarityDto`

NewCreatorRarityDto instantiates a new CreatorRarityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatorRarityDtoWithDefaults

`func NewCreatorRarityDtoWithDefaults() *CreatorRarityDto`

NewCreatorRarityDtoWithDefaults instantiates a new CreatorRarityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *CreatorRarityDto) GetCreator() CreatorPartialDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreatorRarityDto) GetCreatorOk() (*CreatorPartialDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreatorRarityDto) SetCreator(v CreatorPartialDto)`

SetCreator sets Creator field to given value.


### GetRarity

`func (o *CreatorRarityDto) GetRarity() Rarity`

GetRarity returns the Rarity field if non-nil, zero value otherwise.

### GetRarityOk

`func (o *CreatorRarityDto) GetRarityOk() (*Rarity, bool)`

GetRarityOk returns a tuple with the Rarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarity

`func (o *CreatorRarityDto) SetRarity(v Rarity)`

SetRarity sets Rarity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


