# RichTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unsigned** | [**RichUnsignedTx**](RichUnsignedTx.md) |  | 
**ScriptExecutionOk** | **bool** |  | 
**ContractInputs** | [**[]RichContractInput**](RichContractInput.md) |  | 
**GeneratedOutputs** | [**[]Output**](Output.md) |  | 
**InputSignatures** | **[]string** |  | 
**ScriptSignatures** | **[]string** |  | 

## Methods

### NewRichTransaction

`func NewRichTransaction(unsigned RichUnsignedTx, scriptExecutionOk bool, contractInputs []RichContractInput, generatedOutputs []Output, inputSignatures []string, scriptSignatures []string, ) *RichTransaction`

NewRichTransaction instantiates a new RichTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichTransactionWithDefaults

`func NewRichTransactionWithDefaults() *RichTransaction`

NewRichTransactionWithDefaults instantiates a new RichTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsigned

`func (o *RichTransaction) GetUnsigned() RichUnsignedTx`

GetUnsigned returns the Unsigned field if non-nil, zero value otherwise.

### GetUnsignedOk

`func (o *RichTransaction) GetUnsignedOk() (*RichUnsignedTx, bool)`

GetUnsignedOk returns a tuple with the Unsigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsigned

`func (o *RichTransaction) SetUnsigned(v RichUnsignedTx)`

SetUnsigned sets Unsigned field to given value.


### GetScriptExecutionOk

`func (o *RichTransaction) GetScriptExecutionOk() bool`

GetScriptExecutionOk returns the ScriptExecutionOk field if non-nil, zero value otherwise.

### GetScriptExecutionOkOk

`func (o *RichTransaction) GetScriptExecutionOkOk() (*bool, bool)`

GetScriptExecutionOkOk returns a tuple with the ScriptExecutionOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExecutionOk

`func (o *RichTransaction) SetScriptExecutionOk(v bool)`

SetScriptExecutionOk sets ScriptExecutionOk field to given value.


### GetContractInputs

`func (o *RichTransaction) GetContractInputs() []RichContractInput`

GetContractInputs returns the ContractInputs field if non-nil, zero value otherwise.

### GetContractInputsOk

`func (o *RichTransaction) GetContractInputsOk() (*[]RichContractInput, bool)`

GetContractInputsOk returns a tuple with the ContractInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractInputs

`func (o *RichTransaction) SetContractInputs(v []RichContractInput)`

SetContractInputs sets ContractInputs field to given value.


### GetGeneratedOutputs

`func (o *RichTransaction) GetGeneratedOutputs() []Output`

GetGeneratedOutputs returns the GeneratedOutputs field if non-nil, zero value otherwise.

### GetGeneratedOutputsOk

`func (o *RichTransaction) GetGeneratedOutputsOk() (*[]Output, bool)`

GetGeneratedOutputsOk returns a tuple with the GeneratedOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedOutputs

`func (o *RichTransaction) SetGeneratedOutputs(v []Output)`

SetGeneratedOutputs sets GeneratedOutputs field to given value.


### GetInputSignatures

`func (o *RichTransaction) GetInputSignatures() []string`

GetInputSignatures returns the InputSignatures field if non-nil, zero value otherwise.

### GetInputSignaturesOk

`func (o *RichTransaction) GetInputSignaturesOk() (*[]string, bool)`

GetInputSignaturesOk returns a tuple with the InputSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSignatures

`func (o *RichTransaction) SetInputSignatures(v []string)`

SetInputSignatures sets InputSignatures field to given value.


### GetScriptSignatures

`func (o *RichTransaction) GetScriptSignatures() []string`

GetScriptSignatures returns the ScriptSignatures field if non-nil, zero value otherwise.

### GetScriptSignaturesOk

`func (o *RichTransaction) GetScriptSignaturesOk() (*[]string, bool)`

GetScriptSignaturesOk returns a tuple with the ScriptSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSignatures

`func (o *RichTransaction) SetScriptSignatures(v []string)`

SetScriptSignatures sets ScriptSignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


