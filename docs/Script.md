# Script

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**CompilerOptions** | Pointer to [**CompilerOptions**](CompilerOptions.md) |  | [optional] 

## Methods

### NewScript

`func NewScript(code string, ) *Script`

NewScript instantiates a new Script object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWithDefaults

`func NewScriptWithDefaults() *Script`

NewScriptWithDefaults instantiates a new Script object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Script) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Script) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Script) SetCode(v string)`

SetCode sets Code field to given value.


### GetCompilerOptions

`func (o *Script) GetCompilerOptions() CompilerOptions`

GetCompilerOptions returns the CompilerOptions field if non-nil, zero value otherwise.

### GetCompilerOptionsOk

`func (o *Script) GetCompilerOptionsOk() (*CompilerOptions, bool)`

GetCompilerOptionsOk returns a tuple with the CompilerOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompilerOptions

`func (o *Script) SetCompilerOptions(v CompilerOptions)`

SetCompilerOptions sets CompilerOptions field to given value.

### HasCompilerOptions

`func (o *Script) HasCompilerOptions() bool`

HasCompilerOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


