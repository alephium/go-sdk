# FunctionSig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**UsePreapprovedAssets** | **bool** |  | 
**UseAssetsInContract** | **bool** |  | 
**IsPublic** | **bool** |  | 
**ParamNames** | **[]string** |  | 
**ParamTypes** | **[]string** |  | 
**ParamIsMutable** | **[]bool** |  | 
**ReturnTypes** | **[]string** |  | 

## Methods

### NewFunctionSig

`func NewFunctionSig(name string, usePreapprovedAssets bool, useAssetsInContract bool, isPublic bool, paramNames []string, paramTypes []string, paramIsMutable []bool, returnTypes []string, ) *FunctionSig`

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


### GetUsePreapprovedAssets

`func (o *FunctionSig) GetUsePreapprovedAssets() bool`

GetUsePreapprovedAssets returns the UsePreapprovedAssets field if non-nil, zero value otherwise.

### GetUsePreapprovedAssetsOk

`func (o *FunctionSig) GetUsePreapprovedAssetsOk() (*bool, bool)`

GetUsePreapprovedAssetsOk returns a tuple with the UsePreapprovedAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreapprovedAssets

`func (o *FunctionSig) SetUsePreapprovedAssets(v bool)`

SetUsePreapprovedAssets sets UsePreapprovedAssets field to given value.


### GetUseAssetsInContract

`func (o *FunctionSig) GetUseAssetsInContract() bool`

GetUseAssetsInContract returns the UseAssetsInContract field if non-nil, zero value otherwise.

### GetUseAssetsInContractOk

`func (o *FunctionSig) GetUseAssetsInContractOk() (*bool, bool)`

GetUseAssetsInContractOk returns a tuple with the UseAssetsInContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAssetsInContract

`func (o *FunctionSig) SetUseAssetsInContract(v bool)`

SetUseAssetsInContract sets UseAssetsInContract field to given value.


### GetIsPublic

`func (o *FunctionSig) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *FunctionSig) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *FunctionSig) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetParamNames

`func (o *FunctionSig) GetParamNames() []string`

GetParamNames returns the ParamNames field if non-nil, zero value otherwise.

### GetParamNamesOk

`func (o *FunctionSig) GetParamNamesOk() (*[]string, bool)`

GetParamNamesOk returns a tuple with the ParamNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamNames

`func (o *FunctionSig) SetParamNames(v []string)`

SetParamNames sets ParamNames field to given value.


### GetParamTypes

`func (o *FunctionSig) GetParamTypes() []string`

GetParamTypes returns the ParamTypes field if non-nil, zero value otherwise.

### GetParamTypesOk

`func (o *FunctionSig) GetParamTypesOk() (*[]string, bool)`

GetParamTypesOk returns a tuple with the ParamTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamTypes

`func (o *FunctionSig) SetParamTypes(v []string)`

SetParamTypes sets ParamTypes field to given value.


### GetParamIsMutable

`func (o *FunctionSig) GetParamIsMutable() []bool`

GetParamIsMutable returns the ParamIsMutable field if non-nil, zero value otherwise.

### GetParamIsMutableOk

`func (o *FunctionSig) GetParamIsMutableOk() (*[]bool, bool)`

GetParamIsMutableOk returns a tuple with the ParamIsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamIsMutable

`func (o *FunctionSig) SetParamIsMutable(v []bool)`

SetParamIsMutable sets ParamIsMutable field to given value.


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


