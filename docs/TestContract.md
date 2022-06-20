# TestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **int32** |  | [optional] 
**BlockHash** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Bytecode** | **string** |  | 
**InitialFields** | Pointer to [**[]Val**](Val.md) |  | [optional] 
**InitialAsset** | Pointer to [**AssetState**](AssetState.md) |  | [optional] 
**MethodIndex** | Pointer to **int32** |  | [optional] 
**Args** | Pointer to [**[]Val**](Val.md) |  | [optional] 
**ExistingContracts** | Pointer to [**[]ContractState**](ContractState.md) |  | [optional] 
**InputAssets** | Pointer to [**[]TestInputAsset**](TestInputAsset.md) |  | [optional] 

## Methods

### NewTestContract

`func NewTestContract(bytecode string, ) *TestContract`

NewTestContract instantiates a new TestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestContractWithDefaults

`func NewTestContractWithDefaults() *TestContract`

NewTestContractWithDefaults instantiates a new TestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *TestContract) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TestContract) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TestContract) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TestContract) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBlockHash

`func (o *TestContract) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TestContract) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TestContract) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TestContract) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetTxId

`func (o *TestContract) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TestContract) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TestContract) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *TestContract) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetAddress

`func (o *TestContract) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TestContract) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TestContract) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TestContract) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBytecode

`func (o *TestContract) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *TestContract) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *TestContract) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetInitialFields

`func (o *TestContract) GetInitialFields() []Val`

GetInitialFields returns the InitialFields field if non-nil, zero value otherwise.

### GetInitialFieldsOk

`func (o *TestContract) GetInitialFieldsOk() (*[]Val, bool)`

GetInitialFieldsOk returns a tuple with the InitialFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFields

`func (o *TestContract) SetInitialFields(v []Val)`

SetInitialFields sets InitialFields field to given value.

### HasInitialFields

`func (o *TestContract) HasInitialFields() bool`

HasInitialFields returns a boolean if a field has been set.

### GetInitialAsset

`func (o *TestContract) GetInitialAsset() AssetState`

GetInitialAsset returns the InitialAsset field if non-nil, zero value otherwise.

### GetInitialAssetOk

`func (o *TestContract) GetInitialAssetOk() (*AssetState, bool)`

GetInitialAssetOk returns a tuple with the InitialAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAsset

`func (o *TestContract) SetInitialAsset(v AssetState)`

SetInitialAsset sets InitialAsset field to given value.

### HasInitialAsset

`func (o *TestContract) HasInitialAsset() bool`

HasInitialAsset returns a boolean if a field has been set.

### GetMethodIndex

`func (o *TestContract) GetMethodIndex() int32`

GetMethodIndex returns the MethodIndex field if non-nil, zero value otherwise.

### GetMethodIndexOk

`func (o *TestContract) GetMethodIndexOk() (*int32, bool)`

GetMethodIndexOk returns a tuple with the MethodIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIndex

`func (o *TestContract) SetMethodIndex(v int32)`

SetMethodIndex sets MethodIndex field to given value.

### HasMethodIndex

`func (o *TestContract) HasMethodIndex() bool`

HasMethodIndex returns a boolean if a field has been set.

### GetArgs

`func (o *TestContract) GetArgs() []Val`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *TestContract) GetArgsOk() (*[]Val, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *TestContract) SetArgs(v []Val)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *TestContract) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetExistingContracts

`func (o *TestContract) GetExistingContracts() []ContractState`

GetExistingContracts returns the ExistingContracts field if non-nil, zero value otherwise.

### GetExistingContractsOk

`func (o *TestContract) GetExistingContractsOk() (*[]ContractState, bool)`

GetExistingContractsOk returns a tuple with the ExistingContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingContracts

`func (o *TestContract) SetExistingContracts(v []ContractState)`

SetExistingContracts sets ExistingContracts field to given value.

### HasExistingContracts

`func (o *TestContract) HasExistingContracts() bool`

HasExistingContracts returns a boolean if a field has been set.

### GetInputAssets

`func (o *TestContract) GetInputAssets() []TestInputAsset`

GetInputAssets returns the InputAssets field if non-nil, zero value otherwise.

### GetInputAssetsOk

`func (o *TestContract) GetInputAssetsOk() (*[]TestInputAsset, bool)`

GetInputAssetsOk returns a tuple with the InputAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAssets

`func (o *TestContract) SetInputAssets(v []TestInputAsset)`

SetInputAssets sets InputAssets field to given value.

### HasInputAssets

`func (o *TestContract) HasInputAssets() bool`

HasInputAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


