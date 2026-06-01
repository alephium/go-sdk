# \AddressesAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressesAddressBalance**](AddressesAPI.md#GetAddressesAddressBalance) | **Get** /addresses/{address}/balance | Get the balance of an address
[**GetAddressesAddressGroup**](AddressesAPI.md#GetAddressesAddressGroup) | **Get** /addresses/{address}/group | Get the group of an address
[**GetAddressesAddressUtxos**](AddressesAPI.md#GetAddressesAddressUtxos) | **Get** /addresses/{address}/utxos | Get the UTXOs of an address



## GetAddressesAddressBalance

> Balance GetAddressesAddressBalance(ctx, address).Mempool(mempool).Execute()

Get the balance of an address

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
	address := "address_example" // string | 
	mempool := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressesAddressBalance(context.Background(), address).Mempool(mempool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressesAddressBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressesAddressBalance`: Balance
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressesAddressBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressesAddressBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mempool** | **bool** |  | 

### Return type

[**Balance**](Balance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressesAddressGroup

> Group GetAddressesAddressGroup(ctx, address).Execute()

Get the group of an address

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
	address := "address_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressesAddressGroup(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressesAddressGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressesAddressGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressesAddressGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressesAddressGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressesAddressUtxos

> UTXOs GetAddressesAddressUtxos(ctx, address).ErrorIfExceedMaxUtxos(errorIfExceedMaxUtxos).Execute()

Get the UTXOs of an address

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
	address := "address_example" // string | 
	errorIfExceedMaxUtxos := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressesAddressUtxos(context.Background(), address).ErrorIfExceedMaxUtxos(errorIfExceedMaxUtxos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressesAddressUtxos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressesAddressUtxos`: UTXOs
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressesAddressUtxos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressesAddressUtxosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorIfExceedMaxUtxos** | **bool** |  | 

### Return type

[**UTXOs**](UTXOs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

