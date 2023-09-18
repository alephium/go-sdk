# CompileContractResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Name** | **string** |  | 
**Bytecode** | **string** |  | 
**BytecodeDebugPatch** | **string** |  | 
**CodeHash** | **string** |  | 
**CodeHashDebug** | **string** |  | 
**Fields** | [**FieldsSig**](FieldsSig.md) |  | 
**Functions** | [**[]FunctionSig**](FunctionSig.md) |  | 
**Constants** | [**[]Constant**](Constant.md) |  | 
**Enums** | [**[]Enum**](Enum.md) |  | 
**Events** | [**[]EventSig**](EventSig.md) |  | 
**Warnings** | **[]string** |  | 
**StdInterfaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewCompileContractResult

`func NewCompileContractResult(version string, name string, bytecode string, bytecodeDebugPatch string, codeHash string, codeHashDebug string, fields FieldsSig, functions []FunctionSig, constants []Constant, enums []Enum, events []EventSig, warnings []string, ) *CompileContractResult`

NewCompileContractResult instantiates a new CompileContractResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompileContractResultWithDefaults

`func NewCompileContractResultWithDefaults() *CompileContractResult`

NewCompileContractResultWithDefaults instantiates a new CompileContractResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CompileContractResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CompileContractResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CompileContractResult) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *CompileContractResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompileContractResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompileContractResult) SetName(v string)`

SetName sets Name field to given value.


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


### GetBytecodeDebugPatch

`func (o *CompileContractResult) GetBytecodeDebugPatch() string`

GetBytecodeDebugPatch returns the BytecodeDebugPatch field if non-nil, zero value otherwise.

### GetBytecodeDebugPatchOk

`func (o *CompileContractResult) GetBytecodeDebugPatchOk() (*string, bool)`

GetBytecodeDebugPatchOk returns a tuple with the BytecodeDebugPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecodeDebugPatch

`func (o *CompileContractResult) SetBytecodeDebugPatch(v string)`

SetBytecodeDebugPatch sets BytecodeDebugPatch field to given value.


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


### GetCodeHashDebug

`func (o *CompileContractResult) GetCodeHashDebug() string`

GetCodeHashDebug returns the CodeHashDebug field if non-nil, zero value otherwise.

### GetCodeHashDebugOk

`func (o *CompileContractResult) GetCodeHashDebugOk() (*string, bool)`

GetCodeHashDebugOk returns a tuple with the CodeHashDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHashDebug

`func (o *CompileContractResult) SetCodeHashDebug(v string)`

SetCodeHashDebug sets CodeHashDebug field to given value.


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


### GetConstants

`func (o *CompileContractResult) GetConstants() []Constant`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *CompileContractResult) GetConstantsOk() (*[]Constant, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *CompileContractResult) SetConstants(v []Constant)`

SetConstants sets Constants field to given value.


### GetEnums

`func (o *CompileContractResult) GetEnums() []Enum`

GetEnums returns the Enums field if non-nil, zero value otherwise.

### GetEnumsOk

`func (o *CompileContractResult) GetEnumsOk() (*[]Enum, bool)`

GetEnumsOk returns a tuple with the Enums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnums

`func (o *CompileContractResult) SetEnums(v []Enum)`

SetEnums sets Enums field to given value.


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


### GetWarnings

`func (o *CompileContractResult) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CompileContractResult) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CompileContractResult) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.


### GetStdInterfaceId

`func (o *CompileContractResult) GetStdInterfaceId() string`

GetStdInterfaceId returns the StdInterfaceId field if non-nil, zero value otherwise.

### GetStdInterfaceIdOk

`func (o *CompileContractResult) GetStdInterfaceIdOk() (*string, bool)`

GetStdInterfaceIdOk returns a tuple with the StdInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdInterfaceId

`func (o *CompileContractResult) SetStdInterfaceId(v string)`

SetStdInterfaceId sets StdInterfaceId field to given value.

### HasStdInterfaceId

`func (o *CompileContractResult) HasStdInterfaceId() bool`

HasStdInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


