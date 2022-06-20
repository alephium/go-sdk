# TransactionTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unsigned** | [**UnsignedTx**](UnsignedTx.md) |  | 
**InputSignatures** | **[]string** |  | 
**ScriptSignatures** | **[]string** |  | 

## Methods

### NewTransactionTemplate

`func NewTransactionTemplate(unsigned UnsignedTx, inputSignatures []string, scriptSignatures []string, ) *TransactionTemplate`

NewTransactionTemplate instantiates a new TransactionTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTemplateWithDefaults

`func NewTransactionTemplateWithDefaults() *TransactionTemplate`

NewTransactionTemplateWithDefaults instantiates a new TransactionTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsigned

`func (o *TransactionTemplate) GetUnsigned() UnsignedTx`

GetUnsigned returns the Unsigned field if non-nil, zero value otherwise.

### GetUnsignedOk

`func (o *TransactionTemplate) GetUnsignedOk() (*UnsignedTx, bool)`

GetUnsignedOk returns a tuple with the Unsigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsigned

`func (o *TransactionTemplate) SetUnsigned(v UnsignedTx)`

SetUnsigned sets Unsigned field to given value.


### GetInputSignatures

`func (o *TransactionTemplate) GetInputSignatures() []string`

GetInputSignatures returns the InputSignatures field if non-nil, zero value otherwise.

### GetInputSignaturesOk

`func (o *TransactionTemplate) GetInputSignaturesOk() (*[]string, bool)`

GetInputSignaturesOk returns a tuple with the InputSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSignatures

`func (o *TransactionTemplate) SetInputSignatures(v []string)`

SetInputSignatures sets InputSignatures field to given value.


### GetScriptSignatures

`func (o *TransactionTemplate) GetScriptSignatures() []string`

GetScriptSignatures returns the ScriptSignatures field if non-nil, zero value otherwise.

### GetScriptSignaturesOk

`func (o *TransactionTemplate) GetScriptSignaturesOk() (*[]string, bool)`

GetScriptSignaturesOk returns a tuple with the ScriptSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSignatures

`func (o *TransactionTemplate) SetScriptSignatures(v []string)`

SetScriptSignatures sets ScriptSignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


