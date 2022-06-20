# \UtilsApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUtilsVerifySignature**](UtilsApi.md#PostUtilsVerifySignature) | **Post** /utils/verify-signature | Verify the SecP256K1 signature of some data
[**PutUtilsCheckHashIndexing**](UtilsApi.md#PutUtilsCheckHashIndexing) | **Put** /utils/check-hash-indexing | Check and repair the indexing of block hashes



## PostUtilsVerifySignature

> bool PostUtilsVerifySignature(ctx).VerifySignature(verifySignature).Execute()

Verify the SecP256K1 signature of some data

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
    verifySignature := *openapiclient.NewVerifySignature("Data_example", "Signature_example", "PublicKey_example") // VerifySignature | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilsApi.PostUtilsVerifySignature(context.Background()).VerifySignature(verifySignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.PostUtilsVerifySignature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUtilsVerifySignature`: bool
    fmt.Fprintf(os.Stdout, "Response from `UtilsApi.PostUtilsVerifySignature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUtilsVerifySignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifySignature** | [**VerifySignature**](VerifySignature.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUtilsCheckHashIndexing

> PutUtilsCheckHashIndexing(ctx).Execute()

Check and repair the indexing of block hashes

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
    resp, r, err := apiClient.UtilsApi.PutUtilsCheckHashIndexing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.PutUtilsCheckHashIndexing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutUtilsCheckHashIndexingRequest struct via the builder pattern


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

