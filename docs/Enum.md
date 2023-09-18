# Enum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Fields** | [**[]EnumField**](EnumField.md) |  | 

## Methods

### NewEnum

`func NewEnum(name string, fields []EnumField, ) *Enum`

NewEnum instantiates a new Enum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumWithDefaults

`func NewEnumWithDefaults() *Enum`

NewEnumWithDefaults instantiates a new Enum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Enum) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Enum) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Enum) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *Enum) GetFields() []EnumField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Enum) GetFieldsOk() (*[]EnumField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Enum) SetFields(v []EnumField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


