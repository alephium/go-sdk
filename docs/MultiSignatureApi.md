# \MultiSignatureAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMultisigAddress**](MultiSignatureAPI.md#PostMultisigAddress) | **Post** /multisig/address | Create the multisig address and unlock script
[**PostMultisigBuild**](MultiSignatureAPI.md#PostMultisigBuild) | **Post** /multisig/build | Build a multisig unsigned transaction
[**PostMultisigSubmit**](MultiSignatureAPI.md#PostMultisigSubmit) | **Post** /multisig/submit | Submit a multi-signed transaction
[**PostMultisigSweep**](MultiSignatureAPI.md#PostMultisigSweep) | **Post** /multisig/sweep | Sweep all unlocked ALPH and token balances of a multisig address to another address



## PostMultisigAddress

> BuildMultisigAddressResult PostMultisigAddress(ctx).BuildMultisigAddress(buildMultisigAddress).Execute()

Create the multisig address and unlock script

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
	buildMultisigAddress := *openapiclient.NewBuildMultisigAddress([]string{"Keys_example"}, int32(123)) // BuildMultisigAddress | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000  Field fromPublicKeyType can be  `default` or `bip340-schnorr`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiSignatureAPI.PostMultisigAddress(context.Background()).BuildMultisigAddress(buildMultisigAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureAPI.PostMultisigAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultisigAddress`: BuildMultisigAddressResult
	fmt.Fprintf(os.Stdout, "Response from `MultiSignatureAPI.PostMultisigAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultisigAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildMultisigAddress** | [**BuildMultisigAddress**](BuildMultisigAddress.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000  Field fromPublicKeyType can be  &#x60;default&#x60; or &#x60;bip340-schnorr&#x60; | 

### Return type

[**BuildMultisigAddressResult**](BuildMultisigAddressResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMultisigBuild

> BuildTransferTxResult PostMultisigBuild(ctx).BuildMultisig(buildMultisig).Execute()

Build a multisig unsigned transaction

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
	buildMultisig := *openapiclient.NewBuildMultisig("FromAddress_example", []string{"FromPublicKeys_example"}, []openapiclient.Destination{*openapiclient.NewDestination("Address_example")}) // BuildMultisig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiSignatureAPI.PostMultisigBuild(context.Background()).BuildMultisig(buildMultisig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureAPI.PostMultisigBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultisigBuild`: BuildTransferTxResult
	fmt.Fprintf(os.Stdout, "Response from `MultiSignatureAPI.PostMultisigBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultisigBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildMultisig** | [**BuildMultisig**](BuildMultisig.md) |  | 

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


## PostMultisigSubmit

> SubmitTxResult PostMultisigSubmit(ctx).SubmitMultisig(submitMultisig).Execute()

Submit a multi-signed transaction

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
	submitMultisig := *openapiclient.NewSubmitMultisig("UnsignedTx_example", []string{"Signatures_example"}) // SubmitMultisig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiSignatureAPI.PostMultisigSubmit(context.Background()).SubmitMultisig(submitMultisig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureAPI.PostMultisigSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultisigSubmit`: SubmitTxResult
	fmt.Fprintf(os.Stdout, "Response from `MultiSignatureAPI.PostMultisigSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultisigSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitMultisig** | [**SubmitMultisig**](SubmitMultisig.md) |  | 

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


## PostMultisigSweep

> BuildSweepAddressTransactionsResult PostMultisigSweep(ctx).BuildSweepMultisig(buildSweepMultisig).Execute()

Sweep all unlocked ALPH and token balances of a multisig address to another address

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
	buildSweepMultisig := *openapiclient.NewBuildSweepMultisig("FromAddress_example", []string{"FromPublicKeys_example"}, "ToAddress_example") // BuildSweepMultisig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiSignatureAPI.PostMultisigSweep(context.Background()).BuildSweepMultisig(buildSweepMultisig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureAPI.PostMultisigSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultisigSweep`: BuildSweepAddressTransactionsResult
	fmt.Fprintf(os.Stdout, "Response from `MultiSignatureAPI.PostMultisigSweep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultisigSweepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildSweepMultisig** | [**BuildSweepMultisig**](BuildSweepMultisig.md) |  | 

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

