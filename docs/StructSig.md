# StructSig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FieldNames** | **[]string** |  | 
**FieldTypes** | **[]string** |  | 
**IsMutable** | **[]bool** |  | 

## Methods

### NewStructSig

`func NewStructSig(name string, fieldNames []string, fieldTypes []string, isMutable []bool, ) *StructSig`

NewStructSig instantiates a new StructSig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructSigWithDefaults

`func NewStructSigWithDefaults() *StructSig`

NewStructSigWithDefaults instantiates a new StructSig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StructSig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StructSig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StructSig) SetName(v string)`

SetName sets Name field to given value.


### GetFieldNames

`func (o *StructSig) GetFieldNames() []string`

GetFieldNames returns the FieldNames field if non-nil, zero value otherwise.

### GetFieldNamesOk

`func (o *StructSig) GetFieldNamesOk() (*[]string, bool)`

GetFieldNamesOk returns a tuple with the FieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNames

`func (o *StructSig) SetFieldNames(v []string)`

SetFieldNames sets FieldNames field to given value.


### GetFieldTypes

`func (o *StructSig) GetFieldTypes() []string`

GetFieldTypes returns the FieldTypes field if non-nil, zero value otherwise.

### GetFieldTypesOk

`func (o *StructSig) GetFieldTypesOk() (*[]string, bool)`

GetFieldTypesOk returns a tuple with the FieldTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypes

`func (o *StructSig) SetFieldTypes(v []string)`

SetFieldTypes sets FieldTypes field to given value.


### GetIsMutable

`func (o *StructSig) GetIsMutable() []bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *StructSig) GetIsMutableOk() (*[]bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *StructSig) SetIsMutable(v []bool)`

SetIsMutable sets IsMutable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


