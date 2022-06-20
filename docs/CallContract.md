# CallContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **int32** |  | 
**WorldStateBlockHash** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**Address** | **string** |  | 
**MethodIndex** | **int32** |  | 
**Args** | Pointer to [**[]Val**](Val.md) |  | [optional] 
**ExistingContracts** | Pointer to **[]string** |  | [optional] 
**InputAssets** | Pointer to [**[]TestInputAsset**](TestInputAsset.md) |  | [optional] 

## Methods

### NewCallContract

`func NewCallContract(group int32, address string, methodIndex int32, ) *CallContract`

NewCallContract instantiates a new CallContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallContractWithDefaults

`func NewCallContractWithDefaults() *CallContract`

NewCallContractWithDefaults instantiates a new CallContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CallContract) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CallContract) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CallContract) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetWorldStateBlockHash

`func (o *CallContract) GetWorldStateBlockHash() string`

GetWorldStateBlockHash returns the WorldStateBlockHash field if non-nil, zero value otherwise.

### GetWorldStateBlockHashOk

`func (o *CallContract) GetWorldStateBlockHashOk() (*string, bool)`

GetWorldStateBlockHashOk returns a tuple with the WorldStateBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldStateBlockHash

`func (o *CallContract) SetWorldStateBlockHash(v string)`

SetWorldStateBlockHash sets WorldStateBlockHash field to given value.

### HasWorldStateBlockHash

`func (o *CallContract) HasWorldStateBlockHash() bool`

HasWorldStateBlockHash returns a boolean if a field has been set.

### GetTxId

`func (o *CallContract) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *CallContract) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *CallContract) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *CallContract) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetAddress

`func (o *CallContract) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CallContract) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CallContract) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMethodIndex

`func (o *CallContract) GetMethodIndex() int32`

GetMethodIndex returns the MethodIndex field if non-nil, zero value otherwise.

### GetMethodIndexOk

`func (o *CallContract) GetMethodIndexOk() (*int32, bool)`

GetMethodIndexOk returns a tuple with the MethodIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIndex

`func (o *CallContract) SetMethodIndex(v int32)`

SetMethodIndex sets MethodIndex field to given value.


### GetArgs

`func (o *CallContract) GetArgs() []Val`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CallContract) GetArgsOk() (*[]Val, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CallContract) SetArgs(v []Val)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CallContract) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetExistingContracts

`func (o *CallContract) GetExistingContracts() []string`

GetExistingContracts returns the ExistingContracts field if non-nil, zero value otherwise.

### GetExistingContractsOk

`func (o *CallContract) GetExistingContractsOk() (*[]string, bool)`

GetExistingContractsOk returns a tuple with the ExistingContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingContracts

`func (o *CallContract) SetExistingContracts(v []string)`

SetExistingContracts sets ExistingContracts field to given value.

### HasExistingContracts

`func (o *CallContract) HasExistingContracts() bool`

HasExistingContracts returns a boolean if a field has been set.

### GetInputAssets

`func (o *CallContract) GetInputAssets() []TestInputAsset`

GetInputAssets returns the InputAssets field if non-nil, zero value otherwise.

### GetInputAssetsOk

`func (o *CallContract) GetInputAssetsOk() (*[]TestInputAsset, bool)`

GetInputAssetsOk returns a tuple with the InputAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAssets

`func (o *CallContract) SetInputAssets(v []TestInputAsset)`

SetInputAssets sets InputAssets field to given value.

### HasInputAssets

`func (o *CallContract) HasInputAssets() bool`

HasInputAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


