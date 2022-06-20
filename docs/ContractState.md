# ContractState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Bytecode** | **string** |  | 
**CodeHash** | **string** |  | 
**InitialStateHash** | Pointer to **string** |  | [optional] 
**Fields** | [**[]Val**](Val.md) |  | 
**Asset** | [**AssetState**](AssetState.md) |  | 

## Methods

### NewContractState

`func NewContractState(address string, bytecode string, codeHash string, fields []Val, asset AssetState, ) *ContractState`

NewContractState instantiates a new ContractState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractStateWithDefaults

`func NewContractStateWithDefaults() *ContractState`

NewContractStateWithDefaults instantiates a new ContractState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractState) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractState) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractState) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBytecode

`func (o *ContractState) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *ContractState) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *ContractState) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetCodeHash

`func (o *ContractState) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *ContractState) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *ContractState) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.


### GetInitialStateHash

`func (o *ContractState) GetInitialStateHash() string`

GetInitialStateHash returns the InitialStateHash field if non-nil, zero value otherwise.

### GetInitialStateHashOk

`func (o *ContractState) GetInitialStateHashOk() (*string, bool)`

GetInitialStateHashOk returns a tuple with the InitialStateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStateHash

`func (o *ContractState) SetInitialStateHash(v string)`

SetInitialStateHash sets InitialStateHash field to given value.

### HasInitialStateHash

`func (o *ContractState) HasInitialStateHash() bool`

HasInitialStateHash returns a boolean if a field has been set.

### GetFields

`func (o *ContractState) GetFields() []Val`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContractState) GetFieldsOk() (*[]Val, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContractState) SetFields(v []Val)`

SetFields sets Fields field to given value.


### GetAsset

`func (o *ContractState) GetAsset() AssetState`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *ContractState) GetAssetOk() (*AssetState, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *ContractState) SetAsset(v AssetState)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


