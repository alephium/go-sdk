# BuildGrouplessDeployContractTxResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 
**UnsignedTx** | **string** |  | 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 
**TxId** | **string** |  | 
**ContractAddress** | **string** |  | 
**FundingTxs** | Pointer to [**[]BuildSimpleTransferTxResult**](BuildSimpleTransferTxResult.md) |  | [optional] 

## Methods

### NewBuildGrouplessDeployContractTxResult

`func NewBuildGrouplessDeployContractTxResult(fromGroup int32, toGroup int32, unsignedTx string, gasAmount int32, gasPrice string, txId string, contractAddress string, ) *BuildGrouplessDeployContractTxResult`

NewBuildGrouplessDeployContractTxResult instantiates a new BuildGrouplessDeployContractTxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildGrouplessDeployContractTxResultWithDefaults

`func NewBuildGrouplessDeployContractTxResultWithDefaults() *BuildGrouplessDeployContractTxResult`

NewBuildGrouplessDeployContractTxResultWithDefaults instantiates a new BuildGrouplessDeployContractTxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromGroup

`func (o *BuildGrouplessDeployContractTxResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildGrouplessDeployContractTxResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildGrouplessDeployContractTxResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildGrouplessDeployContractTxResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildGrouplessDeployContractTxResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildGrouplessDeployContractTxResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetUnsignedTx

`func (o *BuildGrouplessDeployContractTxResult) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *BuildGrouplessDeployContractTxResult) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *BuildGrouplessDeployContractTxResult) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *BuildGrouplessDeployContractTxResult) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildGrouplessDeployContractTxResult) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildGrouplessDeployContractTxResult) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *BuildGrouplessDeployContractTxResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildGrouplessDeployContractTxResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildGrouplessDeployContractTxResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetTxId

`func (o *BuildGrouplessDeployContractTxResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *BuildGrouplessDeployContractTxResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *BuildGrouplessDeployContractTxResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractAddress

`func (o *BuildGrouplessDeployContractTxResult) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *BuildGrouplessDeployContractTxResult) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *BuildGrouplessDeployContractTxResult) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetFundingTxs

`func (o *BuildGrouplessDeployContractTxResult) GetFundingTxs() []BuildSimpleTransferTxResult`

GetFundingTxs returns the FundingTxs field if non-nil, zero value otherwise.

### GetFundingTxsOk

`func (o *BuildGrouplessDeployContractTxResult) GetFundingTxsOk() (*[]BuildSimpleTransferTxResult, bool)`

GetFundingTxsOk returns a tuple with the FundingTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTxs

`func (o *BuildGrouplessDeployContractTxResult) SetFundingTxs(v []BuildSimpleTransferTxResult)`

SetFundingTxs sets FundingTxs field to given value.

### HasFundingTxs

`func (o *BuildGrouplessDeployContractTxResult) HasFundingTxs() bool`

HasFundingTxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


