# gottx

TTX API Wrapper

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import gottx "github.com/ttxdev/go-ttx"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `gottx.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), gottx.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `gottx.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), gottx.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `gottx.ContextOperationServerIndices` and `gottx.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), gottx.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gottx.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*TTXAPI* | [**CreateCreator**](docs/TTXAPI.md#createcreator) | **Post** /creators | 
*TTXAPI* | [**DiscordCallback**](docs/TTXAPI.md#discordcallback) | **Get** /sessions/discord/callback | 
*TTXAPI* | [**Gamba**](docs/TTXAPI.md#gamba) | **Put** /players/me/lootboxes/{lootBoxId}/open | 
*TTXAPI* | [**GetCreator**](docs/TTXAPI.md#getcreator) | **Get** /creators/{slug} | 
*TTXAPI* | [**GetCreatorTransactions**](docs/TTXAPI.md#getcreatortransactions) | **Get** /creators/{creatorSlug}/transactions | 
*TTXAPI* | [**GetCreators**](docs/TTXAPI.md#getcreators) | **Get** /creators | 
*TTXAPI* | [**GetPlayer**](docs/TTXAPI.md#getplayer) | **Get** /players/{username} | 
*TTXAPI* | [**GetPlayers**](docs/TTXAPI.md#getplayers) | **Get** /players | 
*TTXAPI* | [**GetSelf**](docs/TTXAPI.md#getself) | **Get** /players/me | 
*TTXAPI* | [**LinkDiscordTwitch**](docs/TTXAPI.md#linkdiscordtwitch) | **Post** /sessions/discord/link | 
*TTXAPI* | [**PlaceOrder**](docs/TTXAPI.md#placeorder) | **Post** /transactions | 
*TTXAPI* | [**TwitchCallback**](docs/TTXAPI.md#twitchcallback) | **Get** /sessions/twitch/callback | 


## Documentation For Models

 - [CreateTransactionDto](docs/CreateTransactionDto.md)
 - [CreatorDto](docs/CreatorDto.md)
 - [CreatorOrderBy](docs/CreatorOrderBy.md)
 - [CreatorPartialDto](docs/CreatorPartialDto.md)
 - [CreatorPartialDtoPaginationDto](docs/CreatorPartialDtoPaginationDto.md)
 - [CreatorRarityDto](docs/CreatorRarityDto.md)
 - [CreatorShareDto](docs/CreatorShareDto.md)
 - [CreatorTransactionDto](docs/CreatorTransactionDto.md)
 - [DiscordTokenDto](docs/DiscordTokenDto.md)
 - [LinkDiscordTwitchDto](docs/LinkDiscordTwitchDto.md)
 - [LootBoxDto](docs/LootBoxDto.md)
 - [LootBoxResultDto](docs/LootBoxResultDto.md)
 - [ModelId](docs/ModelId.md)
 - [OrderDirection](docs/OrderDirection.md)
 - [PlayerDto](docs/PlayerDto.md)
 - [PlayerDtoPaginationDto](docs/PlayerDtoPaginationDto.md)
 - [PlayerOrderBy](docs/PlayerOrderBy.md)
 - [PlayerPartialDto](docs/PlayerPartialDto.md)
 - [PlayerShareDto](docs/PlayerShareDto.md)
 - [PlayerTransactionDto](docs/PlayerTransactionDto.md)
 - [PlayerType](docs/PlayerType.md)
 - [PortfolioSnapshotDto](docs/PortfolioSnapshotDto.md)
 - [Rarity](docs/Rarity.md)
 - [StreamStatusDto](docs/StreamStatusDto.md)
 - [TimeStep](docs/TimeStep.md)
 - [TokenDto](docs/TokenDto.md)
 - [TransactionAction](docs/TransactionAction.md)
 - [TwitchUserDto](docs/TwitchUserDto.md)
 - [VoteDto](docs/VoteDto.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author
[NathanRoberts55](https://github.com/nathanroberts55)