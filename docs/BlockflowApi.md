# \BlockflowAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockflowBlocks**](BlockflowAPI.md#GetBlockflowBlocks) | **Get** /blockflow/blocks | List blocks on the given time interval
[**GetBlockflowBlocksBlockHash**](BlockflowAPI.md#GetBlockflowBlocksBlockHash) | **Get** /blockflow/blocks/{block_hash} | Get a block with hash
[**GetBlockflowBlocksWithEvents**](BlockflowAPI.md#GetBlockflowBlocksWithEvents) | **Get** /blockflow/blocks-with-events | List blocks with events on the given time interval
[**GetBlockflowBlocksWithEventsBlockHash**](BlockflowAPI.md#GetBlockflowBlocksWithEventsBlockHash) | **Get** /blockflow/blocks-with-events/{block_hash} | Get a block and events with hash
[**GetBlockflowChainInfo**](BlockflowAPI.md#GetBlockflowChainInfo) | **Get** /blockflow/chain-info | Get infos about the chain from the given groups
[**GetBlockflowHashes**](BlockflowAPI.md#GetBlockflowHashes) | **Get** /blockflow/hashes | Get all block&#39;s hashes at given height for given groups
[**GetBlockflowHeadersBlockHash**](BlockflowAPI.md#GetBlockflowHeadersBlockHash) | **Get** /blockflow/headers/{block_hash} | Get block header
[**GetBlockflowIsBlockInMainChain**](BlockflowAPI.md#GetBlockflowIsBlockInMainChain) | **Get** /blockflow/is-block-in-main-chain | Check if the block is in main chain
[**GetBlockflowMainChainBlockByGhostUncleGhostUncleHash**](BlockflowAPI.md#GetBlockflowMainChainBlockByGhostUncleGhostUncleHash) | **Get** /blockflow/main-chain-block-by-ghost-uncle/{ghost_uncle_hash} | Get a mainchain block by ghost uncle hash
[**GetBlockflowRawBlocksBlockHash**](BlockflowAPI.md#GetBlockflowRawBlocksBlockHash) | **Get** /blockflow/raw-blocks/{block_hash} | Get raw block in hex format
[**GetBlockflowRichBlocks**](BlockflowAPI.md#GetBlockflowRichBlocks) | **Get** /blockflow/rich-blocks | Given a time interval, list blocks containing events and transactions with enriched input information when node indexes are enabled.
[**GetBlockflowRichBlocksBlockHash**](BlockflowAPI.md#GetBlockflowRichBlocksBlockHash) | **Get** /blockflow/rich-blocks/{block_hash} | Get a block containing events and transactions with enriched input information when node indexes are enabled.



## GetBlockflowBlocks

> BlocksPerTimeStampRange GetBlockflowBlocks(ctx).FromTs(fromTs).ToTs(toTs).Execute()

List blocks on the given time interval

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
	fromTs := int64(789) // int64 | 
	toTs := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowBlocks(context.Background()).FromTs(fromTs).ToTs(toTs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowBlocks`: BlocksPerTimeStampRange
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromTs** | **int64** |  | 
 **toTs** | **int64** |  | 

### Return type

[**BlocksPerTimeStampRange**](BlocksPerTimeStampRange.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	blockHash := "blockHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowBlocksBlockHash(context.Background(), blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowBlocksBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowBlocksBlockHash`: BlockEntry
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowBlocksBlockHash`: %v\n", resp)
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


## GetBlockflowBlocksWithEvents

> BlocksAndEventsPerTimeStampRange GetBlockflowBlocksWithEvents(ctx).FromTs(fromTs).ToTs(toTs).Execute()

List blocks with events on the given time interval

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
	fromTs := int64(789) // int64 | 
	toTs := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowBlocksWithEvents(context.Background()).FromTs(fromTs).ToTs(toTs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowBlocksWithEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowBlocksWithEvents`: BlocksAndEventsPerTimeStampRange
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowBlocksWithEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowBlocksWithEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromTs** | **int64** |  | 
 **toTs** | **int64** |  | 

### Return type

[**BlocksAndEventsPerTimeStampRange**](BlocksAndEventsPerTimeStampRange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowBlocksWithEventsBlockHash

> BlockAndEvents GetBlockflowBlocksWithEventsBlockHash(ctx, blockHash).Execute()

Get a block and events with hash

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
	blockHash := "blockHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowBlocksWithEventsBlockHash(context.Background(), blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowBlocksWithEventsBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowBlocksWithEventsBlockHash`: BlockAndEvents
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowBlocksWithEventsBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowBlocksWithEventsBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockAndEvents**](BlockAndEvents.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	fromGroup := int32(56) // int32 | 
	toGroup := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowChainInfo(context.Background()).FromGroup(fromGroup).ToGroup(toGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowChainInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowChainInfo`: ChainInfo
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowChainInfo`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	fromGroup := int32(56) // int32 | 
	toGroup := int32(56) // int32 | 
	height := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowHashes(context.Background()).FromGroup(fromGroup).ToGroup(toGroup).Height(height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowHashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowHashes`: HashesAtHeight
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowHashes`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	blockHash := "blockHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowHeadersBlockHash(context.Background(), blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowHeadersBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowHeadersBlockHash`: BlockHeaderEntry
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowHeadersBlockHash`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	blockHash := "blockHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowIsBlockInMainChain(context.Background()).BlockHash(blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowIsBlockInMainChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowIsBlockInMainChain`: bool
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowIsBlockInMainChain`: %v\n", resp)
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


## GetBlockflowMainChainBlockByGhostUncleGhostUncleHash

> BlockEntry GetBlockflowMainChainBlockByGhostUncleGhostUncleHash(ctx, ghostUncleHash).Execute()

Get a mainchain block by ghost uncle hash

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
	ghostUncleHash := "ghostUncleHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowMainChainBlockByGhostUncleGhostUncleHash(context.Background(), ghostUncleHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowMainChainBlockByGhostUncleGhostUncleHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowMainChainBlockByGhostUncleGhostUncleHash`: BlockEntry
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowMainChainBlockByGhostUncleGhostUncleHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ghostUncleHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowMainChainBlockByGhostUncleGhostUncleHashRequest struct via the builder pattern


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


## GetBlockflowRawBlocksBlockHash

> RawBlock GetBlockflowRawBlocksBlockHash(ctx, blockHash).Execute()

Get raw block in hex format

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
	blockHash := "blockHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowRawBlocksBlockHash(context.Background(), blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowRawBlocksBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowRawBlocksBlockHash`: RawBlock
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowRawBlocksBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowRawBlocksBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RawBlock**](RawBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowRichBlocks

> RichBlocksAndEventsPerTimeStampRange GetBlockflowRichBlocks(ctx).FromTs(fromTs).ToTs(toTs).Execute()

Given a time interval, list blocks containing events and transactions with enriched input information when node indexes are enabled.

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
	fromTs := int64(789) // int64 | 
	toTs := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowRichBlocks(context.Background()).FromTs(fromTs).ToTs(toTs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowRichBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowRichBlocks`: RichBlocksAndEventsPerTimeStampRange
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowRichBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowRichBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromTs** | **int64** |  | 
 **toTs** | **int64** |  | 

### Return type

[**RichBlocksAndEventsPerTimeStampRange**](RichBlocksAndEventsPerTimeStampRange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockflowRichBlocksBlockHash

> RichBlockAndEvents GetBlockflowRichBlocksBlockHash(ctx, blockHash).Execute()

Get a block containing events and transactions with enriched input information when node indexes are enabled.

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
	blockHash := "blockHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockflowAPI.GetBlockflowRichBlocksBlockHash(context.Background(), blockHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockflowAPI.GetBlockflowRichBlocksBlockHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockflowRichBlocksBlockHash`: RichBlockAndEvents
	fmt.Fprintf(os.Stdout, "Response from `BlockflowAPI.GetBlockflowRichBlocksBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockflowRichBlocksBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RichBlockAndEvents**](RichBlockAndEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

