# PlayerShareDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | [**CreatorPartialDto**](CreatorPartialDto.md) |  | 
**Quantity** | **int32** |  | 

## Methods

### NewPlayerShareDto

`func NewPlayerShareDto(creator CreatorPartialDto, quantity int32, ) *PlayerShareDto`

NewPlayerShareDto instantiates a new PlayerShareDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerShareDtoWithDefaults

`func NewPlayerShareDtoWithDefaults() *PlayerShareDto`

NewPlayerShareDtoWithDefaults instantiates a new PlayerShareDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *PlayerShareDto) GetCreator() CreatorPartialDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *PlayerShareDto) GetCreatorOk() (*CreatorPartialDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *PlayerShareDto) SetCreator(v CreatorPartialDto)`

SetCreator sets Creator field to given value.


### GetQuantity

`func (o *PlayerShareDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlayerShareDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlayerShareDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


