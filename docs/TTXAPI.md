# \TTXAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreator**](TTXAPI.md#CreateCreator) | **Post** /creators | 
[**DiscordCallback**](TTXAPI.md#DiscordCallback) | **Get** /sessions/discord/callback | 
[**Gamba**](TTXAPI.md#Gamba) | **Put** /players/me/lootboxes/{lootBoxId}/open | 
[**GetCreator**](TTXAPI.md#GetCreator) | **Get** /creators/{slug} | 
[**GetCreatorTransactions**](TTXAPI.md#GetCreatorTransactions) | **Get** /creators/{creatorSlug}/transactions | 
[**GetCreators**](TTXAPI.md#GetCreators) | **Get** /creators | 
[**GetPlayer**](TTXAPI.md#GetPlayer) | **Get** /players/{username} | 
[**GetPlayers**](TTXAPI.md#GetPlayers) | **Get** /players | 
[**GetSelf**](TTXAPI.md#GetSelf) | **Get** /players/me | 
[**LinkDiscordTwitch**](TTXAPI.md#LinkDiscordTwitch) | **Post** /sessions/discord/link | 
[**PlaceOrder**](TTXAPI.md#PlaceOrder) | **Post** /transactions | 
[**TwitchCallback**](TTXAPI.md#TwitchCallback) | **Get** /sessions/twitch/callback | 



## CreateCreator

> CreatorDto CreateCreator(ctx).Username(username).Ticker(ticker).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	username := "username_example" // string |  (optional)
	ticker := "ticker_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.CreateCreator(context.Background()).Username(username).Ticker(ticker).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.CreateCreator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreator`: CreatorDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.CreateCreator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **ticker** | **string** |  | 

### Return type

[**CreatorDto**](CreatorDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscordCallback

> DiscordTokenDto DiscordCallback(ctx).Code(code).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	code := "code_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.DiscordCallback(context.Background()).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.DiscordCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscordCallback`: DiscordTokenDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.DiscordCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscordCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 

### Return type

[**DiscordTokenDto**](DiscordTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Gamba

> LootBoxResultDto Gamba(ctx, lootBoxId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	lootBoxId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.Gamba(context.Background(), lootBoxId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.Gamba``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Gamba`: LootBoxResultDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.Gamba`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lootBoxId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGambaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LootBoxResultDto**](LootBoxResultDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreator

> CreatorDto GetCreator(ctx, slug).Step(step).After(after).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	slug := "slug_example" // string | 
	step := openapiclient.TimeStep("Minute") // TimeStep |  (optional)
	after := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.GetCreator(context.Background(), slug).Step(step).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.GetCreator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreator`: CreatorDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.GetCreator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **step** | [**TimeStep**](TimeStep.md) |  | 
 **after** | **time.Time** |  | 

### Return type

[**CreatorDto**](CreatorDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreatorTransactions

> []PlayerTransactionDto GetCreatorTransactions(ctx, creatorSlug).Slug(slug).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	creatorSlug := "creatorSlug_example" // string | 
	slug := "slug_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.GetCreatorTransactions(context.Background(), creatorSlug).Slug(slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.GetCreatorTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreatorTransactions`: []PlayerTransactionDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.GetCreatorTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creatorSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slug** | **string** |  | 

### Return type

[**[]PlayerTransactionDto**](PlayerTransactionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreators

> CreatorPartialDtoPaginationDto GetCreators(ctx).Page(page).Limit(limit).Search(search).OrderBy(orderBy).OrderDir(orderDir).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 20)
	search := "search_example" // string |  (optional)
	orderBy := openapiclient.CreatorOrderBy("Name") // CreatorOrderBy |  (optional)
	orderDir := openapiclient.OrderDirection("Ascending") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.GetCreators(context.Background()).Page(page).Limit(limit).Search(search).OrderBy(orderBy).OrderDir(orderDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.GetCreators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreators`: CreatorPartialDtoPaginationDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.GetCreators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]
 **search** | **string** |  | 
 **orderBy** | [**CreatorOrderBy**](CreatorOrderBy.md) |  | 
 **orderDir** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**CreatorPartialDtoPaginationDto**](CreatorPartialDtoPaginationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayer

> PlayerDto GetPlayer(ctx, username).Step(step).After(after).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	username := "username_example" // string | 
	step := openapiclient.TimeStep("Minute") // TimeStep |  (optional)
	after := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.GetPlayer(context.Background(), username).Step(step).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.GetPlayer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayer`: PlayerDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.GetPlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **step** | [**TimeStep**](TimeStep.md) |  | 
 **after** | **time.Time** |  | 

### Return type

[**PlayerDto**](PlayerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayers

> PlayerDtoPaginationDto GetPlayers(ctx).Page(page).Limit(limit).Search(search).OrderBy(orderBy).OrderDir(orderDir).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 20)
	search := "search_example" // string |  (optional)
	orderBy := openapiclient.PlayerOrderBy("Name") // PlayerOrderBy |  (optional)
	orderDir := openapiclient.OrderDirection("Ascending") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.GetPlayers(context.Background()).Page(page).Limit(limit).Search(search).OrderBy(orderBy).OrderDir(orderDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.GetPlayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayers`: PlayerDtoPaginationDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.GetPlayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 20]
 **search** | **string** |  | 
 **orderBy** | [**PlayerOrderBy**](PlayerOrderBy.md) |  | 
 **orderDir** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**PlayerDtoPaginationDto**](PlayerDtoPaginationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelf

> PlayerDto GetSelf(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.GetSelf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.GetSelf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelf`: PlayerDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.GetSelf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfRequest struct via the builder pattern


### Return type

[**PlayerDto**](PlayerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkDiscordTwitch

> TokenDto LinkDiscordTwitch(ctx).LinkDiscordTwitchDto(linkDiscordTwitchDto).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	linkDiscordTwitchDto := *openapiclient.NewLinkDiscordTwitchDto("AccessToken_example", "TwitchId_example") // LinkDiscordTwitchDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.LinkDiscordTwitch(context.Background()).LinkDiscordTwitchDto(linkDiscordTwitchDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.LinkDiscordTwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkDiscordTwitch`: TokenDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.LinkDiscordTwitch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkDiscordTwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkDiscordTwitchDto** | [**LinkDiscordTwitchDto**](LinkDiscordTwitchDto.md) |  | 

### Return type

[**TokenDto**](TokenDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrder

> CreatorTransactionDto PlaceOrder(ctx).CreateTransactionDto(createTransactionDto).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	createTransactionDto := *openapiclient.NewCreateTransactionDto("Creator_example", openapiclient.TransactionAction("Buy"), int32(123)) // CreateTransactionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.PlaceOrder(context.Background()).CreateTransactionDto(createTransactionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.PlaceOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaceOrder`: CreatorTransactionDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.PlaceOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransactionDto** | [**CreateTransactionDto**](CreateTransactionDto.md) |  | 

### Return type

[**CreatorTransactionDto**](CreatorTransactionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwitchCallback

> TokenDto TwitchCallback(ctx).Code(code).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ttxdev/gottx"
)

func main() {
	code := "code_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TTXAPI.TwitchCallback(context.Background()).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TTXAPI.TwitchCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwitchCallback`: TokenDto
	fmt.Fprintf(os.Stdout, "Response from `TTXAPI.TwitchCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwitchCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | 

### Return type

[**TokenDto**](TokenDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

