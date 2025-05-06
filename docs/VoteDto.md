# VoteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorId** | **int32** |  | 
**Value** | **int64** |  | 
**Time** | **time.Time** |  | 

## Methods

### NewVoteDto

`func NewVoteDto(creatorId int32, value int64, time time.Time, ) *VoteDto`

NewVoteDto instantiates a new VoteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteDtoWithDefaults

`func NewVoteDtoWithDefaults() *VoteDto`

NewVoteDtoWithDefaults instantiates a new VoteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorId

`func (o *VoteDto) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *VoteDto) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *VoteDto) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetValue

`func (o *VoteDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VoteDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VoteDto) SetValue(v int64)`

SetValue sets Value field to given value.


### GetTime

`func (o *VoteDto) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VoteDto) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VoteDto) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


