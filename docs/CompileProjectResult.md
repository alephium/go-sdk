# CompileProjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | [**[]CompileContractResult**](CompileContractResult.md) |  | 
**Scripts** | [**[]CompileScriptResult**](CompileScriptResult.md) |  | 
**Structs** | Pointer to [**[]StructSig**](StructSig.md) |  | [optional] 
**Constants** | Pointer to [**[]Constant**](Constant.md) |  | [optional] 
**Enums** | Pointer to [**[]Enum**](Enum.md) |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**TestError** | Pointer to **string** |  | [optional] 

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


### GetStructs

`func (o *CompileProjectResult) GetStructs() []StructSig`

GetStructs returns the Structs field if non-nil, zero value otherwise.

### GetStructsOk

`func (o *CompileProjectResult) GetStructsOk() (*[]StructSig, bool)`

GetStructsOk returns a tuple with the Structs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructs

`func (o *CompileProjectResult) SetStructs(v []StructSig)`

SetStructs sets Structs field to given value.

### HasStructs

`func (o *CompileProjectResult) HasStructs() bool`

HasStructs returns a boolean if a field has been set.

### GetConstants

`func (o *CompileProjectResult) GetConstants() []Constant`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *CompileProjectResult) GetConstantsOk() (*[]Constant, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *CompileProjectResult) SetConstants(v []Constant)`

SetConstants sets Constants field to given value.

### HasConstants

`func (o *CompileProjectResult) HasConstants() bool`

HasConstants returns a boolean if a field has been set.

### GetEnums

`func (o *CompileProjectResult) GetEnums() []Enum`

GetEnums returns the Enums field if non-nil, zero value otherwise.

### GetEnumsOk

`func (o *CompileProjectResult) GetEnumsOk() (*[]Enum, bool)`

GetEnumsOk returns a tuple with the Enums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnums

`func (o *CompileProjectResult) SetEnums(v []Enum)`

SetEnums sets Enums field to given value.

### HasEnums

`func (o *CompileProjectResult) HasEnums() bool`

HasEnums returns a boolean if a field has been set.

### GetWarnings

`func (o *CompileProjectResult) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CompileProjectResult) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CompileProjectResult) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CompileProjectResult) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetTestError

`func (o *CompileProjectResult) GetTestError() string`

GetTestError returns the TestError field if non-nil, zero value otherwise.

### GetTestErrorOk

`func (o *CompileProjectResult) GetTestErrorOk() (*string, bool)`

GetTestErrorOk returns a tuple with the TestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestError

`func (o *CompileProjectResult) SetTestError(v string)`

SetTestError sets TestError field to given value.

### HasTestError

`func (o *CompileProjectResult) HasTestError() bool`

HasTestError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


