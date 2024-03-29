# \TransactionsApi

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionsDetailsTxid**](TransactionsApi.md#GetTransactionsDetailsTxid) | **Get** /transactions/details/{txId} | Get transaction details
[**GetTransactionsStatus**](TransactionsApi.md#GetTransactionsStatus) | **Get** /transactions/status | Get tx status
[**PostTransactionsBuild**](TransactionsApi.md#PostTransactionsBuild) | **Post** /transactions/build | Build an unsigned transaction to a number of recipients
[**PostTransactionsDecodeUnsignedTx**](TransactionsApi.md#PostTransactionsDecodeUnsignedTx) | **Post** /transactions/decode-unsigned-tx | Decode an unsigned transaction
[**PostTransactionsSubmit**](TransactionsApi.md#PostTransactionsSubmit) | **Post** /transactions/submit | Submit a signed transaction
[**PostTransactionsSweepAddressBuild**](TransactionsApi.md#PostTransactionsSweepAddressBuild) | **Post** /transactions/sweep-address/build | Build unsigned transactions to send all unlocked ALPH and token balances of one address to another address



## GetTransactionsDetailsTxid

> Transaction GetTransactionsDetailsTxid(ctx, txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()

Get transaction details

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
    txId := "txId_example" // string | 
    fromGroup := int32(56) // int32 |  (optional)
    toGroup := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionsDetailsTxid(context.Background(), txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionsDetailsTxid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsDetailsTxid`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionsDetailsTxid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsDetailsTxidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromGroup** | **int32** |  | 
 **toGroup** | **int32** |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsStatus

> TxStatus GetTransactionsStatus(ctx).TxId(txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()

Get tx status

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
    txId := "txId_example" // string | 
    fromGroup := int32(56) // int32 |  (optional)
    toGroup := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionsStatus(context.Background()).TxId(txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsStatus`: TxStatus
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionsStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **string** |  | 
 **fromGroup** | **int32** |  | 
 **toGroup** | **int32** |  | 

### Return type

[**TxStatus**](TxStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsBuild

> BuildTransactionResult PostTransactionsBuild(ctx).BuildTransaction(buildTransaction).Execute()

Build an unsigned transaction to a number of recipients

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
    buildTransaction := *openapiclient.NewBuildTransaction("FromPublicKey_example", []openapiclient.Destination{*openapiclient.NewDestination("Address_example", "AttoAlphAmount_example")}) // BuildTransaction | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000  Field fromPublicKeyType can be  `default` or `bip340-schnorr`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionsBuild(context.Background()).BuildTransaction(buildTransaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionsBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionsBuild`: BuildTransactionResult
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionsBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildTransaction** | [**BuildTransaction**](BuildTransaction.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000  Field fromPublicKeyType can be  &#x60;default&#x60; or &#x60;bip340-schnorr&#x60; | 

### Return type

[**BuildTransactionResult**](BuildTransactionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsDecodeUnsignedTx

> DecodeUnsignedTxResult PostTransactionsDecodeUnsignedTx(ctx).DecodeUnsignedTx(decodeUnsignedTx).Execute()

Decode an unsigned transaction

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
    decodeUnsignedTx := *openapiclient.NewDecodeUnsignedTx("UnsignedTx_example") // DecodeUnsignedTx | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionsDecodeUnsignedTx(context.Background()).DecodeUnsignedTx(decodeUnsignedTx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionsDecodeUnsignedTx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionsDecodeUnsignedTx`: DecodeUnsignedTxResult
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionsDecodeUnsignedTx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsDecodeUnsignedTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decodeUnsignedTx** | [**DecodeUnsignedTx**](DecodeUnsignedTx.md) |  | 

### Return type

[**DecodeUnsignedTxResult**](DecodeUnsignedTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsSubmit

> SubmitTxResult PostTransactionsSubmit(ctx).SubmitTransaction(submitTransaction).Execute()

Submit a signed transaction

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
    submitTransaction := *openapiclient.NewSubmitTransaction("UnsignedTx_example", "Signature_example") // SubmitTransaction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionsSubmit(context.Background()).SubmitTransaction(submitTransaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionsSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionsSubmit`: SubmitTxResult
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionsSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTransaction** | [**SubmitTransaction**](SubmitTransaction.md) |  | 

### Return type

[**SubmitTxResult**](SubmitTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsSweepAddressBuild

> BuildSweepAddressTransactionsResult PostTransactionsSweepAddressBuild(ctx).BuildSweepAddressTransactions(buildSweepAddressTransactions).Execute()

Build unsigned transactions to send all unlocked ALPH and token balances of one address to another address

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
    buildSweepAddressTransactions := *openapiclient.NewBuildSweepAddressTransactions("FromPublicKey_example", "ToAddress_example") // BuildSweepAddressTransactions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionsSweepAddressBuild(context.Background()).BuildSweepAddressTransactions(buildSweepAddressTransactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionsSweepAddressBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionsSweepAddressBuild`: BuildSweepAddressTransactionsResult
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionsSweepAddressBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsSweepAddressBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildSweepAddressTransactions** | [**BuildSweepAddressTransactions**](BuildSweepAddressTransactions.md) |  | 

### Return type

[**BuildSweepAddressTransactionsResult**](BuildSweepAddressTransactionsResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

