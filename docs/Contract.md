# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**CompilerOptions** | Pointer to [**CompilerOptions**](CompilerOptions.md) |  | [optional] 

## Methods

### NewContract

`func NewContract(code string, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Contract) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Contract) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Contract) SetCode(v string)`

SetCode sets Code field to given value.


### GetCompilerOptions

`func (o *Contract) GetCompilerOptions() CompilerOptions`

GetCompilerOptions returns the CompilerOptions field if non-nil, zero value otherwise.

### GetCompilerOptionsOk

`func (o *Contract) GetCompilerOptionsOk() (*CompilerOptions, bool)`

GetCompilerOptionsOk returns a tuple with the CompilerOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilerOptions

`func (o *Contract) SetCompilerOptions(v CompilerOptions)`

SetCompilerOptions sets CompilerOptions field to given value.

### HasCompilerOptions

`func (o *Contract) HasCompilerOptions() bool`

HasCompilerOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


