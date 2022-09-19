# \EventsApi

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsBlockHashBlockhash**](EventsApi.md#GetEventsBlockHashBlockhash) | **Get** /events/block-hash/{blockHash} | Get contract events for a block
[**GetEventsContractContractaddress**](EventsApi.md#GetEventsContractContractaddress) | **Get** /events/contract/{contractAddress} | Get events for a contract within a counter range
[**GetEventsContractContractaddressCurrentCount**](EventsApi.md#GetEventsContractContractaddressCurrentCount) | **Get** /events/contract/{contractAddress}/current-count | Get current value of the events counter for a contract
[**GetEventsTxIdTxid**](EventsApi.md#GetEventsTxIdTxid) | **Get** /events/tx-id/{txId} | Get contract events for a transaction



## GetEventsBlockHashBlockhash

> ContractEventsByBlockHash GetEventsBlockHashBlockhash(ctx, blockHash).Group(group).Execute()

Get contract events for a block

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
    blockHash := "blockHash_example" // string | 
    group := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsBlockHashBlockhash(context.Background(), blockHash).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsBlockHashBlockhash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsBlockHashBlockhash`: ContractEventsByBlockHash
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsBlockHashBlockhash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsBlockHashBlockhashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **int32** |  | 

### Return type

[**ContractEventsByBlockHash**](ContractEventsByBlockHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsContractContractaddress

> ContractEvents GetEventsContractContractaddress(ctx, contractAddress).Start(start).Limit(limit).Group(group).Execute()

Get events for a contract within a counter range

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
    contractAddress := "contractAddress_example" // string | 
    start := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    group := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsContractContractaddress(context.Background(), contractAddress).Start(start).Limit(limit).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsContractContractaddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsContractContractaddress`: ContractEvents
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsContractContractaddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsContractContractaddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** |  | 
 **limit** | **int32** |  | 
 **group** | **int32** |  | 

### Return type

[**ContractEvents**](ContractEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsContractContractaddressCurrentCount

> int32 GetEventsContractContractaddressCurrentCount(ctx, contractAddress).Execute()

Get current value of the events counter for a contract

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
    contractAddress := "contractAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsContractContractaddressCurrentCount(context.Background(), contractAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsContractContractaddressCurrentCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsContractContractaddressCurrentCount`: int32
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsContractContractaddressCurrentCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsContractContractaddressCurrentCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsTxIdTxid

> ContractEventsByTxId GetEventsTxIdTxid(ctx, txId).Group(group).Execute()

Get contract events for a transaction

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
    group := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsTxIdTxid(context.Background(), txId).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsTxIdTxid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsTxIdTxid`: ContractEventsByTxId
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsTxIdTxid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsTxIdTxidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **int32** |  | 

### Return type

[**ContractEventsByTxId**](ContractEventsByTxId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

