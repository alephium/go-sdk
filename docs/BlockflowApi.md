# \BlockflowApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockflow**](BlockflowApi.md#GetBlockflow) | **Get** /blockflow | List blocks on the given time interval
[**GetBlockflowBlocksBlockHash**](BlockflowApi.md#GetBlockflowBlocksBlockHash) | **Get** /blockflow/blocks/{block_hash} | Get a block with hash
[**GetBlockflowChainInfo**](BlockflowApi.md#GetBlockflowChainInfo) | **Get** /blockflow/chain-info | Get infos about the chain from the given groups
[**GetBlockflowHashes**](BlockflowApi.md#GetBlockflowHashes) | **Get** /blockflow/hashes | Get all block&#39;s hashes at given height for given groups
[**GetBlockflowHeadersBlockHash**](BlockflowApi.md#GetBlockflowHeadersBlockHash) | **Get** /blockflow/headers/{block_hash} | Get block header
[**GetBlockflowIsBlockInMainChain**](BlockflowApi.md#GetBlockflowIsBlockInMainChain) | **Get** /blockflow/is-block-in-main-chain | Check if the block is in main chain



## GetBlockflow

> FetchResponse GetBlockflow(ctx).FromTs(fromTs).ToTs(toTs).Execute()

List blocks on the given time interval

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
    fromTs := int64(789) // int64 | 
    toTs := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockflowApi.GetBlockflow(context.Background()).FromTs(fromTs).ToTs(toTs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockflowApi.GetBlockflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockflow`: FetchResponse
    fmt.Fprintf(os.Stdout, "Response from `BlockflowApi.GetBlockflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromTs** | **int64** |  | 
 **toTs** | **int64** |  | 

### Return type

[**FetchResponse**](FetchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowBlocksBlockHash

> BlockEntry GetBlockflowBlocksBlockHash(ctx, blockHash).Execute()

Get a block with hash

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockflowApi.GetBlockflowBlocksBlockHash(context.Background(), blockHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockflowApi.GetBlockflowBlocksBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockflowBlocksBlockHash`: BlockEntry
    fmt.Fprintf(os.Stdout, "Response from `BlockflowApi.GetBlockflowBlocksBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowBlocksBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockEntry**](BlockEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowChainInfo

> ChainInfo GetBlockflowChainInfo(ctx).FromGroup(fromGroup).ToGroup(toGroup).Execute()

Get infos about the chain from the given groups

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
    fromGroup := int32(56) // int32 | 
    toGroup := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockflowApi.GetBlockflowChainInfo(context.Background()).FromGroup(fromGroup).ToGroup(toGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockflowApi.GetBlockflowChainInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockflowChainInfo`: ChainInfo
    fmt.Fprintf(os.Stdout, "Response from `BlockflowApi.GetBlockflowChainInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowChainInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromGroup** | **int32** |  | 
 **toGroup** | **int32** |  | 

### Return type

[**ChainInfo**](ChainInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowHashes

> HashesAtHeight GetBlockflowHashes(ctx).FromGroup(fromGroup).ToGroup(toGroup).Height(height).Execute()

Get all block's hashes at given height for given groups

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
    fromGroup := int32(56) // int32 | 
    toGroup := int32(56) // int32 | 
    height := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockflowApi.GetBlockflowHashes(context.Background()).FromGroup(fromGroup).ToGroup(toGroup).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockflowApi.GetBlockflowHashes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockflowHashes`: HashesAtHeight
    fmt.Fprintf(os.Stdout, "Response from `BlockflowApi.GetBlockflowHashes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowHashesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromGroup** | **int32** |  | 
 **toGroup** | **int32** |  | 
 **height** | **int32** |  | 

### Return type

[**HashesAtHeight**](HashesAtHeight.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowHeadersBlockHash

> BlockHeaderEntry GetBlockflowHeadersBlockHash(ctx, blockHash).Execute()

Get block header

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockflowApi.GetBlockflowHeadersBlockHash(context.Background(), blockHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockflowApi.GetBlockflowHeadersBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockflowHeadersBlockHash`: BlockHeaderEntry
    fmt.Fprintf(os.Stdout, "Response from `BlockflowApi.GetBlockflowHeadersBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowHeadersBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockHeaderEntry**](BlockHeaderEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowIsBlockInMainChain

> bool GetBlockflowIsBlockInMainChain(ctx).BlockHash(blockHash).Execute()

Check if the block is in main chain

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockflowApi.GetBlockflowIsBlockInMainChain(context.Background()).BlockHash(blockHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockflowApi.GetBlockflowIsBlockInMainChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockflowIsBlockInMainChain`: bool
    fmt.Fprintf(os.Stdout, "Response from `BlockflowApi.GetBlockflowIsBlockInMainChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowIsBlockInMainChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockHash** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

