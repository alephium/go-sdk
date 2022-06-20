# \ContractsApi

All URIs are relative to *http://127.0.0.1:12973*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractsAddressState**](ContractsApi.md#GetContractsAddressState) | **Get** /contracts/{address}/state | Get contract state
[**PostContractsCallContract**](ContractsApi.md#PostContractsCallContract) | **Post** /contracts/call-contract | Call contract
[**PostContractsCompileContract**](ContractsApi.md#PostContractsCompileContract) | **Post** /contracts/compile-contract | Compile a smart contract
[**PostContractsCompileScript**](ContractsApi.md#PostContractsCompileScript) | **Post** /contracts/compile-script | Compile a script
[**PostContractsTestContract**](ContractsApi.md#PostContractsTestContract) | **Post** /contracts/test-contract | Test contract
[**PostContractsUnsignedTxDeployContract**](ContractsApi.md#PostContractsUnsignedTxDeployContract) | **Post** /contracts/unsigned-tx/deploy-contract | Build an unsigned contract
[**PostContractsUnsignedTxExecuteScript**](ContractsApi.md#PostContractsUnsignedTxExecuteScript) | **Post** /contracts/unsigned-tx/execute-script | Build an unsigned script



## GetContractsAddressState

> ContractState GetContractsAddressState(ctx, address).Group(group).Execute()

Get contract state

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
    address := "address_example" // string | 
    group := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetContractsAddressState(context.Background(), address).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContractsAddressState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractsAddressState`: ContractState
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContractsAddressState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsAddressStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **int32** |  | 

### Return type

[**ContractState**](ContractState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContractsCallContract

> CallContractResult PostContractsCallContract(ctx).CallContract(callContract).Execute()

Call contract

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
    callContract := *openapiclient.NewCallContract(int32(123), "Address_example", int32(123)) // CallContract | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.PostContractsCallContract(context.Background()).CallContract(callContract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.PostContractsCallContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostContractsCallContract`: CallContractResult
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.PostContractsCallContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsCallContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callContract** | [**CallContract**](CallContract.md) |  | 

### Return type

[**CallContractResult**](CallContractResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContractsCompileContract

> CompileContractResult PostContractsCompileContract(ctx).Contract(contract).Execute()

Compile a smart contract

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
    contract := *openapiclient.NewContract("Code_example") // Contract | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.PostContractsCompileContract(context.Background()).Contract(contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.PostContractsCompileContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostContractsCompileContract`: CompileContractResult
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.PostContractsCompileContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsCompileContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contract** | [**Contract**](Contract.md) |  | 

### Return type

[**CompileContractResult**](CompileContractResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContractsCompileScript

> CompileScriptResult PostContractsCompileScript(ctx).Script(script).Execute()

Compile a script

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
    script := *openapiclient.NewScript("Code_example") // Script | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.PostContractsCompileScript(context.Background()).Script(script).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.PostContractsCompileScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostContractsCompileScript`: CompileScriptResult
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.PostContractsCompileScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsCompileScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **script** | [**Script**](Script.md) |  | 

### Return type

[**CompileScriptResult**](CompileScriptResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContractsTestContract

> TestContractResult PostContractsTestContract(ctx).TestContract(testContract).Execute()

Test contract

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
    testContract := *openapiclient.NewTestContract("Bytecode_example") // TestContract | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.PostContractsTestContract(context.Background()).TestContract(testContract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.PostContractsTestContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostContractsTestContract`: TestContractResult
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.PostContractsTestContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsTestContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testContract** | [**TestContract**](TestContract.md) |  | 

### Return type

[**TestContractResult**](TestContractResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContractsUnsignedTxDeployContract

> BuildDeployContractTxResult PostContractsUnsignedTxDeployContract(ctx).BuildDeployContractTx(buildDeployContractTx).Execute()

Build an unsigned contract

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
    buildDeployContractTx := *openapiclient.NewBuildDeployContractTx("FromPublicKey_example", "Bytecode_example") // BuildDeployContractTx | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.PostContractsUnsignedTxDeployContract(context.Background()).BuildDeployContractTx(buildDeployContractTx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.PostContractsUnsignedTxDeployContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostContractsUnsignedTxDeployContract`: BuildDeployContractTxResult
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.PostContractsUnsignedTxDeployContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsUnsignedTxDeployContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildDeployContractTx** | [**BuildDeployContractTx**](BuildDeployContractTx.md) |  | 

### Return type

[**BuildDeployContractTxResult**](BuildDeployContractTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContractsUnsignedTxExecuteScript

> BuildExecuteScriptTxResult PostContractsUnsignedTxExecuteScript(ctx).BuildExecuteScriptTx(buildExecuteScriptTx).Execute()

Build an unsigned script

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
    buildExecuteScriptTx := *openapiclient.NewBuildExecuteScriptTx("FromPublicKey_example", "Bytecode_example") // BuildExecuteScriptTx | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.PostContractsUnsignedTxExecuteScript(context.Background()).BuildExecuteScriptTx(buildExecuteScriptTx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.PostContractsUnsignedTxExecuteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostContractsUnsignedTxExecuteScript`: BuildExecuteScriptTxResult
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.PostContractsUnsignedTxExecuteScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsUnsignedTxExecuteScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildExecuteScriptTx** | [**BuildExecuteScriptTx**](BuildExecuteScriptTx.md) |  | 

### Return type

[**BuildExecuteScriptTxResult**](BuildExecuteScriptTxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

