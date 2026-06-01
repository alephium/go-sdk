# \TransactionsAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionsDetailsTxid**](TransactionsAPI.md#GetTransactionsDetailsTxid) | **Get** /transactions/details/{txId} | Get transaction details
[**GetTransactionsRawTxid**](TransactionsAPI.md#GetTransactionsRawTxid) | **Get** /transactions/raw/{txId} | Get raw transaction in hex format
[**GetTransactionsRichDetailsTxid**](TransactionsAPI.md#GetTransactionsRichDetailsTxid) | **Get** /transactions/rich-details/{txId} | Get transaction with enriched input information when node indexes are enabled.
[**GetTransactionsStatus**](TransactionsAPI.md#GetTransactionsStatus) | **Get** /transactions/status | Get tx status
[**GetTransactionsTxIdFromOutputref**](TransactionsAPI.md#GetTransactionsTxIdFromOutputref) | **Get** /transactions/tx-id-from-outputref | Get transaction id from transaction output ref
[**PostTransactionsBuild**](TransactionsAPI.md#PostTransactionsBuild) | **Post** /transactions/build | Build an unsigned transfer transaction to a number of recipients
[**PostTransactionsBuildChained**](TransactionsAPI.md#PostTransactionsBuildChained) | **Post** /transactions/build-chained | Build a chain of transactions
[**PostTransactionsBuildMultiAddresses**](TransactionsAPI.md#PostTransactionsBuildMultiAddresses) | **Post** /transactions/build-multi-addresses | Build an unsigned transaction with multiple addresses to a number of recipients
[**PostTransactionsBuildTransferFromOneToManyGroups**](TransactionsAPI.md#PostTransactionsBuildTransferFromOneToManyGroups) | **Post** /transactions/build-transfer-from-one-to-many-groups | Build unsigned transfer transactions from an address of one group to addresses of many groups. Each target group requires a dedicated transaction or more in case large number of outputs needed to be split.
[**PostTransactionsDecodeUnsignedTx**](TransactionsAPI.md#PostTransactionsDecodeUnsignedTx) | **Post** /transactions/decode-unsigned-tx | Decode an unsigned transaction
[**PostTransactionsSubmit**](TransactionsAPI.md#PostTransactionsSubmit) | **Post** /transactions/submit | Submit a signed transaction
[**PostTransactionsSweepAddressBuild**](TransactionsAPI.md#PostTransactionsSweepAddressBuild) | **Post** /transactions/sweep-address/build | Build unsigned transactions to send all unlocked ALPH and token balances of one address to another address



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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	txId := "txId_example" // string | 
	fromGroup := int32(56) // int32 |  (optional)
	toGroup := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsDetailsTxid(context.Background(), txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsDetailsTxid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsDetailsTxid`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsDetailsTxid`: %v\n", resp)
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


## GetTransactionsRawTxid

> RawTransaction GetTransactionsRawTxid(ctx, txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()

Get raw transaction in hex format

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
	txId := "txId_example" // string | 
	fromGroup := int32(56) // int32 |  (optional)
	toGroup := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsRawTxid(context.Background(), txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsRawTxid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsRawTxid`: RawTransaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsRawTxid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRawTxidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromGroup** | **int32** |  | 
 **toGroup** | **int32** |  | 

### Return type

[**RawTransaction**](RawTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsRichDetailsTxid

> RichTransaction GetTransactionsRichDetailsTxid(ctx, txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()

Get transaction with enriched input information when node indexes are enabled.

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
	txId := "txId_example" // string | 
	fromGroup := int32(56) // int32 |  (optional)
	toGroup := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsRichDetailsTxid(context.Background(), txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsRichDetailsTxid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsRichDetailsTxid`: RichTransaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsRichDetailsTxid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRichDetailsTxidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromGroup** | **int32** |  | 
 **toGroup** | **int32** |  | 

### Return type

[**RichTransaction**](RichTransaction.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	txId := "txId_example" // string | 
	fromGroup := int32(56) // int32 |  (optional)
	toGroup := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsStatus(context.Background()).TxId(txId).FromGroup(fromGroup).ToGroup(toGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsStatus`: TxStatus
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsStatus`: %v\n", resp)
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


## GetTransactionsTxIdFromOutputref

> string GetTransactionsTxIdFromOutputref(ctx).Hint(hint).Key(key).Execute()

Get transaction id from transaction output ref

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
	hint := int32(56) // int32 | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionsTxIdFromOutputref(context.Background()).Hint(hint).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsTxIdFromOutputref``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionsTxIdFromOutputref`: string
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsTxIdFromOutputref`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsTxIdFromOutputrefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hint** | **int32** |  | 
 **key** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsBuild

> BuildTransferTxResult PostTransactionsBuild(ctx).BuildTransferTx(buildTransferTx).Execute()

Build an unsigned transfer transaction to a number of recipients

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
	buildTransferTx := *openapiclient.NewBuildTransferTx("FromPublicKey_example", []openapiclient.Destination{*openapiclient.NewDestination("Address_example")}) // BuildTransferTx | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000  Field fromPublicKeyType can be  `default` or `bip340-schnorr`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsBuild(context.Background()).BuildTransferTx(buildTransferTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsBuild`: BuildTransferTxResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildTransferTx** | [**BuildTransferTx**](BuildTransferTx.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000  Field fromPublicKeyType can be  &#x60;default&#x60; or &#x60;bip340-schnorr&#x60; | 

### Return type

[**BuildTransferTxResult**](BuildTransferTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsBuildChained

> []BuildChainedTxResult PostTransactionsBuildChained(ctx).BuildChainedTx(buildChainedTx).Execute()

Build a chain of transactions

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
	buildChainedTx := []openapiclient.BuildChainedTx{openapiclient.BuildChainedTx{BuildChainedDeployContractTx: openapiclient.NewBuildChainedDeployContractTx(*openapiclient.NewBuildDeployContractTx("FromPublicKey_example", "Bytecode_example"), "Type_example")}} // []BuildChainedTx | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsBuildChained(context.Background()).BuildChainedTx(buildChainedTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsBuildChained``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsBuildChained`: []BuildChainedTxResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsBuildChained`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsBuildChainedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildChainedTx** | [**[]BuildChainedTx**](BuildChainedTx.md) |  | 

### Return type

[**[]BuildChainedTxResult**](BuildChainedTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsBuildMultiAddresses

> BuildSimpleTransferTxResult PostTransactionsBuildMultiAddresses(ctx).BuildMultiAddressesTransaction(buildMultiAddressesTransaction).Execute()

Build an unsigned transaction with multiple addresses to a number of recipients

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
	buildMultiAddressesTransaction := *openapiclient.NewBuildMultiAddressesTransaction([]openapiclient.Source{*openapiclient.NewSource("FromPublicKey_example", []openapiclient.Destination{*openapiclient.NewDestination("Address_example")})}) // BuildMultiAddressesTransaction | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000  Field fromPublicKeyType can be  `default` or `bip340-schnorr`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsBuildMultiAddresses(context.Background()).BuildMultiAddressesTransaction(buildMultiAddressesTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsBuildMultiAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsBuildMultiAddresses`: BuildSimpleTransferTxResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsBuildMultiAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsBuildMultiAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildMultiAddressesTransaction** | [**BuildMultiAddressesTransaction**](BuildMultiAddressesTransaction.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000  Field fromPublicKeyType can be  &#x60;default&#x60; or &#x60;bip340-schnorr&#x60; | 

### Return type

[**BuildSimpleTransferTxResult**](BuildSimpleTransferTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionsBuildTransferFromOneToManyGroups

> []BuildSimpleTransferTxResult PostTransactionsBuildTransferFromOneToManyGroups(ctx).BuildTransferTx(buildTransferTx).Execute()

Build unsigned transfer transactions from an address of one group to addresses of many groups. Each target group requires a dedicated transaction or more in case large number of outputs needed to be split.

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
	buildTransferTx := *openapiclient.NewBuildTransferTx("FromPublicKey_example", []openapiclient.Destination{*openapiclient.NewDestination("Address_example")}) // BuildTransferTx | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000  Field fromPublicKeyType can be  `default` or `bip340-schnorr`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsBuildTransferFromOneToManyGroups(context.Background()).BuildTransferTx(buildTransferTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsBuildTransferFromOneToManyGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsBuildTransferFromOneToManyGroups`: []BuildSimpleTransferTxResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsBuildTransferFromOneToManyGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionsBuildTransferFromOneToManyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildTransferTx** | [**BuildTransferTx**](BuildTransferTx.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000  Field fromPublicKeyType can be  &#x60;default&#x60; or &#x60;bip340-schnorr&#x60; | 

### Return type

[**[]BuildSimpleTransferTxResult**](BuildSimpleTransferTxResult.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	decodeUnsignedTx := *openapiclient.NewDecodeUnsignedTx("UnsignedTx_example") // DecodeUnsignedTx | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsDecodeUnsignedTx(context.Background()).DecodeUnsignedTx(decodeUnsignedTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsDecodeUnsignedTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsDecodeUnsignedTx`: DecodeUnsignedTxResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsDecodeUnsignedTx`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	submitTransaction := *openapiclient.NewSubmitTransaction("UnsignedTx_example", "Signature_example") // SubmitTransaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsSubmit(context.Background()).SubmitTransaction(submitTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsSubmit`: SubmitTxResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsSubmit`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	buildSweepAddressTransactions := *openapiclient.NewBuildSweepAddressTransactions("FromPublicKey_example", "ToAddress_example") // BuildSweepAddressTransactions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.PostTransactionsSweepAddressBuild(context.Background()).BuildSweepAddressTransactions(buildSweepAddressTransactions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.PostTransactionsSweepAddressBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransactionsSweepAddressBuild`: BuildSweepAddressTransactionsResult
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.PostTransactionsSweepAddressBuild`: %v\n", resp)
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

