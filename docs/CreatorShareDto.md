# CreatorShareDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Player** | [**PlayerPartialDto**](PlayerPartialDto.md) |  | 
**Quantity** | **int32** |  | 

## Methods

### NewCreatorShareDto

`func NewCreatorShareDto(player PlayerPartialDto, quantity int32, ) *CreatorShareDto`

NewCreatorShareDto instantiates a new CreatorShareDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatorShareDtoWithDefaults

`func NewCreatorShareDtoWithDefaults() *CreatorShareDto`

NewCreatorShareDtoWithDefaults instantiates a new CreatorShareDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayer

`func (o *CreatorShareDto) GetPlayer() PlayerPartialDto`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *CreatorShareDto) GetPlayerOk() (*PlayerPartialDto, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *CreatorShareDto) SetPlayer(v PlayerPartialDto)`

SetPlayer sets Player field to given value.


### GetQuantity

`func (o *CreatorShareDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreatorShareDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreatorShareDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


