# \InfosApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfosChainParams**](InfosApi.md#GetInfosChainParams) | **Get** /infos/chain-params | Get key params about your blockchain
[**GetInfosCurrentHashrate**](InfosApi.md#GetInfosCurrentHashrate) | **Get** /infos/current-hashrate | Get average hashrate from &#x60;now - timespan(millis)&#x60; to &#x60;now&#x60;
[**GetInfosDiscoveredNeighbors**](InfosApi.md#GetInfosDiscoveredNeighbors) | **Get** /infos/discovered-neighbors | Get discovered neighbors
[**GetInfosHistoryHashrate**](InfosApi.md#GetInfosHistoryHashrate) | **Get** /infos/history-hashrate | Get history average hashrate on the given time interval
[**GetInfosInterCliquePeerInfo**](InfosApi.md#GetInfosInterCliquePeerInfo) | **Get** /infos/inter-clique-peer-info | Get infos about the inter cliques
[**GetInfosMisbehaviors**](InfosApi.md#GetInfosMisbehaviors) | **Get** /infos/misbehaviors | Get the misbehaviors of peers
[**GetInfosNode**](InfosApi.md#GetInfosNode) | **Get** /infos/node | Get info about that node
[**GetInfosSelfClique**](InfosApi.md#GetInfosSelfClique) | **Get** /infos/self-clique | Get info about your own clique
[**GetInfosUnreachable**](InfosApi.md#GetInfosUnreachable) | **Get** /infos/unreachable | Get the unreachable brokers
[**GetInfosVersion**](InfosApi.md#GetInfosVersion) | **Get** /infos/version | Get version about that node
[**PostInfosDiscovery**](InfosApi.md#PostInfosDiscovery) | **Post** /infos/discovery | Set brokers to be unreachable/reachable
[**PostInfosMisbehaviors**](InfosApi.md#PostInfosMisbehaviors) | **Post** /infos/misbehaviors | Ban/Unban given peers



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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosChainParams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosChainParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosChainParams`: ChainParams
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosChainParams`: %v\n", resp)
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


## GetInfosCurrentHashrate

> string GetInfosCurrentHashrate(ctx).Timespan(timespan).Execute()

Get average hashrate from `now - timespan(millis)` to `now`

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
    timespan := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosCurrentHashrate(context.Background()).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosCurrentHashrate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosCurrentHashrate`: string
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosCurrentHashrate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfosCurrentHashrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timespan** | **int64** |  | 

### Return type

**string**

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosDiscoveredNeighbors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosDiscoveredNeighbors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosDiscoveredNeighbors`: []BrokerInfo
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosDiscoveredNeighbors`: %v\n", resp)
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

> string GetInfosHistoryHashrate(ctx).FromTs(fromTs).ToTs(toTs).Execute()

Get history average hashrate on the given time interval

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
    resp, r, err := apiClient.InfosApi.GetInfosHistoryHashrate(context.Background()).FromTs(fromTs).ToTs(toTs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosHistoryHashrate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosHistoryHashrate`: string
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosHistoryHashrate`: %v\n", resp)
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

**string**

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosInterCliquePeerInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosInterCliquePeerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosInterCliquePeerInfo`: []InterCliquePeerInfo
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosInterCliquePeerInfo`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosMisbehaviors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosMisbehaviors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosMisbehaviors`: []PeerMisbehavior
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosMisbehaviors`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosNode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosNode`: NodeInfo
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosNode`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosSelfClique(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosSelfClique``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosSelfClique`: SelfClique
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosSelfClique`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosUnreachable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosUnreachable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosUnreachable`: []string
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosUnreachable`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.GetInfosVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.GetInfosVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfosVersion`: NodeVersion
    fmt.Fprintf(os.Stdout, "Response from `InfosApi.GetInfosVersion`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    discoveryAction := openapiclient.DiscoveryAction{Reachable: openapiclient.NewReachable([]string{"Peers_example"}, "Type_example")} // DiscoveryAction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.PostInfosDiscovery(context.Background()).DiscoveryAction(discoveryAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.PostInfosDiscovery``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    misbehaviorAction := openapiclient.MisbehaviorAction{Ban: openapiclient.NewBan([]string{"Peers_example"}, "Type_example")} // MisbehaviorAction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfosApi.PostInfosMisbehaviors(context.Background()).MisbehaviorAction(misbehaviorAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfosApi.PostInfosMisbehaviors``: %v\n", err)
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

