# CompileScriptResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Name** | **string** |  | 
**BytecodeTemplate** | **string** |  | 
**BytecodeDebugPatch** | **string** |  | 
**Fields** | [**FieldsSig**](FieldsSig.md) |  | 
**Functions** | [**[]FunctionSig**](FunctionSig.md) |  | 
**Warnings** | **[]string** |  | 

## Methods

### NewCompileScriptResult

`func NewCompileScriptResult(version string, name string, bytecodeTemplate string, bytecodeDebugPatch string, fields FieldsSig, functions []FunctionSig, warnings []string, ) *CompileScriptResult`

NewCompileScriptResult instantiates a new CompileScriptResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompileScriptResultWithDefaults

`func NewCompileScriptResultWithDefaults() *CompileScriptResult`

NewCompileScriptResultWithDefaults instantiates a new CompileScriptResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CompileScriptResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CompileScriptResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CompileScriptResult) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *CompileScriptResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompileScriptResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompileScriptResult) SetName(v string)`

SetName sets Name field to given value.


### GetBytecodeTemplate

`func (o *CompileScriptResult) GetBytecodeTemplate() string`

GetBytecodeTemplate returns the BytecodeTemplate field if non-nil, zero value otherwise.

### GetBytecodeTemplateOk

`func (o *CompileScriptResult) GetBytecodeTemplateOk() (*string, bool)`

GetBytecodeTemplateOk returns a tuple with the BytecodeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecodeTemplate

`func (o *CompileScriptResult) SetBytecodeTemplate(v string)`

SetBytecodeTemplate sets BytecodeTemplate field to given value.


### GetBytecodeDebugPatch

`func (o *CompileScriptResult) GetBytecodeDebugPatch() string`

GetBytecodeDebugPatch returns the BytecodeDebugPatch field if non-nil, zero value otherwise.

### GetBytecodeDebugPatchOk

`func (o *CompileScriptResult) GetBytecodeDebugPatchOk() (*string, bool)`

GetBytecodeDebugPatchOk returns a tuple with the BytecodeDebugPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecodeDebugPatch

`func (o *CompileScriptResult) SetBytecodeDebugPatch(v string)`

SetBytecodeDebugPatch sets BytecodeDebugPatch field to given value.


### GetFields

`func (o *CompileScriptResult) GetFields() FieldsSig`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CompileScriptResult) GetFieldsOk() (*FieldsSig, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CompileScriptResult) SetFields(v FieldsSig)`

SetFields sets Fields field to given value.


### GetFunctions

`func (o *CompileScriptResult) GetFunctions() []FunctionSig`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *CompileScriptResult) GetFunctionsOk() (*[]FunctionSig, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *CompileScriptResult) SetFunctions(v []FunctionSig)`

SetFunctions sets Functions field to given value.


### GetWarnings

`func (o *CompileScriptResult) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CompileScriptResult) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CompileScriptResult) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


