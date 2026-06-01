# \WalletsAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWalletsWalletName**](WalletsAPI.md#DeleteWalletsWalletName) | **Delete** /wallets/{wallet_name} | Delete your wallet file (can be recovered with your mnemonic)
[**GetWallets**](WalletsAPI.md#GetWallets) | **Get** /wallets | List available wallets
[**GetWalletsWalletName**](WalletsAPI.md#GetWalletsWalletName) | **Get** /wallets/{wallet_name} | Get wallet&#39;s status
[**GetWalletsWalletNameAddresses**](WalletsAPI.md#GetWalletsWalletNameAddresses) | **Get** /wallets/{wallet_name}/addresses | List all your wallet&#39;s addresses
[**GetWalletsWalletNameAddressesAddress**](WalletsAPI.md#GetWalletsWalletNameAddressesAddress) | **Get** /wallets/{wallet_name}/addresses/{address} | Get address&#39; info
[**GetWalletsWalletNameBalances**](WalletsAPI.md#GetWalletsWalletNameBalances) | **Get** /wallets/{wallet_name}/balances | Get your total balance
[**PostWallets**](WalletsAPI.md#PostWallets) | **Post** /wallets | Create a new wallet
[**PostWalletsWalletNameChangeActiveAddress**](WalletsAPI.md#PostWalletsWalletNameChangeActiveAddress) | **Post** /wallets/{wallet_name}/change-active-address | Choose the active address
[**PostWalletsWalletNameDeriveNextAddress**](WalletsAPI.md#PostWalletsWalletNameDeriveNextAddress) | **Post** /wallets/{wallet_name}/derive-next-address | Derive your next address
[**PostWalletsWalletNameLock**](WalletsAPI.md#PostWalletsWalletNameLock) | **Post** /wallets/{wallet_name}/lock | Lock your wallet
[**PostWalletsWalletNameRevealMnemonic**](WalletsAPI.md#PostWalletsWalletNameRevealMnemonic) | **Post** /wallets/{wallet_name}/reveal-mnemonic | Reveal your mnemonic. !!! use it with caution !!!
[**PostWalletsWalletNameSign**](WalletsAPI.md#PostWalletsWalletNameSign) | **Post** /wallets/{wallet_name}/sign | Sign the given data and return back the signature
[**PostWalletsWalletNameSweepActiveAddress**](WalletsAPI.md#PostWalletsWalletNameSweepActiveAddress) | **Post** /wallets/{wallet_name}/sweep-active-address | Transfer all unlocked ALPH from the active address to another address
[**PostWalletsWalletNameSweepAllAddresses**](WalletsAPI.md#PostWalletsWalletNameSweepAllAddresses) | **Post** /wallets/{wallet_name}/sweep-all-addresses | Transfer unlocked ALPH from all addresses (including all mining addresses if applicable) to another address
[**PostWalletsWalletNameTransfer**](WalletsAPI.md#PostWalletsWalletNameTransfer) | **Post** /wallets/{wallet_name}/transfer | Transfer ALPH from the active address
[**PostWalletsWalletNameUnlock**](WalletsAPI.md#PostWalletsWalletNameUnlock) | **Post** /wallets/{wallet_name}/unlock | Unlock your wallet
[**PutWallets**](WalletsAPI.md#PutWallets) | **Put** /wallets | Restore a wallet from your mnemonic



## DeleteWalletsWalletName

> DeleteWalletsWalletName(ctx, walletName).Password(password).Execute()

Delete your wallet file (can be recovered with your mnemonic)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	password := "password_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsAPI.DeleteWalletsWalletName(context.Background(), walletName).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.DeleteWalletsWalletName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWalletsWalletNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallets

> []WalletStatus GetWallets(ctx).Execute()

List available wallets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWallets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallets`: []WalletStatus
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWallets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsRequest struct via the builder pattern


### Return type

[**[]WalletStatus**](WalletStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletsWalletName

> WalletStatus GetWalletsWalletName(ctx, walletName).Execute()

Get wallet's status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletsWalletName(context.Background(), walletName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletsWalletName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsWalletName`: WalletStatus
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletsWalletName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsWalletNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletStatus**](WalletStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletsWalletNameAddresses

> Addresses GetWalletsWalletNameAddresses(ctx, walletName).Execute()

List all your wallet's addresses

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletsWalletNameAddresses(context.Background(), walletName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletsWalletNameAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsWalletNameAddresses`: Addresses
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletsWalletNameAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsWalletNameAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Addresses**](Addresses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletsWalletNameAddressesAddress

> AddressInfo GetWalletsWalletNameAddressesAddress(ctx, walletName, address).Execute()

Get address' info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	address := "address_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletsWalletNameAddressesAddress(context.Background(), walletName, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletsWalletNameAddressesAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsWalletNameAddressesAddress`: AddressInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletsWalletNameAddressesAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsWalletNameAddressesAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddressInfo**](AddressInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletsWalletNameBalances

> Balances GetWalletsWalletNameBalances(ctx, walletName).Execute()

Get your total balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletsWalletNameBalances(context.Background(), walletName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletsWalletNameBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsWalletNameBalances`: Balances
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletsWalletNameBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsWalletNameBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Balances**](Balances.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWallets

> WalletCreationResult PostWallets(ctx).WalletCreation(walletCreation).Execute()

Create a new wallet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletCreation := *openapiclient.NewWalletCreation("Password_example", "WalletName_example") // WalletCreation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWallets(context.Background()).WalletCreation(walletCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWallets`: WalletCreationResult
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletCreation** | [**WalletCreation**](WalletCreation.md) |  | 

### Return type

[**WalletCreationResult**](WalletCreationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameChangeActiveAddress

> PostWalletsWalletNameChangeActiveAddress(ctx, walletName).ChangeActiveAddress(changeActiveAddress).Execute()

Choose the active address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	changeActiveAddress := *openapiclient.NewChangeActiveAddress("Address_example") // ChangeActiveAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsAPI.PostWalletsWalletNameChangeActiveAddress(context.Background(), walletName).ChangeActiveAddress(changeActiveAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameChangeActiveAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameChangeActiveAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeActiveAddress** | [**ChangeActiveAddress**](ChangeActiveAddress.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameDeriveNextAddress

> AddressInfo PostWalletsWalletNameDeriveNextAddress(ctx, walletName).Group(group).Execute()

Derive your next address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	group := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWalletsWalletNameDeriveNextAddress(context.Background(), walletName).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameDeriveNextAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsWalletNameDeriveNextAddress`: AddressInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsWalletNameDeriveNextAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameDeriveNextAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **int32** |  | 

### Return type

[**AddressInfo**](AddressInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameLock

> PostWalletsWalletNameLock(ctx, walletName).Execute()

Lock your wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsAPI.PostWalletsWalletNameLock(context.Background(), walletName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameRevealMnemonic

> RevealMnemonicResult PostWalletsWalletNameRevealMnemonic(ctx, walletName).RevealMnemonic(revealMnemonic).Execute()

Reveal your mnemonic. !!! use it with caution !!!

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	revealMnemonic := *openapiclient.NewRevealMnemonic("Password_example") // RevealMnemonic | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWalletsWalletNameRevealMnemonic(context.Background(), walletName).RevealMnemonic(revealMnemonic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameRevealMnemonic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsWalletNameRevealMnemonic`: RevealMnemonicResult
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsWalletNameRevealMnemonic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameRevealMnemonicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revealMnemonic** | [**RevealMnemonic**](RevealMnemonic.md) |  | 

### Return type

[**RevealMnemonicResult**](RevealMnemonicResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameSign

> SignResult PostWalletsWalletNameSign(ctx, walletName).Sign(sign).Execute()

Sign the given data and return back the signature

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	sign := *openapiclient.NewSign("Data_example") // Sign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWalletsWalletNameSign(context.Background(), walletName).Sign(sign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsWalletNameSign`: SignResult
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsWalletNameSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sign** | [**Sign**](Sign.md) |  | 

### Return type

[**SignResult**](SignResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameSweepActiveAddress

> TransferResults PostWalletsWalletNameSweepActiveAddress(ctx, walletName).Sweep(sweep).Execute()

Transfer all unlocked ALPH from the active address to another address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	sweep := *openapiclient.NewSweep("ToAddress_example") // Sweep | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWalletsWalletNameSweepActiveAddress(context.Background(), walletName).Sweep(sweep).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameSweepActiveAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsWalletNameSweepActiveAddress`: TransferResults
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsWalletNameSweepActiveAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameSweepActiveAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sweep** | [**Sweep**](Sweep.md) |  | 

### Return type

[**TransferResults**](TransferResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameSweepAllAddresses

> TransferResults PostWalletsWalletNameSweepAllAddresses(ctx, walletName).Sweep(sweep).Execute()

Transfer unlocked ALPH from all addresses (including all mining addresses if applicable) to another address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	sweep := *openapiclient.NewSweep("ToAddress_example") // Sweep | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWalletsWalletNameSweepAllAddresses(context.Background(), walletName).Sweep(sweep).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameSweepAllAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsWalletNameSweepAllAddresses`: TransferResults
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsWalletNameSweepAllAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameSweepAllAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sweep** | [**Sweep**](Sweep.md) |  | 

### Return type

[**TransferResults**](TransferResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameTransfer

> TransferResult PostWalletsWalletNameTransfer(ctx, walletName).Transfer(transfer).Execute()

Transfer ALPH from the active address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	transfer := *openapiclient.NewTransfer([]openapiclient.Destination{*openapiclient.NewDestination("Address_example")}) // Transfer | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000  Field fromPublicKeyType can be  `default` or `bip340-schnorr`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PostWalletsWalletNameTransfer(context.Background(), walletName).Transfer(transfer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsWalletNameTransfer`: TransferResult
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsWalletNameTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transfer** | [**Transfer**](Transfer.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000  Field fromPublicKeyType can be  &#x60;default&#x60; or &#x60;bip340-schnorr&#x60; | 

### Return type

[**TransferResult**](TransferResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameUnlock

> PostWalletsWalletNameUnlock(ctx, walletName).WalletUnlock(walletUnlock).Execute()

Unlock your wallet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletName := "walletName_example" // string | 
	walletUnlock := *openapiclient.NewWalletUnlock("Password_example") // WalletUnlock | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletsAPI.PostWalletsWalletNameUnlock(context.Background(), walletName).WalletUnlock(walletUnlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsWalletNameUnlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletUnlock** | [**WalletUnlock**](WalletUnlock.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWallets

> WalletRestoreResult PutWallets(ctx).WalletRestore(walletRestore).Execute()

Restore a wallet from your mnemonic

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	walletRestore := *openapiclient.NewWalletRestore("Password_example", "Mnemonic_example", "WalletName_example") // WalletRestore | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.PutWallets(context.Background()).WalletRestore(walletRestore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PutWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWallets`: WalletRestoreResult
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PutWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletRestore** | [**WalletRestore**](WalletRestore.md) |  | 

### Return type

[**WalletRestoreResult**](WalletRestoreResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

