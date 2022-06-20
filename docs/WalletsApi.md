# \WalletsApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWalletsWalletName**](WalletsApi.md#DeleteWalletsWalletName) | **Delete** /wallets/{wallet_name} | Delete your wallet file (can be recovered with your mnemonic)
[**GetWallets**](WalletsApi.md#GetWallets) | **Get** /wallets | List available wallets
[**GetWalletsWalletName**](WalletsApi.md#GetWalletsWalletName) | **Get** /wallets/{wallet_name} | Get wallet&#39;s status
[**GetWalletsWalletNameAddresses**](WalletsApi.md#GetWalletsWalletNameAddresses) | **Get** /wallets/{wallet_name}/addresses | List all your wallet&#39;s addresses
[**GetWalletsWalletNameAddressesAddress**](WalletsApi.md#GetWalletsWalletNameAddressesAddress) | **Get** /wallets/{wallet_name}/addresses/{address} | Get address&#39; info
[**GetWalletsWalletNameBalances**](WalletsApi.md#GetWalletsWalletNameBalances) | **Get** /wallets/{wallet_name}/balances | Get your total balance
[**PostWallets**](WalletsApi.md#PostWallets) | **Post** /wallets | Create a new wallet
[**PostWalletsWalletNameChangeActiveAddress**](WalletsApi.md#PostWalletsWalletNameChangeActiveAddress) | **Post** /wallets/{wallet_name}/change-active-address | Choose the active address
[**PostWalletsWalletNameDeriveNextAddress**](WalletsApi.md#PostWalletsWalletNameDeriveNextAddress) | **Post** /wallets/{wallet_name}/derive-next-address | Derive your next address
[**PostWalletsWalletNameLock**](WalletsApi.md#PostWalletsWalletNameLock) | **Post** /wallets/{wallet_name}/lock | Lock your wallet
[**PostWalletsWalletNameRevealMnemonic**](WalletsApi.md#PostWalletsWalletNameRevealMnemonic) | **Post** /wallets/{wallet_name}/reveal-mnemonic | Reveal your mnemonic. !!! use it with caution !!!
[**PostWalletsWalletNameSign**](WalletsApi.md#PostWalletsWalletNameSign) | **Post** /wallets/{wallet_name}/sign | Sign the given data and return back the signature
[**PostWalletsWalletNameSweepActiveAddress**](WalletsApi.md#PostWalletsWalletNameSweepActiveAddress) | **Post** /wallets/{wallet_name}/sweep-active-address | Transfer all unlocked ALPH from the active address to another address
[**PostWalletsWalletNameSweepAllAddresses**](WalletsApi.md#PostWalletsWalletNameSweepAllAddresses) | **Post** /wallets/{wallet_name}/sweep-all-addresses | Transfer unlocked ALPH from all addresses (including all mining addresses if applicable) to another address
[**PostWalletsWalletNameTransfer**](WalletsApi.md#PostWalletsWalletNameTransfer) | **Post** /wallets/{wallet_name}/transfer | Transfer ALPH from the active address
[**PostWalletsWalletNameUnlock**](WalletsApi.md#PostWalletsWalletNameUnlock) | **Post** /wallets/{wallet_name}/unlock | Unlock your wallet
[**PutWallets**](WalletsApi.md#PutWallets) | **Put** /wallets | Restore a wallet from your mnemonic



## DeleteWalletsWalletName

> DeleteWalletsWalletName(ctx, walletName).WalletDeletion(walletDeletion).Execute()

Delete your wallet file (can be recovered with your mnemonic)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    walletDeletion := *openapiclient.NewWalletDeletion("Password_example") // WalletDeletion | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.DeleteWalletsWalletName(context.Background(), walletName).WalletDeletion(walletDeletion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.DeleteWalletsWalletName``: %v\n", err)
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

 **walletDeletion** | [**WalletDeletion**](WalletDeletion.md) |  | 

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWallets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWallets`: []WalletStatus
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWallets`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWalletsWalletName(context.Background(), walletName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWalletsWalletName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletsWalletName`: WalletStatus
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWalletsWalletName`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWalletsWalletNameAddresses(context.Background(), walletName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWalletsWalletNameAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletsWalletNameAddresses`: Addresses
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWalletsWalletNameAddresses`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    address := "address_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWalletsWalletNameAddressesAddress(context.Background(), walletName, address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWalletsWalletNameAddressesAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletsWalletNameAddressesAddress`: AddressInfo
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWalletsWalletNameAddressesAddress`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWalletsWalletNameBalances(context.Background(), walletName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWalletsWalletNameBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletsWalletNameBalances`: Balances
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWalletsWalletNameBalances`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletCreation := *openapiclient.NewWalletCreation("Password_example", "WalletName_example") // WalletCreation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWallets(context.Background()).WalletCreation(walletCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWallets`: WalletCreationResult
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWallets`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    changeActiveAddress := *openapiclient.NewChangeActiveAddress("Address_example") // ChangeActiveAddress | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameChangeActiveAddress(context.Background(), walletName).ChangeActiveAddress(changeActiveAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameChangeActiveAddress``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    group := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameDeriveNextAddress(context.Background(), walletName).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameDeriveNextAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameDeriveNextAddress`: AddressInfo
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWalletsWalletNameDeriveNextAddress`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameLock(context.Background(), walletName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameLock``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    revealMnemonic := *openapiclient.NewRevealMnemonic("Password_example") // RevealMnemonic | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameRevealMnemonic(context.Background(), walletName).RevealMnemonic(revealMnemonic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameRevealMnemonic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameRevealMnemonic`: RevealMnemonicResult
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWalletsWalletNameRevealMnemonic`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    sign := *openapiclient.NewSign("Data_example") // Sign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameSign(context.Background(), walletName).Sign(sign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameSign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameSign`: SignResult
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWalletsWalletNameSign`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    sweep := *openapiclient.NewSweep("ToAddress_example") // Sweep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameSweepActiveAddress(context.Background(), walletName).Sweep(sweep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameSweepActiveAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameSweepActiveAddress`: TransferResults
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWalletsWalletNameSweepActiveAddress`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    sweep := *openapiclient.NewSweep("ToAddress_example") // Sweep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameSweepAllAddresses(context.Background(), walletName).Sweep(sweep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameSweepAllAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameSweepAllAddresses`: TransferResults
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWalletsWalletNameSweepAllAddresses`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    transfer := *openapiclient.NewTransfer([]openapiclient.Destination{*openapiclient.NewDestination("Address_example", "AttoAlphAmount_example")}) // Transfer | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameTransfer(context.Background(), walletName).Transfer(transfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameTransfer`: TransferResult
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PostWalletsWalletNameTransfer`: %v\n", resp)
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

 **transfer** | [**Transfer**](Transfer.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000&#x60; | 

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
    openapiclient "./openapi"
)

func main() {
    walletName := "walletName_example" // string | 
    walletUnlock := *openapiclient.NewWalletUnlock("Password_example") // WalletUnlock | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PostWalletsWalletNameUnlock(context.Background(), walletName).WalletUnlock(walletUnlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PostWalletsWalletNameUnlock``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    walletRestore := *openapiclient.NewWalletRestore("Password_example", "Mnemonic_example", "WalletName_example") // WalletRestore | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.PutWallets(context.Background()).WalletRestore(walletRestore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.PutWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWallets`: WalletRestoreResult
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.PutWallets`: %v\n", resp)
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

