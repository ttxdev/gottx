# DiscordTokenDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**LinkToken** | **string** |  | 
**TwitchUsers** | [**[]TwitchUserDto**](TwitchUserDto.md) |  | [readonly] 

## Methods

### NewDiscordTokenDto

`func NewDiscordTokenDto(accessToken string, linkToken string, twitchUsers []TwitchUserDto, ) *DiscordTokenDto`

NewDiscordTokenDto instantiates a new DiscordTokenDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordTokenDtoWithDefaults

`func NewDiscordTokenDtoWithDefaults() *DiscordTokenDto`

NewDiscordTokenDtoWithDefaults instantiates a new DiscordTokenDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *DiscordTokenDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DiscordTokenDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DiscordTokenDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetLinkToken

`func (o *DiscordTokenDto) GetLinkToken() string`

GetLinkToken returns the LinkToken field if non-nil, zero value otherwise.

### GetLinkTokenOk

`func (o *DiscordTokenDto) GetLinkTokenOk() (*string, bool)`

GetLinkTokenOk returns a tuple with the LinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToken

`func (o *DiscordTokenDto) SetLinkToken(v string)`

SetLinkToken sets LinkToken field to given value.


### GetTwitchUsers

`func (o *DiscordTokenDto) GetTwitchUsers() []TwitchUserDto`

GetTwitchUsers returns the TwitchUsers field if non-nil, zero value otherwise.

### GetTwitchUsersOk

`func (o *DiscordTokenDto) GetTwitchUsersOk() (*[]TwitchUserDto, bool)`

GetTwitchUsersOk returns a tuple with the TwitchUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchUsers

`func (o *DiscordTokenDto) SetTwitchUsers(v []TwitchUserDto)`

SetTwitchUsers sets TwitchUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


