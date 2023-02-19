# \MempoolApi

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMempoolTransactions**](MempoolApi.md#DeleteMempoolTransactions) | **Delete** /mempool/transactions | Remove all transactions from mempool
[**GetMempoolTransactions**](MempoolApi.md#GetMempoolTransactions) | **Get** /mempool/transactions | List mempool transactions
[**PutMempoolTransactionsRebroadcast**](MempoolApi.md#PutMempoolTransactionsRebroadcast) | **Put** /mempool/transactions/rebroadcast | Rebroadcase a mempool transaction to the network
[**PutMempoolTransactionsValidate**](MempoolApi.md#PutMempoolTransactionsValidate) | **Put** /mempool/transactions/validate | Validate all mempool transactions and remove invalid ones



## DeleteMempoolTransactions

> DeleteMempoolTransactions(ctx).Execute()

Remove all transactions from mempool

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
    resp, r, err := apiClient.MempoolApi.DeleteMempoolTransactions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MempoolApi.DeleteMempoolTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMempoolTransactionsRequest struct via the builder pattern


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


## GetMempoolTransactions

> []MempoolTransactions GetMempoolTransactions(ctx).Execute()

List mempool transactions

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
    resp, r, err := apiClient.MempoolApi.GetMempoolTransactions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MempoolApi.GetMempoolTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMempoolTransactions`: []MempoolTransactions
    fmt.Fprintf(os.Stdout, "Response from `MempoolApi.GetMempoolTransactions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMempoolTransactionsRequest struct via the builder pattern


### Return type

[**[]MempoolTransactions**](MempoolTransactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMempoolTransactionsRebroadcast

> PutMempoolTransactionsRebroadcast(ctx).TxId(txId).Execute()

Rebroadcase a mempool transaction to the network

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MempoolApi.PutMempoolTransactionsRebroadcast(context.Background()).TxId(txId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MempoolApi.PutMempoolTransactionsRebroadcast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMempoolTransactionsRebroadcastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txId** | **string** |  | 

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


## PutMempoolTransactionsValidate

> PutMempoolTransactionsValidate(ctx).Execute()

Validate all mempool transactions and remove invalid ones

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
    resp, r, err := apiClient.MempoolApi.PutMempoolTransactionsValidate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MempoolApi.PutMempoolTransactionsValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutMempoolTransactionsValidateRequest struct via the builder pattern


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

