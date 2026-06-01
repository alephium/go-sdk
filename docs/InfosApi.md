# \InfosAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfosChainParams**](InfosAPI.md#GetInfosChainParams) | **Get** /infos/chain-params | Get key params about your blockchain
[**GetInfosCurrentDifficulty**](InfosAPI.md#GetInfosCurrentDifficulty) | **Get** /infos/current-difficulty | Get the average difficulty of the latest blocks from all shards
[**GetInfosCurrentHashrate**](InfosAPI.md#GetInfosCurrentHashrate) | **Get** /infos/current-hashrate | Get average hashrate from &#x60;now - timespan(millis)&#x60; to &#x60;now&#x60;
[**GetInfosDiscoveredNeighbors**](InfosAPI.md#GetInfosDiscoveredNeighbors) | **Get** /infos/discovered-neighbors | Get discovered neighbors
[**GetInfosHistoryHashrate**](InfosAPI.md#GetInfosHistoryHashrate) | **Get** /infos/history-hashrate | Get history average hashrate on the given time interval
[**GetInfosInterCliquePeerInfo**](InfosAPI.md#GetInfosInterCliquePeerInfo) | **Get** /infos/inter-clique-peer-info | Get infos about the inter cliques
[**GetInfosMisbehaviors**](InfosAPI.md#GetInfosMisbehaviors) | **Get** /infos/misbehaviors | Get the misbehaviors of peers
[**GetInfosNode**](InfosAPI.md#GetInfosNode) | **Get** /infos/node | Get info about that node
[**GetInfosSelfClique**](InfosAPI.md#GetInfosSelfClique) | **Get** /infos/self-clique | Get info about your own clique
[**GetInfosUnreachable**](InfosAPI.md#GetInfosUnreachable) | **Get** /infos/unreachable | Get the unreachable brokers
[**GetInfosVersion**](InfosAPI.md#GetInfosVersion) | **Get** /infos/version | Get version about that node
[**PostInfosDiscovery**](InfosAPI.md#PostInfosDiscovery) | **Post** /infos/discovery | Set brokers to be unreachable/reachable
[**PostInfosMisbehaviors**](InfosAPI.md#PostInfosMisbehaviors) | **Post** /infos/misbehaviors | Ban/Unban given peers



## GetInfosChainParams

> ChainParams GetInfosChainParams(ctx).Execute()

Get key params about your blockchain

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
	resp, r, err := apiClient.InfosAPI.GetInfosChainParams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosChainParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosChainParams`: ChainParams
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosChainParams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosChainParamsRequest struct via the builder pattern


### Return type

[**ChainParams**](ChainParams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosCurrentDifficulty

> CurrentDifficulty GetInfosCurrentDifficulty(ctx).Execute()

Get the average difficulty of the latest blocks from all shards

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
	resp, r, err := apiClient.InfosAPI.GetInfosCurrentDifficulty(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosCurrentDifficulty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosCurrentDifficulty`: CurrentDifficulty
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosCurrentDifficulty`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosCurrentDifficultyRequest struct via the builder pattern


### Return type

[**CurrentDifficulty**](CurrentDifficulty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosCurrentHashrate

> HashRateResponse GetInfosCurrentHashrate(ctx).Timespan(timespan).Execute()

Get average hashrate from `now - timespan(millis)` to `now`

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
	timespan := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfosAPI.GetInfosCurrentHashrate(context.Background()).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosCurrentHashrate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosCurrentHashrate`: HashRateResponse
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosCurrentHashrate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosCurrentHashrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timespan** | **int64** |  | 

### Return type

[**HashRateResponse**](HashRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosDiscoveredNeighbors

> []BrokerInfo GetInfosDiscoveredNeighbors(ctx).Execute()

Get discovered neighbors

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
	resp, r, err := apiClient.InfosAPI.GetInfosDiscoveredNeighbors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosDiscoveredNeighbors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosDiscoveredNeighbors`: []BrokerInfo
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosDiscoveredNeighbors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosDiscoveredNeighborsRequest struct via the builder pattern


### Return type

[**[]BrokerInfo**](BrokerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosHistoryHashrate

> HashRateResponse GetInfosHistoryHashrate(ctx).FromTs(fromTs).ToTs(toTs).Execute()

Get history average hashrate on the given time interval

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
	resp, r, err := apiClient.InfosAPI.GetInfosHistoryHashrate(context.Background()).FromTs(fromTs).ToTs(toTs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosHistoryHashrate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosHistoryHashrate`: HashRateResponse
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosHistoryHashrate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosHistoryHashrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromTs** | **int64** |  | 
 **toTs** | **int64** |  | 

### Return type

[**HashRateResponse**](HashRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosInterCliquePeerInfo

> []InterCliquePeerInfo GetInfosInterCliquePeerInfo(ctx).Execute()

Get infos about the inter cliques

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
	resp, r, err := apiClient.InfosAPI.GetInfosInterCliquePeerInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosInterCliquePeerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosInterCliquePeerInfo`: []InterCliquePeerInfo
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosInterCliquePeerInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosInterCliquePeerInfoRequest struct via the builder pattern


### Return type

[**[]InterCliquePeerInfo**](InterCliquePeerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosMisbehaviors

> []PeerMisbehavior GetInfosMisbehaviors(ctx).Execute()

Get the misbehaviors of peers

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
	resp, r, err := apiClient.InfosAPI.GetInfosMisbehaviors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosMisbehaviors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosMisbehaviors`: []PeerMisbehavior
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosMisbehaviors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosMisbehaviorsRequest struct via the builder pattern


### Return type

[**[]PeerMisbehavior**](PeerMisbehavior.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosNode

> NodeInfo GetInfosNode(ctx).Execute()

Get info about that node

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
	resp, r, err := apiClient.InfosAPI.GetInfosNode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosNode`: NodeInfo
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosNode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosNodeRequest struct via the builder pattern


### Return type

[**NodeInfo**](NodeInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosSelfClique

> SelfClique GetInfosSelfClique(ctx).Execute()

Get info about your own clique

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
	resp, r, err := apiClient.InfosAPI.GetInfosSelfClique(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosSelfClique``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosSelfClique`: SelfClique
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosSelfClique`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosSelfCliqueRequest struct via the builder pattern


### Return type

[**SelfClique**](SelfClique.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosUnreachable

> []string GetInfosUnreachable(ctx).Execute()

Get the unreachable brokers

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
	resp, r, err := apiClient.InfosAPI.GetInfosUnreachable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosUnreachable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosUnreachable`: []string
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosUnreachable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosUnreachableRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfosVersion

> NodeVersion GetInfosVersion(ctx).Execute()

Get version about that node

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
	resp, r, err := apiClient.InfosAPI.GetInfosVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.GetInfosVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfosVersion`: NodeVersion
	fmt.Fprintf(os.Stdout, "Response from `InfosAPI.GetInfosVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosVersionRequest struct via the builder pattern


### Return type

[**NodeVersion**](NodeVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInfosDiscovery

> PostInfosDiscovery(ctx).DiscoveryAction(discoveryAction).Execute()

Set brokers to be unreachable/reachable

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
	discoveryAction := openapiclient.DiscoveryAction{Reachable: openapiclient.NewReachable([]string{"Peers_example"}, "Type_example")} // DiscoveryAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfosAPI.PostInfosDiscovery(context.Background()).DiscoveryAction(discoveryAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.PostInfosDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostInfosDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryAction** | [**DiscoveryAction**](DiscoveryAction.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInfosMisbehaviors

> PostInfosMisbehaviors(ctx).MisbehaviorAction(misbehaviorAction).Execute()

Ban/Unban given peers

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
	misbehaviorAction := openapiclient.MisbehaviorAction{Ban: openapiclient.NewBan([]string{"Peers_example"}, "Type_example")} // MisbehaviorAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InfosAPI.PostInfosMisbehaviors(context.Background()).MisbehaviorAction(misbehaviorAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfosAPI.PostInfosMisbehaviors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostInfosMisbehaviorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **misbehaviorAction** | [**MisbehaviorAction**](MisbehaviorAction.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

