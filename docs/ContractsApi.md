# \ContractsAPI

All URIs are relative to *http://..*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractsAddressParent**](ContractsAPI.md#GetContractsAddressParent) | **Get** /contracts/{address}/parent | Get parent contract address
[**GetContractsAddressState**](ContractsAPI.md#GetContractsAddressState) | **Get** /contracts/{address}/state | Get contract state
[**GetContractsAddressSubContracts**](ContractsAPI.md#GetContractsAddressSubContracts) | **Get** /contracts/{address}/sub-contracts | Get sub-contract addresses
[**GetContractsAddressSubContractsCurrentCount**](ContractsAPI.md#GetContractsAddressSubContractsCurrentCount) | **Get** /contracts/{address}/sub-contracts/current-count | Get current value of the sub-contracts counter for a contract
[**GetContractsCodehashCode**](ContractsAPI.md#GetContractsCodehashCode) | **Get** /contracts/{codeHash}/code | Get contract code by code hash
[**PostContractsCallContract**](ContractsAPI.md#PostContractsCallContract) | **Post** /contracts/call-contract | Call contract
[**PostContractsCallTxScript**](ContractsAPI.md#PostContractsCallTxScript) | **Post** /contracts/call-tx-script | Call TxScript
[**PostContractsCompileContract**](ContractsAPI.md#PostContractsCompileContract) | **Post** /contracts/compile-contract | Compile a smart contract
[**PostContractsCompileProject**](ContractsAPI.md#PostContractsCompileProject) | **Post** /contracts/compile-project | Compile a project
[**PostContractsCompileScript**](ContractsAPI.md#PostContractsCompileScript) | **Post** /contracts/compile-script | Compile a script
[**PostContractsMulticallContract**](ContractsAPI.md#PostContractsMulticallContract) | **Post** /contracts/multicall-contract | Multiple call contract
[**PostContractsTestContract**](ContractsAPI.md#PostContractsTestContract) | **Post** /contracts/test-contract | Test contract
[**PostContractsUnsignedTxDeployContract**](ContractsAPI.md#PostContractsUnsignedTxDeployContract) | **Post** /contracts/unsigned-tx/deploy-contract | Build an unsigned contract
[**PostContractsUnsignedTxExecuteScript**](ContractsAPI.md#PostContractsUnsignedTxExecuteScript) | **Post** /contracts/unsigned-tx/execute-script | Build an unsigned script



## GetContractsAddressParent

> string GetContractsAddressParent(ctx, address).Execute()

Get parent contract address

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
	resp, r, err := apiClient.ContractsAPI.GetContractsAddressParent(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractsAddressParent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsAddressParent`: string
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractsAddressParent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsAddressParentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetContractsAddressState

> ContractState GetContractsAddressState(ctx, address).Execute()

Get contract state

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
	resp, r, err := apiClient.ContractsAPI.GetContractsAddressState(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractsAddressState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsAddressState`: ContractState
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractsAddressState`: %v\n", resp)
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


## GetContractsAddressSubContracts

> SubContracts GetContractsAddressSubContracts(ctx, address).Start(start).Limit(limit).Execute()

Get sub-contract addresses

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
	start := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.GetContractsAddressSubContracts(context.Background(), address).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractsAddressSubContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsAddressSubContracts`: SubContracts
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractsAddressSubContracts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsAddressSubContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**SubContracts**](SubContracts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractsAddressSubContractsCurrentCount

> int32 GetContractsAddressSubContractsCurrentCount(ctx, address).Execute()

Get current value of the sub-contracts counter for a contract

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
	resp, r, err := apiClient.ContractsAPI.GetContractsAddressSubContractsCurrentCount(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractsAddressSubContractsCurrentCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsAddressSubContractsCurrentCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractsAddressSubContractsCurrentCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsAddressSubContractsCurrentCountRequest struct via the builder pattern


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


## GetContractsCodehashCode

> string GetContractsCodehashCode(ctx, codeHash).Execute()

Get contract code by code hash

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
	codeHash := "codeHash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.GetContractsCodehashCode(context.Background(), codeHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractsCodehashCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractsCodehashCode`: string
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractsCodehashCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsCodehashCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	callContract := *openapiclient.NewCallContract(int32(123), "Address_example", int32(123)) // CallContract | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsCallContract(context.Background()).CallContract(callContract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsCallContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsCallContract`: CallContractResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsCallContract`: %v\n", resp)
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


## PostContractsCallTxScript

> CallTxScriptResult PostContractsCallTxScript(ctx).CallTxScript(callTxScript).Execute()

Call TxScript

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
	callTxScript := *openapiclient.NewCallTxScript(int32(123), "Bytecode_example") // CallTxScript | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsCallTxScript(context.Background()).CallTxScript(callTxScript).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsCallTxScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsCallTxScript`: CallTxScriptResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsCallTxScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsCallTxScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callTxScript** | [**CallTxScript**](CallTxScript.md) |  | 

### Return type

[**CallTxScriptResult**](CallTxScriptResult.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	contract := *openapiclient.NewContract("Code_example") // Contract | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsCompileContract(context.Background()).Contract(contract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsCompileContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsCompileContract`: CompileContractResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsCompileContract`: %v\n", resp)
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


## PostContractsCompileProject

> CompileProjectResult PostContractsCompileProject(ctx).Project(project).Execute()

Compile a project

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
	project := *openapiclient.NewProject("Code_example") // Project | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsCompileProject(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsCompileProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsCompileProject`: CompileProjectResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsCompileProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsCompileProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) |  | 

### Return type

[**CompileProjectResult**](CompileProjectResult.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	script := *openapiclient.NewScript("Code_example") // Script | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsCompileScript(context.Background()).Script(script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsCompileScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsCompileScript`: CompileScriptResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsCompileScript`: %v\n", resp)
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


## PostContractsMulticallContract

> MultipleCallContractResult PostContractsMulticallContract(ctx).MultipleCallContract(multipleCallContract).Execute()

Multiple call contract

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
	multipleCallContract := *openapiclient.NewMultipleCallContract([]openapiclient.CallContract{*openapiclient.NewCallContract(int32(123), "Address_example", int32(123))}) // MultipleCallContract | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsMulticallContract(context.Background()).MultipleCallContract(multipleCallContract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsMulticallContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsMulticallContract`: MultipleCallContractResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsMulticallContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContractsMulticallContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipleCallContract** | [**MultipleCallContract**](MultipleCallContract.md) |  | 

### Return type

[**MultipleCallContractResult**](MultipleCallContractResult.md)

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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	testContract := *openapiclient.NewTestContract("Bytecode_example") // TestContract | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsTestContract(context.Background()).TestContract(testContract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsTestContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsTestContract`: TestContractResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsTestContract`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	buildDeployContractTx := *openapiclient.NewBuildDeployContractTx("FromPublicKey_example", "Bytecode_example") // BuildDeployContractTx | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsUnsignedTxDeployContract(context.Background()).BuildDeployContractTx(buildDeployContractTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsUnsignedTxDeployContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsUnsignedTxDeployContract`: BuildDeployContractTxResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsUnsignedTxDeployContract`: %v\n", resp)
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
	openapiclient "github.com/alephium/go-sdk"
)

func main() {
	buildExecuteScriptTx := *openapiclient.NewBuildExecuteScriptTx("FromPublicKey_example", "Bytecode_example") // BuildExecuteScriptTx | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractsAPI.PostContractsUnsignedTxExecuteScript(context.Background()).BuildExecuteScriptTx(buildExecuteScriptTx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.PostContractsUnsignedTxExecuteScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContractsUnsignedTxExecuteScript`: BuildExecuteScriptTxResult
	fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.PostContractsUnsignedTxExecuteScript`: %v\n", resp)
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

