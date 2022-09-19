# CompileProjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | [**[]CompileContractResult**](CompileContractResult.md) |  | 
**Scripts** | [**[]CompileScriptResult**](CompileScriptResult.md) |  | 

## Methods

### NewCompileProjectResult

`func NewCompileProjectResult(contracts []CompileContractResult, scripts []CompileScriptResult, ) *CompileProjectResult`

NewCompileProjectResult instantiates a new CompileProjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompileProjectResultWithDefaults

`func NewCompileProjectResultWithDefaults() *CompileProjectResult`

NewCompileProjectResultWithDefaults instantiates a new CompileProjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContracts

`func (o *CompileProjectResult) GetContracts() []CompileContractResult`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *CompileProjectResult) GetContractsOk() (*[]CompileContractResult, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *CompileProjectResult) SetContracts(v []CompileContractResult)`

SetContracts sets Contracts field to given value.


### GetScripts

`func (o *CompileProjectResult) GetScripts() []CompileScriptResult`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *CompileProjectResult) GetScriptsOk() (*[]CompileScriptResult, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *CompileProjectResult) SetScripts(v []CompileScriptResult)`

SetScripts sets Scripts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


