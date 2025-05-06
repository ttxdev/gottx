# PlayerTransactionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Quantity** | **int32** |  | 
**Value** | **int64** |  | 
**Action** | [**TransactionAction**](TransactionAction.md) |  | 
**CreatorId** | **int32** |  | 
**PlayerId** | **int32** |  | 
**Creator** | [**CreatorPartialDto**](CreatorPartialDto.md) |  | 

## Methods

### NewPlayerTransactionDto

`func NewPlayerTransactionDto(id int32, createdAt time.Time, updatedAt time.Time, quantity int32, value int64, action TransactionAction, creatorId int32, playerId int32, creator CreatorPartialDto, ) *PlayerTransactionDto`

NewPlayerTransactionDto instantiates a new PlayerTransactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerTransactionDtoWithDefaults

`func NewPlayerTransactionDtoWithDefaults() *PlayerTransactionDto`

NewPlayerTransactionDtoWithDefaults instantiates a new PlayerTransactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerTransactionDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerTransactionDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerTransactionDto) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PlayerTransactionDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlayerTransactionDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlayerTransactionDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PlayerTransactionDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlayerTransactionDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlayerTransactionDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetQuantity

`func (o *PlayerTransactionDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlayerTransactionDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlayerTransactionDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetValue

`func (o *PlayerTransactionDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlayerTransactionDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlayerTransactionDto) SetValue(v int64)`

SetValue sets Value field to given value.


### GetAction

`func (o *PlayerTransactionDto) GetAction() TransactionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PlayerTransactionDto) GetActionOk() (*TransactionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PlayerTransactionDto) SetAction(v TransactionAction)`

SetAction sets Action field to given value.


### GetCreatorId

`func (o *PlayerTransactionDto) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PlayerTransactionDto) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PlayerTransactionDto) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetPlayerId

`func (o *PlayerTransactionDto) GetPlayerId() int32`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *PlayerTransactionDto) GetPlayerIdOk() (*int32, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *PlayerTransactionDto) SetPlayerId(v int32)`

SetPlayerId sets PlayerId field to given value.


### GetCreator

`func (o *PlayerTransactionDto) GetCreator() CreatorPartialDto`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *PlayerTransactionDto) GetCreatorOk() (*CreatorPartialDto, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *PlayerTransactionDto) SetCreator(v CreatorPartialDto)`

SetCreator sets Creator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


