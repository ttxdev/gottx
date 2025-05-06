# TwitchUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DisplayName** | **string** |  | 
**Login** | **string** |  | 
**AvatarUrl** | **string** |  | 

## Methods

### NewTwitchUserDto

`func NewTwitchUserDto(id string, displayName string, login string, avatarUrl string, ) *TwitchUserDto`

NewTwitchUserDto instantiates a new TwitchUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwitchUserDtoWithDefaults

`func NewTwitchUserDtoWithDefaults() *TwitchUserDto`

NewTwitchUserDtoWithDefaults instantiates a new TwitchUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TwitchUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TwitchUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TwitchUserDto) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *TwitchUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TwitchUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TwitchUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLogin

`func (o *TwitchUserDto) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *TwitchUserDto) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *TwitchUserDto) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetAvatarUrl

`func (o *TwitchUserDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TwitchUserDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TwitchUserDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


