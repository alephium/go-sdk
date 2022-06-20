# CompileContractResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytecode** | **string** |  | 
**CodeHash** | **string** |  | 
**Fields** | [**FieldsSig**](FieldsSig.md) |  | 
**Functions** | [**[]FunctionSig**](FunctionSig.md) |  | 
**Events** | [**[]EventSig**](EventSig.md) |  | 

## Methods

### NewCompileContractResult

`func NewCompileContractResult(bytecode string, codeHash string, fields FieldsSig, functions []FunctionSig, events []EventSig, ) *CompileContractResult`

NewCompileContractResult instantiates a new CompileContractResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompileContractResultWithDefaults

`func NewCompileContractResultWithDefaults() *CompileContractResult`

NewCompileContractResultWithDefaults instantiates a new CompileContractResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytecode

`func (o *CompileContractResult) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *CompileContractResult) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *CompileContractResult) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetCodeHash

`func (o *CompileContractResult) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *CompileContractResult) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *CompileContractResult) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.


### GetFields

`func (o *CompileContractResult) GetFields() FieldsSig`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CompileContractResult) GetFieldsOk() (*FieldsSig, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CompileContractResult) SetFields(v FieldsSig)`

SetFields sets Fields field to given value.


### GetFunctions

`func (o *CompileContractResult) GetFunctions() []FunctionSig`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *CompileContractResult) GetFunctionsOk() (*[]FunctionSig, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *CompileContractResult) SetFunctions(v []FunctionSig)`

SetFunctions sets Functions field to given value.


### GetEvents

`func (o *CompileContractResult) GetEvents() []EventSig`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CompileContractResult) GetEventsOk() (*[]EventSig, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CompileContractResult) SetEvents(v []EventSig)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


