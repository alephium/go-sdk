# \UtilsAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostUtilsTargetToHashrate**](UtilsAPI.md#PostUtilsTargetToHashrate) | **Post** /utils/target-to-hashrate | Convert a target to hashrate
[**PostUtilsVerifySignature**](UtilsAPI.md#PostUtilsVerifySignature) | **Post** /utils/verify-signature | Verify the SecP256K1 signature of some data
[**PutUtilsCheckHashIndexing**](UtilsAPI.md#PutUtilsCheckHashIndexing) | **Put** /utils/check-hash-indexing | Check and repair the indexing of block hashes



## PostUtilsTargetToHashrate

> Result PostUtilsTargetToHashrate(ctx).TargetToHashrate(targetToHashrate).Execute()

Convert a target to hashrate

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
	targetToHashrate := *openapiclient.NewTargetToHashrate("Target_example") // TargetToHashrate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.PostUtilsTargetToHashrate(context.Background()).TargetToHashrate(targetToHashrate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.PostUtilsTargetToHashrate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUtilsTargetToHashrate`: Result
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.PostUtilsTargetToHashrate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUtilsTargetToHashrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetToHashrate** | [**TargetToHashrate**](TargetToHashrate.md) |  | 

### Return type

[**Result**](Result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	verifySignature := *openapiclient.NewVerifySignature("Data_example", "Signature_example", "PublicKey_example") // VerifySignature | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.PostUtilsVerifySignature(context.Background()).VerifySignature(verifySignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.PostUtilsVerifySignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUtilsVerifySignature`: bool
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.PostUtilsVerifySignature`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilsAPI.PutUtilsCheckHashIndexing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.PutUtilsCheckHashIndexing``: %v\n", err)
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

