# \MultiSignatureApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMultisigAddress**](MultiSignatureApi.md#PostMultisigAddress) | **Post** /multisig/address | Create the multisig address and unlock script
[**PostMultisigBuild**](MultiSignatureApi.md#PostMultisigBuild) | **Post** /multisig/build | Build a multisig unsigned transaction
[**PostMultisigSubmit**](MultiSignatureApi.md#PostMultisigSubmit) | **Post** /multisig/submit | Submit a multi-signed transaction



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
    openapiclient "./openapi"
)

func main() {
    buildMultisigAddress := *openapiclient.NewBuildMultisigAddress([]string{"Keys_example"}, int32(123)) // BuildMultisigAddress | Format 1: `1000000000000000000`  Format 2: `x.y ALPH`, where `1 ALPH = 1000000000000000000`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiSignatureApi.PostMultisigAddress(context.Background()).BuildMultisigAddress(buildMultisigAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureApi.PostMultisigAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMultisigAddress`: BuildMultisigAddressResult
    fmt.Fprintf(os.Stdout, "Response from `MultiSignatureApi.PostMultisigAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultisigAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildMultisigAddress** | [**BuildMultisigAddress**](BuildMultisigAddress.md) | Format 1: &#x60;1000000000000000000&#x60;  Format 2: &#x60;x.y ALPH&#x60;, where &#x60;1 ALPH &#x3D; 1000000000000000000&#x60; | 

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

> BuildTransactionResult PostMultisigBuild(ctx).BuildMultisig(buildMultisig).Execute()

Build a multisig unsigned transaction

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
    buildMultisig := *openapiclient.NewBuildMultisig("FromAddress_example", []string{"FromPublicKeys_example"}, []openapiclient.Destination{*openapiclient.NewDestination("Address_example", "AttoAlphAmount_example")}) // BuildMultisig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiSignatureApi.PostMultisigBuild(context.Background()).BuildMultisig(buildMultisig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureApi.PostMultisigBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMultisigBuild`: BuildTransactionResult
    fmt.Fprintf(os.Stdout, "Response from `MultiSignatureApi.PostMultisigBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultisigBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildMultisig** | [**BuildMultisig**](BuildMultisig.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    submitMultisig := *openapiclient.NewSubmitMultisig("UnsignedTx_example", []string{"Signatures_example"}) // SubmitMultisig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiSignatureApi.PostMultisigSubmit(context.Background()).SubmitMultisig(submitMultisig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiSignatureApi.PostMultisigSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMultisigSubmit`: SubmitTxResult
    fmt.Fprintf(os.Stdout, "Response from `MultiSignatureApi.PostMultisigSubmit`: %v\n", resp)
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

