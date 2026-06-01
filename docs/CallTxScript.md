# CallTxScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **int32** |  | 
**Bytecode** | **string** |  | 
**CallerAddress** | Pointer to **string** |  | [optional] 
**WorldStateBlockHash** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**InputAssets** | Pointer to [**[]TestInputAsset**](TestInputAsset.md) |  | [optional] 
**InterestedContracts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCallTxScript

`func NewCallTxScript(group int32, bytecode string, ) *CallTxScript`

NewCallTxScript instantiates a new CallTxScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTxScriptWithDefaults

`func NewCallTxScriptWithDefaults() *CallTxScript`

NewCallTxScriptWithDefaults instantiates a new CallTxScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CallTxScript) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CallTxScript) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CallTxScript) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetBytecode

`func (o *CallTxScript) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *CallTxScript) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *CallTxScript) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetCallerAddress

`func (o *CallTxScript) GetCallerAddress() string`

GetCallerAddress returns the CallerAddress field if non-nil, zero value otherwise.

### GetCallerAddressOk

`func (o *CallTxScript) GetCallerAddressOk() (*string, bool)`

GetCallerAddressOk returns a tuple with the CallerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerAddress

`func (o *CallTxScript) SetCallerAddress(v string)`

SetCallerAddress sets CallerAddress field to given value.

### HasCallerAddress

`func (o *CallTxScript) HasCallerAddress() bool`

HasCallerAddress returns a boolean if a field has been set.

### GetWorldStateBlockHash

`func (o *CallTxScript) GetWorldStateBlockHash() string`

GetWorldStateBlockHash returns the WorldStateBlockHash field if non-nil, zero value otherwise.

### GetWorldStateBlockHashOk

`func (o *CallTxScript) GetWorldStateBlockHashOk() (*string, bool)`

GetWorldStateBlockHashOk returns a tuple with the WorldStateBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldStateBlockHash

`func (o *CallTxScript) SetWorldStateBlockHash(v string)`

SetWorldStateBlockHash sets WorldStateBlockHash field to given value.

### HasWorldStateBlockHash

`func (o *CallTxScript) HasWorldStateBlockHash() bool`

HasWorldStateBlockHash returns a boolean if a field has been set.

### GetTxId

`func (o *CallTxScript) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *CallTxScript) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *CallTxScript) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *CallTxScript) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetInputAssets

`func (o *CallTxScript) GetInputAssets() []TestInputAsset`

GetInputAssets returns the InputAssets field if non-nil, zero value otherwise.

### GetInputAssetsOk

`func (o *CallTxScript) GetInputAssetsOk() (*[]TestInputAsset, bool)`

GetInputAssetsOk returns a tuple with the InputAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAssets

`func (o *CallTxScript) SetInputAssets(v []TestInputAsset)`

SetInputAssets sets InputAssets field to given value.

### HasInputAssets

`func (o *CallTxScript) HasInputAssets() bool`

HasInputAssets returns a boolean if a field has been set.

### GetInterestedContracts

`func (o *CallTxScript) GetInterestedContracts() []string`

GetInterestedContracts returns the InterestedContracts field if non-nil, zero value otherwise.

### GetInterestedContractsOk

`func (o *CallTxScript) GetInterestedContractsOk() (*[]string, bool)`

GetInterestedContractsOk returns a tuple with the InterestedContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestedContracts

`func (o *CallTxScript) SetInterestedContracts(v []string)`

SetInterestedContracts sets InterestedContracts field to given value.

### HasInterestedContracts

`func (o *CallTxScript) HasInterestedContracts() bool`

HasInterestedContracts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


