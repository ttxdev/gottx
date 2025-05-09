# CreateTransactionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creator** | **string** |  | 
**Action** | [**TransactionAction**](TransactionAction.md) |  | 
**Amount** | **int32** |  | 

## Methods

### NewCreateTransactionDto

`func NewCreateTransactionDto(creator string, action TransactionAction, amount int32, ) *CreateTransactionDto`

NewCreateTransactionDto instantiates a new CreateTransactionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactionDtoWithDefaults

`func NewCreateTransactionDtoWithDefaults() *CreateTransactionDto`

NewCreateTransactionDtoWithDefaults instantiates a new CreateTransactionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreator

`func (o *CreateTransactionDto) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateTransactionDto) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateTransactionDto) SetCreator(v string)`

SetCreator sets Creator field to given value.


### GetAction

`func (o *CreateTransactionDto) GetAction() TransactionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateTransactionDto) GetActionOk() (*TransactionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateTransactionDto) SetAction(v TransactionAction)`

SetAction sets Action field to given value.


### GetAmount

`func (o *CreateTransactionDto) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateTransactionDto) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateTransactionDto) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


