# \MinersApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMinersAddresses**](MinersApi.md#GetMinersAddresses) | **Get** /miners/addresses | List miner&#39;s addresses
[**GetWalletsWalletNameMinerAddresses**](MinersApi.md#GetWalletsWalletNameMinerAddresses) | **Get** /wallets/{wallet_name}/miner-addresses | List all miner addresses per group
[**PostMinersCpuMining**](MinersApi.md#PostMinersCpuMining) | **Post** /miners/cpu-mining | Execute an action on CPU miner. !!! for test only !!!
[**PostWalletsWalletNameDeriveNextMinerAddresses**](MinersApi.md#PostWalletsWalletNameDeriveNextMinerAddresses) | **Post** /wallets/{wallet_name}/derive-next-miner-addresses | Derive your next miner addresses for each group
[**PutMinersAddresses**](MinersApi.md#PutMinersAddresses) | **Put** /miners/addresses | Update miner&#39;s addresses, but better to use user.conf instead



## GetMinersAddresses

> MinerAddresses GetMinersAddresses(ctx).Execute()

List miner's addresses

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
    resp, r, err := apiClient.MinersApi.GetMinersAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinersApi.GetMinersAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMinersAddresses`: MinerAddresses
    fmt.Fprintf(os.Stdout, "Response from `MinersApi.GetMinersAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMinersAddressesRequest struct via the builder pattern


### Return type

[**MinerAddresses**](MinerAddresses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletsWalletNameMinerAddresses

> []MinerAddressesInfo GetWalletsWalletNameMinerAddresses(ctx, walletName).Execute()

List all miner addresses per group



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
    resp, r, err := apiClient.MinersApi.GetWalletsWalletNameMinerAddresses(context.Background(), walletName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinersApi.GetWalletsWalletNameMinerAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletsWalletNameMinerAddresses`: []MinerAddressesInfo
    fmt.Fprintf(os.Stdout, "Response from `MinersApi.GetWalletsWalletNameMinerAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsWalletNameMinerAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MinerAddressesInfo**](MinerAddressesInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMinersCpuMining

> bool PostMinersCpuMining(ctx).Action(action).Execute()

Execute an action on CPU miner. !!! for test only !!!

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
    action := "start-mining" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MinersApi.PostMinersCpuMining(context.Background()).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinersApi.PostMinersCpuMining``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMinersCpuMining`: bool
    fmt.Fprintf(os.Stdout, "Response from `MinersApi.PostMinersCpuMining`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMinersCpuMiningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletsWalletNameDeriveNextMinerAddresses

> []AddressInfo PostWalletsWalletNameDeriveNextMinerAddresses(ctx, walletName).Execute()

Derive your next miner addresses for each group



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
    resp, r, err := apiClient.MinersApi.PostWalletsWalletNameDeriveNextMinerAddresses(context.Background(), walletName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinersApi.PostWalletsWalletNameDeriveNextMinerAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWalletsWalletNameDeriveNextMinerAddresses`: []AddressInfo
    fmt.Fprintf(os.Stdout, "Response from `MinersApi.PostWalletsWalletNameDeriveNextMinerAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsWalletNameDeriveNextMinerAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AddressInfo**](AddressInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMinersAddresses

> PutMinersAddresses(ctx).MinerAddresses(minerAddresses).Execute()

Update miner's addresses, but better to use user.conf instead

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
    minerAddresses := *openapiclient.NewMinerAddresses([]string{"Addresses_example"}) // MinerAddresses | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MinersApi.PutMinersAddresses(context.Background()).MinerAddresses(minerAddresses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MinersApi.PutMinersAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMinersAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minerAddresses** | [**MinerAddresses**](MinerAddresses.md) |  | 

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

