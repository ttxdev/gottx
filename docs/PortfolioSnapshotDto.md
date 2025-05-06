# PortfolioSnapshotDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerId** | **int32** |  | 
**Value** | **int64** |  | 
**Time** | **time.Time** |  | 

## Methods

### NewPortfolioSnapshotDto

`func NewPortfolioSnapshotDto(playerId int32, value int64, time time.Time, ) *PortfolioSnapshotDto`

NewPortfolioSnapshotDto instantiates a new PortfolioSnapshotDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioSnapshotDtoWithDefaults

`func NewPortfolioSnapshotDtoWithDefaults() *PortfolioSnapshotDto`

NewPortfolioSnapshotDtoWithDefaults instantiates a new PortfolioSnapshotDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerId

`func (o *PortfolioSnapshotDto) GetPlayerId() int32`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *PortfolioSnapshotDto) GetPlayerIdOk() (*int32, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *PortfolioSnapshotDto) SetPlayerId(v int32)`

SetPlayerId sets PlayerId field to given value.


### GetValue

`func (o *PortfolioSnapshotDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PortfolioSnapshotDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PortfolioSnapshotDto) SetValue(v int64)`

SetValue sets Value field to given value.


### GetTime

`func (o *PortfolioSnapshotDto) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PortfolioSnapshotDto) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PortfolioSnapshotDto) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


