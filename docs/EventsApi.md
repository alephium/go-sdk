# \EventsApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsContractContractaddress**](EventsApi.md#GetEventsContractContractaddress) | **Get** /events/contract/{contractAddress} | Get events for a contract within a counter range
[**GetEventsContractContractaddressCurrentCount**](EventsApi.md#GetEventsContractContractaddressCurrentCount) | **Get** /events/contract/{contractAddress}/current-count | Get current value of the events counter for a contract
[**GetEventsTxIdTxid**](EventsApi.md#GetEventsTxIdTxid) | **Get** /events/tx-id/{txId} | Get events for a TxScript



## GetEventsContractContractaddress

> ContractEvents GetEventsContractContractaddress(ctx, contractAddress).Start(start).End(end).Group(group).Execute()

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
    end := int32(56) // int32 |  (optional)
    group := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsContractContractaddress(context.Background(), contractAddress).Start(start).End(end).Group(group).Execute()
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
 **end** | **int32** |  | 
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

Get events for a TxScript

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

