# PlayerDtoPaginationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PlayerDto**](PlayerDto.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewPlayerDtoPaginationDto

`func NewPlayerDtoPaginationDto(data []PlayerDto, total int32, ) *PlayerDtoPaginationDto`

NewPlayerDtoPaginationDto instantiates a new PlayerDtoPaginationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerDtoPaginationDtoWithDefaults

`func NewPlayerDtoPaginationDtoWithDefaults() *PlayerDtoPaginationDto`

NewPlayerDtoPaginationDtoWithDefaults instantiates a new PlayerDtoPaginationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PlayerDtoPaginationDto) GetData() []PlayerDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlayerDtoPaginationDto) GetDataOk() (*[]PlayerDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlayerDtoPaginationDto) SetData(v []PlayerDto)`

SetData sets Data field to given value.


### GetTotal

`func (o *PlayerDtoPaginationDto) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PlayerDtoPaginationDto) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PlayerDtoPaginationDto) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


