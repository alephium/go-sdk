# FunctionSig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Signature** | **string** |  | 
**ArgNames** | **[]string** |  | 
**ArgTypes** | **[]string** |  | 
**ReturnTypes** | **[]string** |  | 

## Methods

### NewFunctionSig

`func NewFunctionSig(name string, signature string, argNames []string, argTypes []string, returnTypes []string, ) *FunctionSig`

NewFunctionSig instantiates a new FunctionSig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSigWithDefaults

`func NewFunctionSigWithDefaults() *FunctionSig`

NewFunctionSigWithDefaults instantiates a new FunctionSig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionSig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionSig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionSig) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *FunctionSig) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *FunctionSig) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *FunctionSig) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetArgNames

`func (o *FunctionSig) GetArgNames() []string`

GetArgNames returns the ArgNames field if non-nil, zero value otherwise.

### GetArgNamesOk

`func (o *FunctionSig) GetArgNamesOk() (*[]string, bool)`

GetArgNamesOk returns a tuple with the ArgNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgNames

`func (o *FunctionSig) SetArgNames(v []string)`

SetArgNames sets ArgNames field to given value.


### GetArgTypes

`func (o *FunctionSig) GetArgTypes() []string`

GetArgTypes returns the ArgTypes field if non-nil, zero value otherwise.

### GetArgTypesOk

`func (o *FunctionSig) GetArgTypesOk() (*[]string, bool)`

GetArgTypesOk returns a tuple with the ArgTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgTypes

`func (o *FunctionSig) SetArgTypes(v []string)`

SetArgTypes sets ArgTypes field to given value.


### GetReturnTypes

`func (o *FunctionSig) GetReturnTypes() []string`

GetReturnTypes returns the ReturnTypes field if non-nil, zero value otherwise.

### GetReturnTypesOk

`func (o *FunctionSig) GetReturnTypesOk() (*[]string, bool)`

GetReturnTypesOk returns a tuple with the ReturnTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTypes

`func (o *FunctionSig) SetReturnTypes(v []string)`

SetReturnTypes sets ReturnTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


