# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unsigned** | [**UnsignedTx**](UnsignedTx.md) |  | 
**ScriptExecutionOk** | **bool** |  | 
**ContractInputs** | [**[]OutputRef**](OutputRef.md) |  | 
**GeneratedOutputs** | [**[]Output**](Output.md) |  | 
**InputSignatures** | **[]string** |  | 
**ScriptSignatures** | **[]string** |  | 

## Methods

### NewTransaction

`func NewTransaction(unsigned UnsignedTx, scriptExecutionOk bool, contractInputs []OutputRef, generatedOutputs []Output, inputSignatures []string, scriptSignatures []string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsigned

`func (o *Transaction) GetUnsigned() UnsignedTx`

GetUnsigned returns the Unsigned field if non-nil, zero value otherwise.

### GetUnsignedOk

`func (o *Transaction) GetUnsignedOk() (*UnsignedTx, bool)`

GetUnsignedOk returns a tuple with the Unsigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsigned

`func (o *Transaction) SetUnsigned(v UnsignedTx)`

SetUnsigned sets Unsigned field to given value.


### GetScriptExecutionOk

`func (o *Transaction) GetScriptExecutionOk() bool`

GetScriptExecutionOk returns the ScriptExecutionOk field if non-nil, zero value otherwise.

### GetScriptExecutionOkOk

`func (o *Transaction) GetScriptExecutionOkOk() (*bool, bool)`

GetScriptExecutionOkOk returns a tuple with the ScriptExecutionOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExecutionOk

`func (o *Transaction) SetScriptExecutionOk(v bool)`

SetScriptExecutionOk sets ScriptExecutionOk field to given value.


### GetContractInputs

`func (o *Transaction) GetContractInputs() []OutputRef`

GetContractInputs returns the ContractInputs field if non-nil, zero value otherwise.

### GetContractInputsOk

`func (o *Transaction) GetContractInputsOk() (*[]OutputRef, bool)`

GetContractInputsOk returns a tuple with the ContractInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractInputs

`func (o *Transaction) SetContractInputs(v []OutputRef)`

SetContractInputs sets ContractInputs field to given value.


### GetGeneratedOutputs

`func (o *Transaction) GetGeneratedOutputs() []Output`

GetGeneratedOutputs returns the GeneratedOutputs field if non-nil, zero value otherwise.

### GetGeneratedOutputsOk

`func (o *Transaction) GetGeneratedOutputsOk() (*[]Output, bool)`

GetGeneratedOutputsOk returns a tuple with the GeneratedOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedOutputs

`func (o *Transaction) SetGeneratedOutputs(v []Output)`

SetGeneratedOutputs sets GeneratedOutputs field to given value.


### GetInputSignatures

`func (o *Transaction) GetInputSignatures() []string`

GetInputSignatures returns the InputSignatures field if non-nil, zero value otherwise.

### GetInputSignaturesOk

`func (o *Transaction) GetInputSignaturesOk() (*[]string, bool)`

GetInputSignaturesOk returns a tuple with the InputSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSignatures

`func (o *Transaction) SetInputSignatures(v []string)`

SetInputSignatures sets InputSignatures field to given value.


### GetScriptSignatures

`func (o *Transaction) GetScriptSignatures() []string`

GetScriptSignatures returns the ScriptSignatures field if non-nil, zero value otherwise.

### GetScriptSignaturesOk

`func (o *Transaction) GetScriptSignaturesOk() (*[]string, bool)`

GetScriptSignaturesOk returns a tuple with the ScriptSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSignatures

`func (o *Transaction) SetScriptSignatures(v []string)`

SetScriptSignatures sets ScriptSignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


